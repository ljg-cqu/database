
// RDBConnector package implements relational database connector.
// Based on gorm library, with PostgreSQL supported as default.
// Users can choose to use another RDB with appropriate configurations.
// Besides update import package for database driver accordingly.
package RDBConnector

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	// _ "github.com/jinzhu/gorm/dialects/mssql"
)

// Default relational database configuration
var defaultRDBConfig RDBConfig

// Optionals for user-specified relational database configuration
type RDBConfigFunc func(config *RDBConfig) error

// Data type for relational database configuration
type RDBConfig struct {
    DriverName string `json:"driver_name"` 
	DSN *RDBDSNGenerator `json:"dsn"`
	DSNString string `json:"dsn_string"`
    ConnMaxLifetime time.Duration `json:"conn_max_lifetime"`
    MaxIdleConns int `json:"max_idle_conns"`
	MaxOpenConns int `json:"max_open_conns"`
	Optionals map[string]interface{} `json:"optionals, omitempty"` // Intendedly left for user-specified choices.
}

// Relational database connector that wraps *gorm.DB
type RDBConnector struct {
	config *RDBConfig
	GormDB *gorm.DB
}

// Create relational database connector with configuration
func CreateRDBConnector(Optionals ...RDBConfigFunc) *RDBConnector {
	// Get relational database configuration
    config := createRDBConfig(Optionals...)

    // Connect to relational database based on gorm library
    db, err := gorm.Open(config.DriverName, config.DSNString)
    if err != nil {
    	panic(err)
	}

	// Set connection parameters via generic database interface 
	DB := db.DB()
	DB.SetConnMaxLifetime(config.ConnMaxLifetime)
	DB.SetMaxIdleConns(config.MaxIdleConns)
	DB.SetMaxOpenConns(config.MaxOpenConns)

    return &RDBConnector{config, db}
}

// Create relational database configuration according to user's Optionals,
// or return the default values in case of no Optionals input
func createRDBConfig(Optionals ...RDBConfigFunc) *RDBConfig {
	if len(Optionals) == 0 {
		return &defaultRDBConfig
	}

	// Copy default coniguration parameters
	userRDBConfig := defaultRDBConfig
    for _, option := range Optionals {
    	if err := option(&userRDBConfig); err != nil {
    		panic(err)
		}
	}
    return &userRDBConfig
}

func init() {
	// Set default relational database configuration
	defaultRDBConfig = RDBConfig {
		DriverName: "postgres",
		// DSN: &RDBDSNGenerator{
		// 	Host: "localhost",
		// 	Port: "5432",
		// 	User: "postgres",
		// 	Password: "QrfV2_Pg", 
		// 	SSLMode: "disable",
		// 	DBName: "testDB",
		// },
		DSNString: "host=localhost port=5432 user=postgres password=QrfV2_Pg sslmode=disable dbname=testDB",
		ConnMaxLifetime: time.Second * 10,
		MaxIdleConns: 100,
		MaxOpenConns: 1000,
	}
}