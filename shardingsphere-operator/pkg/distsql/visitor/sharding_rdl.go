// Copyright 2023 Database Mesh Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package visitor

import (
	"fmt"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/distsql/ast"
	parser "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/distsql/visitor_parser/sharding"
)

type ShardingVisitor struct {
	parser.BaseRDLStatementVisitor
}

func (v *ShardingVisitor) VisitCreateShardingTableRule(ctx *parser.CreateShardingTableRuleContext) *ast.CreateShardingTableRule {
	stmt := &ast.CreateShardingTableRule{}
	if ctx.IfNotExists() != nil {
		stmt.IfNotExists = v.VisitIfNotExists(ctx.IfNotExists().(*parser.IfNotExistsContext))
	}
	if ctx.AllShardingTableRuleDefinition() != nil {
		for _, r := range ctx.AllShardingTableRuleDefinition() {
			stmt.AllShardingTableRuleDefinition = append(stmt.AllShardingTableRuleDefinition, v.VisitShardingTableRuleDefinition(r.(*parser.ShardingTableRuleDefinitionContext)))
		}
	}

	return stmt
}

func (v *ShardingVisitor) VisitAlterShardingTableRule(ctx *parser.AlterShardingTableRuleContext) *ast.AlterShardingTableRule {
	stmt := &ast.AlterShardingTableRule{}
	if ctx.AllShardingTableRuleDefinition() != nil {
		for _, r := range ctx.AllShardingTableRuleDefinition() {
			stmt.AllShardingTableRuleDefinition = append(stmt.AllShardingTableRuleDefinition, v.VisitShardingTableRuleDefinition(r.(*parser.ShardingTableRuleDefinitionContext)))
		}
	}
	return stmt
}

func (v *ShardingVisitor) VisitDropShardingTableRule(ctx *parser.DropShardingTableRuleContext) *ast.DropShardingTableRule {
	stmt := &ast.DropShardingTableRule{}
	if ctx.IfExists() != nil {
		stmt.IfExists = v.VisitIfExists(ctx.IfExists().(*parser.IfExistsContext))
	}
	if ctx.AllTableName() != nil {
		for _, t := range ctx.AllTableName() {
			stmt.AllTableName = append(stmt.AllTableName, v.VisitTableName(t.(*parser.TableNameContext)))
		}
	}
	return stmt
}

func (v *ShardingVisitor) VisitCreateShardingTableReferenceRule(ctx *parser.CreateShardingTableReferenceRuleContext) *ast.CreateShardingTableReferenceRule {
	stmt := &ast.CreateShardingTableReferenceRule{}
	if ctx.IfNotExists() != nil {
		stmt.IfNotExists = v.VisitIfNotExists(ctx.IfNotExists().(*parser.IfNotExistsContext))
	}
	if ctx.AllTableReferenceRuleDefinition() != nil {
		for _, r := range ctx.AllTableReferenceRuleDefinition() {
			stmt.AllTableReferenceRuleDefinition = append(stmt.AllTableReferenceRuleDefinition, v.VisitTableReferenceRuleDefinition(r.(*parser.TableReferenceRuleDefinitionContext)))
		}
	}
	return stmt
}

func (v *ShardingVisitor) VisitAlterShardingTableReferenceRule(ctx *parser.AlterShardingTableReferenceRuleContext) *ast.AlterShardingTableReferenceRule {
	stmt := &ast.AlterShardingTableReferenceRule{}
	if ctx.AllTableReferenceRuleDefinition() != nil {
		for _, r := range ctx.AllTableReferenceRuleDefinition() {
			stmt.AllTableReferenceRuleDefinition = append(stmt.AllTableReferenceRuleDefinition, v.VisitTableReferenceRuleDefinition(r.(*parser.TableReferenceRuleDefinitionContext)))
		}
	}
	return stmt
}

func (v *ShardingVisitor) VisitDropShardingTableReferenceRule(ctx *parser.DropShardingTableReferenceRuleContext) *ast.DropShardingTableReferenceRule {
	stmt := &ast.DropShardingTableReferenceRule{}
	if ctx.IfExists() != nil {
		stmt.IfExists = v.VisitIfExists(ctx.IfExists().(*parser.IfExistsContext))
	}
	if ctx.AllRuleName() != nil {
		for _, r := range ctx.AllRuleName() {
			stmt.AllRuleNames = append(stmt.AllRuleNames, v.VisitRuleName(r.(*parser.RuleNameContext)))
		}
	}
	return stmt
}

