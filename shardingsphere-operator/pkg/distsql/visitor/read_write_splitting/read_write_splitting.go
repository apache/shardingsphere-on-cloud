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

package visitor

import (
	"fmt"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/distsql/ast"
	parser "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/distsql/visitor_parser/read_write_splitting"
)

type ReadWriteSplittingVisitor struct {
	parser.BaseRDLStatementVisitor
}

func (v *ReadWriteSplittingVisitor) VisitCreateReadwriteSplittingRule(ctx *parser.CreateReadwriteSplittingRuleContext) *ast.CreateReadwriteSplittingRule {
	stmt := &ast.CreateReadwriteSplittingRule{}
	if ctx.IfNotExists() != nil {
		stmt.IfNotExists = v.VisitIfNotExists(ctx.IfNotExists().(*parser.IfNotExistsContext))
	}
	if ctx.AllReadwriteSplittingRuleDefinition() != nil {
		for _, r := range ctx.AllReadwriteSplittingRuleDefinition() {
			stmt.AllReadwriteSplittingRuleDefinition = append(stmt.AllReadwriteSplittingRuleDefinition, v.VisitReadwriteSplittingRuleDefinition(r.(*parser.ReadwriteSplittingRuleDefinitionContext)))
		}
	}
	return stmt
}

func (v *ReadWriteSplittingVisitor) VisitAlterReadwriteSplittingRule(ctx *parser.AlterReadwriteSplittingRuleContext) *ast.AlterReadwriteSplittingRule {
	stmt := &ast.AlterReadwriteSplittingRule{}
	if ctx.AllReadwriteSplittingRuleDefinition() != nil {
		for _, r := range ctx.AllReadwriteSplittingRuleDefinition() {
			stmt.AllReadwriteSplittingRuleDefinition = append(stmt.AllReadwriteSplittingRuleDefinition, v.VisitReadwriteSplittingRuleDefinition(r.(*parser.ReadwriteSplittingRuleDefinitionContext)))
		}
	}
	return stmt
}

func (v *ReadWriteSplittingVisitor) VisitDropReadwriteSplittingRule(ctx *parser.DropReadwriteSplittingRuleContext) *ast.DropReadwriteSplittingRule {
	stmt := &ast.DropReadwriteSplittingRule{}
	if ctx.IfExists() != nil {
		stmt.IfExists = v.VisitIfExists(ctx.IfExists().(*parser.IfExistsContext))
	}
	if ctx.AllRuleName() != nil {
		for _, r := range ctx.AllRuleName() {
			stmt.AllRuleName = append(stmt.AllRuleName, v.VisitRuleName(r.(*parser.RuleNameContext)))
		}
	}
	return stmt
}

func (v *ReadWriteSplittingVisitor) VisitReadwriteSplittingRuleDefinition(ctx *parser.ReadwriteSplittingRuleDefinitionContext) *ast.ReadWriteSplittingRuleDefinition {
	stmt := &ast.ReadWriteSplittingRuleDefinition{}
	if ctx.RuleName() != nil {
		stmt.RuleName = v.VisitRuleName(ctx.RuleName().(*parser.RuleNameContext))
	}
	if ctx.DataSourceDefinition() != nil {
		stmt.DataSourceDefinition = v.VisitDataSourceDefinition(ctx.DataSourceDefinition().(*parser.DataSourceDefinitionContext))
	}
	if ctx.TransactionalReadQueryStrategy() != nil {
		stmt.TransactionalReadQueryStrategy = v.VisitTransactionalReadQueryStrategy(ctx.TransactionalReadQueryStrategy().(*parser.TransactionalReadQueryStrategyContext))
	}
	if ctx.AlgorithmDefinition() != nil {
		stmt.AlgorithmDefinition = v.VisitAlgorithmDefinition(ctx.AlgorithmDefinition().(*parser.AlgorithmDefinitionContext))
	}
	return stmt
}

func (v *ReadWriteSplittingVisitor) VisitDataSourceDefinition(ctx *parser.DataSourceDefinitionContext) *ast.DataSourceDefinition {
	stmt := &ast.DataSourceDefinition{}
	if ctx.WriteStorageUnit() != nil {
		stmt.WriteStorageUnit = v.VisitWriteStorageUnit(ctx.WriteStorageUnit().(*parser.WriteStorageUnitContext))
	}
	if ctx.ReadStorageUnits() != nil {
		stmt.ReadStorageUnits = v.VisitReadStorageUnits(ctx.ReadStorageUnits().(*parser.ReadStorageUnitsContext))
	}
	return stmt
}

