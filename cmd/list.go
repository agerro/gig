/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"os"
	"io/ioutil"
	"encoding/json"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

type LanguageDefinition struct {
	Id string
	Text string
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the different language template that can be used",
	Long: `Using this command would check for the local copy of the templates that are available to use in your git ignore file.
	If the local copy should be updated, please run the command "gig configure" and it will fetch new (if any) new templates`,
	Run: func(cmd *cobra.Command, args []string) {
		verifyConfigAndLanguageDefinitions()
		printLanguageDefinition()
	},
}

func verifyConfigAndLanguageDefinitions() {

	home, err := homedir.Dir()
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
	}

	fmt.Println("Checking config files...")
	if _, err := os.Stat(home + "/.gig/" + "languages.json"); !os.IsNotExist(err) {
		fmt.Println("Language definition exists, continue...")
	} else {
		fmt.Println(err)
		fmt.Println("No langauge definitions found, please run the command 'gig configure'")
		os.Exit(1)
	}
}

func printLanguageDefinition() {
	fmt.Println("The following language definitions are available:")

	home, err := homedir.Dir()
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
	}

	file, err := ioutil.ReadFile(home + "/.gig/" + "languages.json") // just pass the file name
    if err != nil {
        fmt.Print(err)
	}

	var arr []LanguageDefinition
    _ = json.Unmarshal(file, &arr)

	for _, value := range arr {
        fmt.Printf(value.Id)
        fmt.Println("")
	}
	fmt.Println("")
	fmt.Println("Refine the search by pipeling this output to grep")
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
