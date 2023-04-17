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
	parser "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/distsql/visitor_parser/encrypt"
)

type Visitor struct {
	parser.BaseRDLStatementVisitor
}

func (v *Visitor) VisitCreateEncryptRule(ctx *parser.CreateEncryptRuleContext) *ast.CreateEncryptRule {
	stmt := &ast.CreateEncryptRule{}

	if ctx.IfNotExists() != nil {
		stmt.IfNotExists = v.VisitIfNotExists(ctx.IfNotExists().(*parser.IfNotExistsContext))
	}

	if ctx.EncryptRuleDefinition(0) != nil {
		stmt.EncryptRuleDefinition = v.VisitEncryptRuleDefinition(ctx.EncryptRuleDefinition(0).(*parser.EncryptRuleDefinitionContext))
	}

	if ctx.AllEncryptRuleDefinition() != nil {
		for _, r := range ctx.AllEncryptRuleDefinition() {
			if r != nil {
				stmt.AllEncryptRuleDefinition = append(stmt.AllEncryptRuleDefinition, v.VisitEncryptRuleDefinition(r.(*parser.EncryptRuleDefinitionContext)))
			}
		}
	}

	return stmt
}

func (v *Visitor) VisitIfNotExists(ctx *parser.IfNotExistsContext) *ast.IfNotExists {
	return &ast.IfNotExists{
		IfNotExists: fmt.Sprintf("%s %s %s", ctx.IF().GetText(), ctx.NOT().GetText(), ctx.EXISTS().GetText()),
	}
}

func (v *Visitor) VisitIfExists(ctx *parser.IfExistsContext) *ast.IfExists {
	return &ast.IfExists{
		IfExists: fmt.Sprintf("%s %s", ctx.IF().GetText(), ctx.EXISTS().GetText()),
	}
}

func (v *Visitor) VisitAlterEncryptRule(ctx *parser.AlterEncryptRuleContext) *ast.AlterEncryptRule {
	stmt := &ast.AlterEncryptRule{}
	// if ctx.EncryptRuleDefinition(0) != nil {
	// 	stmt.EncryptRuleDefinition = v.VisitEncryptRuleDefinition(ctx.EncryptRuleDefinition(0).(*parser.EncryptRuleDefinitionContext))
	// }

	if ctx.AllEncryptRuleDefinition() != nil {
		for _, encryptRuleDefinition := range ctx.AllEncryptRuleDefinition() {
			stmt.AllEncryptRuleDefinitionList = append(stmt.AllEncryptRuleDefinitionList, v.VisitEncryptRuleDefinition(encryptRuleDefinition.(*parser.EncryptRuleDefinitionContext)))
		}
	}
	return stmt
}

func (v *Visitor) VisitDropEncryptRule(ctx *parser.DropEncryptRuleContext) *ast.DropEncryptRule {
	stmt := &ast.DropEncryptRule{}

	if ctx.IfExists() != nil {
		stmt.IfExists = v.VisitIfExists(ctx.IfExists().(*parser.IfExistsContext))
	}

	if ctx.AllTableName() != nil {
		for _, tableName := range ctx.AllTableName() {
			stmt.AllTableName = append(stmt.AllTableName, v.VisitTableName(tableName.(*parser.TableNameContext)))
		}
	}
	return stmt
}

func (v *Visitor) VisitEncryptRuleDefinition(ctx *parser.EncryptRuleDefinitionContext) *ast.EncryptRuleDefinition {
	stmt := &ast.EncryptRuleDefinition{}

	if ctx.TableName() != nil {
		stmt.TableName = v.VisitTableName(ctx.TableName().(*parser.TableNameContext))
	}

	if ctx.ResourceDefinition() != nil {
		stmt.ResourceDefinition = v.VisitResourceDefinition(ctx.ResourceDefinition().(*parser.ResourceDefinitionContext))
	}

	// if ctx.EncryptColumnDefinition(0) != nil {
	// 	stmt.EncryptColumnDefinition = v.VisitEncryptColumnDefinition(ctx.EncryptColumnDefinition(0).(*parser.EncryptColumnDefinitionContext))
	// }

	if ctx.AllEncryptColumnDefinition() != nil {
		for _, column := range ctx.AllEncryptColumnDefinition() {
			stmt.AllEncryptColumnDefinition = append(stmt.AllEncryptColumnDefinition, v.VisitEncryptColumnDefinition(column.(*parser.EncryptColumnDefinitionContext)))
		}
	}

	if ctx.QueryWithCipherColumn() != nil {
		stmt.QueryWithCipherColumn = v.VisitQueryWithCipherColumn(ctx.QueryWithCipherColumn().(*parser.QueryWithCipherColumnContext))
	}

	return stmt
}

