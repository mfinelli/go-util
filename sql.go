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

import "database/sql"

// SqliteIntToBool converts an integer returned from sqlite (which doesn't
// have a proper boolean type) into a boolean because I can never remember
// which value is true and which is false. (1 is true and 0 is false)
func SqliteIntToBool(b int64) bool {
	return b == 1
}

// SqliteBoolToInt converts a bool into the correct sqlite integer
func SqliteBoolToInt(b bool) int64 {
	if b {
		return 1
	}

	return 0
}

// NullString converts a string into a sql.NullString that is NULL when the
// string is empty, for columns where NULL means "no value". Nothing is
// trimmed: a string of only whitespace is a value.
func NullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

// NullStringPtr is like NullString for a *string: NULL when s is nil or points
// to an empty string.
func NullStringPtr(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{}
	}

	return NullString(*s)
}

// NullInt64Ptr converts an *int64 into a sql.NullInt64 that is NULL when i is
// nil. Unlike with strings, zero is a value and not NULL.
func NullInt64Ptr(i *int64) sql.NullInt64 {
	if i == nil {
		return sql.NullInt64{}
	}

	return sql.NullInt64{Int64: *i, Valid: true}
}
