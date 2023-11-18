// Copyright 2023 Database Mesh Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package promptutil

import (
	"bufio"
	"fmt"
	"os"

	"github.com/apache/shardingsphere-on-cloud/pitr/cli/internal/pkg/xerr"
	"github.com/apache/shardingsphere-on-cloud/pitr/cli/pkg/logging"
)

func GetUserApproveInTerminal(prompt string) error {
	logging.Warn(fmt.Sprintf("\n%s", prompt))
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		return xerr.NewCliErr(fmt.Sprintf("read user input failed:%s", err.Error()))
	}
	if scanner.Text() != "Y" && scanner.Text() != "y" && scanner.Text() != "yes" && scanner.Text() != "YES" && scanner.Text() != "Yes" {
		return xerr.NewCliErr("User abort")
	}
	return nil
}
