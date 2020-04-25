/*
 * Gravitee.io - Rest API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.22.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package graviteeapim

// AuditEntity struct for AuditEntity
type AuditEntity struct {
	Id            string            `json:"id,omitempty"`
	ReferenceId   string            `json:"referenceId,omitempty"`
	ReferenceType string            `json:"referenceType,omitempty"`
	User          string            `json:"user,omitempty"`
	CreatedAt     int64             `json:"createdAt,omitempty"`
	Event         string            `json:"event,omitempty"`
	Properties    map[string]string `json:"properties,omitempty"`
	Patch         string            `json:"patch,omitempty"`
}
