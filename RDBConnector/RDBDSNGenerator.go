// This file in charge of generating DSN string,
// It's great convenience to create a data source name.
package RDBConnector 

import (
	"strings"
)

// Relational database data source name generator.
// JSON tags here assigned for configuration from JSON file.
// It depends on users' choice to use literals or JSON file.
// In case of configuration from JSON file, please resort to the unified "RDBConfigFuncWrap.go".
type RDBDSNGenerator struct {
	Host string `json:"host"`
	Port string `json:"port"`
	User string `json:"user"`
	Password string `json:"password"`
	SSLMode string `json:"sslmode"`
	DBName string `json:"dbname"`
	Optionals map[string]string `json:"optionals, omitempty"` // Intendedly left for user-specified choices.
}

// Generate DSN string to the RDBDSNGenerator.
// The empty parametrs will get ignored.
func (r *RDBDSNGenerator) String() string {
	var dsn string 
	if r.Host != "" {
		dsn += "host=" + r.Host + " " 
	} 
	if r.Port != "" {
		dsn += "port=" + r.Port + " "
	}
	if r.User != "" {
		dsn += "user=" + r.User + " "
	}
	if r.Password != "" {
        dsn += "password=" + r.Password + " "	
	}
	if r.SSLMode != "" {
		dsn += "sslmode=" + r.SSLMode + " "
	}
	if r.DBName != "" {
         dsn += "dbname=" + r.DBName + " "
	}
	if r.Optionals != nil {
        for key, value := range r.Optionals {
			if key != "" && value != "" {
				dsn += key + "=" + value + " "
			}
		}
	}
	return strings.TrimSpace(dsn)
}