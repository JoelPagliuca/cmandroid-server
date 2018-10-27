```
┏━╸┏┳┓┏━┓┏┓╻╺┳┓┏━┓┏━┓╻╺┳┓   ┏━┓┏━╸┏━┓╻ ╻┏━╸┏━┓
┃  ┃┃┃┣━┫┃┗┫ ┃┃┣┳┛┃ ┃┃ ┃┃╺━╸┗━┓┣╸ ┣┳┛┃┏┛┣╸ ┣┳┛
┗━╸╹ ╹╹ ╹╹ ╹╺┻┛╹┗╸┗━┛╹╺┻┛   ┗━┛┗━╸╹┗╸┗┛ ┗━╸╹┗╸
```

Farm out adb commands using this golang server.

## Getting Started

### Requirements
Android sdk tools ([download](https://developer.android.com/studio/#downloads))

### Installing
Download the latest binary [release](https://github.com/JoelPagliuca/cmandroid-server/releases/latest)

or if you have `go` installed:
```sh
go get github.com/JoelPagliuca/cmandroid-server
```

### Running
First connect some android devices

Then run the server
```
cmandroid server
version: v1.0.0 commit: 73dad32

  -adb string
        Path to adb (default "adb")
  -h    Display this help text
  -p int
        Port to run the server on (default 8080)
```

## Development

### Build Status
| Branch | Status |
| ------ | ------ |
| master | [![Build Status](https://travis-ci.org/JoelPagliuca/cmandroid-server.svg?branch=master)](https://travis-ci.org/JoelPagliuca/cmandroid-server) |
| develop | [![Build Status](https://travis-ci.org/JoelPagliuca/cmandroid-server.svg?branch=develop)](https://travis-ci.org/JoelPagliuca/cmandroid-server) |

### Tasks

All tasks are run through the [`makefile`](./makefile)

```
build                          Builds an executable for host platform
clean                          Cleanup any build binaries or packages
cross                          Builds the cross-compiled binaries
docs                           Builds documentation
help                           Print this message and exit
test                           Runs the go tests
vet                            Verifies `go vet` passes
```

[More docs](./docs)

## License

This project is licensed under the MIT License - see the [LICENSE.md](./LICENSE.md) file for details

