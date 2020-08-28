package handlers

import (
    "log"
    "fmt"
    "strings"
    "strconv"
    tb "gopkg.in/tucnak/telebot.v2"
    "github.com/philippgille/gokv"
)


func Start(bot *tb.Bot, msg *tb.Message) {
    // /start command

    text := `Привет! Это бот для создания заявок в отдел девопс

Чтобы создать заявку - используйте команду /create
Заголовок заявки можно передать сразу в аргументе этой команды. Пример: /create название заявки`
    bot.Send(msg.Sender, text, tb.ParseMode("Markdown"))

    key := strconv.Itoa(msg.Sender.ID)
    log.Println(fmt.Sprintf("%s: sent /start message", key))
}

func Create(bot *tb.Bot, msg *tb.Message, store gokv.Store) {
    // /create command
    if msg.Payload != "" {
        key := strconv.Itoa(msg.Sender.ID)
        err := store.Set(key, msg.Payload)
        if err != nil {
            panic(err)
        }
        log.Println(fmt.Sprintf("key: %s", key))
        askBody(bot, msg)
    } else {
        askTitle(bot, msg)
    }

    key := strconv.Itoa(msg.Sender.ID)
    log.Println(fmt.Sprintf("%s: sent /create message", key))
}

func OnText(bot *tb.Bot, msg *tb.Message, store gokv.Store) {
    // on any text
    selector := &tb.ReplyMarkup{}
    confirmBtn := selector.Data("Подтвердить", "confirm")
    cancelBtn := selector.Data("Отмена", "cancel")
    selector.Inline(
        selector.Row(confirmBtn, cancelBtn),
    )

    title := new(string)
    key := strconv.Itoa(msg.Sender.ID)
    found, _ := store.Get(key, title)
    // if stored something - we've got the body of the task
    // if not - we've got the title
    if !found {
        store.Set(key, strings.ReplaceAll(msg.Text, "\n", " ") )
        log.Println(fmt.Sprintf("%s: title not found, saved new", key))
        askBody(bot, msg)
    } else {
        body := msg.Text
        text := fmt.Sprintf("*%s*\n%s", *title, body)
        bot.Send(msg.Sender, text, tb.ParseMode("Markdown"), selector)
        log.Println(fmt.Sprintf("%s: title found, confirmation sent", key))

        store.Delete(key)
    }
}



func askTitle(bot *tb.Bot, msg *tb.Message) {
    text := "Введите заголовок для новой заявки"
    bot.Send(msg.Sender, text, tb.ParseMode("Markdown"))
}

func askBody(bot *tb.Bot, msg *tb.Message) {
    text := "Введите содержание новой заявки"
    bot.Send(msg.Sender, text, tb.ParseMode("Markdown"))
}

