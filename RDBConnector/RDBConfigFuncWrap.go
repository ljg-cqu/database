// This file aims to wrap RDBConfigFunc from variant place
// It's a convient tool to extend funtionality of RDBConfigFunc.
package RDBConnector 

import (
    "fmt"
	"os"
	"encoding/json"
)

// SetDriverName updates DriverName only.
// If driverName empty, nothing will happen.
func SetDriverName(driverName string) RDBConfigFunc {
	return func(config *RDBConfig) error {
		if driverName != "" {
			config.DriverName = driverName
		}
		return nil
	}
}

// SetDSNString updates DSNString only.
// If dsnString empty, nothing will happen.
func SetDSNString(dsnString string) RDBConfigFunc {
	return func(config *RDBConfig) error {
		if dsnString != "" {
			config.DSNString = dsnString
		}
		return nil
	}
}

// SetDSNStringWithRDBDSNGenerator updates DSNString only.
// If rdbDSNGenerator nil, nothing will happen.
func SetDSNStringWithRDBDSNGenerator(rdbDSNGenerator *RDBDSNGenerator) RDBConfigFunc {
	return func(config *RDBConfig) error {
		if rdbDSNGenerator != nil {
			config.DSNString = rdbDSNGenerator.String()
		}
		return nil
	}
}

// FromConfigObj updates RDBConfig according to userConfig.
// Note DSNString field here must generated against RDBDSNGenerator.
// If users provide both DSNString and RDBDSNGenerator concurrently,
// then DSNString will take priority over RDBDSNGenerator.
func FromConfigObj(userConfig *RDBConfig) RDBConfigFunc {
	return func(config *RDBConfig) error {
        if userConfig == nil {
			return nil
		}
		if userConfig.DriverName != "" {
			config.DriverName = userConfig.DriverName
		}
		if userConfig.DSNString != "" {
			config.DSNString = userConfig.DSNString
		} else if userConfig.DSN != nil {
			config.DSNString = userConfig.DSN.String()
		}
		if userConfig.ConnMaxLifetime != 0 {
			config.ConnMaxLifetime = userConfig.ConnMaxLifetime
		}
		if userConfig.MaxIdleConns != 0 {
			config.MaxIdleConns = userConfig.MaxIdleConns
		}
		if userConfig.MaxOpenConns != 0 {
			config.MaxOpenConns = userConfig.MaxOpenConns
		}
		if userConfig.Optionals != nil {
			for key, value := range userConfig.Optionals {
				config.Optionals[key] = value
			}
		}
		return nil
	} 
}

// FromEnvVar use environmental key and its value as DriverName and DSN respectively.
func FromEnvVar(key string) RDBConfigFunc {
	return func(config *RDBConfig) error {
		dsn, ok :=os.LookupEnv(key)
		if !ok {
			return fmt.Errorf("LookupEnv for DriverName \"%s\"failed.", key)
		}
		config.DriverName = key 
		config.DSNString = dsn
		return nil
	}
}

// FromJSONFile overwrite RDBConfig from JSON file.
// Note DSNString takes priority over RDBDSNGenerator
// That is to say RDBDSNGenerator works only when DSNString is empty.
func FromJSONFile(path string) RDBConfigFunc {
	return func(config *RDBConfig) error {
		// Open coniguration JSON file 
		configFile, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("Failed to open file \"%s\"", path)
		}
		defer configFile.Close()

		// Parse JSON values into config object
		if json.NewDecoder(configFile).Decode(config); err != nil {
			return fmt.Errorf("Fail parse JSON values from file  \"%s\"", path)
		}

		// Generate DSNString, the final DSN value for DB connection
		if config.DSNString == "" {
			config.DSNString = config.DSN.String()
		}
		return nil
	}
}