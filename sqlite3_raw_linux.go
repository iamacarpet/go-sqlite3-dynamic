// Copyright (C) 2016 Samuel Melrose <sam@infitialis.com>.
//
// Based on work by Yasuhiro Matsumoto <mattn.jp@gmail.com>
// https://github.com/mattn/go-sqlite3
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package sqlite3

import (
	"unsafe"

	"github.com/notti/nocgo"
)

func sqlite3_libversion() string {
	strSlice := raw_sqlite3_libversion()
	retVal := nocgo.MakeGoStringFromSlice(strSlice)
	if retVal == "" {
		return "Unknown"
	}
	return retVal
}

func sqlite3_libversion_number() int {
	retInt := raw_sqlite3_libversion_number()
	return int(retInt)
}

func sqlite3_sourceid() string {
	strSlice := raw_sqlite3_sourceid()
	retVal := nocgo.MakeGoStringFromSlice(strSlice)
	if retVal == "" {
		return "Unknown"
	}
	return retVal
}

func sqlite3_errstr(code int) (msg string) {
	strSlice := raw_sqlite3_errstr(int32(code))
	retVal := nocgo.MakeGoStringFromSlice(strSlice)
	if retVal == "" {
		return "Unknown Error"
	}
	return retVal
}

func sqlite3_errcode(db sqlite3) (code int) {
	retInt := raw_sqlite3_errcode(uintptr(db))
	return int(retInt)
}

func sqlite3_extended_errcode(db sqlite3) (code int) {
	retInt := raw_sqlite3_extended_errcode(uintptr(db))
	return int(retInt)
}

func sqlite3_errmsg(db sqlite3) (msg string) {
	strSlice := raw_sqlite3_errmsg(uintptr(db))
	retVal := nocgo.MakeGoStringFromSlice(strSlice)
	if retVal == "" {
		return "Unknown Error"
	}
	return retVal
}

func sqlite3_threadsafe() int {
	retInt := raw_sqlite3_threadsafe()
	return int(retInt)
}

func sqlite3_open_v2(filename string, ppDb *sqlite3, flags int, zVfs string) int {
	retInt := raw_sqlite3_open_v2(
		[]byte(filename),
		uintptr(unsafe.Pointer(ppDb)),
		int32(flags),
		[]byte(zVfs),
	)
	return int(retInt)
}

func sqlite3_busy_timeout(db sqlite3, busyTimeout int) int {
	retInt := raw_sqlite3_busy_timeout(uintptr(db), int32(busyTimeout))
	return int(retInt)
}

func sqlite3_close_v2(db sqlite3) int {
	retInt := raw_sqlite3_close_v2(uintptr(db))
	return int(retInt)
}

func sqlite3_prepare_v2(db sqlite3, zSql string) (retCode int, stmtHandle sqlite3_stmt, tail string) {
	var sql []byte = nocgo.MakeCString(zSql)

	var handle uintptr
	var thandle uintptr

	retInt := raw_sqlite3_prepare_v2(
		uintptr(db),
		sql,
		-1,
		uintptr(unsafe.Pointer(&handle)),
		uintptr(unsafe.Pointer(&thandle)),
	)

	return int(retInt), sqlite3_stmt(handle), BytePtrToString((*byte)(unsafe.Pointer(thandle)))
}

func sqlite3_get_autocommit(db sqlite3) int {
	retInt := raw_sqlite3_get_autocommit(uintptr(db))
	return int(retInt)
}

func sqlite3_finalize(stmt sqlite3_stmt) int {
	retInt := raw_sqlite3_finalize(uintptr(stmt))
	return int(retInt)
}

func sqlite3_bind_parameter_count(stmt sqlite3_stmt) int {
	retInt := raw_sqlite3_bind_parameter_count(uintptr(stmt))
	return int(retInt)
}

func sqlite3_bind_parameter_index(stmt sqlite3_stmt, name string) int {
	retInt := raw_sqlite3_bind_parameter_index(
		uintptr(stmt),
		nocgo.MakeCString(name),
	)

	return int(retInt)
}

func sqlite3_reset(stmt sqlite3_stmt) int {
	retInt := raw_sqlite3_reset(uintptr(stmt))
	return int(retInt)
}

func sqlite3_bind_null(stmt sqlite3_stmt, ord int) int {
	retInt := raw_sqlite3_bind_null(
		uintptr(stmt),
		int32(ord),
	)
	return int(retInt)
}

