/*
*
*Open source agreement:
*   The project is based on the GPLv3 protocol.
*
 */
package gstunnellib

import (
	"fmt"
	"log"
	"net"

	"github.com/panjf2000/ants/v2"
	"github.com/ypcd/gforwarder/gstunnellib/gsbase"
)

const G_Version string = gsbase.G_Version

var p func(...interface{}) (int, error)

//var begin_Dinfo int = 0

var g_commonIV = []byte{171, 158, 1, 73, 31, 98, 64, 85, 209, 217, 131, 150, 104, 219, 33, 220}

var g_logger *log.Logger

const G_Info_protobuf bool = true

const G_Deep_debug = gsbase.G_Deep_debug
const G_RunTime_Debug = gsbase.G_RunTime_Debug

var G_RunTimeDebugInfo1 RunTimeDebugInfo

//var g_key_Default string = gsbase.G_AesKeyDefault

var G_ants_pool *ants.Pool

func Nullprint(v ...interface{}) (int, error)                       { return 1, nil }
func Nullprintf(format string, a ...interface{}) (n int, err error) { return 1, nil }

func init() {
	p = Nullprint

	g_logger = NewLoggerFileAndStdOut("gstunnellib.log")
	//debug_tag = true
	//p = fmt.Println
	if G_RunTime_Debug {
		G_RunTimeDebugInfo1 = NewRunTimeDebugInfo()
	}
	//var err error
	//G_ants_pool, err = ants.NewPool(0)
	//CheckErrorEx_panic(err)
}

func GetNetConnAddrString(str string, conn net.Conn) string {
	return fmt.Sprintf("%s: [localIp:%s  remoteIp:%s]\n",
		str, conn.LocalAddr().String(), conn.RemoteAddr().String())
}
