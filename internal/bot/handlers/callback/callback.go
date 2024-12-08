package callback

import (
	"strconv"

	"github.com/RG1ee/gobot/internal/bot/keyboards/inline"
	"github.com/RG1ee/gobot/internal/repository"
	"github.com/RG1ee/gobot/pkg/domain"
	tele "gopkg.in/telebot.v3"
)

func handlePagination(c tele.Context, allCloth []domain.Cloth, emptyMessage string, isOutgoing bool, uniquePrevBtn string, uniqueNextBtn string) error {
	if len(allCloth) == 0 {
		c.Delete()
		return c.Send(emptyMessage)
	}

	page, _ := strconv.Atoi(c.Callback().Data)
	if page < 0 || page >= len(allCloth) {
		return c.Send("Неверная страница.")
	}

	pageSize := 1
	paginationKeyboard := inline.GeneratePaginationKeyboard(allCloth, page, pageSize, isOutgoing, uniquePrevBtn, uniqueNextBtn)

	photo := &tele.Photo{
		File:    tele.File{FileID: allCloth[page].PhotoId},
		Caption: allCloth[page].Name,
	}
	return c.Edit(photo, &tele.SendOptions{ReplyMarkup: paginationKeyboard})
}

func HandleIncomingPagination(c tele.Context) error {
	db := c.Get("repository").(repository.Cloth)
	allCloth := db.GetIncoming()
	return handlePagination(c, allCloth, "Нет входящих вещей", false, "incoming_prev_btn", "incoming_next_btn")
}

func HandleOutgoingPagination(c tele.Context) error {
	db := c.Get("repository").(repository.Cloth)
	allCloth := db.GetOutgoing()
	return handlePagination(c, allCloth, "Нет отправленных вещей", true, "prev_btn", "next_btn")
}

func IncomingClothHandle(c tele.Context) error {
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
