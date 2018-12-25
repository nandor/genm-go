// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package genm

import (
	"cmd/compile/internal/gc"
	"cmd/compile/internal/ssa"
	"cmd/internal/obj"
	"cmd/internal/obj/genm"
	"log"
)

// markMoves marks any MOVXconst ops that need to avoid clobbering flags.
func ssaMarkMoves(s *gc.SSAGenState, b *ssa.Block) {
}

func ssaGenValue(s *gc.SSAGenState, v *ssa.Value) {
	switch v.Op {
	case ssa.OpGenMLoweredStaticCall:
		log.Printf("not implemented: lowered static call")

	case ssa.OpGenMLoweredInterCall:
		v.Fatalf("ssaGenValue not implemented: %s\n%s", v.Op, v.Block.Func)

	case ssa.OpGenMLoweredWB:
		v.Fatalf("ssaGenValue not implemented: %s\n%s", v.Op, v.Block.Func)

	case ssa.OpGenMLoweredAddr:
		p := s.Prog(genm.AADDR)
		p.Ty = obj.TYPE_I64
		p.From.Type = obj.TYPE_ADDR
		gc.AddAux(&p.From, v)

	case ssa.OpGenMLoweredNilCheck:
		p := s.Prog(genm.ALD)
		p.Ty = obj.TYPE_I8
		p.Size = 1
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg =  genm.REG_T0

	case ssa.OpGenMSELECT_I64:
		panic("Not implemented")

	case ssa.OpGenMCMP_EQ_I8:
		ssaGenBinary(s, v)

	case ssa.OpGenMCONST_I64, ssa.OpGenMCONST_I32, ssa.OpGenMCONST_I16, ssa.OpGenMCONST_I8, ssa.OpGenMCONST_I1:
		p := s.Prog(genm.ACONST)
		p.Ty = getOpType(v)
		p.From = obj.Addr{Type: obj.TYPE_CONST, Offset: v.AuxInt}

	case ssa.OpGenMADD_I64, ssa.OpGenMAND_I64, ssa.OpGenMOR_I64:
		ssaGenBinary(s, v)

	case ssa.OpGenMSLL_I64, ssa.OpGenMSLL_I32, ssa.OpGenMSLL_I16, ssa.OpGenMSLL_I8:
		ssaGenBinary(s, v)

	case ssa.OpGenMSLR_I64, ssa.OpGenMSLR_I32, ssa.OpGenMSLR_I16, ssa.OpGenMSLR_I8:
		ssaGenBinary(s, v)

	case ssa.OpGenMCMP_NE_I64, ssa.OpGenMCMP_NE_I32:
		ssaGenBinary(s, v)

	case ssa.OpGenMCMP_LT_I64, ssa.OpGenMCMP_LT_I32, ssa.OpGenMCMP_LT_I16, ssa.OpGenMCMP_LT_I8, ssa.OpGenMCMP_ULT_I64, ssa.OpGenMCMP_ULT_I32, ssa.OpGenMCMP_ULT_I16, ssa.OpGenMCMP_ULT_I8, ssa.OpGenMCMP_OLT_F64, ssa.OpGenMCMP_OLT_F32:
		ssaGenBinary(s, v)

	case ssa.OpGenMCMP_GT_I64, ssa.OpGenMCMP_GT_I32, ssa.OpGenMCMP_GT_I16, ssa.OpGenMCMP_GT_I8, ssa.OpGenMCMP_UGT_I64, ssa.OpGenMCMP_UGT_I32, ssa.OpGenMCMP_UGT_I16, ssa.OpGenMCMP_UGT_I8, ssa.OpGenMCMP_OGT_F64, ssa.OpGenMCMP_OGT_F32:
		ssaGenBinary(s, v)

	case ssa.OpGenMLD_8_U64, ssa.OpGenMLD_8_I64, ssa.OpGenMLD_4_U32, ssa.OpGenMLD_4_I32, ssa.OpGenMLD_2_U16, ssa.OpGenMLD_2_I16, ssa.OpGenMLD_1_U8, ssa.OpGenMLD_1_I8, ssa.OpGenMLD_4_F32, ssa.OpGenMLD_8_F64:
		if v.AuxInt != 0 {
			panic("Not implemented")
		}
		p := s.Prog(v.Op.Asm())
		p.Ty = getOpType(v)
		p.Size = uint8(v.Type.Size())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()

	case ssa.OpGenMST_8_I, ssa.OpGenMST_4_I, ssa.OpGenMST_2_I, ssa.OpGenMST_1_I, ssa.OpGenMST_8_F, ssa.OpGenMST_4_F:
		if v.AuxInt != 0 {
			panic("Not implemented")
		}
		p := s.Prog(v.Op.Asm())
		p.Ty = obj.TYPE_VOID
		switch v.Op {
			case ssa.OpGenMST_8_I: p.Size = 8
			case ssa.OpGenMST_4_I: p.Size = 4
			case ssa.OpGenMST_2_I: p.Size = 2
			case ssa.OpGenMST_1_I: p.Size = 1
			case ssa.OpGenMST_8_F: p.Size = 8
			case ssa.OpGenMST_4_F: p.Size = 4
		}
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Args[0].Reg()

	case ssa.OpCopy:
		p := s.Prog(genm.ACOPY)
		p.Ty = getOpType(v)
		p.From = obj.Addr{Type: obj.TYPE_REG, Reg: v.Reg()}

	case ssa.OpLoadReg:
		ssaGenLoadReg(s, v)

	default:
		v.Fatalf("ssaGenValue not implemented: %s in %s", v.Op, v.Block.Func)
	}
}

