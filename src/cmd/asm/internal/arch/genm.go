// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package arch defines architecture-specific information and support functions.
package arch

var genMJump = map[string]bool{
}

func jumpGenM(word string) bool {
  return genMJump[word]
}
