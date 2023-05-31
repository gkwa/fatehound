/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test called")
		test()
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func test() {
	// Define the source file path
	sourceFilePath := "C:/ProgramData/Streambox/SpectraUI/settings.xml"

	// Open the source file in read mode
	sourceFile, err := os.OpenFile(sourceFilePath, os.O_RDWR, 0o666)
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}

	// Acquire an exclusive lock on the source file
	if err := sourceFile.Truncate(0); err != nil {
		fmt.Println("Error truncating the source file:", err)
		sourceFile.Close()
		return
	}

	// Create a strings.Builder to store the modified content
	modifiedContent := strings.Builder{}

	// Create a bufio.Scanner to read the source file line by line
	scanner := bufio.NewScanner(sourceFile)
	for scanner.Scan() {
		line := scanner.Text()

		// Replace video_3d="0" with an empty string
		line = strings.Replace(line, `video_3d="0" ?`, "", -1)

		// Write the modified line to the strings.Builder
		modifiedContent.WriteString(line + "\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning source file:", err)
		sourceFile.Close()
		return
	}

	// Close the source file
	sourceFile.Close()

	// Open the source file again, but this time without a lock
	sourceFile, err = os.OpenFile(sourceFilePath, os.O_WRONLY|os.O_TRUNC, 0o666)
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}

	// Write the modified content to the source file
	_, err = sourceFile.WriteString(modifiedContent.String())
	if err != nil {
		fmt.Println("Error writing to source file:", err)
		sourceFile.Close()
		return
	}

	// Close the source file
	sourceFile.Close()

	fmt.Println("File modified successfully!")
}
