package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func handleConnection(c net.Conn) {
	fmt.Fprint(c, "I'm not the one who's always late, nor the one who loves to hate. I don't have bugs in my diet, but without me, there's no quiet. Who am I?\n")
	netData, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	answer := strings.TrimSpace(netData)
	if answer == "hinata" {
		fmt.Fprint(c, "Mount Myōboku.\n")
	} else {
		fmt.Fprint(c, "Beware.\n")
	}
	c.Close()
}

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}
