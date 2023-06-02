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

type CreateReadwriteSplittingRule struct {
	IfNotExists                         *IfNotExists
	AllReadwriteSplittingRuleDefinition []*ReadWriteSplittingRuleDefinition
}

func (createReadwriteSplittingRule *CreateReadwriteSplittingRule) ToString() string {
	var (
		ifNotExists                      string
		readwriteSplittingRuleDefinition []string
	)
	if createReadwriteSplittingRule.IfNotExists != nil {
		ifNotExists = createReadwriteSplittingRule.IfNotExists.ToString()
	}
	if createReadwriteSplittingRule.AllReadwriteSplittingRuleDefinition != nil {
		for _, r := range createReadwriteSplittingRule.AllReadwriteSplittingRuleDefinition {
			readwriteSplittingRuleDefinition = append(readwriteSplittingRuleDefinition, r.ToString())
		}
	}
	return fmt.Sprintf("CREATE READWRITE_SPLITTING RULE %s %s", ifNotExists, strings.Join(readwriteSplittingRuleDefinition, ","))
}

type ReadWriteSplittingRuleDefinition struct {
	RuleName                       *CommonIdentifier
	DataSourceDefinition           *DataSourceDefinition
	TransactionalReadQueryStrategy *TransactionalReadQueryStrategy
	AlgorithmDefinition            *AlgorithmDefinition
}

func (readWriteSplittingRuleDefinition *ReadWriteSplittingRuleDefinition) ToString() string {
	var (
		ruleName                       string
		dataSourceDefinition           string
		transactionalReadQueryStrategy string
		algorithmDefinition            string
	)

	if readWriteSplittingRuleDefinition.RuleName != nil {
		ruleName = readWriteSplittingRuleDefinition.RuleName.ToString()
	}

	if readWriteSplittingRuleDefinition.DataSourceDefinition != nil {
		dataSourceDefinition = readWriteSplittingRuleDefinition.DataSourceDefinition.ToString()
	}

	if readWriteSplittingRuleDefinition.TransactionalReadQueryStrategy != nil {
		transactionalReadQueryStrategy = readWriteSplittingRuleDefinition.TransactionalReadQueryStrategy.ToString()
	}

	if readWriteSplittingRuleDefinition.AlgorithmDefinition != nil {
		algorithmDefinition = readWriteSplittingRuleDefinition.AlgorithmDefinition.ToString()
	}

	return fmt.Sprintf("%s (%s %s %s)", ruleName, dataSourceDefinition, transactionalReadQueryStrategy, algorithmDefinition)
}

type DataSourceDefinition struct {
	WriteStorageUnit *WriteStorageUnit
	ReadStorageUnits *ReadStorageUnits
}

func (dataSourceDefinition DataSourceDefinition) ToString() string {
	return fmt.Sprintf("%s, %s", dataSourceDefinition.WriteStorageUnit.ToString(), dataSourceDefinition.ReadStorageUnits.ToString())
}

type WriteStorageUnit struct {
	WriteStorageUnitName *WriteStorageUnitName
}

func (writeStorageUnit *WriteStorageUnit) ToString() string {
	return fmt.Sprintf("WRITE_STORAGE_UNIT = %s", writeStorageUnit.WriteStorageUnitName.ToString())
}

type WriteStorageUnitName struct {
	StorageUnitName *CommonIdentifier
}

func (writeStorageUnitName *WriteStorageUnitName) ToString() string {
	return writeStorageUnitName.StorageUnitName.ToString()
}

type ReadStorageUnits struct {
	ReadStorageUnitsNames *ReadStorageUnitsNames
}

func (readStorageUnits *ReadStorageUnits) ToString() string {
	return fmt.Sprintf("READ_STORAGE_UNITS (%s)", readStorageUnits.ReadStorageUnitsNames.ToString())
}

type ReadStorageUnitsNames struct {
	AllStorageUnitName []*CommonIdentifier
}

func (readStorageUnitsNames *ReadStorageUnitsNames) ToString() string {
	var (
		storageUnitNames []string
	)
	for _, s := range readStorageUnitsNames.AllStorageUnitName {
		storageUnitNames = append(storageUnitNames, s.ToString())
	}
	return strings.Join(storageUnitNames, ",")
}

type TransactionalReadQueryStrategy struct {
	TransactionalReadQueryStrategyName *TransactionalReadQueryStrategyName
}

func (transactionalReadQueryStrategy *TransactionalReadQueryStrategy) ToString() string {
	return transactionalReadQueryStrategy.TransactionalReadQueryStrategyName.ToString()
}

type AlterReadwriteSplittingRule struct {
	AllReadwriteSplittingRuleDefinition []*ReadWriteSplittingRuleDefinition
}

func (alterReadwriteSplittingRule *AlterReadwriteSplittingRule) ToString() string {
	var (
		allReadwriteSplittingRuleDefinition []string
	)
	for _, r := range alterReadwriteSplittingRule.AllReadwriteSplittingRuleDefinition {
		allReadwriteSplittingRuleDefinition = append(allReadwriteSplittingRuleDefinition, r.ToString())
	}
	return fmt.Sprintf("ALTER READWRITE_SPLITTING RULE %s", strings.Join(allReadwriteSplittingRuleDefinition, ","))
}

type DropReadwriteSplittingRule struct {
	IfExists    *IfExists
	AllRuleName []*CommonIdentifier
}

func (dropReadwriteSplittingRule *DropReadwriteSplittingRule) ToString() string {
	var (
		ifExists    string
		allRuleName []string
	)
	if dropReadwriteSplittingRule.IfExists != nil {
		ifExists = dropReadwriteSplittingRule.IfExists.ToString()
	}

	if dropReadwriteSplittingRule.AllRuleName != nil {
		for _, r := range dropReadwriteSplittingRule.AllRuleName {
			allRuleName = append(allRuleName, r.ToString())
		}
	}
	return fmt.Sprintf("DROP READWRITE_SPLITTING RULE %s %s", ifExists, strings.Join(allRuleName, ","))
}

type TransactionalReadQueryStrategyName struct {
	String string
}

func (transactionalReadQueryStrategyName *TransactionalReadQueryStrategyName) ToString() string {
	return transactionalReadQueryStrategyName.String
}
