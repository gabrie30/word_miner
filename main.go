package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gabrie30/word_miner/trie"
)

func main() {
	fmt.Println("Starting...")
	fmt.Println("To exit the program type exitnow or ctrl + c")
	trie := createTrie()
	userLoop(trie)
	fmt.Println("Exiting...")
}

func userLoop(trie *trie.Node) {
	for {
		fmt.Println("Please enter a word: ")
		reader := bufio.NewReader(os.Stdin)
		value, _ := reader.ReadString('\n')
		value = strings.Replace(value, "\n", "", -1)

		if value == "exitnow" {
			break
		}

		matches, err := trie.WordsBelow(value)

		if err != nil {
			fmt.Println(err)
		} else {
			printMatches(matches)
		}
	}
}

func printMatches(matches *[]string) {
	fmt.Println(matches)
}

func createTrie() *trie.Node {
	return trie.BuildDictionary()
}
