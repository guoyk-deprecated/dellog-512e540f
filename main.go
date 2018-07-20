package main

import (
	"flag"
	"log"
	"os"
	"time"
)

var (
	configFiles string
	dryRun      bool
)

func main() {
	var err error

	// parse flags
	flag.StringVar(&configFiles, "c", "/etc/dellog.d/*.yml", "config file glob pattern")
	flag.BoolVar(&dryRun, "dry-run", false, "set dry-run flag, no file will be deleted")
	flag.Parse()
	log.Println("scan  :", configFiles)

	// load configs
	var configs []Config
	if configs, err = LoadConfigs(configFiles); err != nil {
		return
	}
	log.Println("found :", len(configs), "config(s)")
	// iterate config files
	for _, c := range configs {
		// list files
		var files []string
		if files, err = c.ListFiles(); err != nil {
			log.Println("error : failed to list files for", c)
			err = nil
			continue
		}
		// print informations
		log.Println("------------------------------------")
		log.Println("scan  :", c.Paths)
		log.Println("keep  :", c.Keep, "day(s)")
		log.Println("------------------------------------")
		// iterate files
		for _, file := range files {
			// search date mark, skip if not found
			var t time.Time
			var ok bool
			if t, ok = FindDateMark(file); !ok {
				continue
			}
			// compare time and delete file
			if c.IsExpired(BeginningOfDay(), t) {
				if dryRun {
					log.Println("drydel:", file)
				} else {
					log.Println("delete:", file)
					os.Remove(file)
				}
			} else {
				log.Println("skip  :", file)
			}
		}
	}
}
