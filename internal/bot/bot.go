package bot

import (
	"os"
	"time"

	"github.com/RG1ee/gobot/internal/bot/handlers/callback"
	"github.com/RG1ee/gobot/internal/bot/handlers/message"
	"github.com/RG1ee/gobot/internal/repository"
	"github.com/RG1ee/gobot/internal/repository/repository_backend"
	"github.com/RG1ee/gobot/pkg/component_middlewares"
	"github.com/RG1ee/gobot/pkg/middleware"
	"github.com/avi-gecko/fsm/pkg/fsm"
	tele "gopkg.in/telebot.v3"
)

type TelegramBot struct {
	bot *tele.Bot
	fsm fsm.FSM[component_middlewares.State]
	db  repository.Cloth
}

func NewTelegramBot() (*TelegramBot, error) {
	bot, err := createBot()
	if err != nil {
		panic(err)
	}
	finiteStateMachine, err := fsm.Create[component_middlewares.State](fsm.RAM{})
	if err != nil {
		panic(err)
	}
	// NOTE: for docker
	// db := repository_backend.Sqlite{DB_name: "/volume/db"}
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
		ParseMode: tele.ModeHTML,
	})
	if err != nil {
		return nil, err
	}
	return bot, nil
}

func (tb *TelegramBot) RegisterHandler() {
	// tb.bot.Use(component_middlewares.CleanupMessages())
	tb.bot.Use(component_middlewares.FsmMiddleware(tb.fsm))
	tb.bot.Use(middleware.Repository(tb.db))
	tb.bot.Handle("/start", component_middlewares.SaveLastMessage(message.StartMessageHandler), component_middlewares.CleanupMessages())
	tb.bot.Handle("Отменить и вернуться в главное меню", component_middlewares.SaveLastMessage(message.CancelHandler), component_middlewares.CleanupMessages())
	tb.bot.Handle("Отправить вещь", component_middlewares.SaveLastMessage(message.WriteNewClothMessageHandler), component_middlewares.CleanupMessages())
	tb.bot.Handle("Отправленные вещи", component_middlewares.SaveLastMessage(message.GetListIncomingClothMessageHandler), component_middlewares.CleanupMessages())
	tb.bot.Handle("Пришедшие вещи", component_middlewares.SaveLastMessage(message.GetListOutClothMessageHandler), component_middlewares.CleanupMessages())
	tb.bot.Handle("Сохранить изменения", component_middlewares.StateGate(component_middlewares.SaveLastMessage(message.SaveChangesHandler), component_middlewares.StateSaveChanges), component_middlewares.CleanupMessages())
	tb.bot.Handle(&tele.Btn{Unique: "item_arrived"}, component_middlewares.StateGate(callback.IncomingClothHandle, component_middlewares.StateSaveChanges))
	tb.bot.Handle(&tele.Btn{Unique: "return_item"}, component_middlewares.StateGate(callback.CancelIncomingClothHandle, component_middlewares.StateSaveChanges))
	tb.bot.Handle(tele.OnPhoto, component_middlewares.StateGate(component_middlewares.SaveLastMessage(message.GetPhotoClothMessageHandler), component_middlewares.StateWaitPhoto), component_middlewares.CleanupMessages())
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
