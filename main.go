package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/404th/helloworld/model"
	"github.com/404th/helloworld/pkg/loader"
	"github.com/404th/helloworld/pkg/player"

	_ "github.com/go-bindata/go-bindata" // Include generated bindata.go file
)

var task model.Task

func main() {
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
	task.Time = tm

	fmt.Print("\r" + "🥁 Enter music number from 1 to 5: ")
	m_str, err := reader.ReadString('\n') // Read string until newline character is encountered
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	m_str = m_str[:len(m_str)-1] // Remove the trailing newline character from the string

	m, err := strconv.Atoi(m_str)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("\n")
	task.Music = m

	// functions
	done := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(2)

	go loader.Load(tm, done, &wg)

	go func(task *model.Task, done chan bool, wg *sync.WaitGroup) {
		for {
			select {
			case <-done:
				wg.Done()
				return
			default:
				player.Run(task)
			}
		}
	}(&task, done, &wg)

	wg.Wait()

	fmt.Println("\n\n😅 HURRAY")

	fmt.Printf("\n🚀 Task '%s' is completed in %d minutes\n", task.Name, task.Time)

	fmt.Println("\n⏱️ TIME IS UP!")
}
