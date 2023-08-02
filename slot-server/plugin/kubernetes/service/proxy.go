package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
	"net/url"
	"slot-server/global"
	kubernetesReq "slot-server/plugin/kubernetes/model/kubernetes/request"
	"slot-server/plugin/kubernetes/utils/kubernetes"
	"strings"
	"sync"
)

type ProxyService struct{}

type ItemList []interface{}

type Item struct {
	Metadata metav1.ObjectMeta `json:"metadata"`
}

type K8sListObj struct {
	Kind       string      `json:"kind"`
	ApiVersion string      `json:"apiVersion"`
	Metadata   interface{} `json:"metadata"`
	Items      ItemList    `json:"items"`
}

type K8sObj interface{}

// 参数拆分
func parseResourceName(path string) (string, error) {
	ss := strings.Split(path, "/")
	if len(ss) > 0 {
		return ss[len(ss)-1], nil
	}
	return "", fmt.Errorf("cant not get resource name from url %s", path)
}

// 集群版本判断
func compatibleClusterVersion(minor int, path *string) {
	p := *path

	if minor <= 18 {
		if strings.Contains(p, "networking.k8s.io/v1") && strings.Contains(p, "ingresses") {
			p = strings.Replace(p, "networking.k8s.io/v1", "networking.k8s.io/v1beta1", -1)
		}
	}

	if minor > 18 {
		if strings.Contains(p, "/apis/batch/v1beta1") && strings.Contains(p, "cronjobs") {
			p = strings.Replace(p, "/apis/batch/v1beta1", "/apis/batch/v1", -1)
		}
	}

	*path = p
}

// 字段过滤
type fieldMatcher interface {
	Match(item interface{}) bool
}

// 字段过滤
func fieldFilter(data []interface{}, fms ...fieldMatcher) []interface{} {
	var result []interface{}
	for i := range data {
		for j := range fms {
			if fms[j].Match(data[i]) {
				result = append(result, data[i])
				break
			}
		}
	}
	return result
}

// 分页
func pageFilter(num, size int, data []interface{}) (int64, []interface{}, error) {
	total := len(data)
	result := make([]interface{}, 0)
	if num*size < len(data) {
		result = data[(num-1)*size : (num * size)]
	} else {
		result = data[(num-1)*size:]
	}
	return int64(total), result, nil
}

// 关键字
type keywordsMatcher struct {
	keywords string
}

// 关键字匹配
func (n keywordsMatcher) Match(item interface{}) bool {
	pageItem := item.(map[string]interface{})
	if pageItem["metadata"].(map[string]interface{})["namespace"] != nil && pageItem["metadata"].(map[string]interface{})["namespace"].(string) == n.keywords {
		return true
	}
	if strings.Contains(pageItem["metadata"].(map[string]interface{})["name"].(string), n.keywords) {
		return true
	}
	if pageItem["message"] != nil && strings.Contains(strings.ToLower(pageItem["message"].(string)), strings.ToLower(n.keywords)) {
		return true
	}
	return false
}

func withNamespaceAndNameMatcher(keywords string) fieldMatcher {
	return &keywordsMatcher{
		keywords: keywords,
	}
}

// 分页
func pagerAndSearch(page, pageSize int, items ItemList, keywords string) (*kubernetesReq.ProxyResult, error) {
	var p kubernetesReq.ProxyResult
	if keywords != "" {
		items = fieldFilter(items, withNamespaceAndNameMatcher(keywords))
	}
	if page != 0 && pageSize != 0 {
		tt, items, err := pageFilter(page, pageSize, items)
		if err != nil {
			return nil, err
		}
		p.Total = tt
		p.Items = items
		p.Page, p.PageSize = page, pageSize
		return &p, nil
	}
	p.Total = int64(len(items))
	p.Items = items
	p.Page = page
	p.PageSize = pageSize
	return &p, nil
}

// 地址拼接访问多命名空间资源
func RequestNamespaceRes(path string, ns string) string {
	ss := strings.Split(path, "/")
	resourceName := ss[len(ss)-1]
	per := ss[:len(ss)-1]
	namespacedSs := append(per, "namespaces", ns, resourceName)
	return strings.Join(namespacedSs, "/")
}

