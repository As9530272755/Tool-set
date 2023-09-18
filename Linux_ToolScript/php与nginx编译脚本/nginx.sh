#!/bin/bash
USER=nginx
Uid_nginx=2000
#下载依赖包
yum install -y gcc gcc-c++ glibc glibc-devel pcre pcre-devel openssl openssl-devel systemd-devel net-tools iotop bc zip unzip zlib-devel bash-completion nfs-utils automake libxml2 libxml2-devel libxslt libxslt-devel perl perl-ExtUtils-Embed
cd /usr/local/src/
#下载nginx源码
wget https://nginx.org/download/nginx-1.16.1.tar.gz
tar vxf  nginx-1.16.1.tar.gz
cd nginx-1.16.1/
#生成编译环境
./configure --prefix=/apps/nginx --user=nginx --group=nginx --with-http_ssl_module --with-http_v2_module --with-http_realip_module --with-http_stub_status_module --with-http_gzip_static_module --with-pcre --with-stream --with-stream_ssl_module --with-stream_realip_module
#编译
make && make install
#检测用户nginx
if [ $? -ne 1 ];then
  useradd -s /sbin/nologin $USER -u $Uid_nginx
  else echo"用户存在" 
  exit 9
fi
#支持php
cat > /apps/nginx/conf/nginx.conf << EOF
user  nginx;
worker_processes  1;

events {
    worker_connections  1024;
    }

http {
    include       mime.types;
    default_type  application/octet-stream;

sendfile        on;

 keepalive_timeout  65;

  server {
            listen       80;
	    server_name  localhost;

  location / {
	  root   html;
  	  index index.php index.html index.htm;
  }

  error_page   500 502 503 504  /50x.html;
	location = /50x.html {
		root   html;
			    }

  location ~ \.php$ {
	root          /apps/nginx/html;
	fastcgi_pass   127.0.0.1:9000;
    	fastcgi_index  index.php;
	fastcgi_param  SCRIPT_FILENAME  \$document_root\$fastcgi_script_name;
	include        fastcgi_params;
   }
 }
}
EOF
/apps/nginx/sbin/nginx -V
/apps/nginx/sbin/nginx
echo "已经支持php功能默认路径再/apps/nginx/html下"
