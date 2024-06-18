package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	LogFlag = log.LstdFlags | log.Lshortfile
	Log     = log.New(io.Discard, "", LogFlag)
)

func wcUsage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] [file ...]\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr)
}

var printCountFlags PrintCountFlags = 0

func init() {
	if os.Getenv("WC_LOG") == "1" {
		Log = log.New(os.Stderr, "", LogFlag)
	}

	flag.Usage = wcUsage

	flag.BoolFunc("c", "Print the number of bytes in the input", func(_ string) error {
		printCountFlags &= ^PrintCountChars
		printCountFlags |= PrintCountBytes
		return nil
	})
	flag.BoolFunc("m", "Print the number of chars in the input", func(_ string) error {
		printCountFlags &= ^PrintCountBytes
		printCountFlags |= PrintCountChars
		return nil
	})
	flag.BoolFunc("w", "Print the number of words in the input", func(_ string) error {
		printCountFlags |= PrintCountWords
		return nil
	})
	flag.BoolFunc("l", "Print the number of lines in the input", func(_ string) error {
		printCountFlags |= PrintCountLines
		return nil
	})
}

func main() {
	flag.Parse()

	if printCountFlags == 0 {
		Log.Println("using default show count flags")
		printCountFlags = PrintCountDefault
	}
	Log.Println("show counts for: ", printCountFlags)

	var inputs []Input

	if flag.NArg() == 0 {
		inputs = append(inputs, Input{Reader: os.Stdin})
	} else {
		for _, arg := range flag.Args() {
			inputFile, err := os.Open(arg)
			if err != nil {
				fmt.Println("cannot open file: ", err)
				Log.Fatalf("cannot open file: %v", err)
			}
			Log.Println("opened input file: ", arg)

			defer func() {
				inputFile.Close()
				Log.Println("closed input file: ", arg)
			}()

			inputs = append(inputs, Input{inputFile, inputFile.Name()})
		}
	}

	Log.Printf("inputs: %v", inputs)

	Print(printCountFlags, Count(inputs...))
}
