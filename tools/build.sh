#!/bin/sh

get_os_distribution() {
    if [ -e /etc/fedora-release ]; then
        distri_name="fedora"
    elif [ -e /etc/redhat-release ]; then
        distri_name="redhat"
    else
        distri_name="unkown"
    fi
    echo ${distri_name}
}

SCRIPT_DIR=$(cd $(dirname $0); pwd)
BASE_DIR=${SCRIPT_DIR}/../

cd ${BASE_DIR}/bin
go build ${BASE_DIR}/src/*.go

if [ $(get_os_distribution) == "redhat" -o $(get_os_distribution) == "fedora"  ]; then
  cp -p ${BASE_DIR}/files/usr/lib/systemd/system/anitya-api.service /usr/lib/systemd/system/anitya-api.service
  cd /etc/systemd/system/
  ln -s /usr/lib/systemd/system/anitya-api.service anitya-api.service
fi
