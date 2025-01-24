package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Функция для старта игры
func start() bool {
	var command string
	var isGameStarted bool

	fmt.Println("Press [N]ew to Start game or [E]nd to kill this program")
	fmt.Scanln(&command)

	switch command {
	case "N":
		flag = true
	case "n":
		flag = true
	case "E":
		flag = false
	case "e":
		flag = false
	default:
		fmt.Println("Command not recognized, try again")
		flag = start()
	}

	return flag
}

// Чтение файла со словами
func fileReader() []string {
	file, err := os.Open("words.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	var dataWords []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		dataWords = append(dataWords, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return dataWords
}

func gamePlay(word []string) {
	//fmt.Println(word)
	fmt.Println("Try to guess my word")
	wordLength := len(word)

	for i := 0; i < len(word); i++ {
		fmt.Print("* ")

	}
	fmt.Println()

	// создадим слайс из звездочек, количество которых равно количеству букв в слове
	var guessedLetters = make([]string, wordLength)
	for i := 0; i < wordLength; i++ {
		guessedLetters[i] = "*"
	}
	var guessed []string
	var mistakeCounter int
	var currentLength int // текущая количество отгаданных букв

	// цикл, котором сравнивается написанная буква с буквами загаданного слова
	for {
		var inputLetter string
		isFound := false
		fmt.Print("Enter letter: ")
		fmt.Scanln(&inputLetter)
		fmt.Print("Your word looks like: ")
		guessed = append(guessed, inputLetter)

		for i, letter := range word {
			if letter == inputLetter { // если буква найдена в слове, то обновить слайс "угаданных букв", добавив новую
				guessedLetters[i] = word[i] // добавляем эту новую букву в слайс
				currentLength++             // инкрементируем количество отгаданных букв

				isFound = true // флаг того, что буква была отгадана на этом шаге
			} else if guessedLetters[i] != "*" {
				continue // это условие необходимо для того, чтобы не происходило перезаписи уже отгаданных букв в слайсе
			} else {
				guessedLetters[i] = "*"
			}

		}

		printSlice(guessedLetters) // печатаем слайс отгаданных букв (можно сделать через цикл)
		fmt.Print("Использованные буквы: ")
		printSlice(guessed)

		if !isFound {
			mistakeCounter++
			mistakes := hangThisMan(mistakeCounter)
			if mistakes == 0 {
				fmt.Println("Ты проиграл :(")
				fmt.Print("Загаданное слово: ")
				printSlice(word)
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

func printSlice(slice []string) {
	for _, v := range slice {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func hangThisMan(mistakeCounter int) int {

	switch mistakeCounter {
	case 1:
		fmt.Println()
		fmt.Println("---------")
		fmt.Println("||     ||")
		fmt.Println("       ||")
		fmt.Println("       ||")
		fmt.Println("       ||")
		fmt.Println("       ||")
		return 1
	case 2:
		fmt.Println()
		fmt.Println("---------")
		fmt.Println("||     ||")
		fmt.Println("0      ||")
		fmt.Println("       ||")
		fmt.Println("       ||")
		fmt.Println("       ||")
		return 2
	case 3:
		fmt.Println()
		fmt.Println("---------")
		fmt.Println("||     ||")
		fmt.Println("0      ||")
		fmt.Println("  o    ||")
		fmt.Println(" /|\\   ||")
		fmt.Println(" /\\    ||")
		return 3
	case 4:
		fmt.Println()
		fmt.Println("---------")
		fmt.Println("||     ||")
		fmt.Println("0 o !  ||")
		fmt.Println(" /|\\   ||")
		fmt.Println(" /\\    ||")
		fmt.Println("       ||")
		return 4
	case 5:
		fmt.Println()
		fmt.Println("---------")
		fmt.Println("||     ||")
		fmt.Println(" x     ||")
		fmt.Println("/|\\    ||")
		fmt.Println("/\\     ||")
		fmt.Println("  DEAD ||")
		fmt.Println()
		return 0
	default:
		return 666
	}

}

func game() {
	for {
		words := fileReader()
		rand.Seed(time.Now().UTC().UnixNano())
		word := words[rand.Intn(len(words)-1)]
		var secretWord []string
		for _, letter := range word {
			secretWord = append(secretWord, string(letter))
		}

		startingVar := start()
		if !startingVar {
			return
		}
		gamePlay(secretWord)
	}
}

func main() {
	game()
}
