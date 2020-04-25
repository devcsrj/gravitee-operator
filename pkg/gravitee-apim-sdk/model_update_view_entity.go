/*
 * Gravitee.io - Rest API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.22.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package graviteeapim

// UpdateViewEntity struct for UpdateViewEntity
type UpdateViewEntity struct {
	Id           string `json:"id,omitempty"`
	Name         string `json:"name"`
	Description  string `json:"description,omitempty"`
	DefaultView  bool   `json:"defaultView,omitempty"`
	Hidden       bool   `json:"hidden,omitempty"`
	Order        int32  `json:"order,omitempty"`
	HighlightApi string `json:"highlightApi,omitempty"`
}
