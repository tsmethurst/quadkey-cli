# Quadkey cli tool

Very quick little CLI tool for converting lat/long coordinates into quadkeys.

Based on [this code](https://docs.microsoft.com/en-us/bingmaps/articles/bing-maps-tile-system).

## Build

```shell
go build ./...
```

### Usage

```shell
./quadkey encode -c 51.051509,3.739270 -l 16
```

Output:

```text
1202021230301202
```

Output help text with `quadkey --help`.
