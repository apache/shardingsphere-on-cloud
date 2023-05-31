// Code generated from RDLStatement.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // RDLStatement

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseRDLStatementVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseRDLStatementVisitor) VisitCreateMaskRule(ctx *CreateMaskRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitAlterMaskRule(ctx *AlterMaskRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitDropMaskRule(ctx *DropMaskRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitMaskRuleDefinition(ctx *MaskRuleDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitColumnDefinition(ctx *ColumnDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRDLStatementVisitor) VisitColumnName(ctx *ColumnNameContext) interface{} {
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

func (v *BaseRDLStatementVisitor) VisitBuildInMaskAlgorithmType(ctx *BuildInMaskAlgorithmTypeContext) interface{} {
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
