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

// Define RDL AST
type CreateEncryptRule struct {
	IfNotExists              *IfNotExists
	EncryptRuleDefinition    *EncryptRuleDefinition
	AllEncryptRuleDefinition []*EncryptRuleDefinition
}

func (createEncryptRule *CreateEncryptRule) ToString() string {
	var ifNotExists string
	var allEncryptRuleDefinitionList []string
	if createEncryptRule.IfNotExists != nil {
		ifNotExists = createEncryptRule.IfNotExists.ToString()
	}

	if createEncryptRule.AllEncryptRuleDefinition != nil {
		for _, encryptRuleDefinition := range createEncryptRule.AllEncryptRuleDefinition {
			if encryptRuleDefinition != nil {
				allEncryptRuleDefinitionList = append(allEncryptRuleDefinitionList, encryptRuleDefinition.ToString())
			}
		}
	}
	return fmt.Sprintf("CREATE ENCRYPT RULE%s %s;", ifNotExists, strings.Join(allEncryptRuleDefinitionList, ","))
}

type AlterEncryptRule struct {
	EncryptRuleDefinition        *EncryptRuleDefinition
	AllEncryptRuleDefinitionList []*EncryptRuleDefinition
}

func (alterEncryptRule *AlterEncryptRule) ToString() string {
	var encryptRuleDefinition string
	var encryptRuleDefinitionList []string
	if alterEncryptRule.EncryptRuleDefinition != nil {
		encryptRuleDefinition = alterEncryptRule.EncryptRuleDefinition.ToString()
	}
	if alterEncryptRule.AllEncryptRuleDefinitionList != nil {
		for _, encryptRuleDefinition := range alterEncryptRule.AllEncryptRuleDefinitionList {
			encryptRuleDefinitionList = append(encryptRuleDefinitionList, encryptRuleDefinition.ToString())
		}
	}
	return fmt.Sprintf("ALTER ENCRYPT RULE %s %s;", encryptRuleDefinition, strings.Join(encryptRuleDefinitionList, ","))
}

type DropEncryptRule struct {
	IfExists     *IfExists
	AllTableName []*CommonIdentifier
}

func (dropEncryptRule *DropEncryptRule) ToString() string {
	var (
		ifExists     string
		allTableName []string
	)
	if dropEncryptRule.IfExists != nil {
		ifExists = dropEncryptRule.IfExists.ToString()
	}
	if dropEncryptRule.AllTableName != nil {
		for _, tableName := range dropEncryptRule.AllTableName {
			allTableName = append(allTableName, tableName.ToString())
		}
	}
	return fmt.Sprintf("DROP ENCRYPT RULE %s %s;", ifExists, strings.Join(allTableName, ","))
}

type IfExists struct {
	IfExists string
}

func (ifExists *IfExists) ToString() string {
	return ifExists.IfExists
}

type EncryptRuleDefinition struct {
	TableName                  *CommonIdentifier
	ResourceDefinition         *ResourceDefinition
	EncryptColumnDefinition    *EncryptColumnDefinition
	AllEncryptColumnDefinition []*EncryptColumnDefinition
	QueryWithCipherColumn      *QueryWithCipherColumn
}

func (encryptRuleDefinition *EncryptRuleDefinition) ToString() string {
	var (
		tableName                  string
		resourceDefinition         string
		queryWithCipherColumn      string
		encryptColumnDefinition    string
		allEncryptColumnDefinition []string
	)

	if encryptRuleDefinition.TableName != nil {
		tableName = encryptRuleDefinition.TableName.ToString()
	}

	if encryptRuleDefinition.ResourceDefinition != nil {
		resourceDefinition = encryptRuleDefinition.ResourceDefinition.ToString()
	}

	if encryptRuleDefinition.EncryptColumnDefinition != nil {
		encryptColumnDefinition = encryptRuleDefinition.EncryptColumnDefinition.ToString()
	}

	if encryptRuleDefinition.AllEncryptColumnDefinition != nil {
		for _, rd := range encryptRuleDefinition.AllEncryptColumnDefinition {
			allEncryptColumnDefinition = append(allEncryptColumnDefinition, rd.ToString())
		}
	}

	if encryptRuleDefinition.QueryWithCipherColumn != nil {
		queryWithCipherColumn = fmt.Sprintf(",QUERY_WITH_CIPHER_COLUMN=%s", encryptRuleDefinition.QueryWithCipherColumn.ToString())
	}

	return fmt.Sprintf("%s (%sCOLUMNS(%s%s)%s)",
		tableName,
		resourceDefinition,
		encryptColumnDefinition,
		strings.Join(allEncryptColumnDefinition, ","),
		queryWithCipherColumn)
}

type IfNotExists struct {
	IfNotExists string
}

func (ifNotExists IfNotExists) ToString() string {
	return ifNotExists.IfNotExists
}

type ResourceDefinition struct {
	ResourceName *CommonIdentifier
}

