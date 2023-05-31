// Code generated from RDLStatement.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // RDLStatement

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseRDLStatementVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseRDLStatementVisitor) VisitCreateShadowRule(ctx *CreateShadowRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitAlterShadowRule(ctx *AlterShadowRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitDropShadowRule(ctx *DropShadowRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitDropShadowAlgorithm(ctx *DropShadowAlgorithmContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitCreateDefaultShadowAlgorithm(ctx *CreateDefaultShadowAlgorithmContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitDropDefaultShadowAlgorithm(ctx *DropDefaultShadowAlgorithmContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitAlterDefaultShadowAlgorithm(ctx *AlterDefaultShadowAlgorithmContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitShadowRuleDefinition(ctx *ShadowRuleDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitShadowTableRule(ctx *ShadowTableRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitSource(ctx *SourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitShadow(ctx *ShadowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitTableName(ctx *TableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitAlgorithmName(ctx *AlgorithmNameContext) interface{} {
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

func (v *BaseRDLStatementVisitor) VisitBuildInShadowAlgorithmType(ctx *BuildInShadowAlgorithmTypeContext) interface{} {
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

func (v *BaseRDLStatementVisitor) VisitRuleName(ctx *RuleNameContext) interface{} {
	return v.VisitChildren(ctx)
}
