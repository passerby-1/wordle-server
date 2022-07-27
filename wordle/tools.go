package wordle

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"sort"
	"strings"
	"time"
)

const WORDS_URL = "https://raw.githubusercontent.com/dwyl/english-words/master/words_alpha.txt"

func GetFilledColorVector(color string) [WORD_LENGTH]string {
	colorVector := [WORD_LENGTH]string{}
	for i := range colorVector {
		colorVector[i] = color
	}
	return colorVector
}

func DisplayWord(word string, colorVector [WORD_LENGTH]string) {
	for i, c := range word {
		switch colorVector[i] {
		case "Green":
			fmt.Print("\033[42m\033[1;30m")
		case "Yellow":
			fmt.Print("\033[43m\033[1;30m")
		case "Grey":
			fmt.Print("\033[40m\033[1;37m")
		}
		fmt.Printf(" %c ", c)
		fmt.Print("\033[m\033[m")
	}
	fmt.Println()
}

func GenerateWordList(wordLength int) []string {

	res, err := http.Get(WORDS_URL)
	if err != nil {
		log.Fatalln(err)
	}

	body, _ := ioutil.ReadAll(res.Body)
	words := strings.Split(string(body), "\r\n")

	var wordleWords []string // 指定された文字長の単語のみを格納するスライス
	for _, word := range words {
		if len(word) == wordLength {
			wordleWords = append(wordleWords, strings.ToUpper(word))
		}
	}
	sort.Strings(wordleWords)

	return wordleWords

}

func SelectWord(wordList []string) string {
	rand.Seed(time.Now().Unix())
	selectedWord := wordList[rand.Intn(len(wordList))]
	return selectedWord
}

// ポインタ渡しじゃなくて良いの? と思うかもしれないけど、
// スライス (可変長) はそれ自体が配列へのポインタを持っているので、別にこれで良いらしい
// 配列 (固定長) は、長さによってはポインタ渡しの方が良い
