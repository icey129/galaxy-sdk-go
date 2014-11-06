// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package auth

import (
	"bytes"
	"fmt"
	"github.com/XiaoMi/galaxy-sdk-go/sds/common"
	"github.com/XiaoMi/galaxy-sdk-go/sds/errors"
	"github.com/XiaoMi/galaxy-sdk-go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var _ = errors.GoUnusedProtection__
var _ = common.GoUnusedProtection__
var SIGNATURE_SUPPORT map[UserType]bool

const HK_HOST = "Host"
const HK_TIMESTAMP = "X-Xiaomi-Timestamp"
const HK_CONTENT_MD5 = "X-Xiaomi-Content-MD5"
const HK_AUTHORIZATION = "Authorization"

var SUGGESTED_SIGNATURE_HEADERS []string

const MAX_CONTENT_SIZE = 524288

func init() {
	SIGNATURE_SUPPORT = map[UserType]bool{
		1:  false,
		2:  true,
		10: true,
		11: true,
		12: false,
		13: false,
	}

	SUGGESTED_SIGNATURE_HEADERS = []string{
		"Host", "X-Xiaomi-Timestamp", "X-Xiaomi-Content-MD5"}

}
