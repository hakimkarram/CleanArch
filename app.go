package main

import (
	. "CleanArch/settings"
)

func RunApp() {
	serv := EstablishServer()
	err := serv.Start()
	if err != nil {
		return
	}
	err1 := serv.Launch()

	if err1 != nil {
		return
	}
}