func (v *ShardingVisitor) VisitCreateBroadcastTableRule(ctx *parser.CreateBroadcastTableRuleContext) *ast.CreateBroadcastTableRule {
	stmt := &ast.CreateBroadcastTableRule{}
	if ctx.IfNotExists() != nil {
		stmt.IfNotExists = v.VisitIfNotExists(ctx.IfNotExists().(*parser.IfNotExistsContext))
	}
	if ctx.AllTableName() != nil {
		for _, t := range ctx.AllTableName() {
			stmt.AllTableName = append(stmt.AllTableName, v.VisitTableName(t.(*parser.TableNameContext)))
		}
	}
	return stmt
}

func (v *ShardingVisitor) VisitDropBroadcastTableRule(ctx *parser.DropBroadcastTableRuleContext) *ast.DropBroadcastTableRule {
	stmt := &ast.DropBroadcastTableRule{}
	if ctx.IfExists() != nil {
		stmt.IfExists = v.VisitIfExists(ctx.IfExists().(*parser.IfExistsContext))
	}
	if ctx.AllTableName() != nil {
		for _, t := range ctx.AllTableName() {
			stmt.AllTableName = append(stmt.AllTableName, v.VisitTableName(t.(*parser.TableNameContext)))
		}
	}
	return stmt
}

func (v *ShardingVisitor) VisitDropShardingAlgorithm(ctx *parser.DropShardingAlgorithmContext) *ast.DropShardingAlgorithm {
	stmt := &ast.DropShardingAlgorithm{}
	if ctx.IfExists() != nil {
		stmt.IfExists = v.VisitIfExists(ctx.IfExists().(*parser.IfExistsContext))
	}
	if ctx.AllShardingAlgorithmName() != nil {
		for _, algo := range ctx.AllShardingAlgorithmName() {
			stmt.AllShardingAlgorithmName = append(stmt.AllShardingAlgorithmName, v.VisitShardingAlgorithmName(algo.(*parser.ShardingAlgorithmNameContext)))
		}
	}
	return stmt
}

func (v *ShardingVisitor) VisitCreateDefaultShardingStrategy(ctx *parser.CreateDefaultShardingStrategyContext) *ast.CreateDefaultShardingStrategy {
	stmt := &ast.CreateDefaultShardingStrategy{}
	if ctx.IfNotExists() != nil {
		stmt.IfNotExists = v.VisitIfNotExists(ctx.IfNotExists().(*parser.IfNotExistsContext))
	}
	if ctx.ShardingStrategy() != nil {
		stmt.ShardingStrategy = v.VisitShardingStrategy(ctx.ShardingStrategy().(*parser.ShardingStrategyContext))
	}
	return stmt
}

func (v *ShardingVisitor) VisitAlterDefaultShardingStrategy(ctx *parser.AlterDefaultShardingStrategyContext) *ast.AlterDefaultShardingStrategy {
	stmt := &ast.AlterDefaultShardingStrategy{}
	if ctx.ShardingStrategy() != nil {
		stmt.ShardingStrategy = v.VisitShardingStrategy(ctx.ShardingStrategy().(*parser.ShardingStrategyContext))
	}
	return stmt
}

func (v *ShardingVisitor) VisitDropDefaultShardingStrategy(ctx *parser.DropDefaultShardingStrategyContext) *ast.DropDefaultShardingStrategy {
	stmt := &ast.DropDefaultShardingStrategy{}
	if ctx.IfExists() != nil {
		stmt.IfExists = v.VisitIfExists(ctx.IfExists().(*parser.IfExistsContext))
	}
	return stmt
}

