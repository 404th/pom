package player

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

func Run() error {
	currentFolder, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get current folder path:", err)
		return err
	}

	// Specify a file path relative to the current folder
	filePath := filepath.Join(currentFolder, "music", "b1.mp3")

	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
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
