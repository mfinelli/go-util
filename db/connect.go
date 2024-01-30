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

import (
	"fmt"
	"strings"
)

// The escapePsqlConnectionValue function escapes the provided string so that
// it can be used correctly in a PostgreSQL DSN.
// https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING
func escapePsqlConnectionValue(str string) string {
	if strings.Contains(str, "'") || strings.Contains(str, "\\") ||
		strings.Contains(str, " ") {
		str = strings.ReplaceAll(str, "\\", "\\\\")
		str = strings.ReplaceAll(str, "'", "\\'")
		return fmt.Sprintf("'%s'", str)
	} else if str == "" {
		return "''"
	} else {
		return str
	}
}
