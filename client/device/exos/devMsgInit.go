package main

//The state machine executes this function during the INIT
//phase.
//
//The application performs any initialization it may require to
//support the execution of the state machine.
//
//The state machine adds the fsm client to the upgrade asSet
//list.
//
//The application SHOULD populate the TBD
//the reboot events if the connector rebooted the device.
//
//A return code of OK, and the state machine transitions
//to DISCOVER.
func (dev Device) Init() {
	if !dev.simulation {
		// at startup, call the feature DB loaders to populate the device specific DB
		// Simulations all use the real devices information
		for _, f := range *dev.features {
			f.GetDBConfig()
			f.GetDBStats()
		}
	}
}