func ssaGenBlock(s *gc.SSAGenState, b, next *ssa.Block) {
	switch b.Kind {
	case ssa.BlockPlain:
		panic("BlockPlain")

	case ssa.BlockIf:
		log.Printf("BlockIf: not implemented")
		s.Prog(genm.AJT)

	case ssa.BlockRet:
		log.Printf("BlockRet: not implemented")
		s.Prog(genm.ARET)

	case ssa.BlockRetJmp:
		panic("BlockRetJmp")

	case ssa.BlockExit:
		s.Prog(obj.AUNDEF)

	case ssa.BlockDefer:
		panic("BlockDefer")
	default:
		next.Fatalf("unexpected block: %d", b.Kind)
	}
}

func getOpType(v *ssa.Value) obj.GenMType {
	switch v.Type.Size() {
	case 8:
		if (v.Type.IsFloat()) {
			return obj.TYPE_F64
		}
		if (v.Type.IsInteger()) {
			if (v.Type.IsSigned()) {
				return obj.TYPE_I64
			} else {
				return obj.TYPE_I64
			}
		}
	case 4:
		if (v.Type.IsFloat()) {
			return obj.TYPE_F32
		}
		if (v.Type.IsInteger()) {
			if (v.Type.IsSigned()) {
				return obj.TYPE_I32
			} else {
				return obj.TYPE_I32
			}
		}
	case 2:
		if (v.Type.IsInteger()) {
			if (v.Type.IsSigned()) {
				return obj.TYPE_I16
			} else {
				return obj.TYPE_U16
			}
		}
	case 1:
		if (v.Type.IsInteger()) {
			if (v.Type.IsSigned()) {
				return obj.TYPE_I8
			} else {
				return obj.TYPE_U8
			}
		}
		if (v.Type.IsBoolean()) {
			return obj.TYPE_I1
		}
	}

	v.Fatalf("invalid value type %s, size %d", v.Type.SimpleString(), v.Type.Size())
	return obj.TYPE_VOID
}

func ssaGenBinary(s *gc.SSAGenState, v *ssa.Value) {
	p := s.Prog(v.Op.Asm())
	p.Ty = getOpType(v)
	p.From.Type = obj.TYPE_REG
	p.From.Reg = v.Args[0].Reg()
	p.Reg = v.Args[1].Reg()
	p.To.Type = obj.TYPE_REG
	p.To.Reg = v.Reg()
}

func ssaGenLoadReg(s *gc.SSAGenState, v *ssa.Value) {
	arg := v.Args[0]

	switch arg.Op {
	case ssa.OpArg:
		p := s.Prog(genm.AARG)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()

	default:
		v.Fatalf("unexpected argument to LoadReg")
	}
}
