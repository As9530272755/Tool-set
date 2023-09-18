#♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥>♥♥♥♥♥
# File Name: install_nginx.sh
# Author: 老张
# mail: as953027255@qq.com
# http://1717zgy.com/
# Created Time: Friday, August 07, 2020 PM11:04:54 CST
#♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥♥>♥♥♥♥♥
#!/bin/bash

#声明变量
USER=nginx

DIR=/apps/nginx/

#安装依赖
yum install -y epel-release ;yum install -y libxml2-devel bzip2-devel libmcrypt-devel sqlite-devel gcc gcc-c++ glibc glibc-devel pcre pcre-devel openssl openssl-devel systemd-devel net-tools iotop bc zip unzip zlib-devel bash-completion nfs-utils automake libxml2 libxml2-devel libxslt libxslt-devel perl perl-ExtUtils-Embed

#判断nginx用户是否存在
if id "$USER" &> /dev/null;then
	echo "${USER} is exist"
else 
	useradd -s /sbin/nologin ${USER} -u 2000
	echo "${USER} is Creating a successful"
fi

#判断安装目录是否存在，不存在便创建
if [ -d ${DIR} ];then
	echo "${DIR}不存在"
else
	mkdir -vp ${DIR}
fi


#进入安装目录
cd /usr/local/src/

#下载nginx
wget http://nginx.org/download/nginx-1.19.1.tar.gz

#解压nginx
tar xvf nginx-1.19.1.tar.gz

#进入nginx目录开始准备安装
cd /usr/local/src/nginx-1.19.1

#开始编译
./configure --prefix=/apps/nginx --user=nginx --group=nginx --with-http_ssl_module --with-http_v2_module --with-http_realip_module --with-http_stub_status_module --with-http_gzip_static_module --with-pcre --with-stream --with-stream_ssl_module --with-stream_realip_module

#编译安装
make && make install

#修改nginx配置文件使其支持php-fpm服务。
cat > /apps/nginx/conf/nginx.conf << EOF
user  nginx;
worker_processes  1;

pid        /run/nginx.pid;
events {
    worker_connections  1024;
    }

http {
    include       mime.types;
    default_type  application/octet-stream;

sendfile        on;

 keepalive_timeout  65 65;
 server_tokens off;

  server {
  listen       80;                                                                                                                                                                                                                                                                                                          
  server_name  10.0.0.18;
  charset utf-8;

  location / {
    error_page   500 502 503 504 404  /50x.html;
    root  /apps/nginx/html/;
    index index.php index.html index.htm;
  }

  location = /50x.html {
    root   html;
    }

  location ~ \.php$ {
        root          /apps/nginx/html/;
        fastcgi_pass   127.0.0.1:9000;
        fastcgi_index  index.php;
        fastcgi_param  SCRIPT_FILENAME  \$document_root\$fastcgi_script_name;
        include        fastcgi_params;
   }
 }
}
EOF

#创建测试php页面代码
cat > /apps/nginx/html/test.php << EOF
<?php
phpinfo();
?>
EOF

#编写nginx.service文件通过systemd来管理nginx
cat > /usr/lib/systemd/system/nginx.service << EOF
[Unit]
Description=The nginx HTTP and reverse proxy server
After=network.target remote-fs.target nss-lookup.target

[Service]
Type=forking
PIDFile=/run/nginx.pid
# Nginx will fail to start if /run/nginx.pid already exists but has the wrong
# SELinux context. This might happen when running `nginx -t` from the cmdline.
# https://bugzilla.redhat.com/show_bug.cgi?id=1268621
ExecStartPre=/usr/bin/rm -f /run/nginx.pid
ExecStartPre=/apps/nginx/sbin/nginx -t
ExecStart=/apps/nginx/sbin/nginx
ExecReload=/bin/kill -s HUP \$MAINPID
KillSignal=SIGQUIT
TimeoutStopSec=5                                                                                                    
KillMode=process
PrivateTmp=true

[Install]
WantedBy=multi-user.target
EOF

#启动nginx设置问开机启动
systemctl daemon-reload
systemctl enable --now  nginx

#查看端口80已经打开
ss -ntl|grep 80 && echo "nginx安装完毕"
