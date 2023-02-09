Use links: http://39.105.137.222:8089/?p=2849


# 验证

1.创建 ns 和 pod

```bash
[17:30:06 root@master test]#kubectl create ns dev
namespace/dev created

[17:30:25 root@master test]#kubectl run nginx --image=nginx:1.16 --namespace=dev
pod/nginx created

[17:32:53 root@master test]#kubectl get pod -n dev 
NAME    READY   STATUS              RESTARTS   AGE
nginx   0/1     ContainerCreating   0          2m20s
```



## 9.1 验证 role 绑定 admin

1.创建一个 admin 的 role ，因为刚才在代码中写死了需要绑定 role 为 admin

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



2 通过程序创建 sa

```bash
[16:30:02 root@go sactl]#go run main.go create -n dev sa1
sa1 ServiceAccounts Create Success!
sa1 Rolebinding Create Success!
```



## 9.2 验证基于 SA 获取 POD

1 通过程序来获取 pod

```bash
[17:34:07 root@go sactl]#go run main.go check -n dev sa1 "https://10.0.0.131:6443"
NameSpace：dev  Pod：nginx

# 获取成功
```



### 9.3 验证基于 Token 获取 Pod 信息

通过 token 获取到对应的 Pod 信息



1.获取 token

```bash
# 获取 sa
[15:57:47 root@master ~]#kubectl get sa -n dev 
NAME      SECRETS   AGE
default   1         22h
sa1       1         22h

# 获取 role
[15:58:00 root@master ~]#kubectl get role -n dev 
NAME    CREATED AT
admin   2023-02-08T09:34:22Z

# 查看 secrets
[15:58:13 root@master ~]#kubectl get secrets -n dev 
NAME                  TYPE                                  DATA   AGE
default-token-9bmqc   kubernetes.io/service-account-token   3      22h
sa1-token-c5frn       kubernetes.io/service-account-token   3      22h

# 查看 rolebind
[15:58:22 root@master ~]#kubectl get rolebindings.rbac.authorization.k8s.io -n dev 
NAME   ROLE         AGE
sa1    Role/admin   22h

# 查看 POD
[15:58:31 root@master ~]#kubectl get pod -n dev 
No resources found in dev namespace.

# 创建 POD
[15:58:39 root@master ~]#kubectl run nginx --image=nginx:1.16 -n dev
pod/nginx created

# 获取 pod
[15:58:56 root@master ~]#kubectl get pod -n dev 
NAME    READY   STATUS              RESTARTS   AGE
nginx   0/1     ContainerCreating   0          4s

# 查看 token    
[15:59:00 root@master ~]#kubectl describe secrets -n dev sa1-token-c5frn 
Name:         sa1-token-c5frn
Namespace:    dev
Labels:       <none>
Annotations:  kubernetes.io/service-account.name: sa1
              kubernetes.io/service-account.uid: f12b5cbd-3444-4fd7-ac86-7fa877aa0d87

Type:  kubernetes.io/service-account-token

Data
====
ca.crt:     1099 bytes
namespace:  3 bytes
# 下面就是获取到的 token 信息
token:      		eyJhbGciOiJSUzI1NiIsImtpZCI6IkxlX2JRWEFDNzRCNmFodHp2VzZBanBwMFlDYWxYYk1jMS1CVjMyV3lfOXMifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZXYiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlY3JldC5uYW1lIjoic2ExLXRva2VuLWM1ZnJuIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6InNhMSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6ImYxMmI1Y2JkLTM0NDQtNGZkNy1hYzg2LTdmYTg3N2FhMGQ4NyIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDpkZXY6c2ExIn0.W0iIh7gn3ajuUGN15TvdwDMA3NhhDw-OflrFqZmw6Cvj3_V1kB1o-o1bvgPOdpf758rvSITg96zHSmHguZaRpmUnG7NsEvBHO78hpmzThcWLY8-jcP69nfyeRwzF2M7V3oZ9KyACi1pF5DvYhoUFgpBQ9bIgSGPRlwtM-0YbPOldwPCbImEomHtQhuTZ_tAqm43-xc_jnfnm1kXjP-zhWoq_a_SD28hL_G3CunrMoYajc0VLVcLPIAu53qAHDGp4XQX0Zwo2NQYhU2EGLT4mN7dBZKq36VdS8rTjUegke9dH9zLa6chZUZPX_q6LsW_QfQBK4s84AFdVjnBeAIL7Bg
```



2.工具基于 token 访问 pod

```bash
# token 可以看到是相同的
[17:31:57 root@go sactl]#go run main.go get -n dev  "https://10.0.0.131:6443" "eyJhbGciOiJSUzI1NiIsImtpZCI6IkxlX2JRWEFDNzRCNmFodHp2VzZBanBwMFlDYWxYYk1jMS1CVjMyV3lfOXMifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZXYiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlY3JldC5uYW1lIjoic2ExLXRva2VuLWM1ZnJuIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6InNhMSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6ImYxMmI1Y2JkLTM0NDQtNGZkNy1hYzg2LTdmYTg3N2FhMGQ4NyIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDpkZXY6c2ExIn0.W0iIh7gn3ajuUGN15TvdwDMA3NhhDw-OflrFqZmw6Cvj3_V1kB1o-o1bvgPOdpf758rvSITg96zHSmHguZaRpmUnG7NsEvBHO78hpmzThcWLY8-jcP69nfyeRwzF2M7V3oZ9KyACi1pF5DvYhoUFgpBQ9bIgSGPRlwtM-0YbPOldwPCbImEomHtQhuTZ_tAqm43-xc_jnfnm1kXjP-zhWoq_a_SD28hL_G3CunrMoYajc0VLVcLPIAu53qAHDGp4XQX0Zwo2NQYhU2EGLT4mN7dBZKq36VdS8rTjUegke9dH9zLa6chZUZPX_q6LsW_QfQBK4s84AFdVjnBeAIL7Bg"

token 验证成功！
namespace:dev   pod:nginx
```

