// Copyright 2025 The kmeta Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
)

func main() {
	iris := Load()
	fmt.Println(len(iris))
}
