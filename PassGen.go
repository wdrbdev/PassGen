package main

import (
	"flag"
	"fmt"
	"runtime"

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
	printUsage  bool
	hasUpAlpha  *bool
	hasLowAlpha *bool
	hasDigit    *bool
	hasChar     *bool
	unsimilar   *bool
	unique      *bool
	length      *int
)

func init() {
	flag.BoolVar(&printUsage, "help", false, "Print usage")
	flag.BoolVar(&printUsage, "h", false, "Print usage (shorthand)")
	hasUpAlpha = flag.Bool("upper", false, "Include upper case letters")
	hasLowAlpha = flag.Bool("lower", false, "Include lower case letters")
	hasDigit = flag.Bool("digit", false, "Include digits")
	hasChar = flag.Bool("char", false, "Include special characters")
	unsimilar = flag.Bool("unsimilar", false, "Exclude similar characters (0oO1lI...)")
	unique = flag.Bool("unique", false, "Exclude duplicate characters")
	length = flag.Int("length", 16, "Length of password generated")
}

func main() {
	flag.Parse()

	if printUsage {
		_, fileName, _, _ := runtime.Caller(0)
		fmt.Println("# --- Usage of", fileName, "--- #")
		flag.PrintDefaults()
		return
	}

	if *length <= 0 {
		// TODO use log.fatal instead of panic
		panic("length must be positive integer")
	}

	// Select user-chosen characters according to flag
	if !*hasUpAlpha && !*hasLowAlpha && !*hasDigit && !*hasChar {
		dictMarker.Alphanumeric(dictionary)
	} else {
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
		// TODO use log.fatal instead of panic
		panic("Available characters is less than chosen length.")
	}

	// Select generator algorithm according to flag
	var generate func([]string, int) string
	if !*unique {
		generate = generator.RandIdx
	} else {
		generate = generator.RandIdxDeduplicate
	}

	// generate password
	// TODO generate n password concurrently
	var password = generate(chars, *length)
	fmt.Println(password)
}
