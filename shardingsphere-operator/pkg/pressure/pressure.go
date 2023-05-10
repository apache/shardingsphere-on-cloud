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
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/database-mesh/golang-sdk/pkg/random"

	"github.com/apache/shardingsphere-on-cloud/shardingsphere-operator/api/v1alpha1"
	_ "github.com/go-sql-driver/mysql"
)

type Pressure struct {
	Active         bool
	Name           string
	Result         Result
	Err            error
	Tasks          []v1alpha1.DistSQL
	finishSignalCh chan struct{}
	wg             sync.WaitGroup
}

var (
	db       *sql.DB
	totalReq int
)

type Result struct {
	//total exec req Number
	Total int
	//total success req Number
	Success int
	//todo: get total or get every exec

	//total time in this Pressure execution
	Duration time.Duration
}

func NewPressure(name string, tasks []v1alpha1.DistSQL) *Pressure {
	return &Pressure{
		Active:         false,
		Name:           name,
		Result:         Result{},
		Err:            nil,
		Tasks:          tasks,
		wg:             sync.WaitGroup{},
		finishSignalCh: make(chan struct{}),
	}
}

// todo: get conn args by labels over string
func initDB(connArgs string) error {
	var err error
	db, err = sql.Open("mysql", connArgs)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	db.SetConnMaxLifetime(60 * time.Second)
	return nil
}

func (p *Pressure) Run(ctx context.Context, pressureCfg *v1alpha1.PressureCfg) {
	p.Active = true
	totalReq = 0

	//judge nil for simplify test
	if db == nil {
		if err := initDB(pressureCfg.SsHost); err != nil {
			p.Err = err
			return
		}
		defer func() {
			if err := db.Close(); err != nil {
				p.Err = err
			}
		}()
	}

	result := &p.Result
	pressureCtx, cancel := context.WithTimeout(context.Background(), pressureCfg.Duration.Duration)
	defer cancel()
	ticker := time.NewTicker(pressureCfg.ReqTime.Duration)
	resCh := make(chan bool, 1000)

	//handle result
	go p.handleResponse(pressureCtx, resCh, result)

	//statistics the running time
	start := time.Now()
FOR:
	for {
		select {
		case <-ctx.Done():
			break FOR
		case <-pressureCtx.Done():
			break FOR
		case <-ticker.C:
			for i := 0; i < pressureCfg.ConcurrentNum; i++ {
				totalReq += pressureCfg.ReqNum
				//todo: handle err

				//put wg here to prevent: when root ctx is closed,but some exec task do not start yet
				p.wg.Add(1)
				go p.exec(pressureCtx, pressureCfg.ReqNum, resCh)
			}
		}
	}

	//occur when pressureCtx or root ctx closed

	//wait all exec calls return,we can safely close the result channel
	p.wg.Wait()
	end := time.Now()
	p.Result.Duration = end.Sub(start)
	close(resCh)

	//wait collect results channel finished
	<-p.finishSignalCh

	//when all task finished,update active
	p.Active = false
}

func (p *Pressure) exec(ctx context.Context, times int, res chan bool) {
	defer p.wg.Done()
	for i := 0; i < times; i++ {
		select {
		case <-ctx.Done():
			return
		default:
		}
		if len(p.Tasks) == 0 {
			return
		}
		for i := range p.Tasks {
			//generate diff sql, put result into channel
			args := randomArgs(p.Tasks[i].Args)
			_, err := db.Exec(p.Tasks[i].SQL, args)
			res <- err == nil
		}
	}
}

func (p *Pressure) handleResponse(ctx context.Context, resCh chan bool, result *Result) {
For:
	for {
		select {
		case <-ctx.Done():
			break For
		case ret := <-resCh:
			//todo: add more msg
			handle(ret, result)
		}
	}

	//get left handleResponse
	for ret := range resCh {
		handle(ret, result)
	}

	//when all handle finish,put a signal to finish chan
	p.finishSignalCh <- struct{}{}
}

//todo:add more logic and change ret type(bool ---> struct)
func handle(ret bool, result *Result) {
	if ret {
		result.Success++
	}
	result.Total++
}

func randomArgs(args []string) []string {
	var ret []string
	for i := range args {
		randomArg := fmt.Sprintf("%s-%s", args[i], random.StringN(4))
		ret = append(ret, randomArg)
	}
	return ret
}
