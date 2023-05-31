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
	parser "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/distsql/visitor_parser/shadow"
)

type ShadowVisitor struct {
	parser.BaseRDLStatementVisitor
}

func (v *ShadowVisitor) VisitCreateShadowRule(ctx *parser.CreateShadowRuleContext) *ast.CreateShadowRule {
	stmt := &ast.CreateShadowRule{}
	if ctx.IfNotExists() != nil {
		stmt.IfNotExists = v.VisitIfNotExists(ctx.IfNotExists().(*parser.IfNotExistsContext))
	}
	if ctx.AllShadowRuleDefinition() != nil {
		for _, r := range ctx.AllShadowRuleDefinition() {
			stmt.AllShadowRuleDefinition = append(stmt.AllShadowRuleDefinition, v.VisitShadowRuleDefinition(r.(*parser.ShadowRuleDefinitionContext)))
		}
	}
	return stmt
}

func (v *ShadowVisitor) VisitAlterShadowRule(ctx *parser.AlterShadowRuleContext) *ast.AlterShadowRule {
	stmt := &ast.AlterShadowRule{}
	if ctx.AllShadowRuleDefinition() != nil {
		for _, r := range ctx.AllShadowRuleDefinition() {
			stmt.AllShadowRuleDefinition = append(stmt.AllShadowRuleDefinition, v.VisitShadowRuleDefinition(r.(*parser.ShadowRuleDefinitionContext)))
		}
	}
	return stmt
}

func (v *ShadowVisitor) VisitDropShadowRule(ctx *parser.DropShadowRuleContext) *ast.DropShadowRule {
	stmt := &ast.DropShadowRule{}
	if ctx.IfExists() != nil {
		stmt.IfExists = v.VisitIfExists(ctx.IfExists().(*parser.IfExistsContext))
	}
	if ctx.AllRuleName() != nil {
		for _, rn := range ctx.AllRuleName() {
			stmt.AllRuleName = append(stmt.AllRuleName, v.VisitRuleName(rn.(*parser.RuleNameContext)))
		}
	}
	return stmt
}

func (v *ShadowVisitor) VisitDropShadowAlgorithm(ctx *parser.DropShadowAlgorithmContext) *ast.DropShadowAlgorithm {
	stmt := &ast.DropShadowAlgorithm{}
	if ctx.IfExists() != nil {
		stmt.IfExists = v.VisitIfExists(ctx.IfExists().(*parser.IfExistsContext))
	}
	if ctx.AllAlgorithmName() != nil {
		for _, algorithm := range ctx.AllAlgorithmName() {
			stmt.AllAlgorithmName = append(stmt.AllAlgorithmName, v.VisitAlgorithmName(algorithm.(*parser.AlgorithmNameContext)))
		}
	}
	return stmt
}

func (v *ShadowVisitor) VisitCreateDefaultShadowAlgorithm(ctx *parser.CreateDefaultShadowAlgorithmContext) *ast.CreateDefaultShadowAlgorithm {
	stmt := &ast.CreateDefaultShadowAlgorithm{}
	if ctx.IfNotExists() != nil {
		stmt.IfNotExists = v.VisitIfNotExists(ctx.IfNotExists().(*parser.IfNotExistsContext))
	}
	if ctx.AlgorithmDefinition() != nil {
		stmt.AlgorithmDefinition = v.VisitAlgorithmDefinition(ctx.AlgorithmDefinition().(*parser.AlgorithmDefinitionContext))
	}
	return stmt
}

func (v *ShadowVisitor) VisitDropDefaultShadowAlgorithm(ctx *parser.DropDefaultShadowAlgorithmContext) *ast.DropDefaultShadowAlgorithm {
	stmt := &ast.DropDefaultShadowAlgorithm{}
	if ctx.IfExists() != nil {
		stmt.IfExists = v.VisitIfExists(ctx.IfExists().(*parser.IfExistsContext))
	}
	return stmt
}

