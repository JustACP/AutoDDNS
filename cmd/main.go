package main

import (
	"github.com/JustACP/AutoDDNS/cmd/conf"
	"github.com/JustACP/AutoDDNS/cmd/content"
	"github.com/JustACP/AutoDDNS/logging"
)

func main() {

	logging.Info("DDNS is running")
	conf.ApplyConf()
	for range content.Ticker.C {
		conf.RefreshLocalValue(content.ServerRecordCache)
		for _, job := range content.ServerRecordCache {
			job.DoResolve()
		}
	}
}
