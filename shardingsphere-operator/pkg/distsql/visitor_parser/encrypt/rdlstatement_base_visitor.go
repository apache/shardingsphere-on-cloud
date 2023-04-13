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

type BaseRDLStatementVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseRDLStatementVisitor) VisitCreateEncryptRule(ctx *CreateEncryptRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitAlterEncryptRule(ctx *AlterEncryptRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitDropEncryptRule(ctx *DropEncryptRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitEncryptRuleDefinition(ctx *EncryptRuleDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitResourceDefinition(ctx *ResourceDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitResourceName(ctx *ResourceNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitEncryptColumnDefinition(ctx *EncryptColumnDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitColumnDefinition(ctx *ColumnDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitColumnName(ctx *ColumnNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitDataType(ctx *DataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitPlainColumnDefinition(ctx *PlainColumnDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitPlainColumnName(ctx *PlainColumnNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitCipherColumnDefinition(ctx *CipherColumnDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitCipherColumnName(ctx *CipherColumnNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitAssistedQueryColumnDefinition(ctx *AssistedQueryColumnDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitAssistedQueryColumnName(ctx *AssistedQueryColumnNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitLikeQueryColumnDefinition(ctx *LikeQueryColumnDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitLikeQueryColumnName(ctx *LikeQueryColumnNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitEncryptAlgorithm(ctx *EncryptAlgorithmContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitAssistedQueryAlgorithm(ctx *AssistedQueryAlgorithmContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitLikeQueryAlgorithm(ctx *LikeQueryAlgorithmContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitQueryWithCipherColumn(ctx *QueryWithCipherColumnContext) interface{} {
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

func (v *BaseRDLStatementVisitor) VisitBuildinAlgorithmTypeName(ctx *BuildinAlgorithmTypeNameContext) interface{} {
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
