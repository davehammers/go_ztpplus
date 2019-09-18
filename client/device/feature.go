package device

import (
	msg "ztp"
)

//This interface defines the functions that every feature must implement
type Feature interface {
	GetDBConfig() error                         // feature specific routine to get config data
	GetDBStats() error                          // feature specific routine to get statistics data
	GetConnect(*msg.Connect) error              // feature specific routine to fill in the Connect message
	GetConfig(*msg.Configuration) error         // feature specific routine to fill in the Config message
	SetConfig(*msg.ConfigurationResponse) error // feature specific routine to process the Config response
	GetStats(*msg.Stats) error                  // feature specific routine to fill in Stats message
	SetStats(*msg.StatsResponse) error          // feature specific routine to process in Stats message response
}
