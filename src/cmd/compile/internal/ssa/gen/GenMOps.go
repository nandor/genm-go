// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import "strings"

var regNamesGenM = []string{
	"R00", "R01", "R02", "R03", "R04", "R05", "R06", "R07", "R08", "R09", "R0A", "R0B", "R0C", "R0D", "R0E", "R0F",
	"R10", "R11", "R12", "R13", "R14", "R15", "R16", "R17", "R18", "R19", "R1A", "R1B", "R1C", "R1D", "R1E", "R1F",

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
		r    		 	= buildReg("R00 R01 R02 R03 R04 R05 R06 R07 R08 R09 R0A R0B R0C R0D R0E R0F R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R1A R1B R1C R1D R1E R1F")
	)

	// Common regInfo
	var (
		r01    	= regInfo{inputs: nil, 							  	outputs: []regMask{r}}
		r11    	= regInfo{inputs: []regMask{r},      		outputs: []regMask{r}}
		r21    	= regInfo{inputs: []regMask{r, r},  		outputs: []regMask{r}}
		r31    	= regInfo{inputs: []regMask{r, r, r}, 	outputs: []regMask{r}}
		rload  	= regInfo{inputs: []regMask{r, 0},   		outputs: []regMask{r}}
		rstore	= regInfo{inputs: []regMask{r, r, 0}}
	)

	var GenMOps = []opData{
		// GenM pseudo instructions.
		{name: "LoweredStaticCall",  argLength: 1, reg: regInfo{clobbers: r}, aux: "SymOff", call: true, symEffect: "None"},
		{name: "LoweredClosureCall", argLength: 3, reg: regInfo{clobbers: r, inputs: []regMask{r, r, 0}}, aux: "Int64", call: true},
		{name: "LoweredInterCall",   argLength: 2, reg: regInfo{clobbers: r, inputs: []regMask{r} }, aux: "Int64", call: true},

		{name: "LoweredNilCheck", 	 argLength: 2, reg: regInfo{inputs: []regMask{r}}, nilCheck: true, faultOnNilArg0: true},
		{name: "LoweredAddr", 			 argLength: 1, reg: r11, aux: "SymOff", rematerializeable: true, symEffect: "Addr"},
		{name: "LoweredWB", 				 argLength: 3, reg: regInfo{inputs: []regMask{r, r}}, aux: "Sym", symEffect: "None"},
		{name: "LoweredGetG",				 argLength: 1, reg: r01},
		{name: "LoweredZero", 			 argLength: 2, reg: regInfo{inputs: []regMask{r}}, aux: "Int64"},
		{name: "LoweredMove", 			 argLength: 3, reg: regInfo{inputs: []regMask{r, r}}, aux: "Int64"},

		{name: "LoweredGetClosurePtr", 	reg: r01},                                                                          // returns wasm.REG_CTXT, the closure pointer
		{name: "LoweredGetCallerPC", 		reg: r01, rematerializeable: true},
		{name: "LoweredGetCallerSP", 		reg: r01, rematerializeable: true},

		// GenM instructions.
		{name: "ARG_I64",  		asm: "ARG",			reg: r31, 		typ: "Int64", argLength: 3},

		{name: "SELECT_I64",  asm: "SELECT",	reg: r31, 		typ: "Int64", argLength: 3},
		{name: "SELECT_I32",  asm: "SELECT",	reg: r31, 		typ: "Int32", argLength: 3},
		{name: "SELECT_I16",  asm: "SELECT",	reg: r31, 		typ: "Int16", argLength: 3},
		{name: "SELECT_I8",   asm: "SELECT",	reg: r31, 		typ: "Int8",	argLength: 3},

		{name: "CONST_I1",		asm: "CONST",   reg: r01,    	typ: "Int1",   		aux: "Bool",		rematerializeable: true},
		{name: "CONST_I8",		asm: "CONST",   reg: r01,    	typ: "Int8",   		aux: "Int8",		rematerializeable: true},
		{name: "CONST_I16",		asm: "CONST",   reg: r01,    	typ: "Int16",  		aux: "Int16",		rematerializeable: true},
		{name: "CONST_I32",		asm: "CONST",   reg: r01,    	typ: "Int32",  		aux: "Int32",		rematerializeable: true},
		{name: "CONST_I64",		asm: "CONST",   reg: r01,    	typ: "Int64",  		aux: "Int64",		rematerializeable: true},
		{name: "CONST_F32",		asm: "CONST",   reg: r01,    	typ: "Float32",		aux: "Float32", rematerializeable: true},
		{name: "CONST_F64",		asm: "CONST",   reg: r01,    	typ: "Float64",		aux: "Float64", rematerializeable: true},

		{name: "ADD_F64",			asm: "ADD",     reg: r21,    	typ: "Float64",		argLength: 2},
		{name: "ADD_I64",			asm: "ADD",     reg: r21,    	typ: "Int64",			argLength: 2},
		{name: "ADD_I32",			asm: "ADD",     reg: r21,    	typ: "Int32",			argLength: 2},
		{name: "ADD_I16",			asm: "ADD",     reg: r21,    	typ: "Int16",			argLength: 2},
		{name: "ADD_I8",			asm: "ADD",     reg: r21,    	typ: "Int8",			argLength: 2},

		{name: "SUB_F64",			asm: "SUB",     reg: r21,    	typ: "Float64",		argLength: 2},
		{name: "SUB_F32",			asm: "SUB",     reg: r21,    	typ: "Float32",		argLength: 2},
		{name: "SUB_I64",			asm: "SUB",     reg: r21,    	typ: "Int64",			argLength: 2},
		{name: "SUB_I32",			asm: "SUB",     reg: r21,    	typ: "Int32",			argLength: 2},
		{name: "SUB_I16",			asm: "SUB",     reg: r21,    	typ: "Int16",			argLength: 2},
		{name: "SUB_I8",			asm: "SUB",     reg: r21,    	typ: "Int8",			argLength: 2},

		{name: "MUL_I64",			asm: "MUL", 		reg: r21, 		typ: "Int64",    argLength: 2},
		{name: "MUL_I32",			asm: "MUL", 		reg: r21, 		typ: "Int32",    argLength: 2},
		{name: "MUL_I16",			asm: "MUL", 		reg: r21, 		typ: "Int16",    argLength: 2},
		{name: "MUL_I8",			asm: "MUL", 		reg: r21, 		typ: "Int8",     argLength: 2},
		{name: "MUL_U64",			asm: "MUL", 		reg: r21, 		typ: "Int64",    argLength: 2},
		{name: "MUL_U32",			asm: "MUL", 		reg: r21, 		typ: "Int32",    argLength: 2},
		{name: "MUL_F64",			asm: "MUL", 		reg: r21, 		typ: "Float64",  argLength: 2},

		{name: "DIV_I64",			asm: "DIV", 		reg: r21, 		typ: "Int64",    argLength: 2},
		{name: "DIV_I32",			asm: "DIV", 		reg: r21, 		typ: "Int32",    argLength: 2},
		{name: "DIV_U64",			asm: "DIV", 		reg: r21, 		typ: "Int64",    argLength: 2},
		{name: "DIV_U32",			asm: "DIV", 		reg: r21, 		typ: "Int32",    argLength: 2},
		{name: "DIV_U8",			asm: "DIV", 		reg: r21, 		typ: "Int8",   	 argLength: 2},
		{name: "DIV_F64",			asm: "DIV", 		reg: r21, 		typ: "Float64",  argLength: 2},

		{name: "MOD_I64",			asm: "MOD", 		reg: r21, 		typ: "Int64",    argLength: 2},
		{name: "MOD_I32",			asm: "MOD", 		reg: r21, 		typ: "Int32",    argLength: 2},
		{name: "MOD_U64",			asm: "MOD", 		reg: r21, 		typ: "Int64",    argLength: 2},
		{name: "MOD_U32",			asm: "MOD", 		reg: r21, 		typ: "Int32",    argLength: 2},

		{name: "AND_I64",			asm: "AND",			reg: r21, 		typ: "Int64",    argLength: 2},
		{name: "AND_I32",			asm: "AND",			reg: r21, 		typ: "Int32",    argLength: 2},
		{name: "AND_I16",			asm: "AND",			reg: r21, 		typ: "Int16",    argLength: 2},
		{name: "AND_I8",			asm: "AND",			reg: r21, 		typ: "Int8"	,    argLength: 2},
		{name: "AND_I1",			asm: "AND",			reg: r21, 		typ: "Bool"	,    argLength: 2},

		{name: "OR_I64",			asm: "OR", 			reg: r21, 		typ: "Int64",    argLength: 2},
		{name: "OR_I32",			asm: "OR", 			reg: r21, 		typ: "Int32",    argLength: 2},
		{name: "OR_I16",			asm: "OR", 			reg: r21, 		typ: "Int16",    argLength: 2},
		{name: "OR_I8",				asm: "OR", 			reg: r21, 		typ: "Int8",		 argLength: 2},
		{name: "OR_I1",				asm: "OR",			reg: r21, 		typ: "Bool"	,    argLength: 2},

		{name: "XOR_I64",			asm: "XOR",			reg: r21, 		typ: "Int64",    argLength: 2},
		{name: "XOR_I32",			asm: "XOR",			reg: r21, 		typ: "Int32",    argLength: 2},
		{name: "XOR_I16",			asm: "XOR",			reg: r21, 		typ: "Int16",    argLength: 2},
		{name: "XOR_I8",			asm: "XOR",			reg: r21, 		typ: "Int8",		 argLength: 2},

		{name: "NEG_I64",			asm: "NEG",			reg: r11, 		typ: "Int64",    argLength: 1},
		{name: "NEG_I32",			asm: "NEG",			reg: r11, 		typ: "Int32",    argLength: 1},
		{name: "NEG_F64",			asm: "NEG",			reg: r11, 		typ: "Float64",	 argLength: 1},
		{name: "NEG_F32",			asm: "NEG",			reg: r11, 		typ: "Float32",	 argLength: 1},

		{name: "NOT_I1",			asm: "NOT",			reg: r11, 		typ: "Bool",     argLength: 1},

		{name: "ZEXT_I64",		asm: "ZEXT",		reg: r11,    typ: "Int64",		 argLength: 1},
		{name: "ZEXT_I32",		asm: "ZEXT",		reg: r11,    typ: "Int32",		 argLength: 1},
		{name: "ZEXT_I16",		asm: "ZEXT",		reg: r11,    typ: "Int16",		 argLength: 1},

		{name: "SEXT_F64",		asm: "SEXT",		reg: r11,    typ: "Float64",	 argLength: 1},
		{name: "SEXT_I64",		asm: "SEXT",		reg: r11,    typ: "Int64",		 argLength: 1},
		{name: "SEXT_I32",		asm: "SEXT",		reg: r11,    typ: "Int32",		 argLength: 1},
		{name: "SEXT_I16",		asm: "SEXT",		reg: r11,    typ: "Int16",		 argLength: 1},
		{name: "FEXT_F64",  	asm: "FEXT", 		reg: r11, 	 typ: "Float64",   argLength: 1},

		{name: "TRUNC_I64",		asm: "TRUNC",		reg: r11,    typ: "Int64",		 argLength: 1},
		{name: "TRUNC_I32",		asm: "TRUNC",		reg: r11,    typ: "Int32",		 argLength: 1},
		{name: "TRUNC_I16",		asm: "TRUNC",		reg: r11,    typ: "Int16",		 argLength: 1},
		{name: "TRUNC_I8",		asm: "TRUNC",		reg: r11,    typ: "Int8",		 	 argLength: 1},
		{name: "TRUNC_F32",		asm: "TRUNC",		reg: r11,    typ: "Float32",		 argLength: 1},

		{name: "SAR_I64",			asm: "SAR",     reg: r21,    typ: "Int64",    argLength: 2},
		{name: "SAR_I32",			asm: "SAR",     reg: r21,    typ: "Int32",    argLength: 2},
		{name: "SAR_I16",			asm: "SAR",     reg: r21,    typ: "Int16",    argLength: 2},
		{name: "SAR_I8",			asm: "SAR",     reg: r21,    typ: "Int8",     argLength: 2},

		{name: "SLR_I64",			asm: "SLR",     reg: r21,    typ: "Int64",    argLength: 2},
		{name: "SLR_I32",			asm: "SLR",     reg: r21,    typ: "Int32",    argLength: 2},
		{name: "SLR_I16",			asm: "SLR",     reg: r21,    typ: "Int16",    argLength: 2},
		{name: "SLR_I8",			asm: "SLR",     reg: r21,    typ: "Int8",     argLength: 2},

		{name: "SLL_I64",			asm: "SLL",     reg: r21,    typ: "Int64",    argLength: 2},
		{name: "SLL_I32",			asm: "SLL",     reg: r21,    typ: "Int32",    argLength: 2},
		{name: "SLL_I16",			asm: "SLL",     reg: r21,    typ: "Int16",    argLength: 2},
		{name: "SLL_I8",			asm: "SLL",     reg: r21,    typ: "Int8",     argLength: 2},

		{name: "CMP_NE",			asm: "CMP_NE",  reg: r21,    typ: "Bool",     argLength: 2},
		{name: "CMP_EQ", 			asm: "CMP_EQ",  reg: r21,    typ: "Bool",     argLength: 2},
		{name: "CMP_GE", 			asm: "CMP_GE",  reg: r21,    typ: "Bool",     argLength: 2},
		{name: "CMP_LT",			asm: "CMP_LT",  reg: r21,    typ: "Bool",     argLength: 2},
		{name: "CMP_GT",			asm: "CMP_GT",  reg: r21,    typ: "Bool",     argLength: 2},
		{name: "CMP_LE",			asm: "CMP_LE",  reg: r21,    typ: "Bool",     argLength: 2},
		{name: "CMP_ULT",			asm: "CMP_ULT", reg: r21,    typ: "Bool",     argLength: 2},
		{name: "CMP_UGT",			asm: "CMP_UGT", reg: r21,    typ: "Bool",     argLength: 2},
		{name: "CMP_UGE",			asm: "CMP_UGE", reg: r21,    typ: "Bool",     argLength: 2},
		{name: "CMP_ULE",			asm: "CMP_ULE", reg: r21,    typ: "Bool",     argLength: 2},
		{name: "CMP_ONE",			asm: "CMP_ONE",	reg: r21,    typ: "Bool",     argLength: 2},
		{name: "CMP_OLT",			asm: "CMP_OLT",	reg: r21,    typ: "Bool",     argLength: 2},
		{name: "CMP_OGE",			asm: "CMP_OGE", reg: r21,    typ: "Bool",     argLength: 2},
		{name: "CMP_OGT",			asm: "CMP_OGT", reg: r21,    typ: "Bool",     argLength: 2},
		{name: "CMP_OLE",			asm: "CMP_OLE", reg: r21,    typ: "Bool",     argLength: 2},

		{name: "LD_8_F64",		asm: "LD", 		  reg: rload,  typ: "Float64",  argLength: 2, aux: "Int64"},
		{name: "LD_4_F32",		asm: "LD", 		  reg: rload,  typ: "Float32",  argLength: 2, aux: "Int64"},
		{name: "LD_8_U64",		asm: "LD", 		  reg: rload,  typ: "UInt64",   argLength: 2, aux: "Int64"},
		{name: "LD_8_I64",		asm: "LD", 		  reg: rload,  typ: "Int64",    argLength: 2, aux: "Int64"},
		{name: "LD_4_U32",		asm: "LD", 		  reg: rload,  typ: "UInt32",   argLength: 2, aux: "Int64"},
		{name: "LD_4_I32",		asm: "LD", 		  reg: rload,  typ: "Int32",    argLength: 2, aux: "Int64"},
		{name: "LD_2_U16",		asm: "LD", 		  reg: rload,  typ: "UInt16",   argLength: 2, aux: "Int64"},
		{name: "LD_2_I16",		asm: "LD", 		  reg: rload,  typ: "Int16",    argLength: 2, aux: "Int64"},
		{name: "LD_1_U8",			asm: "LD", 		  reg: rload,  typ: "UInt8",    argLength: 2, aux: "Int64"},
		{name: "LD_1_I8",			asm: "LD", 		  reg: rload,  typ: "Int8",     argLength: 2, aux: "Int64"},

		{name: "ST_8",				asm: "ST_8",		reg: rstore, typ: "Mem",      argLength: 3, aux: "Int64"},
		{name: "ST_4",				asm: "ST_4",		reg: rstore, typ: "Mem",      argLength: 3, aux: "Int64"},
		{name: "ST_2",				asm: "ST_2",		reg: rstore, typ: "Mem",      argLength: 3, aux: "Int64"},
		{name: "ST_1",				asm: "ST_1",		reg: rstore, typ: "Mem",      argLength: 3, aux: "Int64"},
	}

	archs = append(archs, arch{
		name:            "GenM",
		pkg:             "cmd/internal/obj/genm",
		genfile:         "../../genm/ssa.go",
		ops:             GenMOps,
		blocks:          nil,
		regnames:        regNamesGenM,
		gpregmask:       r,
		fpregmask:       r,
		framepointerreg: -1, // not used
		linkreg:         -1, // not used
	})
}
