# surveyor

A Command-Line Interface (CLI) utility for interactive terminal prompts.

Inspired by [enquirer](https://github.com/termapps/enquirer).

**Disclaimer**: This is still under development and not ready for production.
See [critical bug](https://github.com/pbnj/surveyor/issues/1) that defeats the
whole premise of this project.

## Table of Contents

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [Install](#install)
  - [Homebrew](#homebrew)
  - [Go](#go)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Install

`surveyor` is available for Linux, macOS, Windows, and more.

### Homebrew

_Coming soon..._

### Go

`GO111MODULE="on" go get -u -v github.com/pbnj/surveyor`

## Usage

```sh

NAME:
   surveyor - A Command-Line Interface (CLI) utility for interactive terminal prompts

USAGE:
   main [global options] command [command options] [arguments...]

COMMANDS:
   input        Prompt user for text input
   multiline    Prompt user for multiline text input
   password     Prompt user for password input
   confirm      Prompt user for confirmation
   select       Prompt user for single-choice selection
   multiselect  Prompt user for multi-choice selection
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)

```

## Example

```sh

```

### Input

```sh
GH_USER=$(surveyor input --message "What is your GitHub username:")

curl https://api.github.com/users/${GH_USER}
```

### Multiline

```sh

```

### Password

```sh

```

### Confirm

```sh

```

### select

```sh

```

### multiselect

```sh

```

## Contributing

Any and all contributions are welcome!

For bugs, suggestions, or enhancements, please submit an issue.

## Special Thanks

Thank you to authors and maintainers of:

- [AlecAivazis/survey](https://github.com/AlecAivazis/survey)
- [urfave/cli](https://github.com/urfave/cli)

## License

MIT
