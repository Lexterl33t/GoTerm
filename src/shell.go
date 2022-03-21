package shell

import (
	"fmt"
	"strings"
	"syscall"
)

func Exec_shell(commands *Commands_t) []string {

	for _, sh := range commands.Commands {
		/*
			str := strings.Split(sh.Binary+" "+strings.Join(sh.Args, " "), " ")
			fmt.Println(str)
			err := syscall.Exec(sh.Binary, str, nil)
			if err != nil {
				fmt.Println(err)
			}
		*/
		str := strings.Split(sh.Binary+" "+strings.Join(sh.Args, " "), " ")
		fmt.Println(str)

		err := syscall.Exec(sh.Binary, str, nil)
		if err != nil {
			fmt.Println(syscall.Stdout)
		}
	}

	return []string{}
}
