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
	case OpAdd64:
		return rewriteValueGenM_OpAdd64_0(v)
	case OpAddr:
		return rewriteValueGenM_OpAddr_0(v)
	case OpConst16:
		return rewriteValueGenM_OpConst16_0(v)
	case OpConst32:
		return rewriteValueGenM_OpConst32_0(v)
	case OpConst64:
		return rewriteValueGenM_OpConst64_0(v)
	case OpConst8:
		return rewriteValueGenM_OpConst8_0(v)
	case OpEq8:
		return rewriteValueGenM_OpEq8_0(v)
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
	case OpLoad:
		return rewriteValueGenM_OpLoad_0(v)
	case OpLocalAddr:
		return rewriteValueGenM_OpLocalAddr_0(v)
	case OpNeq32:
		return rewriteValueGenM_OpNeq32_0(v)
	case OpNilCheck:
		return rewriteValueGenM_OpNilCheck_0(v)
	case OpOffPtr:
		return rewriteValueGenM_OpOffPtr_0(v)
	case OpStaticCall:
		return rewriteValueGenM_OpStaticCall_0(v)
	case OpStore:
		return rewriteValueGenM_OpStore_0(v)
	case OpWB:
		return rewriteValueGenM_OpWB_0(v)
	}
	return false
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
func rewriteValueGenM_OpEq8_0(v *Value) bool {
	// match: (Eq8 x y)
	// cond:
	// result: (CMP_EQ_I8 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_EQ_I8)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpGreater16_0(v *Value) bool {
	// match: (Greater16 x y)
	// cond:
	// result: (CMP_LT_I16 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_LT_I16)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpGreater16U_0(v *Value) bool {
	// match: (Greater16U x y)
	// cond:
	// result: (CMP_ULT_I16 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_ULT_I16)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpGreater32_0(v *Value) bool {
	// match: (Greater32 x y)
	// cond:
	// result: (CMP_LT_I32 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_LT_I32)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpGreater32F_0(v *Value) bool {
	// match: (Greater32F x y)
	// cond:
	// result: (CMP_OLT_F32 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_OLT_F32)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpGreater32U_0(v *Value) bool {
	// match: (Greater32U x y)
	// cond:
	// result: (CMP_ULT_I32 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_ULT_I32)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpGreater64_0(v *Value) bool {
	// match: (Greater64 x y)
	// cond:
	// result: (CMP_LT_I64 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_LT_I64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpGreater64F_0(v *Value) bool {
	// match: (Greater64F x y)
	// cond:
	// result: (CMP_OLT_F64 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_OLT_F64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpGreater64U_0(v *Value) bool {
	// match: (Greater64U x y)
	// cond:
	// result: (CMP_ULT_I64 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_ULT_I64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpGreater8_0(v *Value) bool {
	// match: (Greater8 x y)
	// cond:
	// result: (CMP_LT_I8 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_LT_I8)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueGenM_OpGreater8U_0(v *Value) bool {
	// match: (Greater8U x y)
	// cond:
	// result: (CMP_ULT_I8 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_ULT_I8)
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
func rewriteValueGenM_OpNeq32_0(v *Value) bool {
	// match: (Neq32 x y)
	// cond:
	// result: (CMP_NE_I64 x y)
	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpGenMCMP_NE_I64)
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
	// result: (ST_8_F ptr val mem)
	for {
		t := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is64BitFloat(t.(*types.Type))) {
			break
		}
		v.reset(OpGenMST_8_F)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: is32BitFloat(t.(*types.Type))
	// result: (ST_4_F ptr val mem)
	for {
		t := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32BitFloat(t.(*types.Type))) {
			break
		}
		v.reset(OpGenMST_4_F)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 8
	// result: (ST_8_I ptr val mem)
	for {
		t := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Size() == 8) {
			break
		}
		v.reset(OpGenMST_8_I)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 4
	// result: (ST_4_I ptr val mem)
	for {
		t := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Size() == 4) {
			break
		}
		v.reset(OpGenMST_4_I)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 2
	// result: (ST_2_I ptr val mem)
	for {
		t := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Size() == 2) {
			break
		}
		v.reset(OpGenMST_2_I)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 1
	// result: (ST_1_I ptr val mem)
	for {
		t := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Size() == 1) {
			break
		}
		v.reset(OpGenMST_1_I)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
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
