/*
 * Gravitee.io - Rest API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.22.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package graviteeapim

// UpdateRatingEntity struct for UpdateRatingEntity
type UpdateRatingEntity struct {
	Id      string `json:"id,omitempty"`
	Api     string `json:"api,omitempty"`
	Rate    string `json:"rate,omitempty"`
	Title   string `json:"title,omitempty"`
	Comment string `json:"comment,omitempty"`
}
