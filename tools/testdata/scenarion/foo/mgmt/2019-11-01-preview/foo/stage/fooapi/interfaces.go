// Copyright 2018 Microsoft Corporation
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

package fooapi

import "github.com/test-repo-arcturus/azure-sdk-for-go/tools/testdata/scenarion/foo/mgmt/2019-11-01-preview/foo"

// GatewaysClientAPI ...
type GatewaysClientAPI interface {
	CreateOrUpdate(resGroup string, parameters foo.Gateway) error
}
