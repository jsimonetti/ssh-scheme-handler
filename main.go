package main

import (
	"flag"
	"net/url"
	"os"
	"os/exec"
)

var sshurl string

func init() {
	flag.StringVar(&sshurl, "url", "", "SSH URL to login to")
}

func main() {
	flag.Parse()

	u, err := url.Parse(sshurl)
	if err != nil {
		print(err.Error() + "\n")
		os.Exit(1)
	}

	scheme := u.Scheme
	if scheme != "ssh" {
		print("not an ssh URL\n")
		os.Exit(1)
	}
	host, port := u.Hostname(), u.Port()
	user := u.User.Username()

	exe := "ssh "
	if port != "" {
		exe = exe + "-p " + port + " "
	}
	if user != "" {
		host = user + "@" + host
	}

	exe = exe + host

	cmd := exec.Command("pantheon-terminal", "-e", exe)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		print(err.Error() + "\n")
		os.Exit(1)
	}
}
