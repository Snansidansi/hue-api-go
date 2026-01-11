package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Snansidansi/hue-api-go"
	"github.com/Snansidansi/hue-api-go/models"
	"github.com/Snansidansi/hue-api-go/util"
	"github.com/joho/godotenv"
	"golang.org/x/term"
)

const (
	borderTL = "╭"
	borderTR = "╮"
	borderBL = "╰"
	borderBR = "╯"
	borderH  = "─"
	borderV  = "│"
	boxWidth = 80
)

type Lamp struct {
	ID         string
	Name       string
	R, G, B    uint8
	Brightness int
	IsOn       bool
}

type Controller struct {
	lamps []*Lamp
	idMap map[string]int

	termW, termH int
	xOffset      int
	yOffset      int
	contentH     int

	state *term.State
}

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Printf("Error while loading .env file: %v", err)
		os.Exit(1)
	}

	bridge := models.Bridge{
		Id:       os.Getenv("HUE_BRIDGE_ID"),
		IPAdress: os.Getenv("HUE_BRIDGE_IP"),
	}

	apiKey := os.Getenv("HUE_BRIDGE_USERNAME")
	client := hueapi.NewClient(bridge, apiKey, nil, false)
	LiveLightDisplay(client)
}

func LiveLightDisplay(client *hueapi.Client) {
	lamps := GetLamps(client)
	controller := NewController(lamps)

	if err := controller.Start(); err != nil {
		panic(err)
	}
	defer controller.Stop()

	client.EventStream.Start()
	defer client.EventStream.Stop()

	stop := make(chan struct{})
	go func() {
		buf := make([]byte, 1)
		for {
			os.Stdin.Read(buf)
			if buf[0] == 'q' || buf[0] == 3 { // 'q' or Strg+C
				close(stop)
				return
			}
		}
	}()

	eventStream := client.EventStream.GetEventStream(100)
	for {
		select {
		case <-stop:
			return
		case event, ok := <-eventStream:
			if !ok {
				return
			}
			handleEvent(controller, event)
		}
	}
}

func GetLamps(client *hueapi.Client) []*Lamp {
	hueResponse, err := client.Lights.GetAllLights()
	if err != nil {
		fmt.Printf("Error while getting lights: %v\n", err)
		os.Exit(1)
	}
	if len(hueResponse.Errors) > 0 {
		for _, hueErr := range hueResponse.Errors {
			fmt.Printf("Hue err: %v\n", hueErr.Description)
		}
	}
	if len(hueResponse.Data) == 0 {
		fmt.Println("No light data available")
		os.Exit(1)
	}

	lamps := make([]*Lamp, len(hueResponse.Data))
	for i, light := range hueResponse.Data {

		lamp := &Lamp{
			ID:         light.ID,
			Name:       *light.Metadata.Name,
			Brightness: int(*light.Dimming.Brightness),
			IsOn:       *light.On.On,
		}

		if light.Color != nil {
			r, g, b := util.XYToRGB(light.Color.XY.X, light.Color.XY.Y, *light.Dimming.Brightness)
			lamp.R = uint8(r)
			lamp.G = uint8(g)
			lamp.B = uint8(b)
		}

		lamps[i] = lamp
	}

	return lamps
}

func handleEvent(c *Controller, event any) {
	if lightEvent, ok := event.(*models.LightChangeEvent); ok {
		targetID := lightEvent.ID

		c.Update(targetID, func(l *Lamp) {
			if lightEvent.On != nil && lightEvent.On.On != nil {
				l.IsOn = *lightEvent.On.On
			}
			if lightEvent.Dimming != nil && lightEvent.Dimming.Brightness != nil {
				l.Brightness = int(*lightEvent.Dimming.Brightness)
			}
			if lightEvent.Color != nil {
				bri := float64(l.Brightness)
				if lightEvent.Dimming != nil && lightEvent.Dimming.Brightness != nil {
					bri = *lightEvent.Dimming.Brightness
				}

				r, g, b := util.XYToRGB(
					lightEvent.Color.XY.X,
					lightEvent.Color.XY.Y,
					bri,
				)

				l.R = uint8(r)
				l.G = uint8(g)
				l.B = uint8(b)
			}
		})
	}
}

