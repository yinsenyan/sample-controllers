/*
Copyright 2022 The Clusternet Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package statefulsetplus

import (
	"encoding/json"

	stspapi "git.woa.com/STKE/statefulsetplus-operator-api/pkg/apis/platform.stke/v1alpha1"
	appsapi "github.com/clusternet/clusternet/pkg/apis/apps/v1alpha1"
	"github.com/clusternet/clusternet/pkg/controllers/apps/feedinventory/utils"
)

const stsp = "StatefulSetPlus"

type Plugin struct {
	name string
}

func NewPlugin() *Plugin {
	return &Plugin{
		name: stsp,
	}
}

func (pl *Plugin) Parser(rawData []byte) (*int32, appsapi.ReplicaRequirements, string, error) {
	var stsp stspapi.StatefulSetPlus
	if err := json.Unmarshal(rawData, &stsp); err != nil {
		return nil, appsapi.ReplicaRequirements{}, "", err
	}
	desiredReplicas := int32(*stsp.Spec.Replicas)

	return &desiredReplicas, utils.GetReplicaRequirements(stsp.Spec.Template.Spec), "/spec/replicas", nil
}

func (pl *Plugin) Name() string {
	return pl.name
}

func (pl *Plugin) Kind() string {
	return stsp
}