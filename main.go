package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/asticode/go-texttospeech/texttospeech"
)

func main() {
	var messages []string
	interval(&messages)
	m := ""
	fmt.Println("Enter messeges to self (just enter \"ext\" to exit) ")
	fmt.Println("Type text/s: ")
	fmt.Print("- ")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		m = scanner.Text()
		messages = append(messages, m)
		fmt.Print("- ")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func interval(messages *[]string) {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				if len(*messages) > 0 {
					randomIndex := rand.Intn(len(*messages))
					pick := (*messages)[randomIndex]
					tts := texttospeech.NewTextToSpeech()
					tts.Say(pick)
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
