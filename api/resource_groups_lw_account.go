//
// Author:: Darren Murray (<darren.murray@lacework.net>)
// Copyright:: Copyright 2021, Lacework Inc.
// License:: Apache License, Version 2.0
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
//

package api

import (
	"encoding/json"
	"strconv"

	"github.com/lacework/go-sdk/lwtime"
	"github.com/pkg/errors"
)

// Wildcard values for selecting all lacework accounts
var (
	LwAccountResourceGroupAllAccounts = []string{"*"}
)

// GetContainer gets a single LwAccount ResourceGroup matching the
// provided resource guid
func (svc *ResourceGroupsService) GetLwAccount(guid string) (
	response LwAccountResourceGroupResponse,
	err error,
) {
	var rawResponse resourceGroupWorkaroundResponse
	err = svc.get(guid, &rawResponse)
	if err != nil {
		return
	}

	return setLwAccountResponse(rawResponse)
}

// UpdateLwAccount updates a single LwAccount ResourceGroup on the Lacework Server
func (svc *ResourceGroupsService) UpdateLwAccount(data ResourceGroup) (
	response LwAccountResourceGroupResponse,
	err error,
) {
	if data == nil {
		err = errors.New("resource group must not be empty")
		return
	}
	guid := data.ID()
	data.ResetResourceGUID()

	err = svc.update(guid, data, &response)
	if err != nil {
		return
	}

	return
}

// CreateLwAccount creates a single LwAccount ResourceGroup on the Lacework Server
func (svc *ResourceGroupsService) CreateLwAccount(data ResourceGroup) (
	response LwAccountResourceGroupResponse,
	err error,
) {
	err = svc.create(data, &response)
	return
}

func setLwAccountResponse(response resourceGroupWorkaroundResponse) (lw LwAccountResourceGroupResponse, err error) {
	var props LwAccountResourceGroupJsonStringProps

	isDefault, err := strconv.Atoi(response.Data.IsDefault)
	if err != nil {
		return
	}

	lw = LwAccountResourceGroupResponse{
		Data: LwAccountResourceGroupData{
			Guid:         response.Data.Guid,
			IsDefault:    isDefault,
			ResourceGuid: response.Data.ResourceGuid,
			Name:         response.Data.Name,
			Type:         response.Data.Type,
			Enabled:      response.Data.Enabled,
		},
	}

	propsString, ok := response.Data.Props.(string)
	if !ok {
		err = errors.New("unable to cast props field from API response")
		return
	}

	err = json.Unmarshal([]byte(propsString), &props)
	if err != nil {
		return
	}
	lw.Data.Props = LwAccountResourceGroupProps(props)
	return
}

type LwAccountResourceGroupResponse struct {
	Data LwAccountResourceGroupData `json:"data"`
}

type LwAccountResourceGroupData struct {
	Guid         string                      `json:"guid,omitempty"`
	IsDefault    int                         `json:"isDefault,omitempty"`
	ResourceGuid string                      `json:"resourceGuid,omitempty"`
	Name         string                      `json:"resourceName"`
	Type         string                      `json:"resourceType"`
	Enabled      int                         `json:"enabled,omitempty"`
	Props        LwAccountResourceGroupProps `json:"props"`
}

type LwAccountResourceGroupProps struct {
	Description string        `json:"description,omitempty"`
	LwAccounts  []string      `json:"lwAccounts"`
	UpdatedBy   string        `json:"updatedBy,omitempty"`
	LastUpdated *lwtime.Epoch `json:"lastUpdated,omitempty"`
}

// Workaround for props being returned as a json string
type LwAccountResourceGroupJsonStringProps struct {
	Description string        `json:"DESCRIPTION,omitempty"`
	LwAccounts  []string      `json:"LW_ACCOUNTS"`
	UpdatedBy   string        `json:"UPDATED_BY,omitempty"`
	LastUpdated *lwtime.Epoch `json:"LAST_UPDATED,omitempty"`
}

func (props LwAccountResourceGroupProps) GetBaseProps() ResourceGroupPropsBase {
	return ResourceGroupPropsBase{
		Description: props.Description,
		UpdatedBy:   props.UpdatedBy,
		LastUpdated: props.LastUpdated,
	}
}

func (props LwAccountResourceGroupProps) MarshalJSON() ([]byte, error) {
	res := struct {
		Description string   `json:"description,omitempty"`
		LwAccounts  []string `json:"lwAccounts"`
		UpdatedBy   string   `json:"updatedBy,omitempty"`
		LastUpdated string   `json:"lastUpdated,omitempty"`
	}{
		Description: props.Description,
		LwAccounts:  props.LwAccounts,
		UpdatedBy:   props.UpdatedBy,
		LastUpdated: props.LastUpdated.String(),
	}
	return json.Marshal(&res)
}
