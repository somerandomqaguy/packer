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

package client

import (
	"encoding/json"
	"errors"
	"github.com/jdcloud-api/jdcloud-sdk-go/core"
	vpc "github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/apis"
)

type VpcClient struct {
	core.JDCloudClient
}

func NewVpcClient(credential *core.Credential) *VpcClient {
	if credential == nil {
		return nil
	}

	config := core.NewConfig()
	config.SetEndpoint("vpc.jdcloud-api.com")

	return &VpcClient{
		core.JDCloudClient{
			Credential:  *credential,
			Config:      *config,
			ServiceName: "vpc",
			Revision:    "0.5.1",
			Logger:      core.NewDefaultLogger(core.LogInfo),
		}}
}

func (c *VpcClient) SetConfig(config *core.Config) {
	c.Config = *config
}

func (c *VpcClient) SetLogger(logger core.Logger) {
	c.Logger = logger
}

/* 删除弹性Ip */
func (c *VpcClient) DeleteElasticIp(request *vpc.DeleteElasticIpRequest) (*vpc.DeleteElasticIpResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.DeleteElasticIpResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 修改VpcPeering接口 */
func (c *VpcClient) ModifyVpcPeering(request *vpc.ModifyVpcPeeringRequest) (*vpc.ModifyVpcPeeringResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.ModifyVpcPeeringResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 修改弹性IP */
func (c *VpcClient) ModifyElasticIp(request *vpc.ModifyElasticIpRequest) (*vpc.ModifyElasticIpResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.ModifyElasticIpResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 修改networkAcl接口 */
func (c *VpcClient) ModifyNetworkAclRules(request *vpc.ModifyNetworkAclRulesRequest) (*vpc.ModifyNetworkAclRulesResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.ModifyNetworkAclRulesResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 删除networkAcl接口 */
func (c *VpcClient) DeleteNetworkAcl(request *vpc.DeleteNetworkAclRequest) (*vpc.DeleteNetworkAclResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.DeleteNetworkAclResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 删除子网 */
func (c *VpcClient) DeleteSubnet(request *vpc.DeleteSubnetRequest) (*vpc.DeleteSubnetResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.DeleteSubnetResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 查询子网列表 */
func (c *VpcClient) DescribeSubnets(request *vpc.DescribeSubnetsRequest) (*vpc.DescribeSubnetsResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.DescribeSubnetsResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 查询弹性网卡列表 */
func (c *VpcClient) DescribeNetworkInterfaces(request *vpc.DescribeNetworkInterfacesRequest) (*vpc.DescribeNetworkInterfacesResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.DescribeNetworkInterfacesResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 移除安全组规则 */
func (c *VpcClient) RemoveNetworkSecurityGroupRules(request *vpc.RemoveNetworkSecurityGroupRulesRequest) (*vpc.RemoveNetworkSecurityGroupRulesResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.RemoveNetworkSecurityGroupRulesResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 修改安全组属性 */
func (c *VpcClient) ModifyNetworkSecurityGroup(request *vpc.ModifyNetworkSecurityGroupRequest) (*vpc.ModifyNetworkSecurityGroupResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.ModifyNetworkSecurityGroupResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 添加安全组规则 */
func (c *VpcClient) AddNetworkSecurityGroupRules(request *vpc.AddNetworkSecurityGroupRulesRequest) (*vpc.AddNetworkSecurityGroupRulesResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.AddNetworkSecurityGroupRulesResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 创建networkAcl接口 */
func (c *VpcClient) CreateNetworkAcl(request *vpc.CreateNetworkAclRequest) (*vpc.CreateNetworkAclResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.CreateNetworkAclResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 查询配额信息 */
func (c *VpcClient) DescribeQuota(request *vpc.DescribeQuotaRequest) (*vpc.DescribeQuotaResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.DescribeQuotaResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 创建路由表 */
func (c *VpcClient) CreateRouteTable(request *vpc.CreateRouteTableRequest) (*vpc.CreateRouteTableResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.CreateRouteTableResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 删除安全组 */
func (c *VpcClient) DeleteNetworkSecurityGroup(request *vpc.DeleteNetworkSecurityGroupRequest) (*vpc.DeleteNetworkSecurityGroupResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.DeleteNetworkSecurityGroupResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 查询路由表信息详情 */
func (c *VpcClient) DescribeRouteTable(request *vpc.DescribeRouteTableRequest) (*vpc.DescribeRouteTableResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.DescribeRouteTableResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 删除私有网络 */
func (c *VpcClient) DeleteVpc(request *vpc.DeleteVpcRequest) (*vpc.DeleteVpcResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.DeleteVpcResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 查询VpcPeering资源列表 */
func (c *VpcClient) DescribeVpcPeerings(request *vpc.DescribeVpcPeeringsRequest) (*vpc.DescribeVpcPeeringsResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.DescribeVpcPeeringsResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 查询Acl列表 */
func (c *VpcClient) DescribeNetworkAcls(request *vpc.DescribeNetworkAclsRequest) (*vpc.DescribeNetworkAclsResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.DescribeNetworkAclsResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 查询子网信息详情 */
func (c *VpcClient) DescribeSubnet(request *vpc.DescribeSubnetRequest) (*vpc.DescribeSubnetResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.DescribeSubnetResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 查询安全组信息详情 */
func (c *VpcClient) DescribeNetworkSecurityGroup(request *vpc.DescribeNetworkSecurityGroupRequest) (*vpc.DescribeNetworkSecurityGroupResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.DescribeNetworkSecurityGroupResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 查询安全组列表 */
func (c *VpcClient) DescribeNetworkSecurityGroups(request *vpc.DescribeNetworkSecurityGroupsRequest) (*vpc.DescribeNetworkSecurityGroupsResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.DescribeNetworkSecurityGroupsResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 查询Vpc信息详情 */
func (c *VpcClient) DescribeVpc(request *vpc.DescribeVpcRequest) (*vpc.DescribeVpcResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.DescribeVpcResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 给子网解绑NetworkAcl接口 */
func (c *VpcClient) DisassociateNetworkAcl(request *vpc.DisassociateNetworkAclRequest) (*vpc.DisassociateNetworkAclResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.DisassociateNetworkAclResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 移除networkAcl规则 */
func (c *VpcClient) RemoveNetworkAclRules(request *vpc.RemoveNetworkAclRulesRequest) (*vpc.RemoveNetworkAclRulesResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.RemoveNetworkAclRulesResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 路由表绑定子网接口 */
func (c *VpcClient) AssociateRouteTable(request *vpc.AssociateRouteTableRequest) (*vpc.AssociateRouteTableResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.AssociateRouteTableResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 查询私有网络列表 */
func (c *VpcClient) DescribeVpcs(request *vpc.DescribeVpcsRequest) (*vpc.DescribeVpcsResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.DescribeVpcsResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 创建安全组 */
func (c *VpcClient) CreateNetworkSecurityGroup(request *vpc.CreateNetworkSecurityGroupRequest) (*vpc.CreateNetworkSecurityGroupResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.CreateNetworkSecurityGroupResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 创建子网 */
func (c *VpcClient) CreateSubnet(request *vpc.CreateSubnetRequest) (*vpc.CreateSubnetResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.CreateSubnetResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 删除路由表 */
func (c *VpcClient) DeleteRouteTable(request *vpc.DeleteRouteTableRequest) (*vpc.DeleteRouteTableResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.DeleteRouteTableResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 查询VpcPeering资源详情 */
func (c *VpcClient) DescribeVpcPeering(request *vpc.DescribeVpcPeeringRequest) (*vpc.DescribeVpcPeeringResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.DescribeVpcPeeringResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 移除路由表规则 */
func (c *VpcClient) RemoveRouteTableRules(request *vpc.RemoveRouteTableRulesRequest) (*vpc.RemoveRouteTableRulesResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.RemoveRouteTableRulesResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 创建VpcPeering接口 */
func (c *VpcClient) CreateVpcPeering(request *vpc.CreateVpcPeeringRequest) (*vpc.CreateVpcPeeringResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.CreateVpcPeeringResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 给网卡分配secondaryIp接口 */
func (c *VpcClient) AssignSecondaryIps(request *vpc.AssignSecondaryIpsRequest) (*vpc.AssignSecondaryIpsResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.AssignSecondaryIpsResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 查询路由表列表 */
func (c *VpcClient) DescribeRouteTables(request *vpc.DescribeRouteTablesRequest) (*vpc.DescribeRouteTablesResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.DescribeRouteTablesResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 修改弹性网卡接口 */
func (c *VpcClient) ModifyNetworkInterface(request *vpc.ModifyNetworkInterfaceRequest) (*vpc.ModifyNetworkInterfaceResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.ModifyNetworkInterfaceResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 查询弹性ip列表 */
func (c *VpcClient) DescribeElasticIps(request *vpc.DescribeElasticIpsRequest) (*vpc.DescribeElasticIpsResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.DescribeElasticIpsResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 修改安全组规则 */
func (c *VpcClient) ModifyNetworkSecurityGroupRules(request *vpc.ModifyNetworkSecurityGroupRulesRequest) (*vpc.ModifyNetworkSecurityGroupRulesResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.ModifyNetworkSecurityGroupRulesResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 修改networkAcl接口 */
func (c *VpcClient) ModifyNetworkAcl(request *vpc.ModifyNetworkAclRequest) (*vpc.ModifyNetworkAclResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.ModifyNetworkAclResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 查询弹性网卡信息详情 */
func (c *VpcClient) DescribeNetworkInterface(request *vpc.DescribeNetworkInterfaceRequest) (*vpc.DescribeNetworkInterfaceResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.DescribeNetworkInterfaceResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 给网卡解绑弹性Ip接口 */
func (c *VpcClient) DisassociateElasticIp(request *vpc.DisassociateElasticIpRequest) (*vpc.DisassociateElasticIpResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.DisassociateElasticIpResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 删除VpcPeering接口 */
func (c *VpcClient) DeleteVpcPeering(request *vpc.DeleteVpcPeeringRequest) (*vpc.DeleteVpcPeeringResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.DeleteVpcPeeringResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 创建网卡接口，只能创建辅助网卡 */
func (c *VpcClient) CreateNetworkInterface(request *vpc.CreateNetworkInterfaceRequest) (*vpc.CreateNetworkInterfaceResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.CreateNetworkInterfaceResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 创建私有网络 */
func (c *VpcClient) CreateVpc(request *vpc.CreateVpcRequest) (*vpc.CreateVpcResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.CreateVpcResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 修改路由表规则 */
func (c *VpcClient) ModifyRouteTableRules(request *vpc.ModifyRouteTableRulesRequest) (*vpc.ModifyRouteTableRulesResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.ModifyRouteTableRulesResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 添加路由表规则 */
func (c *VpcClient) AddRouteTableRules(request *vpc.AddRouteTableRulesRequest) (*vpc.AddRouteTableRulesResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.AddRouteTableRulesResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* ElasticIp资源信息详情 */
func (c *VpcClient) DescribeElasticIp(request *vpc.DescribeElasticIpRequest) (*vpc.DescribeElasticIpResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.DescribeElasticIpResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 给网卡删除secondaryIp接口 */
func (c *VpcClient) UnassignSecondaryIps(request *vpc.UnassignSecondaryIpsRequest) (*vpc.UnassignSecondaryIpsResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.UnassignSecondaryIpsResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 修改路由表属性 */
func (c *VpcClient) ModifyRouteTable(request *vpc.ModifyRouteTableRequest) (*vpc.ModifyRouteTableResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.ModifyRouteTableResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 给网卡绑定弹性Ip接口 */
func (c *VpcClient) AssociateElasticIp(request *vpc.AssociateElasticIpRequest) (*vpc.AssociateElasticIpResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.AssociateElasticIpResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 给路由表解绑子网接口 */
func (c *VpcClient) DisassociateRouteTable(request *vpc.DisassociateRouteTableRequest) (*vpc.DisassociateRouteTableResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.DisassociateRouteTableResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 修改私有网络接口 */
func (c *VpcClient) ModifyVpc(request *vpc.ModifyVpcRequest) (*vpc.ModifyVpcResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.ModifyVpcResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 给子网绑定networkAcl接口 */
func (c *VpcClient) AssociateNetworkAcl(request *vpc.AssociateNetworkAclRequest) (*vpc.AssociateNetworkAclResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.AssociateNetworkAclResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 修改子网接口 */
func (c *VpcClient) ModifySubnet(request *vpc.ModifySubnetRequest) (*vpc.ModifySubnetResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.ModifySubnetResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 创建一个或者多个弹性Ip */
func (c *VpcClient) CreateElasticIps(request *vpc.CreateElasticIpsRequest) (*vpc.CreateElasticIpsResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.CreateElasticIpsResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 删除弹性网卡接口 */
func (c *VpcClient) DeleteNetworkInterface(request *vpc.DeleteNetworkInterfaceRequest) (*vpc.DeleteNetworkInterfaceResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.DeleteNetworkInterfaceResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 查询networkAcl资源详情 */
func (c *VpcClient) DescribeNetworkAcl(request *vpc.DescribeNetworkAclRequest) (*vpc.DescribeNetworkAclResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.DescribeNetworkAclResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 添加networkAcl规则接口 */
func (c *VpcClient) AddNetworkAclRules(request *vpc.AddNetworkAclRulesRequest) (*vpc.AddNetworkAclRulesResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vpc.AddNetworkAclRulesResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}
