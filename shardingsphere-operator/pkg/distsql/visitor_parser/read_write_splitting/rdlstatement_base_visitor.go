// Code generated from RDLStatement.g4 by ANTLR 4.8. DO NOT EDIT.

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

func (v *BaseRDLStatementVisitor) VisitCreateReadwriteSplittingRule(ctx *CreateReadwriteSplittingRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitAlterReadwriteSplittingRule(ctx *AlterReadwriteSplittingRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitDropReadwriteSplittingRule(ctx *DropReadwriteSplittingRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitReadwriteSplittingRuleDefinition(ctx *ReadwriteSplittingRuleDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitDataSourceDefinition(ctx *DataSourceDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitRuleName(ctx *RuleNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitWriteStorageUnit(ctx *WriteStorageUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitReadStorageUnits(ctx *ReadStorageUnitsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitTransactionalReadQueryStrategy(ctx *TransactionalReadQueryStrategyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitWriteStorageUnitName(ctx *WriteStorageUnitNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitReadStorageUnitsNames(ctx *ReadStorageUnitsNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitTransactionalReadQueryStrategyName(ctx *TransactionalReadQueryStrategyNameContext) interface{} {
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

func (v *BaseRDLStatementVisitor) VisitBuildInReadQueryLoadBalanceAlgorithmType(ctx *BuildInReadQueryLoadBalanceAlgorithmTypeContext) interface{} {
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

func (v *BaseRDLStatementVisitor) VisitDatabaseName(ctx *DatabaseNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitGroupName(ctx *GroupNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitStorageUnitName(ctx *StorageUnitNameContext) interface{} {
	return v.VisitChildren(ctx)
}