type NamespaceResourceContainer struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []interface{} `json:"items"`
	Namespaces      []string      `json:"namespaces"`
}

// 异步获取命名空间资源
func fetchMultiNamespaceResource(client *http.Client, namespaces []string, apiUrl url.URL) (*NamespaceResourceContainer, error) {
	wg := &sync.WaitGroup{}
	var mergedContainer NamespaceResourceContainer
	var responses []*http.Response
	var es []error
	for i := range namespaces {
		wg.Add(1)
		ns := namespaces[i]
		go func() {
			newUrl := apiUrl
			newUrl.Path = RequestNamespaceRes(apiUrl.Path, ns)
			resp, err := client.Get(newUrl.String())
			if err != nil {
				es = append(es, err)
				wg.Done()
				return
			}
			responses = append(responses, resp)
			wg.Done()
		}()

	}
	wg.Wait()
	var forbidden int
	var forbiddenMessage []string
	for i := range responses {
		r := responses[i]
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		if r.StatusCode != http.StatusOK {
			if r.StatusCode == http.StatusForbidden {
				forbidden++
				forbiddenMessage = append(forbiddenMessage, string(body))
				continue
			} else {
				return nil, errors.New(string(body))
			}
		}
		var nc NamespaceResourceContainer
		if err := json.Unmarshal(body, &nc); err != nil {
			return nil, err
		}
		mergedContainer.TypeMeta = nc.TypeMeta
		mergedContainer.ListMeta = nc.ListMeta
		mergedContainer.Items = append(mergedContainer.Items, nc.Items...)
	}
	if len(namespaces) == 1 && forbidden == 1 {
		return nil, errors.New(strings.Join(forbiddenMessage, ""))
	}
	return &mergedContainer, nil
}

// 选择器拼接
func Selector(proxy kubernetesReq.ProxyRequest, urlParam kubernetesReq.ProxyParamRequest) (kubernetesReq.ProxyRequest, kubernetesReq.ProxyParamRequest) {

	if proxy.FieldSelector != "" && proxy.LabelSelector != "" {
		urlParam.Path += "?fieldSelector=" + proxy.FieldSelector + "&labelSelector=" + proxy.LabelSelector
	} else if proxy.FieldSelector != "" && proxy.LabelSelector == "" {
		urlParam.Path += "?fieldSelector=" + proxy.FieldSelector
	} else if proxy.LabelSelector != "" {
		urlParam.Path += "?labelSelector=" + proxy.LabelSelector
	}

	return proxy, urlParam
}

// 多命名空间资源获取
func fetchMultiNamespace(c *gin.Context, proxy kubernetesReq.ProxyRequest, httpClient http.Client, apiUrl *url.URL, allowedNamespaces []string) (ret kubernetesReq.ProxyResult, err error) {
	//global.GVA_LOG.Info(fmt.Sprintf("调用多namespace逻辑: 请求类型: %s Url: %s Body: %s", c.Request.Method, apiUrl.String(), c.Request.Body))
	if err != nil {
		global.GVA_LOG.Error("allowedNamespaces request failed: " + err.Error())
		return ret, err
	}

	//获取已有权限的命名空间资源
	resp, err := fetchMultiNamespaceResource(&httpClient, allowedNamespaces, *apiUrl)
	if err != nil {
		global.GVA_LOG.Error("fetchMultiNamespaceResource request failed: " + err.Error())
		return ret, err
	}

	//分页
	PageResult, err := pagerAndSearch(proxy.Page, proxy.PageSize, resp.Items, proxy.Keywords)
	if err != nil {
		global.GVA_LOG.Error("pagerAndSearch  failed: " + err.Error())
		return ret, err
	}

	return *PageResult, err
}

//@function: Option
//@description: 代理操作
//@param: c *gin.Context, proxy kubernetesReq.ProxyRequest, urlParam kubernetesReq.ProxyParamRequest
//@return: ret kubernetesReq.ProxyResult, err error