func (v *Visitor) VisitQueryWithCipherColumn(ctx *parser.QueryWithCipherColumnContext) *ast.QueryWithCipherColumn {
	stmt := &ast.QueryWithCipherColumn{}
	switch {
	case ctx.TRUE() != nil:
		stmt.QueryWithCipherColumn = ctx.TRUE().GetText()
	case ctx.FALSE() != nil:
		stmt.QueryWithCipherColumn = ctx.FALSE().GetText()
	}
	return stmt
}

func (v *Visitor) VisitEncryptColumnDefinition(ctx *parser.EncryptColumnDefinitionContext) *ast.EncryptColumnDefinition {
	stmt := &ast.EncryptColumnDefinition{}

	if ctx.ColumnDefinition() != nil {
		stmt.ColumnDefinition = v.VisitColumnDefinition(ctx.ColumnDefinition().(*parser.ColumnDefinitionContext))
	}

	if ctx.PlainColumnDefinition() != nil {
		stmt.PlainColumnDefinition = v.VisitPlainColumnDefinition(ctx.PlainColumnDefinition().(*parser.PlainColumnDefinitionContext))
	}

	if ctx.CipherColumnDefinition() != nil {
		stmt.CipherColumnDefinition = v.VisitCipherColumnDefinition(ctx.CipherColumnDefinition().(*parser.CipherColumnDefinitionContext))
	}

	if ctx.AssistedQueryColumnDefinition() != nil {
		stmt.AssistedQueryColumnDefinition = v.VisitAssistedQueryColumnDefinition(ctx.AssistedQueryColumnDefinition().(*parser.AssistedQueryColumnDefinitionContext))
	}

	if ctx.LikeQueryColumnDefinition() != nil {
		stmt.LikeQueryColumnDefinition = v.VisitLikeQueryColumnDefinition(ctx.LikeQueryColumnDefinition().(*parser.LikeQueryColumnDefinitionContext))
	}

	if ctx.EncryptAlgorithm() != nil {
		stmt.EncryptAlgorithm = v.VisitEncryptAlgorithm(ctx.EncryptAlgorithm().(*parser.EncryptAlgorithmContext))
	}

	if ctx.AssistedQueryAlgorithm() != nil {
		stmt.AssistedQueryAlgorithm = v.VisitAssistedQueryAlgorithm(ctx.AssistedQueryAlgorithm().(*parser.AssistedQueryAlgorithmContext))
	}

	if ctx.LikeQueryAlgorithm() != nil {
		stmt.LikeQueryAlgorithm = v.VisitLikeQueryAlgorithm(ctx.LikeQueryAlgorithm().(*parser.LikeQueryAlgorithmContext))
	}

	if ctx.QueryWithCipherColumn() != nil {
		stmt.QueryWithCipherColumn = v.VisitQueryWithCipherColumn(ctx.QueryWithCipherColumn().(*parser.QueryWithCipherColumnContext))
	}

	return stmt
}

func (v *Visitor) VisitPlainColumnDefinition(ctx *parser.PlainColumnDefinitionContext) *ast.PlainColumnDefinition {
	stmt := &ast.PlainColumnDefinition{}
	if ctx.PlainColumnName() != nil {
		stmt.PlainColumnName = v.VisitPlainColumnName(ctx.PlainColumnName().(*parser.PlainColumnNameContext))
	}

	if ctx.DataType() != nil {
		stmt.DataType = v.VisitDataType(ctx.DataType().(*parser.DataTypeContext))
	}
	return stmt
}

func (v *Visitor) VisitPlainColumnName(ctx *parser.PlainColumnNameContext) *ast.CommonIdentifier {
	stmt := &ast.CommonIdentifier{}
	if ctx.IDENTIFIER_() != nil {
		stmt.Identifier = ctx.IDENTIFIER_().GetText()
	}
	return stmt
}

func (v *Visitor) VisitCipherColumnDefinition(ctx *parser.CipherColumnDefinitionContext) *ast.CipherColumnDefinition {
	stmt := &ast.CipherColumnDefinition{}
	if ctx.CipherColumnName() != nil {
		stmt.CipherColumnName = v.VisitCipherColumnName(ctx.CipherColumnName().(*parser.CipherColumnNameContext))
	}

	if ctx.DataType() != nil {
		stmt.DataType = v.VisitDataType(ctx.DataType().(*parser.DataTypeContext))
	}

	return stmt
}

