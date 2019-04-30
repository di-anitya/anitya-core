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

echo "** building the source codes.."

cd ${BASE_DIR}/bin
go build -o api-server ${BASE_DIR}/src/api_server/*.go
#go build ${BASE_DIR}/src/grpc_client/*.go
#go build ${BASE_DIR}/src/notification/*.go
#go build ${BASE_DIR}/src/scheduler/*.go


if [ $(get_os_distribution) == "redhat" -o $(get_os_distribution) == "fedora"  ]; then
  echo "** define anitya-api.service"

  \cp -fp ${BASE_DIR}/files/usr/lib/systemd/system/anitya-api.service /usr/lib/systemd/system/anitya-api.service

  echo "** reloading daemon"
  systemctl daemon-reload

  echo "** starting anitya-api.service"
  systemctl start anitya-api.service

  echo "** enabling anitya-api.service"
  systemctl enable anitya-api.service
fi
