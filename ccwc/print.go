package main // TODO: Move it to its own package

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"
)

type PrintCountFlags uint8

func (f PrintCountFlags) String() string {
	var b strings.Builder
	m := map[PrintCountFlags]string{
		PrintCountBytes: "bytes",
		PrintCountChars: "chars",
		PrintCountLines: "lines",
		PrintCountWords: "words",
	}

	for k, v := range m {
		if f&k != 0 {
			b.WriteString(v)
			b.WriteByte(' ')
		}
	}

	return b.String()
}

const (
	PrintCountBytes PrintCountFlags = 1 << iota
	PrintCountLines
	PrintCountWords
	PrintCountChars

	PrintCountDefault = PrintCountBytes | PrintCountLines | PrintCountWords
)

func Print(flags PrintCountFlags, cs []Counts) {
	print(os.Stdout, flags, cs)
}

func print(w io.Writer, flags PrintCountFlags, cs []Counts) {
	tw := tabwriter.NewWriter(w, 8, 8, 2, ' ', tabwriter.AlignRight)

	for _, c := range cs {

		if PrintCountLines&flags != 0 {
			fmt.Fprintf(tw, "%d\t", c.Line)
		}

		if PrintCountWords&flags != 0 {
			fmt.Fprintf(tw, "%d\t", c.Word)
		}

		if PrintCountChars&flags != 0 {
			fmt.Fprintf(tw, "%d\t", c.Char)
		}

		if PrintCountBytes&flags != 0 {
			fmt.Fprintf(tw, "%d\t", c.Byte)
		}

		fmt.Fprintf(tw, " %s\n", c.FileName)
	}

	tw.Flush()
}
