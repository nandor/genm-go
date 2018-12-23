// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package genm

import (
	"cmd/compile/internal/gc"
	"cmd/compile/internal/ssa"
)

// markMoves marks any MOVXconst ops that need to avoid clobbering flags.
func ssaMarkMoves(s *gc.SSAGenState, b *ssa.Block) {
	panic("ssaMarkMoves")
}

func ssaGenValue(s *gc.SSAGenState, v *ssa.Value) {
	panic("ssaGenValue")
}

func ssaGenBlock(s *gc.SSAGenState, b, next *ssa.Block) {
	panic("ssaGenBlock")
}
