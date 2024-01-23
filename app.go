package main

import (
	. "CleanArch/setting"
)

func RunApp() {
	server := EstablishServer()
	err := server.Start()
	if err != nil {
		return
	}
	err1 := server.Launch()
	if err1 != nil {
		return
	}
}
