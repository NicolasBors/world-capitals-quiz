package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"

	"github.com/NicolasBors/world-capitals-quiz/constants"
	"github.com/NicolasBors/world-capitals-quiz/types"
	"github.com/spf13/cobra"
)

var client *http.Client

func init() {
	client = &http.Client{}
	root.AddCommand(start)
}

var start = &cobra.Command{
	Use:   "start",
	Short: "Start the quiz",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get(fmt.Sprint(constants.SERVER_BASE_URL, constants.PATH_GET_QUESTIONS))
		if err != nil {
			fmt.Printf("Error while fetching questions URL: %s", err)
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error while reading questions response data: %s", err)
			return
		}

		var questions []types.Question
		err = json.Unmarshal(body, &questions)
		if err != nil {
			fmt.Printf("Error while parsing JSON data: %s", err)
			return
		}

		var (
			validAnswers = getValidAnswers()
			userAnswers  []int
		)

		for _, question := range questions {
			for {
				fmt.Printf("%s\n", question.Question)
				options := question.Options
				for index, option := range options {
					fmt.Printf("%d - %s\n", index+1, option)
				}

				fmt.Println("Enter your answer:")
				var answer int
				_, err := fmt.Scanln(&answer)
				if err != nil || !slices.Contains(validAnswers, answer) {
					fmt.Printf("Invalid answer, should be on a range of 1 to %d. Please try again\n", constants.OPTIONS_LENGTH)
				} else {
					userAnswers = append(userAnswers, answer)
					break
				}

			}
		}

		submit(userAnswers)
	},
}

func getValidAnswers() []int {
	validAnswers := []int{}
	for i := 1; i <= constants.OPTIONS_LENGTH; i++ {
		validAnswers = append(validAnswers, i)
	}
	return validAnswers
}

func submit(answers []int) {
	jsonData, err := json.Marshal(answers)
	if err != nil {
		fmt.Println("Error while marshalling answers:", err)
		return
	}

	resp, err := client.Post(fmt.Sprint(constants.SERVER_BASE_URL, constants.PATH_POST_ANSWERS), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error while submitting answers:", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error while reading submit response data: %s", err)
		return
	}

	fmt.Println(string(body))
}