func (v *ReadWriteSplittingVisitor) VisitRuleName(ctx *parser.RuleNameContext) *ast.CommonIdentifier {
	stmt := &ast.CommonIdentifier{}
	if ctx.IDENTIFIER_() != nil {
		stmt.Identifier = ctx.IDENTIFIER_().GetText()
	}
	return stmt
}

func (v *ReadWriteSplittingVisitor) VisitWriteStorageUnit(ctx *parser.WriteStorageUnitContext) *ast.WriteStorageUnit {
	stmt := &ast.WriteStorageUnit{}
	if ctx.WriteStorageUnitName() != nil {
		stmt.WriteStorageUnitName = v.VisitWriteStorageUnitName(ctx.WriteStorageUnitName().(*parser.WriteStorageUnitNameContext))
	}
	return stmt
}

func (v *ReadWriteSplittingVisitor) VisitReadStorageUnits(ctx *parser.ReadStorageUnitsContext) *ast.ReadStorageUnits {
	stmt := &ast.ReadStorageUnits{}
	if ctx.ReadStorageUnitsNames() != nil {
		stmt.ReadStorageUnitsNames = v.VisitReadStorageUnitsNames(ctx.ReadStorageUnitsNames().(*parser.ReadStorageUnitsNamesContext))
	}
	return stmt
}

func (v *ReadWriteSplittingVisitor) VisitTransactionalReadQueryStrategy(ctx *parser.TransactionalReadQueryStrategyContext) *ast.TransactionalReadQueryStrategy {
	stmt := &ast.TransactionalReadQueryStrategy{}
	if ctx.TransactionalReadQueryStrategyName() != nil {
		stmt.TransactionalReadQueryStrategyName = v.VisitTransactionalReadQueryStrategyName(ctx.TransactionalReadQueryStrategyName().(*parser.TransactionalReadQueryStrategyNameContext))
	}
	return stmt
}

func (v *ReadWriteSplittingVisitor) VisitWriteStorageUnitName(ctx *parser.WriteStorageUnitNameContext) *ast.WriteStorageUnitName {
	stmt := &ast.WriteStorageUnitName{}
	if ctx.StorageUnitName() != nil {
		stmt.StorageUnitName = v.VisitStorageUnitName(ctx.StorageUnitName().(*parser.StorageUnitNameContext))
	}
	return stmt
}

func (v *ReadWriteSplittingVisitor) VisitReadStorageUnitsNames(ctx *parser.ReadStorageUnitsNamesContext) *ast.ReadStorageUnitsNames {
	stmt := &ast.ReadStorageUnitsNames{}
	if ctx.AllStorageUnitName() != nil {
		for _, s := range ctx.AllStorageUnitName() {
			stmt.AllStorageUnitName = append(stmt.AllStorageUnitName, v.VisitStorageUnitName(s.(*parser.StorageUnitNameContext)))
		}
	}
	return stmt
}

func (v *ReadWriteSplittingVisitor) VisitTransactionalReadQueryStrategyName(ctx *parser.TransactionalReadQueryStrategyNameContext) *ast.TransactionalReadQueryStrategyName {
	return &ast.TransactionalReadQueryStrategyName{
		String: ctx.STRING_().GetText(),
	}
}

func (v *ReadWriteSplittingVisitor) VisitIfExists(ctx *parser.IfExistsContext) *ast.IfExists {
	return &ast.IfExists{
		IfExists: fmt.Sprintf("%s %s", ctx.IF().GetText(), ctx.EXISTS().GetText()),
	}
}

func (v *ReadWriteSplittingVisitor) VisitIfNotExists(ctx *parser.IfNotExistsContext) *ast.IfNotExists {
	return &ast.IfNotExists{
		IfNotExists: fmt.Sprintf("%s %s %s", ctx.IF().GetText(), ctx.NOT().GetText(), ctx.EXISTS().GetText()),
	}
}

