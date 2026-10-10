/*!
 * Copyright 2026 Mario Finelli
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package util

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqliteIntToBool(t *testing.T) {
	assert.True(t, SqliteIntToBool(1))
	assert.False(t, SqliteIntToBool(0))
}

func TestSqliteBoolToInt(t *testing.T) {
	assert.Equal(t, int64(1), SqliteBoolToInt(true))
	assert.Equal(t, int64(0), SqliteBoolToInt(false))
}

func TestNullString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		want sql.NullString
	}{
		{"empty is NULL", "", sql.NullString{}},
		{"text", "hello", sql.NullString{String: "hello", Valid: true}},
		{"whitespace is a value", " ", sql.NullString{String: " ", Valid: true}},
		{"zero as text is a value", "0", sql.NullString{String: "0", Valid: true}},
		{"unicode", "módctl-モッド", sql.NullString{String: "módctl-モッド", Valid: true}},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := NullString(tc.in)
			assert.Equal(t, tc.want, got)

			// what the database driver is actually going to see
			v, err := got.Value()
			assert.NoError(t, err)
			if tc.want.Valid {
				assert.Equal(t, tc.in, v)
			} else {
				assert.Nil(t, v)
			}
		})
	}
}

func TestNullStringPtr(t *testing.T) {
	t.Parallel()

	empty, text, space := "", "hello", " "

	tests := []struct {
		name string
		in   *string
		want sql.NullString
	}{
		{"nil is NULL", nil, sql.NullString{}},
		{"pointer to empty is NULL", &empty, sql.NullString{}},
		{"pointer to text", &text, sql.NullString{String: "hello", Valid: true}},
		{"pointer to whitespace is a value", &space, sql.NullString{String: " ", Valid: true}},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tc.want, NullStringPtr(tc.in))
		})
	}
}
