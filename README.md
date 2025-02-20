# Ignition

A simple CLI tool for fetching `.gitignore` files.

> [!WARNING]
> This project is currently in the early stages of development. Breaking changes might be introduced in future updates.

## Overview

Ignition will be a simple tool to fetch `.gitignore` files for various programming languages and save them a specified location. More details and features will be added as development progresses.

## Installation

### Using `go install`

```sh
go install github.com/notAyrun/ignition@latest
```

### Building from Source

```sh
# Clone the repository
git clone https://github.com/notAyrun/ignition.git

# This runs 'make build' automatically and generates the binary inside the build/ dir
make

# OR, you can simply run 'make install' to build and install the executable binary in your $GOPATH/bin
make install
```

## Usage

Make sure `ignition` is properly installed and available for use by running:

```sh
ignition --version
```

This should output the current version of ignition that is installed in your system.

To generate a `.gitignore` file for a programming language in your current working directory( let's say Python in this case ), run the following command:

```sh
ignition python
```

And here you go! A `.gitignore` file for Python has been successfully downloaded and placed in your current working directory!

You can also specify the output path( `-o` flag ) which will place the `.gitignore` in the specified path:

```sh
ignition python -o path/to/your/directory
```

To view a list of available programming languages for which the `.gitignore` file can be fetched, you can use the `list` command like this:

```sh
ignition list
```

## Roadmap

- [x] Implement fetching `.gitignore` files for different languages
- [x] Add support for saving `.gitignore` to custom locations
- [x] Implement displaying the list of programming languages that are supported
- [ ] Create a Neovim plugin to interact with the CLI tool
- [ ] Write detailed documentation

## Contributing

Contributions are welcome! If you'd like to contribute, follow these steps:

1. **Fork the repository**
2. **Create a new branch** ( `git checkout -b feature-name` )
3. **Make your changes** and commit ( `git commit -m "Describe your changes` )
4. **Push to your fork** ( `git push origin feature-name` )
5. **Open a pull request**

For major changes, please open an issue first to discuss your proposal.

Thank you for contributing!

## License

[MIT](./LICENSE) (c) [notAyrun](https://github.com/notAyrun)
