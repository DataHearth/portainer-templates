# Portainer-templates

## What is it

Portainer-templates is a small HTTP API that aims to ease the template management from `Portainer`.  
It provides a valid HTTP endpoint for `Portainer`, so it can retrieve templates and display it on its UI.  
You can add new templates but triggering a special endpoint and it will be persisted in a SQLite database.

## Endpoints

- `GET /templates` - Portainer endpoint for all templates
- `GET /templates/{type}/{id}` - Get a special template by its database id and type (`container`, `compose`, `stack`)
- `POST /templates/load` - Load templates from a `JSON` file
- `POST /templates/insert` - Insert templates from a `JSON` body
- `/metrics` - Prometheus metrics

*note*: To add templates, you need to pass a valide JSON format (obviously) but also a valid template format
e.g:
```
{
	"version": "2",
	"templates": [{...}]
}
```

## Environment variables

You can customize your `portainer-templates` by setting environment variables:

- `LOG_LEVEL` - Which log level your instance should use (`trace`, `debug`, `info`, `warn`, `error`, `fatal`, `panic`) (default: `info`)
- `PORT` - Which port your instance should listen to (default: `4345`)
- `HOST` - Which address your instance should listen to (default: `0.0.0.0`)
- `DB_FILE` - Where should the database be located (default: `exectuable_workdir/portainer-templates.db`)
- `TEMPLATE_FILE` - Where is located the template file for populating the database. It is mandatory if you want to use the `/templates/load` endpoint

## How to use it

### Build and run

```
git clone https://github.com/DataHearth/portainer-templates/tree/v1.1.0
cd portainer-templates
go build -o portainer-templates cmd/main.go
```
Replace `v1.1.0` by the desired version.

### Go run

```
git clone https://github.com/DataHearth/portainer-templates/tree/v1.1.0
cd portainer-templates
go run cmd/main.go
```
Replace `v1.1.0` by the desired version.

### Docker

```
docker run -e PORT=5555 -e DB_FILE=/data/demo.db -e TEMPLATE_FILE=data/demo.json -e LOG_LEVEL=debug -p 5555:5555 --volume /path/to/data/test:/data:rw datahearth/portainer-templates:latest
```
Replace `latest` by the desired version.

## TO-DO

- Delete templates (by id, by title)
- Update template (by id, by title)
- Add more logs