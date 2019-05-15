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
	"os"
	"fmt"

	"database/sql"

	"github.com/iamacarpet/go-sqlite3-dynamic"
)

func main() {
	resetTime := time.Now()

	fmt.Println(sqlite3.Version())

	db, err := sql.Open(`sqlite3`, "file:"+resetTime.Format("2006-01-02")+"?mode=memory&cache=shared")
	if err != nil {
		panic(err)
	}

	r, err := db.Exec(`CREATE TABLE test (
		id integer PRIMARY KEY NOT NULL,
		name varchar(30)
	)`)
	if err != nil {
		panic(err)
	}

	_ = r

	r, err = db.Exec(`INSERT INTO test(name) VALUES ('first') `)
	if err != nil {
		panic(err)
	}
	_, err = r.LastInsertId()
	if err != nil {
		panic(err)
	}
	_, err = r.RowsAffected()
	if err != nil {
		panic(err)
	}

	r, err = db.Exec(`INSERT INTO test(name) VALUES ('second') `)
	if err != nil {
		panic(err)
	}
	_, err = r.LastInsertId()
	if err != nil {
		panic(err)
	}
	_, err = r.RowsAffected()
	if err != nil {
		panic(err)
	}

	db.Close()
}
```


### On Disk SQLite

```go
package main

import (
	"os"
	"fmt"

	"database/sql"

	"github.com/iamacarpet/go-sqlite3-dynamic"
)

func main() {
	path := "test.db"

	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	f.Close()

	fmt.Println(sqlite3.Version())

	db, err := sql.Open(`sqlite3`, path)
	if err != nil {
		panic(err)
	}

	r, err := db.Exec(`CREATE TABLE test (
		id integer PRIMARY KEY NOT NULL,
		name varchar(30)
	)`)
	if err != nil {
		panic(err)
	}

	_ = r

	r, err = db.Exec(`INSERT INTO test(name) VALUES ('first') `)
	if err != nil {
		panic(err)
	}
	_, err = r.LastInsertId()
	if err != nil {
		panic(err)
	}
	_, err = r.RowsAffected()
	if err != nil {
		panic(err)
	}

	r, err = db.Exec(`INSERT INTO test(name) VALUES ('second') `)
	if err != nil {
		panic(err)
	}
	_, err = r.LastInsertId()
	if err != nil {
		panic(err)
	}
	_, err = r.RowsAffected()
	if err != nil {
		panic(err)
	}

	db.Close()
	os.Remove(`./test.db`)
}
```