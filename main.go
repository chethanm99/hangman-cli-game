package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {

	gamesplayed := 0
	wins := 0
	losses := 0

	for {
		gamesplayed++
		word, err := getRandomword("words.txt")
		if err != nil {
			fmt.Println("Error the words file: ", err)
			return
		}
		var attempts int
		currentWordstate := intializeCurrentwordstate(word)
		scanner := bufio.NewScanner(os.Stdin)
		guessedLetters := make(map[string]bool)
		wrongGuesses := []string{}

		fmt.Println("Welcome to the Hang Man Game, Let's get started")

		var useroption string
		fmt.Println(`Choose your playing level ?
	       1.Easy
	       2.Medium
	       3.Hard
	`)
		fmt.Scanln(&useroption)

		if useroption == "easy" {
			attempts = 10
		} else if useroption == "medium" {
			attempts = 6
		} else {
			attempts = 4
		}

		for attempts > 0 {
			displayCurrentwordstate(currentWordstate, attempts)
			userInput := getUserinput(scanner)
			if !isValidinput(userInput) {
				fmt.Println("Not valid input, please enter a single letter")
				continue
			}

			correctGuess := updateGuessedword(word, currentWordstate, userInput)

			if guessedLetters[userInput] {
				fmt.Println("You have already guesses that letter, try again !")
				continue
			}

			guessedLetters[userInput] = true

			if !correctGuess {
				attempts--
				wrongGuesses = append(wrongGuesses, userInput)
				fmt.Println("Wrong guesses till now : ", wrongGuesses)
			}

			if strings.Join(currentWordstate, "") == word {
				fmt.Println("Congratulations! You guessed the word:", word)
				wins++
				break
			}
		}

		if strings.Join(currentWordstate, "") != word {
			fmt.Println("You ran out of attempts. The word was :", word)
			displayHangman(6 - attempts)
			losses++
		}
		fmt.Println("Do you wish to continue ? if yes then enter 'y' else enter 'quit': ")
		var userchoice string
		fmt.Scanln(&userchoice)

		if userchoice != "y" {
			fmt.Println("Thanks for playing!")
			break
		}
	}

	fmt.Printf(`This is the summary of your gameplay :
	          Total games played :%v
	          Wins: %v
	          Losses: %v
	`, gamesplayed, wins, losses)
}

func intializeCurrentwordstate(word string) []string {
	currentWordstate := make([]string, len(word))

	for i := range word {
		currentWordstate[i] = "_"
	}
	return currentWordstate
}

func isValidinput(input string) bool {
	return utf8.RuneCountInString(input) == 1
}

func displayCurrentwordstate(currentWordstate []string, attempts int) {
	fmt.Println("Current Word State", strings.Join(currentWordstate, " "))
	fmt.Printf("You have %d attempts remaining.\n", attempts)
}

func getUserinput(scanner *bufio.Scanner) string {
	fmt.Println("Start guessing !")
	scanner.Scan()
	return scanner.Text()
}

func updateGuessedword(word string, guessed []string, letter string) bool {
	correctGuess := false
	for i := range word {
		if guessed[i] == letter {
			correctGuess = true
		}
	}
	return correctGuess
}

func displayHangman(incorrectguesses int) {
	if incorrectguesses >= 0 && incorrectguesses < len(hangmanStates) {
		fmt.Println(hangmanStates[incorrectguesses])
	}
}

func getRandomword(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return " ", err
	}
	words := strings.Split(string(data), "\n")
	return words[rand.Intn(len(words))], nil
}
