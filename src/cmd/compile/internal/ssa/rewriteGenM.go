// Code generated from gen/GenM.rules; DO NOT EDIT.
// generated with: cd gen; go run *.go

package ssa

import "fmt"
import "math"
import "cmd/internal/obj"
import "cmd/internal/objabi"
import "cmd/compile/internal/types"

var _ = fmt.Println   // in case not otherwise used
var _ = math.MinInt8  // in case not otherwise used
var _ = obj.ANOP      // in case not otherwise used
var _ = objabi.GOROOT // in case not otherwise used
var _ = types.TypeMem // in case not otherwise used

func rewriteValueGenM(v *Value) bool {
	switch v.Op {
	case OpAdd16:
		return rewriteValueGenM_OpAdd16_0(v)
	case OpAdd32:
		return rewriteValueGenM_OpAdd32_0(v)
	case OpAdd64:
		return rewriteValueGenM_OpAdd64_0(v)
	case OpAdd64F:
		return rewriteValueGenM_OpAdd64F_0(v)
	case OpAdd8:
		return rewriteValueGenM_OpAdd8_0(v)
	case OpAddPtr:
		return rewriteValueGenM_OpAddPtr_0(v)
	case OpAddr:
		return rewriteValueGenM_OpAddr_0(v)
	case OpAnd16:
		return rewriteValueGenM_OpAnd16_0(v)
	case OpAnd32:
		return rewriteValueGenM_OpAnd32_0(v)
	case OpAnd64:
		return rewriteValueGenM_OpAnd64_0(v)
	case OpAnd8:
		return rewriteValueGenM_OpAnd8_0(v)
	case OpAndB:
		return rewriteValueGenM_OpAndB_0(v)
	case OpClosureCall:
		return rewriteValueGenM_OpClosureCall_0(v)
	case OpCom16:
		return rewriteValueGenM_OpCom16_0(v)
	case OpCom32:
		return rewriteValueGenM_OpCom32_0(v)
	case OpCom64:
		return rewriteValueGenM_OpCom64_0(v)
	case OpCom8:
		return rewriteValueGenM_OpCom8_0(v)
	case OpConst16:
		return rewriteValueGenM_OpConst16_0(v)
	case OpConst32:
		return rewriteValueGenM_OpConst32_0(v)
	case OpConst32F:
		return rewriteValueGenM_OpConst32F_0(v)
	case OpConst64:
		return rewriteValueGenM_OpConst64_0(v)
	case OpConst64F:
		return rewriteValueGenM_OpConst64F_0(v)
	case OpConst8:
		return rewriteValueGenM_OpConst8_0(v)
	case OpConstBool:
		return rewriteValueGenM_OpConstBool_0(v)
	case OpConstNil:
		return rewriteValueGenM_OpConstNil_0(v)
	case OpCvt32Fto64:
		return rewriteValueGenM_OpCvt32Fto64_0(v)
	case OpCvt32Fto64F:
		return rewriteValueGenM_OpCvt32Fto64F_0(v)
	case OpCvt32to64F:
		return rewriteValueGenM_OpCvt32to64F_0(v)
	case OpCvt64Fto32:
		return rewriteValueGenM_OpCvt64Fto32_0(v)
	case OpCvt64Fto32F:
		return rewriteValueGenM_OpCvt64Fto32F_0(v)
	case OpCvt64Fto64:
		return rewriteValueGenM_OpCvt64Fto64_0(v)
	case OpCvt64to64F:
		return rewriteValueGenM_OpCvt64to64F_0(v)
	case OpDiv32:
		return rewriteValueGenM_OpDiv32_0(v)
	case OpDiv32u:
		return rewriteValueGenM_OpDiv32u_0(v)
	case OpDiv64:
		return rewriteValueGenM_OpDiv64_0(v)
	case OpDiv64F:
		return rewriteValueGenM_OpDiv64F_0(v)
	case OpDiv64u:
		return rewriteValueGenM_OpDiv64u_0(v)
	case OpDiv8u:
		return rewriteValueGenM_OpDiv8u_0(v)
	case OpEq16:
		return rewriteValueGenM_OpEq16_0(v)
	case OpEq32:
		return rewriteValueGenM_OpEq32_0(v)
	case OpEq32F:
		return rewriteValueGenM_OpEq32F_0(v)
	case OpEq64:
		return rewriteValueGenM_OpEq64_0(v)
	case OpEq64F:
		return rewriteValueGenM_OpEq64F_0(v)
	case OpEq8:
		return rewriteValueGenM_OpEq8_0(v)
	case OpEqB:
		return rewriteValueGenM_OpEqB_0(v)
	case OpEqPtr:
		return rewriteValueGenM_OpEqPtr_0(v)
	case OpGeq16:
		return rewriteValueGenM_OpGeq16_0(v)
	case OpGeq16U:
		return rewriteValueGenM_OpGeq16U_0(v)
	case OpGeq32:
		return rewriteValueGenM_OpGeq32_0(v)
	case OpGeq32F:
		return rewriteValueGenM_OpGeq32F_0(v)
	case OpGeq32U:
		return rewriteValueGenM_OpGeq32U_0(v)
	case OpGeq64:
		return rewriteValueGenM_OpGeq64_0(v)
	case OpGeq64F:
		return rewriteValueGenM_OpGeq64F_0(v)
	case OpGeq64U:
		return rewriteValueGenM_OpGeq64U_0(v)
	case OpGeq8:
		return rewriteValueGenM_OpGeq8_0(v)
	case OpGeq8U:
		return rewriteValueGenM_OpGeq8U_0(v)
	case OpGetCallerPC:
		return rewriteValueGenM_OpGetCallerPC_0(v)
	case OpGetCallerSP:
		return rewriteValueGenM_OpGetCallerSP_0(v)
	case OpGetClosurePtr:
		return rewriteValueGenM_OpGetClosurePtr_0(v)
	case OpGetG:
		return rewriteValueGenM_OpGetG_0(v)
	case OpGreater16:
		return rewriteValueGenM_OpGreater16_0(v)
	case OpGreater16U:
		return rewriteValueGenM_OpGreater16U_0(v)
	case OpGreater32:
		return rewriteValueGenM_OpGreater32_0(v)
	case OpGreater32F:
		return rewriteValueGenM_OpGreater32F_0(v)
	case OpGreater32U:
		return rewriteValueGenM_OpGreater32U_0(v)
	case OpGreater64:
		return rewriteValueGenM_OpGreater64_0(v)
	case OpGreater64F:
		return rewriteValueGenM_OpGreater64F_0(v)
	case OpGreater64U:
		return rewriteValueGenM_OpGreater64U_0(v)
	case OpGreater8:
		return rewriteValueGenM_OpGreater8_0(v)
	case OpGreater8U:
		return rewriteValueGenM_OpGreater8U_0(v)
	case OpInterCall:
		return rewriteValueGenM_OpInterCall_0(v)
	case OpIsInBounds:
		return rewriteValueGenM_OpIsInBounds_0(v)
	case OpIsNonNil:
		return rewriteValueGenM_OpIsNonNil_0(v)
	case OpIsSliceInBounds:
		return rewriteValueGenM_OpIsSliceInBounds_0(v)
	case OpLeq16:
		return rewriteValueGenM_OpLeq16_0(v)
	case OpLeq16U:
		return rewriteValueGenM_OpLeq16U_0(v)
	case OpLeq32:
		return rewriteValueGenM_OpLeq32_0(v)
	case OpLeq32F:
		return rewriteValueGenM_OpLeq32F_0(v)
	case OpLeq32U:
		return rewriteValueGenM_OpLeq32U_0(v)
	case OpLeq64:
		return rewriteValueGenM_OpLeq64_0(v)
	case OpLeq64F:
		return rewriteValueGenM_OpLeq64F_0(v)
	case OpLeq64U:
		return rewriteValueGenM_OpLeq64U_0(v)
	case OpLeq8:
		return rewriteValueGenM_OpLeq8_0(v)
	case OpLeq8U:
		return rewriteValueGenM_OpLeq8U_0(v)
	case OpLess16:
		return rewriteValueGenM_OpLess16_0(v)
	case OpLess16U:
		return rewriteValueGenM_OpLess16U_0(v)
	case OpLess32:
		return rewriteValueGenM_OpLess32_0(v)
	case OpLess32F:
		return rewriteValueGenM_OpLess32F_0(v)
	case OpLess32U:
		return rewriteValueGenM_OpLess32U_0(v)
	case OpLess64:
		return rewriteValueGenM_OpLess64_0(v)
	case OpLess64F:
		return rewriteValueGenM_OpLess64F_0(v)
	case OpLess64U:
		return rewriteValueGenM_OpLess64U_0(v)
	case OpLess8:
		return rewriteValueGenM_OpLess8_0(v)
	case OpLess8U:
		return rewriteValueGenM_OpLess8U_0(v)
	case OpLoad:
		return rewriteValueGenM_OpLoad_0(v)
	case OpLocalAddr:
		return rewriteValueGenM_OpLocalAddr_0(v)
	case OpLsh16x16:
		return rewriteValueGenM_OpLsh16x16_0(v)
	case OpLsh16x32:
		return rewriteValueGenM_OpLsh16x32_0(v)
	case OpLsh16x64:
		return rewriteValueGenM_OpLsh16x64_0(v)
	case OpLsh16x8:
		return rewriteValueGenM_OpLsh16x8_0(v)
	case OpLsh32x16:
		return rewriteValueGenM_OpLsh32x16_0(v)
	case OpLsh32x32:
		return rewriteValueGenM_OpLsh32x32_0(v)
	case OpLsh32x64:
		return rewriteValueGenM_OpLsh32x64_0(v)
	case OpLsh32x8:
		return rewriteValueGenM_OpLsh32x8_0(v)
	case OpLsh64x16:
		return rewriteValueGenM_OpLsh64x16_0(v)
	case OpLsh64x32:
		return rewriteValueGenM_OpLsh64x32_0(v)
	case OpLsh64x64:
		return rewriteValueGenM_OpLsh64x64_0(v)
	case OpLsh64x8:
		return rewriteValueGenM_OpLsh64x8_0(v)
	case OpLsh8x16:
		return rewriteValueGenM_OpLsh8x16_0(v)
	case OpLsh8x32:
		return rewriteValueGenM_OpLsh8x32_0(v)
	case OpLsh8x64:
		return rewriteValueGenM_OpLsh8x64_0(v)
	case OpLsh8x8:
		return rewriteValueGenM_OpLsh8x8_0(v)
	case OpMod32:
		return rewriteValueGenM_OpMod32_0(v)
	case OpMod32u:
		return rewriteValueGenM_OpMod32u_0(v)
	case OpMod64:
		return rewriteValueGenM_OpMod64_0(v)
	case OpMod64u:
		return rewriteValueGenM_OpMod64u_0(v)
	case OpMove:
		return rewriteValueGenM_OpMove_0(v)
	case OpMul16:
		return rewriteValueGenM_OpMul16_0(v)
	case OpMul32:
		return rewriteValueGenM_OpMul32_0(v)
	case OpMul64:
		return rewriteValueGenM_OpMul64_0(v)
	case OpMul64F:
		return rewriteValueGenM_OpMul64F_0(v)
	case OpMul8:
		return rewriteValueGenM_OpMul8_0(v)
	case OpNeg32:
		return rewriteValueGenM_OpNeg32_0(v)
	case OpNeg32F:
		return rewriteValueGenM_OpNeg32F_0(v)
	case OpNeg64:
		return rewriteValueGenM_OpNeg64_0(v)
	case OpNeg64F:
		return rewriteValueGenM_OpNeg64F_0(v)
	case OpNeq16:
		return rewriteValueGenM_OpNeq16_0(v)
	case OpNeq32:
		return rewriteValueGenM_OpNeq32_0(v)
	case OpNeq32F:
		return rewriteValueGenM_OpNeq32F_0(v)
	case OpNeq64:
		return rewriteValueGenM_OpNeq64_0(v)
	case OpNeq64F:
		return rewriteValueGenM_OpNeq64F_0(v)
	case OpNeq8:
		return rewriteValueGenM_OpNeq8_0(v)
	case OpNeqB:
		return rewriteValueGenM_OpNeqB_0(v)
	case OpNeqPtr:
		return rewriteValueGenM_OpNeqPtr_0(v)
	case OpNilCheck:
		return rewriteValueGenM_OpNilCheck_0(v)
	case OpNot:
		return rewriteValueGenM_OpNot_0(v)
	case OpOffPtr:
		return rewriteValueGenM_OpOffPtr_0(v)
	case OpOr16:
		return rewriteValueGenM_OpOr16_0(v)
	case OpOr32:
		return rewriteValueGenM_OpOr32_0(v)
	case OpOr64:
		return rewriteValueGenM_OpOr64_0(v)
	case OpOr8:
		return rewriteValueGenM_OpOr8_0(v)
	case OpOrB:
		return rewriteValueGenM_OpOrB_0(v)
	case OpRsh16x64:
		return rewriteValueGenM_OpRsh16x64_0(v)
	case OpRsh32Ux16:
		return rewriteValueGenM_OpRsh32Ux16_0(v)
	case OpRsh32Ux32:
		return rewriteValueGenM_OpRsh32Ux32_0(v)
	case OpRsh32Ux64:
		return rewriteValueGenM_OpRsh32Ux64_0(v)
	case OpRsh32Ux8:
		return rewriteValueGenM_OpRsh32Ux8_0(v)
	case OpRsh32x64:
		return rewriteValueGenM_OpRsh32x64_0(v)
	case OpRsh64Ux16:
		return rewriteValueGenM_OpRsh64Ux16_0(v)
	case OpRsh64Ux32:
		return rewriteValueGenM_OpRsh64Ux32_0(v)
	case OpRsh64Ux64:
		return rewriteValueGenM_OpRsh64Ux64_0(v)
	case OpRsh64Ux8:
		return rewriteValueGenM_OpRsh64Ux8_0(v)
	case OpRsh64x64:
		return rewriteValueGenM_OpRsh64x64_0(v)
	case OpRsh8Ux16:
		return rewriteValueGenM_OpRsh8Ux16_0(v)
	case OpRsh8Ux32:
		return rewriteValueGenM_OpRsh8Ux32_0(v)
	case OpRsh8Ux64:
		return rewriteValueGenM_OpRsh8Ux64_0(v)
	case OpRsh8Ux8:
		return rewriteValueGenM_OpRsh8Ux8_0(v)
	case OpRsh8x64:
		return rewriteValueGenM_OpRsh8x64_0(v)
	case OpSignExt16to32:
		return rewriteValueGenM_OpSignExt16to32_0(v)
	case OpSignExt16to64:
		return rewriteValueGenM_OpSignExt16to64_0(v)
	case OpSignExt32to64:
		return rewriteValueGenM_OpSignExt32to64_0(v)
	case OpSignExt8to32:
		return rewriteValueGenM_OpSignExt8to32_0(v)
	case OpSignExt8to64:
		return rewriteValueGenM_OpSignExt8to64_0(v)
	case OpSlicemask:
		return rewriteValueGenM_OpSlicemask_0(v)
	case OpStaticCall:
		return rewriteValueGenM_OpStaticCall_0(v)
	case OpStore:
		return rewriteValueGenM_OpStore_0(v)
	case OpSub16:
		return rewriteValueGenM_OpSub16_0(v)
	case OpSub32:
		return rewriteValueGenM_OpSub32_0(v)
	case OpSub32F:
		return rewriteValueGenM_OpSub32F_0(v)
	case OpSub64:
		return rewriteValueGenM_OpSub64_0(v)
	case OpSub64F:
		return rewriteValueGenM_OpSub64F_0(v)
	case OpSub8:
		return rewriteValueGenM_OpSub8_0(v)
	case OpTrunc16to8:
		return rewriteValueGenM_OpTrunc16to8_0(v)
	case OpTrunc32to16:
		return rewriteValueGenM_OpTrunc32to16_0(v)
	case OpTrunc32to8:
		return rewriteValueGenM_OpTrunc32to8_0(v)
	case OpTrunc64to16:
		return rewriteValueGenM_OpTrunc64to16_0(v)
	case OpTrunc64to32:
		return rewriteValueGenM_OpTrunc64to32_0(v)
	case OpTrunc64to8:
		return rewriteValueGenM_OpTrunc64to8_0(v)
	case OpWB:
		return rewriteValueGenM_OpWB_0(v)
	case OpXor16:
		return rewriteValueGenM_OpXor16_0(v)
	case OpXor32:
		return rewriteValueGenM_OpXor32_0(v)
	case OpXor64:
		return rewriteValueGenM_OpXor64_0(v)
	case OpXor8:
		return rewriteValueGenM_OpXor8_0(v)
	case OpZero:
		return rewriteValueGenM_OpZero_0(v)
	case OpZeroExt16to32:
		return rewriteValueGenM_OpZeroExt16to32_0(v)
	case OpZeroExt16to64:
		return rewriteValueGenM_OpZeroExt16to64_0(v)
	case OpZeroExt32to64:
		return rewriteValueGenM_OpZeroExt32to64_0(v)
	case OpZeroExt8to16:
		return rewriteValueGenM_OpZeroExt8to16_0(v)
	case OpZeroExt8to32:
		return rewriteValueGenM_OpZeroExt8to32_0(v)
	case OpZeroExt8to64:
		return rewriteValueGenM_OpZeroExt8to64_0(v)
	}
	return false
}
func rewriteValueGenM_OpAdd16_0(v *Value) bool {
	// match: (Add16 x y)
	// cond:
	// result: (ADD_I16 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMADD_I16)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpAdd32_0(v *Value) bool {
	// match: (Add32 x y)
	// cond:
	// result: (ADD_I32 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMADD_I32)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpAdd64_0(v *Value) bool {
	// match: (Add64 x y)
	// cond:
	// result: (ADD_I64 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMADD_I64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpAdd64F_0(v *Value) bool {
	// match: (Add64F x y)
	// cond:
	// result: (ADD_F64 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMADD_F64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpAdd8_0(v *Value) bool {
	// match: (Add8 x y)
	// cond:
	// result: (ADD_I8 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMADD_I8)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpAddPtr_0(v *Value) bool {
	// match: (AddPtr x y)
	// cond:
	// result: (ADD_I64 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMADD_I64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpAddr_0(v *Value) bool {
	// match: (Addr {sym} base)
	// cond:
	// result: (LoweredAddr {sym} base)
	for {
		sym := v.Aux
		base := v.Args[0]
		v.reset(OpGenMLoweredAddr)
		v.Aux = sym
		v.AddArg(base)
		return true
	}
}
func rewriteValueGenM_OpAnd16_0(v *Value) bool {
	// match: (And16 x y)
	// cond:
	// result: (AND_I16 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMAND_I16)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpAnd32_0(v *Value) bool {
	// match: (And32 x y)
	// cond:
	// result: (AND_I32 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMAND_I32)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpAnd64_0(v *Value) bool {
	// match: (And64 x y)
	// cond:
	// result: (AND_I64 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMAND_I64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpAnd8_0(v *Value) bool {
	// match: (And8 x y)
	// cond:
	// result: (AND_I8 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMAND_I8)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpAndB_0(v *Value) bool {
	// match: (AndB x y)
	// cond:
	// result: (AND_I1 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMAND_I1)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpClosureCall_0(v *Value) bool {
	// match: (ClosureCall [argwid] entry closure mem)
	// cond:
	// result: (LoweredClosureCall [argwid] entry closure mem)
	for {
		argwid := v.AuxInt
		_ = v.Args[2]
		entry := v.Args[0]
		closure := v.Args[1]
		mem := v.Args[2]
		v.reset(OpGenMLoweredClosureCall)
		v.AuxInt = argwid
		v.AddArg(entry)
		v.AddArg(closure)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueGenM_OpCom16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Com16 x)
	// cond:
	// result: (XOR_I16 x (CONST_I64 [-1]))
	for {
		x := v.Args[0]
		v.reset(OpGenMXOR_I16)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpGenMCONST_I64, typ.Int64)
		v0.AuxInt = -1
		v.AddArg(v0)
		return true
	}
}
func rewriteValueGenM_OpCom32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Com32 x)
	// cond:
	// result: (XOR_I32 x (CONST_I64 [-1]))
	for {
		x := v.Args[0]
		v.reset(OpGenMXOR_I32)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpGenMCONST_I64, typ.Int64)
		v0.AuxInt = -1
		v.AddArg(v0)
		return true
	}
}
func rewriteValueGenM_OpCom64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Com64 x)
	// cond:
	// result: (XOR_I64 x (CONST_I64 [-1]))
	for {
		x := v.Args[0]
		v.reset(OpGenMXOR_I64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpGenMCONST_I64, typ.Int64)
		v0.AuxInt = -1
		v.AddArg(v0)
		return true
	}
}
func rewriteValueGenM_OpCom8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Com8 x)
	// cond:
	// result: (XOR_I8 x (CONST_I64 [-1]))
	for {
		x := v.Args[0]
		v.reset(OpGenMXOR_I8)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpGenMCONST_I64, typ.Int64)
		v0.AuxInt = -1
		v.AddArg(v0)
		return true
	}
}
func rewriteValueGenM_OpConst16_0(v *Value) bool {
	// match: (Const16 [val])
	// cond:
	// result: (CONST_I16 [val])
	for {
		val := v.AuxInt
		v.reset(OpGenMCONST_I16)
		v.AuxInt = val
		return true
	}
}
func rewriteValueGenM_OpConst32_0(v *Value) bool {
	// match: (Const32 [val])
	// cond:
	// result: (CONST_I32 [val])
	for {
		val := v.AuxInt
		v.reset(OpGenMCONST_I32)
		v.AuxInt = val
		return true
	}
}
func rewriteValueGenM_OpConst32F_0(v *Value) bool {
	// match: (Const32F [val])
	// cond:
	// result: (CONST_F32 [val])
	for {
		val := v.AuxInt
		v.reset(OpGenMCONST_F32)
		v.AuxInt = val
		return true
	}
}
func rewriteValueGenM_OpConst64_0(v *Value) bool {
	// match: (Const64 [val])
	// cond:
	// result: (CONST_I64 [val])
	for {
		val := v.AuxInt
		v.reset(OpGenMCONST_I64)
		v.AuxInt = val
		return true
	}
}
func rewriteValueGenM_OpConst64F_0(v *Value) bool {
	// match: (Const64F [val])
	// cond:
	// result: (CONST_F64 [val])
	for {
		val := v.AuxInt
		v.reset(OpGenMCONST_F64)
		v.AuxInt = val
		return true
	}
}
func rewriteValueGenM_OpConst8_0(v *Value) bool {
	// match: (Const8 [val])
	// cond:
	// result: (CONST_I8 [val])
	for {
		val := v.AuxInt
		v.reset(OpGenMCONST_I8)
		v.AuxInt = val
		return true
	}
}
func rewriteValueGenM_OpConstBool_0(v *Value) bool {
	// match: (ConstBool [val])
	// cond:
	// result: (CONST_I1 [val])
	for {
		val := v.AuxInt
		v.reset(OpGenMCONST_I1)
		v.AuxInt = val
		return true
	}
}
func rewriteValueGenM_OpConstNil_0(v *Value) bool {
	// match: (ConstNil)
	// cond:
	// result: (CONST_I64 [0])
	for {
		v.reset(OpGenMCONST_I64)
		v.AuxInt = 0
		return true
	}
}
func rewriteValueGenM_OpCvt32Fto64_0(v *Value) bool {
	// match: (Cvt32Fto64 x)
	// cond:
	// result: (TRUNC_I64 x)
	for {
		x := v.Args[0]
		v.reset(OpGenMTRUNC_I64)
		v.AddArg(x)
		return true
	}
}
func rewriteValueGenM_OpCvt32Fto64F_0(v *Value) bool {
	// match: (Cvt32Fto64F x)
	// cond:
	// result: (FEXT_F64 x)
	for {
		x := v.Args[0]
		v.reset(OpGenMFEXT_F64)
		v.AddArg(x)
		return true
	}
}
func rewriteValueGenM_OpCvt32to64F_0(v *Value) bool {
	// match: (Cvt32to64F x)
	// cond:
	// result: (SEXT_F64 x)
	for {
		x := v.Args[0]
		v.reset(OpGenMSEXT_F64)
		v.AddArg(x)
		return true
	}
}
func rewriteValueGenM_OpCvt64Fto32_0(v *Value) bool {
	// match: (Cvt64Fto32 x)
	// cond:
	// result: (TRUNC_I32 x)
	for {
		x := v.Args[0]
		v.reset(OpGenMTRUNC_I32)
		v.AddArg(x)
		return true
	}
}
func rewriteValueGenM_OpCvt64Fto32F_0(v *Value) bool {
	// match: (Cvt64Fto32F x)
	// cond:
	// result: (TRUNC_F32 x)
	for {
		x := v.Args[0]
		v.reset(OpGenMTRUNC_F32)
		v.AddArg(x)
		return true
	}
}
func rewriteValueGenM_OpCvt64Fto64_0(v *Value) bool {
	// match: (Cvt64Fto64 x)
	// cond:
	// result: (TRUNC_I64 x)
	for {
		x := v.Args[0]
		v.reset(OpGenMTRUNC_I64)
		v.AddArg(x)
		return true
	}
}
func rewriteValueGenM_OpCvt64to64F_0(v *Value) bool {
	// match: (Cvt64to64F x)
	// cond:
	// result: (SEXT_F64 x)
	for {
		x := v.Args[0]
		v.reset(OpGenMSEXT_F64)
		v.AddArg(x)
		return true
	}
}
func rewriteValueGenM_OpDiv32_0(v *Value) bool {
	// match: (Div32 x y)
	// cond:
	// result: (DIV_I32 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMDIV_I32)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpDiv32u_0(v *Value) bool {
	// match: (Div32u x y)
	// cond:
	// result: (DIV_U32 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMDIV_U32)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpDiv64_0(v *Value) bool {
	// match: (Div64 x y)
	// cond:
	// result: (DIV_I64 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMDIV_I64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpDiv64F_0(v *Value) bool {
	// match: (Div64F x y)
	// cond:
	// result: (DIV_F64 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMDIV_F64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpDiv64u_0(v *Value) bool {
	// match: (Div64u x y)
	// cond:
	// result: (DIV_U64 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMDIV_U64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpDiv8u_0(v *Value) bool {
	// match: (Div8u x y)
	// cond:
	// result: (DIV_U8 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMDIV_U8)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpEq16_0(v *Value) bool {
	// match: (Eq16 x y)
	// cond:
	// result: (CMP_EQ x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_EQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpEq32_0(v *Value) bool {
	// match: (Eq32 x y)
	// cond:
	// result: (CMP_EQ x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_EQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpEq32F_0(v *Value) bool {
	// match: (Eq32F x y)
	// cond:
	// result: (CMP_EQ x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_EQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpEq64_0(v *Value) bool {
	// match: (Eq64 x y)
	// cond:
	// result: (CMP_EQ x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_EQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpEq64F_0(v *Value) bool {
	// match: (Eq64F x y)
	// cond:
	// result: (CMP_EQ x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_EQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpEq8_0(v *Value) bool {
	// match: (Eq8 x y)
	// cond:
	// result: (CMP_EQ x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_EQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpEqB_0(v *Value) bool {
	// match: (EqB x y)
	// cond:
	// result: (CMP_EQ x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_EQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpEqPtr_0(v *Value) bool {
	// match: (EqPtr x y)
	// cond:
	// result: (CMP_EQ x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_EQ)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpGeq16_0(v *Value) bool {
	// match: (Geq16 x y)
	// cond:
	// result: (CMP_GE x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_GE)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpGeq16U_0(v *Value) bool {
	// match: (Geq16U x y)
	// cond:
	// result: (CMP_UGE x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_UGE)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpGeq32_0(v *Value) bool {
	// match: (Geq32 x y)
	// cond:
	// result: (CMP_GE x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_GE)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpGeq32F_0(v *Value) bool {
	// match: (Geq32F x y)
	// cond:
	// result: (CMP_OGE x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_OGE)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpGeq32U_0(v *Value) bool {
	// match: (Geq32U x y)
	// cond:
	// result: (CMP_UGE x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_UGE)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpGeq64_0(v *Value) bool {
	// match: (Geq64 x y)
	// cond:
	// result: (CMP_GE x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_GE)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpGeq64F_0(v *Value) bool {
	// match: (Geq64F x y)
	// cond:
	// result: (CMP_OGE x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_OGE)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpGeq64U_0(v *Value) bool {
	// match: (Geq64U x y)
	// cond:
	// result: (CMP_UGE x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_UGE)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpGeq8_0(v *Value) bool {
	// match: (Geq8 x y)
	// cond:
	// result: (CMP_GE x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_GE)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpGeq8U_0(v *Value) bool {
	// match: (Geq8U x y)
	// cond:
	// result: (CMP_UGE x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_UGE)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpGetCallerPC_0(v *Value) bool {
	// match: (GetCallerPC)
	// cond:
	// result: (LoweredGetCallerPC)
	for {
		v.reset(OpGenMLoweredGetCallerPC)
		return true
	}
}
func rewriteValueGenM_OpGetCallerSP_0(v *Value) bool {
	// match: (GetCallerSP)
	// cond:
	// result: (LoweredGetCallerSP)
	for {
		v.reset(OpGenMLoweredGetCallerSP)
		return true
	}
}
func rewriteValueGenM_OpGetClosurePtr_0(v *Value) bool {
	// match: (GetClosurePtr)
	// cond:
	// result: (LoweredGetClosurePtr)
	for {
		v.reset(OpGenMLoweredGetClosurePtr)
		return true
	}
}
func rewriteValueGenM_OpGetG_0(v *Value) bool {
	// match: (GetG mem)
	// cond:
	// result: (LoweredGetG mem)
	for {
		mem := v.Args[0]
		v.reset(OpGenMLoweredGetG)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueGenM_OpGreater16_0(v *Value) bool {
	// match: (Greater16 x y)
	// cond:
	// result: (CMP_GT x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_GT)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpGreater16U_0(v *Value) bool {
	// match: (Greater16U x y)
	// cond:
	// result: (CMP_UGT x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_UGT)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpGreater32_0(v *Value) bool {
	// match: (Greater32 x y)
	// cond:
	// result: (CMP_GT x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_GT)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpGreater32F_0(v *Value) bool {
	// match: (Greater32F x y)
	// cond:
	// result: (CMP_OGT x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_OGT)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpGreater32U_0(v *Value) bool {
	// match: (Greater32U x y)
	// cond:
	// result: (CMP_UGT x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_UGT)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpGreater64_0(v *Value) bool {
	// match: (Greater64 x y)
	// cond:
	// result: (CMP_GT x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_GT)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpGreater64F_0(v *Value) bool {
	// match: (Greater64F x y)
	// cond:
	// result: (CMP_OGT x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_OGT)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpGreater64U_0(v *Value) bool {
	// match: (Greater64U x y)
	// cond:
	// result: (CMP_UGT x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_UGT)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpGreater8_0(v *Value) bool {
	// match: (Greater8 x y)
	// cond:
	// result: (CMP_GT x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_GT)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpGreater8U_0(v *Value) bool {
	// match: (Greater8U x y)
	// cond:
	// result: (CMP_UGT x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_UGT)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpInterCall_0(v *Value) bool {
	// match: (InterCall [argwid] entry mem)
	// cond:
	// result: (LoweredInterCall [argwid] entry mem)
	for {
		argwid := v.AuxInt
		_ = v.Args[1]
		entry := v.Args[0]
		mem := v.Args[1]
		v.reset(OpGenMLoweredInterCall)
		v.AuxInt = argwid
		v.AddArg(entry)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueGenM_OpIsInBounds_0(v *Value) bool {
	// match: (IsInBounds idx len)
	// cond:
	// result: (CMP_LT idx len)
	for {
		_ = v.Args[1]
		idx := v.Args[0]
		len := v.Args[1]
		v.reset(OpGenMCMP_LT)
		v.AddArg(idx)
		v.AddArg(len)
		return true
	}
}
func rewriteValueGenM_OpIsNonNil_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (IsNonNil p)
	// cond:
	// result: (CMP_NE p (CONST_I64 [0]))
	for {
		p := v.Args[0]
		v.reset(OpGenMCMP_NE)
		v.AddArg(p)
		v0 := b.NewValue0(v.Pos, OpGenMCONST_I64, typ.Int64)
		v0.AuxInt = 0
		v.AddArg(v0)
		return true
	}
}
func rewriteValueGenM_OpIsSliceInBounds_0(v *Value) bool {
	// match: (IsSliceInBounds idx len)
	// cond:
	// result: (CMP_ULT idx len)
	for {
		_ = v.Args[1]
		idx := v.Args[0]
		len := v.Args[1]
		v.reset(OpGenMCMP_ULT)
		v.AddArg(idx)
		v.AddArg(len)
		return true
	}
}
func rewriteValueGenM_OpLeq16_0(v *Value) bool {
	// match: (Leq16 x y)
	// cond:
	// result: (CMP_LE x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_LE)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpLeq16U_0(v *Value) bool {
	// match: (Leq16U x y)
	// cond:
	// result: (CMP_ULE x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_ULE)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpLeq32_0(v *Value) bool {
	// match: (Leq32 x y)
	// cond:
	// result: (CMP_LE x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_LE)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpLeq32F_0(v *Value) bool {
	// match: (Leq32F x y)
	// cond:
	// result: (CMP_OLE x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_OLE)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpLeq32U_0(v *Value) bool {
	// match: (Leq32U x y)
	// cond:
	// result: (CMP_ULE x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_ULE)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpLeq64_0(v *Value) bool {
	// match: (Leq64 x y)
	// cond:
	// result: (CMP_LE x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_LE)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpLeq64F_0(v *Value) bool {
	// match: (Leq64F x y)
	// cond:
	// result: (CMP_OLE x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_OLE)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpLeq64U_0(v *Value) bool {
	// match: (Leq64U x y)
	// cond:
	// result: (CMP_ULE x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_ULE)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpLeq8_0(v *Value) bool {
	// match: (Leq8 x y)
	// cond:
	// result: (CMP_LE x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_LE)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpLeq8U_0(v *Value) bool {
	// match: (Leq8U x y)
	// cond:
	// result: (CMP_ULE x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_ULE)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpLess16_0(v *Value) bool {
	// match: (Less16 x y)
	// cond:
	// result: (CMP_LT x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_LT)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpLess16U_0(v *Value) bool {
	// match: (Less16U x y)
	// cond:
	// result: (CMP_ULT x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_ULT)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpLess32_0(v *Value) bool {
	// match: (Less32 x y)
	// cond:
	// result: (CMP_LT x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_LT)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpLess32F_0(v *Value) bool {
	// match: (Less32F x y)
	// cond:
	// result: (CMP_OLT x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_OLT)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpLess32U_0(v *Value) bool {
	// match: (Less32U x y)
	// cond:
	// result: (CMP_ULT x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_ULT)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpLess64_0(v *Value) bool {
	// match: (Less64 x y)
	// cond:
	// result: (CMP_LT x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_LT)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpLess64F_0(v *Value) bool {
	// match: (Less64F x y)
	// cond:
	// result: (CMP_OLT x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_OLT)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpLess64U_0(v *Value) bool {
	// match: (Less64U x y)
	// cond:
	// result: (CMP_ULT x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_ULT)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpLess8_0(v *Value) bool {
	// match: (Less8 x y)
	// cond:
	// result: (CMP_LT x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_LT)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpLess8U_0(v *Value) bool {
	// match: (Less8U x y)
	// cond:
	// result: (CMP_ULT x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_ULT)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpLoad_0(v *Value) bool {
	// match: (Load <t> ptr mem)
	// cond: is32BitFloat(t)
	// result: (LD_4_F32 ptr mem)
	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(is32BitFloat(t)) {
			break
		}
		v.reset(OpGenMLD_4_F32)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is64BitFloat(t)
	// result: (LD_8_F64 ptr mem)
	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(is64BitFloat(t)) {
			break
		}
		v.reset(OpGenMLD_8_F64)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.Size() == 8 && !t.IsSigned()
	// result: (LD_8_U64 ptr mem)
	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.Size() == 8 && !t.IsSigned()) {
			break
		}
		v.reset(OpGenMLD_8_U64)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.Size() == 8 && t.IsSigned()
	// result: (LD_8_I64 ptr mem)
	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.Size() == 8 && t.IsSigned()) {
			break
		}
		v.reset(OpGenMLD_8_I64)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.Size() == 4 && !t.IsSigned()
	// result: (LD_4_U32 ptr mem)
	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.Size() == 4 && !t.IsSigned()) {
			break
		}
		v.reset(OpGenMLD_4_U32)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.Size() == 4 && t.IsSigned()
	// result: (LD_4_I32 ptr mem)
	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.Size() == 4 && t.IsSigned()) {
			break
		}
		v.reset(OpGenMLD_4_I32)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.Size() == 2 && !t.IsSigned()
	// result: (LD_2_U16 ptr mem)
	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.Size() == 2 && !t.IsSigned()) {
			break
		}
		v.reset(OpGenMLD_2_U16)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.Size() == 2 && t.IsSigned()
	// result: (LD_2_I16 ptr mem)
	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.Size() == 2 && t.IsSigned()) {
			break
		}
		v.reset(OpGenMLD_2_I16)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.Size() == 1 && !t.IsSigned()
	// result: (LD_1_U8 ptr mem)
	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.Size() == 1 && !t.IsSigned()) {
			break
		}
		v.reset(OpGenMLD_1_U8)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.Size() == 1 && t.IsSigned()
	// result: (LD_1_I8 ptr mem)
	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.Size() == 1 && t.IsSigned()) {
			break
		}
		v.reset(OpGenMLD_1_I8)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueGenM_OpLocalAddr_0(v *Value) bool {
	// match: (LocalAddr {sym} base _)
	// cond:
	// result: (LoweredAddr {sym} base)
	for {
		sym := v.Aux
		_ = v.Args[1]
		base := v.Args[0]
		v.reset(OpGenMLoweredAddr)
		v.Aux = sym
		v.AddArg(base)
		return true
	}
}
func rewriteValueGenM_OpLsh16x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Lsh16x16 x y)
	// cond:
	// result: (SELECT_I16 (SLL_I16 x y) (CONST_I16 [0]) (CMP_ULT y (CONST_I16 [32])))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSELECT_I16)
		v0 := b.NewValue0(v.Pos, OpGenMSLL_I16, typ.Int16)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpGenMCONST_I16, typ.Int16)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpGenMCMP_ULT, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpGenMCONST_I16, typ.Int16)
		v3.AuxInt = 32
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueGenM_OpLsh16x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Lsh16x32 x y)
	// cond:
	// result: (SELECT_I16 (SLL_I16 x y) (CONST_I16 [0]) (CMP_ULT y (CONST_I32 [32])))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSELECT_I16)
		v0 := b.NewValue0(v.Pos, OpGenMSLL_I16, typ.Int16)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpGenMCONST_I16, typ.Int16)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpGenMCMP_ULT, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpGenMCONST_I32, typ.Int32)
		v3.AuxInt = 32
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueGenM_OpLsh16x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Lsh16x64 x y)
	// cond:
	// result: (SELECT_I16 (SLL_I16 x y) (CONST_I16 [0]) (CMP_ULT y (CONST_I64 [64])))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSELECT_I16)
		v0 := b.NewValue0(v.Pos, OpGenMSLL_I16, typ.Int16)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpGenMCONST_I16, typ.Int16)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpGenMCMP_ULT, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpGenMCONST_I64, typ.Int64)
		v3.AuxInt = 64
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueGenM_OpLsh16x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Lsh16x8 x y)
	// cond:
	// result: (SELECT_I16 (SLL_I16 x y) (CONST_I16 [0]) (CMP_ULT y (CONST_I8 [32])))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSELECT_I16)
		v0 := b.NewValue0(v.Pos, OpGenMSLL_I16, typ.Int16)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpGenMCONST_I16, typ.Int16)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpGenMCMP_ULT, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpGenMCONST_I8, typ.Int8)
		v3.AuxInt = 32
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueGenM_OpLsh32x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Lsh32x16 x y)
	// cond:
	// result: (SELECT_I32 (SLL_I32 x y) (CONST_I32 [0]) (CMP_ULT y (CONST_I16 [32])))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSELECT_I32)
		v0 := b.NewValue0(v.Pos, OpGenMSLL_I32, typ.Int32)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpGenMCONST_I32, typ.Int32)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpGenMCMP_ULT, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpGenMCONST_I16, typ.Int16)
		v3.AuxInt = 32
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueGenM_OpLsh32x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Lsh32x32 x y)
	// cond:
	// result: (SELECT_I32 (SLL_I32 x y) (CONST_I32 [0]) (CMP_ULT y (CONST_I32 [32])))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSELECT_I32)
		v0 := b.NewValue0(v.Pos, OpGenMSLL_I32, typ.Int32)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpGenMCONST_I32, typ.Int32)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpGenMCMP_ULT, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpGenMCONST_I32, typ.Int32)
		v3.AuxInt = 32
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueGenM_OpLsh32x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Lsh32x64 x y)
	// cond:
	// result: (SELECT_I32 (SLL_I32 x y) (CONST_I32 [0]) (CMP_ULT y (CONST_I64 [64])))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSELECT_I32)
		v0 := b.NewValue0(v.Pos, OpGenMSLL_I32, typ.Int32)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpGenMCONST_I32, typ.Int32)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpGenMCMP_ULT, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpGenMCONST_I64, typ.Int64)
		v3.AuxInt = 64
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueGenM_OpLsh32x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Lsh32x8 x y)
	// cond:
	// result: (SELECT_I32 (SLL_I32 x y) (CONST_I32 [0]) (CMP_ULT y (CONST_I8 [32])))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSELECT_I32)
		v0 := b.NewValue0(v.Pos, OpGenMSLL_I32, typ.Int32)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpGenMCONST_I32, typ.Int32)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpGenMCMP_ULT, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpGenMCONST_I8, typ.Int8)
		v3.AuxInt = 32
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueGenM_OpLsh64x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Lsh64x16 x y)
	// cond:
	// result: (SELECT_I64 (SLL_I64 x y) (CONST_I64 [0]) (CMP_ULT y (CONST_I16 [32])))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSELECT_I64)
		v0 := b.NewValue0(v.Pos, OpGenMSLL_I64, typ.Int64)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpGenMCONST_I64, typ.Int64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpGenMCMP_ULT, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpGenMCONST_I16, typ.Int16)
		v3.AuxInt = 32
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueGenM_OpLsh64x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Lsh64x32 x y)
	// cond:
	// result: (SELECT_I64 (SLL_I64 x y) (CONST_I64 [0]) (CMP_ULT y (CONST_I32 [32])))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSELECT_I64)
		v0 := b.NewValue0(v.Pos, OpGenMSLL_I64, typ.Int64)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpGenMCONST_I64, typ.Int64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpGenMCMP_ULT, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpGenMCONST_I32, typ.Int32)
		v3.AuxInt = 32
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueGenM_OpLsh64x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Lsh64x64 x y)
	// cond:
	// result: (SELECT_I64 (SLL_I64 x y) (CONST_I64 [0]) (CMP_ULT y (CONST_I64 [64])))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSELECT_I64)
		v0 := b.NewValue0(v.Pos, OpGenMSLL_I64, typ.Int64)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpGenMCONST_I64, typ.Int64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpGenMCMP_ULT, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpGenMCONST_I64, typ.Int64)
		v3.AuxInt = 64
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueGenM_OpLsh64x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Lsh64x8 x y)
	// cond:
	// result: (SELECT_I64 (SLL_I64 x y) (CONST_I64 [0]) (CMP_ULT y (CONST_I8 [32])))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSELECT_I64)
		v0 := b.NewValue0(v.Pos, OpGenMSLL_I64, typ.Int64)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpGenMCONST_I64, typ.Int64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpGenMCMP_ULT, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpGenMCONST_I8, typ.Int8)
		v3.AuxInt = 32
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueGenM_OpLsh8x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Lsh8x16 x y)
	// cond:
	// result: (SELECT_I8 (SLL_I8 x y) (CONST_I8 [0]) (CMP_ULT y (CONST_I16 [32])))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSELECT_I8)
		v0 := b.NewValue0(v.Pos, OpGenMSLL_I8, typ.Int8)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpGenMCONST_I8, typ.Int8)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpGenMCMP_ULT, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpGenMCONST_I16, typ.Int16)
		v3.AuxInt = 32
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueGenM_OpLsh8x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Lsh8x32 x y)
	// cond:
	// result: (SELECT_I8 (SLL_I8 x y) (CONST_I8 [0]) (CMP_ULT y (CONST_I32 [32])))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSELECT_I8)
		v0 := b.NewValue0(v.Pos, OpGenMSLL_I8, typ.Int8)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpGenMCONST_I8, typ.Int8)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpGenMCMP_ULT, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpGenMCONST_I32, typ.Int32)
		v3.AuxInt = 32
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueGenM_OpLsh8x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Lsh8x64 x y)
	// cond:
	// result: (SELECT_I8 (SLL_I8 x y) (CONST_I8 [0]) (CMP_ULT y (CONST_I64 [64])))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSELECT_I8)
		v0 := b.NewValue0(v.Pos, OpGenMSLL_I8, typ.Int8)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpGenMCONST_I8, typ.Int8)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpGenMCMP_ULT, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpGenMCONST_I64, typ.Int64)
		v3.AuxInt = 64
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueGenM_OpLsh8x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Lsh8x8 x y)
	// cond:
	// result: (SELECT_I8 (SLL_I8 x y) (CONST_I8 [0]) (CMP_ULT y (CONST_I8 [32])))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSELECT_I8)
		v0 := b.NewValue0(v.Pos, OpGenMSLL_I8, typ.Int8)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpGenMCONST_I8, typ.Int8)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpGenMCMP_ULT, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpGenMCONST_I8, typ.Int8)
		v3.AuxInt = 32
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueGenM_OpMod32_0(v *Value) bool {
	// match: (Mod32 x y)
	// cond:
	// result: (MOD_I32 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMMOD_I32)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpMod32u_0(v *Value) bool {
	// match: (Mod32u x y)
	// cond:
	// result: (MOD_U32 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMMOD_U32)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpMod64_0(v *Value) bool {
	// match: (Mod64 x y)
	// cond:
	// result: (MOD_I64 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMMOD_I64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpMod64u_0(v *Value) bool {
	// match: (Mod64u x y)
	// cond:
	// result: (MOD_U64 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMMOD_U64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpMove_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Move [8] dst src mem)
	// cond:
	// result: (ST_8 dst (LD_8_I64 src mem) mem)
	for {
		if v.AuxInt != 8 {
			break
		}
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpGenMST_8)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpGenMLD_8_I64, typ.Int64)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [s] dst src mem)
	// cond:
	// result: (LoweredMove [s] dst src mem)
	for {
		s := v.AuxInt
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpGenMLoweredMove)
		v.AuxInt = s
		v.AddArg(dst)
		v.AddArg(src)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueGenM_OpMul16_0(v *Value) bool {
	// match: (Mul16 x y)
	// cond:
	// result: (MUL_I16 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMMUL_I16)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpMul32_0(v *Value) bool {
	// match: (Mul32 x y)
	// cond:
	// result: (MUL_I32 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMMUL_I32)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpMul64_0(v *Value) bool {
	// match: (Mul64 x y)
	// cond:
	// result: (MUL_I64 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMMUL_I64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpMul64F_0(v *Value) bool {
	// match: (Mul64F x y)
	// cond:
	// result: (MUL_F64 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMMUL_F64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpMul8_0(v *Value) bool {
	// match: (Mul8 x y)
	// cond:
	// result: (MUL_I8 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMMUL_I8)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpNeg32_0(v *Value) bool {
	// match: (Neg32 x)
	// cond:
	// result: (NEG_I32 x)
	for {
		x := v.Args[0]
		v.reset(OpGenMNEG_I32)
		v.AddArg(x)
		return true
	}
}
func rewriteValueGenM_OpNeg32F_0(v *Value) bool {
	// match: (Neg32F x)
	// cond:
	// result: (NEG_F32 x)
	for {
		x := v.Args[0]
		v.reset(OpGenMNEG_F32)
		v.AddArg(x)
		return true
	}
}
func rewriteValueGenM_OpNeg64_0(v *Value) bool {
	// match: (Neg64 x)
	// cond:
	// result: (NEG_I64 x)
	for {
		x := v.Args[0]
		v.reset(OpGenMNEG_I64)
		v.AddArg(x)
		return true
	}
}
func rewriteValueGenM_OpNeg64F_0(v *Value) bool {
	// match: (Neg64F x)
	// cond:
	// result: (NEG_F64 x)
	for {
		x := v.Args[0]
		v.reset(OpGenMNEG_F64)
		v.AddArg(x)
		return true
	}
}
func rewriteValueGenM_OpNeq16_0(v *Value) bool {
	// match: (Neq16 x y)
	// cond:
	// result: (CMP_NE x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_NE)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpNeq32_0(v *Value) bool {
	// match: (Neq32 x y)
	// cond:
	// result: (CMP_NE x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_NE)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpNeq32F_0(v *Value) bool {
	// match: (Neq32F x y)
	// cond:
	// result: (CMP_ONE x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_ONE)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpNeq64_0(v *Value) bool {
	// match: (Neq64 x y)
	// cond:
	// result: (CMP_NE x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_NE)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpNeq64F_0(v *Value) bool {
	// match: (Neq64F x y)
	// cond:
	// result: (CMP_ONE x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_ONE)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpNeq8_0(v *Value) bool {
	// match: (Neq8 x y)
	// cond:
	// result: (CMP_NE x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_NE)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpNeqB_0(v *Value) bool {
	// match: (NeqB x y)
	// cond:
	// result: (CMP_NE x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_NE)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpNeqPtr_0(v *Value) bool {
	// match: (NeqPtr x y)
	// cond:
	// result: (CMP_NE x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_NE)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpNilCheck_0(v *Value) bool {
	// match: (NilCheck ptr mem)
	// cond:
	// result: (LoweredNilCheck ptr mem)
	for {
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpGenMLoweredNilCheck)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueGenM_OpNot_0(v *Value) bool {
	// match: (Not x)
	// cond:
	// result: (NOT_I1 x)
	for {
		x := v.Args[0]
		v.reset(OpGenMNOT_I1)
		v.AddArg(x)
		return true
	}
}
func rewriteValueGenM_OpOffPtr_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (OffPtr [off] ptr)
	// cond:
	// result: (ADD_I64 (CONST_I64 [off]) ptr)
	for {
		off := v.AuxInt
		ptr := v.Args[0]
		v.reset(OpGenMADD_I64)
		v0 := b.NewValue0(v.Pos, OpGenMCONST_I64, typ.Int64)
		v0.AuxInt = off
		v.AddArg(v0)
		v.AddArg(ptr)
		return true
	}
}
func rewriteValueGenM_OpOr16_0(v *Value) bool {
	// match: (Or16 x y)
	// cond:
	// result: (OR_I16 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMOR_I16)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpOr32_0(v *Value) bool {
	// match: (Or32 x y)
	// cond:
	// result: (OR_I32 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMOR_I32)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpOr64_0(v *Value) bool {
	// match: (Or64 x y)
	// cond:
	// result: (OR_I64 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMOR_I64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpOr8_0(v *Value) bool {
	// match: (Or8 x y)
	// cond:
	// result: (OR_I8 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMOR_I8)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpOrB_0(v *Value) bool {
	// match: (OrB x y)
	// cond:
	// result: (OR_I1 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMOR_I1)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpRsh16x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh16x64 x y)
	// cond:
	// result: (SAR_I16 x (SELECT_I64 y (CONST_I64 [15]) (CMP_ULT y (CONST_I64 [16]))))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSAR_I16)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpGenMSELECT_I64, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpGenMCONST_I64, typ.Int64)
		v1.AuxInt = 15
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpGenMCMP_ULT, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpGenMCONST_I64, typ.Int64)
		v3.AuxInt = 16
		v2.AddArg(v3)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueGenM_OpRsh32Ux16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh32Ux16 x y)
	// cond:
	// result: (SELECT_I32 (SLR_I32 x y) (CONST_I32 [0]) (CMP_ULT y (CONST_I16 [16])))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSELECT_I32)
		v0 := b.NewValue0(v.Pos, OpGenMSLR_I32, typ.Int32)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpGenMCONST_I32, typ.Int32)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpGenMCMP_ULT, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpGenMCONST_I16, typ.Int16)
		v3.AuxInt = 16
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueGenM_OpRsh32Ux32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh32Ux32 x y)
	// cond:
	// result: (SELECT_I32 (SLR_I32 x y) (CONST_I32 [0]) (CMP_ULT y (CONST_I32 [32])))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSELECT_I32)
		v0 := b.NewValue0(v.Pos, OpGenMSLR_I32, typ.Int32)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpGenMCONST_I32, typ.Int32)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpGenMCMP_ULT, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpGenMCONST_I32, typ.Int32)
		v3.AuxInt = 32
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueGenM_OpRsh32Ux64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh32Ux64 x y)
	// cond:
	// result: (SELECT_I32 (SLR_I32 x y) (CONST_I32 [0]) (CMP_ULT y (CONST_I64 [64])))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSELECT_I32)
		v0 := b.NewValue0(v.Pos, OpGenMSLR_I32, typ.Int32)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpGenMCONST_I32, typ.Int32)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpGenMCMP_ULT, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpGenMCONST_I64, typ.Int64)
		v3.AuxInt = 64
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueGenM_OpRsh32Ux8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh32Ux8 x y)
	// cond:
	// result: (SELECT_I32 (SLR_I32 x y) (CONST_I32 [0]) (CMP_ULT y (CONST_I8 [8])))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSELECT_I32)
		v0 := b.NewValue0(v.Pos, OpGenMSLR_I32, typ.Int32)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpGenMCONST_I32, typ.Int32)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpGenMCMP_ULT, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpGenMCONST_I8, typ.Int8)
		v3.AuxInt = 8
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueGenM_OpRsh32x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh32x64 x y)
	// cond:
	// result: (SAR_I32 x (SELECT_I64 y (CONST_I64 [31]) (CMP_ULT y (CONST_I64 [32]))))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSAR_I32)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpGenMSELECT_I64, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpGenMCONST_I64, typ.Int64)
		v1.AuxInt = 31
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpGenMCMP_ULT, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpGenMCONST_I64, typ.Int64)
		v3.AuxInt = 32
		v2.AddArg(v3)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueGenM_OpRsh64Ux16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh64Ux16 x y)
	// cond:
	// result: (SELECT_I64 (SLR_I64 x y) (CONST_I64 [0]) (CMP_ULT y (CONST_I16 [16])))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSELECT_I64)
		v0 := b.NewValue0(v.Pos, OpGenMSLR_I64, typ.Int64)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpGenMCONST_I64, typ.Int64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpGenMCMP_ULT, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpGenMCONST_I16, typ.Int16)
		v3.AuxInt = 16
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueGenM_OpRsh64Ux32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh64Ux32 x y)
	// cond:
	// result: (SELECT_I64 (SLR_I64 x y) (CONST_I64 [0]) (CMP_ULT y (CONST_I32 [32])))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSELECT_I64)
		v0 := b.NewValue0(v.Pos, OpGenMSLR_I64, typ.Int64)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpGenMCONST_I64, typ.Int64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpGenMCMP_ULT, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpGenMCONST_I32, typ.Int32)
		v3.AuxInt = 32
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueGenM_OpRsh64Ux64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh64Ux64 x y)
	// cond:
	// result: (SELECT_I64 (SLR_I64 x y) (CONST_I64 [0]) (CMP_ULT y (CONST_I64 [64])))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSELECT_I64)
		v0 := b.NewValue0(v.Pos, OpGenMSLR_I64, typ.Int64)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpGenMCONST_I64, typ.Int64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpGenMCMP_ULT, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpGenMCONST_I64, typ.Int64)
		v3.AuxInt = 64
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueGenM_OpRsh64Ux8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh64Ux8 x y)
	// cond:
	// result: (SELECT_I64 (SLR_I64 x y) (CONST_I64 [0]) (CMP_ULT y (CONST_I8 [8])))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSELECT_I64)
		v0 := b.NewValue0(v.Pos, OpGenMSLR_I64, typ.Int64)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpGenMCONST_I64, typ.Int64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpGenMCMP_ULT, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpGenMCONST_I8, typ.Int8)
		v3.AuxInt = 8
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueGenM_OpRsh64x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh64x64 x y)
	// cond:
	// result: (SAR_I64 x (SELECT_I64 y (CONST_I64 [63]) (CMP_ULT y (CONST_I64 [64]))))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSAR_I64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpGenMSELECT_I64, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpGenMCONST_I64, typ.Int64)
		v1.AuxInt = 63
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpGenMCMP_ULT, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpGenMCONST_I64, typ.Int64)
		v3.AuxInt = 64
		v2.AddArg(v3)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueGenM_OpRsh8Ux16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh8Ux16 x y)
	// cond:
	// result: (SELECT_I8 (SLR_I8 x y) (CONST_I8 [0]) (CMP_ULT y (CONST_I16 [16])))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSELECT_I8)
		v0 := b.NewValue0(v.Pos, OpGenMSLR_I8, typ.Int8)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpGenMCONST_I8, typ.Int8)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpGenMCMP_ULT, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpGenMCONST_I16, typ.Int16)
		v3.AuxInt = 16
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueGenM_OpRsh8Ux32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh8Ux32 x y)
	// cond:
	// result: (SELECT_I8 (SLR_I8 x y) (CONST_I8 [0]) (CMP_ULT y (CONST_I32 [32])))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSELECT_I8)
		v0 := b.NewValue0(v.Pos, OpGenMSLR_I8, typ.Int8)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpGenMCONST_I8, typ.Int8)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpGenMCMP_ULT, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpGenMCONST_I32, typ.Int32)
		v3.AuxInt = 32
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueGenM_OpRsh8Ux64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh8Ux64 x y)
	// cond:
	// result: (SELECT_I8 (SLR_I8 x y) (CONST_I8 [0]) (CMP_ULT y (CONST_I64 [64])))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSELECT_I8)
		v0 := b.NewValue0(v.Pos, OpGenMSLR_I8, typ.Int8)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpGenMCONST_I8, typ.Int8)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpGenMCMP_ULT, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpGenMCONST_I64, typ.Int64)
		v3.AuxInt = 64
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueGenM_OpRsh8Ux8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh8Ux8 x y)
	// cond:
	// result: (SELECT_I8 (SLR_I8 x y) (CONST_I8 [0]) (CMP_ULT y (CONST_I8 [8])))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSELECT_I8)
		v0 := b.NewValue0(v.Pos, OpGenMSLR_I8, typ.Int8)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpGenMCONST_I8, typ.Int8)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpGenMCMP_ULT, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpGenMCONST_I8, typ.Int8)
		v3.AuxInt = 8
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueGenM_OpRsh8x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Rsh8x64 x y)
	// cond:
	// result: (SAR_I8 x (SELECT_I64 y (CONST_I64 [7]) (CMP_ULT y (CONST_I64 [8]))))
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSAR_I8)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpGenMSELECT_I64, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpGenMCONST_I64, typ.Int64)
		v1.AuxInt = 7
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpGenMCMP_ULT, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpGenMCONST_I64, typ.Int64)
		v3.AuxInt = 8
		v2.AddArg(v3)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueGenM_OpSignExt16to32_0(v *Value) bool {
	// match: (SignExt16to32 x)
	// cond:
	// result: (SEXT_I32 x)
	for {
		x := v.Args[0]
		v.reset(OpGenMSEXT_I32)
		v.AddArg(x)
		return true
	}
}
func rewriteValueGenM_OpSignExt16to64_0(v *Value) bool {
	// match: (SignExt16to64 x)
	// cond:
	// result: (SEXT_I64 x)
	for {
		x := v.Args[0]
		v.reset(OpGenMSEXT_I64)
		v.AddArg(x)
		return true
	}
}
func rewriteValueGenM_OpSignExt32to64_0(v *Value) bool {
	// match: (SignExt32to64 x)
	// cond:
	// result: (SEXT_I64 x)
	for {
		x := v.Args[0]
		v.reset(OpGenMSEXT_I64)
		v.AddArg(x)
		return true
	}
}
func rewriteValueGenM_OpSignExt8to32_0(v *Value) bool {
	// match: (SignExt8to32 x)
	// cond:
	// result: (SEXT_I32 x)
	for {
		x := v.Args[0]
		v.reset(OpGenMSEXT_I32)
		v.AddArg(x)
		return true
	}
}
func rewriteValueGenM_OpSignExt8to64_0(v *Value) bool {
	// match: (SignExt8to64 x)
	// cond:
	// result: (SEXT_I64 x)
	for {
		x := v.Args[0]
		v.reset(OpGenMSEXT_I64)
		v.AddArg(x)
		return true
	}
}
func rewriteValueGenM_OpSlicemask_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Slicemask <t> x)
	// cond: t.Size() == 8
	// result: (SAR_I64 (NEG_I64 x) (CONST_I64 [63]))
	for {
		t := v.Type
		x := v.Args[0]
		if !(t.Size() == 8) {
			break
		}
		v.reset(OpGenMSAR_I64)
		v0 := b.NewValue0(v.Pos, OpGenMNEG_I64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpGenMCONST_I64, typ.Int64)
		v1.AuxInt = 63
		v.AddArg(v1)
		return true
	}
	return false
}
func rewriteValueGenM_OpStaticCall_0(v *Value) bool {
	// match: (StaticCall [argwid] {target} mem)
	// cond:
	// result: (LoweredStaticCall [argwid] {target} mem)
	for {
		argwid := v.AuxInt
		target := v.Aux
		mem := v.Args[0]
		v.reset(OpGenMLoweredStaticCall)
		v.AuxInt = argwid
		v.Aux = target
		v.AddArg(mem)
		return true
	}
}
func rewriteValueGenM_OpStore_0(v *Value) bool {
	// match: (Store {t} ptr val mem)
	// cond: is64BitFloat(t.(*types.Type))
	// result: (ST_8 ptr val mem)
	for {
		t := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is64BitFloat(t.(*types.Type))) {
			break
		}
		v.reset(OpGenMST_8)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: is32BitFloat(t.(*types.Type))
	// result: (ST_4 ptr val mem)
	for {
		t := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32BitFloat(t.(*types.Type))) {
			break
		}
		v.reset(OpGenMST_4)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 8
	// result: (ST_8 ptr val mem)
	for {
		t := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Size() == 8) {
			break
		}
		v.reset(OpGenMST_8)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 4
	// result: (ST_4 ptr val mem)
	for {
		t := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Size() == 4) {
			break
		}
		v.reset(OpGenMST_4)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 2
	// result: (ST_2 ptr val mem)
	for {
		t := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Size() == 2) {
			break
		}
		v.reset(OpGenMST_2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 1
	// result: (ST_1 ptr val mem)
	for {
		t := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Size() == 1) {
			break
		}
		v.reset(OpGenMST_1)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueGenM_OpSub16_0(v *Value) bool {
	// match: (Sub16 x y)
	// cond:
	// result: (SUB_I16 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSUB_I16)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpSub32_0(v *Value) bool {
	// match: (Sub32 x y)
	// cond:
	// result: (SUB_I32 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSUB_I32)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpSub32F_0(v *Value) bool {
	// match: (Sub32F x y)
	// cond:
	// result: (SUB_F32 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSUB_F32)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpSub64_0(v *Value) bool {
	// match: (Sub64 x y)
	// cond:
	// result: (SUB_I64 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSUB_I64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpSub64F_0(v *Value) bool {
	// match: (Sub64F x y)
	// cond:
	// result: (SUB_F64 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSUB_F64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpSub8_0(v *Value) bool {
	// match: (Sub8 x y)
	// cond:
	// result: (SUB_I8 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMSUB_I8)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpTrunc16to8_0(v *Value) bool {
	// match: (Trunc16to8 x)
	// cond:
	// result: (TRUNC_I8 x)
	for {
		x := v.Args[0]
		v.reset(OpGenMTRUNC_I8)
		v.AddArg(x)
		return true
	}
}
func rewriteValueGenM_OpTrunc32to16_0(v *Value) bool {
	// match: (Trunc32to16 x)
	// cond:
	// result: (TRUNC_I16 x)
	for {
		x := v.Args[0]
		v.reset(OpGenMTRUNC_I16)
		v.AddArg(x)
		return true
	}
}
func rewriteValueGenM_OpTrunc32to8_0(v *Value) bool {
	// match: (Trunc32to8 x)
	// cond:
	// result: (TRUNC_I8 x)
	for {
		x := v.Args[0]
		v.reset(OpGenMTRUNC_I8)
		v.AddArg(x)
		return true
	}
}
func rewriteValueGenM_OpTrunc64to16_0(v *Value) bool {
	// match: (Trunc64to16 x)
	// cond:
	// result: (TRUNC_I16 x)
	for {
		x := v.Args[0]
		v.reset(OpGenMTRUNC_I16)
		v.AddArg(x)
		return true
	}
}
func rewriteValueGenM_OpTrunc64to32_0(v *Value) bool {
	// match: (Trunc64to32 x)
	// cond:
	// result: (TRUNC_I32 x)
	for {
		x := v.Args[0]
		v.reset(OpGenMTRUNC_I32)
		v.AddArg(x)
		return true
	}
}
func rewriteValueGenM_OpTrunc64to8_0(v *Value) bool {
	// match: (Trunc64to8 x)
	// cond:
	// result: (TRUNC_I8 x)
	for {
		x := v.Args[0]
		v.reset(OpGenMTRUNC_I8)
		v.AddArg(x)
		return true
	}
}
func rewriteValueGenM_OpWB_0(v *Value) bool {
	// match: (WB {fn} destptr srcptr mem)
	// cond:
	// result: (LoweredWB {fn} destptr srcptr mem)
	for {
		fn := v.Aux
		_ = v.Args[2]
		destptr := v.Args[0]
		srcptr := v.Args[1]
		mem := v.Args[2]
		v.reset(OpGenMLoweredWB)
		v.Aux = fn
		v.AddArg(destptr)
		v.AddArg(srcptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueGenM_OpXor16_0(v *Value) bool {
	// match: (Xor16 x y)
	// cond:
	// result: (XOR_I16 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMXOR_I16)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpXor32_0(v *Value) bool {
	// match: (Xor32 x y)
	// cond:
	// result: (XOR_I32 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMXOR_I32)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpXor64_0(v *Value) bool {
	// match: (Xor64 x y)
	// cond:
	// result: (XOR_I64 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMXOR_I64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpXor8_0(v *Value) bool {
	// match: (Xor8 x y)
	// cond:
	// result: (XOR_I8 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMXOR_I8)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpZero_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ
	// match: (Zero [4] dst mem)
	// cond:
	// result: (ST_4 dst (CONST_I32 [0]) mem)
	for {
		if v.AuxInt != 4 {
			break
		}
		_ = v.Args[1]
		dst := v.Args[0]
		mem := v.Args[1]
		v.reset(OpGenMST_4)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpGenMCONST_I32, typ.Int32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [s] dst mem)
	// cond:
	// result: (LoweredZero [s] dst mem)
	for {
		s := v.AuxInt
		_ = v.Args[1]
		dst := v.Args[0]
		mem := v.Args[1]
		v.reset(OpGenMLoweredZero)
		v.AuxInt = s
		v.AddArg(dst)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueGenM_OpZeroExt16to32_0(v *Value) bool {
	// match: (ZeroExt16to32 x)
	// cond:
	// result: (ZEXT_I32 x)
	for {
		x := v.Args[0]
		v.reset(OpGenMZEXT_I32)
		v.AddArg(x)
		return true
	}
}
func rewriteValueGenM_OpZeroExt16to64_0(v *Value) bool {
	// match: (ZeroExt16to64 x)
	// cond:
	// result: (ZEXT_I64 x)
	for {
		x := v.Args[0]
		v.reset(OpGenMZEXT_I64)
		v.AddArg(x)
		return true
	}
}
func rewriteValueGenM_OpZeroExt32to64_0(v *Value) bool {
	// match: (ZeroExt32to64 x)
	// cond:
	// result: (ZEXT_I64 x)
	for {
		x := v.Args[0]
		v.reset(OpGenMZEXT_I64)
		v.AddArg(x)
		return true
	}
}
func rewriteValueGenM_OpZeroExt8to16_0(v *Value) bool {
	// match: (ZeroExt8to16 x)
	// cond:
	// result: (ZEXT_I16 x)
	for {
		x := v.Args[0]
		v.reset(OpGenMZEXT_I16)
		v.AddArg(x)
		return true
	}
}
func rewriteValueGenM_OpZeroExt8to32_0(v *Value) bool {
	// match: (ZeroExt8to32 x)
	// cond:
	// result: (ZEXT_I32 x)
	for {
		x := v.Args[0]
		v.reset(OpGenMZEXT_I32)
		v.AddArg(x)
		return true
	}
}
func rewriteValueGenM_OpZeroExt8to64_0(v *Value) bool {
	// match: (ZeroExt8to64 x)
	// cond:
	// result: (ZEXT_I64 x)
	for {
		x := v.Args[0]
		v.reset(OpGenMZEXT_I64)
		v.AddArg(x)
		return true
	}
}
func rewriteBlockGenM(b *Block) bool {
	config := b.Func.Config
	_ = config
	fe := b.Func.fe
	_ = fe
	typ := &config.Types
	_ = typ
	switch b.Kind {
	}
	return false
}
