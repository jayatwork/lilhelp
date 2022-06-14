package cmdline

import (
	"flag"
	"fmt"
)

type selection struct {
	address, shortnm, longnm, help string
	needArg                        bool
}

func Init() string {
	fmt.Println("\n\nUsage: \nlilhelp utility usage guide")
	all := flag.String("all", "-all", "Invoker help util across ALL projects")
//	ccoe := flag.String("ccoe", "-ccoe", "Invoker help util across CCOE projects")
	// paas := flag.String("paas", "-paas", "Invoker help util across PAAS projects")
	// devops := flag.String("devops", "-devops", "Invoker help util across PAAS projects")
	// sprinter := flag.String("sprinter", "-sprinter", "Invoker help util for sprint board across ")

	flag.Parse()

	fmt.Println("All helper:", *all)
	// fmt.Println("CCOE helper:", *ccoe)
	// fmt.Println("PAAS helper:", *paas)
	// fmt.Println("DEVOPS helper:", *devops)
	// fmt.Println("VersionOne helper:", *sprinter)

	// TODO follow on the array of selection type struct
	basenameOpts := []selection{
		selection{
			address: "https://www16.v1host.com/Delta_Air_Lines/TeamRoom.mvc/Show/3036441",
			shortnm: "s",
			longnm:  *sprinter,
			needArg: true,
			help:    "Usage for s",
		},
		selection{
			address: "https://www16.v1host.com/Delta_Air_Lines/TeamRoom.mvc/Show/3036441",
			shortnm: "a",
			longnm:  "ccoe",
			needArg: false,
			help:    "Usage for a",
		},
	}

	choice := "all"
	for _, options := range basenameOpts {
		if options.shortnm == "s" {
			fmt.Printf("%#v\n", basenameOpts)
			choice = options.address
		}
	}
	return choice

}
