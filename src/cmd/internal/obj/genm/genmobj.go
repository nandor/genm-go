// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package genm

import (
	"cmd/internal/obj"
	"cmd/internal/sys"
	"log"
	"fmt"
	"strings"
)

var Register = map[string]int16{
}

var unaryDst = map[obj.As]bool{
}

var LinkGenM = obj.LinkArch{
	Arch:       sys.ArchGenM,
	Init:       instinit,
	Preprocess: preprocess,
	Assemble:   assemble,
	Progedit:   progedit,
	UnaryDst:   unaryDst,
}

func instinit(ctxt *obj.Link) {
}

func preprocess(ctxt *obj.Link, s *obj.LSym, newprog obj.ProgAlloc) {
	log.Printf("not implemented: preprocess")
}

func assemble(ctxt *obj.Link, s *obj.LSym, newprog obj.ProgAlloc) {
	var sb strings.Builder

	switch s.Name {

	}

	for p := s.Func.Text; p != nil; p = p.Link {
		switch p.As {
		case obj.AFUNCDATA:
			// ignored
			continue

		case obj.APCDATA:
			// ignored
			continue

		case obj.ATEXT:
			// ignored
			continue

  	case obj.AUNDEF:
  		sb.WriteString(fmt.Sprintf("\tATRAP\n"))

		case ACONST:
			sb.WriteString(fmt.Sprintf("\tACONST\n"))
  	case AADDR:
  		sb.WriteString(fmt.Sprintf("\tAADDR\n"))
  	case ALD:
  		sb.WriteString(fmt.Sprintf("\tALD\n"))
  	case AST_1:
  		sb.WriteString(fmt.Sprintf("\tAST_1\n"))
  	case AST_8:
  		sb.WriteString(fmt.Sprintf("\tAST_8\n"))
  	case ACMP_EQ:
  		sb.WriteString(fmt.Sprintf("\tACMP_EQ\n"))
  	case ACMP_ULT:
  		sb.WriteString(fmt.Sprintf("\tACMP_ULT\n"))
  	case AJT:
  		sb.WriteString(fmt.Sprintf("\tAJT\n"))
  	case ARET:
  		sb.WriteString(fmt.Sprintf("\tARET\n"))
  	case AARG:
  		sb.WriteString(fmt.Sprintf("\tAARG\n"))
		default:
			log.Printf("assemble not implemented: %d %d\n", p.As, obj.AUNDEF)
		}
	}
	log.Printf("\n%s\n", sb.String())
	s.P = []byte(sb.String())
}

func progedit(ctxt *obj.Link, p *obj.Prog, newprog obj.ProgAlloc) {
	log.Printf("not implemented: progedit")
}
