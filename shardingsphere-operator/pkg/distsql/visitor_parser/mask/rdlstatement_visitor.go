// Code generated from RDLStatement.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // RDLStatement

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by RDLStatementParser.
type RDLStatementVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by RDLStatementParser#createMaskRule.
	VisitCreateMaskRule(ctx *CreateMaskRuleContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#alterMaskRule.
	VisitAlterMaskRule(ctx *AlterMaskRuleContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#dropMaskRule.
	VisitDropMaskRule(ctx *DropMaskRuleContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#maskRuleDefinition.
	VisitMaskRuleDefinition(ctx *MaskRuleDefinitionContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#columnDefinition.
	VisitColumnDefinition(ctx *ColumnDefinitionContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#columnName.
	VisitColumnName(ctx *ColumnNameContext) interface{}

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

	// Visit a parse tree produced by RDLStatementParser#buildInMaskAlgorithmType.
	VisitBuildInMaskAlgorithmType(ctx *BuildInMaskAlgorithmTypeContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#propertiesDefinition.
	VisitPropertiesDefinition(ctx *PropertiesDefinitionContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#properties.
	VisitProperties(ctx *PropertiesContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#property.
	VisitProperty(ctx *PropertyContext) interface{}

	// Visit a parse tree produced by RDLStatementParser#ruleName.
	VisitRuleName(ctx *RuleNameContext) interface{}
}
