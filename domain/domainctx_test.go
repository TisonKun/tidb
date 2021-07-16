// Copyright 2015 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package domain

import (
	"testing"

	"github.com/pingcap/tidb/util/mock"
	"github.com/stretchr/testify/assert"
)

func TestDomain(t *testing.T) {
	ctx := mock.NewContext()
	assert.NotEqual(t, "", domainKey.String())

	BindDomain(ctx, nil)
	v := GetDomain(ctx)
	assert.Nil(t, v)

	ctx.ClearValue(domainKey)
	v = GetDomain(ctx)
	assert.Nil(t, v)
}
