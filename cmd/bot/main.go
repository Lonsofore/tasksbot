package main

import (
    "os"
    "fmt"
    "log"
    "time"
    "strconv"
    handlers "tasksbot/pkg/handlers"
    tb "gopkg.in/tucnak/telebot.v2"
    gokv "github.com/philippgille/gokv"
    gomap "github.com/philippgille/gokv/gomap"
    redis "github.com/philippgille/gokv/redis"
)

func main() {
    // get envs
    token := os.Getenv("BOT_TOKEN")
    channelID, err := strconv.Atoi(os.Getenv("BOT_CHANNEL_ID"))
    if err != nil {
        log.Fatal(err)
        return
    }

    // init kv storage
    store, closable := getStore()
    if closable {
        defer store.Close()
    }

    // init telebot
    channel := tb.ChatID(channelID)
    bot, err := tb.NewBot(tb.Settings{
        Token:  token,
        Poller: &tb.LongPoller{Timeout: 10*time.Second},
    })
    if err != nil {
        log.Fatal(err)
        return
    }

    // message handlers
    bot.Handle("/start", func(msg *tb.Message) {
        handlers.Start(bot, msg)
    })
    bot.Handle("/create", func(msg *tb.Message) {
        handlers.Create(bot, msg, store)
    })
    bot.Handle(tb.OnText, func(msg *tb.Message) {
        handlers.OnText(bot, msg, store)
    })

    // callback handlers
    bot.Handle(tb.OnCallback, func(cb *tb.Callback) {
        switch cb.Data {
            case "\fconfirm":
                handlers.Confirm(bot, cb, channel)
            case "\fcancel":
                handlers.Cancel(bot, cb)
            case "\faccept":
                handlers.Accept(bot, cb)
        }
    })

    log.Println("bot started")
    bot.Start()
}

func getStore() (gokv.Store, bool) {
    if os.Getenv("REDIS_ADDRESS") != "" {
        db := 0
        if os.Getenv("REDIS_DB") != "" {
            db, _ = strconv.Atoi(os.Getenv("REDIS_DB"))
        }
        options := redis.Options{
            Address: os.Getenv("REDIS_ADDRESS"),
            Password: os.Getenv("REDIS_PASSWORD"),
            DB: db}
        store, _ := redis.NewClient(options)

        passwordInfo := "password is empty"
        if os.Getenv("REDIS_PASSWORD") != "" {
            passwordInfo = "password is not empty"
        }
        log.Println(fmt.Sprintf("chosed redis as storage. address: %s, db: %d, %s", options.Address, db, passwordInfo))

        return store, true
    } else {
        options := gomap.DefaultOptions
        store := gomap.NewStore(options)
        log.Println("chosed gomap as storage")
        return store, false
    }
}
