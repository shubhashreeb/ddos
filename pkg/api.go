package pkg

import (
	"context"
	"fmt"

	ddospb "github.com/cloverway/schema/pbgo/v2/ddos"
	"github.com/shubhashreeb/ddos/store"
)

type DdosServer struct {
	dbstore store.Store
	ddospb.UnimplementedDdoSServiceServer
}

func NewDdosServer() *DdosServer {
	return &DdosServer{
		dbstore: store.NewDbStore(),
	}
}

func (d *DdosServer) CreateDDos(ctx context.Context, req *ddospb.CreateRequest) (*ddospb.CreateHttpResponse, error) {
	fmt.Println("CreateDDos is invoked")
	res, err := d.dbstore.CreateDdos(store.DdosConfigReq{
		Url:            req.Url,
		NumberRequests: req.NumberRequests,
		Duration:       req.Duration,
	})
	if err != nil {
		fmt.Println("Error in creating db entry", err)
		return nil, err
	}
	return &ddospb.CreateHttpResponse{
		Uuid: res.Uuid,
		Url:  res.Url,
	}, nil
}
