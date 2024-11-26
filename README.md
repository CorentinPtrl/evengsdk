# evengsdk

`evengsdk` is a Go client library for interacting with the EVE-NG API. It provides methods to manage labs, nodes, networks, and folders within an EVE-NG environment.

## Installation

Install the library using `go get`:

```sh
go get github.com/CorentinPtrl/evengsdk
```

## Usage

### Basic Authentication Client

Create a new client with basic authentication:

```go
package main

import (
    "github.com/CorentinPtrl/evengsdk"
    "log"
)

func main() {
    client, err := evengsdk.NewBasicAuthClient("username", "password", "0", "http://your-eve-ng-host")
    if err != nil {
        log.Fatal(err)
    }

    folder, err := client.Folder.GetFolder("/")
    if err != nil {
        log.Fatal(err)
    }
    log.Println(folder)
}
```

## Testing

Run the tests using:

```sh
go test ./test
```

## Credits
Sander van Harmelen for [go-gitlab](https://github.com/xanzy/go-gitlab) which was used as a reference for the structure of this library.
