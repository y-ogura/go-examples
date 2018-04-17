package main

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type repository struct {
	Conn *gorm.DB
}

// NewRepository mount repository
func NewRepository(db *gorm.DB) Repository {
	return &repository{
		Conn: db,
	}
}

// Repository repository interface
type Repository interface {
	Update(model interface{}, param map[string]interface{}) error
	NullTimeUpdate(model *NullTime) error
}

// PointerTime struct
type PointerTime struct {
	Time  *time.Time
	Valid bool
}

// NullTime pq.NullTime struct
type NullTime struct {
	ID       int
	JoinDate pq.NullTime
}

// NullPointer pointer struct
type NullPointer struct {
	ID       int
	JoinDate *time.Time
}

// NullPointerTime null pointer time struct
type NullPointerTime struct {
	ID       int
	JoinDate NullPointer
}

func main() {
	db, _ := stubDB("mock db")
	defer db.Close()
	db.LogMode(true)
	r := NewRepository(db)

	nullTime := &NullTime{}
	param1 := map[string]interface{}{"id": 1, "join_date": pq.NullTime{Time: time.Now(), Valid: false}}
	r.Update(nullTime, param1)

	nullPointer := &NullPointer{}
	param2 := map[string]interface{}{"id": 2, "join_date": pq.NullTime{Time: time.Now(), Valid: false}}
	r.Update(nullPointer, param2)

	pointerTime := &NullPointerTime{}
	param3 := map[string]interface{}{"id": 3, "join_date": PointerTime{Time: nil, Valid: true}}
	r.Update(pointerTime, param3)

	param4 := map[string]interface{}{"id": 4, "join_date": PointerTime{Time: nil, Valid: false}}
	r.Update(pointerTime, param4)

	now := time.Now()
	param5 := map[string]interface{}{"id": 5, "join_date": PointerTime{Time: &now, Valid: true}}
	r.Update(pointerTime, param5)

	param6 := NullTime{
		ID: 6,
		JoinDate: pq.NullTime{
			Valid: false,
		},
	}
	r.NullTimeUpdate(&param6)

	return
}

func stubDB(dsn string) (*gorm.DB, sqlmock.Sqlmock) {
	var db *gorm.DB
	_, mock, err := sqlmock.NewWithDSN(dsn)
	if err != nil {
		panic("Got an unexpected error.")
	}
	db, err = gorm.Open("sqlmock", dsn)
	if err != nil {
		panic("Got an unexpected error.")
	}
	return db, mock
}

func (m *repository) Update(model interface{}, param map[string]interface{}) error {
	err := m.Conn.Model(model).Updates(param).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *repository) NullTimeUpdate(model *NullTime) error {
	err := m.Conn.Model(model).Updates(model).Error
	if err != nil {
		return err
	}
	return nil
}
