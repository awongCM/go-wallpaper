# go-wallpaper

A simple program that switches your OS desktop wallpaper in Go.

## Installation

Do a `git clone` on this project url on your machine.

It is assumed you have the latest Go runtime libraries in your system in order to run this program.

## Usage

1. In the project folder, run `go build -o bin/go-wallpaper .` to compile
2. Run `./bin/go-wallpaper`

On startup the program tries to read your current wallpaper path. When you stop it with Ctrl+C (SIGINT) or SIGTERM, it restores that path if it was saved successfully.

Restore is skipped when the program cannot read the current wallpaper (for example `gsettings` errors) or when the desktop uses a non-file wallpaper (GNOME `resource://` URIs, solid colors, etc.).

## OS Version Supported:

- [x] Mac OS
- [x] Linux (GNOME with `gsettings` only)
- [ ] Windows - BOO!(why are you so cumbersome..)

## Development

```sh
go test ./...
```

## TODOS

* Add integration tests with mocked `exec` for darwin/gnome backends

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
