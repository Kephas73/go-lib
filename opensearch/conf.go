package opensearch

import (
    "encoding/json"
    "fmt"
    "github.com/spf13/viper"
    "golib/constant"
    "golib/util"
    "strings"
    "time"
)

type OpenSearchConfig struct {
    Hosts       []string
    Username    string
    Password    string
    IndexFormat string
}

var openSearchConf *OpenSearchConfig

func createConfigFromEnv(configKeys ...string) {
    configKey := "OpenSearch"
    for _, envKey := range configKeys {
        envKeyTrim := strings.TrimSpace(envKey)
        if envKeyTrim != "" {
            configKey = envKeyTrim
        }
    }

    openSearchConf = &OpenSearchConfig{}

    if err := viper.UnmarshalKey(configKey, openSearchConf); err != nil {
        err := fmt.Errorf("not found config name with env %q for open search with error: %+v", configKey, err)
        panic(err)
    }

    if len(openSearchConf.Hosts) == 0 {
        err := fmt.Errorf("not found hosts for open search with env %q", fmt.Sprintf("%s.Hosts", configKey))
        panic(err)
    }

    if openSearchConf.Username == "" {
        openSearchConf.Username = constant.UsernameDefault
    }

    if openSearchConf.Password == "" {
        openSearchConf.Password = constant.PasswordDefault
    }

    if openSearchConf.IndexFormat == "" {
        openSearchConf.IndexFormat = constant.IndexDefault
    }

}

type DataLogBase interface {
    FromJSON(data []byte) error
    ToJSON() string
    SetEventName(eventName string) DataLogBase
    SetLogDataJSON(dataJSON string) DataLogBase
    SetDescription(description string) DataLogBase
}

// DefaultDataLog type;
type DefaultDataLog struct {
    EventName    string      `json:"event_name,omitempty"`
    LogDataJSON  interface{} `json:"log_data_json,omitempty"`
    Description  string      `json:"description,omitempty"`
    TimeStarted  int64       `json:"time_started,omitempty"`
    TimeFinished int64       `json:"time_finished,omitempty"`
    TimeExecute  int64       `json:"time_execute,omitempty"`
}

// FromJSON func;
func (obj *DefaultDataLog) FromJSON(data []byte) error {
    return json.Unmarshal(data, obj)
}

// ToJSON func;
func (obj *DefaultDataLog) ToJSON() string {
    return util.JSONDebugDataString(obj)
}

// SetEventName func;
func (obj *DefaultDataLog) SetEventName(eventName string) DataLogBase {
    obj.EventName = eventName

    return obj
}

// SetLogDataJSON func;
func (obj *DefaultDataLog) SetLogDataJSON(dataJSON string) DataLogBase {
    obj.LogDataJSON = dataJSON

    return obj
}

// SetDescription func;
func (obj *DefaultDataLog) SetDescription(description string) DataLogBase {
    obj.Description = description

    return obj
}

func NewDefaultDataLog() *DefaultDataLog {
    timeNow := time.Now().Unix()
    return &DefaultDataLog{
        EventName:    "DefaultLog",
        TimeStarted:  timeNow,
        TimeFinished: timeNow,
        TimeExecute:  0,
    }
}
