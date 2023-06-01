/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		doit()
	},
}

var logger zerolog.Logger

func init() {
	rootCmd.AddCommand(testCmd)

	consoleWriter := zerolog.ConsoleWriter{
		Out: os.Stderr,
		PartsExclude: []string{
			zerolog.TimestampFieldName,
		},
		FormatLevel: func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("[%s]", i))
		},
		FormatCaller: func(i interface{}) string {
			return filepath.Base(fmt.Sprintf("%s", i))
		},
	}
	logger = zerolog.New(consoleWriter).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Logger()
}

func doit() {
	paths := viper.GetStringSlice("path")

	for _, path := range paths {
		doit1(path)
	}
}

func doit1(sourceFilePath string) {
	// Open the source file in read mode
	sourceFile, err := os.OpenFile(sourceFilePath, os.O_RDWR, 0o666)
	if err != nil {
		logger.Error().Msgf("Error opening source file: %v", err)
		return
	}

	// Create a strings.Builder to store the modified content
	modifiedContent := strings.Builder{}

	// Create a bufio.Scanner to read the source file line by line
	scanner := bufio.NewScanner(sourceFile)

	foundMatch := false
	pattern := regexp.MustCompile(`video_3d *= *"[^"]+" ?`)
	for scanner.Scan() {
		line := scanner.Text()

		if pattern.MatchString(line) {
			logger.Trace().Msgf("found match\n")
			foundMatch = true
		}
		logger.Trace().Msgf("Original line:%s", line)
		line = pattern.ReplaceAllString(line, "")
		logger.Trace().Msgf("Modified line:%s", line)

		modifiedContent.WriteString(line + "\n")
	}

	if err := scanner.Err(); err != nil {
		logger.Error().Msgf("error scanning source file: %v", err)
		return
	}

	// Close the source file
	sourceFile.Close()

	if !foundMatch {
		logger.Debug().Msgf("%s already clean", sourceFile.Name())
		return
	}

	// Open the source file again, but this time without a lock
	sourceFile, err = os.OpenFile(sourceFilePath, os.O_WRONLY|os.O_TRUNC, 0o666)
	if err != nil {
		logger.Error().Msgf("error opening source file: %v", err)
		return
	}

	// Write the modified content to the source file
	_, err = sourceFile.WriteString(modifiedContent.String())
	if err != nil {
		logger.Error().Msgf("error writing to source file %s: %v", sourceFile.Name(), err)
		sourceFile.Close()
		return
	}
	sourceFile.Close()

	logger.Info().Msgf("%s updated successfully!", sourceFile.Name())
}
