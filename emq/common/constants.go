// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package common

import (
	"bytes"
	"fmt"
	"github.com/XiaoMi/galaxy-sdk-go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var ERROR_BACKOFF map[ErrorCode]int64
var ERROR_RETRY_TYPE map[ErrorCode]RetryType

const MAX_RETRY = 3

func init() {
	ERROR_BACKOFF = map[ErrorCode]int64{
		15: 1000,
		18: 1000,
		16: 1000,
		17: 1000,
		7:  1000,
		14: 1000,
		19: 1000,
	}

	ERROR_RETRY_TYPE = map[ErrorCode]RetryType{
		15: 0,
		18: 0,
		16: 0,
		17: 0,
		7:  0,
		14: 1,
		19: 2,
	}

}