/*
 * Gravitee.io - Rest API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.22.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package graviteeapim

// NewExternalUserEntity struct for NewExternalUserEntity
type NewExternalUserEntity struct {
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Email     string `json:"email"`
	Source    string `json:"source,omitempty"`
	Picture   string `json:"picture,omitempty"`
	SourceId  string `json:"sourceId,omitempty"`
}
