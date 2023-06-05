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

package ast

import (
	"fmt"
	"strings"
)

type CreateShardingTableRule struct {
	IfNotExists                    *IfNotExists
	AllShardingTableRuleDefinition []*ShardingTableRuleDefinition
}

func (createShardingTableRule *CreateShardingTableRule) ToString() string {
	var (
		ifNotExists                    string
		allShardingTableRuleDefinition []string
	)
	if createShardingTableRule.IfNotExists != nil {
		ifNotExists = createShardingTableRule.IfNotExists.ToString()
	}
	if createShardingTableRule.AllShardingTableRuleDefinition != nil {
		for _, r := range createShardingTableRule.AllShardingTableRuleDefinition {
			allShardingTableRuleDefinition = append(allShardingTableRuleDefinition, r.ToString())
		}
	}
	return fmt.Sprintf("CREATE SHARDING TABLE RULE %s %s", ifNotExists, strings.Join(allShardingTableRuleDefinition, ","))
}

type ShardingTableRuleDefinition struct {
	ShardingAutoTableRule *ShardingAutoTableRule
	ShardingTableRule     *ShardingTableRule
}

func (shardingTableRuleDefinition *ShardingTableRuleDefinition) ToString() string {
	switch {
	case shardingTableRuleDefinition.ShardingAutoTableRule != nil:
		return shardingTableRuleDefinition.ShardingAutoTableRule.ToString()
	case shardingTableRuleDefinition.ShardingTableRule != nil:
		return shardingTableRuleDefinition.ShardingTableRule.ToString()
	}
	return ""
}

type CreateShardingTableReferenceRule struct {
	IfNotExists                     *IfNotExists
	AllTableReferenceRuleDefinition []*TableReferenceRuleDefinition
}

func (createShardingTableReferenceRule *CreateShardingTableReferenceRule) ToString() string {
	var (
		ifNotExists                     string
		allTableReferenceRuleDefinition []string
	)
	if createShardingTableReferenceRule.IfNotExists != nil {
		ifNotExists = createShardingTableReferenceRule.IfNotExists.ToString()
	}
	if createShardingTableReferenceRule.AllTableReferenceRuleDefinition != nil {
		for _, r := range createShardingTableReferenceRule.AllTableReferenceRuleDefinition {
			allTableReferenceRuleDefinition = append(allTableReferenceRuleDefinition, r.ToString())
		}
	}
	return fmt.Sprintf("CREATE SHARDING TABLE REFERENCE RULE %s %s", ifNotExists, strings.Join(allTableReferenceRuleDefinition, ","))
}

type TableReferenceRuleDefinition struct {
	RuleName     *CommonIdentifier
	AllTableName []*CommonIdentifier
}

func (tableReferenceRuleDefinition *TableReferenceRuleDefinition) ToString() string {
	var (
		ruleName     string
		allTableName []string
	)
	if tableReferenceRuleDefinition.RuleName != nil {
		ruleName = tableReferenceRuleDefinition.RuleName.ToString()
	}
	if tableReferenceRuleDefinition.AllTableName != nil {
		for _, t := range tableReferenceRuleDefinition.AllTableName {
			allTableName = append(allTableName, t.ToString())
		}
	}
	return fmt.Sprintf("%s (%s)", ruleName, strings.Join(allTableName, ","))
}

type ShardingAutoTableRule struct {
	TableName                    *CommonIdentifier
	StorageUnits                 *StorageUnits
	AutoShardingColumnDefinition *AutoShardingColumnDefinition
	AlgorithmDefinition          *ShardingAlgorithmDefinition
	KeyGenerateDefinition        *KeyGenerateDefinition
	AuditDefinition              *AuditDefinition
}

