package RDBConnector 

import (
	"testing"
)

func Test_RDBConfigFuncWrap_FromOjb(t *testing.T) {
	// Test happy path 
	t.Run("nil userConfig bject", func(t *testing.T){
		fromObj := FromConfigObj(nil) 
		rdbConn := CreateRDBConnector(fromObj)

		driverNameGot := rdbConn.config.DriverName 
		driverNameWant := "postgres"
		if driverNameGot != driverNameWant  {
			t.Errorf("Got:%s, want: %s", driverNameGot, driverNameWant)
		}

		dsnGot := rdbConn.config.DSNString
		dsnWant := "host=localhost port=5432 user=postgres password=QrfV2_Pg sslmode=disable dbname=testDB"
		if dsnGot != dsnWant {
			t.Errorf("Got:%s, want: %s", dsnGot, dsnWant)
		}
	})

	t.Run("switch to sqlite", func(t *testing.T){
		sqliteDB := "C:/Workspace/Go/src/github.com/ljg_cqu/database/RDBConnector/testDB.db"
		myConfig := &RDBConfig{
			DriverName: "sqlite3", 
			DSN: &RDBDSNGenerator{
				Host: "localhost",
			},
			DSNString: sqliteDB,
		}

		fromObj := FromConfigObj(myConfig)
		rdbConn := CreateRDBConnector(fromObj)
		driverNameGot := rdbConn.config.DriverName 
		driverNameWant := "sqlite3"
		if driverNameGot != driverNameWant  {
			t.Errorf("Got:%s, want: %s", driverNameGot, driverNameWant)
		}

		dsnGot := rdbConn.config.DSNString
		dsnWant := sqliteDB
		if dsnGot != dsnWant {
			t.Errorf("Got:%s, want: %s", dsnGot, dsnWant)
		}
	})
}

func Test_RDBConfigFuncWrap_FromEnv(t *testing.T)  {
	// Test happy path
	t.Run("happy path", func(t *testing.T){
		  fromEnv := FromEnvVar("postgres")
		  rdbConn := CreateRDBConnector(fromEnv)
		  if rdbConn.config.DriverName != "postgres" || rdbConn.config.DSNString != "host=localhost port=5432 user=postgres password=QrfV2_Pg sslmode=disable dbname=testDB" {
			  t.Errorf("")
		  }
	})

	// Test input errors 
	t.Run("input errors", func(t *testing.T){
		fromEnv := FromEnvVar("postgres_invalid")
		rdbConn := CreateRDBConnector(fromEnv)
		if rdbConn.config.DriverName != "postgres" || rdbConn.config.DSNString != "host=localhost port=5432 user=postgres password=QrfV2_Pg sslmode=disable dbname=testDB" {
			t.Errorf("")
		}
	})
}

func Test_RDBConfigFuncWrap_FromFile_postgreSQL(t *testing.T)  {
	// Test happy path
	t.Run("happy path", func(t *testing.T){
		path := "C:/Workspace/Go/src/github.com/ljg_cqu/database/RDBConnector/RDBConfig_postgres.json"
		fromFile := FromJSONFile(path)
		rdbConn := CreateRDBConnector(fromFile)

		driverNameGot := rdbConn.config.DriverName 
		driverNameWant := "postgres"
		if driverNameGot != driverNameWant  {
			t.Errorf("dGot:%s, want: %s", driverNameGot, driverNameWant)
		}

		dsnGot := rdbConn.config.DSNString
		dsnWant := "host=localhost port=5432 user=postgres password=QrfV2_Pg sslmode=disable dbname=testDB"
		if dsnGot != dsnWant {
			t.Errorf("Got:%s, want: %s", dsnGot, dsnWant)
		}

		dsnOpGot := rdbConn.config.DSN.Optionals["optional_key"]
		dsnOpWant := ""
		if dsnOpGot != dsnOpWant {
			t.Errorf("Got: %s, want: %s", dsnOpGot, dsnOpWant)
		}

		confOpGot := rdbConn.config.Optionals["optional_key"] 
		confOpWant := "optional_value"
		if confOpGot !=  confOpWant {
			t.Errorf("Got: %s, want: %s", confOpGot, dsnOpWant)
		}
	})

	// Test input errors 
	t.Run("input errors ", func(t *testing.T){
		path := "C:/Workspace/Go/src/github.com/ljg_cqu/database/RDBConnector/RDBConfig_postgres_notexist.json"
		fromFile := FromJSONFile(path)
		_ = CreateRDBConnector(fromFile)
	})
}

func Test_RDBConfigFuncWrap_FromFile_sqlite(t *testing.T)  {
	// Test happy path
	t.Run("happy path", func(t *testing.T){
		path := "C:/Workspace/Go/src/github.com/ljg_cqu/database/RDBConnector/RDBConfig_sqlite.json"
		fromFile := FromJSONFile(path)
		rdbConn := CreateRDBConnector(fromFile)

		driverNameGot := rdbConn.config.DriverName 
		driverNameWant := "sqlite3"
		if driverNameGot != driverNameWant  {
			t.Errorf("dGot:%s, want: %s", driverNameGot, driverNameWant)
		}

		dsnGot := rdbConn.config.DSNString
		dsnWant := "C:/Workspace/Go/src/github.com/ljg_cqu/database/RDBConnector/testDB.db"
		if dsnGot != dsnWant {
			t.Errorf("Got:%s, want: %s", dsnGot, dsnWant)
		}

		confOpGot := rdbConn.config.Optionals["optional_key_sqlite"] 
		confOpWant := "optional_value_sqlite"
		if confOpGot !=  confOpWant {
			t.Errorf("Got: %s, want: %s", confOpGot, confOpWant)
		}
	})

	// Test input errors 
	t.Run("input errors ", func(t *testing.T){
		path := "C:/Workspace/Go/src/github.com/ljg_cqu/database/RDBConnector/RDBConfig_sqlite_notexist.json"
		fromFile := FromJSONFile(path)
		_ = CreateRDBConnector(fromFile)
	})
}