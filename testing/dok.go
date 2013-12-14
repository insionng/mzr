package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func Superkill(port int64) error {

	if s, e := Getpid(port); e != nil {
		return e
	} else {

		cmd := exec.Command("kill", "-9", s)

		if e := cmd.Run(); e != nil {
			fmt.Println("Removal process fails,pid:", s, "!")
			return e
		} else {
			fmt.Println("Remove the success of the process,pid:", s, "!")

			return nil
		}

	}

}

func ListPidLine(port int64) (string, error) {

	cmd1 := exec.Command("netstat", "-nlp")

	if buf, err := cmd1.Output(); err != nil {
		fmt.Println(err)
		return "", err
	} else {
		s := string(buf)
		sp := strings.SplitN(s, "\n", -1)
		for _, v := range sp {
			//	fmt.Println(k, v)
			if strings.Contains(v, ":"+strconv.Itoa(int(port))) {
				s = v

			}
		}
		if s != "" {
			return s, nil
		} else {
			return "", errors.New("No such port number corresponding to the PID line!")
		}

	}

	return "", errors.New("Find PID line error!")
}

func Getpid(port int64) (string, error) {
	s, e := ListPidLine(port)
	if e != nil {
		return "", e
	} else {

		sf := []string{}

		sf = strings.Fields(s)
		for k, v := range sf {
			if k == 6 {

				sf = strings.SplitN(v, "/", -1)
				s = sf[0]
				return s, nil
			}
		}
		return "", errors.New("Find PID number error")
	}
}

func main() {
	al := len(os.Args)
	if al > 1 {

		switch {
		case al == 3:
			cmdname := os.Args[1]
			port, e := strconv.Atoi(os.Args[2])

			if e != nil {
				fmt.Println("Your argument is not numeric port!ERR:", e)
				return
			}
			p := int64(port)

			if cmdname == "kill" {

				if e := Superkill(p); e != nil {
					fmt.Println("::dok:: Error:", e)
				} else {
					fmt.Println("::dok:: Okay!")
				}
			}

			if cmdname == "list" {

				if s, e := ListPidLine(p); e != nil {
					fmt.Println("::dok:: List Pid Line Error:", e)
				} else {
					fmt.Println("::dok:: List Pid Line Okay:")
					fmt.Println(s)
				}
			}

			if cmdname == "get" {

				if s, e := Getpid(p); e != nil {
					fmt.Println("::dok:: Get Pid Error:", e)
				} else {
					fmt.Println("::dok:: Get Pid Okay:")
					fmt.Println(s)
				}
			}

		case al == 2:

			port, e := strconv.Atoi(os.Args[1])
			if e != nil {
				fmt.Println("Your argument is not numeric port!ERR:", e)
				return
			}
			p := int64(port)

			if e := Superkill(p); e != nil {
				fmt.Println("::dok:: Error:", e)
			} else {
				fmt.Println("::dok:: Okay!")
			}
		default:
			fmt.Println("usage: dok port")

		}

	} else {
		fmt.Println("usage: dok port")
	}

}