func (v *Visitor) VisitCipherColumnName(ctx *parser.CipherColumnNameContext) *ast.CommonIdentifier {
	stmt := &ast.CommonIdentifier{}
	if ctx.IDENTIFIER_() != nil {
		stmt.Identifier = ctx.IDENTIFIER_().GetText()
	}
	return stmt
}

func (v *Visitor) VisitAssistedQueryColumnDefinition(ctx *parser.AssistedQueryColumnDefinitionContext) *ast.AssistedQueryColumnDefinition {
	stmt := &ast.AssistedQueryColumnDefinition{}
	if ctx.AssistedQueryColumnName() != nil {
		stmt.AssistedQueryColumnName = v.VisitAssistedQueryColumnName(ctx.AssistedQueryColumnName().(*parser.AssistedQueryColumnNameContext))
	}

	if ctx.DataType() != nil {
		stmt.DataType = v.VisitDataType(ctx.DataType().(*parser.DataTypeContext))
	}

	return stmt
}

func (v *Visitor) VisitAssistedQueryColumnName(ctx *parser.AssistedQueryColumnNameContext) *ast.CommonIdentifier {
	stmt := &ast.CommonIdentifier{}
	if ctx.IDENTIFIER_() != nil {
		stmt.Identifier = ctx.IDENTIFIER_().GetText()
	}
	return stmt
}

func (v *Visitor) VisitLikeQueryColumnDefinition(ctx *parser.LikeQueryColumnDefinitionContext) *ast.LikeQueryColumnDefinition {
	stmt := &ast.LikeQueryColumnDefinition{}
	if ctx.LikeQueryColumnName() != nil {
		stmt.LikeQueryColumnName = v.VisitLikeQueryColumnName(ctx.LikeQueryColumnName().(*parser.LikeQueryColumnNameContext))
	}

	if ctx.DataType() != nil {
		stmt.DataType = v.VisitDataType(ctx.DataType().(*parser.DataTypeContext))
	}
	return stmt
}

func (v *Visitor) VisitLikeQueryColumnName(ctx *parser.LikeQueryColumnNameContext) *ast.CommonIdentifier {
	stmt := &ast.CommonIdentifier{}
	if ctx.IDENTIFIER_() != nil {
		stmt.Identifier = ctx.IDENTIFIER_().GetText()
	}
	return stmt
}

func (v *Visitor) VisitAssistedQueryAlgorithm(ctx *parser.AssistedQueryAlgorithmContext) *ast.AssistedQueryAlgorithm {
	stmt := &ast.AssistedQueryAlgorithm{}
	if ctx.AlgorithmDefinition() != nil {
		stmt.AlgorithmDefinition = v.VisitAlgorithmDefinition(ctx.AlgorithmDefinition().(*parser.AlgorithmDefinitionContext))
	}
	return stmt
}

func (v *Visitor) VisitLikeQueryAlgorithm(ctx *parser.LikeQueryAlgorithmContext) *ast.LikeQueryAlgorithm {
	stmt := &ast.LikeQueryAlgorithm{}
	if ctx.AlgorithmDefinition() != nil {
		stmt.AlgorithmDefinition = v.VisitAlgorithmDefinition(ctx.AlgorithmDefinition().(*parser.AlgorithmDefinitionContext))
	}
	return stmt
}

func (v *Visitor) VisitEncryptAlgorithm(ctx *parser.EncryptAlgorithmContext) *ast.EncryptAlgorithm {
	stmt := &ast.EncryptAlgorithm{}
	if ctx.AlgorithmDefinition() != nil {
		stmt.AlgorithmDefinition = v.VisitAlgorithmDefinition(ctx.AlgorithmDefinition().(*parser.AlgorithmDefinitionContext))
	}
	return stmt
}

func (v *Visitor) VisitColumnDefinition(ctx *parser.ColumnDefinitionContext) *ast.ColumnDefinition {
	stmt := &ast.ColumnDefinition{}
	if ctx.ColumnName() != nil {
		stmt.ColumnName = v.VisitColumnName(ctx.ColumnName().(*parser.ColumnNameContext))
	}

	if ctx.DataType() != nil {
		stmt.DataType = v.VisitDataType(ctx.DataType().(*parser.DataTypeContext))
	}
	return stmt
}

func (v *Visitor) VisitColumnName(ctx *parser.ColumnNameContext) *ast.CommonIdentifier {
	stmt := &ast.CommonIdentifier{}
	if ctx.IDENTIFIER_() != nil {
		stmt.Identifier = ctx.IDENTIFIER_().GetText()
	}
	return stmt
}

func (v *Visitor) VisitDataType(ctx *parser.DataTypeContext) *ast.DataType {
	stmt := &ast.DataType{}
	if ctx.STRING_() != nil {
		stmt.String = ctx.STRING_().GetText()
	}
	return stmt
}

