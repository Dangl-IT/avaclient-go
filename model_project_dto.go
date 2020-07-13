/*
 * AVACloud API 1.16.0
 *
 * AVACloud API specification
 *
 * API version: 1.16.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package avaclient

// ProjectDto A Project contains all relevant information for a construction project.
type ProjectDto struct {
	// Elements GUID identifier.
	Id string `json:"id"`
	// This property controls the accuracy of all price properties, meaning how many decimal places are preserved in calculations. It defaults to DEFAULT_PRICE_ACCURACY. Please see the Dangl.AVA documentation for further information about decimal precision.
	PriceAccuracy int32 `json:"priceAccuracy"`
	// This forces total prices to be the strict product of quantities times unit price in positions. It is enabled by default. If this is disabled, both the unit price and the total price of positions is calculated from the non-rounded values. Please see the documentation for a more detailed explanation of this setting.
	ForceStrictTotals  bool                  `json:"forceStrictTotals"`
	PriceRoundingMode  PriceRoundingModeDto  `json:"priceRoundingMode"`
	ProjectInformation ProjectInformationDto `json:"projectInformation,omitempty"`
	// The ServiceSpecifications in this Project.
	ServiceSpecifications []ServiceSpecificationDto `json:"serviceSpecifications,omitempty"`
	// This is used to store the GAEB XML Id within this Project. This data is not used for any calculations or evaluations but only for GAEB serialization and deserialization.
	GaebXmlId string `json:"gaebXmlId,omitempty"`
}
