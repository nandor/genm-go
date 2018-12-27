// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package genm

import (
	"cmd/compile/internal/gc"
	"cmd/compile/internal/ssa"
	"cmd/internal/obj"
	"cmd/internal/obj/genm"
)

// markMoves marks any MOVXconst ops that need to avoid clobbering flags.
func ssaMarkMoves(s *gc.SSAGenState, b *ssa.Block) {
}

func ssaGenValue(s *gc.SSAGenState, v *ssa.Value) {
	switch v.Op {
	case ssa.OpGenMLoweredStaticCall:
		v.Logf("ssaGenValue not implemented: lowered static call\n")

	case ssa.OpGenMLoweredInterCall:
		v.Logf("ssaGenValue not implemented: lowered inter call\n")

	case ssa.OpGenMLoweredClosureCall:
		v.Logf("ssaGenValue not implemented: lowered closure call\n")

	case ssa.OpGenMLoweredWB:
		v.Logf("ssaGenValue not implemented: lowered writeback\n")

	case ssa.OpGenMLoweredGetG:
		v.Logf("ssaGenValue not implemented: lowered writeback\n")

	case ssa.OpGenMLoweredZero:
		v.Logf("ssaGenValue not implemented: lowered zeroing\n")

	case ssa.OpGenMLoweredMove:
		v.Logf("ssaGenValue not implemented: lowered moving\n")

	case ssa.OpGenMLoweredGetCallerPC:
		v.Logf("ssaGenValue not implemented: lowered get caller PC\n")

	case ssa.OpGenMLoweredGetCallerSP:
		v.Logf("ssaGenValue not implemented: lowered get caller SP\n")

	case ssa.OpGenMLoweredGetClosurePtr:
		v.Logf("ssaGenValue not implemented: lowered get closure ptr\n")

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

	case ssa.OpGenMARG_I64:
		v.Logf("ssaGenValue not implemented: ARG\n")

	case ssa.OpGenMSELECT_I64, ssa.OpGenMSELECT_I32, ssa.OpGenMSELECT_I16, ssa.OpGenMSELECT_I8:
		v.Logf("ssaGenValue not implemented: SELECT\n")

	case ssa.OpGenMCONST_I64, ssa.OpGenMCONST_I32, ssa.OpGenMCONST_I16, ssa.OpGenMCONST_I8, ssa.OpGenMCONST_I1:
		p := s.Prog(genm.ACONST)
		p.Ty = getOpType(v)
		p.From = obj.Addr{Type: obj.TYPE_CONST, Offset: v.AuxInt}

	case ssa.OpGenMCONST_F32, ssa.OpGenMCONST_F64:
		p := s.Prog(genm.ACONST)
		p.Ty = getOpType(v)
		p.From = obj.Addr{Type: obj.TYPE_FCONST, Val: v.AuxFloat()}

	case ssa.OpGenMDIV_F64, ssa.OpGenMMUL_F64:
		ssaGenBinary(s, v)

	case ssa.OpGenMMUL_I32, ssa.OpGenMMUL_U32, ssa.OpGenMDIV_I32, ssa.OpGenMDIV_U32, ssa.OpGenMMOD_I32, ssa.OpGenMMOD_U32:
		ssaGenBinary(s, v)

	case ssa.OpGenMMUL_I64, ssa.OpGenMMUL_U64, ssa.OpGenMDIV_I64, ssa.OpGenMDIV_U64, ssa.OpGenMMOD_I64, ssa.OpGenMMOD_U64:
		ssaGenBinary(s, v)

	case ssa.OpGenMMUL_I16:
		ssaGenBinary(s, v)

	case ssa.OpGenMMUL_I8, ssa.OpGenMDIV_U8:
		ssaGenBinary(s, v)

	case ssa.OpGenMADD_F64, ssa.OpGenMSUB_F64:
		ssaGenBinary(s, v)

	case ssa.OpGenMSUB_F32:
		ssaGenBinary(s, v)

	case ssa.OpGenMADD_I64, ssa.OpGenMSUB_I64, ssa.OpGenMAND_I64, ssa.OpGenMOR_I64, ssa.OpGenMXOR_I64:
		ssaGenBinary(s, v)

	case ssa.OpGenMADD_I32, ssa.OpGenMSUB_I32, ssa.OpGenMAND_I32, ssa.OpGenMOR_I32, ssa.OpGenMXOR_I32:
		ssaGenBinary(s, v)

	case ssa.OpGenMADD_I16, ssa.OpGenMSUB_I16, ssa.OpGenMAND_I16, ssa.OpGenMOR_I16, ssa.OpGenMXOR_I16:
		ssaGenBinary(s, v)

	case ssa.OpGenMADD_I8, ssa.OpGenMSUB_I8, ssa.OpGenMAND_I8, ssa.OpGenMOR_I8, ssa.OpGenMXOR_I8:
		ssaGenBinary(s, v)

	case ssa.OpGenMAND_I1, ssa.OpGenMOR_I1:
		ssaGenBinary(s, v)

	case ssa.OpGenMSEXT_I64, ssa.OpGenMSEXT_I32, ssa.OpGenMSEXT_I16, ssa.OpGenMSEXT_F64:
		ssaGenUnary(s, v)

	case ssa.OpGenMZEXT_I64, ssa.OpGenMZEXT_I32, ssa.OpGenMZEXT_I16:
		ssaGenUnary(s, v)

	case ssa.OpGenMTRUNC_I64, ssa.OpGenMTRUNC_I32, ssa.OpGenMTRUNC_I16, ssa.OpGenMTRUNC_I8, ssa.OpGenMTRUNC_F32:
		ssaGenUnary(s, v)

	case ssa.OpGenMFEXT_F64:
		ssaGenUnary(s, v)

	case ssa.OpGenMNEG_I64, ssa.OpGenMNEG_I32, ssa.OpGenMNEG_F64, ssa.OpGenMNEG_F32:
		ssaGenUnary(s, v)

	case ssa.OpGenMNOT_I1:
		ssaGenUnary(s, v)

	case ssa.OpGenMSAR_I64, ssa.OpGenMSAR_I32, ssa.OpGenMSAR_I16, ssa.OpGenMSAR_I8:
		ssaGenBinary(s, v)

	case ssa.OpGenMSLR_I64, ssa.OpGenMSLR_I32, ssa.OpGenMSLR_I16, ssa.OpGenMSLR_I8:
		ssaGenBinary(s, v)

	case ssa.OpGenMSLL_I64, ssa.OpGenMSLL_I32, ssa.OpGenMSLL_I16, ssa.OpGenMSLL_I8:
		ssaGenBinary(s, v)

	case ssa.OpGenMCMP_EQ, ssa.OpGenMCMP_NE, ssa.OpGenMCMP_GE, ssa.OpGenMCMP_LT, ssa.OpGenMCMP_GT, ssa.OpGenMCMP_LE:
		ssaGenBinary(s, v)

	case ssa.OpGenMCMP_UGE, ssa.OpGenMCMP_ULT, ssa.OpGenMCMP_UGT, ssa.OpGenMCMP_ULE:
		ssaGenBinary(s, v)

	case ssa.OpGenMCMP_OGE, ssa.OpGenMCMP_OGT, ssa.OpGenMCMP_OLT, ssa.OpGenMCMP_OLE, ssa.OpGenMCMP_ONE:
		ssaGenBinary(s, v)

	case ssa.OpGenMLD_8_U64, ssa.OpGenMLD_8_I64, ssa.OpGenMLD_4_U32, ssa.OpGenMLD_4_I32, ssa.OpGenMLD_2_U16, ssa.OpGenMLD_2_I16, ssa.OpGenMLD_1_U8, ssa.OpGenMLD_1_I8, ssa.OpGenMLD_4_F32, ssa.OpGenMLD_8_F64:
		if v.AuxInt != 0 {
			panic("Not implemented: LD")
		}
		p := s.Prog(v.Op.Asm())
		p.Ty = getOpType(v)
		p.Size = uint8(v.Type.Size())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()

	case ssa.OpGenMST_8, ssa.OpGenMST_4, ssa.OpGenMST_2, ssa.OpGenMST_1:
		if v.AuxInt != 0 {
			panic("Not implemented: ST")
		}
		p := s.Prog(v.Op.Asm())
		p.Ty = obj.TYPE_VOID
		switch v.Op {
			case ssa.OpGenMST_8: p.Size = 8
			case ssa.OpGenMST_4: p.Size = 4
			case ssa.OpGenMST_2: p.Size = 2
			case ssa.OpGenMST_1: p.Size = 1
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

	case ssa.OpStoreReg:
		ssaGenStoreReg(s, v)

	default:
		v.Fatalf("ssaGenValue not implemented: %s in %s\n", v.Op, v.Block.Func)
	}
}

func ssaGenBlock(s *gc.SSAGenState, b, next *ssa.Block) {
	switch b.Kind {
	case ssa.BlockPlain:
		b.Logf("ssaGenBlock BlockPlain: not implemented\n")
		s.Prog(genm.AJMP)

	case ssa.BlockIf:
		b.Logf("ssaGenBlock BlockIf: not implemented\n")
		s.Prog(genm.AJT)
		s.Prog(genm.AJMP)

	case ssa.BlockRet:
		b.Logf("ssaGenBlock BlockRet: not implemented\n")
		s.Prog(genm.ARET)

	case ssa.BlockRetJmp:
		b.Logf("ssaGenBlock BlockRetJmp: not implemented\n")
		s.Prog(genm.ARET)

	case ssa.BlockExit:
		s.Prog(obj.AUNDEF)

	case ssa.BlockDefer:
		b.Logf("ssaGenBlock BlockDefer: not implemented\n")

	default:
		b.Fatalf("unexpected block: %d", b.Kind)
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
				return obj.TYPE_U64
			}
		}
		if (v.Type.IsPtrShaped()) {
			return obj.TYPE_I64
		}
	case 4:
		if (v.Type.IsFloat()) {
			return obj.TYPE_F32
		}
		if (v.Type.IsInteger()) {
			if (v.Type.IsSigned()) {
				return obj.TYPE_I32
			} else {
				return obj.TYPE_U32
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

	v.Fatalf("invalid value type %s, size %d\n", v.Type.SimpleString(), v.Type.Size())
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

func ssaGenUnary(s *gc.SSAGenState, v *ssa.Value) {
	p := s.Prog(v.Op.Asm())
	p.Ty = getOpType(v)
	p.From.Type = obj.TYPE_REG
	p.From.Reg = v.Args[0].Reg()
	p.To.Type = obj.TYPE_REG
	p.To.Reg = v.Reg()
}

func ssaGenLoadReg(s *gc.SSAGenState, v *ssa.Value) {
	arg := v.Args[0]

	switch arg.Op {
	case ssa.OpArg:
		// Create an argument instruction.
		p := s.Prog(genm.AARG)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()

	case ssa.OpStoreReg:
		// Stack restore - move from old temp register.
		v.Logf("StoreReg to LoadReg: %s\n", arg.LongString())

	case ssa.OpPhi:
		// Stack restore - move from old temp register.
		v.Logf("Phi to LoadReg: %s\n", arg.LongString())

	default:
		v.Fatalf("unexpected argument to LoadReg: %s\n%s\n", arg.LongString(), arg.Block.Func)
	}
}

func ssaGenStoreReg(s *gc.SSAGenState, v *ssa.Value) {
	v.Logf("not implemented: StoreReg\n")
}
