/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"encoding/json"
	"fmt"
	"github.com/abhi18av/twrgo/api"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
)

// serviceInfoCmd represents the serviceInfo command
var serviceInfoCmd = &cobra.Command{
	Use:   "serviceInfo",
	Short: "Brief information for the Tower installation.",
	Long: `This command can be used to obtain information about a Tower installation.`,
	Run: func(cmd *cobra.Command, args []string) {

		TOWER_API_ENDPOINT := viper.GetString("TOWER_API_ENDPOINT")
		res, err := json.MarshalIndent(api.ServiceInfo(TOWER_API_ENDPOINT),  "", "  ")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(res))
	},
}

func init() {
	rootCmd.AddCommand(serviceInfoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serviceInfoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serviceInfoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