func (shardingAutoTableRule *ShardingAutoTableRule) ToString() string {
	var (
		tableName                    string
		storageUnits                 string
		autoShardingColumnDefinition string
		algorithmDefinition          string
		keyGenerateDefinition        string
		auditDefinition              string
	)

	if shardingAutoTableRule.TableName != nil {
		tableName = shardingAutoTableRule.TableName.ToString()
	}
	if shardingAutoTableRule.StorageUnits != nil {
		storageUnits = shardingAutoTableRule.StorageUnits.ToString()
	}
	if shardingAutoTableRule.AutoShardingColumnDefinition != nil {
		autoShardingColumnDefinition = shardingAutoTableRule.AutoShardingColumnDefinition.ToString()
	}
	if shardingAutoTableRule.AlgorithmDefinition != nil {
		algorithmDefinition = shardingAutoTableRule.AlgorithmDefinition.ToString()
	}
	if shardingAutoTableRule.KeyGenerateDefinition != nil {
		keyGenerateDefinition = shardingAutoTableRule.KeyGenerateDefinition.ToString()
	}
	if shardingAutoTableRule.AuditDefinition != nil {
		auditDefinition = shardingAutoTableRule.AuditDefinition.ToString()
	}
	return fmt.Sprintf("%s (%s,%s,%s,%s,%s)", tableName, storageUnits, autoShardingColumnDefinition, algorithmDefinition, keyGenerateDefinition, auditDefinition)
}

type StorageUnits struct {
	AllStorageUnit []*StorageUnit
}

func (storageUnits *StorageUnits) ToString() string {
	var (
		allStorageUnit []string
	)
	if storageUnits.AllStorageUnit != nil {
		for _, s := range storageUnits.AllStorageUnit {
			allStorageUnit = append(allStorageUnit, s.ToString())
		}
	}
	return fmt.Sprintf("STORAGE_UNITS (%s)", strings.Join(allStorageUnit, ","))
}

type StorageUnit struct {
	String     string
	Identifier string
}

func (storageUnit *StorageUnit) ToString() string {
	switch {
	case storageUnit.String != "":
		return storageUnit.String
	case storageUnit.Identifier != "":
		return storageUnit.Identifier
	}
	return ""
}

type AutoShardingColumnDefinition struct {
	ShardingColumn *ShardingColumn
}

func (autoShardingColumnDefinition *AutoShardingColumnDefinition) ToString() string {
	return autoShardingColumnDefinition.ShardingColumn.ToString()
}

type ShardingColumn struct {
	ColumnName *CommonIdentifier
}

func (shardingColumn *ShardingColumn) ToString() string {
	return fmt.Sprintf("SHARDING_COLUMN EQ_ %s", shardingColumn.ColumnName.ToString())
}

type KeyGenerateDefinition struct {
	ColumnName          *CommonIdentifier
	AlgorithmDefinition *ShardingAlgorithmDefinition
}

func (keyGenerateDefinition *KeyGenerateDefinition) ToString() string {
	var (
		columnName          string
		algorithmDefinition string
	)
	if keyGenerateDefinition.ColumnName != nil {
		columnName = keyGenerateDefinition.ColumnName.ToString()
	}
	if keyGenerateDefinition.AlgorithmDefinition != nil {
		algorithmDefinition = keyGenerateDefinition.AlgorithmDefinition.ToString()
	}
	return fmt.Sprintf("KEY_GENERATE_STRATEGY (COLUMN = %s, %s)", columnName, algorithmDefinition)
}

type AuditDefinition struct {
	MultiAuditDefinition  *MultiAuditDefinition
	AuditAllowHintDisable *AuditAllowHintDisable
}

func (auditDefinition *AuditDefinition) ToString() string {
	return fmt.Sprintf("AUDIT_STRATEGY (%s, ALLOW_HINT_DISABLE = %s)",
		auditDefinition.MultiAuditDefinition.ToString(),
		auditDefinition.AuditAllowHintDisable.ToString())
}

type MultiAuditDefinition struct {
	AllSingleAuditDefinition []*SingleAuditDefinition
}

func (multiAuditDefinition *MultiAuditDefinition) ToString() string {
	var allSingleAuditDefinition []string
	for _, d := range multiAuditDefinition.AllSingleAuditDefinition {
		allSingleAuditDefinition = append(allSingleAuditDefinition, d.ToString())
	}
	return strings.Join(allSingleAuditDefinition, ",")
}

