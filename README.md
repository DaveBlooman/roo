# Docker Tool

### Installation

Download from releases page or building use Go

To build locally
```sh
make build
```

### Usage

```sh
Usage: deliveroo [--version] [--help] <command> [<args>]

Available commands are:
    deploy
    stop
```

 - Deploy command accepts a hash from the github repo - awdaw.  It will clone the code from that hash locally, build the image and create a docker container for the app to run in and link to a postgres container.

 - Stop command will completely remove the containers created.

### Tests

```sh
make tests
```
