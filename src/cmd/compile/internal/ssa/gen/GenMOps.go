// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import "strings"

var regNamesGenM = []string{
	"R0", "R1", "R2", "R3", "R4", "R5", "R6", "R7", "R8", "R9", "RA", "RB", "RC", "RD", "RE", "RF",
	"F0", "F1", "F2", "F3", "F4", "F5", "F6", "F7", "F8", "F9", "FA", "FB", "FC", "FD", "FE", "FF",

	"SP",
	"SB",
}

func init() {
	// Make map from reg names to reg integers.
	if len(regNamesGenM) > 64 {
		panic("too many registers")
	}
	num := map[string]int{}
	for i, name := range regNamesGenM {
		num[name] = i
	}
	buildReg := func(s string) regMask {
		m := regMask(0)
		for _, r := range strings.Split(s, " ") {
			if n, ok := num[r]; ok {
				m |= regMask(1) << uint(n)
				continue
			}
			panic("register " + r + " not found")
		}
		return m
	}

	var (
		gp    		 = buildReg("R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 RA RB RC RD RE RF")
		fp     		 = buildReg("F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 FA FB FC FD FE FF")
	)

	// Common regInfo
	var (
		gp01    = regInfo{inputs: nil, 							  		outputs: []regMask{gp}}
		gp11    = regInfo{inputs: []regMask{gp},      		outputs: []regMask{gp}}
		gp21    = regInfo{inputs: []regMask{gp, gp},  		outputs: []regMask{gp}}
		gp31    = regInfo{inputs: []regMask{gp, gp, gp},  outputs: []regMask{gp}}
		gpload  = regInfo{inputs: []regMask{gp, 0},   		outputs: []regMask{gp}}
		fpload  = regInfo{inputs: []regMask{gp, 0},   		outputs: []regMask{fp}}
		gpstore = regInfo{inputs: []regMask{gp, gp, 0}}
		fpstore = regInfo{inputs: []regMask{gp, fp, 0}}
	)

	var GenMOps = []opData{
		// GenM pseudo instructions.
		{name: "LoweredStaticCall",  argLength: 1, aux: "SymOff", call: true, symEffect: "None"},
		{name: "LoweredInterCall",   argLength: 2, reg: regInfo{inputs: []regMask{gp} }, aux: "Int64", call: true},
		{name: "LoweredNilCheck", 	 argLength: 2, reg: regInfo{inputs: []regMask{gp}}, nilCheck: true, faultOnNilArg0: true},
		{name: "LoweredAddr", 			 argLength: 1, reg: gp11, aux: "SymOff", rematerializeable: true, symEffect: "Addr"},
		{name: "LoweredWB", 				 argLength: 3, reg: regInfo{inputs: []regMask{gp, gp}}, aux: "Sym", symEffect: "None"},

		// GenM instructions.
		{name: "SELECT_I64",  asm: "SELECT",  argLength: 3, reg: gp31},

		{name: "CONST_I1",		asm: "CONST",   reg: gp01,    typ: "Int1",   aux: "Bool",  rematerializeable: true},
		{name: "CONST_I8",		asm: "CONST",   reg: gp01,    typ: "Int8",   aux: "Int8",  rematerializeable: true},
		{name: "CONST_I16",		asm: "CONST",   reg: gp01,    typ: "Int16",  aux: "Int16", rematerializeable: true},
		{name: "CONST_I32",		asm: "CONST",   reg: gp01,    typ: "Int32",  aux: "Int32", rematerializeable: true},
		{name: "CONST_I64",		asm: "CONST",   reg: gp01,    typ: "Int64",  aux: "Int64", rematerializeable: true},

		{name: "ADD_I64",			asm: "ADD",     reg: gp21,    typ: "Int64",    argLength: 2},
		{name: "AND_I64",			asm: "AND",			reg: gp21, 		typ: "Int64",    argLength: 2},
		{name: "OR_I64",			asm: "AND",			reg: gp21, 		typ: "Int64",    argLength: 2},

		{name: "SLR_I64",			asm: "SLL",     reg: gp21,    typ: "Int64",    argLength: 2},
		{name: "SLR_I32",			asm: "SLL",     reg: gp21,    typ: "Int32",    argLength: 2},
		{name: "SLR_I16",			asm: "SLL",     reg: gp21,    typ: "Int16",    argLength: 2},
		{name: "SLR_I8",			asm: "SLL",     reg: gp21,    typ: "Int8",     argLength: 2},

		{name: "SLL_I64",			asm: "SLL",     reg: gp21,    typ: "Int64",    argLength: 2},
		{name: "SLL_I32",			asm: "SLL",     reg: gp21,    typ: "Int32",    argLength: 2},
		{name: "SLL_I16",			asm: "SLL",     reg: gp21,    typ: "Int16",    argLength: 2},
		{name: "SLL_I8",			asm: "SLL",     reg: gp21,    typ: "Int8",     argLength: 2},

		{name: "CMP_NE_I32",	asm: "CMP_NE",  reg: gp21,    typ: "Bool",     argLength: 2},
		{name: "CMP_NE_I64",	asm: "CMP_NE",  reg: gp21,    typ: "Bool",     argLength: 2},

		{name: "CMP_EQ_I8", 	asm: "CMP_EQ",  reg: gp21,    typ: "Bool",     argLength: 2},

		{name: "CMP_LT_I64",	asm: "CMP_LT",  reg: gp21,    typ: "Bool",     argLength: 2},
		{name: "CMP_LT_I32",	asm: "CMP_LT",  reg: gp21,    typ: "Bool",     argLength: 2},
		{name: "CMP_LT_I16",	asm: "CMP_LT",  reg: gp21,    typ: "Bool",     argLength: 2},
		{name: "CMP_LT_I8",		asm: "CMP_LT",  reg: gp21,    typ: "Bool",     argLength: 2},
		{name: "CMP_ULT_I64",	asm: "CMP_ULT", reg: gp21,    typ: "Bool",     argLength: 2},
		{name: "CMP_ULT_I32",	asm: "CMP_ULT", reg: gp21,    typ: "Bool",     argLength: 2},
		{name: "CMP_ULT_I16",	asm: "CMP_ULT", reg: gp21,    typ: "Bool",     argLength: 2},
		{name: "CMP_ULT_I8",	asm: "CMP_ULT", reg: gp21,    typ: "Bool",     argLength: 2},
		{name: "CMP_OLT_F64",	asm: "CMP_OLT", reg: gp21,    typ: "Bool",     argLength: 2},
		{name: "CMP_OLT_F32",	asm: "CMP_OLT", reg: gp21,    typ: "Bool",     argLength: 2},

		{name: "CMP_GT_I64",	asm: "CMP_GT",  reg: gp21,    typ: "Bool",     argLength: 2},
		{name: "CMP_GT_I32",	asm: "CMP_GT",  reg: gp21,    typ: "Bool",     argLength: 2},
		{name: "CMP_GT_I16",	asm: "CMP_GT",  reg: gp21,    typ: "Bool",     argLength: 2},
		{name: "CMP_GT_I8",		asm: "CMP_GT",  reg: gp21,    typ: "Bool",     argLength: 2},
		{name: "CMP_UGT_I64",	asm: "CMP_UGT", reg: gp21,    typ: "Bool",     argLength: 2},
		{name: "CMP_UGT_I32",	asm: "CMP_UGT", reg: gp21,    typ: "Bool",     argLength: 2},
		{name: "CMP_UGT_I16",	asm: "CMP_UGT", reg: gp21,    typ: "Bool",     argLength: 2},
		{name: "CMP_UGT_I8",	asm: "CMP_UGT", reg: gp21,    typ: "Bool",     argLength: 2},
		{name: "CMP_OGT_F64",	asm: "CMP_OGT", reg: gp21,    typ: "Bool",     argLength: 2},
		{name: "CMP_OGT_F32",	asm: "CMP_OGT", reg: gp21,    typ: "Bool",     argLength: 2},

		{name: "LD_8_U64",		asm: "LD", 		  reg: gpload,  typ: "UInt64",   argLength: 2, aux: "Int64"},
		{name: "LD_8_I64",		asm: "LD", 		  reg: gpload,  typ: "Int64",    argLength: 2, aux: "Int64"},
		{name: "LD_4_U32",		asm: "LD", 		  reg: gpload,  typ: "UInt32",   argLength: 2, aux: "Int64"},
		{name: "LD_4_I32",		asm: "LD", 		  reg: gpload,  typ: "Int32",    argLength: 2, aux: "Int64"},
		{name: "LD_2_U16",		asm: "LD", 		  reg: gpload,  typ: "UInt16",   argLength: 2, aux: "Int64"},
		{name: "LD_2_I16",		asm: "LD", 		  reg: gpload,  typ: "Int16",    argLength: 2, aux: "Int64"},
		{name: "LD_1_U8",			asm: "LD", 		  reg: gpload,  typ: "UInt8",    argLength: 2, aux: "Int64"},
		{name: "LD_1_I8",			asm: "LD", 		  reg: gpload,  typ: "Int8",     argLength: 2, aux: "Int64"},
		{name: "LD_4_F32",		asm: "LD", 		  reg: fpload,  typ: "Float32",  argLength: 2, aux: "Int64"},
		{name: "LD_8_F64",		asm: "LD", 		  reg: fpload,  typ: "Float64",  argLength: 2, aux: "Int64"},

		{name: "ST_8_I",			asm: "ST_8", 	  reg: gpstore, typ: "Mem",      argLength: 3, aux: "Int64"},
		{name: "ST_4_I",			asm: "ST_4", 	  reg: gpstore, typ: "Mem",      argLength: 3, aux: "Int64"},
		{name: "ST_2_I",			asm: "ST_2", 	  reg: gpstore, typ: "Mem",      argLength: 3, aux: "Int64"},
		{name: "ST_1_I",			asm: "ST_1", 	  reg: gpstore, typ: "Mem",      argLength: 3, aux: "Int64"},
		{name: "ST_8_F",			asm: "ST_8", 	  reg: fpstore, typ: "Mem",      argLength: 3, aux: "Int64"},
		{name: "ST_4_F",			asm: "ST_4", 	  reg: fpstore, typ: "Mem",      argLength: 3, aux: "Int64"},
	}

	archs = append(archs, arch{
		name:            "GenM",
		pkg:             "cmd/internal/obj/genm",
		genfile:         "../../genm/ssa.go",
		ops:             GenMOps,
		blocks:          nil,
		regnames:        regNamesGenM,
		gpregmask:       gp,
		fpregmask:       fp,
		framepointerreg: -1, // not used
		linkreg:         -1, // not used
	})
}
