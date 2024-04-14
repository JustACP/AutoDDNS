package config

type Platform struct {
	Cloud           string   `json:"cloud"`
	AccessKeyId     string   `json:"access_key_id"`
	AccessKeySecret string   `json:"access_key_secret"`
	Endpoint        string   `json:"endpoint"`
	DnsRecords      []Record `json:"dns_record"`
}

type Record struct {
	Domain  string `json:"domain"`
	Type    string `json:"Type"`
	BindMAC string `json:"bind_mac"`
}

type Config struct {
	LogPath      string     `json:"log_path"`
	TickDuration int64      `json:"tick_duration"`
	Platforms    []Platform `json:"platforms"`
}

var record = Record{
	Domain:  "a.b.com",
	Type:    "ipv6",
	BindMAC: "",
}

var templateConfig = Config{
	LogPath:      "",
	TickDuration: int64(10),
	Platforms: []Platform{{
		Cloud:           "aliyun",
		AccessKeyId:     "key",
		AccessKeySecret: "secret",
		DnsRecords:      []Record{record},
	}},
}

func GetTemplateConfig() Config {
	return templateConfig
}
