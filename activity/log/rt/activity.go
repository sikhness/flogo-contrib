package log

import (
	"fmt"
	"os"
	"strconv"

	"github.com/TIBCOSoftware/flogo-lib/core/ext/activity"
	"github.com/op/go-logging"
)

// activityLog is the default logger for the Log Activity
var activityLog = logging.MustGetLogger("activity-tibco-log")

// format is the log format for the Activity log
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{color:reset} %{message}`,
)

var backend = logging.NewLogBackend(os.Stdout, "", 0)
var backendFormatter = logging.NewBackendFormatter(backend, format)
var backendLeveled = logging.AddModuleLevel(backendFormatter)

func init() {
	backendLeveled.SetLevel(logging.INFO, "")
	activityLog.SetBackend(backendLeveled)
}

// LogActivity is an Activity that is used to log a message to the console
// inputs : {message, flowInfo}
// outputs: none
type LogActivity struct {
	metadata *activity.Metadata
}

func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(&LogActivity{metadata: md})
}

// Metadata returns the activity's metadata
func (a *LogActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *LogActivity) Eval(context activity.Context) bool {

	message, _ := context.GetAttrValue("message")
	flowInfo, _ := context.GetAttrValue("flowInfo")

	msg := message

	showInfo, _ := strconv.ParseBool(flowInfo)

	if showInfo {

		msg = fmt.Sprintf("%s - FlowInstanceID [%s], Flow [%s], Task [%s]", msg,
			context.FlowInstanceID(), context.FlowName(), context.TaskName())
	}

	activityLog.Info(msg)

	//log.Debugf("%s: %s\n", time.Now(), msg)

	return true
}