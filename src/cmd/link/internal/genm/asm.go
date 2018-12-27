// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package genm

import (
	"cmd/link/internal/ld"
	"cmd/link/internal/sym"
)

func gentext(ctxt *ld.Link) {
}

func assignAddress(ctxt *ld.Link, sect *sym.Section, n int, s *sym.Symbol, va uint64, isTramp bool) (*sym.Section, int, uint64) {
	s.Sect = sect
	s.Value = int64(va)
	va += 1
	return sect, n, va
}

func asmb(ctxt *ld.Link) {
}
