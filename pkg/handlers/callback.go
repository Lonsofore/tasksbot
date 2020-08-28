package handlers

import (
    "log"
    "fmt"
    "strconv"
    tb "gopkg.in/tucnak/telebot.v2"
    utils "tasksbot/pkg/utils"
)


func Confirm(bot *tb.Bot, cb *tb.Callback, channel tb.ChatID) {
    // confirm task and send to the channel

    selector := &tb.ReplyMarkup{}
    acceptBtn := selector.Data("Взять в работу", "accept")
    selector.Inline(
        selector.Row(acceptBtn),
    )


    // accept the task in bot chat
    text := utils.FixMessage(cb.Message.Text)
    bot.Edit(cb.Message, text, tb.ParseMode("Markdown"))
    bot.Send(cb.Sender, "Задача создана!", tb.ParseMode("Markdown"))

    key := strconv.Itoa(cb.Sender.ID)
    log.Println(fmt.Sprintf("%s: task has been confirmed", key))

    // send the task to channel
    taskText := fmt.Sprintf("Задача от @%s\n%s", cb.Sender.Username, text)
    bot.Send(channel, taskText, tb.ParseMode("Markdown"), selector)

    log.Println(fmt.Sprintf("%s: task has been sent to the channel", key))

    // respond to callback
    bot.Respond(cb, &tb.CallbackResponse{})
}

func Cancel(bot *tb.Bot, cb *tb.Callback) {
    // cancel task to prevent sending

    text := utils.FixMessage(cb.Message.Text) // because bot lose format
    bot.Edit(cb.Message, text, tb.ParseMode("Markdown"))
    bot.Send(cb.Sender, "Создание задачи отменено", tb.ParseMode("Markdown"))
    bot.Respond(cb, &tb.CallbackResponse{})

    key := strconv.Itoa(cb.Sender.ID)
    log.Println(fmt.Sprintf("%s: task has been canceled", key))
}

func Accept(bot *tb.Bot, cb *tb.Callback) {
    // accept task in the channel

    text := fmt.Sprintf("%s\n\nЗадача взята в работу @%s",
    utils.FixMessage(cb.Message.Text), cb.Sender.Username)

    bot.Edit(cb.Message, text, tb.ParseMode("Markdown"))
    bot.Respond(cb, &tb.CallbackResponse{})

    key := strconv.Itoa(cb.Sender.ID)
    log.Println(fmt.Sprintf("%s: task has been accepted", key))
}
