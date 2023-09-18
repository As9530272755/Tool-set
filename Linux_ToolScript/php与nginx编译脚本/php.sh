#!/bin/bash
#
#********************************************************************
#Author:		zhangfeng
#QQ: 			553645084
#Date: 			2019-11-09
#FileName：		php.sh
#URL: 			https://blog.51cto.com/14451122
#Description：		The test script
#Copyright (C): 	2019 All rights reserved
#********************************************************************
USER=nginx
Uid_nginx=2000
PWD=`pwd`
if [ $? -ne 1 ];then
    useradd -s /sbin/nologin $USER -u $Uid_nginx
      else
          echo "用户存在"
	      exit 3
	        fi
yum install libxml2-devel bzip2-devel libmcrypt-devel -y
tar xvf php-7.2.24.tar.gz
cp -f ${PWD}/www.conf ${PWD}/php-7.2.24
cd php-7.2.24
./configure --prefix=/app/php \
--enable-mysqlnd \
--with-mysqli=mysqlnd \
--with-pdo-mysql=mysqlnd \
--with-openssl \
--with-freetype-dir \
--with-jpeg-dir \
--with-png-dir \
--with-zlib \
--with-libxml-dir=/usr \
--with-config-file-path=/etc \
--with-config-file-scan-dir=/etc/php.d \
--enable-mbstring \
--enable-xml \
--enable-sockets \
--enable-fpm \
--enable-maintainer-zts \
--disable-fileinfo
make  && make install
cp  -f php.ini-production /etc/php.ini
cp -f sapi/fpm/php-fpm.service /usr/lib/systemd/system/
cp -f ${PWD}/www.conf /app/php/etc/php-fpm.d
cd /app/php/etc/
cp php-fpm.conf.default php-fpm.conf
systemctl daemon-reload 
systemctl start php-fpm.service