func (resourceDefinition *ResourceDefinition) ToString() string {
	if resourceDefinition.ResourceName != nil {
		return fmt.Sprintf("RESOURCE=%s", resourceDefinition.ResourceName.ToString())
	}
	return ""
}

type EncryptColumnDefinition struct {
	ColumnDefinition              *ColumnDefinition
	PlainColumnDefinition         *PlainColumnDefinition
	CipherColumnDefinition        *CipherColumnDefinition
	AssistedQueryColumnDefinition *AssistedQueryColumnDefinition
	LikeQueryColumnDefinition     *LikeQueryColumnDefinition
	EncryptAlgorithm              *EncryptAlgorithm
	AssistedQueryAlgorithm        *AssistedQueryAlgorithm
	LikeQueryAlgorithm            *LikeQueryAlgorithm
	QueryWithCipherColumn         *QueryWithCipherColumn
}

func (encryptColumnDefinition *EncryptColumnDefinition) ToString() string {
	var (
		plainColumnDefinition         string
		cipherColumnDefinition        string
		assistedQueryColumnDefinition string
		likeQueryColumnDefinition     string
		assistedQueryAlgorithm        string
		likeQueryAlgorithm            string
		queryWithCipherColumn         string
	)

	if encryptColumnDefinition.PlainColumnDefinition != nil {
		plainColumnDefinition = fmt.Sprintf(",%s", encryptColumnDefinition.PlainColumnDefinition.ToString())
	}

	if encryptColumnDefinition.CipherColumnDefinition != nil {
		cipherColumnDefinition = encryptColumnDefinition.CipherColumnDefinition.ToString()
	}

	if encryptColumnDefinition.AssistedQueryColumnDefinition != nil {
		assistedQueryColumnDefinition = fmt.Sprintf(",%s", encryptColumnDefinition.AssistedQueryColumnDefinition.ToString())
	}

	if encryptColumnDefinition.LikeQueryAlgorithm != nil {
		likeQueryColumnDefinition = fmt.Sprintf(",%s", encryptColumnDefinition.LikeQueryAlgorithm.ToString())
	}

	if encryptColumnDefinition.AssistedQueryAlgorithm != nil {
		assistedQueryAlgorithm = fmt.Sprintf(",%s", encryptColumnDefinition.AssistedQueryAlgorithm.ToString())
	}

	if encryptColumnDefinition.LikeQueryAlgorithm != nil {
		likeQueryAlgorithm = fmt.Sprintf(",%s", encryptColumnDefinition.LikeQueryAlgorithm.ToString())
	}

	if encryptColumnDefinition.QueryWithCipherColumn != nil {
		queryWithCipherColumn = fmt.Sprintf(",QUERY_WITH_CIPHER_COLUMN=%s", encryptColumnDefinition.QueryWithCipherColumn.ToString())
	}

	return fmt.Sprintf("(%s%s,%s%s%s,%s%s%s%s)",
		encryptColumnDefinition.ColumnDefinition.ToString(),
		plainColumnDefinition,
		cipherColumnDefinition,
		assistedQueryColumnDefinition,
		likeQueryColumnDefinition,
		encryptColumnDefinition.EncryptAlgorithm.ToString(),
		assistedQueryAlgorithm,
		likeQueryAlgorithm,
		queryWithCipherColumn)
}

type ColumnDefinition struct {
	ColumnName *CommonIdentifier
	DataType   *DataType
}

func (columnDefinition *ColumnDefinition) ToString() string {
	var dataType string
	if columnDefinition.DataType != nil {
		dataType = fmt.Sprintf(",DATA_TYP=%s", columnDefinition.DataType.ToString())
	}

	return fmt.Sprintf("NAME=%s%s", columnDefinition.ColumnName.ToString(), dataType)
}

type PlainColumnDefinition struct {
	PlainColumnName *CommonIdentifier
	DataType        *DataType
}

func (plainColumnDefinition *PlainColumnDefinition) ToString() string {
	var (
		plainColumnName string
		dataType        string
	)
	if plainColumnDefinition.PlainColumnName != nil {
		plainColumnName = fmt.Sprintf("PLAIN=%s", plainColumnDefinition.PlainColumnName.ToString())
	}

	if plainColumnDefinition.DataType != nil {
		dataType = plainColumnDefinition.DataType.ToString()
	}
	return fmt.Sprintf("PLAIN = %s, PLAIN_DATA_TYPE = %s", plainColumnName, dataType)
}

type CipherColumnDefinition struct {
	CipherColumnName *CommonIdentifier
	DataType         *DataType
}

func (cipherColumnDefinition *CipherColumnDefinition) ToString() string {
	var dataType string
	if cipherColumnDefinition.DataType != nil {
		dataType = fmt.Sprintf(",CIPHER_DATA_TYPE=%s", dataType)
	}
	return fmt.Sprintf("CIPHER=%s%s", cipherColumnDefinition.CipherColumnName.ToString(), dataType)
}

