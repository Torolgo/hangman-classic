//===================================================================
//PROJECT : Hangman
//DOMAIN : Game programming
//TITLE : hangman-basic
//DEVELOPERS : Dessenne Ylan
//===================================================================

// -------------------------IMPORTS-----------------------------------

package internal

import (
	"fmt"
	"math/rand"
	"time"
)

//-------------------------RELEASE-----------------------------------

// The notFound function checks how many times a guessed letter appears in the target word.
// It returns the number of times the guessed letter appears in the word.
func notFound(word string, guessedLetter rune) int {
	var count = 0
	for _, letter := range word { // Iterates over each letter in the word
		if guessedLetter == letter {
			count += 1
		} else {
			count += 0
		}
	}
	return count
}

//-------------------------------------------------------------------

// The revealLetter function reveals a certain number of letters in the target word.
// It randomly selects positions in the word and marks those letters as guessed.
func revealLetter(word string, guessedLetter map[rune]bool) {
	var n = len(word)/2 - 1
	rand.Seed(time.Now().UnixNano()) // Initializes the random number generator with the current time
	for k := 0; k < n; k++ {         // Loops to reveal 'n' letters
		place := rand.Intn(len(word))
		guessedLetter[rune(word[place])] = true
	}
}

//-------------------------------------------------------------------

// The wordVerification function checks if the guessed word is identical to the target word.
func wordVerification(word string, guessedWord string) bool {
	return word == guessedWord
}

//-------------------------------------------------------------------

// The wordReal function checks if the guessed word consists only of lowercase letters.
func wordReal(guessedWord string) bool {
	for i := 0; i < len(guessedWord); i++ { // Iterates over each character in the guessed word
		if guessedWord[i] < 97 || guessedWord[i] > 122 { // Checks if the character is not a lowercase letter
			return false
		}
	}
	return true
}

//-------------------------------------------------------------------

// The testedLetters function checks if the chosen letter has already been tried by the player.
func testedLetters(testLetter []string, choice string) bool {
	for i := 0; i < len(testLetter); i++ { // Iterates over each element in the testLetter slice
		if testLetter[i] == choice {
			fmt.Println("You have already tried this ! ")
			return false
		}
	}
	return true
}

//-------------------------------------------------------------------

// The stopDetection function checks if the word STOP is entered by the player,
// asks for confirmation, and then stops or continues the game accordingly
func stopDetection(choice string) string {
	if choice == "STOP" {
		fmt.Println("You want to leave me in the middle of the game")
		var confirm string
		fmt.Println("Yes or No")
		fmt.Scan(&confirm)
		if confirm == "Yes" {
			return "Yes"
		} else if confirm == "No" {
			return "No"
		} else {
			return "err"
		}
	}
	return "no stop"
}
