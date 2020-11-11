/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.2
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// CreateSecurityGroupRequest struct for CreateSecurityGroupRequest
type CreateSecurityGroupRequest struct {
	// A description for the security group, with a maximum length of 255 [ASCII printable characters](https://en.wikipedia.org/wiki/ASCII#Printable_characters).
	Description string `json:"Description"`
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun bool `json:"DryRun,omitempty"`
	// The ID of the Net for the security group.
	NetId string `json:"NetId,omitempty"`
	// (Public Cloud only) The name of the security group.<br /> This name must not start with `sg-`.</br> This name must be unique and contain between 1 and 255 ASCII characters. Accented letters are not allowed.
	SecurityGroupName string `json:"SecurityGroupName"`
}
