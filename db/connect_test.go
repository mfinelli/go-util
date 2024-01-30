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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildDSN(t *testing.T) {
	tests := []struct {
		input  PostgresqlConnectionDetails
		output string
	}{
		{
			PostgresqlConnectionDetails{
				"host", "user", "pass", "db", 5432, "", "",
			},
			"host=host user=user password=pass dbname=db " +
				"port=5432",
		},
		{
			PostgresqlConnectionDetails{
				"host", "user", "", "db", 5432, "require",
				"UTC",
			},
			"host=host user=user dbname=db port=5432 " +
				"sslmode=require TimeZone=UTC",
		},
		{
			PostgresqlConnectionDetails{
				"host", "", "", "db", 5432, "", "",
			},
			"host=host user='' dbname=db port=5432",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.output, test.input.BuildDSN())
	}
}

func TestEscapePsqlConnectionValue(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"normal", "normal"},
		{"normal with spaces", "'normal with spaces'"},
		{"", "''"},
		{"it's \\ complicated", "'it\\'s \\\\ complicated'"},
	}

	for _, test := range tests {
		assert.Equal(t, test.output,
			escapePsqlConnectionValue(test.input))
	}
}
