package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-yaml/yaml"
)

// Config config struct
type Config struct {
	Enabled bool     `yaml:"enabled"`
	Paths   []string `yaml:"paths"`
	Keep    int      `yaml:"keep"`
}

// ListFiles search all files in Config.Paths
func (c Config) ListFiles() (ret []string, err error) {
	// allocate slice
	ret = []string{}
	// iterate paths
	for _, path := range c.Paths {
		// skip empty
		path = strings.TrimSpace(path)
		if len(path) == 0 {
			continue
		}
		// glob filenames
		var filenames []string
		if filenames, err = filepath.Glob(path); err != nil {
			return
		}
		// iterate file
		for _, filename := range filenames {
			// skip duplicates
			if StrSliceContains(ret, filename) {
				continue
			}
			// skip dir and unaccessable file
			if fi, err1 := os.Stat(filename); err1 != nil || fi.IsDir() {
				continue
			}
			ret = append(ret, filename)
		}
	}
	return
}

// IsExpired check expired with current time and target time given
func (c Config) IsExpired(n, t time.Time) bool {
	return n.Sub(t)/(time.Hour*24) > time.Duration(c.Keep)
}

// LoadConfigs load configs from file glob pattern
func LoadConfigs(pattern string) (ret []Config, err error) {
	// allocate slice
	ret = []Config{}

	// glob filenames
	var filenames []string
	if filenames, err = filepath.Glob(pattern); err != nil {
		return
	}

	// iterate filenames
	for _, filename := range filenames {
		var buf []byte
		var cw struct {
			Configs []Config `yaml:"configs"`
		}
		// read file
		if buf, err = ioutil.ReadFile(filename); err != nil {
			return
		}
		// unmarshal
		if err = yaml.Unmarshal(buf, &cw); err != nil {
			return
		}
		// iterate config entries
		for i, c := range cw.Configs {
			// skip unless enabled
			if !c.Enabled {
				log.Println("skipping disabled config in file", filename, i)
				continue
			}
			// skip invalid keep number
			if c.Keep < 0 {
				log.Println("invalid field 'keep' in file", filename, i)
				continue
			}
			// default keep 7 days
			if c.Keep == 0 {
				c.Keep = 7
			}
			// skip empty paths
			if len(c.Paths) == 0 {
				log.Println("empty field 'paths' in file", filename, i)
			}
			// append ret
			ret = append(ret, c)
		}
	}

	return
}
