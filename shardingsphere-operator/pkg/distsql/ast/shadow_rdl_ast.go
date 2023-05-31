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

type CreateShadowRule struct {
	IfNotExists             *IfNotExists
	AllShadowRuleDefinition []*ShadowRuleDefinition
}

func (createShadowRule *CreateShadowRule) ToString() string {
	var (
		distSQL = "CREATE SHADOW RULE"
		allRule = []string{}
	)
	if createShadowRule.IfNotExists != nil {
		distSQL = fmt.Sprintf("%s %s", distSQL, createShadowRule.IfNotExists)
	}
	if createShadowRule.AllShadowRuleDefinition != nil {
		for _, r := range createShadowRule.AllShadowRuleDefinition {
			allRule = append(allRule, r.ToString())
		}
	}
	return fmt.Sprintf("%s %s", distSQL, strings.Join(allRule, ","))
}

type ShadowRuleDefinition struct {
	RuleName           *CommonIdentifier
	Source             *CommonIdentifier
	Shadow             *CommonIdentifier
	AllShadowTableRule []*ShadowTableRule
}

func (shadowRuleDefinition *ShadowRuleDefinition) ToString() string {
	var (
		distSQL          = ""
		shadowTableRules = []string{}
	)
	if shadowRuleDefinition.RuleName != nil {
		distSQL = fmt.Sprintf("%s (", shadowRuleDefinition.RuleName.ToString())
	}

	if shadowRuleDefinition.Source != nil {
		distSQL = fmt.Sprintf("%s SOURCE = %s, ", distSQL, shadowRuleDefinition.ToString())
	}

	if shadowRuleDefinition.Shadow != nil {
		distSQL = fmt.Sprintf("%s SHADOW = %s, ", distSQL, shadowRuleDefinition.Shadow.ToString())
	}

	if shadowRuleDefinition.AllShadowTableRule != nil {
		for _, r := range shadowRuleDefinition.AllShadowTableRule {
			shadowTableRules = append(shadowTableRules, r.ToString())
		}
	}
	return fmt.Sprintf("%s %s)", distSQL, strings.Join(shadowTableRules, ","))
}

type ShadowTableRule struct {
	TableName              *CommonIdentifier
	AllAlgorithmDefinition []*AlgorithmDefinition
}

func (shadowTableRule *ShadowTableRule) ToString() string {
	var (
		tableName           = ""
		algorithmDefinition = []string{}
	)
	if shadowTableRule.TableName != nil {
		tableName = shadowTableRule.TableName.ToString()
	}
	if shadowTableRule.AllAlgorithmDefinition != nil {
		for _, algo := range shadowTableRule.AllAlgorithmDefinition {
			algorithmDefinition = append(algorithmDefinition, algo.ToString())
		}
	}
	return fmt.Sprintf("%s (%s)", tableName, strings.Join(algorithmDefinition, ","))
}

type AlterShadowRule struct {
	AllShadowRuleDefinition []*ShadowRuleDefinition
}

func (alterShadowRule *AlterShadowRule) ToString() string {
	var allrules []string
	if alterShadowRule.AllShadowRuleDefinition != nil {
		for _, r := range alterShadowRule.AllShadowRuleDefinition {
			allrules = append(allrules, r.ToString())
		}
	}
	return fmt.Sprintf("ALTER SHADOW RULE %s", strings.Join(allrules, ","))
}

type DropShadowRule struct {
	IfExists    *IfExists
	AllRuleName []*CommonIdentifier
}

func (dropShadowRule *DropShadowRule) ToString() string {
	var (
		ifExists    = ""
		allRuleName = []string{}
	)
	if dropShadowRule.IfExists != nil {
		ifExists = dropShadowRule.IfExists.ToString()
	}
	for _, r := range dropShadowRule.AllRuleName {
		allRuleName = append(allRuleName, r.ToString())
	}
	return fmt.Sprintf("DROP SHADOW RULE %s %s", ifExists, strings.Join(allRuleName, ","))
}

type DropShadowAlgorithm struct {
	IfExists         *IfExists
	AllAlgorithmName []*CommonIdentifier
}

func (dropShadowAlgorithm *DropShadowAlgorithm) ToString() string {
	var (
		ifExists         = ""
		allAlgorithmName = []string{}
	)

	if dropShadowAlgorithm.IfExists != nil {
		ifExists = dropShadowAlgorithm.IfExists.ToString()
	}

	if dropShadowAlgorithm.AllAlgorithmName != nil {
		for _, algo := range dropShadowAlgorithm.AllAlgorithmName {
			allAlgorithmName = append(allAlgorithmName, algo.ToString())
		}
	}

	return fmt.Sprintf("DROP SHADOW ALGORITHM %s %s", ifExists, strings.Join(allAlgorithmName, ","))
}

type CreateDefaultShadowAlgorithm struct {
	IfNotExists         *IfNotExists
	AlgorithmDefinition *AlgorithmDefinition
}

func (createDefaultShadowAlgorithm *CreateDefaultShadowAlgorithm) ToString() string {
	return fmt.Sprintf("CREATE DEFAULT SHADOW ALGORITHM %s %s", createDefaultShadowAlgorithm.IfNotExists.ToString(), createDefaultShadowAlgorithm.AlgorithmDefinition.ToString())
}

type DropDefaultShadowAlgorithm struct {
	IfExists *IfExists
}

func (dropDefaultShadowAlgorithm *DropDefaultShadowAlgorithm) ToString() string {
	return fmt.Sprintf("DROP DEFAULT SHADOW ALGORITHM %s", dropDefaultShadowAlgorithm.IfExists.ToString())
}

type AlterDefaultShadowAlgorithm struct {
	AlgorithmDefinition *AlgorithmDefinition
}

func (alterDefaultShadowAlgorithm *AlterDefaultShadowAlgorithm) ToString() string {
	return fmt.Sprintf("ALTER DEFAULT SHADOW ALGORITHM %s", alterDefaultShadowAlgorithm.AlgorithmDefinition.ToString())
}
