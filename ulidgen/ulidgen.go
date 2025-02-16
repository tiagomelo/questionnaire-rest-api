// Copyright (c) 2025 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package ulidgen

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

// GenerateULIDs generates n ULIDs with a 300ms interval between each one.
func GenerateULIDs(n int) []string {
	var ulids []string
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	for i := 0; i < n; i++ {
		ulids = append(ulids, ulid.MustNew(ulid.Timestamp(t), entropy).String())
		time.Sleep(300 * time.Millisecond)
	}
	return ulids
}