func (v *ShardingVisitor) VisitDropShardingKeyGenerator(ctx *parser.DropShardingKeyGeneratorContext) *ast.DropShardingKeyGenerator {
	stmt := &ast.DropShardingKeyGenerator{}
	if ctx.IfExists() != nil {
		stmt.IfExists = v.VisitIfExists(ctx.IfExists().(*parser.IfExistsContext))
	}
	if ctx.AllKeyGeneratorName() != nil {
		for _, n := range ctx.AllKeyGeneratorName() {
			stmt.AllKeyGeneratorName = append(stmt.AllKeyGeneratorName, v.VisitKeyGeneratorName(n.(*parser.KeyGeneratorNameContext)))
		}
	}
	return stmt
}

func (v *ShardingVisitor) VisitDropShardingAuditor(ctx *parser.DropShardingAuditorContext) *ast.DropShardingAuditor {
	stmt := &ast.DropShardingAuditor{}
	if ctx.IfExists() != nil {
		stmt.IfExists = v.VisitIfExists(ctx.IfExists().(*parser.IfExistsContext))
	}
	if ctx.AllAuditorName() != nil {
		for _, n := range ctx.AllAuditorName() {
			stmt.AllAuditorName = append(stmt.AllAuditorName, v.VisitAuditorName(n.(*parser.AuditorNameContext)))
		}
	}
	return stmt
}

func (v *ShardingVisitor) VisitShardingTableRuleDefinition(ctx *parser.ShardingTableRuleDefinitionContext) *ast.ShardingTableRuleDefinition {
	stmt := &ast.ShardingTableRuleDefinition{}
	return stmt
}

func (v *ShardingVisitor) VisitShardingAutoTableRule(ctx *parser.ShardingAutoTableRuleContext) *ast.ShardingAutoTableRule {
	stmt := &ast.ShardingAutoTableRule{}
	if ctx.TableName() != nil {
		stmt.TableName = v.VisitTableName(ctx.TableName().(*parser.TableNameContext))
	}
	if ctx.StorageUnits() != nil {
		stmt.StorageUnits = v.VisitStorageUnits(ctx.StorageUnits().(*parser.StorageUnitsContext))
	}
	if ctx.AutoShardingColumnDefinition() != nil {
		stmt.AutoShardingColumnDefinition = v.VisitAutoShardingColumnDefinition(ctx.AutoShardingColumnDefinition().(*parser.AutoShardingColumnDefinitionContext))
	}
	if ctx.AlgorithmDefinition() != nil {
		stmt.AlgorithmDefinition = v.VisitAlgorithmDefinition(ctx.AlgorithmDefinition().(*parser.AlgorithmDefinitionContext))
	}
	if ctx.KeyGenerateDefinition() != nil {
		stmt.KeyGenerateDefinition = v.VisitKeyGenerateDefinition(ctx.KeyGenerateDefinition().(*parser.KeyGenerateDefinitionContext))
	}
	if ctx.AuditDefinition() != nil {
		stmt.AuditDefinition = v.VisitAuditDefinition(ctx.AuditDefinition().(*parser.AuditDefinitionContext))
	}
	return stmt
}

func (v *ShardingVisitor) VisitShardingTableRule(ctx *parser.ShardingTableRuleContext) *ast.ShardingTableRule {
	stmt := &ast.ShardingTableRule{}
	if ctx.TableName() != nil {
		stmt.TableName = v.VisitTableName(ctx.TableName().(*parser.TableNameContext))
	}
	if ctx.DataNodes() != nil {
		stmt.DataNodes = v.VisitDataNodes(ctx.DataNodes().(*parser.DataNodesContext))
	}
	if ctx.DatabaseStrategy() != nil {
		stmt.DatabaseStrategy = v.VisitDatabaseStrategy(ctx.DatabaseStrategy().(*parser.DatabaseStrategyContext))
	}
	if ctx.TableStrategy() != nil {
		stmt.TableStrategy = v.VisitTableStrategy(ctx.TableStrategy().(*parser.TableStrategyContext))
	}
	if ctx.KeyGenerateDefinition() != nil {
		stmt.KeyGenerateDefinition = v.VisitKeyGenerateDefinition(ctx.KeyGenerateDefinition().(*parser.KeyGenerateDefinitionContext))
	}
	if ctx.AuditDefinition() != nil {
		stmt.AuditDefinition = v.VisitAuditDefinition(ctx.AuditDefinition().(*parser.AuditDefinitionContext))
	}
	return stmt
}

