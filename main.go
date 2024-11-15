package main

import (
	"fmt"
	"strings"

	"github.com/alicebob/miniredis/v2"
	"github.com/alicebob/miniredis/v2/server"
)

func main() {
	// f, err := os.ReadFile("a.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(f))
	// a, _ := server.NewServer("127.0.0.1:6379")
	z := miniredis.NewMiniRedis()
	if err := z.StartAddr(":6379"); err != nil {
		panic(err)
	}
	defer z.Close()

	// cb := func(c *server.Peer, cmd string, args []string) {
	// 	fmt.Println(cmd, strings.Join(args, " "))
	// 	c.WriteStrings([]string{"OK"})
	// }

	// cmdCommand := func(c *server.Peer, cmd string, args []string) {
	// 	fmt.Println(cmd, strings.Join(args, " "))
	// 	c.WriteRaw("*2\r\n*10\r\n$14\r\nJSON.ARRINSERT\r\n:-5\r\n*4\r\n:0\r\n*0\r\n")
	// }

	// if err := z.Server().Register("COMMAND", cmdCommand); err != nil {
	// 	panic(err)
	// }

	// if err := z.Server().Register("COMMAND", server.Cmd(cb)); err != nil {
	// 	panic(err)
	// }

	// z.SetError("err")
	cb := func(c *server.Peer, cmd string, args ...string) bool {
		if !(cmd == "COMMAND" && len(args) == 0) {
			fmt.Println(">> ", cmd, strings.Join(args, " "))
		}
		return false
	}

	// var c server.Hook
	// func(*Peer, string, ...string) bool
	z.Server().SetPreHook(server.Hook(cb))

	fmt.Println("listening on", z.Addr())
	select {}
}
