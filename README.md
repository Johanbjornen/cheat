cheat
====

## Description
cheat sheet cli tool.

## Usage

Add a command and its description.
```bash
$ cheat add "show list of directories & files" "ls"
```

Show cheat list.

```bash
$ cheat ls
# => 1. show list of directories & files
# => ls
# =>
# => 2. remove directory
# => rm -r
# =>
```

Delete a cheat.

```bash
$ cheat rm 1
```

### Note
Cheat sheet is stored in `~/.cheat.json`.

## Install

To install, use `go get`:

```bash
$ go get github.com/takady/cheat
```

## Contribution

1. Fork ([https://github.com/takady/cheat/fork](https://github.com/takady/cheat/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create new Pull Request

## Author

[takady](https://github.com/takady)
