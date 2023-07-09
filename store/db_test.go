package store

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDbStoreCrateGet(t *testing.T) {
	d := NewDbStore()

	//Create
	res, err := d.CreateDdos(DdosConfigReq{
		Url:            "www.google.com",
		NumberRequests: 20,
		Duration:       50,
	})
	if err != nil {
		fmt.Print("Error : ", err)
	}
	uuid := res.Uuid
	fmt.Println("UUid received -- ", res.Uuid)
	assert.NotNil(t, res)
	assert.NotNil(t, res.Uuid)
	assert.Equal(t, "www.google.com", res.Url)

	//Read
	res, err = d.GetDdos(res.Uuid)
	if err != nil {
		fmt.Print("Error in getting data : ", err)
	}
	assert.Equal(t, uuid, res.Uuid)

	//delete
	err = d.Delete(res.Uuid)
	if err != nil {
		fmt.Println("Error in deleting the data", err)
	}
}

func TestDbStoreUpdate(t *testing.T) {
	d := NewDbStore()

	//Create
	res, err := d.CreateDdos(DdosConfigReq{
		Url:            "www.google.com",
		NumberRequests: 20,
		Duration:       50,
	})
	if err != nil {
		fmt.Print("Error : ", err)
	}

	//Update
	_, err = d.UpdateDdos(DdosConfigReq{
		Url:            "www.google1.com",
		NumberRequests: 10,
		Duration:       5,
	}, res.Uuid)
	if err != nil {
		fmt.Print("Error in getting data : ", err)
	}

	res, err = d.GetDdos(res.Uuid)
	if err != nil {
		fmt.Print("Error in getting data : ", err)
	}
	fmt.Println("UUid received after update -- ", res)
	assert.Equal(t, int(res.NumberRequests), 10)
	assert.Equal(t, int(res.Duration), 5)

	//delete
	err = d.Delete(res.Uuid)
	if err != nil {
		fmt.Println("Error in deleting the data", err)
	}
}
