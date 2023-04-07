package player

import (
	"bytes"
	"fmt"
	"io"

	"github.com/404th/helloworld/biny"
	"github.com/404th/helloworld/model"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

func Run(task *model.Task) error {
	road_to_music := fmt.Sprintf("music/%d.mp3", task.Music)

	f, err := biny.Asset(road_to_music)
	if err != nil {
		return err
	}

	r := bytes.NewReader(f)

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
