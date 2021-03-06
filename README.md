Influx Stack Playground
=======================

This docker compose file brings up a full Influx db test stack consisting of:

* InfluxDB
* Kapacitor
* Telegraf
* Grafana

To bring the stack up, simply run:

    docker-compose up

Which will run all the services in the foreground, with logs going to your terminal, or:

    docker-compose up -d

To background them and get your terminal back.

## Grafana Setup

The first time you run it, you need to add a datasource to Grafana. The `grafana-setup` script in this directory will do that for you, simply run:

    ./grafana-setup

Once the services have started.

## Service URLs

The services will be available on the following URL:

* InfluxDB Simple Web UI [http://localhost:8083/](http://localhost:8083/)
* Grafana [http://localhost:3000/](http://localhost:3000/) (admin/admin login)
* Kapacitor Rest API [http://localhost:9092/](http://localhost:9092/)

## Documentation

* [Kapacitor](https://docs.influxdata.com/kapacitor/v1.1//)
* [Telegraf](https://docs.influxdata.com/telegraf/v1.1/)
* [Telegraf Plugins](https://github.com/influxdata/telegraf/tree/master/plugins/inputs)
* [InfluxDB](https://docs.influxdata.com/influxdb/v1.0/)
