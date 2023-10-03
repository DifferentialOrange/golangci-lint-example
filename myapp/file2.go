package main

import "github.com/tarantool/go-tarantool/v2"

func IsClosed(conn *tarantool.Connection) bool {
	return conn.ClosedNow()
}
