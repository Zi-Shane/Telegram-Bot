package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// Bot TOKEN
	botToken := os.Getenv("TELEGRAM_APITOKEN")

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Set Webhook
	production, _ := strconv.ParseBool(os.Getenv("PRODUCTION"))
	var cert interface{}
	var webhookUrl string
	if production {
		webhookUrl = os.Getenv("DOMAINNAME") + ":" + os.Getenv("PORT") + "/" + bot.Token
		cert = "./secret/" + os.Getenv("DOMAINNAME") + "/tls.crt"
		_, err = bot.SetWebhook(tgbotapi.NewWebhookWithCert(webhookUrl, cert))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		webhookUrl = os.Getenv("DOMAINNAME") + "/" + bot.Token
		_, err = bot.SetWebhook(tgbotapi.NewWebhook(webhookUrl))
		if err != nil {
			log.Fatal(err)
		}
	}

	// Test Webhook
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	// Setup Webhook server
	updates := bot.ListenForWebhook("/" + bot.Token)
	if production {
		go http.ListenAndServeTLS(
			":"+os.Getenv("PORT"),
			"./secret/"+os.Getenv("DOMAINNAME")+"/tls.crt",
			"./secret/"+os.Getenv("DOMAINNAME")+"/tls.key",
			nil)
	} else {
		go http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	}

	// Get Updates from telegram
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Text == "Hi" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi!, "+update.Message.Chat.FirstName+update.Message.Chat.LastName)

			msg.ReplyToMessageID = update.Message.MessageID
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}
	}
}
