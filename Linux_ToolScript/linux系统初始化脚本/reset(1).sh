#!/bin/bash
#
#********************************************************************
#Author:		张桂元
#QQ: 			953027255
#Date: 			2019-08-03
#FileName：		reset.sh
#URL: 			https://1717zgy.com
#Description：		The test script
#Copyright (C): 	2019 All rights reserved
#********************************************************************



#颜色配置
COLOR="\033[1;$[RANDOM%7+31]m"                                                                         
COLOREND='\033[0m'

echo -e "${COLOR}开机提示配置 $COLOREND"
cat > /etc/issue << EOF
\S
time is \t
tty is \l
hostname is \n 
Kernel \r on an \m
EOF
sleep 1

echo -e "${COLOR}别名环境变量配置 $COLOREND"
cat > /etc/profile.d/env.sh << EOF
alias df="df -h"
alias cdnet="cd /etc/sysconfig/network-scripts"
alias editnet="vim /etc/sysconfig/network-scripts/ifcfg-ens33"
alias scandisk="echo '- - -' > /sys/class/scsi_host/host2/scan;echo '- - -' > /sys/class/scsi_host/host0/scan" 
export PATH=/app/bin:$PATH
export EDITOR=vim
export HISTTIMEFORMAT="%F %T "
export HISTCONTROL=ignoreboth
EOF
#~/.bashrc

echo -e "${COLOR}提示符配置 $COLOREND"
echo 'PS1="\[\e[1;32m\][\u@\h \W]\\$\[\e[0m\]"' >> /etc/profile.d/env.sh
sleep 1 

echo -e "${COLOR}VIM配置$COLOREND"
cat > ~/.vimrc << EOF
set ignorecase
set cursorline
set autoindent
autocmd BufNewFile *.sh exec ":call SetTitle()"
func SetTitle()
	if expand("%:e") == 'sh'
	call setline(1,"#!/bin/bash") 
	call setline(2,"#") 
	call setline(3,"#********************************************************************") 
	call setline(4,"#Author:		张桂元") 
	call setline(5,"#QQ: 			953027255") 
	call setline(6,"#Date: 			".strftime("%Y-%m-%d"))
	call setline(7,"#FileName：		".expand("%"))
	call setline(8,"#URL: 			https://1717zgy.com")
	call setline(9,"#Description：		The test script") 
	call setline(10,"#Copyright (C): 	".strftime("%Y")." All rights reserved")
	call setline(11,"#********************************************************************") 
	call setline(12,"") 
	endif
endfunc
autocmd BufNewFile * normal G
EOF
sleep 1 

#关闭selinux
echo -e "${COLOR}selinux已经关闭 $COLOREND"
sed -i  's/SELINUX=enforcing/SELINUX=disabled/ ' /etc/selinux/config
sleep 1 

#把ens33改为eth0
echo -e "${COLOR}网卡名改为eth0 $COLOREND"
sed -ri '/GRUB_CMDLINE_LINUX=/s@"$@ net.ifnames=0"@' /etc/default/grub
sed -ri '/^[[:space:]]+linux16/s#(.*)#\1 net.ifnames=0#' /boot/grub2/grub.cfg
sleep 1 

#7版本关闭防火墙
echo -e "${COLOR}关闭7版本防火墙 $COLOREND"
systemctl disable firewalld.service
sleep 1 

#6版本关闭防火墙
echo -e "${COLOR}关闭6版本防火墙 $COLOREND"
chkconfig iptables off
sleep 1 

#主机时间同步
echo -e "${COLOR}主机时间和阿里云同步 $COLOREND"
ntpdate time1.aliyum.com

#配置阿里云YUM源适合7版本                                                            
echo -e "${COLOR}7版本阿里云repo源配置 $COLOREND"
#wget http://mirrors.aliyun.com/repo/Centos-7.repo
mkdir /etc/yum.repos.d/bak
mv /etc/yum.repos.d/*.repo /etc/yum.repos.d/bak/
#mv Centos-7.repo /etc/yum.repos.d/
#touch /etc/yum.repos.d/mirrorlist
#cat > /etc/yum.repos.d/mirrorlist <<EOF
#https://mirrors.aliyun.com/epel/$releasever/$basearch/
#EOF
touch /etc/yum.repos.d/epel.repo
cat > /etc/yum.repos.d/epel.repo <<EOF
[epel]
name=epel
baseurl=https://mirrors.aliyun.com/epel/7/x86_64/
gpgcheck=0
enabled=1
EOF
sleep 1 

echo -e "${COLOR}阿里云普通源配置 $COLOREND"
touch /etc/yum.repos.d/base.repo
cat > /etc/yum.repos.d/base.repo <<EOF
[base]
name=base
baseurl=file:///misc/cd
	https://mirrors.aliyun.com/centos/7/os/x86_64/                                                 
gpgcheck=0
EOF
sleep 1 

yum clean all


echo -e "${COLOR}自动下载软件 $COLOREND"
yum install tree -y
yum install ftp -y
yum install lftp -y
yum install telnet -y
yum install autofs gcc gcc-c++ glibc glibc-devel pcre pcre-devel openssl openssl-devel systemd-devel zlib-devel vim lrzsz tree screen lsof tcpdump wget ntpdate net-tools iotop bc zip unzip nfs-utils -y

sleep 1 


echo -e "${COLOR}服务器基本信息 $COLOREND"
echo -e  "CPU type is ${COLOR}`cat /proc/cpuinfo|grep -im1 'model name'|cut -d: -f2|tr -s ' '`$COLOREND"
echo -e "Mem is ${COLOR}`cat /proc/meminfo |head -1 | tr -s " " |cut -d: -f2`$COLOREND"
echo -e "Disk size is ${COLOR}`cat /proc/partitions |grep "sda$" |tr -s " " |cut -d" " -f4` kB$COLOREND"

echo -e "OS version is ${COLOR}`cat /etc/redhat-release`$COLOREND"
echo -e "Kernel version is ${COLOR}`uname -r`$COLOREND"
