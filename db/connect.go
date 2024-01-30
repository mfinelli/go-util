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

// The PostgresqlConnectionDetails type is a representation of the connection
// options that we care about for connecting to a PostgreSQL server that we
// can transform into a data source name (DSN).
type PostgresqlConnectionDetails struct {
	Hostname string // PostgreSQL server hostname
	Username string // PostgreSQL server username
	Password string // PostgreSQL server password
	Database string // PostgreSQL database name
	Port     int    // PostgreSQL server port
	Sslmode  string // PostgreSQL SSL mode
	Timezone string // PostgreSQL optional timezone connection parameter
}

// The BuildDSN function creates a DSN from the connection details that can
// be passed directly to the database open functions.
func (d *PostgresqlConnectionDetails) BuildDSN() string {
	parts := []string{}

	parts = append(parts, fmt.Sprintf("host=%s",
		escapePsqlConnectionValue(d.Hostname)))
	parts = append(parts, fmt.Sprintf("user=%s",
		escapePsqlConnectionValue(d.Username)))

	if d.Password != "" {
		parts = append(parts, fmt.Sprintf("password=%s",
			escapePsqlConnectionValue(d.Password)))
	}

	parts = append(parts, fmt.Sprintf("dbname=%s",
		escapePsqlConnectionValue(d.Database)))
	parts = append(parts, fmt.Sprintf("port=%d", d.Port))

	if d.Sslmode != "" {
		parts = append(parts, fmt.Sprintf("sslmode=%s",
			escapePsqlConnectionValue(d.Sslmode)))
	}

	if d.Timezone != "" {
		parts = append(parts, fmt.Sprintf("TimeZone=%s",
			escapePsqlConnectionValue(d.Timezone)))
	}

	return strings.Join(parts, " ")
}

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
