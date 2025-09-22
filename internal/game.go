//===================================================================
//PROJECT : Hangman
//DOMAIN : Game programming
//TITLE : hangman-basic
//DEVELOPERS : Dessenne Ylan
//===================================================================

// -------------------------IMPORTS-----------------------------------

package internal

import "fmt"

//-------------------------RELEASE-----------------------------------

// StandardGame starts the Hangman game
func StandardGame() {
	var finished = false
	guessedLetters := make(map[rune]bool) // Map to track guessed letters
	var word = chooseWord()
	var testLetter []string
	var triesNumber int = 10
	var startLine int = 1
	var endLine int = 8

	revealLetter(word, guessedLetters)
	clear()

	for finished == false {
		var valid = false
		for valid == false {
			printJose(startLine, endLine)
			fmt.Println("You have", triesNumber, "chance to found the word !!!")
			fmt.Print("THE GUESS WORD :")
			fmt.Println(getWord(word, guessedLetters)) // Displays the word with guessed letters
			fmt.Println("The letters you tested")
			fmt.Println(testLetter)
			var choice string
			fmt.Println("Enter your choice: ")
			fmt.Scan(&choice) // Reads the player's choice

			// Checks if the player wants to stop the game
			stopResult := stopDetection(choice)
			if stopResult == "Yes" {
				return
			} else if stopResult == "err" {
				clear()
				fmt.Println("Error : You have to say yes or no. Return to game.")
			} else if stopResult == "No" {
				clear()
				fmt.Println("You have to say no. Return to game.")
			} else if stopResult == "no stop" {

				if len(choice) > len(word) {
					clear()
					fmt.Println("Your choice is too long")
				} else {
					if len(choice) == 1 {
						if choice[0] >= 97 && choice[0] <= 122 { // Checks if the choice is a lowercase letter
							clear()
							if testedLetters(testLetter, choice) == true {

								if notFound(word, rune(choice[0])) == 0 { // Checks if the letter is not in the word
									// Decreases the number of tries by 1
									triesNumber--
									startLine += 8
									endLine += 8
								}
								guessedLetters[rune(choice[0])] = true // Adds the letter to guessed letters
								valid = true
								testLetter = append(testLetter, string(rune(choice[0]))) // Adds the letter to tested letters
							}
						} else {
							clear()
							fmt.Println("Error: This is not a letter")
						}
					} else {
						if wordReal(choice) == true && wordVerification(word, choice) == true { // Checks if the word is valid and correct
							fmt.Println("Well done, you've found the word. You have WON !!!!")
							return
						} else if wordReal(choice) == true && wordVerification(word, choice) == false {
							// Decreases the number of tries by 2
							startLine += 8
							endLine += 8
							triesNumber -= 2
							startLine += 8
							endLine += 8
							clear()
							if triesNumber == 0 { // Checks if the number of tries is exhausted
								fmt.Println("You have reached the maximum number of tries", "GAME OVER")
								fmt.Println("The word was : ", word)
								return
							}
						} else if wordReal(choice) == false {
							clear()
							fmt.Println("Error: This is not a word")
						}
					}
				}
			}
		}
		if triesNumber == 0 {
			fmt.Println("You have reached the maximum number of tries", "GAME OVER")
			fmt.Println("The word was : ", word)
			return
		}
		finished = getWord(word, guessedLetters) == word // Checks if the word has been fully guessed
	}
	fmt.Println("Well done, you've found the word. You have WON !!!!")
}
