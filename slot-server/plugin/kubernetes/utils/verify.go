package utils

import "slot-server/utils"

var (
	ProxyVerify    = utils.Rules{"ClusterId": {utils.NotEmpty()}, "Path": {utils.NotEmpty()}}
	TerminalVerify = utils.Rules{"ClusterId": {utils.NotEmpty()}, "Name": {utils.NotEmpty()}, "PodName": {utils.NotEmpty()}, "Namespace": {utils.NotEmpty()}}
)
