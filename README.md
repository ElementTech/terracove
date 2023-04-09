
<h1 align="center">
  <br>
  <img src="./assets/logo/logo.png" alt="terracove" width="200">
</h1>

<h4 align="center">A recursive terraform repository tester powered by <a href="https://terratest.gruntwork.io/" target="_blank">Terratest</a>.</h4>

<p align="center">
<img src="https://github.com/jatalocks/terracove/actions/workflows/test.yml/badge.svg" alt="drawing"/>
<img src="https://github.com/jatalocks/terracove/actions/workflows/lint.yml/badge.svg" alt="drawing"/>
<img src="https://pkg.go.dev/badge/github.com/jatalocks/terracove.svg" alt="drawing"/>
<img src="https://codecov.io/gh/jatalocks/terracove/branch/main/graph/badge.svg" alt="drawing"/>
<img src="https://img.shields.io/github/v/release/jatalocks/terracove" alt="drawing"/>
<img src="https://img.shields.io/github/downloads/jatalocks/terracove/total.svg" alt="drawing"/>
</p>

<p align="center">
  <a href="#key-features">Key Features</a> •
  <a href="#install">Install</a> •
  <a href="#usage">Usage</a> •
  <a href="#quickstart">Quickstart</a> •
  <a href="#credits">Credits</a> •
  <a href="#support">Support</a> •
  <a href="#license">License</a>
</p>

<!-- ![screenshot](https://raw.githubusercontent.com/amitmerchant1990/electron-markdownify/master/app/img/markdownify.gif) -->

## Key Features

* Test in parallel multiple directory paths
* Generate a `junit.xml` with `--junit`
* Generate a `json` summary with `--json`
* Generate `%` coverage for each module and root directory
* Supports terraform and terragrunt in the same directory tree


## Install

The recommended way to install on MacOS is via brew:

```sh
brew tap jatalocks/jatalocks
brew install terracove
```

Or, you can install directly from release:
```sh
curl -sS https://raw.githubusercontent.com/jatalocks/terracove/main/install.sh | bash
```

If you'd like to use Docker, you can use the official image:
```sh
docker run --rm -v $(pwd):/app ghcr.io/jatalocks/terracove --json --junit /app
```
## Usage

```sh
Usage:
  terracove [paths]... [flags]

Flags:
  -e, --exclude strings         Exclude directories while parsing tree
  -h, --help                    help for terracove
  -j, --json                    Output JSON
  -x, --junit                   Output Junit XML
      --o-json string           Output JSON (default "terracove.json")
      --o-junit string          Output Junit XML (default "terracove.xml")
  -t, --validate-tf-by string   validate terraform by the existence of [filename] in a directory (default "main.tf")
  -g, --validate-tg-by string   validate terragrunt by the existence of [filename] in a directory (default "terragrunt.hcl")
  -v, --version                 version for terracove
```

## Quickstart
> Note that you must have terraform/terragrunt binaries installed on your machine

> The [examples](./examples) directory contains **4 modules**. 2 of them are [terraform](./examples/terraform) and 2 are [terragrunt](./examples/terragrunt). 
> 
> **Oh no!** It appears some of them have some problems. Let's see exactly what is going on.
> Clone this repository and give it a try.

```sh
git clone https://github.com/jatalocks/terracove.git
terracove --junit --json . # Same as terracove --junit --json examples
```

Open the **terracove.xml** or **terracove.json** and observe the results. You should see the following:

```json
[
 {
  "Timestamp": "1984-01-01T19:32:58+05:00",
  "Path": ".",
  "Results": [
   {
    "Path": "examples/terragrunt/no-resources",
    "ResourceCount": 0,
    "Coverage": 100,
    ...
   },
   {
    "Path": "examples/terragrunt/error",
    "Coverage": 0,
    ...
   },
   {
    "Path": "examples/terraform/tfstate-diff",
    "ResourceCount": 2,
    "ResourceCountDiff": 1,
    "Coverage": 50,
    ...
   },
   {
    "Path": "examples/terraform/success",
    "ResourceCount": 2,
    "ResourceCountExists": 2,
    "Coverage": 100,
    ...
   }
  ],
  "Coverage": 62.5
 }
]
```


## Credits

This software uses the following open source projects:

- [golang-cli-template](https://github.com/FalcoSuessgott/golang-cli-template)
- [terratest](https://terratest.gruntwork.io/)
## Support

<a href="https://www.buymeacoffee.com/jatalocks" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/purple_img.png" alt="Buy Me A Coffee" style="height: 41px !important;width: 174px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;" ></a>

## License

[MIT](LICENSE)
