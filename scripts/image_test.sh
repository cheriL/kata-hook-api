#!/bin/bash

set -e
set -x

sudo mkdir -p /mnt/disk

## image file
img_file=./kata-containers-hooktest.img

sudo losetup  -f $img_file

dev=$(losetup  | grep kata-containers | awk '{ print $1}')
f=$(basename $dev)

sudo kpartx -a $dev
sudo mount /dev/mapper/${f}p1 /mnt/disk

sudo rm -rf /mnt/disk/usr/share/oci/hooks
sudo mkdir -p /mnt/disk/usr/share/oci/hooks/prestart

sudo cp -f ./prestart /mnt/disk/usr/share/oci/hooks/prestart/prestart

sudo umount /mnt/disk/
sudo kpartx -d $dev
sudo losetup -d $dev
sudo losetup
