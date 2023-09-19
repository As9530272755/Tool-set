#!/bin/bash
# 该脚本用于修改内核参数

ConfiFile="/etc/sysctl.d/kernel.conf"

# 修改配置文件
Edit_Config() {

cat > "$ConfiFile" << EOF
vm.max_map_count = 262144
fs.inotify.max_user_instances=8192
fs.inotify.max_user_watches=524288
kernel.pid_max = 65535
fs.pipe-max-size=4194304
net.core.netdev_max_backlog=65535
net.core.rmem_max = 33554432
net.core.wmem_max = 33554432
net.ipv4.tcp_max_syn_backlog = 1048576
net.ipv4.neigh.default.gc_thresh1= 512
net.ipv4.neigh.default.gc_thresh2 = 2048
net.ipv4.neigh.default.gc_thresh3 = 4096
net.core.somaxconn = 32768
fs.aio-max-nr=262144
net.ipv4.tcp_max_tw_buckets = 1048576
net.ipv4.udp_rmem_min = 131072
net.ipv4.udp_wmem_min = 131072
EOF

}

Run_sysctl() {
    sysctl -p "$ConfiFile"
}

Edit_Config
Run_sysctl