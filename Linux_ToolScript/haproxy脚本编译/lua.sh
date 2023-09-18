#!/bin/bash
#
#********************************************************************
#Author:		zhangfeng
#QQ: 			553645084
#Date: 			2019-11-08
#FileName：		lua.sh
#URL: 			https://blog.51cto.com/14451122
#Description：		The test script
#Copyright (C): 	2019 All rights reserved
#********************************************************************
PWD=`pwd`
cp ${PWD}/lua-5.3.5.tar.gz /usr/local/src/
cd /usr/local/src/
tar xvf	lua-5.3.5.tar.gz
cd lua-5.3.5
make linux test
/usr/local/src/lua-5.3.5/src/lua -v
