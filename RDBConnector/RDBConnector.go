
// RDBConnector package implements relational database connector.
// Based on gorm library, with PostgreSQL supported as default.
package RDBConnector

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	// _ "github.com/jinzhu/gorm/dialects/mssql"
)

// Default relational database configuration
var defaultRDBConfig RDBConfig

// Options for user-specified relational database configuration
type RDBConfigOption func(config *RDBConfig) error

// Data type for relational database configuration
type RDBConfig struct {
    DriverName string 
    DSN string 
    ConnMaxLifetime time.Duration
    MaxIdleConns int
    MaxOpenConns int
}

// Relational database connector that wraps *gorm.DB
type RDBConnector struct {
	config *RDBConfig
	GormDB *gorm.DB
}

// Create relational database connector with configuration
func CreateRDBConnector(options ...RDBConfigOption) *RDBConnector {
	// Get relational database configuration
    config := createRDBConfig(options...)

    // Connect to relational database based on gorm library
    db, err := gorm.Open(config.DriverName, config.DSN)
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

// Create relational database configuration according to user's options,
// or return the default values in case of no options input
func createRDBConfig(options ...RDBConfigOption) *RDBConfig {
	if len(options) == 0 {
		return &defaultRDBConfig
	}

	// Copy default coniguration parameters
	userRDBConfig := defaultRDBConfig
    for _, option := range options {
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
		// dbname default as postgres
		DSN: "host=localhost port=5432 user=postgres password=QrfV2_Pg sslmode=disable",
		ConnMaxLifetime: time.Second * 10,
		MaxIdleConns: 100,
		MaxOpenConns: 1000,
	}
}