# Fatehound

Fatehound is a command-line tool designed to clean up specific XML configuration files. It primarily targets settings files used by Streambox SpectraUI.

## Purpose

The main purpose of Fatehound is to remove the `video_3d` attribute from specified XML files. This can be useful for maintenance or cleanup operations on Streambox SpectraUI configurations.

## Features

- Processes multiple XML files in a single run
- Configurable through command-line flags and a configuration file
- Supports various logging levels for detailed output control
- Default paths set to common Streambox SpectraUI configuration locations

## Usage

```
fatehound test [flags]
```

### Flags

- `--log-level`: Set the logging level (trace, debug, info, warn, error, fatal, panic)
- `--path`: Specify custom paths to XML files (can be used multiple times)

## Default Behavior

By default, Fatehound will attempt to process the following files:

1. `C:\ProgramData\Streambox\SpectraUI\settings.xml.bak`
2. `C:\ProgramData\Streambox\SpectraUI\settings.xml`

## Building and Running

To build and run the project:

1. Ensure you have Go installed on your system.
2. Clone the repository.
3. Run `go build` in the project directory.
4. Execute the resulting binary: `./fatehound test`
