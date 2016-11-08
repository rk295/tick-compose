#!/usr/bin/env bash
# set -x

function runInflux {

    dataDir="influx-data"

    if [[ ! -d "$dataDir" ]]; then
        mkdir $dataDir
    fi

    docker run \
           -d \
           -p 8083:8083 \
           -p 8086:8086 \
           -v "${PWD}/$dataDir/:/var/lib/influxdb" \
           --name influxdb \
           influxdb
}

function runTelegraf {
    docker run \
           -d \
           -v "${PWD}/telegraf.conf:/etc/telegraf/telegraf.conf:ro" \
           --link influxdb:influxdb \
           --name telegraf \
           telegraf
}

function runKapacitor {
    docker run \
           -d \
           -p 9092:9092 \
           -v "${PWD}/kapacitor.conf:/etc/kapacitor/kapacitor.conf:ro" \
           --link influxdb:influxdb \
           --name kapacitor \
           kapacitor
}

function runContainer {
    case "$1" in
        influxdb)
            runInflux
            ;;
        telegraf)
            runTelegraf
            ;;
        kapacitor)
            runKapacitor
            ;;
    esac
}

function killContainer {
    containerId="$1"

    if [[ -n "$containerId" ]]; then
        echo "Killing container $containerId"
        docker kill "$containerId"
        docker rm -f "$containerId"
    fi

}

function getContainerId {
    containerId=$(docker ps --filter name="$1" -q)
    echo "$containerId"
}

function stopContainers {

    for c in kapacitor telegraf influxdb; do
        killContainer "$(getContainerId $c)"
    done

}

function logs {
    docker logs -f "$(getContainerId "$1")"
}

function status {
    for c in influxdb kapacitor telegraf; do

        echo -n "Checking status of $c - "
        id=$(getContainerId "$c")

        if [[ -n "$id" ]]; then
            echo "running ($id)"
        else
            echo "stoped"
        fi
    done
}

function usage {
    echo "$0 [start|stop|status|logs]."
    echo ""
    echo "logs option takes one of two options: influxdb, kapacitor or telegraf"
    echo ""
    echo "* Kapacitor is at http://localhost:9092/"
    echo "* Influx is at http://localhost:8083/ and http://localhost:8086/"
}

action="$1"
case "$action" in
    start)
        for c in influxdb telegraf kapacitor; do
            containerId="$(runContainer $c)"
            echo "$c running in container id: $containerId"
        done
        ;;
    stop)
        stopContainers
        ;;
    status)
        status
        ;;
    logs)
        shift
        if [[ -z "$1" ]]; then
            echo "Error: logs sub command requires a containr to get logs from."
            echo ""
            usage
            exit 1
        fi
        logs "$1"
        ;;
    *)
        usage
        exit 1
        ;;
esac
