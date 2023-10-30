package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/PassGen/concurrent"
	"github.com/PassGen/dictMarker"
	"github.com/PassGen/generator"
)

var dictionary = map[string]bool{
	"a":  false,
	"b":  false,
	"c":  false,
	"d":  false,
	"e":  false,
	"f":  false,
	"g":  false,
	"h":  false,
	"i":  false,
	"j":  false,
	"k":  false,
	"l":  false,
	"m":  false,
	"n":  false,
	"o":  false,
	"p":  false,
	"q":  false,
	"r":  false,
	"s":  false,
	"t":  false,
	"u":  false,
	"v":  false,
	"w":  false,
	"x":  false,
	"y":  false,
	"z":  false,
	"A":  false,
	"B":  false,
	"C":  false,
	"D":  false,
	"E":  false,
	"F":  false,
	"G":  false,
	"H":  false,
	"I":  false,
	"J":  false,
	"K":  false,
	"L":  false,
	"M":  false,
	"N":  false,
	"O":  false,
	"P":  false,
	"Q":  false,
	"R":  false,
	"S":  false,
	"T":  false,
	"U":  false,
	"V":  false,
	"W":  false,
	"X":  false,
	"Y":  false,
	"Z":  false,
	"1":  false,
	"2":  false,
	"3":  false,
	"4":  false,
	"5":  false,
	"6":  false,
	"7":  false,
	"8":  false,
	"9":  false,
	"0":  false,
	"!":  false,
	"\"": false,
	"#":  false,
	"$":  false,
	"%":  false,
	"&":  false,
	"'":  false,
	"(":  false,
	")":  false,
	"*":  false,
	"+":  false,
	",":  false,
	"-":  false,
	".":  false,
	"/":  false,
	":":  false,
	";":  false,
	"<":  false,
	"=":  false,
	">":  false,
	"?":  false,
	"@":  false,
	"[":  false,
	"\\": false,
	"]":  false,
	"^":  false,
	"_":  false,
	"`":  false,
	"{":  false,
	"|":  false,
	"}":  false,
	"~":  false,
}

var (
	fileName      string
	printUsage    bool
	passwordCount int
	hasUpAlpha    *bool
	hasLowAlpha   *bool
	hasDigit      *bool
	hasChar       *bool
	unsimilar     *bool
	unique        *bool
	length        *int
	delimiter     *string
	output        *string
)

func init() {
	_, fileName, _, _ = runtime.Caller(0)
	fileName = filepath.Base(fileName)

	flag.BoolVar(&printUsage, "help", false, "Print usage")
	flag.BoolVar(&printUsage, "h", false, "Print usage (shorthand)")
	flag.IntVar(&passwordCount, "number", 1, "Total count of passwords generated")
	flag.IntVar(&passwordCount, "n", 1, "Total count of passwords generated (shorthand)")
	hasUpAlpha = flag.Bool("upper", false, "Include upper case letters")
	hasLowAlpha = flag.Bool("lower", false, "Include lower case letters")
	hasDigit = flag.Bool("digit", false, "Include digits")
	hasChar = flag.Bool("char", false, "Include special characters")
	unsimilar = flag.Bool("unsimilar", false, "Exclude similar characters (0oO1lI...)")
	unique = flag.Bool("unique", false, "Exclude duplicate characters")
	length = flag.Int("length", 16, "Length of password generated. Default is 16.")
	delimiter = flag.String("delimiter", "\n", "Delimiter of passwords generated. Default is `\n`.")
	output = flag.String("output", "", "The file path as output destination. Default is stdout.")

	flag.Parse()
}

func main() {
	if printUsage {
		fmt.Println("["+fileName+"]:", "Generate random passwords offline.")
		fmt.Println("# --- Usage of", fileName, "--- #")
		flag.PrintDefaults()
		return
	}

	if *length <= 0 {
		log.Fatal("[", fileName, "] ", "length must be non-zero positive integer")
	}

	// Select user-chosen characters according to flag
	if *hasUpAlpha {
		dictMarker.Upper(dictionary)
	}
	if *hasLowAlpha {
		dictMarker.Lower(dictionary)
	}
	if *hasDigit {
		dictMarker.Digit(dictionary)
	}
	if *hasChar {
		dictMarker.Char(dictionary)
	}
	if !*hasUpAlpha && !*hasLowAlpha && !*hasDigit && !*hasChar {
		dictMarker.Alphanumeric(dictionary)
	}
	if *unsimilar {
		dictMarker.Unsimilar(dictionary)
	}

	var chars []string
	for char, included := range dictionary {
		if !included {
			continue
		}
		chars = append(chars, char)
	}
	if *unique && len(chars) < *length {
		log.Fatal("[", fileName, "] ", "Available characters is less than chosen length.")
	}

	// Select generator algorithm according to flag
	var generate func([]string, int) string
	if !*unique {
		generate = generator.RandIdx
	} else {
		generate = generator.RandIdxDeduplicate
	}

	// select output destination
	var writer = os.Stdout
	if *output != "" {
		var err error
		if _, err := os.Stat(*output); err == nil {
			log.Fatal("[", fileName, "] ", "file `", *output, "` already existed or ", err)
		} else if !errors.Is(err, os.ErrNotExist) {
			log.Fatal("[", fileName, "] ", err)
		}

		writer, err = os.Create(*output)
		if err != nil {
			log.Fatal("[", fileName, "] ", err)
		}
		defer writer.Close()
	}

	// generate password
	passwordChan := make(chan string, 16)
	stopChan := make(chan bool)
	start := time.Now()
	concurrent.Generator(generate, chars, *length, passwordChan, stopChan)
	// concurrent.FanIn(generate, chars, *length, passwordChan, stopChan)
	// concurrent.WorkerPool(generate, chars, *length, passwordChan, stopChan)
	for i := 0; i < passwordCount; i++ {
		password := <-passwordChan
		fmt.Fprint(writer, password)
		if i != passwordCount-1 {
			fmt.Fprint(writer, *delimiter)
		}
	}
	stopChan <- true
	fmt.Printf("\nIt took %v to generate %v passwords.\n", time.Since(start), passwordCount)
}