type SingleAuditDefinition struct {
	AlgorithmDefinition *ShardingAlgorithmDefinition
}

func (singleAuditDefinition *SingleAuditDefinition) ToString() string {
	return singleAuditDefinition.AlgorithmDefinition.ToString()
}

type AuditAllowHintDisable struct {
	AuditAllowHintDisable string
}

func (auditAllowHintDisable *AuditAllowHintDisable) ToString() string {
	return auditAllowHintDisable.AuditAllowHintDisable
}

type ShardingTableRule struct {
	TableName             *CommonIdentifier
	DataNodes             *DataNodes
	DatabaseStrategy      *DatabaseStrategy
	TableStrategy         *TableStrategy
	KeyGenerateDefinition *KeyGenerateDefinition
	AuditDefinition       *AuditDefinition
}

func (shardingTableRule *ShardingTableRule) ToString() string {
	var (
		tableName             string
		datanodes             string
		databaseStrategy      string
		tableStrategy         string
		keyGenerateDefinition string
		auditDefinition       string
	)
	if shardingTableRule.TableName != nil {
		tableName = shardingTableRule.TableName.ToString()
	}
	if shardingTableRule.DataNodes != nil {
		datanodes = shardingTableRule.DataNodes.ToString()
	}
	if shardingTableRule.DatabaseStrategy != nil {
		databaseStrategy = shardingTableRule.DatabaseStrategy.ToString()
	}
	if shardingTableRule.TableStrategy != nil {
		tableStrategy = shardingTableRule.TableStrategy.ToString()
	}
	if shardingTableRule.KeyGenerateDefinition != nil {
		keyGenerateDefinition = shardingTableRule.KeyGenerateDefinition.ToString()
	}
	if shardingTableRule.AuditDefinition != nil {
		auditDefinition = shardingTableRule.AuditDefinition.ToString()
	}
	return fmt.Sprintf("%s(%s,%s,%s,%s,%s)", tableName, datanodes, databaseStrategy, tableStrategy, keyGenerateDefinition, auditDefinition)
}

type DataNode struct {
	String string
}

func (dataNode *DataNode) ToString() string {
	return dataNode.String
}

type DatabaseStrategy struct {
	ShardingStrategy *ShardingStrategy
}

func (databaseStrategy *DatabaseStrategy) ToString() string {
	return fmt.Sprintf("DATABASE_STRATEGY (%s)", databaseStrategy.ShardingStrategy.ToString())
}

type TableStrategy struct {
	ShardingStrategy *ShardingStrategy
}

func (tableStrategy *TableStrategy) ToString() string {
	return fmt.Sprintf("TABLE_STRATEGY (%s)", tableStrategy.ShardingStrategy.ToString())
}

type ShardingStrategy struct {
	StrategyType             *StrategyType
	ShardingColumnDefinition *ShardingColumnDefinition
	ShardingAlgorithm        *ShardingAlgorithm
}

func (shardingStrategy *ShardingStrategy) ToString() string {
	var (
		strategyType             string
		ShardingColumnDefinition string
		ShardingAlgorithm        string
	)
	if shardingStrategy.StrategyType != nil {
		strategyType = shardingStrategy.StrategyType.ToString()
	}
	if shardingStrategy.ShardingColumnDefinition != nil {
		ShardingColumnDefinition = shardingStrategy.ShardingColumnDefinition.ToString()
	}
	if shardingStrategy.ShardingAlgorithm != nil {
		ShardingAlgorithm = shardingStrategy.ShardingAlgorithm.ToString()
	}

	return fmt.Sprintf("TYPE = %s,%s,%s", strategyType, ShardingColumnDefinition, ShardingAlgorithm)
}

type StrategyType struct {
	BuildInStrategyType *BuildInStrategyType
	String              string
}

func (strategyType *StrategyType) ToString() string {
	switch {
	case strategyType.String != "":
		return strategyType.String
	case strategyType.BuildInStrategyType != nil:
		return strategyType.BuildInStrategyType.ToString()
	}
	return ""
}

