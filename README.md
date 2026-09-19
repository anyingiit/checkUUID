<!-- Source: Best-README-Template BLANK_README (Unlicense) — https://github.com/othneildrew/Best-README-Template -->
<a id="readme-top"></a>

# checkUUID

A Go command-line program that generates 200,000,000 UUIDs in a single run and reports whether any of them collide.

**English** · [简体中文](README.zh-CN.md)

[![CI](https://github.com/anyingiit/checkUUID/actions/workflows/ci.yml/badge.svg)](https://github.com/anyingiit/checkUUID/actions/workflows/ci.yml)
[![License](https://img.shields.io/github/license/anyingiit/checkUUID)](LICENSE)

[Report a bug](https://github.com/anyingiit/checkUUID/issues/new?template=bug_report.yml) · [Request a feature](https://github.com/anyingiit/checkUUID/issues/new?template=feature_request.yml)

<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#about-the-project">About The Project</a></li>
    <li><a href="#getting-started">Getting Started</a></li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

## About The Project

checkUUID is a small Go program that stress-tests the `google/uuid` package.
`main.go` generates 200,000,000 version-4 UUIDs in a loop, keeping every value
it has already seen so it can spot a repeat, while a worker goroutine defined
in `woker.go` prints progress in ten-percent steps as the run goes. When the
loop finishes, it reports either that all 200,000,000 UUIDs were unique or, for
every UUID it saw more than once, the value and how many times it recurred.

See the [open issues](https://github.com/anyingiit/checkUUID/issues) for planned features and known issues.

## Getting Started

### Prerequisites

- Go 1.19 or newer, the version declared in `go.mod`
- Git, to clone the repository

### Installation

```sh
git clone https://github.com/anyingiit/checkUUID.git
cd checkUUID
go build .
```

## Usage

The program takes no flags and no input; running it immediately starts the
fixed, 200,000,000-iteration collision check:

```sh
go run .
```

It prints its progress every ten percent, and finishes by printing
`test uuid of 200000000 ok!!!!!` if every UUID it generated was unique, or a
list of the UUIDs (and their repeat counts) it saw more than once.

## Contributing

Contributions are welcome. Read [CONTRIBUTING.md](CONTRIBUTING.md) for how to open an issue or a pull request, and [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) for the standards expected of everyone taking part.

Please do not report security issues in public issues or pull requests. [SECURITY.md](SECURITY.md) explains how to report them privately.

## License

Distributed under the MIT License. See [LICENSE](LICENSE) for details.

## Contact

Project link: [https://github.com/anyingiit/checkUUID](https://github.com/anyingiit/checkUUID)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
