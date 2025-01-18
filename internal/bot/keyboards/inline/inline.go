package inline

import (
	"strconv"

	"github.com/RG1ee/gobot/pkg/domain"
	tele "gopkg.in/telebot.v3"
)

func GeneratePaginationKeyboard(items []domain.Cloth, page int, pageSize int, isOutgoing bool, uniquePrevBtn string, uniqueNextBtn string) *tele.ReplyMarkup {
	keyboard := &tele.ReplyMarkup{}
	var rows []tele.Row

	start := page * pageSize
	end := start + pageSize
	if end > len(items) {
		end = len(items)
	}

	for _, item := range items[start:end] {
		if !isOutgoing {
			row := tele.Row{
				tele.Btn{
					Unique: "outCloth",
					Text:   "Вещь " + item.Name + " пришла",
					Data:   strconv.Itoa(int(item.ID)),
				},
			}
			rows = append(rows, row)
		}
	}

	var navRow tele.Row
	if page > 0 {
		navRow = append(navRow, tele.Btn{
			Unique: uniquePrevBtn,
			Text:   "Назад",
			Data:   strconv.Itoa(page - 1),
		})
	}
	if end < len(items) {
		navRow = append(navRow, tele.Btn{
			Unique: uniqueNextBtn,
			Text:   "Вперед",
			Data:   strconv.Itoa(page + 1),
		})
	}
	if len(navRow) > 0 {
		rows = append(rows, navRow)
	}

	keyboard.Inline(rows...)

	return keyboard
}

func DeleteKeyboard(idCloth int) *tele.ReplyMarkup {
	keyboard := &tele.ReplyMarkup{}
	deleteButton := tele.InlineButton{
		Text:   "Вещь пришла",
		Unique: "item_arrived",
		Data:   strconv.Itoa(idCloth),
	}
	keyboard.InlineKeyboard = [][]tele.InlineButton{
		{deleteButton},
	}

	return keyboard
}

func ReturnKeyboard(idCloth int) *tele.ReplyMarkup {
	keyboard := &tele.ReplyMarkup{}
	deleteButton := tele.InlineButton{
		Unique: "return_item",
		Text:   "Отмена",
		Data:   strconv.Itoa(idCloth),
	}
	keyboard.InlineKeyboard = [][]tele.InlineButton{
		{deleteButton},
	}

	return keyboard
}
