package response

import (
	"slot-server/plugin/kubernetes/model"
)

type ClusterResponse struct {
	Cluster model.Cluster `json:"cluster"`
}
