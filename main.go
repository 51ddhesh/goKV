package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	store := NewKVStore("data.json")
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Init: ")

	for {
		fmt.Print(">>> ")
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		parts := strings.SplitN(line, " ", 3)
		if len(parts) == 0 {
			continue
		}

		switch parts[0] {
		case "set":
			if len(parts) < 3 {
				fmt.Println("Usage: set <key> <value>")
				continue
			}
			store.Set(parts[1], parts[2])
			fmt.Println("OK")

		case "get":
			if len(parts) < 2 {
				fmt.Println("Usage: get <key>")
				continue
			}
			if val, ok := store.Get(parts[1]); ok {
				fmt.Println(val)
			} else {
				fmt.Println("Key not found")
			}

		case "del":
			if len(parts) < 2 {
				fmt.Println("Usage: del <key>")
				continue
			}
			store.Delete(parts[1])
			fmt.Println("Deleted")

		case "exit":
			fmt.Println("Exiting...")
			return

		case "help":
			fmt.Println("Commands: set <k> <v>, get <k>, del <k>, exit")
		default:
			fmt.Println("Unknown command, enter 'help' for help with commands")
		}
	}

}
