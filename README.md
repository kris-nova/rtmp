# UNOFFICIAL FORK

Hi - I did not write the majority of this code.

This is mostly a copy from [here](https://github.com/todostreaming/rtmp).

I have ported the code over and retained the original creative commons license.

I will be making changes to make this code (and it's imports) compile.  

This comes with no service guarantee and you should never use this code, especially in production.

--- 

# rtmp (ORIGINAL)

[![Build
Status](https://travis-ci.org/kris-nova/rtmp.svg?branch=master)](https://travis-ci.org/kris-nova/rtmp)
[![GoDoc](https://godoc.org/github.com/kris-nova/rtmp?status.svg)](https://godoc.org/github.com/kris-nova/rtmp)

`rtmp` is a Golang implementation of the RTMP specification, found
[here](http://www.adobe.com/devnet/rtmp.html). It is entirely specification
compliant, and implements all modern parts of the RTMP spec, most used in the
wild.

It is currently used in Beam's internal RTMP ingest server.

# Getting Started

## Installation

`rtmp` is distributed as a Golang library. It is easily accessible from within
other Go packages by simply importing it. To make `rtmp` available in your
environment, simply `go get` it:

```
go get github.com/kris-nova/rtmp
```

Alternatively, you can lock `rtmp` as a dependency, or fetch it using
[gopkg.in](http://labix.org/gopkg.in), applying the standard conventions.

## Core Concepts

At its most basic form, the `rtmp` package provides a simple RTMP server
implementation, that returns a `<-chan *client.Client`. The `client.Client` type
has several methods defined on it which reference other packages in the library
to preform more composable operations, such as chunk reading, data streams, and
so forth.

### Server

An RTMP has only one job, which is to return RTMP clients, such that the caller
may do with them what they want. An example implementation of processing clients
from a server follows:

```go
package main

import (
        "github.com/kris-nova/rtmp/server"
)

func main() {
        s := server.New(":1935")

        go s.Accept()
        defer s.Close()

        for {
                select {
                case c := <-s.Clients():
                        // ...
                case err := <-s.Errs():
                        // ...
                }
        }
}
```

### Client

The client, on the other hand, has many jobs. It serves as the top node in a
tree, capable of calling into all other packages dealing with particular parts
of the RTMP specification. For example, the client knows how to:

  * Handshake itself by using the `github.com/kris-nova/rtmp/handshake` package
  * Receive RTMP control sequences using the `github.com/kris-nova/rtmp/control`
    package
  * More to come...

For more information on all of the things that the client can do, see the
[relevant
documentation](https://godoc.org/github.com/kris-nova/rtmp/client#Client).

## License

Creative Commons Attribution-NonCommercial 4.0 International.
