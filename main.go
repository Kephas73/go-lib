package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "github.com/Kephas73/go-lib/document"
    "github.com/Kephas73/go-lib/logger"
    "github.com/Kephas73/go-lib/opensearch"
    "github.com/Kephas73/go-lib/util"
    "github.com/spf13/viper"
    "math"
    "time"
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

//func main() {
//
//    logPath := viper.GetString("Log.Path")
//    logPrefix := viper.GetString("Log.Prefix")
//    logger.NewLogger(logPath, logPrefix)
//
//    /*// Logstash
//      logstash.InstallLogStashClient()
//
//      doc := document.MakeDocument()
//      data := document.NewDefaultData()
//      data.EventName = "ES_Pub"
//      data.Data = ES{Name: "123456"}
//
//      doc.Document = data
//
//      for {
//      logstash.GetLogStashClient().InsertDocument(doc)
//      time.Sleep(time.Second)
//      }*/
//
//    /*// Open search
//      opensearch.InstallOpenSearchClient()
//      for {
//       doc := document.MakeDocument()
//       doc.RandomIDDoc()
//       data := document.NewDefaultData()
//       data.EventName = "OP_Pub"
//       data.Data = ES2{Es: "123456"}
//       doc.Document = data
//
//       fmt.Println(doc)
//       opensearch.GetOpenSearchClient().InsertDocument(doc)
//       time.Sleep(time.Second * 3)
//      }*/
//
//     // Query open search
//     opensearch.InstallOpenSearchClient()
//     rangeStringBuilder := new(document.RangeStringBuilder)
//     rangeStringBuilder.Range = make(map[string]interface{})
//     rangeStringBuilder.Range["@timestamp"] = map[string]interface{}{
//         "gte": "now-1h",
//     }
//    
//     termsStringBuilder := new(document.TermsStringBuilder)
//     termsStringBuilder.Terms = make(map[string]interface{})
//     termsStringBuilder.Terms["field"] = "document.data.PID"
//     termsStringBuilder.Terms["size"] = 1000000000
//    
//     aggsConditionBuilder := new(document.AggsCondition)
//     aggsConditionBuilder.ResponseCodes = termsStringBuilder
//    
//     query := new(document.QueryBuilder)
//     query.Query.Bool.Must = append(query.Query.Bool.Must, rangeStringBuilder)
//     query.Aggs = aggsConditionBuilder
//    
//     var queryMapping map[string]interface{}
//     _ = json.Unmarshal(util.JSONDebugData(query), &queryMapping)
//    
//     var buf bytes.Buffer
//     if err := json.NewEncoder(&buf).Encode(queryMapping); err != nil {
//         panic(err)
//     }
//    
//     res, err := opensearch.GetOpenSearchClient().SearchDocument([]string{"watching_time-*"}, &buf)
//     if err != nil {
//         fmt.Println("AAAAAA")
//         fmt.Println(err)
//     }
//    
//     //fmt.Println(res)
//     //fmt.Println(res.Aggregations.ResponseCodes.Buckets)
//     fmt.Println(res.Count)
//     fmt.Println(len(res.Aggregations.ResponseCodes.Buckets))
//     fmt.Println(time.Now().Add(time.Duration(math.MaxInt32) * time.Second).Unix())
//}

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

    //a := map[string]interface{}{
    //    "cardinality" : map[string]interface{}{
    //        "field":"document.data.PID",
    //    },
    //}
    //
    //fmt.Println(a)
    
    // Query open search
    opensearch.InstallOpenSearchClient()
    rangeStringBuilder := new(document.RangeStringBuilder)
    rangeStringBuilder.Range = make(map[string]interface{})
    rangeStringBuilder.Range["@timestamp"] = map[string]interface{}{
        "gte": "now-1h",
    }

    
    cardinalityStringBuilder := new(document.CardinalityStringBuilder)
    cardinalityStringBuilder.Cardinality = make(map[string]interface{})
    cardinalityStringBuilder.Cardinality["field"] = "document.data.PID"
    // termsStringBuilder.Terms["size"] = 1000000000

    aggsConditionBuilder := new(document.AggsCondition)
    aggsConditionBuilder.ResponseCodes = cardinalityStringBuilder

    query := new(document.QueryBuilder)
    query.Query.Bool.Must = append(query.Query.Bool.Must, rangeStringBuilder)
    query.Aggs = aggsConditionBuilder

    var queryMapping map[string]interface{}
    _ = json.Unmarshal(util.JSONDebugData(query), &queryMapping)

    var buf bytes.Buffer
    if err := json.NewEncoder(&buf).Encode(queryMapping); err != nil {
        panic(err)
    }

    res, err := opensearch.GetOpenSearchClient().SearchDocument([]string{"watching_time-*"}, &buf)
    if err != nil {
        fmt.Println("AAAAAA")
        fmt.Println(err)
    }

    //fmt.Println(res)
    //fmt.Println(res.Aggregations.ResponseCodes.Buckets)
    fmt.Println(res.Count)
    fmt.Println(len(res.Aggregations.ResponseCodes.Buckets))
    fmt.Println(res.Aggregations.ResponseCodes.Value)
    fmt.Println(time.Now().Add(time.Duration(math.MaxInt32) * time.Second).Unix())
}

type ES struct {
    Name string `json:"name"`
}

type ES2 struct {
    Es string `json:"es"`
}