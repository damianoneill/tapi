#!/bin/sh

mkdir -p OAS

# https://github.com/mbj4668/yanger
docker run -it -v "${PWD}":/workdir mbj4668/yanger -p TAPI/YANG -t expand -f swagger TAPI/YANG/tapi-topology@2019-03-31.yang -o OAS/tapi-topology.json
docker run -it -v "${PWD}":/workdir mbj4668/yanger -p TAPI/YANG -t expand -f swagger TAPI/YANG/tapi-connectivity@2019-03-31.yang -o OAS/tapi-connectivity.json
docker run -it -v "${PWD}":/workdir mbj4668/yanger -p TAPI/YANG -t expand -f swagger TAPI/YANG/tapi-equipment@2019-03-31.yang -o OAS/tapi-equipment.json

# non-standard service
docker run -it -v "${PWD}":/workdir mbj4668/yanger -p TAPI/YANG -t expand -f swagger tapi-network-element@2019-03-31.yang -o OAS/tapi-network-element.json

# https://github.com/go-swagger/go-swagger
swagger generate model -f OAS/tapi-connectivity.json
swagger generate model -f OAS/tapi-equipment.json
swagger generate model -f OAS/tapi-topology.json
swagger generate model -f OAS/tapi-network-element.json

go build ./...
