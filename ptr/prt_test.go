// Copyright (c) 2025 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package ptr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestP(t *testing.T) {
	s := "hello"
	ps := P(s)
	require.NotNil(t, ps)
	require.Equal(t, s, *ps)
}