func (v *ShardingVisitor) VisitKeyGeneratorName(ctx *parser.KeyGeneratorNameContext) *ast.CommonIdentifier {
	stmt := &ast.CommonIdentifier{}
	if ctx.IDENTIFIER_() != nil {
		stmt.Identifier = ctx.IDENTIFIER_().GetText()
	}
	return stmt
}

func (v *ShardingVisitor) VisitAuditorDefinition(ctx *parser.AuditorDefinitionContext) *ast.AuditorDefinition {
	stmt := &ast.AuditorDefinition{}
	if ctx.AuditorName() != nil {
		stmt.AuditorName = v.VisitAuditorName(ctx.AuditorName().(*parser.AuditorNameContext))
	}
	if ctx.AlgorithmDefinition() != nil {
		stmt.AlgorithmDefinition = v.VisitAlgorithmDefinition(ctx.AlgorithmDefinition().(*parser.AlgorithmDefinitionContext))
	}
	return stmt
}

func (v *ShardingVisitor) VisitAuditorName(ctx *parser.AuditorNameContext) *ast.CommonIdentifier {
	stmt := &ast.CommonIdentifier{}
	if ctx.IDENTIFIER_() != nil {
		stmt.Identifier = ctx.IDENTIFIER_().GetText()
	}
	return stmt
}

func (v *ShardingVisitor) VisitStorageUnits(ctx *parser.StorageUnitsContext) *ast.StorageUnits {
	stmt := &ast.StorageUnits{}
	if ctx.AllStorageUnit() != nil {
		for _, s := range ctx.AllStorageUnit() {
			stmt.AllStorageUnit = append(stmt.AllStorageUnit, v.VisitStorageUnit(s.(*parser.StorageUnitContext)))
		}
	}
	return stmt
}

func (v *ShardingVisitor) VisitStorageUnit(ctx *parser.StorageUnitContext) *ast.StorageUnit {
	stmt := &ast.StorageUnit{}
	switch {
	case ctx.STRING_() != nil:
		stmt.String = ctx.STRING_().GetText()
	case ctx.IDENTIFIER_() != nil:
		stmt.Identifier = ctx.IDENTIFIER_().GetText()
	}
	return stmt
}

func (v *ShardingVisitor) VisitDataNodes(ctx *parser.DataNodesContext) *ast.DataNodes {
	stmt := &ast.DataNodes{}
	if ctx.AllDataNode() != nil {
		for _, d := range ctx.AllDataNode() {
			stmt.AllDataNode = append(stmt.AllDataNode, v.VisitDataNode(d.(*parser.DataNodeContext)))
		}
	}
	return stmt
}

func (v *ShardingVisitor) VisitDataNode(ctx *parser.DataNodeContext) *ast.CommonIdentifier {
	stmt := &ast.CommonIdentifier{}
	if ctx.STRING_() != nil {
		stmt.Identifier = ctx.STRING_().GetText()
	}
	return stmt
}

func (v *ShardingVisitor) VisitAutoShardingColumnDefinition(ctx *parser.AutoShardingColumnDefinitionContext) *ast.AutoShardingColumnDefinition {
	stmt := &ast.AutoShardingColumnDefinition{}
	if ctx.ShardingColumn() != nil {
		stmt.ShardingColumn = v.VisitShardingColumn(ctx.ShardingColumn().(*parser.ShardingColumnContext))
	}
	return stmt
}

func (v *ShardingVisitor) VisitShardingColumnDefinition(ctx *parser.ShardingColumnDefinitionContext) *ast.ShardingColumnDefinition {
	stmt := &ast.ShardingColumnDefinition{}
	switch {
	case ctx.ShardingColumn() != nil:
		stmt.ShardingColumn = v.VisitShardingColumn(ctx.ShardingColumn().(*parser.ShardingColumnContext))
	case ctx.ShardingColumns() != nil:
		stmt.ShardingColumns = v.VisitShardingColumns(ctx.ShardingColumns().(*parser.ShardingColumnsContext))
	}
	return stmt
}

