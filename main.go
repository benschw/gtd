package main

import (
	"flag"
	"fmt"
	"log"
	"log/syslog"

	"github.com/benschw/ghgtd/gtd"
)

func main() {

	useSyslog := flag.Bool("syslog", false, "log to syslog")
	flag.Parse()

	if *useSyslog {
		logwriter, err := syslog.New(syslog.LOG_NOTICE, "todo")
		if err == nil {
			log.SetOutput(logwriter)
		}
	}

	// pull desired command/operation from args
	if flag.NArg() == 0 {
		flag.Usage()
		log.Fatal("Command argument required")
	}

	args := make([]string, 0)

	for i := 0; i < flag.NArg(); i++ {
		args = append(args, flag.Arg(i))
	}

	d := &gtd.Dispatcher{}
	out, err := d.Dispatch(&gtd.Meta{Context: "@global"}, args)
	if err != nil {
		log.Println(err)
	}
	fmt.Print(out)

}
