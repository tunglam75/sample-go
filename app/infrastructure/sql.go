package infrastructure

import (
	"os"
	"sample/app/shared/utils"

	"github.com/jinzhu/gorm"
	// blank import.
	_ "github.com/lib/pq"
)

const (
	// DBMaster set master database string.
	DBMaster = "master"
	// DBRead set read replica database string.
	DBRead = "read"
	// DBMS postgres type
	DBMS = "postgres"
)

// SQL struct.
type SQL struct {
	// Master connections master database.
	Master *gorm.DB
	// Read connections read replica database.
	Read *gorm.DB
}

type dbInfo struct {
	host    string
	user    string
	pass    string
	name    string
	port    string
	logmode bool
}

// NewSQL returns new SQL.
// NewSQL returns new SQL.
func NewSQL() (*SQL, error) {
	info := map[string]dbInfo{}
	info[DBMaster] = dbInfo{
		host: os.Getenv("DB_MASTER_HOST"),
		user: os.Getenv("DB_MASTER_USER"),
		pass: os.Getenv("DB_MASTER_PASSWORD"),
		name: os.Getenv("DB_NAME"),
		port: os.Getenv("DB_PORT"),
		// logmode: GetenvBool("DB_MASTER_LOG_MODE"),
	}
	info[DBRead] = dbInfo{
		host: os.Getenv("DB_READ_HOST"),
		user: os.Getenv("DB_READ_USER"),
		pass: os.Getenv("DB_READ_PASSWORD"),
		name: os.Getenv("DB_NAME"),
		port: os.Getenv("DB_PORT"),
		// logmode: GetenvBool("DB_READ_LOG_MODE"),
	}
	var master, read *gorm.DB

	for i, v := range info {
		connect := "host=" + v.host + " port=" + v.port + " user=" + v.user + " dbname=" + v.name + " sslmode=disable password=" + v.pass
		db, err := gorm.Open(DBMS, connect)
		if err != nil {
			return nil, utils.ErrorsWrap(err, "can't open database")
		}
		db.LogMode(v.logmode)
		// Disable table name's pluralization globally
		// if set this to true, `User`'s default table name will be `user`, table name setted with `TableName` won't be affected
		db.SingularTable(true)

		if i == DBMaster {
			master = db
		} else if i == DBRead {
			// can't create/update/delete read replica database.
			// db.Callback().Create().Before("gorm:create").Register("read_create", database.CreateErrorCallback)
			// db.Callback().Update().Before("gorm:update").Register("read_update", database.UpdateErrorCallback)
			// db.Callback().Delete().Before("gorm:delete").Register("read_delete", database.DeleteErrorCallback)
			read = db
		}
	}
	return &SQL{master, read}, nil
}

// CloseSQL close DB
func CloseSQL(db *gorm.DB) error {
	if db != nil {
		err := db.Close()
		if err != nil {
			return utils.ErrorsWrap(err, "can't close db")
		}
	}
	return nil
}