func (v *ShardingVisitor) VisitShardingColumn(ctx *parser.ShardingColumnContext) *ast.ShardingColumn {
	stmt := &ast.ShardingColumn{}
	if ctx.ColumnName() != nil {
		stmt.ColumnName = v.VisitColumnName(ctx.ColumnName().(*parser.ColumnNameContext))
	}
	return stmt
}

func (v *ShardingVisitor) VisitShardingColumns(ctx *parser.ShardingColumnsContext) *ast.ShardingColumns {
	stmt := &ast.ShardingColumns{}
	if ctx.AllColumnName() != nil {
		for _, c := range ctx.AllColumnName() {
			stmt.AllColumnName = append(stmt.AllColumnName, v.VisitColumnName(c.(*parser.ColumnNameContext)))
		}
	}
	if ctx.ColumnName(0) != nil {
		stmt.ColumnName = v.VisitColumnName(ctx.ColumnName(0).(*parser.ColumnNameContext))
	}
	return stmt
}

func (v *ShardingVisitor) VisitShardingAlgorithm(ctx *parser.ShardingAlgorithmContext) *ast.ShardingAlgorithm {
	stmt := &ast.ShardingAlgorithm{}
	if ctx.AlgorithmDefinition() != nil {
		stmt.AlgorithmDefinition = v.VisitAlgorithmDefinition(ctx.AlgorithmDefinition().(*parser.AlgorithmDefinitionContext))
	}
	return stmt
}

func (v *ShardingVisitor) VisitShardingStrategy(ctx *parser.ShardingStrategyContext) *ast.ShardingStrategy {
	stmt := &ast.ShardingStrategy{}
	if ctx.StrategyType() != nil {
		stmt.StrategyType = v.VisitStrategyType(ctx.StrategyType().(*parser.StrategyTypeContext))
	}
	if ctx.ShardingColumnDefinition() != nil {
		stmt.ShardingColumnDefinition = v.VisitShardingColumnDefinition(ctx.ShardingColumnDefinition().(*parser.ShardingColumnDefinitionContext))
	}
	if ctx.ShardingAlgorithm() != nil {
		stmt.ShardingAlgorithm = v.VisitShardingAlgorithm(ctx.ShardingAlgorithm().(*parser.ShardingAlgorithmContext))
	}
	return stmt
}

func (v *ShardingVisitor) VisitDatabaseStrategy(ctx *parser.DatabaseStrategyContext) *ast.DatabaseStrategy {
	stmt := &ast.DatabaseStrategy{}
	if ctx.ShardingStrategy() != nil {
		stmt.ShardingStrategy = v.VisitShardingStrategy(ctx.ShardingStrategy().(*parser.ShardingStrategyContext))
	}
	return stmt
}

func (v *ShardingVisitor) VisitTableStrategy(ctx *parser.TableStrategyContext) *ast.TableStrategy {
	stmt := &ast.TableStrategy{}
	if ctx.ShardingStrategy() != nil {
		stmt.ShardingStrategy = v.VisitShardingStrategy(ctx.ShardingStrategy().(*parser.ShardingStrategyContext))
	}
	return stmt
}

func (v *ShardingVisitor) VisitKeyGenerateDefinition(ctx *parser.KeyGenerateDefinitionContext) *ast.KeyGenerateDefinition {
	stmt := &ast.KeyGenerateDefinition{}
	if ctx.ColumnName() != nil {
		stmt.ColumnName = v.VisitColumnName(ctx.ColumnName().(*parser.ColumnNameContext))
	}
	if ctx.AlgorithmDefinition() != nil {
		stmt.AlgorithmDefinition = v.VisitAlgorithmDefinition(ctx.AlgorithmDefinition().(*parser.AlgorithmDefinitionContext))
	}
	return stmt
}

func (v *ShardingVisitor) VisitAuditDefinition(ctx *parser.AuditDefinitionContext) *ast.AuditDefinition {
	stmt := &ast.AuditDefinition{}
	if ctx.MultiAuditDefinition() != nil {
		stmt.MultiAuditDefinition = v.VisitMultiAuditDefinition(ctx.MultiAuditDefinition().(*parser.MultiAuditDefinitionContext))
	}
	if ctx.AuditAllowHintDisable() != nil {
		stmt.AuditAllowHintDisable = v.VisitAuditAllowHintDisable(ctx.AuditAllowHintDisable().(*parser.AuditAllowHintDisableContext))
	}
	return stmt
}

