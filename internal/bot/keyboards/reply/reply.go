package reply

import (
	tele "gopkg.in/telebot.v3"
)

func StartKeyboard() *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	sendClothButton := menu.Text("Отправить вещь")
	inDryCleaningClothButton := menu.Text("Отправленные вещи")
	outDryCleaningClothLastSevenDaysButton := menu.Text("Пришедшие вещи за последние 7 дней")
	outDryCleaningClothAllTimeButton := menu.Text("Пришедшие вещи за все время")
	menu.Reply(menu.Row(sendClothButton), menu.Row(inDryCleaningClothButton), menu.Row(outDryCleaningClothLastSevenDaysButton), menu.Row(outDryCleaningClothAllTimeButton))
	return menu
}

func CancelKeyboard() *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	cancelButton := menu.Text("Отменить и вернуться в главное меню")
	menu.Reply(menu.Row(cancelButton))
	return menu
}

func SaveChangesKeyboard() *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	saveChangesButton := menu.Text("Сохранить изменения")
	cancelButton := menu.Text("Отменить и вернуться в главное меню")
	menu.Reply(menu.Row(saveChangesButton, cancelButton))
	return menu
}
