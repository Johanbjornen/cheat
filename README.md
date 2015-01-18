cheat
====

## Description
cheat sheet cli tool.

## Usage

Add a command and its description.
```bash
$ cheat add "ls" "show list of directories & files"
```

Show cheat list.

```bash
$ cheat list
# => 1) ls
# =>   show list of directories & files
# => 2) rm -r
# =>   remove directory
```

Delete a command.

```bash
$ cheat delete 1
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
