/*
Copyright © 2021 CELLA, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var (
	Version = ""
	Date    = ""
)

// GetVersion get CLI version
func GetVersion() string {
	if Version == "" {
		if info, ok := debug.ReadBuildInfo(); ok {
			return info.Main.Version
		}
		return "(none)"
	}
	if Date == "" {
		return fmt.Sprintf("%s", Version)
	}
	return fmt.Sprintf("%s(%s)", Version, Date)
}

// GetRuntimeVersion get yomo runtime version
func GetRuntimeVersion() (v string) {
	v = "(none)"
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, d := range info.Deps {
			if d.Path == "github.com/yomorun/yomo" {
				if d.Replace != nil {
					return d.Replace.Version + "[Replace]"
				}
				return d.Version
			}
		}
	}
	return
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print CLI version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("YoMo CLI Version:", GetVersion())
		fmt.Println("Runtime Version:", GetRuntimeVersion())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
