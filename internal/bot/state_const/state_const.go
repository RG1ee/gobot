package stateconst

import (
	userstate "github.com/RG1ee/gobot/pkg/user_state"
	tele "gopkg.in/telebot.v3"
)

type State struct {
	UserState        userstate.UserState
	CurrentMessages  []tele.Editable
	PreviousMessages *[]tele.Editable
}

const (
	NullState userstate.UserState = iota
	StateWaitPhoto
)