type ShardingColumnDefinition struct {
	ShardingColumn  *ShardingColumn
	ShardingColumns *ShardingColumns
}

func (shardingColumnDefinition *ShardingColumnDefinition) ToString() string {
	switch {
	case shardingColumnDefinition.ShardingColumn != nil:
		return shardingColumnDefinition.ShardingColumn.ToString()
	case shardingColumnDefinition.ShardingColumns != nil:
		return shardingColumnDefinition.ShardingColumns.ToString()
	}
	return ""
}

type ShardingColumns struct {
	ColumnName    *CommonIdentifier
	AllColumnName []*CommonIdentifier
}

func (shardingColumns *ShardingColumns) ToString() string {
	var (
		columnName    string
		allColumnName []string
	)
	if shardingColumns.ColumnName != nil {
		columnName = shardingColumns.ColumnName.ToString()
	}
	if shardingColumns.AllColumnName != nil {
		for _, n := range shardingColumns.AllColumnName {
			allColumnName = append(allColumnName, n.ToString())
		}
	}
	return fmt.Sprintf("SHARDING_COLUMNS = %s,%s", columnName, strings.Join(allColumnName, ","))
}

type ShardingAlgorithm struct {
	AlgorithmDefinition *ShardingAlgorithmDefinition
}

func (shardingAlgorithm *ShardingAlgorithm) ToString() string {
	return fmt.Sprintf("SHARDING_ALGORITHM (%s)", shardingAlgorithm.AlgorithmDefinition.ToString())
}

type AlterShardingTableRule struct {
	AllShardingTableRuleDefinition []*ShardingTableRuleDefinition
}

func (alterShardingTableRule *AlterShardingTableRule) ToString() string {
	var allRule []string
	for _, r := range alterShardingTableRule.AllShardingTableRuleDefinition {
		allRule = append(allRule, r.ToString())
	}
	return fmt.Sprintf("ALTER SHARDING TABLE RULE %s", strings.Join(allRule, ","))
}

type AlterShardingTableReferenceRule struct {
	AllTableReferenceRuleDefinition []*TableReferenceRuleDefinition
}

func (alterShardingTableReferenceRule *AlterShardingTableReferenceRule) ToString() string {
	var allRule []string
	for _, r := range alterShardingTableReferenceRule.AllTableReferenceRuleDefinition {
		allRule = append(allRule, r.ToString())
	}
	return fmt.Sprintf("ALTER SHARDING TABLE REFERENCE RULE %s", strings.Join(allRule, ","))
}

type DropShardingTableReferenceRule struct {
	IfExists     *IfExists
	AllRuleNames []*CommonIdentifier
}

func (dropShardingTableReferenceRule *DropShardingTableReferenceRule) ToString() string {
	var (
		ifExists string
		allRule  []string
	)

	if dropShardingTableReferenceRule.IfExists != nil {
		ifExists = dropShardingTableReferenceRule.IfExists.ToString()
	}
	if dropShardingTableReferenceRule.AllRuleNames != nil {
		for _, r := range dropShardingTableReferenceRule.AllRuleNames {
			allRule = append(allRule, r.ToString())
		}
	}
	return fmt.Sprintf("DROP SHARDING TABLE REFERENCE RULE %s %s", ifExists, strings.Join(allRule, ","))
}

type DropShardingTableRule struct {
	IfExists     *IfExists
	AllTableName []*CommonIdentifier
}

func (dropShardingTableRule *DropShardingTableRule) ToString() string {
	var (
		ifExists     string
		allTableName []string
	)
	if dropShardingTableRule.IfExists != nil {
		ifExists = dropShardingTableRule.IfExists.ToString()
	}
	if dropShardingTableRule.AllTableName != nil {
		for _, t := range dropShardingTableRule.AllTableName {
			allTableName = append(allTableName, t.ToString())
		}
	}
	return fmt.Sprintf("DROP SHARDING TABLE RULE %s %s", ifExists, strings.Join(allTableName, ","))
}

