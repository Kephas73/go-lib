package sql_client

import (
    "fmt"
    "github.com/Kephas73/go-lib/logger"
    "sync"
)

// sqlClientManager type;
type sqlClientManager struct {
    sqlClients sync.Map
}

var sqlClientsManagerInstance = &sqlClientManager{}

// default value env key is "MySQL";
// if configKeys was set, key env will be first value (not empty) of this;
func InstallSQLClientManager(configKeys ...string) {
    getConfigFromEnv(configKeys...)

    for _, config := range configs {
        client := NewSqlxDB(config)
        if client == nil {
            err := fmt.Errorf("InstallSQLClientsManager - NewSqlxDB {%v} error", config)
            logger.Error("InstallSQLClientsManager - Error: %v", err)

            panic(err)
        }

        if config.Name == "" {
            err := fmt.Errorf("InstallSQLClientsManager - config error: config.Name is empty")
            logger.Error("InstallSQLClientsManager - Error: %v", err)

            panic(err)
        }
        if val, ok := sqlClientsManagerInstance.sqlClients.Load(config.Name); ok {
            err := fmt.Errorf("InstallSQLClientsManager - config error: duplicated config.Name {%v}", val)
            logger.Error("InstallSQLClientsManager - Error: %v", err)

            panic(err)
        }

        sqlClientsManagerInstance.sqlClients.Store(config.Name, client)
    }
}

// GetSQLClient type;
func GetSQLClient(dbName string) (client *SQLClient) {
    if val, ok := sqlClientsManagerInstance.sqlClients.Load(dbName); ok {
        if client, ok = val.(*SQLClient); ok {
            return
        }
    }

    logger.Info("GetSQLClient - Not found client: %s", dbName)
    return
}

// GetSQLClientManager type;
func GetSQLClientManager() sync.Map {
    return sqlClientsManagerInstance.sqlClients
}
