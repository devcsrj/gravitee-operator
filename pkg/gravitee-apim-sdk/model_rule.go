/*
 * Gravitee.io - Rest API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.22.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package graviteeapim

// Rule struct for Rule
type Rule struct {
	Methods     []string `json:"methods,omitempty"`
	Policy      Policy   `json:"policy,omitempty"`
	Description string   `json:"description,omitempty"`
	Enabled     bool     `json:"enabled,omitempty"`
}
