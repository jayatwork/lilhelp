package cmdline

import (
	"flag"
	"fmt"
)

func Init() {
	all := flag.String("all", "foo", "Invoker help util across ALL projects")
	ccoe := flag.String("ccoe", "foo", "Invoker help util across CCOE projects")
	paas := flag.String("paas", "foo", "Invoker help util across PAAS projects")
	devops := flag.String("devops", "foo", "Invoker help util across PAAS projects")

	flag.Parse()

	fmt.Println("All helper:", *all)
	fmt.Println("CCOE helper:", *ccoe)
	fmt.Println("PAAS helper:", *paas)
	fmt.Println("DEVOPS helper:", *devops)
}
