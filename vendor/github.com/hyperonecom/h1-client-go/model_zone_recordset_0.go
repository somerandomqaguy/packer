/*
 * HyperOne API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ZoneRecordset0 struct for ZoneRecordset0
type ZoneRecordset0 struct {
	Record []ZoneRecordset0Record `json:"record,omitempty"`
	Type   string                 `json:"type,omitempty"`
	Name   string                 `json:"name,omitempty"`
	Ttl    float32                `json:"ttl,omitempty"`
	Id     string                 `json:"id,omitempty"`
}