func (v *ShadowVisitor) VisitAlterDefaultShadowAlgorithm(ctx *parser.AlterDefaultShadowAlgorithmContext) *ast.AlterDefaultShadowAlgorithm {
	stmt := &ast.AlterDefaultShadowAlgorithm{}
	if ctx.AlgorithmDefinition() != nil {
		stmt.AlgorithmDefinition = v.VisitAlgorithmDefinition(ctx.AlgorithmDefinition().(*parser.AlgorithmDefinitionContext))
	}
	return stmt
}

func (v *ShadowVisitor) VisitShadowRuleDefinition(ctx *parser.ShadowRuleDefinitionContext) *ast.ShadowRuleDefinition {
	stmt := &ast.ShadowRuleDefinition{}
	if ctx.RuleName() != nil {
		stmt.RuleName = v.VisitRuleName(ctx.RuleName().(*parser.RuleNameContext))
	}
	if ctx.Source() != nil {
		stmt.Source = v.VisitSource(ctx.Source().(*parser.SourceContext))
	}
	if ctx.SHADOW() != nil {
		stmt.Shadow = v.VisitShadow(ctx.Shadow().(*parser.ShadowContext))
	}
	if ctx.AllShadowTableRule() != nil {
		for _, r := range ctx.AllShadowTableRule() {
			stmt.AllShadowTableRule = append(stmt.AllShadowTableRule, v.VisitShadowTableRule(r.(*parser.ShadowTableRuleContext)))
		}
	}
	return stmt
}

func (v *ShadowVisitor) VisitShadowTableRule(ctx *parser.ShadowTableRuleContext) *ast.ShadowTableRule {
	stmt := &ast.ShadowTableRule{}
	if ctx.TableName() != nil {
		stmt.TableName = v.VisitTableName(ctx.TableName().(*parser.TableNameContext))
	}
	if ctx.AllAlgorithmDefinition() != nil {
		for _, algorithm := range ctx.AllAlgorithmDefinition() {
			stmt.AllAlgorithmDefinition = append(stmt.AllAlgorithmDefinition, v.VisitAlgorithmDefinition(algorithm.(*parser.AlgorithmDefinitionContext)))
		}
	}
	return stmt
}

func (v *ShadowVisitor) VisitSource(ctx *parser.SourceContext) *ast.CommonIdentifier {
	stmt := &ast.CommonIdentifier{}
	if ctx.IDENTIFIER_() != nil {
		stmt.Identifier = ctx.IDENTIFIER_().GetText()
	}
	return stmt
}

func (v *ShadowVisitor) VisitShadow(ctx *parser.ShadowContext) *ast.CommonIdentifier {
	stmt := &ast.CommonIdentifier{}
	if ctx.IDENTIFIER_() != nil {
		stmt.Identifier = ctx.IDENTIFIER_().GetText()
	}
	return stmt
}

func (v *ShadowVisitor) VisitTableName(ctx *parser.TableNameContext) *ast.CommonIdentifier {
	stmt := &ast.CommonIdentifier{}
	if ctx.IDENTIFIER_() != nil {
		stmt.Identifier = ctx.IDENTIFIER_().GetText()
	}
	return stmt
}

func (v *ShadowVisitor) VisitAlgorithmName(ctx *parser.AlgorithmNameContext) *ast.CommonIdentifier {
	stmt := &ast.CommonIdentifier{}
	if ctx.IDENTIFIER_() != nil {
		stmt.Identifier = ctx.IDENTIFIER_().GetText()
	}
	return stmt
}

func (v *ShadowVisitor) VisitIfExists(ctx *parser.IfExistsContext) *ast.IfExists {
	return &ast.IfExists{
		IfExists: fmt.Sprintf("%s %s", ctx.IF().GetText(), ctx.EXISTS().GetText()),
	}
}

