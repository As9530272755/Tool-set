#!/bin/bash
#
#********************************************************************
#Author:		zhangfeng
#QQ: 			553645084
#Date: 			2019-11-08
#FileName：		haproxy.sh
#URL: 			https://blog.51cto.com/14451122
#Description：		The test script
#Copyright (C): 	2019 All rights reserved
#********************************************************************
PWD=`pwd`
cp ${PWD}/haproxy.service /usr/lib/systemd/system/
cp ${PWD}/haproxy-2.0.8.tar.gz /usr/local/src/
mkdir -p /etc/haproxy
cp ${PWD}/haproxy.cfg /etc/haproxy
cd /usr/local/src
yum install gcc gcc-c++ glibc glibc-devel pcre pcre-devel openssl	openssl-devel systemd-devel net-tools vim iotop bc	zip unzip zlib-devel lrzsz tree screen lsof tcpdump wget ntpdate  -y
tar vxf  haproxy-2.0.8.tar.gz
cd haproxy-2.0.8/
make ARCH=x86_64 TARGET=linux-glibc USE_PCRE=1 USE_OPENSSL=1 USE_ZLIB=1 USE_SYSTEMD=1 USE_CPU_AFFINITY=1 USE_LUA=1 LUA_INC=/usr/local/src/lua-5.3.5/src/ LUA_LIB=/usr/local/src/lua-5.3.5/src/ PREFIX=/usr/local/haproxy
make install PREFIX=/usr/local/haproxy
./haproxy -v
cp haproxy /usr/share/
cp haproxy /usr/sbin/
mv /usr/share/haproxy /tmp/
mkdir -p /var/lib/haproxy
systemctl daemon-reload
systemctl start haproxy
ss -nutl|grep 9999
