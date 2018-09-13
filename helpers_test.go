// Copyright 2017 Improbable. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package grpcweb_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	"github.com/yannic/grpcweb"
	"github.com/yannic/grpcweb/testproto"
)

func TestListGRPCResources(t *testing.T) {
	server := grpc.NewServer()
	testproto.RegisterTestServiceServer(server, &testServiceImpl{})
	expected := []string{
		"/improbable.grpcweb.test.TestService/PingEmpty",
		"/improbable.grpcweb.test.TestService/Ping",
		"/improbable.grpcweb.test.TestService/PingError",
		"/improbable.grpcweb.test.TestService/PingList",
		"/improbable.grpcweb.test.TestService/Echo",
		"/improbable.grpcweb.test.TestService/PingPongBidi",
		"/improbable.grpcweb.test.TestService/PingStream",
	}
	actual := grpcweb.ListGRPCResources(server)
	sort.Strings(expected)
	sort.Strings(actual)
	assert.EqualValues(t,
		expected,
		actual,
		"list grpc resources must provide an exhaustive list of all registered handlers")
}
