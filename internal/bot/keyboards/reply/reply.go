package reply

import (
	tele "gopkg.in/telebot.v3"
)

func StartKeyboard() *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	sendClothButton := menu.Text("Отправить вещь")
	inDryCleaningClothButton := menu.Text("Отправленные вещи")
	outDryCleaningClothButton := menu.Text("Пришедшие вещи")
	menu.Reply(menu.Row(sendClothButton, inDryCleaningClothButton, outDryCleaningClothButton))
	return menu
}

func CancelKeyboard() *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	cancelButton := menu.Text("Отмена")
	menu.Reply(menu.Row(cancelButton))
	return menu
}
