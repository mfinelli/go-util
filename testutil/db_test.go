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
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
)

func TestDropAllTables(t *testing.T) {
	conn := "postgres://goutil:goutil@localhost:5432/goutil"
	ctx := context.Background()
	db, err := pgxpool.New(ctx, conn)
	defer db.Close()
	assert.Nil(t, err)

	query := "SELECT COUNT(1) FROM information_schema.tables " +
		"WHERE table_schema = 'public';"

	{
		var count int
		err = db.QueryRow(ctx, query).Scan(&count)
		assert.Nil(t, err)
		assert.Zero(t, count)
	}

	_, err = db.Exec(ctx, "CREATE TABLE testtable (id INT GENERATED " +
		"ALWAYS AS IDENTITY, CONSTRAINT pkey PRIMARY KEY (id));")
	assert.Nil(t, err)

	{
		var count int
		err = db.QueryRow(ctx, query).Scan(&count)
		assert.Nil(t, err)
		assert.Equal(t, 1, count)
	}

	err = DropAllTables(ctx, db)
	assert.Nil(t, err)

	{
		var count int
		err = db.QueryRow(ctx, query).Scan(&count)
		assert.Nil(t, err)
		assert.Zero(t, count)
	}
}
