package no3-refactoring

import (
	"fmt"
	"strings"
)

func isNotNull(str string) bool{
	return len(str) > 0 
}

func firstBracketFound (str string) string{
	indexFirstBracketFound := strings.Index(str ,"(")
	if indexFirstBracketFound >= 0 {
		runes := []rune(str)
		
		wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
		return wordsAfterFirstBracket 
		
	}else{
		return ""
	}
}

func findFirstStringInBracket(str string) string {
	getIsNotNull := isNotNull(str)
	indexClosingBracketFound := strings.Index(firstBracketFound (str),")")
	
	if getIsNotNull && indexClosingBracketFound >= 0 {
		runes := []rune(firstBracketFound (str))
		return string(runes[1:indexClosingBracketFound-1])
	}else{
		return ""
	}		
		
}

//Fungsi dibawaw merupakan fungsi sebelum dilakukan refactoring
func oldfindFirstStringInBracket(str string) string { 
	if (len(str) > 0) {
		indexFirstBracketFound := strings.Index(str,"(") 
		if indexFirstBracketFound >= 0 {
			runes := []rune(str)
			wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)]) 
			indexClosingBracketFound := strings.Index(wordsAfterFirstBracket,")") 
			if indexClosingBracketFound >= 0 {
				runes := []rune(wordsAfterFirstBracket)
				return string(runes[1:indexClosingBracketFound-1])

			}else{
				return ""

			}

		}else{
			return ""

		}
	}else{
		return ""
	}
}



func main() {
	var words string = "Saya tidak (pernah) melakukan perbuatan tercela (itu)"
	fmt.Println("After Refactor: " + findFirstStringInBracket(words))
	fmt.Println("Before Refactor: " + oldfindFirstStringInBracket(words))

	/**Result */
	/** perna (karena string -1 pada line 31 dan 48) */
}
