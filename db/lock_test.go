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

func TestGenerateLockKey(t *testing.T) {
	tests := []struct {
		input  string
		output int64
	}{
		{"example", 3821494549623970275},
		{"testing", 8025471639231305070},
	}

	for _, test := range tests {
		assert.Equal(t, test.output, GenerateLockKey(test.input))
	}
}
