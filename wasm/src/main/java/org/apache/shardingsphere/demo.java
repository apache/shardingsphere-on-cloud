package org.apache.shardingsphere;

import org.apache.shardingsphere.infra.config.props.ConfigurationProperties;
import org.apache.shardingsphere.infra.datanode.DataNodeInfo;
import org.apache.shardingsphere.sharding.WasmShardingAlgorithm;
import org.apache.shardingsphere.sharding.api.sharding.standard.StandardShardingAlgorithm;
import org.apache.shardingsphere.sharding.route.engine.condition.value.ListShardingConditionValue;
import org.apache.shardingsphere.sharding.route.engine.condition.value.ShardingConditionValue;
import org.apache.shardingsphere.sharding.route.strategy.ShardingStrategy;
import org.apache.shardingsphere.sharding.route.strategy.type.standard.StandardShardingStrategy;

import java.util.*;

public class demo {
    public static void main(String[] args) {
        String shardingColumn = "id";
        StandardShardingAlgorithm<?> shardingAlgorithm = new WasmShardingAlgorithm();
        Properties props = new Properties();
        props.setProperty("sharding-count", "3");
        shardingAlgorithm.init(props);
        ShardingStrategy strategy = new StandardShardingStrategy(shardingColumn, shardingAlgorithm);

        Collection<String> targetNames = new ArrayList<>();
        targetNames.add("t_order_0");
        targetNames.add("t_order_1");
        targetNames.add("t_order_2");

        DataNodeInfo dataNodeInfo = new DataNodeInfo("t_order_", 1, '0');

        Collection<ShardingConditionValue> conditionValues = Collections.singletonList(
                new ListShardingConditionValue<>("id", "t_order", Collections.singletonList(10)));

        Collection<String> result = strategy.doSharding(targetNames, conditionValues, dataNodeInfo, new ConfigurationProperties(props));
        System.out.println(result);
    }
}
