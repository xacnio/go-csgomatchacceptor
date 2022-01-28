package main

import (
	"encoding/json"
	"fmt"
	"image"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
	tb "gopkg.in/tucnak/telebot.v2"
)

const VERSION = "1.0.3"

type Config struct {
	TelegramBotToken string `json:"tgtoken"`
	TelegramUserID   int    `json:"tguserid"`
	Test             bool   `json:"test"`
}

var config = Config{}
var WIN_DETECTED bool = false
var BOT *tb.Bot = nil

func main() {
	fmt.Printf("go-csgomatchacceptor - v%s\n", VERSION)
	fmt.Println("Started! Switch to CS:GO's window then you can go away from your pc.")

	config = LoadConfiguration("config.json")

	if strings.Contains(config.TelegramBotToken, ":") && config.TelegramUserID != 0 {
		b, err := tb.NewBot(tb.Settings{
			Token:  config.TelegramBotToken,
			Poller: &tb.LongPoller{Timeout: 10 * time.Second},
		})
		if err != nil {
			log.Fatal(err)
			return
		}
		BOT = b
		fmt.Println("Telegram Bot Started!")
		go b.Start()
	}
	DetectThread()
}

func DetectThread() {
	for {
		Detect()
		time.Sleep(time.Second * 1)
	}
}

func Detect() {
	title := robotgo.GetTitle()
	detect := "Counter-Strike: Global Offensive"
	if strings.HasPrefix(title, detect) {
		if !WIN_DETECTED {
			WIN_DETECTED = true
			fmt.Println("CS:GO Window is Active")
		}
		sx, sy := robotgo.GetScreenSize()
		bitmap := robotgo.CaptureScreen(44*sx/100, 11*sy/29, sx/8, sy/15)
		defer robotgo.FreeBitmap(bitmap)
		robotgo.SaveBitmap(bitmap, "csgomatchacceptor_tmp.png")
		file, err := os.Open("csgomatchacceptor_tmp.png")

		if err != nil {
			fmt.Println("Error: File could not be opened")
			return
		}

		defer file.Close()
		pixels, err := getPixels(file, 10)

		if err != nil {
			fmt.Println("Error: Image could not be decoded")
			return
		}
		count := 0
		total := 0
		for i := 0; i < len(pixels); i++ {
			for _, x := range pixels[i] {
				if x.check() {
					count++
				}
				total++
			}
		}
		percent := (count * 100) / total
		if percent >= 60 {
			fmt.Printf("\n\n%%%v\n", percent)
			fmt.Println("==============================================================")
			fmt.Println("Detected! Clicked to the accept button. Waiting 20 seconds...")
			fmt.Print("Then the program will run again because maybe accepted match\nmay be not accepted by everybody.")
			fmt.Println("If you saw the match is\naccepted, you can close the program.")
			fmt.Println("==============================================================")
			if BOT != nil {
				BOT.Send(tb.ChatID(config.TelegramUserID), "*CSGO MATCH DETECTED*", &tb.SendOptions{ParseMode: "MarkdownV2"})
			}
			robotgo.MoveMouse(sx/2, 20*sy/50)
			time.Sleep(time.Millisecond * 500)
			if !config.Test {
				robotgo.MouseClick()
			}
			time.Sleep(time.Second * 20)
		}
	} else if WIN_DETECTED {
		WIN_DETECTED = false
		fmt.Println("CS:GO Window is Hidden")
	}
}

func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

type Pixel struct {
	R int
	G int
	B int
	A int
}

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}

func (c Pixel) check() bool {
	if c.R >= 75 && c.R <= 108 && c.G >= 175 && c.G <= 215 && c.B >= 75 && c.B <= 115 {
		return true
	}
	return false
}

func getPixels(file io.Reader, maxHeight int) ([][]Pixel, error) {
	img, _, err := image.Decode(file)

	if err != nil {
		return nil, err
	}
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	if height > maxHeight {
		height = maxHeight
	}
	var pixels [][]Pixel
	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}
	return pixels, nil
}
