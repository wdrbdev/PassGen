package main

import (
	"fmt"

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

var lowAlpha = []string{}
var upAlpha = []string{}
var numbers = []string{}
var specialChars = []string{}

func main() {
	// TODO takes command line arguments

	// TODO to be removed
	for char := range dictionary {
		dictionary[char] = true
	}

	var chars []string
	for char, included := range dictionary {
		if !included {
			continue
		}
		chars = append(chars, char)
	}

	var generate func([]string, int) string
	generate = generator.Shuffle
	var password = generate(chars, 16)
	fmt.Println(password)
}
