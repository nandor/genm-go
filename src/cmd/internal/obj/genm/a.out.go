// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package genm

import "cmd/internal/obj"

//go:generate go run ../stringer.go -i $GOFILE -o anames.go -p genm

const (
	ACALL = obj.ABaseGenM + obj.A_ARCHSPECIFIC + iota

	ASELECT

  ACONST
  AADDR
  ACOPY
  ALD
  AST_1
  AST_2
  AST_4
  AST_8

	AADD
	AAND
	ASLL
	ASLR

  ACMP_NE

  ACMP_EQ

  ACMP_LT
  ACMP_ULT
  ACMP_OLT

  ACMP_GT
  ACMP_UGT
  ACMP_OGT

  AJT

  ARET

  AARG
)

const (
	// globals
	REG_SP = obj.RBaseGenM + iota
  REG_SB

  REG_R0
  REG_R1
  REG_R2
  REG_R3
  REG_R4
  REG_R5
  REG_R6
  REG_R7
  REG_R8
  REG_R9
  REG_RA
  REG_RB
  REG_RC
  REG_RD
  REG_RE
  REG_RF

  REG_F0
  REG_F1
  REG_F2
  REG_F3
  REG_F4
  REG_F5
  REG_F6
  REG_F7
  REG_F8
  REG_F9
  REG_FA
  REG_FB
  REG_FC
  REG_FD
  REG_FE
  REG_FF

  REG_T0

  REGSP = REG_SP
)
