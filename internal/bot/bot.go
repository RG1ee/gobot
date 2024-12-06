package bot

import (
	"os"
	"time"

	"github.com/RG1ee/gobot/internal/bot/handlers/callback"
	"github.com/RG1ee/gobot/internal/bot/handlers/message"
	stateconst "github.com/RG1ee/gobot/internal/bot/state_const"
	"github.com/RG1ee/gobot/internal/repository"
	"github.com/RG1ee/gobot/internal/repository/repository_backend"
	"github.com/RG1ee/gobot/pkg/handler_decorator"
	"github.com/RG1ee/gobot/pkg/middleware"
	"github.com/avi-gecko/fsm/pkg/fsm"
	tele "gopkg.in/telebot.v3"
)

type TelegramBot struct {
	bot *tele.Bot
	fsm fsm.FSM
	db  repository.Cloth
}

func NewTelegramBot() (*TelegramBot, error) {
	bot, err := createBot()
	if err != nil {
		panic(err)
	}
	finiteStateMachine, err := fsm.Create(fsm.RAM{})
	if err != nil {
		panic(err)
	}
	db := repository_backend.Sqlite{DB_name: "db"}
	db.Init()
	return &TelegramBot{
		bot: bot,
		fsm: finiteStateMachine,
		db:  &db,
	}, nil
}

func createBot() (*tele.Bot, error) {
	bot, err := tele.NewBot(tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		return nil, err
	}
	return bot, nil
}

func (tb *TelegramBot) RegisterHandler() {
	tb.bot.Use(middleware.FsmMiddleware(tb.fsm))
	tb.bot.Use(middleware.Repository(tb.db))
	tb.bot.Handle("/start", message.StartMessageHandler)
	tb.bot.Handle("Отмена", message.CancelHandler)
	tb.bot.Handle("Отправить в химчистку", message.WriteNewClothMessageHandler)
	tb.bot.Handle("В химчистке", message.GetListIncomingClothMessageHandler)
	tb.bot.Handle(&tele.Btn{Unique: "next_btn"}, callback.HandlePagination)
	tb.bot.Handle(&tele.Btn{Unique: "prev_btn"}, callback.HandlePagination)
	tb.bot.Handle(&tele.Btn{Unique: "outCloth"}, callback.OutgoingClothHandle)
	tb.bot.Handle(tele.OnPhoto, handlerdecorator.DecoratorHandle(message.GetPhotoClothMessageHandler, stateconst.StateWaitPhoto))
}

func Start() error {
	tb, err := NewTelegramBot()
	if err != nil {
		return err
	}

	tb.RegisterHandler()

	tb.bot.Start()
	return nil
}
