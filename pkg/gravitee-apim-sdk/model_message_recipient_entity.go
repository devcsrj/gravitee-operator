/*
 * Gravitee.io - Rest API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.22.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package graviteeapim

// MessageRecipientEntity struct for MessageRecipientEntity
type MessageRecipientEntity struct {
	Url       string   `json:"url,omitempty"`
	RoleScope string   `json:"role_scope,omitempty"`
	RoleValue []string `json:"role_value,omitempty"`
}
