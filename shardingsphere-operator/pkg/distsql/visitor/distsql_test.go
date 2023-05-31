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

package visitor

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/distsql/ast"
	parser "github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/pkg/distsql/visitor_parser/encrypt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Distsql", func() {
	var (
		encryptDistSQL = "CREATE ENCRYPT RULE t_encrypt (COLUMNS((NAME=user_id,PLAIN=user_plain,CIPHER=user_cipher,ENCRYPT_ALGORITHM(TYPE(NAME='AES',PROPERTIES('aes-key-value'='123456abc')))),(NAME=order_id,CIPHER=order_cipher,ENCRYPT_ALGORITHM(TYPE(NAME='MD5')))),QUERY_WITH_CIPHER_COLUMN=true);"
		visitor        = EncryptVisitor{}
		ast            = &ast.CreateEncryptRule{}
	)

	BeforeEach(func() {
		inputStream := antlr.NewInputStream(encryptDistSQL)
		lexer := parser.NewRDLStatementLexer(inputStream)
		tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		distSQLParser := parser.NewRDLStatementParser(tokens)
		createEncryptRule := distSQLParser.CreateEncryptRule()
		ast = visitor.VisitCreateEncryptRule(createEncryptRule.(*parser.CreateEncryptRuleContext))
	})

	Context("parse distSQL to AST", func() {
		It("should encrypt distSQL parse correctly", func() {
			Expect(ast.AllEncryptRuleDefinition[0].TableName.Identifier).To(Equal("t_encrypt"))
		})
	})

	Context("covert distSQL AST to string", func() {
		It("should encrypt distsql parse correctly", func() {
			Expect(ast.AllEncryptRuleDefinition[0].TableName.ToString()).To(Equal("t_encrypt"))
		})
	})
})
