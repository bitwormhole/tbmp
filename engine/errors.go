package engine

import "github.com/starter-go/vlog"

// HandleError ...
func HandleError(err error) {
	vlog.HandleError(err)
}

// HandleErrorX ...
func HandleErrorX(x any) {
	vlog.HandleErrorX(x)
}
