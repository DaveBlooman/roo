# Dave's Docker Tool

This tool will retrieve the code for a given hash from the github repo - https://github.com/DaveBlooman/rubydockerapp.  This will build the hash and then create a docker image based on the Dockerfile in that repo.  A blank Postgres database will also be started along side this container and linked to the application container.  You can deploy different commit hashes, but you can only deploy a commit once.

### Requirements
  - Up to date version of Docker

### Installation

Download from releases page or building use Go

To build locally
```sh
make build
```

### Usage

```sh
Usage: roo [--version] [--help] <command> [<args>]

Available commands are:
    deploy
    stop
```

 - Deploy command accepts a hash from the github repo - https://github.com/DaveBlooman/go-app.  It will clone the code from that hash locally, build the image and create a docker container for the app to run in and link to a postgres container.

 - Stop command will completely remove the containers created.

### Tests

```sh
make tests
```
