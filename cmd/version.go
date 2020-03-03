// Copyright 2020 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/c4pt0r/tiup/pkg/version"
	"github.com/spf13/cobra"
)

func newVersionCmd() *cobra.Command {
	cmdVersion := &cobra.Command{
		Use:   "version",
		Short: "Show tiup version and quit",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version.NewTiUPVersion())
		},
	}
	return cmdVersion
}
