package batch

import (
	"errors"
	"strings"
	"unicode"
)

// Split breaks a command line into arguments, like zmprov does for the lines of a -f file.
// Arguments are separated by whitespace. Single quotes keep their content as is, double quotes
// allow \" and \\ escapes, and a backslash outside quotes escapes the next character.
// An empty pair of quotes gives an empty argument.
func Split(line string) ([]string, error) {
	var args []string
	var current strings.Builder

	inArg := false
	runes := []rune(line)

	for i := 0; i < len(runes); i++ {
		r := runes[i]

		switch {
		case unicode.IsSpace(r):
			if inArg {
				args = append(args, current.String())
				current.Reset()
				inArg = false
			}
		case r == '\'':
			inArg = true
			end := indexRune(runes, i+1, '\'')

			if end < 0 {
				return nil, errors.New("unterminated single quote")
			}

			current.WriteString(string(runes[i+1 : end]))
			i = end
		case r == '"':
			inArg = true
			i++

			for ; i < len(runes) && runes[i] != '"'; i++ {
				if runes[i] == '\\' && i+1 < len(runes) && (runes[i+1] == '"' || runes[i+1] == '\\') {
					i++
				}

				current.WriteRune(runes[i])
			}

			if i >= len(runes) {
				return nil, errors.New("unterminated double quote")
			}
		case r == '\\':
			inArg = true

			if i+1 >= len(runes) {
				return nil, errors.New("trailing backslash")
			}

			i++
			current.WriteRune(runes[i])
		default:
			inArg = true
			current.WriteRune(r)
		}
	}

	if inArg {
		args = append(args, current.String())
	}

	return args, nil
}

func indexRune(runes []rune, from int, target rune) int {
	for i := from; i < len(runes); i++ {
		if runes[i] == target {
			return i
		}
	}

	return -1
}
