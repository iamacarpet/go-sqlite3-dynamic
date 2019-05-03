# Go-SQLite3-Dynamic
## GoLang `database/sql` driver without CGo

This package provides an alternative to [go-sqlite3](https://github.com/mattn/go-sqlite3) without CGo.

It relies on calls to external libraries, using either DLL syscalls on Windows, or [nocgo](https://github.com/notti/nocgo) on Linux.

This should allow really easy cross compile support when building from another OS (e.g. Linux).

Basic functionality is implemented, but it doesn't support things like user defined functions that call back to Go code.

Otherwise, usage should be the same as for go-sqlite3.

## Windows

### DLL Search Path

By default, we'll look for `sqlite3.dll` in either the same directory as your executable, or in a `support` folder.

You'll need [sqlite3.dll](https://sqlite.org/download.html) to match your compiled architecture (32/64bit).

## Linux

Linux support is expermimental, as [nocgo](https://github.com/notti/nocgo)'s implementation of dlopen is still considered experimental.

It'll use the SQLite3 library installed on your system by the package manager.

### Installing on Ubuntu

```
apt-get -y install libsqlite3-0
```

## Examples

### In-Memory SQLite

```go
package main

import (

)

func main() {

}
```


### On Disk SQLite

```go
package main

import (

)

func main() {

}
```