func (v *ShardingVisitor) VisitMultiAuditDefinition(ctx *parser.MultiAuditDefinitionContext) *ast.MultiAuditDefinition {
	stmt := &ast.MultiAuditDefinition{}
	if ctx.AllSingleAuditDefinition() != nil {
		for _, a := range ctx.AllSingleAuditDefinition() {
			stmt.AllSingleAuditDefinition = append(stmt.AllSingleAuditDefinition, v.VisitSingleAuditDefinition(a.(*parser.SingleAuditDefinitionContext)))
		}
	}
	return stmt
}

func (v *ShardingVisitor) VisitSingleAuditDefinition(ctx *parser.SingleAuditDefinitionContext) *ast.SingleAuditDefinition {
	stmt := &ast.SingleAuditDefinition{}
	if ctx.AlgorithmDefinition() != nil {
		stmt.AlgorithmDefinition = v.VisitAlgorithmDefinition(ctx.AlgorithmDefinition().(*parser.AlgorithmDefinitionContext))
	}
	return stmt
}

func (v *ShardingVisitor) VisitAuditAllowHintDisable(ctx *parser.AuditAllowHintDisableContext) *ast.AuditAllowHintDisable {
	stmt := &ast.AuditAllowHintDisable{}
	switch {
	case ctx.TRUE() != nil:
		stmt.AuditAllowHintDisable = ctx.TRUE().GetText()
	case ctx.FALSE() != nil:
		stmt.AuditAllowHintDisable = ctx.FALSE().GetText()
	}
	return stmt
}

func (v *ShardingVisitor) VisitColumnName(ctx *parser.ColumnNameContext) *ast.CommonIdentifier {
	stmt := &ast.CommonIdentifier{}
	if ctx.IDENTIFIER_() != nil {
		stmt.Identifier = ctx.IDENTIFIER_().GetText()
	}
	return stmt
}

func (v *ShardingVisitor) VisitTableReferenceRuleDefinition(ctx *parser.TableReferenceRuleDefinitionContext) *ast.TableReferenceRuleDefinition {
	stmt := &ast.TableReferenceRuleDefinition{}
	if ctx.RuleName() != nil {
		stmt.RuleName = v.VisitRuleName(ctx.RuleName().(*parser.RuleNameContext))
	}
	if ctx.AllTableName() != nil {
		for _, t := range ctx.AllTableName() {
			stmt.AllTableName = append(stmt.AllTableName, v.VisitTableName(t.(*parser.TableNameContext)))
		}
	}
	return stmt
}

func (v *ShardingVisitor) VisitStrategyType(ctx *parser.StrategyTypeContext) *ast.StrategyType {
	stmt := &ast.StrategyType{}
	switch {
	case ctx.STRING_() != nil:
		stmt.String = ctx.STRING_().GetText()
	}
	if ctx.BuildInStrategyType() != nil {
		stmt.BuildInStrategyType = v.VisitBuildInStrategyType(ctx.BuildInStrategyType().(*parser.BuildInStrategyTypeContext))
	}
	return stmt
}

func (v *ShardingVisitor) VisitBuildInStrategyType(ctx *parser.BuildInStrategyTypeContext) *ast.BuildInStrategyType {
	stmt := &ast.BuildInStrategyType{}
	switch {
	case ctx.STANDARD() != nil:
		stmt.BuildInStrategyType = ctx.STANDARD().GetText()
	case ctx.COMPLEX() != nil:
		stmt.BuildInStrategyType = ctx.COMPLEX().GetText()
	case ctx.HINT() != nil:
		stmt.BuildInStrategyType = ctx.HINT().GetText()
	case ctx.NONE() != nil:
		stmt.BuildInStrategyType = ctx.NONE().GetText()
	}
	return stmt
}

func (v *ShardingVisitor) VisitIfExists(ctx *parser.IfExistsContext) *ast.IfExists {
	return &ast.IfExists{
		IfExists: fmt.Sprintf("%s %s", ctx.IF().GetText(), ctx.EXISTS().GetText()),
	}
}

