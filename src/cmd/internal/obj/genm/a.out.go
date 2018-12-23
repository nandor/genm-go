// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package genm

import "cmd/internal/obj"

//go:generate go run ../stringer.go -i $GOFILE -o anames.go -p wasm

const (
	SET = obj.ABaseGenM + obj.A_ARCHSPECIFIC + iota

	ADD
	AND
	CMP
	DIV
	REM
	MUL
	OR
	ROTL
	SLL
	SRA
	SRL
	SUB
	XOR
	POW
)

const (
	// globals
	REG_SP = obj.RBaseGenM + iota

	REGSP = REG_SP
)
