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

package testutil

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DropAllTables does it what it says on the tin... it drops all of the tables
// that currently exist. Only suitable for testing, if you try to run this
// function outside of test it will return an error.
func DropAllTables(ctx context.Context, db *pgxpool.Pool) error {
	if !testing.Testing() {
		return fmt.Errorf("only drop all tables during testing")
	}

	query := "SELECT table_name AS table FROM information_schema.tables " +
		"WHERE table_schema = $1 AND table_type = $2 ORDER BY " +
		"table_name ASC;"

	rows, err := db.Query(ctx, query, "public", "BASE TABLE")
	if err != nil {
		return err
	}

	tables := []string{}
	for rows.Next() {
		var table string
		e := rows.Scan(&table)
		if e != nil {
			return e
		}
		tables = append(tables, table)
	}

	if rows.Err() != nil {
		return rows.Err()
	}

	if len(tables) == 0 {
		return nil
	}

	_, err = db.Exec(ctx, fmt.Sprintf("DROP TABLE %s;",
		strings.Join(tables, ", ")))
	if err != nil {
		return err
	}

	return nil
}