func sqlite3_bind_text(stmt sqlite3_stmt, ord int, data string) int {
	retInt := raw_sqlite3_bind_text(
		uintptr(stmt),
		int32(ord),
		[]byte(data),
		int32(len(data)),
		uintptr(SQLITE_TRANSIENT),
	)
	return int(retInt)
}

func sqlite3_bind_int64(stmt sqlite3_stmt, ord int, data int64) int {
	retInt := raw_sqlite3_bind_int64(
		uintptr(stmt),
		int32(ord),
		int64(data),
	)
	return int(retInt)
}

func sqlite3_bind_int(stmt sqlite3_stmt, ord int, data int32) int {
	retInt := raw_sqlite3_bind_int(
		uintptr(stmt),
		int32(ord),
		int32(data),
	)
	return int(retInt)
}

func sqlite3_bind_double(stmt sqlite3_stmt, ord int, data float64) int {
	retInt := raw_sqlite3_bind_double(
		uintptr(stmt),
		int32(ord),
		float64(data),
	)
	return int(retInt)
}

func sqlite3_bind_blob(stmt sqlite3_stmt, ord int, data []byte) int {
	var pData uintptr
	if len(data) == 0 {
		pData = 0
	} else {
		pData = uintptr(unsafe.Pointer(&data[0]))
	}
	retInt := raw_sqlite3_bind_blob(
		uintptr(stmt),
		int32(ord),
		pData,
		int32(len(data)),
		uintptr(SQLITE_TRANSIENT),
	)
	return int(retInt)
}

func sqlite3_column_count(stmt sqlite3_stmt) int {
	retInt := raw_sqlite3_column_count(uintptr(stmt))
	return int(retInt)
}

func sqlite3_column_name(stmt sqlite3_stmt, index int) string {
	strSlice := raw_sqlite3_column_name(
		uintptr(stmt),
		int32(index),
	)
	return nocgo.MakeGoStringFromSlice(strSlice)
}

func sqlite3_interrupt(db sqlite3) {
	raw_sqlite3_interrupt(uintptr(db))
}

func sqlite3_clear_bindings(stmt sqlite3_stmt) {
	raw_sqlite3_clear_bindings(uintptr(stmt))
}

func sqlite3_step(stmt sqlite3_stmt, rowid *int64, changes *int64) int {
	retInt := raw_sqlite3_step(
		uintptr(stmt),
	)

	dbHandle := raw_sqlite3_db_handle(uintptr(stmt))

	*rowid = raw_sqlite3_last_insert_rowid(dbHandle)
	*changes = int64(raw_sqlite3_changes(dbHandle))

	return int(retInt)
}

func sqlite3_column_decltype(stmt sqlite3_stmt, index int) string {
	strSlice := raw_sqlite3_column_decltype(
		uintptr(stmt),
		int32(index),
	)
	return nocgo.MakeGoStringFromSlice(strSlice)
}

func sqlite3_column_type(stmt sqlite3_stmt, index int) int {
	retInt := raw_sqlite3_column_type(
		uintptr(stmt),
		int32(index),
	)
	return int(retInt)
}

func sqlite3_column_int64(stmt sqlite3_stmt, index int) int64 {
	return raw_sqlite3_column_int64(uintptr(stmt), int32(index))
}

func sqlite3_column_double(stmt sqlite3_stmt, index int) float64 {
	return raw_sqlite3_column_double(uintptr(stmt), int32(index))
}

func sqlite3_column_bytes(stmt sqlite3_stmt, index int) int {
	retInt := raw_sqlite3_column_bytes(
		uintptr(stmt),
		int32(index),
	)
	return int(retInt)
}

func sqlite3_column_blob(stmt sqlite3_stmt, index int) []byte {
	bytesPtr := raw_sqlite3_column_blob(
		uintptr(stmt),
		int32(index),
	)

	n := sqlite3_column_bytes(stmt, index)

	slice := make([]byte, n)
	copy(slice[:], (*[1 << 30]byte)(unsafe.Pointer(bytesPtr))[0:n])
	return slice
}

func sqlite3_column_text(stmt sqlite3_stmt, index int) string {
	bytesPtr := raw_sqlite3_column_text(
		uintptr(stmt),
		int32(index),
	)

	n := sqlite3_column_bytes(stmt, index)

	slice := make([]byte, n)
	copy(slice[:], (*[1 << 30]byte)(unsafe.Pointer(bytesPtr))[0:n])
	return string(slice)
}