// nolint
func (v *ReadWriteSplittingVisitor) VisitLiteral(ctx *parser.LiteralContext) *ast.Literal {
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

func (v *ReadWriteSplittingVisitor) VisitAlgorithmDefinition(ctx *parser.AlgorithmDefinitionContext) *ast.AlgorithmDefinition {
	stmt := &ast.AlgorithmDefinition{}
	if ctx.AlgorithmTypeName() != nil {
		stmt.AlgorithmTypeName = v.VisitAlgorithmTypeName(ctx.AlgorithmTypeName().(*parser.AlgorithmTypeNameContext))
	}

	if ctx.PropertiesDefinition() != nil {
		stmt.PropertiesDefinition = v.VisitPropertiesDefinition(ctx.PropertiesDefinition().(*parser.PropertiesDefinitionContext))
	}
	return stmt
}

func (v *ReadWriteSplittingVisitor) VisitAlgorithmTypeName(ctx *parser.AlgorithmTypeNameContext) *ast.AlgorithmTypeName {
	stmt := &ast.AlgorithmTypeName{}
	switch {
	case ctx.STRING_() != nil:
		stmt.String = ctx.STRING_().GetText()
	case ctx.BuildInReadQueryLoadBalanceAlgorithmType() != nil:
		stmt.BuildinAlgorithmTypeName = v.VisitBuildInReadQueryLoadBalanceAlgorithmType(ctx.BuildInReadQueryLoadBalanceAlgorithmType().(*parser.BuildInReadQueryLoadBalanceAlgorithmTypeContext))
	}
	return stmt
}

// nolint
func (v *ReadWriteSplittingVisitor) VisitBuildInReadQueryLoadBalanceAlgorithmType(ctx *parser.BuildInReadQueryLoadBalanceAlgorithmTypeContext) *ast.BuildinAlgorithmTypeName {
	stmt := &ast.BuildinAlgorithmTypeName{}
	switch {
	case ctx.ROUND_ROBIN() != nil:
		stmt.AlgorithmTypeName = ctx.ROUND_ROBIN().GetText()
	case ctx.RANDOM() != nil:
		stmt.AlgorithmTypeName = ctx.RANDOM().GetText()
	case ctx.WEIGHT() != nil:
		stmt.AlgorithmTypeName = ctx.WEIGHT().GetText()
	}
	return stmt
}

func (v *ReadWriteSplittingVisitor) VisitPropertiesDefinition(ctx *parser.PropertiesDefinitionContext) *ast.PropertiesDefinition {
	stmt := &ast.PropertiesDefinition{}
	if ctx.Properties() != nil {
		stmt.Properties = v.VisitProperties(ctx.Properties().(*parser.PropertiesContext))
	}
	return stmt
}

func (v *ReadWriteSplittingVisitor) VisitProperties(ctx *parser.PropertiesContext) *ast.Properties {
	stmt := &ast.Properties{}
	for _, p := range ctx.AllProperty() {
		stmt.Properties = append(stmt.Properties, v.VisitProperty(p.(*parser.PropertyContext)))
	}
	return stmt
}

func (v *ReadWriteSplittingVisitor) VisitProperty(ctx *parser.PropertyContext) *ast.Property {
	stmt := &ast.Property{}
	if ctx.STRING_() != nil {
		stmt.Key = ctx.STRING_().GetText()
	}
	if ctx.Literal() != nil {
		stmt.Literal = v.VisitLiteral(ctx.Literal().(*parser.LiteralContext))
	}
	return stmt
}

func (v *ReadWriteSplittingVisitor) VisitDatabaseName(ctx *parser.DatabaseNameContext) *ast.CommonIdentifier {
	stmt := &ast.CommonIdentifier{}
	if ctx.IDENTIFIER_() != nil {
		stmt.Identifier = ctx.IDENTIFIER_().GetText()
	}
	return stmt
}

func (v *ReadWriteSplittingVisitor) VisitGroupName(ctx *parser.GroupNameContext) *ast.CommonIdentifier {
	stmt := &ast.CommonIdentifier{}
	if ctx.IDENTIFIER_() != nil {
		stmt.Identifier = ctx.IDENTIFIER_().GetText()
	}
	return stmt
}

func (v *ReadWriteSplittingVisitor) VisitStorageUnitName(ctx *parser.StorageUnitNameContext) *ast.CommonIdentifier {
	stmt := &ast.CommonIdentifier{}
	if ctx.IDENTIFIER_() != nil {
		stmt.Identifier = ctx.IDENTIFIER_().GetText()
	}
	return stmt
}