func (v *Visitor) VisitResourceDefinition(ctx *parser.ResourceDefinitionContext) *ast.ResourceDefinition {
	stmt := &ast.ResourceDefinition{}
	if ctx.ResourceName() != nil {
		stmt.ResourceName = v.VisitResourceName(ctx.ResourceName().(*parser.ResourceNameContext))
	}
	return stmt
}

func (v *Visitor) VisitResourceName(ctx *parser.ResourceNameContext) *ast.CommonIdentifier {
	stmt := &ast.CommonIdentifier{}
	if ctx.IDENTIFIER_() != nil {
		stmt.Identifier = ctx.IDENTIFIER_().GetText()
	}
	return stmt
}

func (v *Visitor) VisitLiteral(ctx *parser.LiteralContext) *ast.Literal {
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

func (v *Visitor) VisitAlgorithmDefinition(ctx *parser.AlgorithmDefinitionContext) *ast.AlgorithmDefinition {
	stmt := &ast.AlgorithmDefinition{}
	if ctx.AlgorithmTypeName() != nil {
		stmt.AlgorithmTypeName = v.VisitAlgorithmTypeName(ctx.AlgorithmTypeName().(*parser.AlgorithmTypeNameContext))
	}

	if ctx.PropertiesDefinition() != nil {
		stmt.PropertiesDefinition = v.VisitPropertiesDefinition(ctx.PropertiesDefinition().(*parser.PropertiesDefinitionContext))
	}
	return stmt
}

func (v *Visitor) VisitAlgorithmTypeName(ctx *parser.AlgorithmTypeNameContext) *ast.AlgorithmTypeName {
	stmt := &ast.AlgorithmTypeName{}
	switch {
	case ctx.STRING_() != nil:
		stmt.String = ctx.STRING_().GetText()
	case ctx.BuildinAlgorithmTypeName() != nil:
		stmt.BuildinAlgorithmTypeName = v.VisitBuildinAlgorithmTypeName(ctx.BuildinAlgorithmTypeName().(*parser.BuildinAlgorithmTypeNameContext))
	}
	return stmt
}

func (v *Visitor) VisitBuildinAlgorithmTypeName(ctx *parser.BuildinAlgorithmTypeNameContext) *ast.BuildinAlgorithmTypeName {
	stmt := &ast.BuildinAlgorithmTypeName{}
	switch {
	case ctx.MD5() != nil:
		stmt.AlgorithmTypeName = ctx.MD5().GetText()
	case ctx.AES() != nil:
		stmt.AlgorithmTypeName = ctx.AES().GetText()
	case ctx.RC4() != nil:
		stmt.AlgorithmTypeName = ctx.RC4().GetText()
	case ctx.SM3() != nil:
		stmt.AlgorithmTypeName = ctx.SM3().GetText()
	case ctx.SM4() != nil:
		stmt.AlgorithmTypeName = ctx.SM4().GetText()
	case ctx.CHAR_DIGEST_LIKE() != nil:
		stmt.AlgorithmTypeName = ctx.CHAR_DIGEST_LIKE().GetText()
	}
	return stmt
}

func (v *Visitor) VisitPropertiesDefinition(ctx *parser.PropertiesDefinitionContext) *ast.PropertiesDefinition {
	stmt := &ast.PropertiesDefinition{}
	if ctx.Properties() != nil {
		stmt.Properties = v.VisitProperties(ctx.Properties().(*parser.PropertiesContext))
	}
	return stmt
}

func (v *Visitor) VisitProperties(ctx *parser.PropertiesContext) *ast.Properties {
	stmt := &ast.Properties{}
	for _, p := range ctx.AllProperty() {
		stmt.Properties = append(stmt.Properties, v.VisitProperty(p.(*parser.PropertyContext)))
	}
	return stmt
}

func (v *Visitor) VisitProperty(ctx *parser.PropertyContext) *ast.Property {
	stmt := &ast.Property{}
	if ctx.STRING_() != nil {
		stmt.Key = ctx.STRING_().GetText()
	}
	if ctx.Literal() != nil {
		stmt.Literal = v.VisitLiteral(ctx.Literal().(*parser.LiteralContext))
	}
	return stmt
}

func (v *Visitor) VisitTableName(ctx *parser.TableNameContext) *ast.CommonIdentifier {
	stmt := &ast.CommonIdentifier{}
	if ctx.IDENTIFIER_() != nil {
		stmt.Identifier = ctx.IDENTIFIER_().GetText()
	}
	return stmt
}
