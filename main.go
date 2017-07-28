package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	ui "github.com/gizak/termui"
	"net/url"
)

type Config struct {
	Twitter *anaconda.TwitterApi
}

var config Config

func main() {
	if err := ui.Init(); err != nil {
		panic(err)
	}

	defer ui.Close()

	config.Twitter = getTwitterClient()

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
	ui.Handle("/sys/wnd/resize", func(ui.Event) {
		ui.Body.Width = ui.TermWidth()
		ui.Body.Align()
		redraw()
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
	magic := magicRealismBox()
	tweet2 := makeWeather()

	ui.Body.Rows = []*ui.Row{}

	ui.Body.AddRows(
		ui.NewRow(ui.NewCol(6, 0, today), ui.NewCol(6, 0, weather, magic, tweet2)),
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

func magicRealismBox() (w *ui.Par) {
	v := url.Values{}
	v.Set("count", "5")
	v.Set("screen_name", "magicrealismbot")

	result, err := config.Twitter.GetUserTimeline(v)
	if err != nil {
		panic(err)
	}

	w = ui.NewPar(fmt.Sprintf("%v", result))
	w.Height = (int)(ui.TermHeight() / 6)
	w.BorderLabel = " [@](fg-green)f0sdf09df "
	w.BorderFg = ui.ColorBlue
	w.BorderLabelFg = ui.ColorWhite

	return
}

func getTwitterClient() *anaconda.TwitterApi {
	anaconda.SetConsumerKey("consume")
	anaconda.SetConsumerSecret("secret")

	return anaconda.NewTwitterApi("access", "secret")
}
