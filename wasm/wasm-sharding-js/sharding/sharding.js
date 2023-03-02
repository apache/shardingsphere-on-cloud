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

import * as c from './config.js'
import * as crc32 from './crc32.js'

function signed_crc_table() {
    var c = 0, table = new Array(256);

    for(var n =0; n != 256; ++n){
        c = n;
        c = ((c&1) ? (-306674912 ^ (c >>> 1)) : (c >>> 1));
        c = ((c&1) ? (-306674912 ^ (c >>> 1)) : (c >>> 1));
        c = ((c&1) ? (-306674912 ^ (c >>> 1)) : (c >>> 1));
        c = ((c&1) ? (-306674912 ^ (c >>> 1)) : (c >>> 1));
        c = ((c&1) ? (-306674912 ^ (c >>> 1)) : (c >>> 1));
        c = ((c&1) ? (-306674912 ^ (c >>> 1)) : (c >>> 1));
        c = ((c&1) ? (-306674912 ^ (c >>> 1)) : (c >>> 1));
        c = ((c&1) ? (-306674912 ^ (c >>> 1)) : (c >>> 1));
        table[n] = c;
    }

    return typeof Int32Array !== 'undefined' ? new Int32Array(table) : table;
}

var T0 = signed_crc_table();

function crc32_str(str, seed) {
    var C = seed ^ -1;
    for(var i = 0, L = str.length, c = 0, d = 0; i < L;) {
        c = str.charCodeAt(i++);
        if(c < 0x80) {
            C = (C>>>8) ^ T0[(C^c)&0xFF];
        } else if(c < 0x800) {
            C = (C>>>8) ^ T0[(C ^ (192|((c>>6)&31)))&0xFF];
            C = (C>>>8) ^ T0[(C ^ (128|(c&63)))&0xFF];
        } else if(c >= 0xD800 && c < 0xE000) {
            c = (c&1023)+64; d = str.charCodeAt(i++)&1023;
            C = (C>>>8) ^ T0[(C ^ (240|((c>>8)&7)))&0xFF];
            C = (C>>>8) ^ T0[(C ^ (128|((c>>2)&63)))&0xFF];
            C = (C>>>8) ^ T0[(C ^ (128|((d>>6)&15)|((c&3)<<4)))&0xFF];
            C = (C>>>8) ^ T0[(C ^ (128|(d&63)))&0xFF];
        } else {
            C = (C>>>8) ^ T0[(C ^ (224|((c>>12)&15)))&0xFF];
            C = (C>>>8) ^ T0[(C ^ (128|((c>>6)&63)))&0xFF];
            C = (C>>>8) ^ T0[(C ^ (128|(c&63)))&0xFF];
        }
    }
    return ~C;
}

function crc32_buf(B, seed) {
    var C = seed ^ -1, L = B.length - 15, i = 0;
    for(; i < L;) C =
        Tf[B[i++] ^ (C & 255)] ^
        Te[B[i++] ^ ((C >> 8) & 255)] ^
        Td[B[i++] ^ ((C >> 16) & 255)] ^
        Tc[B[i++] ^ (C >>> 24)] ^
        Tb[B[i++]] ^ Ta[B[i++]] ^ T9[B[i++]] ^ T8[B[i++]] ^
        T7[B[i++]] ^ T6[B[i++]] ^ T5[B[i++]] ^ T4[B[i++]] ^
        T3[B[i++]] ^ T2[B[i++]] ^ T1[B[i++]] ^ T0[B[i++]];
    L += 15;
    while(i < L) C = (C>>>8) ^ T0[(C^B[i++])&0xFF];
    return ~C;
}

function crc32_bstr(bstr, seed) {
    var C = seed ^ -1;
    for(var i = 0, L = bstr.length; i < L;) C = (C>>>8) ^ T0[(C^bstr.charCodeAt(i++))&0xFF];
    return ~C;
}

function textToArrayBuffer(s) {
  var i = s.length;
  var n = 0;
  var ba = new Array()
  for (var j = 0; j < i;) {
    var c = s.codePointAt(j);
    if (c < 128) {
      ba[n++] = c;
      j++;
    }
    else if ((c > 127) && (c < 2048)) {
      ba[n++] = (c >> 6) | 192;
      ba[n++] = (c & 63) | 128;
      j++;
    }
    else if ((c > 2047) && (c < 65536)) {
      ba[n++] = (c >> 12) | 224;
      ba[n++] = ((c >> 6) & 63) | 128;
      ba[n++] = (c & 63) | 128;
      j++;
    }
    else {
      ba[n++] = (c >> 18) | 240;
      ba[n++] = ((c >> 12) & 63) | 128;
      ba[n++] = ((c >> 6) & 63) | 128;
      ba[n++] = (c & 63) | 128;
      j+=2;
    }
  }
  return ba
}


function pisa_crc32(str, mod) {
    let c2 = crc32_str(str)
    let m = c2 % mod
    return m < 256?0:m<512?1:m<768?2:3
}

let re = /(\w+)\$\{(\w+)\((\w+)\)\}/
let cc = c.cc_crc32.match(re)

let fileds = ["id", "user_name", "address", "password", "email"]

var expresses_crc32 = new Array();
for (var i = 0; i < 5; i++) {
    if (cc[2] == "crc32") {
        let m = crc32_buf(textToArrayBuffer(fileds[i]))
        let h = m % 1024
        expresses_crc32[i] = cc[1] + pisa_crc32(fileds[i], 1024)
    }
}

print(c.cc_crc32, ' => [' + expresses_crc32 + ']')
