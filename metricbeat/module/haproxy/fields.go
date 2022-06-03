// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package haproxy

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "haproxy", asset.ModuleFieldsPri, AssetHaproxy); err != nil {
		panic(err)
	}
}

// AssetHaproxy returns asset data.
// This is the base64 encoded gzipped contents of module/haproxy.
func AssetHaproxy() string {
	return "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"
}