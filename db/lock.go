/*!
 * Copyright 2024 Mario Finelli
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

package db

import "hash/crc32"

// A query to try and obtain an advisory lock (for e.g., running database
// migrations).
const ADVISORY_LOCK_QUERY = "SELECT pg_try_advisory_lock($1) as unlocked"

// A query to release an advisory lock (e.g., after migrations are complete).
const ADVISORY_UNLOCK_QUERY = "SELECT pg_advisory_unlock($1) as unlocked"

// The GenerateLockKey function returns an integer that can be used as the
// parameter to the PostgreSQL advisory lock functions based on the name of
// the database for which we want to acquire the lock. Note that advisory lock
// keys are supposed to be 64 bits, and manual testing says that even something
// as simple as "123" will work, but the approach that Ruby on Rails takes is
// to hash the database name and then multiply it by a known salt to get a big
// number. I've decided to take their approach (and their salt).
// See: https://github.com/rails/rails/blob/main/activerecord/lib/active_record/migration.rb#L1608
func GenerateLockKey(dbname string) int64 {
	salt := int64(2053462845)
	hash := crc32.ChecksumIEEE([]byte(dbname))
	return salt * int64(hash)
}
