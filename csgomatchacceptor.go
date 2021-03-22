package main

import (
	"bufio"
	"csgomatchacceptor/excepts"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-vgo/robotgo"
)

var POINT_FILENAME string = "point.txt"
var WIN_DETECTED bool = false
var VALID_RGBS = []RGB{{76, 175, 80}, {76, 176, 80}, {90, 203, 94}, {90, 203, 95}, {107, 210, 110}}

type Point struct {
	x, y int
}

type RGB struct {
	R, G, B uint8
}

func (c RGB) check() bool {
	for _, v := range VALID_RGBS {
		if v == c {
			return true
		}
	}
	return false
}

func Detect(point Point) {
	title := robotgo.GetTitle()
	detect := "Counter-Strike: Global Offensive"
	if title == detect {
		if !WIN_DETECTED {
			WIN_DETECTED = true
			fmt.Println("CS:GO Window is Active")
		}
		color := robotgo.GetPixelColor(point.x, point.y)
		values, err := strconv.ParseUint(color, 16, 32)
		if err == nil {
			red := uint8(values >> 16)
			green := uint8((values >> 8) & 0xFF)
			blue := uint8(values & 0xFF)
			rgb := RGB{red, green, blue}
			if rgb.check() {
				robotgo.MoveMouse(point.x, point.y)
				time.Sleep(time.Millisecond * 500)
				robotgo.MouseClick()
				fmt.Println("\n\n==============================================================")
				fmt.Println("Detected! Clicked to the accept button. Waiting 20 seconds...")
				fmt.Print("Then the program will run again because maybe accepted match\nmay be not accepted by everybody.")
				fmt.Println("If you saw the match is\naccepted, you can close the program.")
				fmt.Println("==============================================================")
				time.Sleep(time.Second * 20)
			}
		}
	} else if WIN_DETECTED {
		WIN_DETECTED = false
		fmt.Println("CS:GO Window is Hidden")
	}
}

func DetectThread(point Point) {
	for {
		Detect(point)
		time.Sleep(time.Second)
	}
}

func GetPoint() (Point, error) {
	file, err := os.Open(POINT_FILENAME)
	if err != nil {
		return Point{-1, -1}, &excepts.PointNotFound{File: POINT_FILENAME, Err: err.Error()}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	x, y := -1, -1
	_, err = fmt.Sscanf(line, "%d,%d", &x, &y)
	if err != nil {
		return Point{-1, -1}, &excepts.InvalidPointFormat{File: POINT_FILENAME}
	}
	if x < 0 || y < 0 {
		return Point{-1, -1}, &excepts.InvalidPointCo{File: POINT_FILENAME}
	}
	return Point{x, y}, nil
}

func main() {
	point, err := GetPoint()
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	fmt.Printf("Point axis has been fetched from point.txt: (%d, %d)\n", point.x, point.y)
	fmt.Println("The program started. Open CS:GO's window then you can go away from your pc.")
	DetectThread(point)
}
