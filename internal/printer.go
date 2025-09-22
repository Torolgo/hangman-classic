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
	"fmt"
	"os"
	"os/exec"
)

//-------------------------RELEASE-----------------------------------

// The printJose function displays the progress of the hangman based on its arguments
func printJose(startLine int, endLine int) {
	file, err := os.Open("data/textfile/hangman.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file) // Creates a scanner to read the file line by line
	var lines []string
	currentLine := 1
	for scanner.Scan() {
		if currentLine >= startLine && currentLine <= endLine { // Checks if the current line is within the specified range
			lines = append(lines, scanner.Text())
		}
		currentLine++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
	}
	for _, line := range lines {
		fmt.Println(line)
	}
}

//-------------------------------------------------------------------

// The clear function clears the terminal screen on Windows
func clear() {
	cmd := exec.Command("cmd", "/c", "cls") // Redirects the command's standard output to the program's standard output
	cmd.Stdout = os.Stdout
	cmd.Run()
}

//-------------------------------------------------------------------

// The getWord function returns the word with guessed letters and underscores
func getWord(word string, guessedLetter map[rune]bool) string {
	var str = ""
	for _, letter := range word { // Iterates over each letter in the word
		if guessedLetter[letter] { // Checks if the letter has been guessed
			str += string(letter)
		} else {
			str += "_"
		}
	}
	return str
}
