package RDBConnector

import (
	"fmt"
	"errors"
	"testing"
	"github.com/jinzhu/gorm"
)

func Test_CreateRDBConnector(t *testing.T) {
	// Test happy path
	t.Run("happy path", func(t *testing.T){
		t.Run("use default config", func(t *testing.T){
			c := CreateRDBConnector()
			if c.config.DriverName != "postgres" ||  c.GormDB == nil {
				t.Error("Failed creating RDBConnector")
				return
			}
			stats := c.GormDB.DB().Stats()
			fmt.Printf("%#v", stats)

			type Product struct {
				gorm.Model
				Name string
			}
			c.GormDB = c.GormDB.AutoMigrate(&Product{})
			if c.GormDB.Error != nil {
				t.Error(c.GormDB.Error)
			}
		})

		t.Run("switch to another existing database",func(t *testing.T){
			// testDB database should be created beforehand
			newDSN := "host=localhost port=5432 user=postgres password=QrfV2_Pg sslmode=disable dbname=testDB"
			switchDB := func(config *RDBConfig) error {
				config.DSNString = newDSN
				return nil
			}

			c := CreateRDBConnector(switchDB)
			if c.config.DSNString != newDSN || c.GormDB == nil {
				t.Error("Failed creating RDBConnector")
				return
			} 

			type Product struct {
				gorm.Model
				Name string
			}
			c.GormDB = c.GormDB.CreateTable(&Product{})
			if c.GormDB.Error != nil {
				t.Error("Failed creating table. Err: ", c.GormDB.Error)
			}
		})
	})

	// Test input errors
	t.Run("input errors", func(t *testing.T){
		t.Run("driver name error", func(t *testing.T){
			driverError := func(config *RDBConfig) error {
				config.DriverName = "wrongDriverName"
				return nil
			}
			CreateRDBConnector(driverError)
		})

		t.Run("driver name error 2", func(t *testing.T){
			driverError := func(config *RDBConfig) error {
				return errors.New("Wrong driver name")
			}
			CreateRDBConnector(driverError)
		})
	})

	// Test dependency issues
	t.Run("dependency issues", func(t *testing.T){
		t.Run("database not ready", func(t *testing.T){
			// testDB2 database does not be created beforehand
			newDSN := "host=localhost port=5432 user=postgres password=QrfV2_Pg sslmode=disable dbname=testDB2"
			switchDB := func(config *RDBConfig) error {
				config.DSNString = newDSN
				return nil
			}

			c := CreateRDBConnector(switchDB)
			if c.config.DSNString != newDSN || c.GormDB == nil {
				t.Error("Failed creating RDBConnector")
				return
			} 
		})
	})
}