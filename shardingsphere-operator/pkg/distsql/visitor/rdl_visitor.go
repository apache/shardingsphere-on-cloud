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
	"atlt/encrypt_visitor/ast"
	parser "atlt/encrypt_visitor_parser"
	"fmt"
)

type Visitor struct {
	parser.BaseRDLStatementVisitor
}

func (v *Visitor) VisitCreateEncryptRule(ctx *parser.CreateEncryptRuleContext) *ast.CreateEncryptRule {
	stmt := &ast.CreateEncryptRule{}
	stmt.Create = ctx.CREATE().GetText()
	stmt.Encrypt = ctx.ENCRYPT().GetText()
	stmt.EncryptName = ctx.RULE().GetText()

	if ctx.IfNotExists() != nil {
		stmt.IfNotExists = v.VisitIfNotExists(ctx.IfNotExists().(*parser.IfNotExistsContext))
	}

	if ctx.AllEncryptRuleDefinition() != nil {
		for _, r := range ctx.AllEncryptRuleDefinition() {
			v.VisitEncryptRuleDefinition(r.(*parser.EncryptRuleDefinitionContext))
		}
	}

	return stmt
}

func (v *Visitor) VisitAlterEncryptRule(ctx *parser.AlterEncryptRuleContext) interface{} {
	if ctx.AllEncryptRuleDefinition() != nil {
		for _, encryptRuleDefinition := range ctx.AllEncryptRuleDefinition() {
			v.VisitEncryptRuleDefinition(encryptRuleDefinition.(*parser.EncryptRuleDefinitionContext))
		}
	}
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitDropEncryptRule(ctx *parser.DropEncryptRuleContext) interface{} {
	if ctx.IfExists() != nil {
		v.VisitIfExists(ctx.IfExists().(*parser.IfExistsContext))
	}

	if ctx.AllTableName() != nil {
		for _, tableName := range ctx.AllTableName() {
			v.VisitTableName(tableName.(*parser.TableNameContext))
		}
	}
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitIfNotExists(ctx *parser.IfNotExistsContext) *ast.IfNotExists {
	return &ast.IfNotExists{
		IfNotExists: fmt.Sprintf("%s %s %s", ctx.IF().GetText(), ctx.NOT().GetText(), ctx.EXISTS().GetText()),
	}
}

func (v *Visitor) VisitEncryptRuleDefinition(ctx *parser.EncryptRuleDefinitionContext) interface{} {
	// TODO: get table name set AST with ctx.TableName().GetTxt()

	if ctx.ResourceDefinition() != nil {
		v.VisitResourceDefinition(ctx.ResourceDefinition().(*parser.ResourceDefinitionContext))
	}

	if ctx.AllEncryptColumnDefinition() != nil {
		for _, column := range ctx.AllEncryptColumnDefinition() {
			v.VisitEncryptColumnDefinition(column.(*parser.EncryptColumnDefinitionContext))
		}
	}

	if ctx.QueryWithCipherColumn() != nil {
		v.VisitQueryWithCipherColumn(ctx.QueryWithCipherColumn().(*parser.QueryWithCipherColumnContext))
	}

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitQueryWithCipherColumn(ctx *parser.QueryWithCipherColumnContext) interface{} {
	switch {
	case ctx.TRUE() != nil:
		//TODO: set ctx.TRUE().GetText() to AST
	case ctx.FALSE() != nil:
		//TODO: set ctx.FALSE().GetText() to AST
	}
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitEncryptColumnDefinition(ctx *parser.EncryptColumnDefinitionContext) interface{} {

	if ctx.ColumnDefinition() != nil {
		v.VisitColumnDefinition(ctx.ColumnDefinition().(*parser.ColumnDefinitionContext))
	}

	if ctx.PlainColumnDefinition() != nil {
		v.VisitPlainColumnDefinition(ctx.PlainColumnDefinition().(*parser.PlainColumnDefinitionContext))
	}

	if ctx.CipherColumnDefinition() != nil {
		v.VisitCipherColumnDefinition(ctx.CipherColumnDefinition().(*parser.CipherColumnDefinitionContext))
	}

	if ctx.AssistedQueryColumnDefinition() != nil {
		v.VisitAssistedQueryColumnDefinition(ctx.AssistedQueryColumnDefinition().(*parser.AssistedQueryColumnDefinitionContext))
	}

	if ctx.LikeQueryColumnDefinition() != nil {
		v.VisitLikeQueryColumnDefinition(ctx.LikeQueryColumnDefinition().(*parser.LikeQueryColumnDefinitionContext))
	}

	if ctx.EncryptAlgorithm() != nil {
		v.VisitEncryptAlgorithm(ctx.EncryptAlgorithm().(*parser.EncryptAlgorithmContext))
	}

	if ctx.AssistedQueryAlgorithm() != nil {
		v.VisitAssistedQueryAlgorithm(ctx.AssistedQueryAlgorithm().(*parser.AssistedQueryAlgorithmContext))
	}

	if ctx.LikeQueryAlgorithm() != nil {
		v.VisitLikeQueryAlgorithm(ctx.LikeQueryAlgorithm().(*parser.LikeQueryAlgorithmContext))
	}

	if ctx.QueryWithCipherColumn() != nil {
		v.VisitQueryWithCipherColumn(ctx.QueryWithCipherColumn().(*parser.QueryWithCipherColumnContext))
	}

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPlainColumnDefinition(ctx *parser.PlainColumnDefinitionContext) interface{} {
	if ctx.PlainColumnName() != nil {
		v.VisitPlainColumnName(ctx.PlainColumnName().(*parser.PlainColumnNameContext))
	}

	if ctx.DataType() != nil {
		v.VisitDataType(ctx.DataType().(*parser.DataTypeContext))
	}
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPlainColumnName(ctx *parser.PlainColumnNameContext) interface{} {
	if ctx.IDENTIFIER_() != nil {
		//TODO:set ctx.IDENTIFIER_().GetText() to AST
	}
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitCipherColumnDefinition(ctx *parser.CipherColumnDefinitionContext) interface{} {
	if ctx.CipherColumnName() != nil {
		v.VisitCipherColumnName(ctx.CipherColumnName().(*parser.CipherColumnNameContext))
	}

	if ctx.DataType() != nil {
		v.VisitDataType(ctx.DataType().(*parser.DataTypeContext))
	}

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitCipherColumnName(ctx *parser.CipherColumnNameContext) interface{} {
	if ctx.IDENTIFIER_() != nil {
		//TODO: set ctx.IDENTIFIER_().GetText() to AST
	}
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAssistedQueryColumnDefinition(ctx *parser.AssistedQueryColumnDefinitionContext) interface{} {
	if ctx.AssistedQueryColumnName() != nil {
		v.VisitAssistedQueryColumnName(ctx.AssistedQueryColumnName().(*parser.AssistedQueryColumnNameContext))
	}

	if ctx.DataType() != nil {
		v.VisitDataType(ctx.DataType().(*parser.DataTypeContext))
	}

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAssistedQueryColumnName(ctx *parser.AssistedQueryColumnNameContext) interface{} {
	if ctx.IDENTIFIER_() != nil {
		//TODO: set ctx.IDENTIFIER_().GetText() to AST
	}
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitLikeQueryColumnDefinition(ctx *parser.LikeQueryColumnDefinitionContext) interface{} {
	if ctx.LikeQueryColumnName() != nil {
		v.VisitLikeQueryColumnName(ctx.LikeQueryColumnName().(*parser.LikeQueryColumnNameContext))
	}

	if ctx.DataType() != nil {
		v.VisitDataType(ctx.DataType().(*parser.DataTypeContext))
	}
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitLikeQueryColumnName(ctx *parser.LikeQueryColumnNameContext) interface{} {
	if ctx.IDENTIFIER_() != nil {
		//TODO: set ctx.IDENTIFIER_() to AST
	}
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAssistedQueryAlgorithm(ctx *parser.AssistedQueryAlgorithmContext) interface{} {
	if ctx.AlgorithmDefinition() != nil {
		v.VisitAlgorithmDefinition(ctx.AlgorithmDefinition().(*parser.AlgorithmDefinitionContext))
	}
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitLikeQueryAlgorithm(ctx *parser.LikeQueryAlgorithmContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitEncryptAlgorithm(ctx *parser.EncryptAlgorithmContext) interface{} {
	if ctx.AlgorithmDefinition() != nil {
		v.VisitAlgorithmDefinition(ctx.AlgorithmDefinition().(*parser.AlgorithmDefinitionContext))
	}
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitColumnDefinition(ctx *parser.ColumnDefinitionContext) interface{} {
	if ctx.ColumnName() != nil {
		v.VisitColumnName(ctx.ColumnName().(*parser.ColumnNameContext))
	}

	if ctx.DataType() != nil {
		v.VisitDataType(ctx.DataType().(*parser.DataTypeContext))
	}
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitColumnName(ctx *parser.ColumnNameContext) interface{} {
	if ctx.IDENTIFIER_() != nil {
		// TODO: set ctx.IDENTIFIER_().GetText() to AST
	}
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitDataType(ctx *parser.DataTypeContext) interface{} {
	if ctx.STRING_() != nil {
		//TODO: set ctx.STRING_().GetText() to AST
	}
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitResourceDefinition(ctx *parser.ResourceDefinitionContext) interface{} {
	fmt.Println("rrrr >>> ", ctx.ResourceName().GetText())
	if ctx.ResourceName() != nil {
		v.VisitResourceName(ctx.ResourceName().(*parser.ResourceNameContext))
	}
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitResourceName(ctx *parser.ResourceNameContext) interface{} {
	if ctx.IDENTIFIER_() != nil {
		// TODO: set ctx.IDENTIFIER_().GetText() to AST
	}
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	switch {
	case ctx.STRING_() != nil:
		// TODO: set ctx.STRING_().GetText() to AST
	case ctx.MINUS_() != nil:
		// TODO: set ctx.MINUS_().GetText() to AST
	case ctx.INT_() != nil:
		// TODO: set ctx.INT_().GetText() to AST
	case ctx.TRUE() != nil:
		// TODO: set ctx.TRUE().GetText() to AST
	case ctx.FALSE() != nil:
		// TODO: set ctx.FALSE().GetText() to AST
	}
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAlgorithmDefinition(ctx *parser.AlgorithmDefinitionContext) interface{} {
	if ctx.AlgorithmTypeName() != nil {
		v.VisitAlgorithmTypeName(ctx.AlgorithmTypeName().(*parser.AlgorithmTypeNameContext))
	}

	if ctx.PropertiesDefinition() != nil {
		v.VisitPropertiesDefinition(ctx.PropertiesDefinition().(*parser.PropertiesDefinitionContext))
	}
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAlgorithmTypeName(ctx *parser.AlgorithmTypeNameContext) interface{} {
	switch {
	case ctx.STRING_() != nil:
		//TODO: set ctx.STRING_().GetText() to AST
	case ctx.BuildinAlgorithmTypeName() != nil:
		v.VisitBuildinAlgorithmTypeName(ctx.BuildinAlgorithmTypeName().(*parser.BuildinAlgorithmTypeNameContext))
	}
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitBuildinAlgorithmTypeName(ctx *parser.BuildinAlgorithmTypeNameContext) interface{} {
	switch {
	case ctx.MD5() != nil:
		//TODO: set ctx.MD5().GetText() to AST
	case ctx.AES() != nil:
		//TODO: set ctx.AES().GetText() to AST
	case ctx.RC4() != nil:
		//TODO: set ctx.RC4().GetText() to AST
	case ctx.SM3() != nil:
		//TODO: set ctx.SM3().GetText() to AST
	case ctx.SM4() != nil:
		//TODO: set ctx.SM4().GetText() to AST
	case ctx.CHAR_DIGEST_LIKE() != nil:
		//TODO: set ctx.CHAR_DIGEST_LIKE().GetText() to AST
	}
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPropertiesDefinition(ctx *parser.PropertiesDefinitionContext) interface{} {
	if ctx.Properties() != nil {
		v.VisitProperties(ctx.Properties().(*parser.PropertiesContext))
	}
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitProperties(ctx *parser.PropertiesContext) interface{} {
	for _, p := range ctx.AllProperty() {
		v.VisitProperty(p.(*parser.PropertyContext))
	}
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitProperty(ctx *parser.PropertyContext) interface{} {
	if ctx.STRING_() != nil {
		//TODO: set ctx.STRING_() to AST
	}
	if ctx.Literal() != nil {
		v.VisitLiteral(ctx.Literal().(*parser.LiteralContext))
	}
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitTableName(ctx *parser.TableNameContext) interface{} {
	if ctx.IDENTIFIER_() != nil {
		// TODO: set ctx.IDENTIFIER_().GetText() to AST
	}
	return v.VisitChildren(ctx)
}
