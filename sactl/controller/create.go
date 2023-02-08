package controller

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	config2 "sactl/config"
	"sactl/linK8S"
	"sactl/logs"
	"sactl/model"
)

// create ServiceAccount
func Create_SA(newSa *model.SA) {

	config, err := config2.ParseConfig()
	if err != nil {
		logs.ErrInfo("Create_SA", "create SA config error!", err)
	}
	clientSet, err := linK8S.Link_K8s(config.KubeConfig.ConfigPath)
	if err != nil {
		logs.ErrInfo("Create_SA", "create SA clientSet error!", err)
	}

	sa, err := clientSet.CoreV1().ServiceAccounts(newSa.NameSpace).Create(context.TODO(), &v1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      newSa.Name,
			Namespace: newSa.NameSpace,
		},
	}, metav1.CreateOptions{})
	if err != nil {
		logs.ErrInfo("Create_SA", "创建 SA Create 时操作失败", err)
		fmt.Println("Failed to create ServiceAccounts!\n", err)
	} else {
		fmt.Printf("%v ServiceAccounts Create Success!\n", sa.Name)
	}
}

// kubectl create sa and rolebind admin role
func Role_Bind(newRoleBind *model.RoleBind) {
	config, err := config2.ParseConfig()
	if err != nil {
		logs.ErrInfo("Create_SA", "创建 SA config 位置失败", err)
	}
	clientSet, err := linK8S.Link_K8s(config.KubeConfig.ConfigPath)
	if err != nil {
		logs.ErrInfo("Create_SA", "创建 SA clientSet 位置失败", err)
	}
	roleBind, err := clientSet.RbacV1().RoleBindings(newRoleBind.NameSpace).Create(context.TODO(), &rbacv1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      newRoleBind.Name,
			Namespace: newRoleBind.NameSpace,
		},
		Subjects: []rbacv1.Subject{
			{Kind: "ServiceAccount",
				Name:      newRoleBind.Name,
				Namespace: newRoleBind.NameSpace},
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "Role",
			Name:     "admin",
		},
	}, metav1.CreateOptions{})
	if err != nil {
		logs.ErrInfo("roleBind", "create roleBind error!", err)
		fmt.Println("Failed to create roleBind!\n", err)
	} else {
		fmt.Printf("%v Rolebinding Create Success!\n", roleBind.Name)
	}
}
