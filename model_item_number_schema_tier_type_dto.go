/*
 * AVACloud API 1.16.0
 *
 * AVACloud API specification
 *
 * API version: 1.16.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package avaclient

// ItemNumberSchemaTierTypeDto Specifies the type an ItemNumberSchemaTier represents. For example, a tier may indicate to be used for positions or for groups.
type ItemNumberSchemaTierTypeDto string

// List of ItemNumberSchemaTierTypeDto
const (
	ITEMNUMBERSCHEMATIERTYPEDTO_UNDEFINED ItemNumberSchemaTierTypeDto = "Undefined"
	ITEMNUMBERSCHEMATIERTYPEDTO_INDEX     ItemNumberSchemaTierTypeDto = "Index"
	ITEMNUMBERSCHEMATIERTYPEDTO_POSITION  ItemNumberSchemaTierTypeDto = "Position"
	ITEMNUMBERSCHEMATIERTYPEDTO_GROUP     ItemNumberSchemaTierTypeDto = "Group"
	ITEMNUMBERSCHEMATIERTYPEDTO_LOT       ItemNumberSchemaTierTypeDto = "Lot"
)
