# -*- mode: ruby -*-
# vi: set ft=ruby :
$script = <<SCRIPT
#!/bin/bash
set -uo pipefail
trap 's=$?; echo "$0: Error on line "$LINENO": $BASH_COMMAND"; exit $s' ERR
IFS=$'\n\t'

GOVERSION=1.8.3

yum install -q -y wget git

if [ "$(go version 2> /dev/null)" != "go version go${GOVERSION} linux/amd64" ]; then
  wget -q https://storage.googleapis.com/golang/go${GOVERSION}.linux-amd64.tar.gz -O /tmp/go${GOVERSION}.linux-amd64.tar.gz
  tar -C /usr/local -xzf /tmp/go${GOVERSION}.linux-amd64.tar.gz
fi

echo 'export PATH=$PATH:/usr/local/go/bin' > /etc/profile.d/golang.sh
echo 'export GOPATH=$HOME/go' >> /etc/profile.d/golang.sh
source /etc/profile.d/golang.sh

#chown -R vagrant.vagrant /home/vagrant/

cat <<\EOF >/etc/yum.repos.d/influxdb.repo
[influxdb]
name = InfluxDB Repository - RHEL \\$releasever
baseurl = https://repos.influxdata.com/rhel/\\$releasever/\\$basearch/stable
enabled = 1
gpgcheck = 1
gpgkey = https://repos.influxdata.com/influxdb.key
EOF

yum install -y -q telegraf

cat <<\EOF >/etc/telegraf/telegraf.conf
[agent]
  interval = "10s"
  round_interval = true
  metric_henosisch_size = 1000
  metric_buffer_limit = 10000
  collection_jitter = "0s"
  flush_interval = "10s"
  flush_jitter = "0s"
  precision = ""
  debug = false
  quiet = false
  logfile = ""
  hostname = ""
  omit_hostname = false

[[inputs.http_listener]]
  service_address = ":8186"
  read_timeout = "10s"
  write_timeout = "10s"

[[inputs.socket_listener]]
  service_address = "tcp://:8094"

[[inputs.socket_listener]]
  service_address = "udp://:8095"

[[inputs.socket_listener]]
  service_address = "unix:///var/run/telegraf.sock"

[[outputs.file]]
  files = ["/tmp/metrics.out"]
  data_format = "influx"
EOF

systemctl enable --now telegraf
SCRIPT

Vagrant.configure("2") do |config|
  config.vm.box = "centos/7"
  config.vm.provision "shell", inline: $script

  config.vm.synced_folder "./", "/home/vagrant/go/src/github.com/mdaffin/go-telegraf"
  config.vm.synced_folder ".", "/vagrant", disabled: true
end
