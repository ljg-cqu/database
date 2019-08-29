// This file present an comprehensive example to basic use of RDBConnector
package main

import (
	"fmt"
	
	"github.com/jinzhu/gorm"
	rdb "github.com/ljg_cqu/database/RDBConnector"
)

// Note normally it's ok with one approach to configure RDB connection.
// But here we show multiple ways in one place, just for not to omit important usages.
// Besides, users can only use default configurations, with postgreSQL and testDB required.
func main() {
	// Set DriverName 
	setDriverName_sqlite := rdb.SetDriverName("sqlite3")

	// Set DSNString 
	sqliteDBPath := "C:/Workspace/Go/src/github.com/ljg_cqu/database/examples/testDB.db"
	setDSNString_sqlite := rdb.SetDSNString(sqliteDBPath)

	// Set DSNString With RDBDSNGenerator
	dsnGen := &rdb.RDBDSNGenerator{
		Host: "localhost",
		Port: "5432",
		User: "postgres",
		Password: "QrfV2_Pg",
		SSLMode: "disable",
		DBName: "testDB",
	}
	setDSNStringWithGen := rdb.SetDSNStringWithRDBDSNGenerator(dsnGen)

	// Update configurations from environment 
	fromEnv := rdb.FromEnvVar("postgres")

	// Overite configurations from JSON file
	pgPath := "C:/Workspace/Go/src/github.com/ljg_cqu/database/examples/RDBConfig_postgres.json"
	fromFile := rdb.FromJSONFile(pgPath)

	// ----------------------------------------------------------------------------------------
	// Check whether it got connected successfully,
	// from gorm.DB to generic database/sql.DB
	// rdbConn := rdb.CreateRDBConnector(setDriverName_sqlite, setDSNString_sqlite, setDSNStringWithGen, fromEnv, fromFile)
	rdbConn := rdb.CreateRDBConnector(setDSNStringWithGen, fromEnv, fromFile, setDriverName_sqlite, setDSNString_sqlite)
	fmt.Println("RDB connector created.")
	fmt.Println(rdbConn)
	
	fmt.Println("Retrive generic database.sql.DB")
	db := rdbConn.GormDB.DB()
	fmt.Println(db)

	// Create test_tables via AutoMigrate 
	// besides insert one record.
	type TestTable struct {
		gorm.Model
		Name string `gorm:"NOT NULL"`
	}  
	rdbConn.GormDB = rdbConn.GormDB.AutoMigrate(&TestTable{})
	rdbConn.GormDB.Create(&TestTable{
		Name: "ljg_cqu",
	})
}