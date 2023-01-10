# Development

## Update all Dependencies

```sh
go get -u ./...
```

## Update a Dependency

```sh
go get <dependency name>
```

### Example

```sh
go get github.com/spf13/cobra
```

## Build Releases

```sh
VERSION=0.4.0 ./build-releases.sh
```