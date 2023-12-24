package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"time"
)

var alphaNumToMorse = map[string]string{
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

var morseToAlphaNum map[string]string

func start() {
	morseToAlphaNum = map[string]string{}
	for k, v := range alphaNumToMorse {
		morseToAlphaNum[v] = k
	}
}

func print(morse string) {
	vl, found := morseToAlphaNum[morse]
	if !found {
		fmt.Print("\\u")
	} else {
		fmt.Print(vl)
	}

}

func main() {
	start()
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	var f time.Time
	var s time.Duration
	var morse string
	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if key == keyboard.KeyEnter {
			if f.IsZero() {
				f = time.Now()
			} else {
				s = time.Since(f)
				f = time.Now()
				//fmt.Println(s)
				if s.Milliseconds() > 300 {
					morse += "-"
				} else {
					morse += "."
				}
			}
		} else {
			s = time.Since(f)
			f = time.Now()
			//fmt.Println(s)
			if s.Milliseconds() > 300 {
				morse += "-"
			} else {
				morse += "."
			}
			go print(morse)
			morse = ""
			f = time.Time{}
		}
	}
}
