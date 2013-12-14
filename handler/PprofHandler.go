package handler

import (
	"mzr/lib"
	"net/http/pprof"
)

type PprofHandler struct {
	lib.BaseHandler
}

func (self *PprofHandler) Get() {

	switch self.Ctx.Input.Params(":pp") {
	default:
		pprof.Index(self.Ctx.ResponseWriter, self.Ctx.Request)
	case "":
		pprof.Index(self.Ctx.ResponseWriter, self.Ctx.Request)
	case "cmdline":
		pprof.Cmdline(self.Ctx.ResponseWriter, self.Ctx.Request)
	case "profile":
		pprof.Profile(self.Ctx.ResponseWriter, self.Ctx.Request)
	case "symbol":
		pprof.Symbol(self.Ctx.ResponseWriter, self.Ctx.Request)
	}
	self.Ctx.ResponseWriter.WriteHeader(200)
}
