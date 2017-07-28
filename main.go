package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	ui "github.com/gizak/termui"
	"github.com/gnumast/tiny-care-terminal/git"
	"net/url"
	"os"
)

type Config struct {
	Twitter      *anaconda.TwitterApi
	Repositories []*git.Repository
}

var config Config

func main() {
	if err := ui.Init(); err != nil {
		panic(err)
	}

	defer ui.Close()

	config.Twitter = getTwitterClient()
	config.Repositories = git.ToRepositories(os.Getenv("GIT_REPOSITORIES"))

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
	today := getCommits()
	week := getCommits()
	water := getCommits()
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

func getCommits() (ls *ui.List) {
	commits := []string{}

	ls = ui.NewList()
	ls.Items = commits
	ls.ItemFgColor = ui.ColorWhite
	ls.BorderLabel = " Commits "
	ls.BorderLabelFg = ui.ColorWhite
	ls.BorderFg = ui.ColorBlue
	ls.Height = int(ui.TermHeight() / 2)

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

	// Use the height of the box to determine how many tweets to display
	height := int(ui.TermHeight() / 6)

	numberOfTweets := height

	if height > len(result) {
		numberOfTweets = len(result)
	}

	tweets := ""

	for _, tweet := range result[0:numberOfTweets] {
		tweets += fmt.Sprintf(" - %s\n", tweet.Text)
	}

	w = ui.NewPar(tweets)
	w.Height = height
	w.BorderLabel = " [@](fg-green)f0sdf09df "
	w.BorderFg = ui.ColorBlue
	w.BorderLabelFg = ui.ColorWhite

	return
}

func getTwitterClient() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(os.Getenv("TWITTER_CONSUMER"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_CONSUMER_SECRET"))

	return anaconda.NewTwitterApi(os.Getenv("TWITTER_ACCESS"), os.Getenv("TWITTER_SECRET"))
}
