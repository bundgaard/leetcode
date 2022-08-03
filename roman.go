package roman_number

import (
	"bufio"
	"io"
	"log"
	"strings"
)

var romanLiteral = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	result := 0

	br := bufio.NewReader(strings.NewReader(s))
	for {
		current, err := br.ReadByte()
		if err != nil {
			if err == io.EOF {
				result += romanLiteral[current]
				break
			}
			log.Println(err)
		}
		peek, _ := br.ReadByte()

		if current == 'I' && (peek == 'V' || peek == 'X') {
			result += romanLiteral[peek] - romanLiteral[current]
		} else if current == 'X' && (peek == 'L' || peek == 'C') {
			result += romanLiteral[peek] - romanLiteral[current]
		} else if current == 'C' && (peek == 'D' || peek == 'M') {
			result += romanLiteral[peek] - romanLiteral[current]
		} else {
			if peek != 0 {
				if err := br.UnreadByte(); err != nil {
					log.Println(err)
				}
			}

			result += romanLiteral[current]
		}
	}

	// ij
	// MCMXCIV

	return result

}
