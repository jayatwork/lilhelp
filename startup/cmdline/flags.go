package cmdline

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
)

type selection struct {
	address string
}

func Init() []selection {
	fmt.Println("\n\nUsage: \nlilhelp utility usage guide")
	all := flag.String("all", "-all", "Invoker help util across ALL projects")
	ccoe := flag.String("ccoe", "-ccoe", "Invoker help util across CCOE projects")
	paas := flag.String("paas", "-paas", "Invoker help util across PAAS projects")
	devops := flag.String("devops", "-devops", "Invoker help util across PAAS projects")

	flag.Parse()

	input := []selection{""}
	fmt.Println("All helper:", *all)
	fmt.Println("CCOE helper:", *ccoe)
	fmt.Println("PAAS helper:", *paas)
	fmt.Println("DEVOPS helper:", *devops)

	// TODO follow on the array of selection type struct
	basenameOpts = []opt{
		opt{
			shortnm: 'a',
			longnm:  "multiple",
			needArg: false,
			help:    "Usage for a",
		},
		opt{
			shortnm: 'b',
			longnm:  "b-option",
			needArg: false,
			help:    "Usage for b",
		},
	}

	switch flag.Getter.Get() {
	case "all":
		return []selection{"url": "somestring"}
	case "ccoe":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "paas":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported helper domain selected")
	}
	if err != nil {
		log.Fatal(err)
	}
}
