package main

import (
	"log"
	"strconv"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {
	pref := tele.Settings{
		Token:  "6392888294:AAHw6D0LESoKHcC3xoseU4hwvbSLNrb-p1I",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	//b.Handle("/hello", func(c tele.Context) error {
	//	return c.Send("Hello!")
	//})
	//b.Handle(tele.OnText, func(c tele.Context) error {
	//	// All the text messages that weren't
	//	// captured by existing handlers.
	//
	//	var (
	//		user = c.Sender()
	//		text = c.Text()
	//	)
	//
	//	// Use full-fledged bot's functions
	//	// only if you need a result:
	//	msg, err := b.Send(user, text)
	//	if err != nil {
	//		return err
	//	}
	//	fmt.Println(msg)
	//
	//	// Instead, prefer a context short-hand:
	//	return c.Send(text)
	//})
	//
	//b.Handle(tele.OnChannelPost, func(c tele.Context) error {
	//	fmt.Println("OnChannelPost")
	//	// Channel posts only.
	//	msg := c.Message()
	//	fmt.Println(msg)
	//	return nil
	//})
	//
	//b.Handle(tele.OnPhoto, func(c tele.Context) error {
	//	// Photos only.
	//	photo := c.Message().Photo
	//	fmt.Println(photo)
	//	return nil
	//})
	//
	//b.Handle(tele.OnQuery, func(c tele.Context) error {
	//	// Incoming inline queries.
	//	qr := tele.QueryResponse{}
	//	//fmt.Println(tele.OnQuery)
	//	fmt.Println("fuck pig", c.Text())
	//
	//	return c.Answer(&qr)
	//
	//})
	//b.Handle("/tags", func(c tele.Context) error {
	//	tags := c.Args() // list of arguments splitted by a space
	//	for _, tag := range tags {
	//		fmt.Println(tag)
	//		// iterate through passed arguments
	//	}
	//	return nil
	//})

	//var (
	//	// Universal markup builders.
	//	menu     = &tele.ReplyMarkup{ResizeKeyboard: true}
	//	selector = &tele.ReplyMarkup{}
	//
	//	// Reply buttons.
	//	btnHelp     = menu.Text("ℹ Help")
	//	btnSettings = menu.Text("⚙ Settings")
	//
	//	// Inline buttons.
	//	//
	//	// Pressing it will cause the client to
	//	// send the bot a callback.
	//	//
	//	// Make sure Unique stays unique as per button kind
	//	// since it's required for callback routing to work.
	//	//
	//	btnPrev = selector.Data("⬅", "prev")
	//	btnNext = selector.Data("➡", "next")
	//)
	//
	//menu.Reply(
	//	menu.Row(btnHelp),
	//	menu.Row(btnSettings),
	//)
	//selector.Inline(
	//	selector.Row(btnPrev, btnNext),
	//)
	//
	//b.Handle("/ok", func(c tele.Context) error {
	//	return c.Send("ok!", menu)
	//})
	//
	//// On reply button pressed (message)
	//b.Handle(&btnHelp, func(c tele.Context) error {
	//	return c.Edit("Here is some help: ...")
	//})
	//
	//// On inline button pressed (callback)
	//b.Handle(&btnPrev, func(c tele.Context) error {
	//	return c.Respond()
	//})
	//r := b.NewMarkup()
	//
	//// Reply buttons:
	//r.Text("Hello!")
	//r.Contact("Send phone number")
	//r.Location("Send location")
	//r.Poll("adfasdf", (tele.PollQuiz))
	//
	//// Inline buttons:
	//r.Data("Show help", "help") // data is optional
	////r.Data("Delete item", "delete", item.ID)
	//r.URL("Visit", "https://google.com")
	//query := "123"
	//r.Query("Search", query)
	//r.QueryChat("Share", query)

	//r.Login("Login", &tele.Login{...})
	b.Handle(tele.OnQuery, func(c tele.Context) error {
		urls := []string{
			"https://pbs.twimg.com/media/F45ceLXbIAA6Wpz?format=jpg&name=medium",
			"https://pbs.twimg.com/media/F42K0tfakAAHnSD?format=jpg&name=medium",
		}

		results := make(tele.Results, len(urls)) // []tele.Result
		for i, url := range urls {
			result := &tele.PhotoResult{
				URL:      url,
				ThumbURL: url, // required for photos
			}

			results[i] = result
			// needed to set a unique string ID for each result
			results[i].SetResultID(strconv.Itoa(i))
		}

		return c.Answer(&tele.QueryResponse{
			Results:   results,
			CacheTime: 60, // a minute
		})
	})
	b.Start()
}
