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
	parser "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/distsql/visitor_parser/mask"
)

type MaskVisitor struct {
	parser.BaseRDLStatementVisitor
}

func (v *MaskVisitor) VisitCreateMaskRule(ctx *parser.CreateMaskRuleContext) *ast.CreateMaskRule {
	stmt := &ast.CreateMaskRule{}
	if ctx.IfNotExists() != nil {
		stmt.IfNotExists = v.VisitIfNotExists(ctx.IfNotExists().(*parser.IfNotExistsContext))
	}
	if ctx.TABLE() != nil {
		stmt.Table = ctx.TABLE().GetText()
	}
	if ctx.AllMaskRuleDefinition() != nil {
		for _, m := range ctx.AllMaskRuleDefinition() {
			stmt.AllMaskRuleDefinition = append(stmt.AllMaskRuleDefinition, v.VisitMaskRuleDefinition(m.(*parser.MaskRuleDefinitionContext)))
		}
	}
	return stmt
}

func (v *MaskVisitor) VisitAlterMaskRule(ctx *parser.AlterMaskRuleContext) *ast.AlterMaskRule {
	stmt := &ast.AlterMaskRule{}
	if ctx.TABLE() != nil {
		stmt.Table = ctx.TABLE().GetText()
	}
	if ctx.AllMaskRuleDefinition() != nil {
		for _, m := range ctx.AllMaskRuleDefinition() {
			stmt.AllMaskRuleDefinition = append(stmt.AllMaskRuleDefinition, v.VisitMaskRuleDefinition(m.(*parser.MaskRuleDefinitionContext)))
		}
	}
	return stmt
}

