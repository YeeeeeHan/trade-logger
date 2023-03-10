package handlers

import (
	"NFTracker/cmd/message"
	"NFTracker/pkg/opensea"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const PopularCollectionNumOwners = 400

func Introduction(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, message.WelcomeMessage)
	msg.ParseMode = "Markdown"
	msg.DisableWebPagePreview = true
	if _, e := bot.Send(msg); e != nil {
		log.Printf("Error sending message to telegram.\nMessage: %v\nError: %v", msg, e)
	}
	return
}

func PriceCheck(bot *tgbotapi.BotAPI, chat *tgbotapi.Chat, t, userName string) {
	userLink := fmt.Sprintf("https://t.me/%s", userName)
	msg := fmt.Sprintf("[From %s](%s)\n", userName, userLink) + t[5:]
	message.SendMessage(bot, -1001744918879, msg)
	message.SendMessage(bot, chat.ID, msg)
	return

	//// If empty slugQuery, send "No slugQuery detected"
	//if slugQuery == "" {
	//	msg := "No slugQuery detected."
	//	message.SendMessage(bot, chat.ID, msg)
	//	return
	//}
	//
	//// else query web to get osResponse
	//osResponse, err := opensea.QueryAPI(slugQuery)
	//
	//// If slug is invalid
	//if err == custError.InvalidSlugErr {
	//	// Find the closest match from list of popular collections from DB
	//	match, err := closestMatchHelper(pgdb, slugQuery)
	//	if err != nil {
	//		log.Printf(fmt.Sprintf("[closestMatchHelper] Err: %s", err))
	//		return
	//	}
	//
	//	// Send notification of 404 and suggest the closest match
	//	msg := fmt.Sprintf("⚠️ \"%s\" does not exist, please double-check `<slug>`", slugQuery)
	//	message.SendInlineSlugMissMessage(bot, chat.ID, msg, slugQuery, match)
	//	return
	//}
	//if err != nil {
	//	// Send notification of 404 but suggest the closest match
	//	msg := fmt.Sprintf("⚠️  Internal error, please give us some time to solve it.")
	//	message.SendMessage(bot, chat.ID, msg)
	//}
	//
	//// If collection is popular (> certain number of owners), return osResponse
	//if osResponse.Collection.Stats.NumOwners > PopularCollectionNumOwners {
	//	// Send price check message
	//	msg := message.PriceCheckMessage(osResponse.Collection.Name, opensea.CreateOpenseaUrlFromSlug(slugQuery), osResponse)
	//	message.SendMessage(bot, chat.ID, msg)
	//
	//	// Update DB with popular collection
	//	popularCollectionHelper(pgdb, osResponse, slugQuery)
	//
	//	return
	//
	//	// If collection is not popular, return osResponse AND suggest the closest match
	//} else {
	//	// Return query
	//	msg := message.PriceCheckMessage(osResponse.Collection.Name, opensea.CreateOpenseaUrlFromSlug(slugQuery), osResponse)
	//
	//	// But suggest the closest match from list of popular NFTs from database
	//	match, err := closestMatchHelper(pgdb, slugQuery)
	//	if err != nil {
	//		log.Printf(fmt.Sprintf("[closestMatchHelper] Err: %s", err))
	//		return
	//	}
	//
	//	message.SendInlineSlugMissMessage(bot, chat.ID, msg, slugQuery, match)
	//	return
	//}
}

// EditMessage updates the message(ID: messageID) when user clicks on the match suggestion
func EditMessage(bot *tgbotapi.BotAPI, messageID int, chatID int64, slugQuery string) {
	// Query web
	osResponse, err := opensea.QueryAPI(slugQuery)
	if err != nil {
		msg := "⚠️  Collection does not exist!"
		message.SendMessage(bot, chatID, msg)
		log.Printf("[opensea.QueryAPI] %v", err)
		return
	}

	// Send price check message
	msg := message.PriceCheckMessage(osResponse.Collection.Name, opensea.CreateOpenseaUrlFromSlug(slugQuery), osResponse)
	message.SendEditMessage(bot, chatID, messageID, msg)
}
