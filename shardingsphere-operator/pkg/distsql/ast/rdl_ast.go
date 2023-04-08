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

// Define RDL AST
type CreateEncryptRule struct {
	Create                   string
	Encrypt                  string
	EncryptName              string
	IfNotExists              *IfNotExists
	AllEncryptRuleDefinition []*EncryptRuleDefinition
}

type AlterEncryptRule struct {
	EncryptRuleDefinition []*EncryptRuleDefinition
}

func (alterEncryptRule *AlterEncryptRule) ToString() string {
	return ""
}

type DropEncryptRule struct {
	IfExists     *IfExists
	AllTableName []*CommonIdentifier
}

type IfExists struct {
	IfExists string
}

func (dropEncryptRule *DropEncryptRule) ToString() string {
	return ""
}

type EncryptRuleDefinition struct {
	TableName                  *CommonIdentifier
	ResourceDefinition         *ResourceDefinition
	AllEncryptColumnDefinition []*EncryptColumnDefinition
	QueryWithCipherColumn      *QueryWithCipherColumn
}

func (encryptRuleDefinition *EncryptRuleDefinition) ToString() string {
	return ""
}

type IfNotExists struct {
	IfNotExists string
}

func (ifNotExists IfNotExists) ToString() string {
	return ""
}

type ResourceDefinition struct {
	ResourceName *CommonIdentifier
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

type ColumnDefinition struct {
	ColumnName *CommonIdentifier
	DataType   *DataType
}

type PlainColumnDefinition struct {
	PlainColumnName *CommonIdentifier
	DataType        *DataType
}

type CipherColumnDefinition struct {
	CipherColumnName *CommonIdentifier
	DataType         *DataType
}

type AssistedQueryColumnDefinition struct {
	AssistedQueryColumnName *CommonIdentifier
	DataType                *DataType
}

type LikeQueryColumnDefinition struct {
	LikeQueryColumnName *CommonIdentifier
	DataType            *DataType
}

type EncryptAlgorithm struct {
	AlgorithmDefinition *AlgorithmDefinition
}

type AssistedQueryAlgorithm struct {
	AlgorithmDefinition *AlgorithmDefinition
}

type AlgorithmDefinition struct {
	AlgorithmTypeName    *AlgorithmTypeName
	PropertiesDefinition *PropertiesDefinition
}

type PropertiesDefinition struct {
	Properties *Properties
}

type Properties struct {
	Properties []*Property
}

type LikeQueryAlgorithm struct {
	AlgorithmDefinition *AlgorithmDefinition
}

type QueryWithCipherColumn struct {
	QueryWithCipherColumn string
}

type CommonIdentifier struct {
	Identifier string
}

type Property struct {
	Key     string
	Literal *Literal
}

type Literal struct {
	Literal string
}

type BuildinAlgorithmTypeName struct {
	AlgorithmTypeName string
}

type DataType struct {
	String string
}

type AlgorithmTypeName struct {
	BuildinAlgorithmTypeName *BuildinAlgorithmTypeName
	String                   string
}
