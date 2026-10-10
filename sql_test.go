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
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqliteIntToBool(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   int64
		want bool
	}{
		{"zero is false", 0, false},
		{"one is true", 1, true},
		{"any other positive value is true", 2, true},
		{"a negative value is true", -1, true},
		{"the largest value is true", math.MaxInt64, true},
		{"the smallest value is true", math.MinInt64, true},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tc.want, SqliteIntToBool(tc.in))
		})
	}
}

func TestSqliteBoolToInt(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   bool
		want int64
	}{
		{"true is one", true, 1},
		{"false is zero", false, 0},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tc.want, SqliteBoolToInt(tc.in))
		})
	}
}

func TestSqliteBoolRoundTrip(t *testing.T) {
	t.Parallel()

	t.Run("a bool survives the trip through an integer", func(t *testing.T) {
		t.Parallel()

		for _, b := range []bool{true, false} {
			assert.Equal(t, b, SqliteIntToBool(SqliteBoolToInt(b)))
		}
	})

	t.Run("an integer comes back as zero or one", func(t *testing.T) {
		t.Parallel()

		for _, i := range []int64{-5, -1, 0, 1, 2, 100} {
			want := int64(1)
			if i == 0 {
				want = 0
			}

			assert.Equal(t, want, SqliteBoolToInt(SqliteIntToBool(i)), "input %d", i)
		}
	})
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

func TestNullInt64Ptr(t *testing.T) {
	t.Parallel()

	zero, positive, negative, largest := int64(0), int64(42), int64(-7), int64(math.MaxInt64)

	tests := []struct {
		name string
		in   *int64
		want sql.NullInt64
	}{
		{"nil is NULL", nil, sql.NullInt64{}},
		{"zero is a value", &zero, sql.NullInt64{Int64: 0, Valid: true}},
		{"positive", &positive, sql.NullInt64{Int64: 42, Valid: true}},
		{"negative", &negative, sql.NullInt64{Int64: -7, Valid: true}},
		{"largest", &largest, sql.NullInt64{Int64: math.MaxInt64, Valid: true}},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := NullInt64Ptr(tc.in)
			assert.Equal(t, tc.want, got)

			// what the database driver is actually going to see
			v, err := got.Value()
			assert.NoError(t, err)
			if tc.want.Valid {
				assert.Equal(t, tc.want.Int64, v)
			} else {
				assert.Nil(t, v)
			}
		})
	}
}
