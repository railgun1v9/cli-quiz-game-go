/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strings"
)

// quizCmd represents the quiz command
var quizCmd = &cobra.Command{
	Use:   "quiz",
	Short: "Begin quiz",
	Long: `Takes a CSV of quiz questions and answers and asks the user one question
		at a time in a timed quiz.`,
	Run: func(cmd *cobra.Command, args []string) {
		correctCount := 0
		questions := make(map[string]string)

		// Open
		csvFile, err := os.Open("problems.csv")
		if err != nil {
			fmt.Println("An error occurred while opening csv.", err)
			return
		}

		r := csv.NewReader(csvFile)

		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}

			if err != nil {
				fmt.Println("An error occurred reading from csv.", err)
				return
			}

			// Add question as key and answer as value to questions map
			questions[record[0]] = record[1]
		}

		for k, v := range questions {
			msg := fmt.Sprintf("Question: %s = ? - Answer: ", k)
			fmt.Print(msg)
			reader := bufio.NewReader(os.Stdin)
			// ReadString will block until the delimiter is entered
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("An error occurred while reading input.", err)
				return
			}

			// remove the delimiter from the string
			input = strings.TrimSuffix(input, "\n")
			fmt.Println(input)

			if input == v {
				correctCount++
			}
		}

		fmt.Println(fmt.Sprintf("You have answered %d out of %d questions correctly!", correctCount, len(questions)))
	},
}

func init() {
	rootCmd.AddCommand(quizCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// quizCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// quizCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