type CreateBroadcastTableRule struct {
	IfNotExists  *IfNotExists
	AllTableName []*CommonIdentifier
}

func (createBroadcastTableRule *CreateBroadcastTableRule) ToString() string {
	var (
		ifNotExists  string
		allTableName []string
	)
	if createBroadcastTableRule.IfNotExists != nil {
		ifNotExists = createBroadcastTableRule.IfNotExists.ToString()
	}
	if createBroadcastTableRule.AllTableName != nil {
		for _, t := range createBroadcastTableRule.AllTableName {
			allTableName = append(allTableName, t.ToString())
		}
	}
	return fmt.Sprintf("CREATE BROADCAST TABLE RULE %s %s", ifNotExists, strings.Join(allTableName, ","))
}

type DropBroadcastTableRule struct {
	IfExists     *IfExists
	AllTableName []*CommonIdentifier
}

func (dropBroadcastTableRule *DropBroadcastTableRule) ToString() string {
	var (
		ifExists string
		allTable []string
	)
	if dropBroadcastTableRule.IfExists != nil {
		ifExists = dropBroadcastTableRule.IfExists.ToString()
	}
	if dropBroadcastTableRule.AllTableName != nil {
		for _, t := range dropBroadcastTableRule.AllTableName {
			allTable = append(allTable, t.ToString())
		}
	}
	return fmt.Sprintf("DROP BROADCAST TABLE RULE %s %s", ifExists, strings.Join(allTable,","))
}

type DropShardingAlgorithm struct {
	IfExists                 *IfExists
	AllShardingAlgorithmName []*CommonIdentifier
}

func (dropShardingAlgorithm DropShardingAlgorithm) ToString() string {
	var (
		IfExists string
		allAlgo  []string
	)
	if dropShardingAlgorithm.IfExists != nil {
		IfExists = dropShardingAlgorithm.IfExists.ToString()
	}
	if dropShardingAlgorithm.AllShardingAlgorithmName != nil {
		for _, t := range dropShardingAlgorithm.AllShardingAlgorithmName {
			allAlgo = append(allAlgo, t.ToString())
		}
	}
	return fmt.Sprintf("DROP BROADCAST TABLE RULE %s %s", IfExists, strings.Join(allAlgo, ","))
}

type CreateDefaultShardingStrategy struct {
	IfNotExists      *IfNotExists
	ShardingStrategy *ShardingStrategy
}

func (createDefaultShardingStrategy *CreateDefaultShardingStrategy) ToString() string {
	var (
		ifNotExists      string
		shardingStrategy string
	)
	if createDefaultShardingStrategy.IfNotExists != nil {
		ifNotExists = createDefaultShardingStrategy.IfNotExists.ToString()
	}
	if createDefaultShardingStrategy.ShardingStrategy != nil {
		shardingStrategy = createDefaultShardingStrategy.ShardingStrategy.ToString()
	}
	return fmt.Sprintf("CREATE DEFAULT SHARDING type=(DATABASE | TABLE) STRATEGY %s(%s)", ifNotExists, shardingStrategy)
}

type BuildInStrategyType struct {
	BuildInStrategyType string
}

func (buildInStrategyType *BuildInStrategyType) ToString() string {
	return buildInStrategyType.BuildInStrategyType
}

type ShardingAlgorithmDefinition struct {
	ShardingAlgorithmTypeName *ShardingAlgorithmTypeName
	PropertiesDefinition      *PropertiesDefinition
}

func (shardingAlgorithmDefinition *ShardingAlgorithmDefinition) ToString() string {
	var (
		shardingAlgorithmTypeName string
		propertiesDefinition      string
	)
	if shardingAlgorithmDefinition.ShardingAlgorithmTypeName != nil {
		shardingAlgorithmTypeName = shardingAlgorithmDefinition.ShardingAlgorithmTypeName.ToString()
	}
	if shardingAlgorithmDefinition.PropertiesDefinition != nil {
		propertiesDefinition = shardingAlgorithmDefinition.PropertiesDefinition.ToString()
	}
	return fmt.Sprintf("TYPE ( NAME = %s ,%s)", shardingAlgorithmTypeName, propertiesDefinition)
}

