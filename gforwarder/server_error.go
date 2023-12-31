package main

import (
	"errors"
	"strconv"

	"github.com/ypcd/gforwarder/gstunnellib"
)

func checkError(err error) {
	gstunnellib.CheckErrorEx_exit(err, g_Logger)
}

func checkError_NoExit(err error) {
	gstunnellib.CheckErrorEx(err, g_Logger)
}

func checkError_info(err error) {
	gstunnellib.CheckErrorEx_info(err, g_Logger)
}

func checkError_panic(err error) {
	gstunnellib.CheckErrorEx_panic(err)
}

func checkError_GsCtx(err error, gctx gstunnellib.GsContext) {
	if err != nil {
		err2 := errors.New("[" + strconv.FormatUint(gctx.GetGsId(), 10) + "] " + err.Error())
		gstunnellib.CheckErrorEx_exit(err2, g_Logger)
	}
}

func checkError_NoExit_GsCtx(err error, gctx gstunnellib.GsContext) {
	if err != nil {
		err2 := errors.New("[" + strconv.FormatUint(gctx.GetGsId(), 10) + "] " + err.Error())
		gstunnellib.CheckErrorEx(err2, g_Logger)
	}
}

func checkError_info_GsCtx(err error, gctx gstunnellib.GsContext) {
	if err != nil {
		err2 := errors.New("[" + strconv.FormatUint(gctx.GetGsId(), 10) + "] " + err.Error())
		gstunnellib.CheckErrorEx_info(err2, g_Logger)
	}
}

func checkError_panic_GsCtx(err error, gctx gstunnellib.GsContext) {
	if err != nil {
		err2 := errors.New("[" + strconv.FormatUint(gctx.GetGsId(), 10) + "] " + err.Error())
		gstunnellib.CheckErrorEx_panic(err2)
	}
}
