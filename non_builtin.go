// +build !builtin_static

package grocksdb

// #cgo LDFLAGS: -pthread -lstdc++ -ldl -lm -lzstd -llz4 -lz -lsnappy
import "C"
