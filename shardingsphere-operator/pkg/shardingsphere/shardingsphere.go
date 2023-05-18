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

	_ "github.com/go-sql-driver/mysql"
)

const (
	// DistSQLCreateDatabase create database if not exists.
	DistSQLCreateDatabase = `CREATE DATABASE IF NOT EXISTS %s;`
	// DistSQLRegisterStorageUnit register database to shardingsphere by storage unit name and database info.
	DistSQLRegisterStorageUnit = `REGISTER STORAGE UNIT IF NOT EXISTS %s (HOST="%s",PORT=%d,DB="%s",USER="%s",PASSWORD="%s");`
	// DistSQLShowRulesUsed show all rules used by storage unit name.
	DistSQLShowRulesUsed = `SHOW RULES USED STORAGE UNIT %s FROM %s;`
	// DistSQLUnRegisterStorageUnit unregister database from shardingsphere by storage unit name.
	DistSQLUnRegisterStorageUnit = `UNREGISTER STORAGE UNIT %s;`
	// DistSQLDropRule drop rule by rule type and rule name.
	DistSQLDropRule = `DROP %s RULE %s;`
	// DistSQLDropTable drop table by table name.
	DistSQLDropTable = `DROP TABLE %s;`
)

var ruleTypeMap = map[string]string{}

type Rule struct {
	Type string
	Name string
}

type server struct {
	db *sql.DB
}

type IServer interface {
	CreateDatabase(dbName string) error
	RegisterStorageUnit(dsName, host string, port uint, dbName, user, password string) error
	UnRegisterStorageUnit(dsName string) error
	Close() error
}

var _ IServer = (*server)(nil)

func NewServer(driver, host string, port uint, user, password string) (IServer, error) {
	if driver != "mysql" && driver != "postgres" {
		return nil, fmt.Errorf("unsupported database driver: %s", driver)
	}

	if host == "" || port == 0 || user == "" || password == "" {
		return nil, fmt.Errorf("invalid database config, host=%s, port=%d, user=%s, password=%s", host, port, user, password)
	}

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/", user, password, host, port)

	db, err := sql.Open(driver, dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("open database=%s error: %w", dataSourceName, err)
	}

	// check database connection
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("ping database=%s error: %w", dataSourceName, err)
	}

	return &server{db: db}, nil
}

func (s *server) Close() error {
	return s.db.Close()
}

func (s *server) CreateDatabase(dbName string) error {
	distSQL := fmt.Sprintf(DistSQLCreateDatabase, dbName)

	_, err := s.db.Exec(distSQL)
	if err != nil {
		return fmt.Errorf("create database error: %w", err)
	}

	return nil
}

func (s *server) RegisterStorageUnit(dsName, host string, port uint, dbName, user, password string) error {
	distSQL := fmt.Sprintf(DistSQLRegisterStorageUnit, dsName, host, port, dbName, user, password)

	_, err := s.db.Exec(distSQL)
	if err != nil {
		return fmt.Errorf("register database error: %w", err)
	}

	return nil
}

// getRulesUsed returns all rules used by storage unit name.
func (s *server) getRulesUsed(dsName, dbName string) (rules []*Rule, err error) {
	rules = make([]*Rule, 0)
	distSQL := fmt.Sprintf(DistSQLShowRulesUsed, dsName, dbName)

	rows, err := s.db.Query(distSQL)
	if err != nil {
		return nil, fmt.Errorf("get rules used error: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var ruleT, ruleN string
		if err := rows.Scan(&ruleT, &ruleN); err != nil {
			return nil, fmt.Errorf("scan rules used error: %w", err)
		}
		rules = append(rules, &Rule{Type: ruleT, Name: ruleN})
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}
	return rules, nil
}

func (s *server) UnRegisterStorageUnit(dsName string) error {
	rules, err := s.getRulesUsed(dsName, "")
	if err != nil {
		return fmt.Errorf("get rules used error: %w", err)
	}

	// clean all rules used by storage unit
	for _, rule := range rules {
		if err := s.dropRule(rule.Type, rule.Name); err != nil {
			return fmt.Errorf("drop rule error: %w", err)
		}
	}

	distSQL := fmt.Sprintf(DistSQLUnRegisterStorageUnit, dsName)

	_, err = s.db.Exec(distSQL)
	if err != nil {
		return fmt.Errorf("unregister database error: %w", err)
	}

	return nil
}

func (s *server) dropRule(ruleType, ruleName string) error {
	// convert rule type
	ruleType = ruleTypeMap[ruleType]
	distSQL := fmt.Sprintf(DistSQLDropRule, ruleType, ruleName)
	_, err := s.db.Exec(distSQL)
	if err != nil {
		return fmt.Errorf("drop rule fail, err: %s", err)
	}
	return nil
}

func init() {
	// init rule type map
	// implement more rule type if needed
	ruleTypeMap = map[string]string{
		"sharding": "SHARDING TABLE",
	}
}
