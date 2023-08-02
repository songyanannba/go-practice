package utils

import (
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"slot-server/global"
	"slot-server/plugin/kubernetes/model"
)

type kubeclient struct {
	Pod *Pod
}

func NewKubeClient(id int) (kube *kubeclient, err error) {
	var cluster model.Cluster
	if err := global.GVA_DB.Where("id = ?", id).First(&cluster).Error; err != nil {
		global.GVA_LOG.Error("cluster get failed, err:", zap.Any("err", err))
		return nil, err
	}

	var config *rest.Config
	if cluster.KubeType == 1 {
		config, err = clientcmd.RESTConfigFromKubeConfig([]byte(cluster.KubeConfig))
		if err != nil {
			global.GVA_LOG.Error("config get failed, err:", zap.Any("err", err))
		}
	} else if cluster.KubeType == 2 {
		config = &rest.Config{
			Host:            cluster.ApiAddress,
			BearerToken:     cluster.KubeConfig,
			TLSClientConfig: rest.TLSClientConfig{Insecure: true},
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		global.GVA_LOG.Error("kubernetes client init failed, err:", zap.Any("err", err))
	}

	client := &kubeclient{
		Pod: NewPod(clientset, config),
	}

	return client, nil
}
