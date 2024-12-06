package inline

import (
	"strconv"

	"github.com/RG1ee/gobot/pkg/domain"
	tele "gopkg.in/telebot.v3"
)

func GeneratePaginationKeyboard(items []domain.Cloth, page, pageSize int) *tele.ReplyMarkup {
	keyboard := &tele.ReplyMarkup{}
	var rows []tele.Row

	start := page * pageSize
	end := start + pageSize
	if end > len(items) {
		end = len(items)
	}

	for _, item := range items[start:end] {
		row := tele.Row{
			tele.Btn{
				Unique: "outCloth",
				Text:   "Вещь " + item.Name + " пришла",
				Data:   strconv.Itoa(int(item.ID)),
			},
		}
		rows = append(rows, row)
	}

	var navRow tele.Row
	if page > 0 {
		navRow = append(navRow, tele.Btn{
			Unique: "prev_btn",
			Text:   "Назад",
			Data:   strconv.Itoa(page - 1),
		})
	}
	if end < len(items) {
		navRow = append(navRow, tele.Btn{
			Unique: "next_btn",
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
