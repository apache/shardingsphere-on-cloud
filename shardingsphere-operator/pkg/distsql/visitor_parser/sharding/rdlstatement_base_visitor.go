// Code generated from RDLStatement.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // RDLStatement

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseRDLStatementVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseRDLStatementVisitor) VisitCreateShardingTableRule(ctx *CreateShardingTableRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitAlterShardingTableRule(ctx *AlterShardingTableRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitDropShardingTableRule(ctx *DropShardingTableRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitCreateShardingTableReferenceRule(ctx *CreateShardingTableReferenceRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitAlterShardingTableReferenceRule(ctx *AlterShardingTableReferenceRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitDropShardingTableReferenceRule(ctx *DropShardingTableReferenceRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitCreateBroadcastTableRule(ctx *CreateBroadcastTableRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitDropBroadcastTableRule(ctx *DropBroadcastTableRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitDropShardingAlgorithm(ctx *DropShardingAlgorithmContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitCreateDefaultShardingStrategy(ctx *CreateDefaultShardingStrategyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitAlterDefaultShardingStrategy(ctx *AlterDefaultShardingStrategyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitDropDefaultShardingStrategy(ctx *DropDefaultShardingStrategyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitDropShardingKeyGenerator(ctx *DropShardingKeyGeneratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitDropShardingAuditor(ctx *DropShardingAuditorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitShardingTableRuleDefinition(ctx *ShardingTableRuleDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitShardingAutoTableRule(ctx *ShardingAutoTableRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitShardingTableRule(ctx *ShardingTableRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitKeyGeneratorName(ctx *KeyGeneratorNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitAuditorDefinition(ctx *AuditorDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitAuditorName(ctx *AuditorNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitStorageUnits(ctx *StorageUnitsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitStorageUnit(ctx *StorageUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitDataNodes(ctx *DataNodesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitDataNode(ctx *DataNodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitAutoShardingColumnDefinition(ctx *AutoShardingColumnDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitShardingColumnDefinition(ctx *ShardingColumnDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitShardingColumn(ctx *ShardingColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitShardingColumns(ctx *ShardingColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitShardingAlgorithm(ctx *ShardingAlgorithmContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitShardingStrategy(ctx *ShardingStrategyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitDatabaseStrategy(ctx *DatabaseStrategyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitTableStrategy(ctx *TableStrategyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitKeyGenerateDefinition(ctx *KeyGenerateDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitAuditDefinition(ctx *AuditDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitMultiAuditDefinition(ctx *MultiAuditDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitSingleAuditDefinition(ctx *SingleAuditDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitAuditAllowHintDisable(ctx *AuditAllowHintDisableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitColumnName(ctx *ColumnNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitTableReferenceRuleDefinition(ctx *TableReferenceRuleDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitStrategyType(ctx *StrategyTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitBuildInStrategyType(ctx *BuildInStrategyTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitIfExists(ctx *IfExistsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitIfNotExists(ctx *IfNotExistsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitAlgorithmDefinition(ctx *AlgorithmDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitAlgorithmTypeName(ctx *AlgorithmTypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitBuildInShardingAlgorithmType(ctx *BuildInShardingAlgorithmTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitBuildInKeyGenerateAlgorithmType(ctx *BuildInKeyGenerateAlgorithmTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitBuildInShardingAuditAlgorithmType(ctx *BuildInShardingAuditAlgorithmTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitPropertiesDefinition(ctx *PropertiesDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitProperties(ctx *PropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitProperty(ctx *PropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitTableName(ctx *TableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitShardingAlgorithmName(ctx *ShardingAlgorithmNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitRuleName(ctx *RuleNameContext) interface{} {
	return v.VisitChildren(ctx)
}
