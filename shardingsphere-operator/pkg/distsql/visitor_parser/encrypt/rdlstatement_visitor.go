// Code generated from encrypt_g4/RDLStatement.g4 by ANTLR 4.8. DO NOT EDIT.

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

package parser // RDLStatement

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by RDLStatementParser.
type RDLStatementVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by RDLStatementParser#createEncryptRule.
	VisitCreateEncryptRule(ctx *CreateEncryptRuleContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#alterEncryptRule.
	VisitAlterEncryptRule(ctx *AlterEncryptRuleContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#dropEncryptRule.
	VisitDropEncryptRule(ctx *DropEncryptRuleContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#encryptRuleDefinition.
	VisitEncryptRuleDefinition(ctx *EncryptRuleDefinitionContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#resourceDefinition.
	VisitResourceDefinition(ctx *ResourceDefinitionContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#resourceName.
	VisitResourceName(ctx *ResourceNameContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#encryptColumnDefinition.
	VisitEncryptColumnDefinition(ctx *EncryptColumnDefinitionContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#columnDefinition.
	VisitColumnDefinition(ctx *ColumnDefinitionContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#columnName.
	VisitColumnName(ctx *ColumnNameContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#dataType.
	VisitDataType(ctx *DataTypeContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#plainColumnDefinition.
	VisitPlainColumnDefinition(ctx *PlainColumnDefinitionContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#plainColumnName.
	VisitPlainColumnName(ctx *PlainColumnNameContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#cipherColumnDefinition.
	VisitCipherColumnDefinition(ctx *CipherColumnDefinitionContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#cipherColumnName.
	VisitCipherColumnName(ctx *CipherColumnNameContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#assistedQueryColumnDefinition.
	VisitAssistedQueryColumnDefinition(ctx *AssistedQueryColumnDefinitionContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#assistedQueryColumnName.
	VisitAssistedQueryColumnName(ctx *AssistedQueryColumnNameContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#likeQueryColumnDefinition.
	VisitLikeQueryColumnDefinition(ctx *LikeQueryColumnDefinitionContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#likeQueryColumnName.
	VisitLikeQueryColumnName(ctx *LikeQueryColumnNameContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#encryptAlgorithm.
	VisitEncryptAlgorithm(ctx *EncryptAlgorithmContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#assistedQueryAlgorithm.
	VisitAssistedQueryAlgorithm(ctx *AssistedQueryAlgorithmContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#likeQueryAlgorithm.
	VisitLikeQueryAlgorithm(ctx *LikeQueryAlgorithmContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#queryWithCipherColumn.
	VisitQueryWithCipherColumn(ctx *QueryWithCipherColumnContext) interface{}

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

	// Visit a parse tree produced by RDLStatementParser#buildinAlgorithmTypeName.
	VisitBuildinAlgorithmTypeName(ctx *BuildinAlgorithmTypeNameContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#propertiesDefinition.
	VisitPropertiesDefinition(ctx *PropertiesDefinitionContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#properties.
	VisitProperties(ctx *PropertiesContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#property.
	VisitProperty(ctx *PropertyContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#tableName.
	VisitTableName(ctx *TableNameContext) interface{}
}
