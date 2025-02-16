// Copyright (c) 2025 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package ptr

// P returns a pointer to the value passed as argument.
func P[T any](v T) *T {
	return &v
}
