package main

import (
	"testing"
	"time"
)

func TestLoadConfigs(t *testing.T) {
	var cfgs []Config
	var err error
	if cfgs, err = LoadConfigs("testdata/config.d/*.yml"); err != nil {
		t.Fatal(err)
	}
	if len(cfgs) != 1 {
		t.Fatal("failed")
	}
}

func TestConfigListFiles(t *testing.T) {
	var cfgs []Config
	cfgs, _ = LoadConfigs("testdata/config.yml")
	files, err := cfgs[0].ListFiles()
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 6 {
		t.Fatal("failed")
	}
}

func TestConfigIsExpired(t *testing.T) {
	n := time.Date(2018, time.July, 23, 0, 0, 0, 0, time.UTC)
	cfg := Config{Keep: 10}
	ft := time.Date(2018, time.July, 24, 0, 0, 0, 0, time.UTC)
	if cfg.IsExpired(n, ft) {
		t.Fatal("failed")
	}
	ft = time.Date(2018, time.July, 23, 0, 0, 0, 0, time.UTC)
	if cfg.IsExpired(n, ft) {
		t.Fatal("failed")
	}
	ft = time.Date(2018, time.July, 13, 0, 0, 0, 0, time.UTC)
	if cfg.IsExpired(n, ft) {
		t.Fatal("failed")
	}
	ft = time.Date(2018, time.July, 12, 0, 0, 0, 0, time.UTC)
	if !cfg.IsExpired(n, ft) {
		t.Fatal("failed")
	}
}
