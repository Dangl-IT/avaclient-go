/*
 * AVACloud API 1.16.0
 *
 * AVACloud API specification
 *
 * API version: 1.16.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package avaclient

// RegisterPost struct for RegisterPost
type RegisterPost struct {
	Username           string   `json:"username"`
	Email              string   `json:"email"`
	Password           string   `json:"password"`
	PreferredLanguages []string `json:"preferredLanguages,omitempty"`
}