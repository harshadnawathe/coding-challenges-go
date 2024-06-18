package main // TODO: Move this to it's own package

import (
	"bufio"
	"io"
	"unicode"
)

type Counts struct {
	FileName               string
	Byte, Char, Word, Line int
}

type Input struct {
	Reader   io.Reader
	FileName string
}


func Count(inputs... Input) []Counts {
	if len(inputs) == 0 {
		return nil
	}

	counts := make([]Counts, 0, len(inputs)+1)
	
	for _, input := range inputs {
		c := count(input.FileName, input.Reader)
		counts = append(counts, c)
	}

	if len(inputs) == 1 {
		return counts
	}

	total := Counts{FileName: "total"}
  for _, c := range counts {
    total.Byte += c.Byte
    total.Char += c.Char
    total.Word += c.Word
    total.Line += c.Line
  }

	counts = append(counts, total)

	return counts
}


func count(fileName string, reader io.Reader) Counts {
	c := Counts{FileName: fileName}
	r := bufio.NewReader(reader)

	inWord := false
	for {
		char, size, err := r.ReadRune()
		if err != nil {
			break
		}

		c.Byte += size

		c.Char += 1

		switch char {
		case '\n':
			c.Line += 1
			inWord = false
		case '\r':
			inWord = false
		case '\f':
			inWord = false
		case ' ':
			inWord = false
		case '\t':
			inWord = false
		case '\v':
			inWord = false
		default:
			isCharNotSpace := !unicode.IsSpace(char)
			if !inWord && isCharNotSpace {
				c.Word += 1
			}
			inWord = isCharNotSpace
		}
	}

	return c
}
