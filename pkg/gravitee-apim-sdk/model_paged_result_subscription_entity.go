/*
 * Gravitee.io - Rest API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.22.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package graviteeapim

// PagedResultSubscriptionEntity struct for PagedResultSubscriptionEntity
type PagedResultSubscriptionEntity struct {
	Data     []SubscriptionEntity                         `json:"data,omitempty"`
	Metadata map[string]map[string]map[string]interface{} `json:"metadata,omitempty"`
	Page     Page                                         `json:"page,omitempty"`
}
