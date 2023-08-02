package kubernetes

import (
	"context"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"net/http"
	"regexp"
	"slot-server/plugin/kubernetes/model"
	"sort"
	"strconv"
)

type Interface interface {
	Config() (*rest.Config, error)
	Client() (*kubernetes.Clientset, error)
	VersionMinor() (int, error)
	IsNamespacedResource(resourceName string) (bool, error)
	GetUserNamespaceNames(c *gin.Context) ([]string, error)
}

type Kubernetes struct {
	*model.Cluster
}

func NewKubernetes(cluster *model.Cluster) Interface {
	return &Kubernetes{cluster}
}

func (k *Kubernetes) Config() (config *rest.Config, err error) {
	if k.KubeType == 1 {
		return clientcmd.RESTConfigFromKubeConfig([]byte(k.KubeConfig))
	} else if k.KubeType == 2 {
		return &rest.Config{
			Host:            k.ApiAddress,
			BearerToken:     k.KubeConfig,
			TLSClientConfig: rest.TLSClientConfig{Insecure: true},
		}, err
	}

	return clientcmd.RESTConfigFromKubeConfig([]byte(k.KubeConfig))
}

func (k *Kubernetes) VersionMinor() (int, error) {
	v, err := k.Version()
	if err != nil {
		return 0, err
	}
	reg := regexp.MustCompile("[^0-9]")
	minor, err := strconv.Atoi(reg.ReplaceAllString(v.Minor, ""))
	return minor, err
}

func (k *Kubernetes) Version() (*version.Info, error) {
	client, err := k.Client()
	if err != nil {
		return nil, err
	}
	return client.ServerVersion()
}

func (k *Kubernetes) Client() (*kubernetes.Clientset, error) {
	cfg, err := k.Config()
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(cfg)
}

func (k *Kubernetes) GetUserNamespaceNames(c *gin.Context) ([]string, error) {
	client, err := k.Client()
	if err != nil {
		return nil, err
	}

	ns, err := client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var nsList []string
	for i := range ns.Items {
		if ns.Items[i].Status.Phase == "Active" {
			nsList = append(nsList, ns.Items[i].Name)
		}
	}

	sort.Strings(nsList)
	return nsList, nil
}

func (k *Kubernetes) IsNamespacedResource(resourceName string) (bool, error) {
	if resourceName == "events" || resourceName == "nodes" {
		return false, nil
	}
	client, err := k.Client()
	if err != nil {
		return false, err
	}
	apiList, err := client.ServerPreferredNamespacedResources()
	if err != nil && len(apiList) == 0 {
		return false, err
	}
	for i := range apiList {
		for j := range apiList[i].APIResources {
			if apiList[i].APIResources[j].Name == resourceName {
				return true, nil
			}
		}
	}
	return false, nil
}

func GenerateTLSTransport(c *model.Cluster) (http.RoundTripper, error) {
	kube := NewKubernetes(c)
	kubeconfig, err := kube.Config()
	if err != nil {
		return nil, err
	}

	return rest.TransportFor(kubeconfig)
}
