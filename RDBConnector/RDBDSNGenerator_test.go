package RDBConnector 

import (
	"testing"
)

func Test_RDBDSNGenerator(t *testing.T) {
	t.Run("String", func(t *testing.T){
		// Test happy path
		t.Run("happy path", func(t *testing.T){
			testObj := &RDBDSNGenerator{
				Host: "localhost",
				Port: "5432",
				User: "ljg_cqu",
				Password: "139850", 
				SSLMode: "disable",
				DBName: "testDB",
				Optionals: map[string]string{
					"optionalKey": "optionalValue",
				},
			}

			dsn := testObj.String()
			if dsn != "host=localhost port=8000 user=ljg_cqu " + 
			"password=139850 sslmode=disable dbname=testDB optionalKey=optionalValue" {
                t.Error("Failed generating DSN.", dsn)
			}
		})

		// Test input errors
		t.Run("input errors", func(t *testing.T){
			testObj := &RDBDSNGenerator{}
			dsn := testObj.String()
			if dsn != "" {
				t.Error("DSN not empty string.")
			}
		})
	})
}