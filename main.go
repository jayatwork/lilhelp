package main

import (
	"flag"
	"fmt"
	b "lilhelp/startup/browser"
)

func main() {
	//initialize and prep the commandline flags
	all := flag.String("all", "foo", "Invoker help util across ALL projects")
	ccoe := flag.String("ccoe", "foo", "Invoker help util across CCOE projects")
	paas := flag.String("paas", "foo", "Invoker help util across PAAS projects")
	devops := flag.String("devops", "foo", "Invoker help util across PAAS projects")

	flag.Parse()

	fmt.Println("All helper:", *all)
	fmt.Println("CCOE helper:", *ccoe)
	fmt.Println("PAAS helper:", *paas)
	fmt.Println("DEVOPS helper:", *devops)

	b.Openbrowser("https://www16.v1host.com/Delta_Air_Lines/TeamRoom.mvc/Show/3036441") //open browser with native cmd

}
