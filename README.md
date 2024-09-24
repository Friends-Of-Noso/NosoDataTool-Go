# Noso Data Tool in Go
A tool that provides a `CLI` interface for `Noso` data files.

## Usage

```console
$ nosodatatoolgo help
This tool allows to query Noso data files

Usage:
  nosodatatoolgo [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  display     Displays various file contents in text format of JSON format
  export      Exports all data into SQL or a database
  help        Help about any command

Flags:
      --config string           config file (default is $HOME/.nosodatatoolgo)
  -h, --help                    help for nosodatatoolgo
  -d, --noso-directory string   root folder where to find Noso files

Use "nosodatatoolgo [command] --help" for more information about a command.
```

```console
$ nosodatatoolgo help display
Displays various file contents in text format of JSON format

Usage:
  nosodatatoolgo display [command]

Available Commands:
  block       Display the contents of blocks by block number
  gvt         Display the contents of a GVT file
  pso         Display the contents of a PSO file
  summary     Display the contents of a summary file
  wallet      Display the contents of a wallet file

Flags:
  -h, --help   help for display
  -j, --json   Outputs in JSON

Global Flags:
      --config string           config file (default is $HOME/.nosodatatoolgo)
  -d, --noso-directory string   root folder where to find Noso files

Use "nosodatatoolgo display [command] --help" for more information about a command.
```

```console
$ nosodatatoolgo help export
Exports all data into SQL or a database

Usage:
  nosodatatoolgo export [command]

Available Commands:
  sql         Exports all data to SQL

Flags:
  -h, --help   help for export

Global Flags:
      --config string           config file (default is $HOME/.nosodatatoolgo)
  -d, --noso-directory string   root folder where to find Noso files

Use "nosodatatoolgo export [command] --help" for more information about a command.
```