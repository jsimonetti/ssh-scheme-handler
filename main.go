package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"os/exec"
)

type TermProg struct {
	name string
	mkargs func() []string
}

var sshurl string
var progs []TermProg = []TermProg{
	{"konsole",        func() []string {return []string{"--new-tab", "--hold", "-e"}}},
	{"st",             func() []string {return []string{"-e"}}},
	{"tilix",          func() []string {return []string{"-t", sshurl, "-a", "app-new-session", "-e"}}},
	{"alacritty",      func() []string {return []string{"-e"}}},
	{"kitty",          func() []string {return []string{"-e"}}},
	{"gnome-terminal", func() []string {return []string{"--"}}},
}

func init() {
	flag.StringVar(&sshurl, "url", "", "SSH URL to login to")
}

func main() {
	flag.Parse()

	u, err := url.Parse(sshurl)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	scheme := u.Scheme
	if scheme != "ssh" {
		print("not an ssh URL\n")
		os.Exit(1)
	}
	host, port := u.Hostname(), u.Port()
	user := u.User.Username()

	var prog string = ""
	var prognames []string
	var args []string
	for _, p := range progs {
		prognames = append(prognames, p.name)
		_, err := exec.LookPath(p.name)
		if err == nil {
			prog, args = p.name, p.mkargs()
			break
		}
	}
	if prog == "" {
		fmt.Fprintln(os.Stderr, "no terminal program found, tried", prognames)
		os.Exit(1)
	}

	args = append(args, "ssh")
	if port != "" {
		args = append(args, "-p", port)
	}
	if user != "" {
		args = append(args, user + "@" + host)
	} else {
		args = append(args, host)
	}

	cmd := exec.Command(prog, args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		print(err.Error() + "\n")
		os.Exit(1)
	}
}
