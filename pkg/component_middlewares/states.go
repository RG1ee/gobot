package component_middlewares

import (
	userstate "github.com/RG1ee/gobot/pkg/user_state"
	tele "gopkg.in/telebot.v3"
)

type State struct {
	UserState         userstate.UserState
	CurrentMessages   []tele.Editable
	PreviousMessages  *[]tele.Editable
	DeletedIdMessages []int
}

const (
	NullState userstate.UserState = iota
	StateWaitPhoto
	StateSaveChanges
)
