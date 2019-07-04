#!/usr/bin/env bash

#sudo apt-get update

wget https://www.openvswitch.org/releases/openvswitch-2.11.1.tar.gz

ls -ltr

tar -xvf openvswitch-2.11.1.tar.gz

cd openvswitch-2.11.1

sudo su

apt-get install python-simplejson python-qt4 libssl-dev python-twisted-conch automake autoconf gcc uml-utilities libtool build-essential pkg-config

uname -r

sudo apt-cache search linux-headers

apt-get install linux-headers-4.9.0-6-rpi

./configure --with-linux=/lib/modules/4.9.0-6-rpi/build

make && make install

cd datapath/linux

modprobe openvswitch

cat /etc/modules

echo "openvswitch" >> /etc/modules

cat /etc/modules

cd ../..

touch /usr/local/etc/ovs-vswitchd.conf

mkdir -p /usr/local/etc/openvswitch

ovsdb-tool create /usr/local/etc/openvswitch/conf.db vswitchd/vswitch.ovsschema

sudo /usr/local/share/openvswitch/scripts/ovs-ctl start

cat > script << EOF1
#! /bin/bash

ovsdb-server --remote=punix:/usr/local/var/run/openvswitch/db.sock --remote=db:Open_vSwitch,Open_vSwitch,manager_options --private-key=db:Open_vSwitch,SSL,private_key --certificate=db:Open_vSwitch,SSL,certificate --bootstrap-ca-cert=db:Open_vSwitch,SSL,ca_cert --pidfile --detach

ovs-vswitchd --pidfile --detach
ovs-vsctl --no-wait init
ovs-vsctl show

EOF1

sudo ./script

## TODO() If you want to remove that version of openvswitch

#sudo /usr/local/share/openvswitch/scripts/ovs-ctl stop
#sudo ovs-dpctl del-dp ovs-system
#sudo modprobe -r openvswitch


#sudo ovs-vsctl add-br br0
#sudo ovs-vsctl add=port br0 eth1

#set the sdn controller using
#sudo ovs-vsctl set-controller br0 tcp:192.168.1.205:6633
#
#sudo ovs-vsctl show