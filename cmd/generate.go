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
	"net/http"
	"os"
	"strings"
	"io"
	"io/ioutil"
	"encoding/json"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// parameters
var languageInput []string
var outputPath string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a .gitignore file at your current location",
	Long: `This command will use the input of languages provided and generate a .gitignore file based on that at your current location.
	The path of where to store the .gitignore file can be overridden by passing the argument for -p, --path.`,
	Run: func(cmd *cobra.Command, args []string) {
		validateInput(languageInput)
		languageString := strings.Join(languageInput, ",")
		downloadIgoreFile(outputPath + ".gitignore", "https://gitignore.io/api/" + languageString)
	},
}

func validateInput(input []string) {
	for _,value := range input {
        validateLanguage(value)
    }
}

func validateLanguage(language string) {
	fmt.Println("Validating language:", language, "...")

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

	sum := 0
	for _, value := range arr {
        if language == value.Id {
			sum += 1
		} else {
			sum += 0
		}
	}

	if sum != 0 {
		fmt.Println("VALID")
	 } else {
		 fmt.Println("INVALID")
	 }
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func downloadIgoreFile(filepath string, url string) error {

	fmt.Println("Generating ignore file based upon valid input languages")

	// Get the data
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // Create the file
    out, err := os.Create(filepath)
    if err != nil {
        return err
    }
    defer out.Close()

    // Write the body to file
    _, err = io.Copy(out, resp.Body)
    return err
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")
	generateCmd.Flags().StringVarP(&outputPath, "path", "p", "", "Path to output ignore file to")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	generateCmd.Flags().StringSliceVarP(&languageInput, "languages", "l", []string{}, "Input each language that you want to be a part of the ignore file. Input one language as one '-l' parameter (case sensitive against the output of 'gig list'). Available languages can be found running the command 'gig list'")
}
