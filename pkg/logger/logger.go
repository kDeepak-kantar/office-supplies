package logger

import (
	"encoding/json"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// LogValue is used to send additional information when logging.
type LogValue map[string]interface{}

// SeverityStatus is used to describe the severity level when logging.
type SeverityStatus byte

const (
	// SeverityInfo is used to log information.
	SeverityInfo SeverityStatus = iota
	// SeverityWarning is used to log warrnings.
	SeverityWarning
	// SeverityError is used to log errors.
	SeverityError
	// SeverityCritical is used to log critical errors.
	SeverityCritical
)

// severityStatus names is the internal structure for severity level pretty names.
var severityStatusName = [...]string{
	SeverityInfo:     "INFO",
	SeverityWarning:  "WARNING",
	SeverityError:    "ERROR",
	SeverityCritical: "CRITICAL",
}

// severityStatusLevel is the internal structure for severity level integer level.
var severityStatusLevel = [...]int{
	SeverityInfo:     20,
	SeverityWarning:  30,
	SeverityError:    40,
	SeverityCritical: 50,
}

// systemLog is an internal structure used to specify a system log entry.
type systemLog struct{}

// systemInfoLog is an internal structure used to specify a system info log entry.
type systemInfoLog struct {
	Time               string   `json:"asctime"`
	GroupName          string   `json:"group,omitempty"`
	Message            string   `json:"message,omitempty"`
	SeverityStatusName string   `json:"levelname"`
	SeverityStatusNo   int      `json:"levelno"`
	AdditionalInfo     LogValue `json:"additionalInfo"`
}

// systemErrorLog is an internal structure used to specify a system error log entry.
type systemErrorLog struct {
	Time               string      `json:"asctime"`
	GroupName          string      `json:"group,omitempty"`
	SeverityStatusName string      `json:"levelname"`
	SeverityStatusNo   int         `json:"levelno"`
	ErrorMsg           string      `json:"message"`
	Error              string      `json:"error"`
	AdditionalInfo     LogValue    `json:"additionalInfo"`
	Stack              interface{} `json:"stack"`
}

// systemContextErrorLog is an internal structure used to specify a system context error log entry.
type systemContextErrorLog struct {
	Time               string      `json:"asctime"`
	SeverityStatusName string      `json:"levelname"`
	SeverityStatusNo   int         `json:"levelno"`
	ErrorMsg           string      `json:"message"`
	Error              string      `json:"error"`
	Stack              interface{} `json:"stack,omitempty"`
}

// statusName is used to look-up the status name of a severity level.
func (s SeverityStatus) statusName() string {
	return severityStatusName[s]
}

// statusNumber is used to look-up the status number of a severity level.
func (s SeverityStatus) statusNumber() int {
	return severityStatusLevel[s]
}

// stdErr is used to output errors and critical messages.
var stdErr = log.New(os.Stderr, "", 0)

// stdOut is used to output warnings and information.
var stdOut = log.New(os.Stdout, "", 0)

// mu is a mutex lock used to serialize the logging safely.
var mu = &sync.Mutex{}

// Debug is used to debug the output when logging.
var Debug bool

// Log will log an error and associated values to the reporting and monitor systems.
func Log(group string, severity SeverityStatus, msg string, err error, val LogValue) {
	mu.Lock()
	defer func() {
		mu.Unlock()
	}()

	logSystem(group, severity, msg, err, val)
}

// LogInfo will log basic info to the reporting and monitor systems.
func LogInfo(group string, msg string) {
	mu.Lock()
	defer func() {
		mu.Unlock()
	}()

	logSystem(group, SeverityInfo, msg, nil, nil)
}

// LogError will log an error and associated values to the reporting and monitor systems.
func LogError(group string, err error) {
	mu.Lock()
	defer func() {
		mu.Unlock()
	}()

	logSystem(group, SeverityError, err.Error(), err, nil)
}

// LogError will log an error and associated values to the reporting and monitor systems.
func LogCriticalError(group string, err error) {
	mu.Lock()
	defer func() {
		mu.Unlock()
	}()

	logSystem(group, SeverityCritical, err.Error(), err, nil)
}

// LogContextError is used to log context errors.
func LogContextError(err error, c *gin.Context) {
	mu.Lock()
	defer func() {
		mu.Unlock()
	}()

	if err != nil && c != nil {
		logContextToSystem("Context Error", err)
	}
}

// logContextToSystem is one of the main helping functions used to serialize error related
// to a context.
func logContextToSystem(msg string, err error) {
	marshalledErr := string(marshalError(err))

	sysLog := systemContextErrorLog{
		getLogTime(),
		SeverityError.statusName(),
		SeverityError.statusNumber(),
		msg,
		marshalledErr,
		getStack(),
	}

	var bSysLog []byte
	if Debug {
		bSysLog, _ = json.MarshalIndent(sysLog, "", "\t")
	} else {
		bSysLog, _ = json.Marshal(sysLog)
	}

	stdErr.Println(string(bSysLog))
}

// logSystem is the main helping function used to serialize log information to system output.
func logSystem(group string, severity SeverityStatus, msg string, err error, val LogValue) {
	var bSysLog []byte

	// Do not print trace if we report no errors or the log is not marked as error.
	if err == nil || (severity != SeverityCritical && severity != SeverityError) {
		sysLog := systemInfoLog{
			getLogTime(),
			group,
			msg,
			severity.statusName(),
			severity.statusNumber(),
			val,
		}

		if Debug {
			bSysLog, _ = json.MarshalIndent(sysLog, "", "\t")
		} else {
			bSysLog, _ = json.Marshal(sysLog)
		}
	} else {
		if val == nil {
			val = LogValue{}
		}

		stack := getStack()
		marshalledErr := string(marshalError(err))

		val["error"] = marshalledErr
		val["stack"] = stack

		sysLog := systemErrorLog{
			GroupName:          group,
			Time:               getLogTime(),
			SeverityStatusName: severity.statusName(),
			SeverityStatusNo:   severity.statusNumber(),
			ErrorMsg:           msg,
			Error:              marshalledErr,
			AdditionalInfo:     val,
			Stack:              stack,
		}

		if Debug {
			bSysLog, _ = json.MarshalIndent(sysLog, "", "\t")
		} else {
			bSysLog, _ = json.Marshal(sysLog)
		}
	}

	switch severity {
	case SeverityWarning, SeverityInfo:
		stdOut.Println(string(bSysLog))
	default:
		stdErr.Println(string(bSysLog))
	}
}

// marshalError serializes an error to a byte array.
func marshalError(err error) []byte {
	if err == nil {
		return []byte{}
	}

	if _, ok := err.(json.Marshaler); ok {
		var b []byte

		if Debug {
			b, _ = json.MarshalIndent(err, "", "\t")
		} else {
			b, _ = json.Marshal(err)
		}

		// Return the marshalled error.
		return b
	}

	// Fallback to the error string.
	return []byte(err.Error())
}

// getLogTime returns the log time in the format described by RFC3339.
func getLogTime() string {
	t := time.Now().UnixNano()
	unix := time.Unix(0, t)
	rft := unix.Format(time.RFC3339Nano)
	return rft
}

// getStack returns a stack trace of current stack.
func getStack() interface{} {
	buf := make([]byte, 1024)
	for {
		n := runtime.Stack(buf, false)
		if n < len(buf) {
			buf = buf[:n]
			break
		}
		buf = make([]byte, 2*len(buf))
	}

	if Debug {
		trace := strings.Split(string(buf), "\n")
		trace = trace[:len(trace)-2]
		traceTree := make([]interface{}, len(trace))
		for i := range trace {
			if strings.HasPrefix(trace[i], "\t") {
				traceTree[i] = strings.Split(trace[i], "\t")[1:]
			} else {
				traceTree[i] = trace[i]
			}
		}
		return traceTree
	} else {
		return string(buf)
	}
}
