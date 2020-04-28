package sql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"sync"
)

var (
	dbPool sync.Pool
)

var (
	dbType = "sqlite3"
	args   = []interface{}{
		"test.db",
	}
)

func initPool() {
	dbPool = sync.Pool{
		New: createDBConnect,
	}
}

func createDBConnect() interface{} {
	db, err := gorm.Open(dbType, args...)
	if err != nil {
		log.Printf("Error: Create DB connect error: %v\n", err)
		for i := 0; i < 3; i++ {
			db, err = gorm.Open(dbType, args...)
			if err == nil {
				return db
			}
			log.Printf("Error: Create DB connect error: %v, Retry: %d\n", err, i+1)
		}
		return nil
	}
	return db
}

func GetDBConnect() (db *gorm.DB, err error) {
	var ok bool
	if db, ok = dbPool.Get().(*gorm.DB); !ok {
		return nil, fmt.Errorf("get db connection error, pool may be dirty")
	}
	if db == nil {
		return nil, fmt.Errorf("get db connection error")
	}
	return db, nil
}

func PutDBConnect(db *gorm.DB) {
	if db == nil {
		return
	}
	dbPool.Put(db)
}

func init() {
	initPool()
	models := []interface{}{
		&User{},
		&Project{},
		&Entry{},
	}

	db, err := GetDBConnect()
	defer PutDBConnect(db)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(models...)
}

func createTestInitData() {
	db, err := GetDBConnect()
	defer PutDBConnect(db)
	if err != nil {
		panic(err)
	}

	e1 := Entry{
		Key:   "abc",
		Value: "def",
	}

	e2 := Entry{
		Key:   "k2",
		Value: "{\"k1\": 1}",
		Type:  "json",
	}

	u0 := User{
		Name:     "daisy",
		Password: "3306",
	}

	p := Project{
		Name:      "prj name",
		AccessKey: "0001",
		Operators: []User{u0},
		Entries:   []Entry{e1, e2},
	}

	u := User{
		Name:     "Tobias",
		Password: "abcd",
		Projects: []Project{p},
	}
	db.Create(&u)
}
