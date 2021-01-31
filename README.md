# go-config

Simple configuration utility to use settings from local configuration file.

Configuration will be loaded from `config.json` and `config-{profile}.json` for each profile afterwards.
Each file will replace already loaded settings.


## Usage
```golang
package main

import (
	"flag"
	"github.com/dreske/go-config"
)

type AppConfig struct {
	Host *string
	Port int
}

var appConfig AppConfig

func main() {
	flag.Parse()
	if err := go_config.Load(&appConfig); err != nil {
		panic(err)
	}
}
```

## Profiles
Application may be started with `-profile` command line argument. Multiple active profiles may be specified using a comma separated list.
A config file will be loaded for each profile in the specified order (if one exists).

## Example
config.json
```json
{
  "host": "example.com",
  "port": 1234
}
```

config-local.json
```json
{
  "host": "localhost"
}
```

Calling the program with `-profiles=local` will use the following effective configuration
```json
{
  "host": "localhost",
  "port": 1234
}
```
