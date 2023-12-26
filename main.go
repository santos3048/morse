package main

import (
	"github.com/eiannone/keyboard"
	"time"
	"os"
	"strconv"
	//"fmt"
)

func main() {
	if len(os.Args) > 2  {
		i, err := strconv.Atoi(os.Args[2])
		if err != nil{
			panic(err)
		}
		ToMorse(i)
	}
	Start()
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
			go Print(morse)
			morse = ""
			f = time.Time{}
		}
	}
}
