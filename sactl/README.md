Use links: http://39.105.137.222:8089/?p=2849

1.create ns and pod

```bash
[17:30:06 root@master test]#kubectl create ns dev
namespace/dev created

[17:30:25 root@master test]#kubectl run nginx --image=nginx:1.16 --namespace=dev
pod/nginx created

[17:32:53 root@master test]#kubectl get pod -n dev 
NAME    READY   STATUS              RESTARTS   AGE
nginx   0/1     ContainerCreating   0          2m20s
```



2.create admin a role ，Because it was written dead in the code just now, you need to bind the role to admin


```yaml
[17:30:35 root@master test]#cat test_role.yaml 
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  namespace: dev
  name: admin	   # 指定 role 为 admin
rules:
- apiGroups: ["*"] # "" 标明 core API 组以及所有资源和操作的权限， * 表示所有
  resources: ["*"]
  verbs: ["*"]

[17:32:55 root@master test]#kubectl apply -f test_role.yaml 
role.rbac.authorization.k8s.io/admin created
```



3 Create sa by program


```bash
[16:30:02 root@go sactl]#go run main.go create -n dev sa1
sa1 ServiceAccounts Create Success!
sa1 Rolebinding Create Success!
```



4 Get pod through program


```bash
[17:34:07 root@go sactl]#go run main.go check -n dev sa1 "https://10.0.0.131:6443"
NameSpace：dev  Pod：nginx
```

