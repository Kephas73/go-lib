package main
//
//import (
//	"flag"
//	"log"
//
//	"github.com/spf13/viper"
//	"ggithub.com/Kephas73/go-lib/baselib/env"
//	"ggithub.com/Kephas73/go-lib/baselib/sql_client"
//)
//
//func init() {
//	flag.Parse()
//
//	viper.SetConfigFile(`config.json`)
//	if err := viper.ReadInConfig(); err != nil {
//		panic(err)
//	}
//
//	env.SetupConfigEnv()
//
//	logPath := env.LogPath
//	logPrefix := env.LogPrefix
//
//	if viper.GetBool(`Debug`) || env.LogPrintLevel == 5 {
//		log.Println("Service RUN on DEBUG mode")
//	} else {
//		log.Println("Service RUN on PRODUCTION mode")
//	}
//
//	if logPath == "" {
//		logPath = "/var/log/GtvPlus"
//	}
//
//	if logPrefix == "" {
//		logPrefix = "backend-game"
//	}
//}
//
//func main() {
//	sql_client.InstallSQLClientManager("SQL")
//	stats := sql_client.GetSQLClient("auto_update_dev").DB.Stats()
//	log.Printf("%+v", stats)
//}
