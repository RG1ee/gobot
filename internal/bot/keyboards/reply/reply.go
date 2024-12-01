package reply

import (
	tele "gopkg.in/telebot.v3"
)

func StartKeyboard() *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	sendClothButton := menu.Text("Отправить в химчистку")
	getListClothButton := menu.Text("Список вещей")
	menu.Reply(menu.Row(sendClothButton, getListClothButton))
	return menu
}