type ShardingAlgorithmTypeName struct {
	String                            string
	BuildInShardingAlgorithmType      *BuildInCommon
	BuildInKeyGenerateAlgorithmType   *BuildInCommon
	BuildInShardingAuditAlgorithmType *BuildInCommon
}

func (shardingAlgorithmTypeName *ShardingAlgorithmTypeName) ToString() string {
	switch {
	case shardingAlgorithmTypeName.String != "":
		return shardingAlgorithmTypeName.String
	case shardingAlgorithmTypeName.BuildInShardingAlgorithmType != nil:
		return shardingAlgorithmTypeName.BuildInShardingAlgorithmType.ToString()
	case shardingAlgorithmTypeName.BuildInKeyGenerateAlgorithmType != nil :
		return shardingAlgorithmTypeName.BuildInKeyGenerateAlgorithmType.ToString()
	case shardingAlgorithmTypeName.BuildInShardingAuditAlgorithmType != nil:
		return shardingAlgorithmTypeName.BuildInShardingAuditAlgorithmType.ToString()
	}
}

type BuildInCommon struct {
	String string
}

func (buildInCommon *BuildInCommon) ToString() string {
	return buildInCommon.String
}

type DropDefaultShardingStrategy struct {
	IfExists *IfExists
}

func (dropDefaultShardingStrategy *DropDefaultShardingStrategy) ToString() string {
	return fmt.Sprintf("DROP DEFAULT SHARDING type=(DATABASE | TABLE) STRATEGY %s", dropDefaultShardingStrategy.IfExists.ToString())
}

type DropShardingKeyGenerator struct {
	IfExists            *IfExists
	AllKeyGeneratorName []*CommonIdentifier
}

func (DropShardingKeyGenerator *DropShardingKeyGenerator) ToString() string {
	var (
		ifExists string
		allKey []string
	)
	if DropShardingKeyGenerator.IfExists != nil {
		ifExists = DropShardingKeyGenerator.IfExists.ToString()
	}
	if DropShardingKeyGenerator.AllKeyGeneratorName != nil {
		for _, k := range DropShardingKeyGenerator.AllKeyGeneratorName {
			allKey = append(allKey, k.ToString())
		}
	}
	return fmt.Sprintf("DROP SHARDING KEY GENERATOR %s %s", ifExists,strings.Join(allKey,","))
}

type DropShardingAuditor struct {
	IfExists       *IfExists
	AllAuditorName []*CommonIdentifier
}

func (dropShardingAuditor *DropShardingAuditor) ToString() string {
	var (
		ifExists string
		allName  []string
	)
	if dropShardingAuditor.IfExists != nil {
		ifExists = dropShardingAuditor.IfExists.ToString()
	}
	if dropShardingAuditor.AllAuditorName != nil {
		for _, r := range dropShardingAuditor.AllAuditorName {
			allName = append(allName, r.ToString())
		}
	}
	return fmt.Sprintf("DROP SHARDING AUDITOR %s %s", ifExists,allName[])
}

type DataNodes struct {
	AllDataNode []*CommonIdentifier
}

func (dataNodes *DataNodes) ToString() string {
	var allDataNode []string
	for _, d := range dataNodes.AllDataNode {
		allDataNode = append(allDataNode, d.ToString())
	}
	return fmt.Sprintf("DATANODES(%s)", strings.Join(allDataNode, ","))
}

type AlterDefaultShardingStrategy struct {
	ShardingStrategy *ShardingStrategy
}

func (alterDefaultShardingStrategy *AlterDefaultShardingStrategy) ToString() string {
	return fmt.Sprintf("ALTER DEFAULT SHARDING type=(DATABASE | TABLE) STRATEGY (%s)", alterDefaultShardingStrategy.ShardingStrategy.ToString())
}
