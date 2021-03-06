# Dcrextdata
- Maintains an up-to-date database of ticker data from various dcr exchanges
- Collects and store mining pool stats

## Requirements
- [Go](golang.org/dl) 1.11
- [Postgresql](postgresql.org/download)

## Building

- [Install Go](http://golang.org/doc/install)

- Verify Go installation:

      go env GOROOT GOPATH

- Ensure `$GOPATH/bin` is on your `$PATH`.

- Clone the dcrextdata repository. It is conventional to put it under `GOPATH`, but
  this is no longer necessary with go module.

  ```sh
  git clone https://github.com/raedahgroup/dcrextdata $GOPATH/src/github.com/raedahgroup/dcrextdata
  ```

 - If building inside GOPATH it is nessasary to set `GO111MODULE=ON` before building otherwise just build.

 - `go build`, run in the dcrextdata repo, builds the executabe `dcrextdata`

## Configuring `dcrextdata`
`dcrextdata` can be configured via command-line options or a config file located in the same diretcory as the executable. Start with the sample config file:
```sh
cp sample-dcrextdata.conf dcrextdata.conf
```
Then edit `dcrextdata.conf` with your postgres settings.  See the output of `dcrextdata --help`
for a list of all options and their default values.

## Running `dcrextdata`
Simply run `dcrextdata` with your flags in the same directory as it's config file and you're good to go. You can perform a reset by running with the `-R` or `--reset` flag.

## Quick start for  Postgresql
If you have a new postgresql install and you want a quick setup for dcrextdata, you can start postgresql command-line client with `sudo -u postgres psql` or you could `su` into the postgres user and run `psql` then execute the following sql statements to create a user and database:
```sql
    CREATE USER {username} WITH PASSWORD '{password}' CREATEDB;
    CREATE DATABASE {databasename} OWNER {user};
```







