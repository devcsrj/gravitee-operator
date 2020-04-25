/*
 * Gravitee.io - Rest API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.22.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package graviteeapim

// MonitoringData struct for MonitoringData
type MonitoringData struct {
	Cpu     MonitoringCpu     `json:"cpu,omitempty"`
	Process MonitoringProcess `json:"process,omitempty"`
	Jvm     MonitoringJvm     `json:"jvm,omitempty"`
	Thread  MonitoringThread  `json:"thread,omitempty"`
	Gc      MonitoringGc      `json:"gc,omitempty"`
}
