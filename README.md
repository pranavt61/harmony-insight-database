Harmony Explorer Backend
=======

What is the Harmony ONE Explorer?
--------------

[Harmony](https://www.harmony.one/) is a fast and secure blockchain for decentralized applications.
Our production mainnet supports 4 shards of 1000 nodes, producing blocks in 5 seconds with finality.

[The Explorer](http://54.187.20.215/:8080/#/) allows users from all backgrounds to view and explorer the Harmony Blockchain by
making visible the latest blocks, pending transactions, top validators, and much more!

## Requirments

1) Install latest version of [Go](https://golang.org/dl/)

2) Install [mySQL](https://www.mysql.com/) database
  - Default location is localhost:3306 (Change in './sql/ManageConnection.go')
  - Default user in root:password@tcp (Change in './sql/ManageConnection.go')

3) Make sure Harmony's nodes are up and running. Check [API guide](https://api.hmny.io/)

4) Made sure port 8081 is unused. 
  - Used by dataServer (Change in './dataServer/Server.go')

## Build Setup

```bash
# Run SQL set up in mySQL CLI
source {PROJECT_DIR}/sql/CreateDatabase.sql
```

```bash
# Build source
go build

# Run server
./harmony-insight-database
```

## Harmony Docs

Please see our [user guide](https://docs.harmony.one/home/).

Please see our [API guide](https://api.hmny.io/)
