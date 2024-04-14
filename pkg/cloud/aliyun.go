package cloud

import (
	"errors"
	"github.com/JustACP/AutoDDNS/config"
	"github.com/JustACP/AutoDDNS/logging"
	"github.com/JustACP/AutoDDNS/pkg/util"
	"github.com/alibabacloud-go/alidns-20150109/v4/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"sync"
)

type AliyunServer struct {
	Client   *client.Client
	Platform *config.Platform
	once     sync.Once
}

func (s *AliyunServer) InitServer(platform *config.Platform) error {
	var err error = nil
	//s.once.Do(func() {
	newClient, err := client.NewClient(&openapi.Config{
		AccessKeyId:     &platform.AccessKeyId,
		AccessKeySecret: &platform.AccessKeySecret,
		Endpoint:        &platform.Endpoint,
	})
	if err != nil {
		logging.Warn("init aliyun server error, %v", err)
		return nil
	}
	s.Platform = platform
	s.Client = newClient

	err = nil
	//})
	return err
}

func (s *AliyunServer) GetDomain(domain string) (*ServerRecord, error) {
	request := &client.DescribeDomainRecordsRequest{
		DomainName: util.GetPtr[string](util.GetMainDomain(domain)),
		PageNumber: util.GetPtr[int64](1),
		PageSize:   util.GetPtr[int64](100),
		KeyWord:    util.GetPtr[string](util.GetPrefix(domain)),
		SearchMode: util.GetPtr[string]("EXACT"),
	}

	result, err := s.Client.DescribeDomainRecords(request)
	if err != nil {
		logging.Warn("aliyun get domain: %s error: %v", domain, err)
		return nil, err
	}

	var matchRecord *client.DescribeDomainRecordsResponseBodyDomainRecordsRecord
	for _, record := range result.Body.DomainRecords.Record {
		if *record.RR+"."+*record.DomainName == domain {
			matchRecord = record
		}
	}
	if matchRecord == nil {
		return nil, nil
	}
	var recordType ServerRecordType
	var exist bool
	if recordType, exist = RecordTypeMap[*matchRecord.Type]; !exist {
		logging.Warn("record type is not supported, current type: %s", recordType)
		return nil, errors.New("record type is not supported")
	}

	return &ServerRecord{
		MainDomain: util.GetMainDomain(domain),
		FullDomain: domain,
		Value:      *matchRecord.Value,
		RecordType: recordType,
		TTL:        *matchRecord.TTL,
		RecordId:   *matchRecord.RecordId,
		Status:     *matchRecord.Status,
	}, nil
}

func (s *AliyunServer) CreateDomainRecord(record *ServerRecord) error {
	var TTL int64 = 600
	if record.TTL <= 0 {
		record.TTL = TTL
	}
	request := &client.AddDomainRecordRequest{
		DomainName: &record.MainDomain,
		RR:         util.GetPtr[string](util.GetPrefix(record.FullDomain)),
		Type:       util.GetPtr[string](string(record.RecordType)),
		Value:      &record.Value,
		TTL:        &record.TTL,
	}
	result, err := s.Client.AddDomainRecord(request)
	if err != nil {
		logging.Warn("create domain record error, record:%v , err:%v", record, err)
		return err
	}
	record.RecordId = *result.Body.RecordId
	return nil

}

func (s *AliyunServer) UpdateDomainRecord(record *ServerRecord) error {
	var TTL int64 = 600
	if record.TTL <= 0 {
		record.TTL = TTL
	}
	request := &client.UpdateDomainRecordRequest{
		RR:       util.GetPtr[string](util.GetPrefix(record.FullDomain)),
		Value:    &record.Value,
		RecordId: &record.RecordId,
		TTL:      &record.TTL,
		Type:     util.GetPtr[string](string(record.RecordType)),
	}
	result, err := s.Client.UpdateDomainRecord(request)
	if err != nil {
		logging.Warn("update domain record error, record:%v , err:%v", record, err)
		return err
	}
	record.RecordId = *result.Body.RecordId
	return nil
}