func (v *ShardingVisitor) VisitIfNotExists(ctx *parser.IfNotExistsContext) *ast.IfNotExists {
	return &ast.IfNotExists{
		IfNotExists: fmt.Sprintf("%s %s %s", ctx.IF().GetText(), ctx.NOT().GetText(), ctx.EXISTS().GetText()),
	}
}

// nolint
func (v *ShardingVisitor) VisitLiteral(ctx *parser.LiteralContext) *ast.Literal {
	stmt := &ast.Literal{}
	switch {
	case ctx.STRING_() != nil:
		stmt.Literal = ctx.STRING_().GetText()
	case ctx.MINUS_() != nil:
		stmt.Literal = ctx.MINUS_().GetText()
	case ctx.INT_() != nil:
		stmt.Literal = ctx.INT_().GetText()
	case ctx.TRUE() != nil:
		stmt.Literal = ctx.TRUE().GetText()
	case ctx.FALSE() != nil:
		stmt.Literal = ctx.FALSE().GetText()
	}
	return stmt
}

func (v *ShardingVisitor) VisitAlgorithmDefinition(ctx *parser.AlgorithmDefinitionContext) *ast.ShardingAlgorithmDefinition {
	stmt := &ast.ShardingAlgorithmDefinition{}
	if ctx.AlgorithmTypeName() != nil {
		stmt.ShardingAlgorithmTypeName = v.VisitAlgorithmTypeName(ctx.AlgorithmTypeName().(*parser.AlgorithmTypeNameContext))
	}
	if ctx.PropertiesDefinition() != nil {
		stmt.PropertiesDefinition = v.VisitPropertiesDefinition(ctx.PropertiesDefinition().(*parser.PropertiesDefinitionContext))
	}
	return stmt
}

// nolint
func (v *ShardingVisitor) VisitAlgorithmTypeName(ctx *parser.AlgorithmTypeNameContext) *ast.ShardingAlgorithmTypeName {
	stmt := &ast.ShardingAlgorithmTypeName{}
	switch {
	case ctx.STRING_() != nil:
		stmt.String = ctx.STRING_().GetText()
	case ctx.BuildInShardingAlgorithmType() != nil:
		stmt.BuildInShardingAlgorithmType = v.VisitBuildInShardingAlgorithmType(ctx.BuildInShardingAlgorithmType().(*parser.BuildInShardingAlgorithmTypeContext))
	case ctx.BuildInKeyGenerateAlgorithmType() != nil:
		stmt.BuildInKeyGenerateAlgorithmType = v.VisitBuildInKeyGenerateAlgorithmType(ctx.BuildInKeyGenerateAlgorithmType().(*parser.BuildInKeyGenerateAlgorithmTypeContext))
	case ctx.BuildInShardingAuditAlgorithmType() != nil:
		stmt.BuildInShardingAuditAlgorithmType = v.VisitBuildInShardingAuditAlgorithmType(ctx.BuildInShardingAuditAlgorithmType().(*parser.BuildInShardingAuditAlgorithmTypeContext))
	}
	return stmt
}

// nolint
func (v *ShardingVisitor) VisitBuildInShardingAlgorithmType(ctx *parser.BuildInShardingAlgorithmTypeContext) *ast.BuildInCommon {
	stmt := &ast.BuildInCommon{}
	switch {
	case ctx.MOD() != nil:
		stmt.String = ctx.MOD().GetText()
	case ctx.HASH_MOD() != nil:
		stmt.String = ctx.HASH_MOD().GetText()
	case ctx.VOLUME_RANGE() != nil:
		stmt.String = ctx.VOLUME_RANGE().GetText()
	case ctx.BOUNDARY_RANGE() != nil:
		stmt.String = ctx.BOUNDARY_RANGE().GetText()
	case ctx.AUTO_INTERVAL() != nil:
		stmt.String = ctx.AUTO_INTERVAL().GetText()
	case ctx.INLINE() != nil:
		stmt.String = ctx.INLINE().GetText()
	case ctx.INTERVAL() != nil:
		stmt.String = ctx.INTERVAL().GetText()
	case ctx.COSID_MOD() != nil:
		stmt.String = ctx.COSID_MOD().GetText()
	case ctx.COSID_INTERVAL() != nil:
		stmt.String = ctx.COSID_INTERVAL().GetText()
	case ctx.COSID_INTERVAL_SNOWFLAKE() != nil:
		stmt.String = ctx.COSID_INTERVAL_SNOWFLAKE().GetText()
	case ctx.COMPLEX_INLINE() != nil:
		stmt.String = ctx.COMPLEX_INLINE().GetText()
	case ctx.HINT_INLINE() != nil:
		stmt.String = ctx.HINT_INLINE().GetText()
	case ctx.CLASS_BASED() != nil:
		stmt.String = ctx.CLASS_BASED().GetText()
	}
	return stmt
}

