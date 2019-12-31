package mylog

const (
	DEBUG    = iota
	TRACE
	INFO
	WARN
	ERROR
	CRITICAL
)

func GetLevelStr(level int) string {
	leverStr := ""
	switch level {
	case DEBUG:
		leverStr = "DEBUG"
	case TRACE:
		leverStr = "TRACE"
	case INFO:
		leverStr = "INFO"
	case WARN:
		leverStr = "WARN"
	case ERROR:
		leverStr = "ERROR"
	case CRITICAL:
		leverStr = "CRITICAL"
	default:
		leverStr = "DEBUG"
	}

	return leverStr
}
