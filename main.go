package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/nektro/go-util/util"
	etc "github.com/nektro/go.etc"
	"github.com/rakyll/statik/fs"
)

type DHTData struct {
	Hash        string `json:"infoHash"`
	Name        string
	Peers       int
	ScrapedDate int `json:"scraped_date"`
}

type TrackerData struct {
	Hash        string `json:"infoHash"`
	Seeders     int    `json:"complete"`
	Leechers    int    `json:"incomplete"`
	ScrapedDate int    `json:"scraped_date"`
}

type Torrent struct {
	Hash        string `json:"_id"`
	Name        string
	Size        int `json: "size_bytes"`
	Link        string
	Seeders     int
	Leechers    int
	DHTData     DHTData
	TrackerData map[string]*TrackerData
}

func setInterval(someFunc func(), milliseconds int, async bool) chan bool {

	// How often to fire the passed in function
	// in milliseconds
	interval := time.Duration(milliseconds) * time.Millisecond

	// Setup the ticker and the channel to signal
	// the ending of the interval
	ticker := time.NewTicker(interval)
	clear := make(chan bool)

	// Put the selection in a go routine
	// so that the for loop is non-blocking
	go func() {
		for {

			select {
			case <-ticker.C:
				if async {
					// This won't block
					go someFunc()
				} else {
					// This will block
					someFunc()
				}
			case <-clear:
				ticker.Stop()
				return
			}

		}
	}()

	// We return the channel so we can pass in
	// a value to it to clear the interval
	return clear

}

var torrents []*Torrent

func updateStats() {
	util.Log("Updating stats")
    resp, err := http.Get("https://phillm.net/libgen-stats2.php")
	if err != nil {
		util.LogError("Failed to fetch stats", err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		util.LogError("Failed to read body", err.Error())
		return
	}
    json.Unmarshal(body, &torrents)

	for _, torrent := range torrents {
		seeders := 0
		leechers := 0
		for _, v := range torrent.TrackerData {
			if v.Seeders > seeders {
				seeders = v.Seeders
				leechers = v.Leechers
			}
		}
		torrent.Seeders = seeders
		torrent.Leechers = leechers
	}

	util.Log("Torrent count:", len(torrents))
}

func main() {
	updateStats()

    setInterval(updateStats, 1800*1000, true)

	etc.MFS.Add(http.Dir("www"))

	statikFS, err := fs.New()
	if err == nil {
        etc.MFS.Add(statikFS)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        results := map[string]*Torrent{}
		for _, t := range torrents {
			results[t.Hash] = t
		}
		etc.WriteHandlebarsFile(r, w, "/index.hbs", map[string]interface{}{
			"torrents": results,
		})
	})

    util.Log("Starting on Port 8080")
    util.LogError(http.ListenAndServe(":8080", nil))
}
