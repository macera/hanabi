package main

import (
	"github.com/nsf/termbox-go"
	"math"
	"math/rand"
	"time"
)

// import "strconv"
// import "fmt"

func burn(posx, posy, starty, size int) {

	for y := starty; y > posy; y -= 2 {
		termbox.SetCell(posx, y, '●', termbox.Attribute(8), termbox.ColorDefault)
		termbox.Flush()
		time.Sleep(50 * time.Millisecond)
	}
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	//爆発
	//円の数
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
			//termbox.SetCell(x1, y1, '●', termbox.Attribute(f+1), termbox.ColorDefault)
		}
		termbox.Flush()
		time.Sleep(30 * time.Millisecond)

	}

	//termbox.Flush()
	//time.Sleep(5 * time.Millisecond)

}

func monsterball(posx, posy, starty int) {

	for y := starty; y > posy; y -= 2 {
		termbox.SetCell(posx, y, '●', termbox.Attribute(8), termbox.ColorDefault)
		termbox.Flush()
		time.Sleep(100 * time.Millisecond)
	}
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	//爆発
	rx := 2.2 * 4.0
	ry := 1.0 * 4.0
	colors := []int{2, 3, 4, 5, 6, 7, 8}
	var startlinex1 int
	var liney1 int
	var endlinex1 int
	var startlinex2 int
	var liney2 int
	var endlinex2 int

	for f := 4; f < 8; f += 1 {
		color := choice(colors)
		rx *= 1.0 + 0.1*float64(f)
		ry *= 1.0 + 0.1*float64(f)
		if f == 6 {
			continue
		}
		for i := 0; i < 360; i += 10 {
			x1 := int(float64(posx) + math.Cos(float64(i))*rx)
			y1 := int(float64(posy) + math.Sin(float64(i))*ry)
			termbox.SetCell(x1, y1, '●', termbox.Attribute(color), termbox.ColorDefault)
			if f == 5 {
				if i == 120 {
					startlinex1 = x1
					liney1 = y1
				}
				if i == 260 {
					startlinex2 = x1
					liney2 = y1
				}
			}
			if f == 7 {
				if i == 120 {
					endlinex1 = x1

				}
				if i == 260 {
					endlinex2 = x1

				}
			}
		}
	}
	moveRight(startlinex1, endlinex1, liney1)
	moveLeft(startlinex2, endlinex2, liney2)

	termbox.Flush()
	time.Sleep(700 * time.Millisecond)

}

// func burn() {
// w, h := termbox.Size()
// termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
// //打ち上げ
// x := w / 2
// for y := h; y > h/2; y -= 2 {
// 	termbox.SetCell(x, y, '●', termbox.Attribute(8), termbox.ColorDefault)
// 	termbox.Flush()
// 	time.Sleep(100 * time.Millisecond)

// 	// termbox.SetCell(x, y+2, ' ', termbox.ColorDefault,
// 	// 	termbox.Attribute(1))
// 	// a := strconv.Itoa(y)
// 	// fmt.Print(a)
// }
// termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
// //爆発
// y := h / 2
// //円の数
// rx := 2.2
// ry := 1.0
// colors := []int{2, 3, 4, 5, 6, 7, 8}
// for f := 1; f < 9; f += 1 {
// 	color := choice(colors)
// 	rx *= 1.0 + 0.1*float64(f)
// 	ry *= 1.0 + 0.1*float64(f)
// 	for i := 0; i < 360; i += 10 {
// 		x1 := int(float64(x) + math.Cos(float64(i))*rx)
// 		y1 := int(float64(y) + math.Sin(float64(i))*ry)
// 		termbox.SetCell(x1, y1, '●', termbox.Attribute(color), termbox.ColorDefault)
// 		//termbox.SetCell(x1, y1, '●', termbox.Attribute(f+1), termbox.ColorDefault)
// 	}
// 	termbox.Flush()
// 	time.Sleep(30 * time.Millisecond)

// }

// termbox.Flush()
// time.Sleep(700 * time.Millisecond)

// }
func moveRight(fromx, tox, y int) {
	for x := fromx; x < tox; x += 2 {
		termbox.SetCell(x, y, '●', termbox.Attribute(8), termbox.ColorDefault)
		termbox.SetCell(x+8, y-5, '●', termbox.Attribute(8), termbox.ColorDefault)
	}
}
func moveLeft(fromx, tox, y int) {
	for x := fromx; x > tox; x -= 2 {
		termbox.SetCell(x, y, '●', termbox.Attribute(8), termbox.ColorDefault)
		termbox.SetCell(x-8, y-5, '●', termbox.Attribute(8), termbox.ColorDefault)
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	event_queue := make(chan termbox.Event)
	go func() {
		for {
			event_queue <- termbox.PollEvent()
		}
	}()

loop:
	for {
		select {
		case ev := <-event_queue:
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
				break loop
			}
		default:

			w, h := termbox.Size()
			//termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			x := w / 2
			y := h / 2

			funcs := []func(){
				func() { burn(x, y, h, 9) },
				func() { burn(x-30, y+10, h, 8) },
				func() { burn(x+30, y+10, h, 8) },
				func() { burn(x+60, y-15, h, 7) },
				func() { burn(x-50, y-15, h, 7) },
			}
			shuffle(funcs)
			for _, f := range funcs {
				f()
			}

			//monsterball(x+30, y, h)

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
