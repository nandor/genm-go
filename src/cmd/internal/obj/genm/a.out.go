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
	ASUB
	AMUL
	ADIV
	AMOD
	AAND
	AOR
	AXOR
	ANEG
	ASAR
	ASLR
	ASLL

  ACMP_NE
  ACMP_GE
  ACMP_EQ
  ACMP_LT
  ACMP_GT
  ACMP_LE

  ACMP_OLT
  ACMP_OGE
  ACMP_OGT
  ACMP_OLE
  ACMP_ONE

  ACMP_ULT
  ACMP_UGE
  ACMP_UGT
  ACMP_ULE

  AZEXT
  ASEXT
  AFEXT
  ATRUNC
  ANOT

  AJT
  AJF
  AJMP

  ARET

  AARG
)

const (
	// globals
	REG_SP = obj.RBaseGenM + iota
  REG_SB

  REG_R00
  REG_R01
  REG_R02
  REG_R03
  REG_R04
  REG_R05
  REG_R06
  REG_R07
  REG_R08
  REG_R09
  REG_R0A
  REG_R0B
  REG_R0C
  REG_R0D
  REG_R0E
  REG_R0F

  REG_R10
  REG_R11
  REG_R12
  REG_R13
  REG_R14
  REG_R15
  REG_R16
  REG_R17
  REG_R18
  REG_R19
  REG_R1A
  REG_R1B
  REG_R1C
  REG_R1D
  REG_R1E
  REG_R1F

  REG_T0

  REGSP = REG_SP
)
