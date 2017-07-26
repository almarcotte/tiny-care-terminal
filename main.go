package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	ui "github.com/gizak/termui"
)

func main() {
	if err := ui.Init(); err != nil {
		panic(err)
	}

	defer ui.Close()

	bindQuit()
	bindResize()
	redraw()

	ui.Loop()
}

func bindQuit() {
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})
}

func bindResize() {
	ui.Handle("/sys/wnd/resize", func(e ui.Event) {
		ui.Body.Width = ui.TermWidth()
		ui.Body.Align()
		ui.Clear()
		ui.Render(ui.Body)
	})
}

func redraw() {
	initDashboard()
}

func initDashboard() {
	today := dailyCommits()
	week := dailyCommits()
	water := dailyCommits()
	weather := makeWeather()
	tweet1 := makeWeather()
	tweet2 := makeWeather()

	ui.Body.AddRows(
		ui.NewRow(ui.NewCol(6, 0, today), ui.NewCol(6, 0, weather, tweet1, tweet2)),
		ui.NewRow(ui.NewCol(6, 0, week), ui.NewCol(6, 0, water)),
	)

	ui.Body.Align()
	ui.Render(ui.Body)
}

func dailyCommits() (ls *ui.List) {
	strs := []string{
		"[/Users/alex/go/src/github.com/gnumast/tiny-care-terminal](fg-red)",
		"[e9da701b](fg-green) - Fixed a thing",
		"[d825f5e0](fg-green) - Added a feature",
		"[2a84fe96](fg-green) - Initial commit",
		fmt.Sprintf("Body width: %d", ui.Body.Width),
		fmt.Sprintf("Body height: %d", ui.TermHeight()),
	}

	ls = ui.NewList()
	ls.Items = strs
	ls.ItemFgColor = ui.ColorWhite
	ls.BorderLabel = "Today"
	ls.BorderLabelFg = ui.ColorWhite
	ls.BorderFg = ui.ColorBlue
	ls.Height = (int)(ui.TermHeight() / 2)

	return
}

func makeWeather() (w *ui.Par) {
	w = ui.NewPar("Simple colored text\nwith label. It [can be](fg-red) multilined with \\n or something!")
	w.Height = (int)(ui.TermHeight() / 6)
	w.BorderLabel = " Weather "
	w.BorderFg = ui.ColorBlue
	w.BorderLabelFg = ui.ColorWhite

	return
}

func getTwitterClient() *twitter.Client {
	config := oauth1.NewConfig("consumerKey", "consumerSecret")
	token := oauth1.NewToken("accessToken", "accessSecret")
	httpClient := config.Client(oauth1.NoContext, token)

	return twitter.NewClient(httpClient)
}
