/*
 * Gravitee.io - Rest API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.22.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package graviteeapim

// MonitoringProcess struct for MonitoringProcess
type MonitoringProcess struct {
	OpenFileDescriptors int32 `json:"open_file_descriptors,omitempty"`
	MaxFileDescriptors  int32 `json:"max_file_descriptors,omitempty"`
}
