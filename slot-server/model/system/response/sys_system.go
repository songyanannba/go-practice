package response

import "slot-server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
