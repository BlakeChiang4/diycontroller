package main

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
	"k8s.io/kubectl/pkg/scheme"
)

var (
	configFile = "/root/.kube/config"
	ApiPath    = "api"
	nameSpace  = "kube-system"
	resource   = "pods"
)

func main() {
	defer klog.Flush()
	klog.Info("this is the first RestClient")
	// 生成config
	config, err := clientcmd.BuildConfigFromFlags("", configFile)
	if err != nil {
		panic(err)
	}
	config.APIPath = ApiPath
	config.GroupVersion = &corev1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs

	// 生成restClient
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}
	// 声明空结构体
	rest := &corev1.PodList{}
	if err = restClient.Get().Namespace(nameSpace).Resource(resource).VersionedParams(&metav1.ListOptions{Limit: 500},
		scheme.ParameterCodec).Do(context.TODO()).Into(rest); err != nil {
		panic(err)
	}
	for _, v := range rest.Items {
		fmt.Printf("NameSpace: %v  Name: %v  Status: %v \n", v.Namespace, v.Name, v.Status.Phase)
	}
}
