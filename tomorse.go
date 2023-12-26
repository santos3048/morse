package main

import (
	"fmt"
	//"math"
	"github.com/eiannone/keyboard"
	"time"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/mp3"
	"os"
)

func ToMorse(interval int){
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	for {
		rn, _, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		fmt.Print(string(rn))
		val, ok := AlphaNumToMorse[string(rn)]
		if !ok{
			continue
		}
		for _, char := range val{
			if string(char) == "."{
				Beep(interval)
			} else {
				Beep(2 * interval)
			}
		}
	}
}

func Beep(duration int){
	f, _ := os.Open("beep.mp3")
	defer f.Close()
	streamer, format, _ := mp3.Decode(f)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)
	defer streamer.Close()
	time.Sleep(time.Duration(duration) * time.Millisecond)
}