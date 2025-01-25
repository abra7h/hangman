package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Чтение файла со словами
func fileReader() []string {
	file, err := os.Open("words.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var dataWords []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		dataWords = append(dataWords, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	err = file.Close()
	if err != nil {
		fmt.Println(err)
	}

	return dataWords
}

func launchingTheGame(words []string) {
	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	for {
		var secretWord []rune
		word := words[rand.Intn(len(words))]
		for _, letter := range word {
			secretWord = append(secretWord, letter)
		}

		startingVar := startMenu()
		if !startingVar {
			return
		}

		gamePlay(secretWord)
	}

}

// Функция для старта игры
func startMenu() bool {
	var command string
	var isGameStarted bool

	fmt.Println("Press [N]ew to Start game or [E]nd to kill this program")
	if _, err := fmt.Scanln(&command); err != nil {
		fmt.Println(err)
	}

	switch command {
	case "N", "n":
		isGameStarted = true
	case "E", "e":
		isGameStarted = false
	default:
		fmt.Println("Command not recognized, try again")
		isGameStarted = startMenu()
	}

	return isGameStarted
}

func gamePlay(word []rune) {
	fmt.Println(string(word))

	wordLength := len(word)
	guessedLetters := printGuessedWord(wordLength)

	var guessed []rune
	var mistakeCounter int
	var currentLength int // текущая количество отгаданных букв

	for {
		isFound, currentLength := checkTheLetterInWord(&word, &guessedLetters, &guessed, &currentLength)

		if !isFound {
			mistakeCounter++
			hangThisMan(mistakeCounter)
			if mistakeCounter == 5 {
				fmt.Println("Ты проиграл :(")
				fmt.Print("Загаданное слово: ")
				printSlice(&word)
				break
			}
		}

		if wordLength == currentLength { // выходим из цикла, если все слово отгадано
			fmt.Println("You guessed my word! Congratulations!")
			break
		}
		fmt.Println()
	}

}

func enterTheLetter() rune {
	var inputLetter string

	fmt.Print("Enter the letter: ")
	_, err := fmt.Scanln(&inputLetter)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	temp := []rune(inputLetter)
	if _, err := strconv.Atoi(inputLetter); err == nil {
		fmt.Println("You can't enter a number, please enter a letter")
		temp[0] = enterTheLetter()
	} else if len(temp) != 1 {
		fmt.Println("Enter only one letter")
		temp[0] = enterTheLetter()
	} else {
		fmt.Print("Your word looks like: ")
	}
	return temp[0]

}

func checkTheLetterInWord(word *[]rune, guessedLetters *[]string, guessed *[]rune, currentLength *int) (bool, int) {

	isFound := false
	inputLetter := enterTheLetter()

	for i, letter := range *word {
		if letter == inputLetter && !isElementInSlice(*guessed, inputLetter) { // если буква найдена в слове, то обновить слайс "угаданных букв", добавив новую
			(*guessedLetters)[i] = string((*word)[i]) // добавляем эту новую букву в слайс
			*currentLength++                          // инкрементируем количество отгаданных букв

			isFound = true // флаг того, что буква была отгадана на этом шаге
		} else if (*guessedLetters)[i] != "*" {
			continue // это условие необходимо для того, чтобы не происходило перезаписи уже отгаданных букв в слайсе
		} else {
			(*guessedLetters)[i] = "*"
		}

	}

	printSlice(guessedLetters) // печатаем слайс отгаданных букв (можно сделать через цикл)

	if !isElementInSlice(*guessed, inputLetter) {
		*guessed = append(*guessed, inputLetter)
	} else {
		fmt.Println("Вы уже вводили такую букву...")
		isFound = true
	}

	fmt.Print("Использованные буквы: ")
	printSlice(guessed)

	return isFound, *currentLength
}

func hangThisMan(mistakeCounter int) {
	switch mistakeCounter {
	case 1:
		fmt.Println()
		fmt.Println("---------")
		fmt.Println("||     ||")
		fmt.Println("       ||")
		fmt.Println("       ||")
		fmt.Println("       ||")
		fmt.Println("       ||")
	case 2:
		fmt.Println()
		fmt.Println("---------")
		fmt.Println("||     ||")
		fmt.Println("0      ||")
		fmt.Println("       ||")
		fmt.Println("       ||")
		fmt.Println("       ||")
	case 3:
		fmt.Println()
		fmt.Println("---------")
		fmt.Println("||     ||")
		fmt.Println("0      ||")
		fmt.Println("  o    ||")
		fmt.Println(" /|\\   ||")
		fmt.Println(" /\\    ||")
	case 4:
		fmt.Println()
		fmt.Println("---------")
		fmt.Println("||     ||")
		fmt.Println("0 o !  ||")
		fmt.Println(" /|\\   ||")
		fmt.Println(" /\\    ||")
		fmt.Println("       ||")
	case 5:
		fmt.Println()
		fmt.Println("---------")
		fmt.Println("||     ||")
		fmt.Println(" x     ||")
		fmt.Println("/|\\    ||")
		fmt.Println("/\\     ||")
		fmt.Println("  DEAD ||")
		fmt.Println()
	}

}

func isElementInSlice(slice []rune, element rune) bool {
	for _, value := range slice {
		if value == element {
			return true
		}
	}
	return false
}

func printGuessedWord(length int) []string {

	fmt.Println("Try to guess my word")

	// создадим слайс из звездочек, количество которых равно количеству букв в слове
	var wordMadeOfStars = make([]string, length)

	for i := 0; i < length; i++ {
		wordMadeOfStars[i] = "*"
		fmt.Print(wordMadeOfStars[i], " ")
	}
	fmt.Println()

	return wordMadeOfStars
}

func printSlice[A rune | string](slice *[]A) {
	for _, v := range *slice {
		fmt.Print(string(v), " ")
	}
	fmt.Println()
}

func main() {
	dataSet := fileReader()
	launchingTheGame(dataSet)
}
