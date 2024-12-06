package callback

import (
	"strconv"

	"github.com/RG1ee/gobot/internal/bot/keyboards/inline"
	"github.com/RG1ee/gobot/internal/repository"
	tele "gopkg.in/telebot.v3"
)

func HandlePagination(c tele.Context) error {
	db := c.Get("repository").(repository.Cloth)
	allCloth := db.GetIncoming()
	if len(allCloth) == 0 {
		c.Delete()
		return c.Send("Нет отправленных вещей")
	}

	page, _ := strconv.Atoi(c.Callback().Data)
	pageSize := 1
	paginationKeyboard := inline.GeneratePaginationKeyboard(allCloth, page, pageSize)

	photo := &tele.Photo{File: tele.File{FileID: allCloth[page].PhotoId}, Caption: allCloth[page].Name}
	return c.Edit(photo, &tele.SendOptions{ReplyMarkup: paginationKeyboard})
}

func OutgoingClothHandle(c tele.Context) error {
	db := c.Get("repository").(repository.Cloth)
	clothId, _ := strconv.Atoi(c.Callback().Data)

	cloth, err := db.GetById(clothId)
	if err != nil {
		return c.Send("Такой записи уже нет")
	}

	db.Out(cloth)

	c.Delete()
	return c.Send("Вещь " + cloth.Name + " успешное пришла!")
}
