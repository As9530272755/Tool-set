#!/bin/bash
#
#********************************************************************
#Author:		zhangfeng
#QQ: 			553645084
#Date: 			2019-11-08
#FileName：		keepalived.sh
#URL: 			https://blog.51cto.com/14451122
#Description：		The test script
#Copyright (C): 	2019 All rights reserved
#********************************************************************
PWD=`pwd`
cp -f ${PWD}/keepalived-2.0.18.tar.gz /usr/local/src/
mkdir -p /etc/keepalived
cp -f  ${PWD}/keepalived.conf /etc/keepalived
cd /usr/local/src/
tar xvf keepalived-2.0.18.tar.gz
cd keepalived-2.0.18/
yum  -y install libnfnetlink-devel libnfnetlink ipvsadm  libnl libnl-devel  libnl3 libnl3-devel   lm_sensors-libs net-snmp-agent-libs net-snmp-libs  openssh-server openssh-clients  openssl openssl-devel automake iproute 
mkdir -p /usr/local/keepalived
./configure --prefix=/usr/local/keepalived --disable-fwmark
make && make install
cp /usr/local/src/keepalived-2.0.18/keepalived/etc/init.d/keepalived.rh.init /etc/sysconfig/keepalived.sysconfig
cp  /usr/local/src/keepalived-2.0.18/bin/keepalived /usr/sbin/
systemctl daemon-reload
systemctl start keepalived
ip a