func (p ProxyService) Option(c *gin.Context, proxy kubernetesReq.ProxyRequest, urlParam kubernetesReq.ProxyParamRequest) (ret kubernetesReq.ProxyResult, err error) {
	// 获取集群信息
	cluster, err := clusterService.GetGlusterById(urlParam.ClusterId)
	if err != nil {
		global.GVA_LOG.Error("get cluster info failed: " + err.Error())
		return ret, err
	}

	search := false
	if proxy.Search != "" {
		search = true
	}

	// 生成transport
	ts, err := kubernetes.GenerateTLSTransport(&cluster)
	if err != nil {
		global.GVA_LOG.Error("transport generate failed: " + err.Error())
		return ret, err
	}

	// 生成httpClient
	httpClient := http.Client{Transport: ts}
	k := kubernetes.NewKubernetes(&cluster)
	clusterVersionMinor, err := k.VersionMinor()
	if err != nil {
		global.GVA_LOG.Error("Minor Version cluster failed: " + err.Error())
		return ret, err
	}

	compatibleClusterVersion(clusterVersionMinor, &urlParam.Path)

	//判断是否已经包含了namespace的查询
	resourceName, err := parseResourceName(strings.Split(urlParam.Path, "?")[0])
	if err != nil {
		global.GVA_LOG.Error("Get  resourceName failed: " + err.Error())
		return ret, err
	}

	//判断资源类型是否是namespace级别的
	namespaced, err := k.IsNamespacedResource(resourceName)
	if err != nil {
		global.GVA_LOG.Error("IsNamespacedResource failed: " + err.Error())
		return ret, err
	}

	// 判断是否包含监控接口，或者包含命名空间
	if strings.Contains(urlParam.Path, "namespaces") || strings.Contains(urlParam.Path, "metrics.k8s.io") {
		namespaced = false
	}

	//选择器拼接
	proxy, urlParam = Selector(proxy, urlParam)

	//url格式化
	apiUrl, err := url.Parse(fmt.Sprintf("%s%s", cluster.ApiAddress, urlParam.Path))
	if err != nil {
		global.GVA_LOG.Error("url Parse failed: " + err.Error())
		return ret, err
	}

	// 调用多namespace逻辑
	if http.MethodGet == c.Request.Method && proxy.Namespace == "" && namespaced {
		//已获权的命名空间
		allowedNamespaces, err := k.GetUserNamespaceNames(c)
		if err != nil {
			return ret, err
		}

		return fetchMultiNamespace(c, proxy, httpClient, apiUrl, allowedNamespaces)
	}

	//创建http请求
	//global.GVA_LOG.Info(fmt.Sprintf("调用单namespace逻辑: 请求类型: %s Url: %s", c.Request.Method, apiUrl.String()))
	req, err := http.NewRequest(c.Request.Method, apiUrl.String(), c.Request.Body)
	if err != nil {
		global.GVA_LOG.Error("http request failed: " + err.Error())
		return ret, err
	}

	//修改头信息
	if c.Request.Method == "PATCH" {
		req.Header.Set("Content-Type", "application/merge-patch+json")
	}

	//发起请求
	resp, err := httpClient.Do(req)
	if err != nil {
		global.GVA_LOG.Error("http request do failed: " + err.Error())
		return ret, err
	}

	//取出数据
	rawResp, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode == http.StatusForbidden {
		resp.StatusCode = http.StatusInternalServerError
	}

	//解析出数据
	if req.Method == http.MethodGet && search {
		var listObj K8sListObj
		if err = json.Unmarshal(rawResp, &listObj); err != nil {
			global.GVA_LOG.Error("json Unmarshal failed: " + string(rawResp))
			return ret, errors.New(string(rawResp))
		}

		//分页
		PageResult, err := pagerAndSearch(proxy.Page, proxy.PageSize, listObj.Items, proxy.Keywords)
		if err != nil {
			global.GVA_LOG.Error("pagerAndSearch  failed: " + err.Error())
			return ret, err
		}

		return *PageResult, err
	}

	//解析数据
	if req.Method == http.MethodGet || req.Method == http.MethodPut || req.Method == http.MethodPost || req.Method == http.MethodPatch || req.Method == http.MethodDelete {
		var k8sObj K8sObj
		if err = json.Unmarshal(rawResp, &k8sObj); err != nil {
			global.GVA_LOG.Error("json Unmarshal failed: " + string(rawResp))
			return ret, errors.New(string(rawResp))
		}

		// 单个数据返回
		var PageResult kubernetesReq.ProxyResult
		PageResult.Items = k8sObj
		return PageResult, err
	}

	return ret, err
}
