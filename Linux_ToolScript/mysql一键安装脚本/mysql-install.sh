#!/bin/bash
DIR=`pwd`
NAME="mysql-5.6.34-linux-glibc2.5-x86_64.tar.gz"
FULL_NAME=${DIR}/${NAME}
DATA_DIR="/data/mysql"

yum install vim gcc gcc-c++ wget autoconf  net-tools lrzsz iotop lsof iotop bash-completion -y
yum install curl policycoreutils openssh-server openssh-clients postfix -y

if [ -f ${FULL_NAME} ];then
    echo "安装文件存在"
else
    echo "安装文件不存在"
    exit 3
fi
if [ -h /usr/local/mysql ];then
    echo "Mysql 已经安装"
    exit 3 
else
    tar xvf ${FULL_NAME}   -C /usr/local/src
    ln -sv /usr/local/src/mysql-5.6.34-linux-glibc2.5-x86_64  /usr/local/mysql
    if id  mysql;then
        echo "mysql 用户已经存在，跳过创建用户过程"
    fi
        useradd  mysql  -s /sbin/nologin
    if  id  mysql;then
    	chown  -R mysql.mysql  /usr/local/mysql/* -R
        if [ ! -d  /data/mysql ];then
            mkdir -pv /data/mysql && chown  -R mysql.mysql  /data   -R
            /usr/local/mysql/scripts/mysql_install_db  --user=mysql --datadir=/data/mysql  --basedir=/usr/local/mysql/
	    cp  /usr/local/src/mysql-5.6.34-linux-glibc2.5-x86_64/support-files/mysql.server /etc/init.d/mysqld
	    chmod a+x /etc/init.d/mysqld
 	    cp ${DIR}/my.cnf   /etc/my.cnf
	    ln -sv /usr/local/mysql/bin/mysql  /usr/bin/mysql
	    /etc/init.d/mysqld start
	else
            echo "MySQL数据目录已经存在,"
			exit 3
	fi
    fi
fi

