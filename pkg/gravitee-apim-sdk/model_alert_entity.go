/*
 * Gravitee.io - Rest API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.22.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package graviteeapim

// AlertEntity struct for AlertEntity
type AlertEntity struct {
	Id            string  `json:"id,omitempty"`
	Name          string  `json:"name,omitempty"`
	Description   string  `json:"description,omitempty"`
	Type          string  `json:"type,omitempty"`
	Enabled       bool    `json:"enabled,omitempty"`
	MetricType    string  `json:"metricType,omitempty"`
	Metric        string  `json:"metric,omitempty"`
	ThresholdType string  `json:"thresholdType,omitempty"`
	Threshold     float64 `json:"threshold,omitempty"`
	Plan          string  `json:"plan,omitempty"`
	ReferenceType string  `json:"reference_type,omitempty"`
	ReferenceId   string  `json:"reference_id,omitempty"`
	CreatedAt     int64   `json:"created_at,omitempty"`
	UpdatedAt     int64   `json:"updated_at,omitempty"`
}
