//
// Copyright © 2017-2019 Solus Project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"libferry"
	"os"
)

var (
	fromRepo string
	toRepo string
	deepClone bool
)

var cloneRepoCmd = &cobra.Command{
	Use:   "clone",
	Short: "clone an existing repository",
	Long:  "Clone an existing repository into a new repository",
	Run:   cloneRepo,
}

func init() {
	cloneRepoCmd.Flags().StringVarP(&fromRepo, "from", "f", "", "Source Repo")
	cloneRepoCmd.Flags().StringVarP(&toRepo, "to", "t", "", "Destination Repo")
	cloneRepoCmd.Flags().BoolVarP(&deepClone, "deep", "d", false, "Perform a deep clone")
	cloneRepoCmd.MarkFlagRequired("from")
	cloneRepoCmd.MarkFlagRequired("to")

	RootCmd.AddCommand(cloneRepoCmd)
}

func cloneRepo(cmd *cobra.Command, args []string) {
	if (fromRepo == "") || (toRepo == "") {
		fmt.Fprintf(os.Stderr, cmd.UsageString())
		return
	}

	client := libferry.NewClient(socketPath)
	defer client.Close()

	if err := client.CloneRepo(fromRepo, toRepo, deepClone); err != nil {
		fmt.Fprintf(os.Stderr, "Error while cloning repo: %v\n", err)
		return
	}
}
