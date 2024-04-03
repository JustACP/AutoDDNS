package cloud

import "github.com/JustACP/AutoDDNS/config"

type ServerRecordType string

const (
	RecordType_MX    ServerRecordType = "MX"
	RecordType_A     ServerRecordType = "A"
	RecordType_AAAA  ServerRecordType = "AAAA"
	RecordType_CNAME ServerRecordType = "CNAME"
)

var (
	RecordTypeMap = map[string]ServerRecordType{
		string(RecordType_AAAA):  RecordType_AAAA,
		string(RecordType_A):     RecordType_A,
		string(RecordType_CNAME): RecordType_CNAME,
		string(RecordType_MX):    RecordType_MX,
	}
)

type DNSServer interface {
	InitServer(server *config.Platform) error
	GetDomain(domain string) (*ServerRecord, error)
	CreateDomainRecord(record *ServerRecord) error
	UpdateDomainRecord(record *ServerRecord) error
}

type ServerRecord struct {
	RecordId   string
	MainDomain string
	FullDomain string
	Value      string
	RecordType ServerRecordType
	TTL        int64
	Status     string
	MAC        string
}

func (s *ServerRecord) RefreshRecordFromRemote(remoteRecord *ServerRecord) {
	s.RecordId = remoteRecord.RecordId
	s.Value = remoteRecord.Value
	s.TTL = remoteRecord.TTL
	s.Status = remoteRecord.Status
}
