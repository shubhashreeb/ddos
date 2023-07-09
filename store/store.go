package store

type Store interface {
	CreateDdos(DdosConfigReq) (*DdosConfig, error)
	GetDdos(string) (*DdosConfig, error)
	UpdateDdos(DdosConfigReq, string) (*DdosConfig, error)
	Delete(string) error
}

type DdosConfigReq struct {
	Url            string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	NumberRequests int64  `protobuf:"varint,2,opt,name=number_requests,json=number_requests,proto3" json:"number_requests,omitempty"`
	Duration       int64  `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
}

type DdosConfig struct {
	Uuid           string `db:"uuid,omitempty"`
	Url            string `db:"url"`
	NumberRequests int64  `db:"number_requests"`
	Duration       int64  `db:"duration" `
}
