/*
 * Gravitee.io - Rest API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.22.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package graviteeapim

// InstanceEntity struct for InstanceEntity
type InstanceEntity struct {
	Id               string            `json:"id,omitempty"`
	Hostname         string            `json:"hostname,omitempty"`
	Ip               string            `json:"ip,omitempty"`
	Port             string            `json:"port,omitempty"`
	Tenant           string            `json:"tenant,omitempty"`
	Version          string            `json:"version,omitempty"`
	Tags             []string          `json:"tags,omitempty"`
	State            string            `json:"state,omitempty"`
	SystemProperties map[string]string `json:"systemProperties,omitempty"`
	Plugins          []PluginEntity    `json:"plugins,omitempty"`
	StartedAt        int64             `json:"started_at,omitempty"`
	LastHeartbeatAt  int64             `json:"last_heartbeat_at,omitempty"`
	StoppedAt        int64             `json:"stopped_at,omitempty"`
}
