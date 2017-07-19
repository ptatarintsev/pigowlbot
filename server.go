package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"pigowlbot/token"
	"strings"

	"gopkg.in/telegram-bot-api.v4"
)

type PackPhrase struct {
	Phrase      string         `json:"phrase"`
	Complexity  float32        `json:"complexity"`
	Description string         `json:"description"`
	Reviews     map[string]int `json:"reviews"`
}

type Pack struct {
	ID          int          `json:"id"`
	Language    string       `json:"language"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Phrases     []PackPhrase `json:"phrases"`
	Version     int          `json:"version"`
	Paid        bool         `json:"paid"`
}

type PackResponse struct {
	Pack  Pack `json:"pack"`
	Count int  `json:"count"`
}

type GetPacksResponse struct {
	Packs  []PackResponse `json:"packs"`
}

func MainHandler(resp http.ResponseWriter, _ *http.Request) {
	resp.Write([]byte("Hi there! I'm PigowlTestBot!"))
}

func main() {
	url := "http://pigowl.com:8080/getPacks"

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	
	response := new(GetPacksResponse) 
	json.NewDecoder(res.Body).Decode(response)

	var parts []string
	for _, pack := range response.Packs {
		parts = append(parts, pack.Pack.Name)
	}
	log.Printf("%s", strings.Join(parts," "))

	bot, err := tgbotapi.NewBotAPI(token.BotToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	//_, err = bot.SetWebhook(tgbotapi.NewWebhookWithCert("https://www.pigowl.com:8443/"+bot.Token, "fullchain.pem"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//u := tgbotapi.NewUpdate(0)
	//u.Timeout = 60

	//updates, err := bot.GetUpdatesChan(u)
	updates := bot.ListenForWebhook("/" + token.BotToken)

	http.HandleFunc("/", MainHandler)
	go http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	//updates := bot.ListenForWebhook("/" + bot.Token)
	//go http.ListenAndServeTLS(":8443", "fullchain.pem", "privkey.pem", nil)

	for update := range updates {
		command := update.Message.Command()
		if command == nil || commnad == "" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi there!")
			//msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		} else {
			switch command {
				case "getpackages":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, strings.Join(parts,"\n"))
					bot.Send(msg)
				}
		}
	}
}
