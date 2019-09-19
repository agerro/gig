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
	"io"
	"net/http"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Init and fetch new template definitions",
	Long: `This command will initialize and fetch the current up to date template that can be used for creating yout .gitignore file`,
	Run: func(cmd *cobra.Command, args []string) {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Configuring GitIGnore")
		verifyConfigDirExistsElseCreate(home + "/.gig")
		downloadFile(home + "/.gig/" + "languages.json", "https://gitignore.io/dropdown/templates.json")
		verifyConfigFilesExists(home + "/.gig/")
	},
}

func verifyConfigFilesExists(dirPath string) {
	fmt.Println("Checking config files...")
	if _, err := os.Stat(dirPath + "languages.json"); !os.IsNotExist(err) {
		fmt.Println("Language definition exists, continue...")
	} else {
		fmt.Println(err)
		fmt.Println("No langauge definitions found, please run the command 'gig configure'")
		os.Exit(1)
	}
}

func verifyConfigDirExistsElseCreate(dirPath string){
	fmt.Println("Checking if config directory exists...")
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		fmt.Println("Creating config directory at " + dirPath)
		os.Mkdir(dirPath, 0777)
	} else {
		fmt.Println("Config directory already exists at " + dirPath)
	}
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func downloadFile(filepath string, url string) error {

	fmt.Println("Attmpeting to download new language definitions")

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
	rootCmd.AddCommand(configureCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
