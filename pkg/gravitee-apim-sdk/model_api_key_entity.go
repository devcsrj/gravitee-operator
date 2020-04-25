/*
 * Gravitee.io - Rest API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.22.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package graviteeapim

// ApiKeyEntity struct for ApiKeyEntity
type ApiKeyEntity struct {
	Key          string `json:"key,omitempty"`
	Subscription string `json:"subscription,omitempty"`
	Application  string `json:"application,omitempty"`
	Plan         string `json:"plan,omitempty"`
	Revoked      bool   `json:"revoked,omitempty"`
	Paused       bool   `json:"paused,omitempty"`
	ExpireAt     int64  `json:"expire_at,omitempty"`
	CreatedAt    int64  `json:"created_at,omitempty"`
	UpdatedAt    int64  `json:"updated_at,omitempty"`
	RevokedAt    int64  `json:"revoked_at,omitempty"`
}