func NewController(lamps []*Lamp) *Controller {

	idMap := make(map[string]int)
	for i, l := range lamps {
		idMap[l.ID] = i
	}

	return &Controller{
		lamps: lamps,
		idMap: idMap,
	}
}

func (c *Controller) Start() error {
	var err error

	c.termW, c.termH, err = term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return err
	}

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return err
	}
	c.state = oldState

	fmt.Print("\x1b[?25l")
	fmt.Print("\x1b[2J")

	c.contentH = len(c.lamps)
	boxHeight := c.contentH + 2

	c.xOffset = (c.termW - boxWidth) / 2
	c.yOffset = (c.termH - boxHeight) / 2

	if c.xOffset < 0 {
		c.xOffset = 0
	}
	if c.yOffset < 0 {
		c.yOffset = 0
	}

	c.drawBorder()

	for i := range c.lamps {
		c.renderLine(i)
	}

	return nil
}

func (c *Controller) drawBorder() {

	borderColor := "\x1b[38;2;80;80;120m"
	reset := "\x1b[0m"

	fmt.Printf("\x1b[%d;%dH%s%s%s%s%s", c.yOffset+1, c.xOffset+1, borderColor, borderTL, strings.Repeat(borderH, boxWidth-2), borderTR, reset)

	for i := 0; i < c.contentH; i++ {
		fmt.Printf("\x1b[%d;%dH%s%s%s", c.yOffset+2+i, c.xOffset+1, borderColor, borderV, reset)
		fmt.Printf("\x1b[%d;%dH%s%s%s", c.yOffset+2+i, c.xOffset+boxWidth, borderColor, borderV, reset)
	}

	fmt.Printf("\x1b[%d;%dH%s%s%s%s%s", c.yOffset+c.contentH+2, c.xOffset+1, borderColor, borderBL, strings.Repeat(borderH, boxWidth-2), borderBR, reset)

	title := " HUE LIGHTS "
	fmt.Printf("\x1b[%d;%dH\x1b[1;37m%s\x1b[0m", c.yOffset+1, c.xOffset+4, title)
}

func (c *Controller) Stop() {
	fmt.Print("\x1b[?25h")
	fmt.Print("\x1b[2J")
	fmt.Print("\x1b[H")

	if c.state != nil {
		term.Restore(int(os.Stdin.Fd()), c.state)
	}
}

func (c *Controller) Update(id string, modifier func(*Lamp)) {

	index, ok := c.idMap[id]
	if !ok {

		return
	}

	modifier(c.lamps[index])
	c.renderLine(index)
}

func (c *Controller) renderLine(index int) {
	l := c.lamps[index]

	posY := c.yOffset + 2 + index
	posX := c.xOffset + 3

	fmt.Printf("\x1b[%d;%dH", posY, posX)

	contentWidth := boxWidth - 4

	var colorCode string
	if !l.IsOn {
		colorCode = "\x1b[38;2;100;100;100m"
	} else {
		if l.R == 0 && l.G == 0 && l.B == 0 {
			colorCode = "\x1b[38;2;255;255;255m"
		} else {
			colorCode = fmt.Sprintf("\x1b[38;2;%d;%d;%dm", l.R, l.G, l.B)
		}
	}

	statusIcon := "○"
	if l.IsOn {
		statusIcon = "●"
	}

	nameDisplay := l.Name
	if len(nameDisplay) > 15 {
		nameDisplay = nameDisplay[:14] + "…"
	}

	barLen := 20
	filledLen := min(max((l.Brightness*barLen)/100, 0), barLen)

	bar := strings.Repeat("█", filledLen) + strings.Repeat("░", barLen-filledLen)

	output := fmt.Sprintf("%s [%s] %-15s [%s] %3d%%",
		colorCode, statusIcon, nameDisplay, bar, l.Brightness)

	visibleLen := 1 + 3 + 15 + 1 + barLen + 1 + 5
	padding := max(contentWidth-visibleLen, 0)

	fmt.Print(output + "\x1b[0m" + strings.Repeat(" ", padding))

	borderColor := "\x1b[38;2;80;80;120m"
	fmt.Printf("\x1b[%d;%dH%s%s\x1b[0m", posY, c.xOffset+boxWidth, borderColor, borderV)
}
