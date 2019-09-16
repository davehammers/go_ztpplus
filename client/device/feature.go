package device

import (
	msg "ztp"
)

//This interface defines the functions that every feature must implement
type Feature interface {
	GetConnect(*msg.Connect) error
	GetConfig(*msg.Configuration) error
	SetConfig(*msg.ConfigurationResponse) error
	GetStats(*msg.Stats) error
	SetStats(*msg.StatsResponse) error
}
