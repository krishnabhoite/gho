package gonam

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {
	// do eval

	//Get activity data from the Context
	name := context.GetInput("name").(string)
	salutation := context.GetInput("salutation").(string)

	//use the log object to log the debug messages
	logger.Debugf("The flogo engine says [%s] to [%s]", salutation, name)

	//set the result as part of the Context
	context.SetOutput("result", "The flogo engines says "+salutation+" to "+name)

	//signal the flog engines that operaiton is complete
	return true, nil
}
