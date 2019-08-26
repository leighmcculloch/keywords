package main

import "github.com/gizak/termui"

var labels = []string{"Lua", "Go", "Erlang", "C", "Python", "Ruby", "JS", "Java", "Rust", "Dart", "Swift", "C#", "C++"}
var values = []int{22, 25, 26, 32, 35, 36, 49, 51, 52, 54, 89, 100, 109}

func main() {
	if err := termui.Init(); err != nil {
		panic(err)
	}
	defer termui.Close()

	c := termui.NewBarChart()
	c.DataLabels = labels
	c.Data = values
	c.BarWidth = 6
	c.BarGap = 1
	c.Width = c.BarWidth*len(labels) + len(labels) + 2
	c.Height = 20
	c.CellChar = '.'

	termui.Render(c)

	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.StopLoop()
	})
	termui.Loop()
}