// nolint
func (v *ShardingVisitor) VisitBuildInKeyGenerateAlgorithmType(ctx *parser.BuildInKeyGenerateAlgorithmTypeContext) *ast.BuildInCommon {
	stmt := &ast.BuildInCommon{}
	switch {
	case ctx.SNOWFLAKE() != nil:
		stmt.String = ctx.SNOWFLAKE().GetText()
	case ctx.NANOID() != nil:
		stmt.String = ctx.NANOID().GetText()
	case ctx.UUID() != nil:
		stmt.String = ctx.UUID().GetText()
	case ctx.COSID() != nil:
		stmt.String = ctx.COSID().GetText()
	case ctx.COSID_SNOWFLAKE() != nil:
		stmt.String = ctx.COSID_SNOWFLAKE().GetText()
	}
	return stmt
}

// nolint
func (v *ShardingVisitor) VisitBuildInShardingAuditAlgorithmType(ctx *parser.BuildInShardingAuditAlgorithmTypeContext) *ast.BuildInCommon {
	stmt := &ast.BuildInCommon{}
	switch {
	case ctx.DML_SHARDING_CONDITIONS() != nil:
		stmt.String = ctx.DML_SHARDING_CONDITIONS().GetText()
	}
	return stmt
}

// nolint
func (v *ShardingVisitor) VisitPropertiesDefinition(ctx *parser.PropertiesDefinitionContext) *ast.PropertiesDefinition {
	stmt := &ast.PropertiesDefinition{}
	if ctx.Properties() != nil {
		stmt.Properties = v.VisitProperties(ctx.Properties().(*parser.PropertiesContext))
	}
	return stmt
}

func (v *ShardingVisitor) VisitProperties(ctx *parser.PropertiesContext) *ast.Properties {
	stmt := &ast.Properties{}
	for _, p := range ctx.AllProperty() {
		stmt.Properties = append(stmt.Properties, v.VisitProperty(p.(*parser.PropertyContext)))
	}
	return stmt
}

func (v *ShardingVisitor) VisitProperty(ctx *parser.PropertyContext) *ast.Property {
	stmt := &ast.Property{}
	if ctx.STRING_() != nil {
		stmt.Key = ctx.STRING_().GetText()
	}
	if ctx.Literal() != nil {
		stmt.Literal = v.VisitLiteral(ctx.Literal().(*parser.LiteralContext))
	}
	return stmt
}

func (v *ShardingVisitor) VisitTableName(ctx *parser.TableNameContext) *ast.CommonIdentifier {
	stmt := &ast.CommonIdentifier{}
	if ctx.IDENTIFIER_() != nil {
		stmt.Identifier = ctx.IDENTIFIER_().GetText()
	}
	return stmt
}

func (v *ShardingVisitor) VisitShardingAlgorithmName(ctx *parser.ShardingAlgorithmNameContext) *ast.CommonIdentifier {
	stmt := &ast.CommonIdentifier{}
	if ctx.IDENTIFIER_() != nil {
		stmt.Identifier = ctx.IDENTIFIER_().GetText()
	}
	return stmt
}

func (v *ShardingVisitor) VisitRuleName(ctx *parser.RuleNameContext) *ast.CommonIdentifier {
	stmt := &ast.CommonIdentifier{}
	if ctx.IDENTIFIER_() != nil {
		stmt.Identifier = ctx.IDENTIFIER_().GetText()
	}
	return stmt
}
