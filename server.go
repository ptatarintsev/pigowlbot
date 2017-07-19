package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"pigowlbot/token"
	"strconv"
	"strings"
	"time"

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

type PackStatResponse struct {
	Timestamp int64  `json:"timestamp"`
	ID int  `json:"id"`
}

type GetPacksResponse struct {
	Packs  []PackResponse `json:"packs"`
}

type GetPacksStatResponse struct {
	PacksStat  []PackStatResponse
}

func MainHandler(resp http.ResponseWriter, _ *http.Request) {
	resp.Write([]byte("Hi there! I'm PigowlTestBot!"))
}

func getPacksResponse() *GetPacksResponse {
	url := "http://pigowl.com:8080/getPacks"

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	response := new(GetPacksResponse) 
	json.NewDecoder(res.Body).Decode(response)
	return response
}

func getPacksStatResponse() *GetPacksStatResponse {
	url := "http://pigowl.com:8090/getPacksStat"

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	response := new(GetPacksStatResponse) 
	json.NewDecoder(res.Body).Decode(response)
	return response
}

func getPackages() string {
	response := getPacksResponse()

	var parts []string
	for _, pack := range response.Packs {
		parts = append(parts, pack.Pack.Name)
	}
	return strings.Join(parts,"\n")
}

func getDownloads() string {
	packsResponse := getPacksResponse()
	packsStatResponse := getPacksStatResponse()

	packsMap := make(map[int]string)
	for _, pack := range packsResponse.Packs {
		packsMap[pack.Pack.ID] = pack.Pack.Name
	}

	var parts []string
	weekAgo := time.Now().Add(-7*24*time.Hour).Truncate(24 * time.Hour).Unix()
	parts = append(parts, strconv.FormatInt(weekAgo, 10))

	var packStats []PackStatResponse
	for _, packStat := range packsStatResponse.PacksStat {
		parts = append(parts, strconv.FormatInt(packStat.Timestamp, 10))
		if packStat.Timestamp >= weekAgo {
			packStats = append(packStats, packStat)
		}
	}
	//var parts []string
	for _, packStat := range packStats {
		parts = append(parts, packsMap[packStat.ID])
	}
	return strings.Join(parts,"\n")
}

func main() {
	bot, err := tgbotapi.NewBotAPI(token.BotToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	updates := bot.ListenForWebhook("/" + token.BotToken)

	http.HandleFunc("/", MainHandler)
	go http.ListenAndServe(":"+os.Getenv("PORT"), nil)

	for update := range updates {
		command := update.Message.Command()
		if command == "" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi there!")
			bot.Send(msg)
		} else {
			switch command {
				case "getpackages":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, getPackages())
					bot.Send(msg)
				case "getdownloads":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, getDownloads())
					bot.Send(msg)
				}
		}
	}
}
