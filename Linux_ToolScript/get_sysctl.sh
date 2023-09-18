#!/bin/bash
Dir="./sysctl"
Kernel_parameters="vm.max_map_count \
fs.inotify.max_user_instances \
fs.inotify.max_user_watches \
kernel.pid_max \
fs.pipe-max-size \
net.core.netdev_max_backlog \
net.core.rmem_max \
net.core.wmem_max \
net.ipv4.tcp_max_syn_backlog \
net.ipv4.neigh.default.gc_thresh1 \
net.ipv4.neigh.default.gc_thresh2 \
net.ipv4.neigh.default.gc_thresh3 \
net.core.somaxconn \
fs.aio-max-nr \
net.ipv4.tcp_max_tw_buckets \
net.ipv4.udp_rmem_min \
net.ipv4.udp_wmem_min \
net.ipv4.conf.all.rp_filter \
net.ipv4.conf.default.rp_filter \
net.ipv4.conf.eth0.arp_accept"


MkFile() {
    if [ ! -d "$Dir" ] ; then
        mkdir $Dir
    else 
        echo "$Dir exist!"
        exit 0
    fi
}

List() {
    for i in $Kernel_parameters ; do
        sysctl "$i" >> "$Dir"/now_sysctl.txt
    done
    echo "当前现有内核参数值如下:"
    cat "$Dir"/now_sysctl.txt
}
MkFile
List