package shell

import (
	"strings"
)

type Command_t struct {
	Binary string
	Args   []string
}

type Commands_t struct {
	Commands []Command_t
}

const (
	BINARY = iota
	ARGS
	AND
	OR
	XOR
	PIPE
)

var OP = map[string]int32{
	"&&": AND,
	"^":  XOR,
	"|":  PIPE,
	"||": OR,
}

func Parses(command string) (*Commands_t, error) {
	lol := strings.FieldsFunc(command, func(r rune) bool {
		if r == '&' || r == '|' {
			return true
		} else {

		}
		return false
	})

	var res [][]string
	for _, l := range lol {
		lel := strings.FieldsFunc(l, func(r rune) bool {
			if r == ' ' {
				return true
			}

			return false
		})
		res = append(res, lel)
	}

	var commandd Commands_t

	for _, d := range res {
		var comman Command_t
		for i, j := range d {
			if i == 0 {
				if len(j) > 6 {
					if j[0:5] == "/bin/" {
						comman.Binary = j
					}
				} else {
					comman.Binary = "/bin/" + j
				}
				continue
			}

			comman.Args = append(comman.Args, j)
		}

		commandd.Commands = append(commandd.Commands, comman)

	}

	return &commandd, nil
}
