package controller

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sactl/config"
	"sactl/linK8S"
	"sactl/logs"
	"sactl/model"
	"time"
)

var (
	CA        []byte
	Token     []byte
	ClientSet *kubernetes.Clientset
	Secrets   string
)

func Get_Secrets(newSA *model.SA) {
	config, err := config.ParseConfig()
	if err != nil {
		fmt.Println("Error from server Config:", err)
		logs.ErrInfo("GetSecrets", "Config Fail", err)
	}

	ClientSet, err = linK8S.Link_K8s(config.KubeConfig.ConfigPath)
	if err != nil {
		fmt.Println("Error from server ClientSet:", err)
		logs.ErrInfo("GetSecrets", "ClientSet", err)
	}

	saSet := ClientSet.CoreV1().ServiceAccounts(newSA.NameSpace)
	saList, err := saSet.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error from server SAList:", err)
		logs.ErrInfo("GetSecrets 函数", "saList 报错", err)
	}

	// get secrets name
	for _, sa := range saList.Items {
		if sa.Name == newSA.Name {
			Secrets = sa.Secrets[0].Name
		}
	}

}

func Get_Token(newSA *model.SA) {
	secretsSet := ClientSet.CoreV1().Secrets(newSA.NameSpace)

	se, err := secretsSet.Get(context.TODO(), Secrets, metav1.GetOptions{})
	if err != nil {
		fmt.Println("Error from server Get_Token:", err)
		logs.ErrInfo("Get_Token", "Secrets", err)
	}

	CA = se.Data["ca.crt"]
	Token = se.Data["token"]
}

// 通过 token 和 CA 获取资源用于验证
func Get_RS(newSA *model.SA) {
	// 通过CA 来获取 client
	tlsClientConfig := rest.TLSClientConfig{
		CAData: CA,
	}

	config := rest.Config{
		Host:            newSA.API,
		BearerToken:     string(Token),
		TLSClientConfig: tlsClientConfig,
		Timeout:         20 * time.Second,
	}

	RSclientSet, err := kubernetes.NewForConfig(&config)
	if err != nil {
		fmt.Println("Error from server (NotFound) 基于 token CA 检查失败:", err)
		logs.ErrInfo("GetRS", "基于 token CA 检查失败", err)
	}

	pods := RSclientSet.CoreV1().Pods(newSA.NameSpace)
	podsList, err := pods.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error from server (NotFound):", err)
		logs.ErrInfo("GetRS podsList", "基于 token CA 获取 POD 失败", err)
	}

	for _, pod := range podsList.Items {
		fmt.Printf("NameSpace：%s\tPod：%s\n", pod.Namespace, pod.Name)
	}
}
