package trie

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// Node represents a node of the trie
type Node struct {
	Word     bool
	Children map[string]*Node
}

func newNode() *Node {
	var n Node
	n.Children = make(map[string]*Node)
	return &n
}

// WordsBelow finds all words below a word in a trie
func (n *Node) WordsBelow(word string) (*[]string, error) {
	ln, err := lastNode(word, n)

	if err != nil {
		return nil, err
	}
	// will need to take word and append it to all the letter sequesnces found
	return depthFirstWord(ln, word), nil
}

func depthFirstWord(node *Node, word string) *[]string {
	strings := &[]string{}
	suffix := word
	return depthFirstWordHelper(node, suffix, strings)
}

func depthFirstWordHelper(node *Node, suffix string, strings *[]string) *[]string {

	if len(node.Children) == 0 {
		return strings
	}

	for letter, node := range node.Children {
		if node.Word {
			fmt.Println("Found a word: " + suffix + letter)
			*strings = append(*strings, suffix+letter)
		}
		depthFirstWordHelper(node, suffix+letter, strings)
	}

	return strings
}

func lastNode(word string, n *Node) (*Node, error) {
	letters := strings.Split(word, "")
	trie := n
	for _, letter := range letters {

		if val, ok := trie.Children[letter]; ok {
			trie = val
		} else {
			return nil, errors.New("Word does not exist")
		}
	}

	return trie, nil
}

func createTrie() *Node {
	fmt.Println("Im Trieing")
	words := getWords()
	trie := newNode()
	for i := range words {
		word := words[i]
		fmt.Println("Attempting to add: " + word)
		// For each word, add it to the trie by calling insert, starting with the root node
		insert(word, trie)
	}

	return trie
}

func insert(word string, trie *Node) {
	letters := strings.Split(word, "")
	for i, l := range letters {
		seeking := strings.ToLower(l)
		fmt.Println("-- letter: " + seeking)
		_, ok := trie.Children[seeking]

		if ok {
			// letter is in node, go into that node
			trie = trie.Children[seeking]
		} else {
			// add letter to trie and continue
			if i == len(word)-1 {
				// this is the last letter of the word and node should be marked as a word

				n := newNode()
				n.Word = true
				trie.Children[l] = n

			} else {
				trie.Children[l] = newNode()
			}
			trie = trie.Children[l]
		}
	}
	fmt.Println("Added : " + word)
	// see if word is in the trie, or if you hit the end of a trie build the new word
}

func getWords() []string {
	words := []string{}

	file, err := os.Open("/usr/share/dict/words")
	// file, err := os.Open("test-words")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return words
}

// BuildDictionary creates a trie containing all words in Dictionary
func BuildDictionary() *Node {
	return createTrie()
}