func (v *ShadowVisitor) VisitIfNotExists(ctx *parser.IfNotExistsContext) *ast.IfNotExists {
	return &ast.IfNotExists{
		IfNotExists: fmt.Sprintf("%s %s %s", ctx.IF().GetText(), ctx.NOT().GetText(), ctx.EXISTS().GetText()),
	}
}

// nolint
func (v *ShadowVisitor) VisitLiteral(ctx *parser.LiteralContext) *ast.Literal {
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

func (v *ShadowVisitor) VisitAlgorithmDefinition(ctx *parser.AlgorithmDefinitionContext) *ast.AlgorithmDefinition {
	stmt := &ast.AlgorithmDefinition{}
	if ctx.AlgorithmTypeName() != nil {
		stmt.AlgorithmTypeName = v.VisitAlgorithmTypeName(ctx.AlgorithmTypeName().(*parser.AlgorithmTypeNameContext))
	}

	if ctx.PropertiesDefinition() != nil {
		stmt.PropertiesDefinition = v.VisitPropertiesDefinition(ctx.PropertiesDefinition().(*parser.PropertiesDefinitionContext))
	}
	return stmt
}

func (v *ShadowVisitor) VisitAlgorithmTypeName(ctx *parser.AlgorithmTypeNameContext) *ast.AlgorithmTypeName {
	stmt := &ast.AlgorithmTypeName{}
	switch {
	case ctx.STRING_() != nil:
		stmt.String = ctx.STRING_().GetText()
	case ctx.BuildInShadowAlgorithmType() != nil:
		stmt.BuildinAlgorithmTypeName = v.VisitBuildInShadowAlgorithmType(ctx.BuildInShadowAlgorithmType().(*parser.BuildInShadowAlgorithmTypeContext))
	}
	return stmt
}

// nolint
func (v *ShadowVisitor) VisitBuildInShadowAlgorithmType(ctx *parser.BuildInShadowAlgorithmTypeContext) *ast.BuildinAlgorithmTypeName {
	stmt := &ast.BuildinAlgorithmTypeName{}
	switch {
	case ctx.VALUE_MATCH() != nil:
		stmt.AlgorithmTypeName = ctx.VALUE_MATCH().GetText()
	case ctx.REGEX_MATCH() != nil:
		stmt.AlgorithmTypeName = ctx.REGEX_MATCH().GetText()
	case ctx.SQL_HINT() != nil:
		stmt.AlgorithmTypeName = ctx.SQL_HINT().GetText()
	}
	return stmt
}

func (v *ShadowVisitor) VisitPropertiesDefinition(ctx *parser.PropertiesDefinitionContext) *ast.PropertiesDefinition {
	stmt := &ast.PropertiesDefinition{}
	if ctx.Properties() != nil {
		stmt.Properties = v.VisitProperties(ctx.Properties().(*parser.PropertiesContext))
	}
	return stmt
}

func (v *ShadowVisitor) VisitProperties(ctx *parser.PropertiesContext) *ast.Properties {
	stmt := &ast.Properties{}
	for _, p := range ctx.AllProperty() {
		stmt.Properties = append(stmt.Properties, v.VisitProperty(p.(*parser.PropertyContext)))
	}
	return stmt
}

func (v *ShadowVisitor) VisitProperty(ctx *parser.PropertyContext) *ast.Property {
	stmt := &ast.Property{}
	if ctx.STRING_() != nil {
		stmt.Key = ctx.STRING_().GetText()
	}
	if ctx.Literal() != nil {
		stmt.Literal = v.VisitLiteral(ctx.Literal().(*parser.LiteralContext))
	}
	return stmt
}

func (v *ShadowVisitor) VisitRuleName(ctx *parser.RuleNameContext) *ast.CommonIdentifier {
	stmt := &ast.CommonIdentifier{}
	if ctx.IDENTIFIER_() != nil {
		stmt.Identifier = ctx.IDENTIFIER_().GetText()
	}
	return stmt
}
