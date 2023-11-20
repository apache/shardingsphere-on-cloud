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

package prettyoutput

import (
	"time"

	"github.com/jedib0t/go-pretty/v6/progress"
)

func NewPW(totalNum int) progress.Writer {
	pw := progress.NewWriter()

	pw.SetTrackerLength(25)
	pw.SetAutoStop(true)
	pw.SetNumTrackersExpected(totalNum)
	pw.SetSortBy(progress.SortByPercentDsc)
	pw.SetTrackerPosition(progress.PositionRight)

	style := progress.StyleDefault
	style.Options.PercentIndeterminate = "running"
	pw.SetStyle(style)

	return pw
}

type ProgressPrintOption struct {
	NumTrackersExpected int
}

type ProgressPrinter struct {
	progress.Writer
}

func NewProgressPrinter(opt ProgressPrintOption) *ProgressPrinter {
	p := &ProgressPrinter{
		Writer: progress.NewWriter(),
	}

	// passed printer options
	p.SetNumTrackersExpected(opt.NumTrackersExpected)

	// default printer options
	p.SetTrackerLength(25)
	p.SetAutoStop(true)
	p.SetSortBy(progress.SortByPercentDsc)
	p.SetTrackerPosition(progress.PositionRight)
	style := progress.StyleDefault
	style.Options.PercentIndeterminate = "running"
	p.SetStyle(style)

	return p
}

func (p *ProgressPrinter) BlockedRendered() {
	time.Sleep(time.Millisecond * 100)
	for p.IsRenderInProgress() {
		time.Sleep(time.Millisecond * 100)
	}
}

func (p *ProgressPrinter) UpdateProgress(tracker *progress.Tracker, updateF func() (bool, error)) {
	var (
		done   = make(chan struct{})
		ticker = time.NewTicker(time.Second * 2)
	)

	for !tracker.IsDone() {
		select {
		case <-done:
			return
		case <-ticker.C:
			finished, err := updateF()
			if err != nil {
				tracker.MarkAsErrored()
				done <- struct{}{}
			}

			if finished {
				tracker.MarkAsDone()
				done <- struct{}{}
			}
		}
	}
}
