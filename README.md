<p align="center">
  <img width="200" src="teamplanner-spa/logo-default.png"></img>
  <div style="font-weight: bold; text-align: center">Quickly get an overview of you team's schedule and lineup!</div>
</p>
<hr>

## Overview

Teamplanner stores your teams lineup and matches and provides an easy method to record your teammates' availability for each match.

It consists of a RESTful API written in Go and backed by [buntdb](https://github.com/tidwall/buntdb). It can be used with the included [PWA](/teamplanner-spa).

## Quickstart

Build the container with
```
docker build -t teamplanner .
```
and run it:
```
docker run -dp 8042:8042 --name teamplanner teamplanner
```

## Configuration

### Networking
The app listens on `:8042` by default and port 8042 is exposed from the container. You can override it by setting the `LISTENADDR` environment variable. 
You should use a reverse proxy to handle HTTP termination if you want to expose it to the internet.

### Persistence
The database is by default stored at `/data/teamplanner.db`. Mount `/data` to a volume or use a bind mount to persist it over container restarts.

You can also change the location of the database by setting the environment variable `DBPATH` to the absolute path of the file.

Technically it's also possible to store the database completely in memory by setting `DBPATH` to `:memory:`.
