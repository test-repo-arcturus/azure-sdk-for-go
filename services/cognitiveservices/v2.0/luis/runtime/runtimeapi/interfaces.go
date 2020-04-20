package runtimeapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/satori/go.uuid"
	"github.com/test-repo-arcturus/azure-sdk-for-go/services/cognitiveservices/v2.0/luis/runtime"
)

// PredictionClientAPI contains the set of methods on the PredictionClient type.
type PredictionClientAPI interface {
	Resolve(ctx context.Context, appID uuid.UUID, query string, timezoneOffset *float64, verbose *bool, staging *bool, spellCheck *bool, bingSpellCheckSubscriptionKey string, logParameter *bool) (result runtime.LuisResult, err error)
}

var _ PredictionClientAPI = (*runtime.PredictionClient)(nil)
