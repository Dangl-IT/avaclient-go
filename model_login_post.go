/*
 * AVACloud API 1.16.0
 *
 * AVACloud API specification
 *
 * API version: 1.16.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package avaclient

// LoginPost struct for LoginPost
type LoginPost struct {
	Identifier   string `json:"identifier"`
	Password     string `json:"password"`
	StaySignedIn bool   `json:"staySignedIn"`
}