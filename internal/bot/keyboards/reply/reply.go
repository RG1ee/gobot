package reply

import (
	tele "gopkg.in/telebot.v3"
)

func StartKeyboard() *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	sendClothButton := menu.Text("Отправить в химчистку")
	inDryCleaningClothButton := menu.Text("В химчистке")
	outDryCleaningClothButton := menu.Text("Из химчистке")
	menu.Reply(menu.Row(sendClothButton, inDryCleaningClothButton, outDryCleaningClothButton))
	return menu
}

func CancelKeyboard() *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	cancelButton := menu.Text("Отмена")
	menu.Reply(menu.Row(cancelButton))
	return menu
}
