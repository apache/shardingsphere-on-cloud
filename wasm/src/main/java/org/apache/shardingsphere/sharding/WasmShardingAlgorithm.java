/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package org.apache.shardingsphere.sharding;

import com.google.common.base.Strings;
import io.github.kawamuray.wasmtime.*;
import io.github.kawamuray.wasmtime.wasi.WasiCtx;
import io.github.kawamuray.wasmtime.wasi.WasiCtxBuilder;
import static io.github.kawamuray.wasmtime.WasmValType.I32;
import static io.github.kawamuray.wasmtime.WasmValType.I64;
import io.github.kawamuray.wasmtime.Store;
import lombok.Getter;
import org.apache.shardingsphere.infra.datanode.DataNodeInfo;
import org.apache.shardingsphere.sharding.api.sharding.standard.PreciseShardingValue;
import org.apache.shardingsphere.sharding.api.sharding.standard.RangeShardingValue;
import org.apache.shardingsphere.sharding.api.sharding.standard.StandardShardingAlgorithm;

import java.math.BigInteger;
import java.nio.ByteBuffer;
import java.util.*;
import java.util.concurrent.atomic.AtomicReference;

/**
 * Modulo sharding algorithm.
 */
public final class WasmShardingAlgorithm implements StandardShardingAlgorithm<Comparable<?>> {

    @Getter
    private Properties props;

    private static final String WASM_PATH = "./wasm-sharding/target/wasm32-wasi/debug/wasm_sharding.wasm";

    @Override
    public void init(final Properties props) {
        this.props = props;
    }

    @Override
    public String doSharding(final Collection<String> availableTargetNames, final PreciseShardingValue<Comparable<?>> shardingValue) {
        return wasmDoSharding(availableTargetNames, shardingValue);
    }

    @Override
    public Collection<String> doSharding(final Collection<String> availableTargetNames, final RangeShardingValue<Comparable<?>> shardingValue) {
        return Collections.EMPTY_LIST;
    }

    private String wasmDoSharding(final Collection<String> availableTargetNames, final PreciseShardingValue<Comparable<?>> shardingValue) {
        String availableTargetNamesArgs = String.join(",", availableTargetNames);
        String conditionValuesArgs = String.format("%s,%s,%s", shardingValue.getColumnName(), shardingValue.getLogicTableName(), shardingValue.getValue());
        AtomicReference<Memory> memRef = new AtomicReference<>();
        final Integer[] offset = {1};
        try (WasiCtx wasi = new WasiCtxBuilder().inheritStdout().inheritStderr().build();
             Store<Void> store = Store.withoutData(wasi);
             Linker linker = new Linker(store.engine());
             Func pollWordFn = WasmFunctions.wrap(store, I64, I32, I32, (addr, len) -> {
                 ByteBuffer buf = memRef.get().buffer(store);
                 buf.put(addr.intValue(), (byte) availableTargetNamesArgs.length());

                 for (int j = 0; j < availableTargetNamesArgs.length(); j++) {
                    buf.put(addr.intValue() + j + offset[0], (byte) availableTargetNamesArgs.charAt(j));
                 }

                 offset[0] += availableTargetNamesArgs.length();

                 buf.put(addr.intValue() + offset[0], (byte) conditionValuesArgs.length());
                 offset[0] += 1;
                 for (int j = 0; j < conditionValuesArgs.length(); j++) {
                    buf.put(addr.intValue() + j + offset[0], (byte) conditionValuesArgs.charAt(j));
                 }

                 offset[0] += conditionValuesArgs.length();
                 return Math.min(offset[0], len);
             });
             Module module = Module.fromFile(store.engine(), WASM_PATH)) {

            WasiCtx.addToLinker(linker);
            linker.define("sharding", "poll_table", Extern.fromFunc(pollWordFn));
            linker.module(store, "", module);


            try (Memory mem = linker.get(store, "", "memory").get().memory();
                 Func doWorkFn = linker.get(store, "", "do_work").get().func()) {
                WasmFunctions.Function0<Long> doWork = WasmFunctions.func(store, doWorkFn,  I64);

                memRef.set(mem);

                ByteBuffer buf = mem.buffer(store);

                int ptr = doWork.call().intValue();
                buf.position(ptr + offset[0]);
                StringBuilder result = new StringBuilder(10);

                while (true) {
                    Byte byt = buf.get();
                    char a = (char) byt.byteValue();
                    if (byt.intValue() == 0)  {
                        break;
                    }

                    result.append(a);
                }

                return result.toString();
            }
        }
    }

    @Override
    public String getType() {
        return "WASM";
    }
}


