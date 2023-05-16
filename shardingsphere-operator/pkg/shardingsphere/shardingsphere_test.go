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

package shardingsphere

import (
	"database/sql"
	"fmt"
	"regexp"

	"bou.ke/monkey"
	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test ShardingSphere Server", func() {
	var (
		db     *sql.DB
		dbmock sqlmock.Sqlmock
		err    error
		s      IServer
	)
	BeforeEach(func() {
		db, dbmock, err = sqlmock.New()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(dbmock).ShouldNot(BeNil())

		monkey.Patch(sql.Open, func(driverName, dataSourceName string) (*sql.DB, error) {
			return db, nil
		})

		s, err = NewServer("mysql", "localhost", uint(3307), "user", "password")
		Expect(err).ShouldNot(HaveOccurred())
	})

	AfterEach(func() {
		monkey.Unpatch(sql.Open)
		db.Close()
	})

	Context("Test Create database", func() {
		It("should create success", func() {
			dbmock.ExpectExec(regexp.QuoteMeta("CREATE DATABASE")).WillReturnResult(sqlmock.NewResult(1, 1))

			err = s.CreateDatabase("test_db")
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	// test register storage unit
	Context("Test register storage unit", func() {
		It("should register success", func() {
			// mock db and return register storage unit success
			dbmock.ExpectExec(regexp.QuoteMeta("REGISTER STORAGE UNIT")).WillReturnResult(sqlmock.NewResult(1, 1))

			// create server
			err = s.RegisterStorageUnit("ds_0", "localhost", uint(3307), "sharding_db", "user", "password")
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	// test get rules used by storage units
	Context("Test get rules used by storage units", func() {
		// should return a sharding table rule named 't_order'.
		It("should return a sharding table rule named 't_order'", func() {
			// mock db and return sharding table rule
			dbmock.ExpectQuery(regexp.QuoteMeta("SHOW RULES USED STORAGE UNIT")).WillReturnRows(sqlmock.NewRows([]string{"type", "name"}).AddRow("sharding", "t_order"))

			result, err := s.(*server).getRulesUsed("ds_0", "sharding_db")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).Should(Equal([]*Rule{{Type: "sharding", Name: "t_order"}}))
		})
	})

	// test drop rule by rule type 'sharding' and rule name 't_order'
	Context("Test drop rule by rule type 'sharding' and rule name 't_order'", func() {
		It("should drop success", func() {
			// mock db and return drop rule success
			dbmock.ExpectExec("DROP SHARDING TABLE RULE").WillReturnResult(sqlmock.NewResult(1, 1))

			err := s.(*server).dropRule("sharding", "t_order")
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("Test unregister storage node", func() {
		It("should unregister success", func() {
			dbmock.ExpectQuery(regexp.QuoteMeta("SHOW RULES USED STORAGE UNIT")).WillReturnRows(sqlmock.NewRows([]string{"type", "name"}))
			dbmock.ExpectExec(regexp.QuoteMeta("UNREGISTER STORAGE UNIT")).WillReturnResult(sqlmock.NewResult(1, 1))

			err = s.UnRegisterStorageUnit("ds_0")
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})

var _ = Describe("Test ShardingSphere Server Manually", func() {
	var (
		driver string
		host   string
		port   uint
		user   string
		pass   string
	)

	Context("Test create database", func() {
		It("should create success", func() {
			if driver == "" || host == "" || port == 0 || user == "" || pass == "" {
				Skip("skip test")
			}
			dbName := "test_db"
			s, err := NewServer(driver, host, port, user, pass)
			Expect(err).ShouldNot(HaveOccurred())
			err = s.CreateDatabase(dbName)
			Expect(err).ShouldNot(HaveOccurred())

			s.(*server).db.Exec(fmt.Sprintf(`DROP DATABASE %s`, dbName))
		})
	})
})
