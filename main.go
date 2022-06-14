package main

import (
	b "lilhelp/startup/browser"
	t "lilhelp/startup/timer"
)

func main() {
	//initialize and prep the commandline flags
	//f.Init()
	//Iterate through brower collections depending on flag selected
	t.Start()
	b.Openbrowser("https://www16.v1host.com/Delta_Air_Lines/TeamRoom.mvc/Show/3036441") //open browser with native cmd

}
