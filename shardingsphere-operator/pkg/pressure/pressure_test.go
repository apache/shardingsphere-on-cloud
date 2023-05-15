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

package pressure

import (
	"bou.ke/monkey"
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"regexp"
	"testing"
	"time"
)

func TestPressure(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controllers Suite")
}

var _ = Describe("test pressure", func() {
	var (
		dbmock sqlmock.Sqlmock
		err    error
	)
	BeforeEach(func() {
		db, dbmock, err = sqlmock.New()
		Expect(err).To(BeNil())
		Expect(db).NotTo(BeNil())
		Expect(dbmock).NotTo(BeNil())

		monkey.Patch(sql.Open, func(driverName, dataSourceName string) (*sql.DB, error) {
			return db, nil
		})
	})

	AfterEach(func() {
		monkey.Unpatch(sql.Open)
		db.Close()
	})

	Context("test Run function", func() {
		It("should Run successfully", func() {
			registerStorageUnitCase := &v1alpha1.PressureCfg{
				ZkHost:   "",
				SsHost:   "test",
				Duration: metav1.Duration{Duration: 20 * time.Second},
				ReqTime:  metav1.Duration{Duration: 5 * time.Second},
				DistSQLs: []v1alpha1.DistSQL{
					{
						SQL: "REGISTER STORAGE UNIT ?",
						Args: []string{
							"**",
						},
					},
				},
				ConcurrentNum: 2,
				ReqNum:        5,
			}

			dbmock.ExpectExec(regexp.QuoteMeta("REGISTER STORAGE UNIT")).WillReturnResult(sqlmock.NewResult(1, 1))
			pressure := NewPressure("verify", registerStorageUnitCase.DistSQLs)
			pressure.Run(context.TODO(), registerStorageUnitCase)

			Expect(pressure.Result.Total > 0).To(BeTrue())
			Expect(pressure.Result.Success >= 0).To(BeTrue())
			Expect(pressure.Result.Total >= pressure.Result.Success).To(BeTrue())
			Expect(pressure.Result.Duration.Milliseconds() >= registerStorageUnitCase.Duration.Milliseconds()).To(BeTrue())
			Expect(pressure.Active).To(BeFalse())
		})
	})

})
