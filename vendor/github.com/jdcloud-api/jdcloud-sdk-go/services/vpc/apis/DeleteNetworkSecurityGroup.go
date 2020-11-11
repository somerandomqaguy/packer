// Copyright 2018 JDCLOUD.COM
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
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
	"github.com/jdcloud-api/jdcloud-sdk-go/core"
)

type DeleteNetworkSecurityGroupRequest struct {
	core.JDCloudRequest

	/* Region ID  */
	RegionId string `json:"regionId"`

	/* NetworkSecurityGroup ID  */
	NetworkSecurityGroupId string `json:"networkSecurityGroupId"`
}

/*
 * param regionId: Region ID (Required)
 * param networkSecurityGroupId: NetworkSecurityGroup ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteNetworkSecurityGroupRequest(
	regionId string,
	networkSecurityGroupId string,
) *DeleteNetworkSecurityGroupRequest {

	return &DeleteNetworkSecurityGroupRequest{
		JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/networkSecurityGroups/{networkSecurityGroupId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
		RegionId:               regionId,
		NetworkSecurityGroupId: networkSecurityGroupId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param networkSecurityGroupId: NetworkSecurityGroup ID (Required)
 */
func NewDeleteNetworkSecurityGroupRequestWithAllParams(
	regionId string,
	networkSecurityGroupId string,
) *DeleteNetworkSecurityGroupRequest {

	return &DeleteNetworkSecurityGroupRequest{
		JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/networkSecurityGroups/{networkSecurityGroupId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
		RegionId:               regionId,
		NetworkSecurityGroupId: networkSecurityGroupId,
	}
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteNetworkSecurityGroupRequestWithoutParam() *DeleteNetworkSecurityGroupRequest {

	return &DeleteNetworkSecurityGroupRequest{
		JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/networkSecurityGroups/{networkSecurityGroupId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
	}
}

/* param regionId: Region ID(Required) */
func (r *DeleteNetworkSecurityGroupRequest) SetRegionId(regionId string) {
	r.RegionId = regionId
}

/* param networkSecurityGroupId: NetworkSecurityGroup ID(Required) */
func (r *DeleteNetworkSecurityGroupRequest) SetNetworkSecurityGroupId(networkSecurityGroupId string) {
	r.NetworkSecurityGroupId = networkSecurityGroupId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteNetworkSecurityGroupRequest) GetRegionId() string {
	return r.RegionId
}

type DeleteNetworkSecurityGroupResponse struct {
	RequestID string                           `json:"requestId"`
	Error     core.ErrorResponse               `json:"error"`
	Result    DeleteNetworkSecurityGroupResult `json:"result"`
}

type DeleteNetworkSecurityGroupResult struct {
}
