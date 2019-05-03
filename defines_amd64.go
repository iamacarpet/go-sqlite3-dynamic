// Copyright (C) 2017 Hank Shen <swh@admpub.com>.
// - Thanks @admpub from @iamacarpet!!
//
// Based on work by Yasuhiro Matsumoto <mattn.jp@gmail.com>
// https://github.com/mattn/go-sqlite3
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package sqlite3

import "math"

const (
	SQLITE_TRANSIENT = math.MaxUint64 // Can't do -1 for overflow like in C, so use largest unsigned 64bit int.
)
