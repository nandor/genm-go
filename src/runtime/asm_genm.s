// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_asm.h"
#include "go_tls.h"
#include "funcdata.h"
#include "textflag.h"


TEXT _rt0_amd64(SB),NOSPLIT,$0
  ARG_I64     R0, 0

TEXT runtime·rt0_go(SB),NOSPLIT,$0
  UNDEF
