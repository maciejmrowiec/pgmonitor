## **PGMonitor** [![Build Status](https://drone.io/github.com/maciejmrowiec/pgmonitor/status.png)](https://drone.io/github.com/maciejmrowiec/pgmonitor/latest)

New Relic plugin for monitoring PostgreSQL.

#### Features

###### Monitoring tuple distribution within tables (active tuples, dead tuples, free space). Intended to monitor table bloat in detail.

* **Component/tuple/table/`table_name`/percent/active** relation of live tuples to total table size
* **Component/tuple/table/`table_name`/percent/dead** relation of dead tuples to total table size
* **Component/tuple/table/`table_name`/percent/free** relation of free space to total table size
* **Component/tuple/summary/percent/active** live tuples to total table size; average over all tables in database
* **Component/tuple/summary/percent/dead** dead tuples to total table size; average over all tables in database
* **Component/tuple/summary/percent/free** free space to total table size; average over all tables in database

###### Monitoring table and index size on disk.

* **Component/disksize/table/`table_name`/table** disk space used by table; expressed in Bytes
* **Component/disksize/table/`table_name`/index** disk space used by table indexes; expressed in Bytes
* **Component/disksize/summary/table/`table_name`** disk space used by table; expressed in Bytes
* **Component/disksize/summary/index/`table_name`** disk space used by table indexes; expressed in Bytes


#### Installation

###### Dependencies

Requires golang toolchain.

```
sudo apt-get install golang
```

###### Build

```
go get github.com/lib/pq
go get github.com/yvasiyarov/newrelic_platform_go
go get github.com/maciejmrowiec/pg_monitor
go build
```

#### Usage

**Usage of pgmonitor:**
* -database="": Database name (required)
* -interval=1: Sampling interval [min]
* -key="": Newrelic license key (required)
* -user="postgres": Database user name
* -verbose=false: Verbose mode

To deamonize in backgrund you can use:

```
nohup ./pgmonitor -database=<mydatabase> -key=<my_newrelic_key> >/dev/null 2>&1 &
```

**Note:** PGMonitor requires [pgstattuple](http://www.postgresql.org/docs/9.3/static/pgstattuple.html) extension to work correctly.

To enable execute in postgres shell:

```
CREATE EXTENSION pgstattuple;
```
