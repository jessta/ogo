package ogo

import (
	"github.com/jonstout/ogo/openflow/ofp10"
)

var Applications map[string]Application
var messageChans map[uint8][]chan ofp10.OfpMsg

type Application interface {
	InitApplication(args map[string]string)
	Name() string
	Receive()
}

func SubscribeTo(msg uint8) chan ofp10.OfpMsg {
	ch := make(chan ofp10.OfpMsg)
	messageChans[msg] = append(messageChans[msg], ch)
	return ch
}
