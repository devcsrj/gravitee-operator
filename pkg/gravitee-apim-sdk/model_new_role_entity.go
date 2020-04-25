/*
 * Gravitee.io - Rest API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.22.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package graviteeapim

// NewRoleEntity struct for NewRoleEntity
type NewRoleEntity struct {
	Name        string              `json:"name"`
	Description string              `json:"description,omitempty"`
	Scope       string              `json:"scope"`
	Permissions map[string][]string `json:"permissions,omitempty"`
	Default     bool                `json:"default,omitempty"`
}
