package constant

const (
    LogBackEndPrefix          string = "[BACKEND]"
    LogBackEndModulePrefix    string = "[golib]"
    LogBackEndMainInfoPrefix  string = "[Main_Info]"
    LogBackEndMainErrorPrefix string = "[Main_Error]"
)

const (
    LogInfoPrefix  = LogBackEndPrefix + LogBackEndModulePrefix + LogBackEndMainInfoPrefix
    LogErrorPrefix = LogBackEndPrefix + LogBackEndModulePrefix + LogBackEndMainErrorPrefix
)
