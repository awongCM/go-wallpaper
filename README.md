# go-wallpaper

A simple program that switches your OS desktop wallpaper in Go.

## Installation

Do a `git clone` on this project url on your machine.

It is assumed you have the latest Go runtime libraries in your system in order to run this program.

## Usage

1. In the project folder, run `go build -o bin/go-wallpaper .` to compile
2. Run `./bin/go-wallpaper`

On startup the program saves your current wallpaper path. When you stop it with Ctrl+C (SIGINT) or SIGTERM, it restores that wallpaper.

## OS Version Supported:

- [x] Mac OS
- [x] Linux (GNOME with `gsettings` only)
- [ ] Windows - BOO!(why are you so cumbersome..)

## Development

```sh
go test ./...
```

## TODOS

* Add some decent unit testing tools (integration tests with mocked exec)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
