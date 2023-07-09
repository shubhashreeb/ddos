package store

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

type DbStore struct {
	db *sqlx.DB
}

func NewDbStore() *DbStore {
	dbConfig := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		"localhost", "5432", "postgres", "default", "postgres")

	newDb, err := sqlx.Connect("postgres", dbConfig)

	if err != nil {
		fmt.Println("Failed to connect to database", err)
		panic(err)
	}
	return &DbStore{
		db: newDb,
	}
}

func (store *DbStore) CreateDdos(req DdosConfigReq) (*DdosConfig, error) {
	uuid := uuid.NewV4()
	dbQuery := `INSERT INTO ddos_db (uuid, url, number_requests, duration) VALUES ($1, $2, $3, $4)`
	_, err := store.db.Exec(dbQuery, uuid, req.Url, req.NumberRequests, req.Duration)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &DdosConfig{
		Uuid:           uuid.String(),
		Url:            req.Url,
		NumberRequests: req.NumberRequests,
		Duration:       req.Duration,
	}, nil
}

func (store *DbStore) UpdateDdos(req DdosConfigReq, uuid string) (*DdosConfig, error) {
	dbQuery := `UPDATE ddos_db SET url = $2, number_requests = $3, duration = $4 WHERE uuid = $1;`
	r, err := store.db.Exec(dbQuery, uuid, req.Url, req.NumberRequests, req.Duration)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("result is", r)
	return &DdosConfig{
		Uuid:           uuid,
		Url:            req.Url,
		NumberRequests: req.NumberRequests,
		Duration:       req.Duration,
	}, nil
}

func (store *DbStore) GetDdos(uuid string) (*DdosConfig, error) {
	var result []DdosConfig
	result = make([]DdosConfig, 0)
	dbQuery := `SELECT uuid, url, number_requests, duration FROM ddos_db WHERE uuid = $1;`
	err := store.db.Select(&result, dbQuery, uuid)
	if err != nil {
		log.Println("Error in executing db query", err)
		return nil, err
	}
	log.Println("Here is the request details", result, err)
	return &result[0], err
}

func (store *DbStore) Delete(uuid string) error {
	sqlStatement := `DELETE FROM ddos_db WHERE uuid = $1;`
	_, err := store.db.Exec(sqlStatement, uuid)
	if err != nil {
		panic(err)
	}
	return nil
}
