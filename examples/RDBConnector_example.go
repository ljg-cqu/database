// This file present an comprehensive example to basic use of RDBConnector
package main

import (
	"fmt"
	"time"
	
	"github.com/jinzhu/gorm"
	conn "github.com/ljg_cqu/database/RDBConnector"
)

// Struct types for illustration.
type Comment struct {
	ID uint
	Author string
	Content string
	PostID uint
	CreateAt time.Time
}

type Post struct {
	ID uint 
	Author string
	Comments []Comment 
	CreateAt time.Time
} 

type TestTable struct {
	gorm.Model
	Name string `gorm:"NOT NULL"`
}  

// Note normally it's ok with one approach to configure RDB connection.
// But here we show multiple ways in one place, just for not to omit important usages.
// Besides, users can only use default configurations, with postgreSQL and testDB required.
func main() {
	// Set DriverName 
	setDriverName_sqlite := conn.SetDriverName("sqlite3")

	// Set DSNString 
	sqliteDBPath := "C:/Workspace/Go/src/github.com/ljg_cqu/database/examples/testDB.db"
	setDSNString_sqlite := conn.SetDSNString(sqliteDBPath)

	// Set DSNString With RDBDSNGenerator
	dsnGen := &conn.RDBDSNGenerator{
		Host: "localhost",
		Port: "5432",
		User: "postgres",
		Password: "QrfV2_Pg",
		SSLMode: "disable",
		DBName: "testDB",
	}
	setDSNStringWithGen := conn.SetDSNStringWithRDBDSNGenerator(dsnGen)

	// Update configurations from environment 
	fromEnv := conn.FromEnvVar("postgres")

	// Overite configurations from JSON file
	pgPath := "C:/Workspace/Go/src/github.com/ljg_cqu/database/examples/RDBConfig_postgres.json"
	fromFile := conn.FromJSONFile(pgPath)

	// ----------------------------------------------------------------------------------------
	// Check whether it got connected successfully,
	// from gorm.DB to generic database/sql.DB
	// rdbConn := conn.CreateRDBConnector(setDriverName_sqlite, setDSNString_sqlite, setDSNStringWithGen, fromEnv, fromFile)
	rdbConn := conn.CreateRDBConnector(setDSNStringWithGen, fromEnv, fromFile, setDriverName_sqlite, setDSNString_sqlite)
	fmt.Println("RDB connector created.")
	fmt.Println(rdbConn)
	
	fmt.Println("Retrive generic database.sql.DB")
	db := rdbConn.GormDB.DB()
	fmt.Println(db)

	// Create test_tables via AutoMigrate 
	// besides insert one record.
	rdbConn.GormDB = rdbConn.GormDB.AutoMigrate(&TestTable{})
	rdbConn.GormDB.Create(&TestTable{
		Name: "ljg_cqu",
	})

	// // Create posts and comments tables with an foreignkey added.
	gormDB := rdbConn.GormDB
	gormDB.AutoMigrate(&Post{})
	gormDB.AutoMigrate(&Post{} ,&Comment{}).AddForeignKey("post_id", "posts(id)", "RESTRICT", "RESTRICT")	
	if gormDB.Error != nil {
		panic(gormDB.Error)
	}

	gormDB.Close()
}