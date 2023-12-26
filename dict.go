package main

import "fmt"

var AlphaNumToMorse = map[string]string{
	"A":  ".-",
	"B":  "-...",
	"C":  "-.-.",
	"D":  "-..",
	"E":  ".",
	"F":  "..-.",
	"G":  "--.",
	"H":  "....",
	"I":  "..",
	"J":  ".---",
	"K":  "-.-",
	"L":  ".-..",
	"M":  "--",
	"N":  "-.",
	"O":  "---",
	"P":  ".--.",
	"Q":  "--.-",
	"R":  ".-.",
	"S":  "...",
	"T":  "-",
	"U":  "..-",
	"V":  "...-",
	"W":  ".--",
	"X":  "-..-",
	"Y":  "-.--",
	"Z":  "--..",
	"1":  ".----",
	"2":  "..---",
	"3":  "...--",
	"4":  "....-",
	"5":  ".....",
	"6":  "-....",
	"7":  "--...",
	"8":  "---..",
	"9":  "----.",
	"0":  "-----",
	".":  ".-.-.-",  // period
	":":  "---...",  // colon
	",":  "--..--",  // comma
	";":  "-.-.-",   // semicolon
	"?":  "..--..",  // question
	"=":  "-...-",   // equals
	"'":  ".----.",  // apostrophe
	"/":  "-..-.",   // slash
	"!":  "-.-.--",  // exclamation
	"-":  "-....-",  // dash
	"_":  "..--.-",  // underline
	"\"": ".-..-.",  // quotation marks
	"(":  "-.--.",   // parenthesis (open)
	")":  "-.--.-",  // parenthesis (close)
	"()": "-.--.-",  // parentheses
	"$":  "...-..-", // dollar
	"&":  ".-...",   // ampersand
	"@":  ".--.-.",  // at
	"+":  ".-.-.",   // plus
	"Á":  ".--.-",   // A with acute accent
	"Ä":  ".-.-",    // A with diaeresis
	"É":  "..-..",   // E with acute accent
	"Ñ":  "--.--",   // N with tilde
	"Ö":  "---.",    // O with diaeresis
	"Ü":  "..--",    // U with diaeresis
	" ":  ".......", // word interval
}

var MorseToAlphaNum map[string]string

func Start() {
	MorseToAlphaNum = map[string]string{}
	for k, v := range AlphaNumToMorse {
		MorseToAlphaNum[v] = k
	}
}

func Print(morse string) {
	vl, found := MorseToAlphaNum[morse]
	if !found {
		fmt.Print("\\u")
	} else {
		fmt.Print(vl)
	}
}
