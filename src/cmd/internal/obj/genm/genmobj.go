// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package genm

import (
	"cmd/internal/obj"
	"cmd/internal/sys"
)


var unaryDst = map[obj.As]bool{
}

var LinkGenM = obj.LinkArch{
	Arch:       sys.ArchGenM,
	Init:       instinit,
	Preprocess: preprocess,
	Assemble:   assemble,
	Progedit:   progedit,
	UnaryDst:   unaryDst,
}

func instinit(ctxt *obj.Link) {
	panic("instinit")
}

func preprocess(ctxt *obj.Link, s *obj.LSym, newprog obj.ProgAlloc) {
	panic("preprocess")
}

func assemble(ctxt *obj.Link, s *obj.LSym, newprog obj.ProgAlloc) {
	panic("assemble")
}

func progedit(ctxt *obj.Link, p *obj.Prog, newprog obj.ProgAlloc) {
	panic("progedit")
}
