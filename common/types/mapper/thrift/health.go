// Copyright (c) 2017-2020 Uber Technologies Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package thrift

import (
	"github.com/uber/cadence/.gen/go/health"
	"github.com/uber/cadence/common/types"
)

// FromHealthStatus converts internal HealthStatus type to thrift
func FromHealthStatus(t *types.HealthStatus) *health.HealthStatus {
	if t == nil {
		return nil
	}
	return &health.HealthStatus{
		Ok:  t.Ok,
		Msg: t.Msg,
	}
}

// ToHealthStatus converts thrift HealthStatus type to internal
func ToHealthStatus(t *health.HealthStatus) *types.HealthStatus {
	if t == nil {
		return nil
	}
	return &types.HealthStatus{
		Ok:  t.Ok,
		Msg: t.Msg,
	}
}

// FromMetaHealthArgs converts internal Meta_Health_Args type to thrift
func FromMetaHealthArgs(t *types.MetaHealthArgs) *health.Meta_Health_Args {
	if t == nil {
		return nil
	}
	return &health.Meta_Health_Args{}
}

// ToMetaHealthArgs converts thrift Meta_Health_Args type to internal
func ToMetaHealthArgs(t *health.Meta_Health_Args) *types.MetaHealthArgs {
	if t == nil {
		return nil
	}
	return &types.MetaHealthArgs{}
}

// FromMetaHealthResult converts internal Meta_Health_Result type to thrift
func FromMetaHealthResult(t *types.MetaHealthResult) *health.Meta_Health_Result {
	if t == nil {
		return nil
	}
	return &health.Meta_Health_Result{
		Success: FromHealthStatus(t.Success),
	}
}

// ToMetaHealthResult converts thrift Meta_Health_Result type to internal
func ToMetaHealthResult(t *health.Meta_Health_Result) *types.MetaHealthResult {
	if t == nil {
		return nil
	}
	return &types.MetaHealthResult{
		Success: ToHealthStatus(t.Success),
	}
}
