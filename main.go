package main

import (
	"github.com/nsf/termbox-go"
	"math"
	"math/rand"
	"time"
)

// import "strconv"
//import "fmt" //fmt.Println

func hanabi(posx, posy, starty, size int) {

	//打ち上げ
	for y := starty; y > posy; y -= 2 {
		termbox.SetCell(posx, y, '●', termbox.ColorWhite, termbox.ColorDefault)
		termbox.Flush()
		time.Sleep(50 * time.Millisecond)
	}
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	//爆発
	rx := 2.2
	ry := 1.0
	colors := []int{2, 3, 4, 5, 6, 7, 8}
	for f := 1; f < size; f += 1 {
		color := choice(colors)
		rx *= 1.0 + 0.1*float64(f)
		ry *= 1.0 + 0.1*float64(f)
		for i := 0; i < 360; i += 10 {
			x1 := int(float64(posx) + math.Cos(float64(i))*rx)
			y1 := int(float64(posy) + math.Sin(float64(i))*ry)
			termbox.SetCell(x1, y1, '●', termbox.Attribute(color), termbox.ColorDefault)
		}
		termbox.Flush()
		time.Sleep(30 * time.Millisecond)
	}

}

func main() {
	//初期化・エラー処理
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	//終了処理
	defer termbox.Close()

	event_queue := make(chan termbox.Event)
	go func() {
		for {
			//fmt.Println("Event用のgorutine")
			event_queue <- termbox.PollEvent() //受け取られるまでブロック
		}
	}()

loop:
	for {
		select {
		case ev := <-event_queue: //送信されるまでブロック
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
				break loop //for文を抜けるためのラベル
			}
		default:

			//fmt.Println("mainのgorutine")
			w, h := termbox.Size()
			x := w / 2
			y := h / 2

			funcs := []func(){
				func() { hanabi(x, y, h, 9) },
				func() { hanabi(x-30, y+10, h, 8) },
				func() { hanabi(x+30, y+10, h, 8) },
				func() { hanabi(x+60, y-15, h, 7) },
				func() { hanabi(x-50, y-15, h, 7) },
			}
			shuffle(funcs)
			for _, f := range funcs {
				f()
			}

		}
	}
}

func choice(s []int) int {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(s))
	return s[i]
}

func shuffle(list []func()) {
	for i := len(list); i > 1; i-- {
		j := rand.Intn(i) // 0～(i-1) の乱数発生
		list[i-1], list[j] = list[j], list[i-1]
	}
}
