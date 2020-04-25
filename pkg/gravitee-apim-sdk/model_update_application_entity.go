/*
 * Gravitee.io - Rest API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.22.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package graviteeapim

// UpdateApplicationEntity struct for UpdateApplicationEntity
type UpdateApplicationEntity struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Type        string   `json:"type,omitempty"`
	ClientId    string   `json:"clientId,omitempty"`
	Groups      []string `json:"groups,omitempty"`
}
