/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-10-18T06:16:13Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type RebootServerInstancesRequest struct {

	// 서버인스턴스번호리스트
	ServerInstanceNoList []*string `json:"serverInstanceNoList"`
}
