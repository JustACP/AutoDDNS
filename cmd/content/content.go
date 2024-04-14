package content

import (
	"github.com/JustACP/AutoDDNS/logging"
	"github.com/JustACP/AutoDDNS/pkg/cloud"
	"time"
)

var (
	ServerRecordCache map[string]*ResolveJob = make(map[string]*ResolveJob)
	CloudClientCache                         = map[string]cloud.DNSServer{
		"aliyun": &cloud.AliyunServer{},
	}
	Ticker *time.Ticker
)

type ResolveJob struct {
	Server cloud.DNSServer
	Record *cloud.ServerRecord
}

func (s *ResolveJob) InitResolveJob(record *cloud.ServerRecord) {
	s.Record = record
}

func (s *ResolveJob) DoResolve() {
	if s.Record == nil {
		return
	}

	newRecord, err := s.Server.GetDomain(s.Record.FullDomain)
	if err != nil {
		logging.Warn("get record domain %s error %v", s.Record.FullDomain, err)
		return
	}

	if newRecord == nil {
		err := s.Server.CreateDomainRecord(s.Record)
		if err != nil {
			logging.Warn("create record domain %s error %v", s.Record.FullDomain, err)
			return
		}
		logging.Info("create record domain %s value %s success", s.Record.FullDomain, s.Record.Value)
	}

	if s.Record.Value != newRecord.Value {
		err := s.Server.UpdateDomainRecord(s.Record)
		if err != nil {
			s.Record.RefreshRecordFromRemote(newRecord)
			logging.Warn("do update record domain %s error %v", s.Record.FullDomain, err)
		}
		logging.Info("update record domain %s value %s success", s.Record.FullDomain, s.Record.Value)
		return
	}
	logging.Info("domain %s value %s no change", s.Record.FullDomain)

}
