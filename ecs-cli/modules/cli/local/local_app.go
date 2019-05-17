// Copyright 2015-2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Package local implements the subcommands to run ECS task definitions locally
// (See: https://github.com/aws/containers-roadmap/issues/180).
package local

import (
	"fmt"
	"github.com/urfave/cli"
)

func Create(c *cli.Context) {
	// 1. read in task def (from file or arn)
	// 2. parse task def into go object
	// 3. write to docker-compose.local.yml file
	fmt.Println("foo") // placeholder
}
