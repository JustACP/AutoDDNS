package conf

import (
	"bufio"
	"github.com/JustACP/AutoDDNS/cmd/content"
	"github.com/JustACP/AutoDDNS/config"
	"github.com/JustACP/AutoDDNS/logging"
	"github.com/JustACP/AutoDDNS/pkg/cloud"
	"github.com/JustACP/AutoDDNS/pkg/util"
	"io"
	"log"
	"os"
	"time"
)

func ApplyConf() {

	conf := LoadConfig()
	setLogOutput(conf)
	initServerAndCache(conf)
	content.Ticker = time.NewTicker(time.Second * time.Duration(conf.TickDuration))

}

func setLogOutput(conf *config.Config) {

	if conf.LogPath == "std" {
		log.SetOutput(os.Stdout)
		logging.Info("set log output as %s success", conf.LogPath)
		return
	}
	logging.Info("set log output as %s success", conf.LogPath)
	log.SetOutput(GetOrCreateLog(conf))
}

func GetOrCreateLog(conf *config.Config) io.Writer {
	logFile, err := os.OpenFile(conf.LogPath, os.O_APPEND|os.O_RDWR, 0664)
	if err != nil {
		logging.Error("create log error, %v", err)
	}
	return bufio.NewWriter(logFile)
}

func initServerAndCache(conf *config.Config) {

	for _, platform := range conf.Platforms {
		dns := content.CloudClientCache[platform.Cloud]
		err := dns.InitServer(&platform)

		if err != nil {
			logging.Error("init %s dns server request client error: %v", platform.Cloud, err)
		}

		for _, record := range platform.DnsRecords {
			job := &content.ResolveJob{
				Server: dns,
				Record: &cloud.ServerRecord{
					RecordId:   "",
					MainDomain: util.GetMainDomain(record.Domain),
					FullDomain: record.Domain,
					Value:      "",
					RecordType: cloud.RecordTypeMap[record.Type],
					MAC:        record.BindMAC,
					TTL:        600,
					Status:     "",
				},
			}
			content.ServerRecordCache[record.Domain] = job
			logging.Info("platform %s domain %s ddns job has been load", platform.Cloud, record.Domain)

		}

	}

	RefreshLocalValue(content.ServerRecordCache)
}

func RefreshLocalValue(iter map[string]*content.ResolveJob) {
	util.GetLocalAddr()

	for _, job := range iter {
		ips := util.LocalAddress[job.Record.MAC]
		var v4, v6 *string

		for _, ip := range ips {

			if ip.To4() != nil && v4 == nil {
				v4 = util.GetPtr[string](ip.To4().String())
			} else if ip.To16() != nil && v6 == nil {
				v6 = util.GetPtr[string](ip.To16().String())
			}
			if v4 != nil && v6 != nil {
				break
			}
		}

		switch job.Record.RecordType {
		case cloud.RecordType_AAAA:
			job.Record.Value = *v6
		case cloud.RecordType_A:
			job.Record.Value = *v4
		default:
		}
	}
}
