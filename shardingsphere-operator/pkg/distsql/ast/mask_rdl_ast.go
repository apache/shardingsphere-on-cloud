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

package ast

import (
	"fmt"
	"strings"
)

type CreateMaskRule struct {
	AllMaskRuleDefinition []*MaskRuleDefinition
	Table                 string
	IfNotExists           *IfNotExists
}

func (createMaskRule *CreateMaskRule) ToString() string {
	var (
		distSQL            = "CREATE MASK"
		maskRuleDefinition = []string{}
	)

	if createMaskRule.Table != "" {
		distSQL = fmt.Sprintf("%s %s", distSQL, createMaskRule.Table)
	}

	if createMaskRule.IfNotExists != nil {
		distSQL = fmt.Sprintf("%s %s", distSQL, createMaskRule.IfNotExists.ToString())
	}

	if createMaskRule.AllMaskRuleDefinition != nil {
		for _, mr := range createMaskRule.AllMaskRuleDefinition {
			maskRuleDefinition = append(maskRuleDefinition, mr.ToString())
		}
	}
	return fmt.Sprintf("%s %s", distSQL, strings.Join(maskRuleDefinition, ","))
}

type MaskRuleDefinition struct {
	RuleName         *CommonIdentifier
	ColumnDefinition []*ColumnDefinition
}

func (maskRuleDefinition *MaskRuleDefinition) ToString() string {
	var (
		columnDefinition []string
	)

	if maskRuleDefinition.ColumnDefinition != nil {
		for _, cd := range maskRuleDefinition.ColumnDefinition {
			columnDefinition = append(columnDefinition, cd.ToString())
		}
	}

	return fmt.Sprintf("%s (( %s ))", maskRuleDefinition.RuleName.ToString(), strings.Join(columnDefinition, ","))
}

type AlterMaskRule struct {
	Table                 string
	AllMaskRuleDefinition []*MaskRuleDefinition
}

func (alterMaskRule *AlterMaskRule) ToString() string {
	var (
		distSQL         = fmt.Sprintf("ALTER MASK %s RULE", alterMaskRule.Table)
		ruleDefinitions []string
	)
	if alterMaskRule.AllMaskRuleDefinition != nil {
		for _, rule := range alterMaskRule.AllMaskRuleDefinition {
			ruleDefinitions = append(ruleDefinitions, rule.ToString())
		}
	}
	return fmt.Sprintf("%s %s", distSQL, strings.Join(ruleDefinitions, ","))
}

type DropMaskRule struct {
	IfExists    *IfExists
	Table       string
	AllRuleName []*CommonIdentifier
}

func (dropMaskRule *DropMaskRule) ToString() string {
	var (
		distSQL   = "DROP MASK"
		ruleNames []string
	)
	if dropMaskRule.Table != "" {
		distSQL = fmt.Sprintf("%s %s", distSQL, dropMaskRule.Table)
	}
	if dropMaskRule.IfExists != nil {
		distSQL = fmt.Sprintf("%s %s", distSQL, dropMaskRule.IfExists.ToString())
	}
	if dropMaskRule.AllRuleName != nil {
		for _, ruleName := range dropMaskRule.AllRuleName {
			ruleNames = append(ruleNames, ruleName.ToString())
		}
	}
	return fmt.Sprintf("%s %s", distSQL, strings.Join(ruleNames, ","))
}
