package cloudfunction_todo

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	wh "github.com/mytodolist1/webhook"
)

func init() {
	functions.HTTP("Webhook", wh.Post)
}