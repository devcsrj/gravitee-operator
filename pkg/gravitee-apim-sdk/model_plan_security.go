/*
 * Gravitee.io - Rest API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.22.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package graviteeapim

// PlanSecurity struct for PlanSecurity
type PlanSecurity struct {
	Apikey  Enabled `json:"apikey,omitempty"`
	Oauth2  Enabled `json:"oauth2,omitempty"`
	Keyless Enabled `json:"keyless,omitempty"`
	Jwt     Enabled `json:"jwt,omitempty"`
}
