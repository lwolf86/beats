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

package docker

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "docker", asset.ModuleFieldsPri, AssetDocker); err != nil {
		panic(err)
	}
}

// AssetDocker returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/docker.
func AssetDocker() string {
	return "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"
}
