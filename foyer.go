package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	msg_dir_name := "/opt/foyer"
	mm := NewMessageManager(msg_dir_name)

	// parse command line to see if we are adding
	if len(os.Args) > 1 {
		if os.Args[1] == "add" {
			fmt.Println("Adding a message to the foyer.")
			fmt.Println("Write the short version of your message below. Press Enter to finish")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			short_text := scanner.Text()

			fmt.Println("Write the long version of your message below. You can add newlines.")
			fmt.Println("To complete the message, add a line consisting only of '~~~'")

			long_text := ""
			for scanner.Scan() {
				this_line := scanner.Text()
				if this_line == "~~~" {
					break
				} else {
					long_text = long_text + this_line + "\n"
				}
			}
			mm.CreateMessage(short_text, long_text)

		} else {
			fmt.Println("I don't understand those arguments")
			fmt.Println("You can run foyer with no arguments, or with 'add' to add a message")
		}
		return
	}

	for {
		if len(mm.messages) > 0 {
			fmt.Print("\033[H\033[2J")
			fmt.Println("You are in the foyer.")
			mm.PrintMessages()
			fmt.Println("Enter to exit, message number to view full details, or 'd <N>' to delete message N")
			var command string
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			command = scanner.Text()

			if len(command) == 0 {
				fmt.Println("Bye!")
				break
			} else if command[:1] == "d" {
				msg_no, err := strconv.Atoi(command[2:])
				if err != nil {
					fmt.Println("I don't understand that message number")
					continue
				}
				mm.DeleteMessage(msg_no)
			} else {
				msg_no, err := strconv.Atoi(command)
				if err != nil {
					fmt.Println("I don't understand that message number")
					continue
				}
				fmt.Print("\033[H\033[2J")
				mm.PrintFullMessage(msg_no)
				fmt.Println("Press Enter to return to the foyer")
				scanner.Scan()
			}
		} else {
			fmt.Println("There are no messages. Run 'foyer add' to add one.")
			break
		}
	}
}
