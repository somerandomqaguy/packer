/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-10-18T06:16:13Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type InstanceTagParameter struct {

	// 태그키
	TagKey *string `json:"tagKey,omitempty"`

	// 태그값
	TagValue *string `json:"tagValue,omitempty"`
}
