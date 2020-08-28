package utils

import (
//    "log"
    "fmt"
    "strings"
//    "strconv"
    tb "gopkg.in/tucnak/telebot.v2"
//    "github.com/philippgille/gokv"
)

func AskTitle(bot *tb.Bot, msg *tb.Message) {
    text := "Введите заголовок для новой заявки"
    bot.Send(msg.Sender, text, tb.ParseMode("Markdown"))
}

func AskBody(bot *tb.Bot, msg *tb.Message) {
    text := "Введите содержание новой заявки"
    bot.Send(msg.Sender, text, tb.ParseMode("Markdown"))
}


func ParseMessage(msg string) (string, string) {
    arr := strings.SplitN(msg, "\n", 2)
    title, body := arr[0], arr[1]
    return title, body
}

func FormatMessage(title string, body string) string {
    return fmt.Sprintf("*%s*\n%s", title, body)
}

func FixMessage(text string) string {
    return FormatMessage(ParseMessage(text))
}
