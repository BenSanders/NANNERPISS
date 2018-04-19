package main

import (
	"github.com/NANNERPISS/NANNERPISS"
	"github.com/NANNERPISS/NANNERPISS/config"
	"github.com/NANNERPISS/NANNERPISS/db"

	"github.com/ChimeraCoder/anaconda"
	"gopkg.in/telegram-bot-api.v4"
)

func main() {
	b := bot.New()

	var err error
	b.Config, err = config.Load("config.json")
	if err != nil {
		panic(err)
	}
	b.DB, err = db.New(b.Config.DB.Driver, b.Config.DB.Source)
	if err != nil {
		panic(err)
	}
	b.TG, err = tgbotapi.NewBotAPI(b.Config.TG.Token)
	if err != nil {
		panic(err)
	}
	b.TW = anaconda.NewTwitterApiWithCredentials(
		b.Config.TW.AccessToken, b.Config.TW.AccessSecret,
		b.Config.TW.ConsumerKey, b.Config.TW.ConsumerSecret,
	)

	b.Run()
}
