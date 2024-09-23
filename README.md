# Noso Data Tool in Go
A tool that provides a `CLI` interface for `Noso` data files.

## Usage

```console
$ nosodatatoolgo help
This tool allows to query Noso data files

Usage:
  nosodatatoolgo [command]

Available Commands:
  block       This command allows to query blocks by block number
  completion  Generate the autocompletion script for the specified shell
  gvt         This command allows to query the contents of a GVT file
  help        Help about any command
  summary     This command allows to query the contents of a summary file
  wallet      This command allows to query the contents of a wallet file

Flags:
      --config string           config file (default is $HOME/.nosodatatoolgo)
  -h, --help                    help for nosodatatoolgo
  -d, --noso-directory string   root folder where to find Noso files

Use "nosodatatoolgo [command] --help" for more information about a command.
```