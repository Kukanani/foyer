package main

import (
	"log"
	"path"
	"os"
	"io/fs"
	"sort"
	"strconv"
	"fmt"
	"bufio"
	"strings"
	"time"
)

type Message struct {
	file_info  fs.FileInfo
	short_text string
	full_text  string
}

type message_manager struct {
	msg_dir_name string
	messages []Message
}

func TimeStamp() string {
	ts := time.Now().UTC().Format(time.RFC3339)
	return strings.Replace(ts, ":", "_", -1)
}

func (mm *message_manager) CreateMessage(short_text string, long_text string) {
	out_file := path.Join(mm.msg_dir_name, TimeStamp()+".txt")
	err := os.WriteFile(out_file, []byte(short_text+"\n"+long_text), 0666)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Message #" + strconv.Itoa(len(mm.messages)) + " created at " + out_file)
	}
}

func (mm *message_manager) PrintMessages() {
	fmt.Println("There are", len(mm.messages), "message(s):")

	for msg_idx, msg := range mm.messages {
		short_text := msg.short_text
		max_width := 76
		consumed := 0
		var lines []string
		for {
			end_idx := len(short_text)
			if consumed+max_width < end_idx {
				end_idx = consumed + max_width
			}
			lines = append(lines, short_text[consumed:end_idx])
			if max_width+consumed >= len(short_text) {
				break
			}
			consumed += max_width
		}
		for line_idx, line := range lines {
			// Padding for lines after the first one
			left_col := "    "
			if line_idx == 0 {
				left_col = "[" + strconv.Itoa(msg_idx) + "] "
			}
			fmt.Println(left_col + line)
		}
	}
}

func (mm *message_manager) DeleteMessage(msg_no int) {
	if msg_no >= len(mm.messages) || msg_no < 0 {
		fmt.Println("Message index out of bounds")
		return
	}
	file_name := path.Join(mm.msg_dir_name, mm.messages[msg_no].file_info.Name())
	err := os.Remove(file_name)
	if err != nil {
		log.Fatal(err)
	} else {
		mm.messages = append(mm.messages[:msg_no], mm.messages[msg_no+1:]...)
		fmt.Println("message #" + strconv.Itoa(msg_no) + " (" + file_name + ") deleted")
	}
}

func (mm *message_manager) PrintFullMessage(msg_no int) {
	if msg_no >= len(mm.messages) || msg_no < 0 {
		fmt.Println("Message index out of bounds")
		return
	}
	fmt.Println("Details for message #" + strconv.Itoa(msg_no) + ":")
	fmt.Println("Short text:   ", mm.messages[msg_no].short_text)
	fmt.Println("Filename:     ", path.Join(mm.msg_dir_name, mm.messages[msg_no].file_info.Name()))
	fmt.Println("Last modified:", mm.messages[msg_no].file_info.ModTime().Format("2006-01-02 15:04:05"))
	fmt.Println("Full text:")
	fmt.Println()
	print(mm.messages[msg_no].full_text)
	fmt.Println()
}

func NewMessageManager(msg_dir_name string) message_manager {
	messages := LoadMessages(msg_dir_name)

	mm := message_manager{msg_dir_name, messages}
	return mm
}

func LoadMessages(msg_dir_name string) []Message {

	// load all stored messages
	var messages []Message

	err := fs.WalkDir(os.DirFS(msg_dir_name), ".", func(this_path string, entry fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if this_path == "." {
			return nil
		}
		file_info, err := entry.Info()
		if err != nil {
			log.Fatal(err)
		}

		file, err := os.Open(path.Join(msg_dir_name, this_path))
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		scanner.Scan()
		line1 := scanner.Text()
		the_rest := ""
		for scanner.Scan() {
			the_rest = the_rest + scanner.Text() + "\n"
		}

		var message Message
		message.file_info = file_info
		message.short_text = line1
		message.full_text = the_rest
		messages = append(messages, message)
		return nil
	})
	sort.Slice(messages, func(i, j int) bool {
		return messages[i].file_info.ModTime().Before(messages[j].file_info.ModTime())
	})
	if err != nil {
		log.Fatal(err)
	}
	return messages
}