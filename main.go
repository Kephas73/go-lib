package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "github.com/spf13/viper"
    "golib/document"
    "golib/logger"
    "golib/opensearch"
    "golib/util"
)

func init() {
    viper.SetConfigFile(`config.json`)

    err := viper.ReadInConfig()

    if err != nil {
        panic(err)
    }

    if viper.GetBool(`Debug`) {
        fmt.Println("Service RUN on DEBUG mode")
    } else {
        fmt.Println("Service RUN on PRODUCTION mode")
    }
}

func main() {

    logPath := viper.GetString("Log.Path")
    logPrefix := viper.GetString("Log.Prefix")
    logger.NewLogger(logPath, logPrefix)

    /*// Logstash
      logstash.InstallLogStashClient()

      doc := document.MakeDocument()
      data := document.NewDefaultData()
      data.EventName = "ES_Pub"
      data.Data = ES{Name: "123456"}

      doc.Document = data

      for {
      logstash.GetLogStashClient().InsertDocument(doc)
      time.Sleep(time.Second)
      }*/

    /*// Open search
     opensearch.InstallOpenSearchClient()
     for {
      doc := document.MakeDocument()
      doc.RandomIDDoc()
      data := document.NewDefaultData()
      data.EventName = "OP_Pub"
      data.Data = ES2{Es: "123456"}
      doc.Document = data
    
      fmt.Println(doc)
      opensearch.GetOpenSearchClient().InsertDocument(doc)
      time.Sleep(time.Second * 3)
     }*/

    // Query open search
    opensearch.InstallOpenSearchClient()

    termString := new(document.TermStringBuilder)
    termString.Term = make(map[string]interface{})
    termString.Term["document.event_name.keyword"] = "ES_Pub"

    query := new(document.QueryBuilder)
    query.Query.Bool.Must = append(query.Query.Bool.Must, termString)

    var queryMapping map[string]interface{}
    _ = json.Unmarshal(util.JSONDebugData(query), &queryMapping)

    var buf bytes.Buffer
    if err := json.NewEncoder(&buf).Encode(queryMapping); err != nil {
        panic(err)
    }
    
    res, err := opensearch.GetOpenSearchClient().CountDocument([]string{"cms-*"}, &buf)
    if err != nil {
        fmt.Println(err)
    }
    
    fmt.Println(res.Count)

}

type ES struct {
    Name string `json:"name"`
}

type ES2 struct {
    Es string `json:"es"`
}
