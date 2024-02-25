package app

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	"log"

	"github.com/google/generative-ai-go/genai"
	"github.com/said-boy/talkify/api"
	"github.com/said-boy/talkify/assets"
	"github.com/said-boy/talkify/controller"
	entitiy "github.com/said-boy/talkify/entity"
)

func Run() {
	assets.ShowBanner()
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("Recording...")
		controller.StartRecord()
		
		fmt.Println("Wait...")
		resp := api.SendWitVoice("data/rec.wav")

		var entityWit entitiy.Wit

		err := json.Unmarshal([]byte(resp), &entityWit)
		if err != nil {
			log.Fatal(err)
		}

		message := entityWit.TEXT

		if message == "Exit" || message == "exit" {
			os.Exit(1)
		}

		fmt.Println("You : ", message)

		var aiText genai.Part
		for _, v := range api.SendGeminiMessage(message) {
			aiText = v.Parts[0]
		}

		fullText := fmt.Sprintf("%v", aiText)

		fmt.Println("Ai Answered...")
		api.GetWitVoiceFromMessage(fullText)

		fmt.Println("Ai : ", aiText)
		api.GeminiToSayAnswer()
		fmt.Println(" ")
	}
}
