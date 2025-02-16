// Copyright (c) 2025 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package ulidgen

import "testing"

func TestGenerateULIDs(t *testing.T) {
	ulids := GenerateULIDs(3)
	if len(ulids) != 3 {
		t.Errorf("GenerateULIDs(5) returned %d ULIDs, expected 3", len(ulids))
	}
	for _, ulid := range ulids {
		if len(ulid) != 26 {
			t.Errorf("ULID %s has length %d, expected 26", ulid, len(ulid))
		}
	}
}
