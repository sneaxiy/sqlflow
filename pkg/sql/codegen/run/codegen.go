// Copyright 2020 The SQLFlow Authors. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package run

import (
	"sqlflow.org/sqlflow/pkg/ir"
	"sqlflow.org/sqlflow/pkg/log"
	pb "sqlflow.org/sqlflow/pkg/proto"
)

// Run is the unified entry for the function in `TO RUN` clause
func Run(runStmt *ir.RunStmt, session *pb.Session) (string, error) {
	logger := log.GetDefaultLogger()
	logger.Infof("Execute codegen.Run: %v", runStmt)

	return "", nil
}