func (v *MaskVisitor) VisitDropMaskRule(ctx *parser.DropMaskRuleContext) *ast.DropMaskRule {
	stmt := &ast.DropMaskRule{}
	if ctx.TABLE() != nil {
		stmt.Table = ctx.TABLE().GetText()
	}
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

func (v *MaskVisitor) VisitMaskRuleDefinition(ctx *parser.MaskRuleDefinitionContext) *ast.MaskRuleDefinition {
	stmt := &ast.MaskRuleDefinition{}
	if ctx.RuleName() != nil {
		stmt.RuleName = v.VisitRuleName(ctx.RuleName().(*parser.RuleNameContext))
	}
	return stmt
}

func (v *MaskVisitor) VisitColumnDefinition(ctx *parser.ColumnDefinitionContext) *ast.ColumnDefinition {
	stmt := &ast.ColumnDefinition{}
	if ctx.ColumnName() != nil {
		stmt.ColumnName = v.VisitColumnName(ctx.ColumnName().(*parser.ColumnNameContext))
	}
	return stmt
}

func (v *MaskVisitor) VisitColumnName(ctx *parser.ColumnNameContext) *ast.CommonIdentifier {
	stmt := &ast.CommonIdentifier{}
	if ctx.IDENTIFIER_() != nil {
		stmt.Identifier = ctx.IDENTIFIER_().GetText()
	}
	return stmt
}

func (v *MaskVisitor) VisitIfExists(ctx *parser.IfExistsContext) *ast.IfExists {
	return &ast.IfExists{
		IfExists: fmt.Sprintf("%s %s", ctx.IF().GetText(), ctx.EXISTS().GetText()),
	}
}

func (v *MaskVisitor) VisitIfNotExists(ctx *parser.IfNotExistsContext) *ast.IfNotExists {
	return &ast.IfNotExists{
		IfNotExists: fmt.Sprintf("%s %s %s", ctx.IF().GetText(), ctx.NOT().GetText(), ctx.EXISTS().GetText()),
	}
}

// nolint
func (v *MaskVisitor) VisitLiteral(ctx *parser.LiteralContext) *ast.Literal {
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

func (v *MaskVisitor) VisitAlgorithmDefinition(ctx *parser.AlgorithmDefinitionContext) *ast.AlgorithmDefinition {
	stmt := &ast.AlgorithmDefinition{}
	if ctx.AlgorithmTypeName() != nil {
		stmt.AlgorithmTypeName = v.VisitAlgorithmTypeName(ctx.AlgorithmTypeName().(*parser.AlgorithmTypeNameContext))
	}

	if ctx.PropertiesDefinition() != nil {
		stmt.PropertiesDefinition = v.VisitPropertiesDefinition(ctx.PropertiesDefinition().(*parser.PropertiesDefinitionContext))
	}
	return stmt
}

func (v *MaskVisitor) VisitAlgorithmTypeName(ctx *parser.AlgorithmTypeNameContext) *ast.AlgorithmTypeName {
	stmt := &ast.AlgorithmTypeName{}
	switch {
	case ctx.STRING_() != nil:
		stmt.String = ctx.STRING_().GetText()
	case ctx.BuildInMaskAlgorithmType() != nil:
		stmt.BuildinAlgorithmTypeName = v.VisitBuildInMaskAlgorithmType(ctx.BuildInMaskAlgorithmType().(*parser.BuildInMaskAlgorithmTypeContext))
	}
	return stmt
}

func (v *MaskVisitor) VisitBuildInMaskAlgorithmType(ctx *parser.BuildInMaskAlgorithmTypeContext) *ast.BuildinAlgorithmTypeName {
	stmt := &ast.BuildinAlgorithmTypeName{}
	switch {
	case ctx.MD5() != nil:
		stmt.AlgorithmTypeName = ctx.MD5().GetText()
	case ctx.KEEP_FIRST_N_LAST_M() != nil:
		stmt.AlgorithmTypeName = ctx.KEEP_FIRST_N_LAST_M().GetText()
	case ctx.KEEP_FROM_X_TO_Y() != nil:
		stmt.AlgorithmTypeName = ctx.KEEP_FROM_X_TO_Y().GetText()
	case ctx.MASK_FIRST_N_LAST_M() != nil:
		stmt.AlgorithmTypeName = ctx.MASK_FIRST_N_LAST_M().GetText()
	case ctx.MASK_FROM_X_TO_Y() != nil:
		stmt.AlgorithmTypeName = ctx.MASK_FROM_X_TO_Y().GetText()
	case ctx.MASK_BEFORE_SPECIAL_CHARS() != nil:
		stmt.AlgorithmTypeName = ctx.MASK_BEFORE_SPECIAL_CHARS().GetText()
	case ctx.MASK_AFTER_SPECIAL_CHARS() != nil:
		stmt.AlgorithmTypeName = ctx.MASK_AFTER_SPECIAL_CHARS().GetText()
	case ctx.PERSONAL_IDENTITY_NUMBER_RANDOM_REPLACE() != nil:
		stmt.AlgorithmTypeName = ctx.PERSONAL_IDENTITY_NUMBER_RANDOM_REPLACE().GetText()
	case ctx.MILITARY_IDENTITY_NUMBER_RANDOM_REPLACE() != nil:
		stmt.AlgorithmTypeName = ctx.MILITARY_IDENTITY_NUMBER_RANDOM_REPLACE().GetText()
	case ctx.LANDLINE_NUMBER_RANDOM_REPLACE() != nil:
		stmt.AlgorithmTypeName = ctx.LANDLINE_NUMBER_RANDOM_REPLACE().GetText()
	case ctx.TELEPHONE_RANDOM_REPLACE() != nil:
		stmt.AlgorithmTypeName = ctx.TELEPHONE_RANDOM_REPLACE().GetText()
	case ctx.UNIFIED_CREDIT_CODE_RANDOM_REPLACE() != nil:
		stmt.AlgorithmTypeName = ctx.UNIFIED_CREDIT_CODE_RANDOM_REPLACE().GetText()
	case ctx.GENERIC_TABLE_RANDOM_REPLACE() != nil:
		stmt.AlgorithmTypeName = ctx.GENERIC_TABLE_RANDOM_REPLACE().GetText()
	}
	return stmt

}

func (v *MaskVisitor) VisitPropertiesDefinition(ctx *parser.PropertiesDefinitionContext) *ast.PropertiesDefinition {
	stmt := &ast.PropertiesDefinition{}
	if ctx.Properties() != nil {
		stmt.Properties = v.VisitProperties(ctx.Properties().(*parser.PropertiesContext))
	}
	return stmt
}

func (v *MaskVisitor) VisitProperties(ctx *parser.PropertiesContext) *ast.Properties {
	stmt := &ast.Properties{}
	for _, p := range ctx.AllProperty() {
		stmt.Properties = append(stmt.Properties, v.VisitProperty(p.(*parser.PropertyContext)))
	}
	return stmt
}

func (v *MaskVisitor) VisitProperty(ctx *parser.PropertyContext) *ast.Property {
	stmt := &ast.Property{}
	if ctx.STRING_() != nil {
		stmt.Key = ctx.STRING_().GetText()
	}
	if ctx.Literal() != nil {
		stmt.Literal = v.VisitLiteral(ctx.Literal().(*parser.LiteralContext))
	}
	return stmt
}

func (v *MaskVisitor) VisitRuleName(ctx *parser.RuleNameContext) *ast.CommonIdentifier {
	stmt := &ast.CommonIdentifier{}
	if ctx.IDENTIFIER_() != nil {
		stmt.Identifier = ctx.IDENTIFIER_().GetText()
	}
	return stmt
}
