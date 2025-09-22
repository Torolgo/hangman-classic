//===================================================================
//PROJECT : Hangman
//DOMAIN : Game programming
//TITLE : hangman-basic
//DEVELOPERS : Dessenne Ylan
//===================================================================

// -------------------------IMPORTS-----------------------------------

package internal

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

//-------------------------RELEASE-----------------------------------

// chooseWord selects a random word from the French dictionary and returns it
func chooseWord() string {
	fileName, err := os.Open("data/textfile/frenchDictionary.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fileName.Close()
	scanner := bufio.NewScanner(fileName) // Creates a scanner to read the file line by line
	var words []string                    // Initializes a slice to store the words
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	rand.Seed(time.Now().UnixNano()) // Seeds the random number generator with the current time
	chosenWord := words[rand.Intn(len(words))]
	return chosenWord
}
