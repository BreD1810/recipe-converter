Recipe Converter
===

Converts online recipes to a markdown format to be saved as a note.

I store all my recipes in [Joplin](https://joplinapp.org), for access across all of my devices.
Other markdown editors are available.

## CLI Tool
To install:
```shell
go install github.com/BreD1810/recipe-converter/cmd/convert-recipe@latest
```
or to install from the repository:
```shell
make install-cli
```

To run:
```shell
convert-recipe <recipe-url>
```
or from the repository:
```shell
make build-cli
./convert-recipe <recipe-url>
```
