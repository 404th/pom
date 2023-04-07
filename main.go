package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/404th/helloworld/biny"
	"github.com/404th/helloworld/model"
	"github.com/404th/helloworld/pkg/loader"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"

	_ "github.com/go-bindata/go-bindata" // Include generated bindata.go file
)

func main() {
	var task model.Task

	fmt.Println("💂 Welcome to mini 'POMODORO'")
	fmt.Printf("\n")

	reader := bufio.NewReader(os.Stdin) // Create a new buffered reader that reads from the standard input (terminal)
	fmt.Print("\r" + "⚽ Enter name of task: ")
	name, err := reader.ReadString('\n') // Read string until newline character is encountered
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	name = name[:len(name)-1] // Remove the trailing newline character from the string
	task.Name = name

	fmt.Print("\r" + "⌛ Enter time you need to complete the task (minutes): ")
	tm_str, err := reader.ReadString('\n') // Read string until newline character is encountered
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	tm_str = tm_str[:len(tm_str)-1] // Remove the trailing newline character from the string

	tm, err := strconv.Atoi(tm_str)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("\n")
	task.Time = tm

	loader.Load(tm)
	// if err := player.Run(); err != nil {
	if err := run(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n\n😅 HURRAY")

	fmt.Printf("\n🚀 Task '%s' is completed in %d minutes\n", task.Name, task.Time)

	fmt.Println("\n⏱️ TIME IS UP!")
}

func run() error {
	f, err := biny.Asset("music/b1.mp3")
	if err != nil {
		panic(err)
	}

	r := bytes.NewReader(f)

	// currentFolder, err := os.Getwd()
	// if err != nil {
	// 	fmt.Println("Failed to get current folder path:", err)
	// 	return err
	// }

	// Specify a file path relative to the current folder
	// filePath := filepath.Join(currentFolder, "music", "b1.mp3")

	// f, err := os.Open(filePath)
	// if err != nil {
	// 	return err
	// }
	// defer f.Close()

	d, err := mp3.NewDecoder(r)
	if err != nil {
		return err
	}

	c, err := oto.NewContext(d.SampleRate(), 2, 2, 8192)
	if err != nil {
		return err
	}
	defer c.Close()

	p := c.NewPlayer()
	defer p.Close()

	if _, err := io.Copy(p, d); err != nil {
		return err
	}
	return nil
}
