// Code generated from RDLStatement.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // RDLStatement

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by RDLStatementParser.
type RDLStatementVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by RDLStatementParser#createShardingTableRule.
	VisitCreateShardingTableRule(ctx *CreateShardingTableRuleContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#alterShardingTableRule.
	VisitAlterShardingTableRule(ctx *AlterShardingTableRuleContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#dropShardingTableRule.
	VisitDropShardingTableRule(ctx *DropShardingTableRuleContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#createShardingTableReferenceRule.
	VisitCreateShardingTableReferenceRule(ctx *CreateShardingTableReferenceRuleContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#alterShardingTableReferenceRule.
	VisitAlterShardingTableReferenceRule(ctx *AlterShardingTableReferenceRuleContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#dropShardingTableReferenceRule.
	VisitDropShardingTableReferenceRule(ctx *DropShardingTableReferenceRuleContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#createBroadcastTableRule.
	VisitCreateBroadcastTableRule(ctx *CreateBroadcastTableRuleContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#dropBroadcastTableRule.
	VisitDropBroadcastTableRule(ctx *DropBroadcastTableRuleContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#dropShardingAlgorithm.
	VisitDropShardingAlgorithm(ctx *DropShardingAlgorithmContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#createDefaultShardingStrategy.
	VisitCreateDefaultShardingStrategy(ctx *CreateDefaultShardingStrategyContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#alterDefaultShardingStrategy.
	VisitAlterDefaultShardingStrategy(ctx *AlterDefaultShardingStrategyContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#dropDefaultShardingStrategy.
	VisitDropDefaultShardingStrategy(ctx *DropDefaultShardingStrategyContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#dropShardingKeyGenerator.
	VisitDropShardingKeyGenerator(ctx *DropShardingKeyGeneratorContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#dropShardingAuditor.
	VisitDropShardingAuditor(ctx *DropShardingAuditorContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#shardingTableRuleDefinition.
	VisitShardingTableRuleDefinition(ctx *ShardingTableRuleDefinitionContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#shardingAutoTableRule.
	VisitShardingAutoTableRule(ctx *ShardingAutoTableRuleContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#shardingTableRule.
	VisitShardingTableRule(ctx *ShardingTableRuleContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#keyGeneratorName.
	VisitKeyGeneratorName(ctx *KeyGeneratorNameContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#auditorDefinition.
	VisitAuditorDefinition(ctx *AuditorDefinitionContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#auditorName.
	VisitAuditorName(ctx *AuditorNameContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#storageUnits.
	VisitStorageUnits(ctx *StorageUnitsContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#storageUnit.
	VisitStorageUnit(ctx *StorageUnitContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#dataNodes.
	VisitDataNodes(ctx *DataNodesContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#dataNode.
	VisitDataNode(ctx *DataNodeContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#autoShardingColumnDefinition.
	VisitAutoShardingColumnDefinition(ctx *AutoShardingColumnDefinitionContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#shardingColumnDefinition.
	VisitShardingColumnDefinition(ctx *ShardingColumnDefinitionContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#shardingColumn.
	VisitShardingColumn(ctx *ShardingColumnContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#shardingColumns.
	VisitShardingColumns(ctx *ShardingColumnsContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#shardingAlgorithm.
	VisitShardingAlgorithm(ctx *ShardingAlgorithmContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#shardingStrategy.
	VisitShardingStrategy(ctx *ShardingStrategyContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#databaseStrategy.
	VisitDatabaseStrategy(ctx *DatabaseStrategyContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#tableStrategy.
	VisitTableStrategy(ctx *TableStrategyContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#keyGenerateDefinition.
	VisitKeyGenerateDefinition(ctx *KeyGenerateDefinitionContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#auditDefinition.
	VisitAuditDefinition(ctx *AuditDefinitionContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#multiAuditDefinition.
	VisitMultiAuditDefinition(ctx *MultiAuditDefinitionContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#singleAuditDefinition.
	VisitSingleAuditDefinition(ctx *SingleAuditDefinitionContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#auditAllowHintDisable.
	VisitAuditAllowHintDisable(ctx *AuditAllowHintDisableContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#columnName.
	VisitColumnName(ctx *ColumnNameContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#tableReferenceRuleDefinition.
	VisitTableReferenceRuleDefinition(ctx *TableReferenceRuleDefinitionContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#strategyType.
	VisitStrategyType(ctx *StrategyTypeContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#buildInStrategyType.
	VisitBuildInStrategyType(ctx *BuildInStrategyTypeContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#ifExists.
	VisitIfExists(ctx *IfExistsContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#ifNotExists.
	VisitIfNotExists(ctx *IfNotExistsContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#algorithmDefinition.
	VisitAlgorithmDefinition(ctx *AlgorithmDefinitionContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#algorithmTypeName.
	VisitAlgorithmTypeName(ctx *AlgorithmTypeNameContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#buildInShardingAlgorithmType.
	VisitBuildInShardingAlgorithmType(ctx *BuildInShardingAlgorithmTypeContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#buildInKeyGenerateAlgorithmType.
	VisitBuildInKeyGenerateAlgorithmType(ctx *BuildInKeyGenerateAlgorithmTypeContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#buildInShardingAuditAlgorithmType.
	VisitBuildInShardingAuditAlgorithmType(ctx *BuildInShardingAuditAlgorithmTypeContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#propertiesDefinition.
	VisitPropertiesDefinition(ctx *PropertiesDefinitionContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#properties.
	VisitProperties(ctx *PropertiesContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#property.
	VisitProperty(ctx *PropertyContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#tableName.
	VisitTableName(ctx *TableNameContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#shardingAlgorithmName.
	VisitShardingAlgorithmName(ctx *ShardingAlgorithmNameContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#ruleName.
	VisitRuleName(ctx *RuleNameContext) interface{}
}