type AssistedQueryColumnDefinition struct {
	AssistedQueryColumnName *CommonIdentifier
	DataType                *DataType
}

func (assistedQueryColumnDefinition *AssistedQueryColumnDefinition) ToString() string {
	var dataType string
	if assistedQueryColumnDefinition.DataType != nil {
		dataType = fmt.Sprintf(",ASSISTED_QUERY_DATA_TYPE=%s", assistedQueryColumnDefinition.DataType.ToString())
	}
	return fmt.Sprintf("ASSISTED_QUERY_COLUMN=%s%s", assistedQueryColumnDefinition.AssistedQueryColumnName.ToString(), dataType)
}

type LikeQueryColumnDefinition struct {
	LikeQueryColumnName *CommonIdentifier
	DataType            *DataType
}

func (likeQueryColumnDefinition *LikeQueryColumnDefinition) ToString() string {
	var dataType string
	if likeQueryColumnDefinition.DataType != nil {
		dataType = fmt.Sprintf(", LIKE_QUERY_DATA_TYPE=%s", likeQueryColumnDefinition.DataType.ToString())
	}
	return fmt.Sprintf("LIKE_QUERY_COLUMN=%s%s", likeQueryColumnDefinition.LikeQueryColumnName.ToString(), dataType)
}

type EncryptAlgorithm struct {
	AlgorithmDefinition *AlgorithmDefinition
}

func (encryptAlgorithm *EncryptAlgorithm) ToString() string {
	return fmt.Sprintf("ENCRYPT_ALGORITHM(%s)", encryptAlgorithm.AlgorithmDefinition.ToString())
}

type AssistedQueryAlgorithm struct {
	AlgorithmDefinition *AlgorithmDefinition
}

func (assistedQueryAlgorithm *AssistedQueryAlgorithm) ToString() string {
	return assistedQueryAlgorithm.AlgorithmDefinition.ToString()
}

type AlgorithmDefinition struct {
	AlgorithmTypeName    *AlgorithmTypeName
	PropertiesDefinition *PropertiesDefinition
}

func (algorithmDefinition AlgorithmDefinition) ToString() string {
	var propertiesDefinition string

	if algorithmDefinition.PropertiesDefinition != nil {
		propertiesDefinition = fmt.Sprintf(",%s", algorithmDefinition.PropertiesDefinition.ToString())
	}

	return fmt.Sprintf("TYPE(NAME=%s%s)", algorithmDefinition.AlgorithmTypeName.ToString(), propertiesDefinition)
}

type PropertiesDefinition struct {
	Properties *Properties
}

func (propertiesDefinition *PropertiesDefinition) ToString() string {
	if propertiesDefinition.Properties != nil {
		return fmt.Sprintf("PROPERTIES(%s)", propertiesDefinition.Properties.ToString())
	}
	return ""
}

type Properties struct {
	Properties []*Property
}

func (properties *Properties) ToString() (sql string) {
	for _, property := range properties.Properties {
		sql += property.ToString()
	}
	return
}

type LikeQueryAlgorithm struct {
	AlgorithmDefinition *AlgorithmDefinition
}

func (likeQueryAlgorithm *LikeQueryAlgorithm) ToString() (sql string) {
	if likeQueryAlgorithm.AlgorithmDefinition != nil {
		sql += likeQueryAlgorithm.ToString()
	}
	return
}

type QueryWithCipherColumn struct {
	QueryWithCipherColumn string
}

func (queryWithAlgorithm *QueryWithCipherColumn) ToString() string {
	return queryWithAlgorithm.QueryWithCipherColumn
}

type CommonIdentifier struct {
	Identifier string
}

func (commonIdentifier *CommonIdentifier) ToString() string {
	return commonIdentifier.Identifier
}

type Property struct {
	Key     string
	Literal *Literal
}

func (property *Property) ToString() string {
	if property.Literal != nil {
		return fmt.Sprintf("%s=%s", property.Key, property.Literal.ToString())
	}
	return ""
}

type Literal struct {
	Literal string
}

func (literal *Literal) ToString() string {
	return literal.Literal
}

type BuildinAlgorithmTypeName struct {
	AlgorithmTypeName string
}

func (buildinAlgorithmTypeName *BuildinAlgorithmTypeName) ToString() string {
	return buildinAlgorithmTypeName.AlgorithmTypeName
}

type DataType struct {
	String string
}

func (dataType *DataType) ToString() string {
	return dataType.String
}

type AlgorithmTypeName struct {
	BuildinAlgorithmTypeName *BuildinAlgorithmTypeName
	String                   string
}

func (algorithmTypeName *AlgorithmTypeName) ToString() string {
	switch {
	case algorithmTypeName.BuildinAlgorithmTypeName != nil:
		return algorithmTypeName.BuildinAlgorithmTypeName.ToString()
	case algorithmTypeName.String != "":
		return algorithmTypeName.String
	}
	return ""
}
