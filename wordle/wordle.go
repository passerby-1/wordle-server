package wordle

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

const WORD_LENGTH = 5
const MAX_GUESSES = 6

// mainでやっていることが多すぎる

func Wordle() {
	wordList := GenerateWordList(WORD_LENGTH)
	selectedWord := SelectWord(wordList)
	reader := bufio.NewReader(os.Stdin)

	var guesses []map[string][WORD_LENGTH]string

	var guessCount int
	for guessCount = 0; guessCount < MAX_GUESSES; guessCount++ { // MAX_GUESSESの回数繰り返し
		fmt.Printf("Enter your guess (%v/%v): ", guessCount+1, MAX_GUESSES)

		guessWord, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		guessWord = strings.ToUpper(guessWord[:len(guessWord)-1])

		if guessWord == selectedWord { // 正解なら
			fmt.Println("You guessed right!")
			colorVector := GetFilledColorVector("Green")

			guesses = append(guesses, map[string][WORD_LENGTH]string{guessWord: colorVector})

			fmt.Println("Your wordle matrix is: ")
			for _, guess := range guesses {
				for guessWord, colorVector := range guess {
					DisplayWord(guessWord, colorVector)
				}
			}
			break
		} else { // 単語が正解ではなかったら

			i := sort.SearchStrings(wordList, guessWord)
			// 昇順にソートされた文字列スライスwordleWordsと文字列guessWordを受け取り、このスライスの中にguessWordがあるかを二分探索する
			// 存在する場合、その順番を返す (if文で引っかかっていないものだけなので、ここでは存在しない場合のみとなる)
			// 存在しない場合、昇順の文字列スライス中何番目にguessWordを挿入すれば良いかのインデックスを返す
			if i < len(wordList) && wordList[i] == guessWord {
				colorVector := GetFilledColorVector("Grey")

				// stores whether an index is allowed to cause another index to be yellow
				yellowLock := [WORD_LENGTH]bool{}

				for j, guessLetter := range guessWord {
					for k, letter := range selectedWord {
						if guessLetter == letter && j == k {
							colorVector[j] = "Green"
							// now the kth index can no longer cause another index to be yellow
							yellowLock[k] = true
							break

						}
					}
				}
				for j, guessLetter := range guessWord {
					for k, letter := range selectedWord {
						if guessLetter == letter && colorVector[j] != "Green" && yellowLock[k] == false {
							colorVector[j] = "Yellow"
							yellowLock[k] = true
						}
					}
				}
				guesses = append(guesses, map[string][WORD_LENGTH]string{guessWord: colorVector})
				DisplayWord(guessWord, colorVector)
			} else {
				guessCount--
				fmt.Printf("Please guess a valid %v letter word from the wordlist", WORD_LENGTH)
				fmt.Println()
			}
		}
	}

	if guessCount == MAX_GUESSES {
		fmt.Println("Better luck next time!")
		colorVector := GetFilledColorVector("Green")
		fmt.Print("The correct word is : ")
		DisplayWord(selectedWord, colorVector)
	}
}
