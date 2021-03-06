/*
 * AVACloud API 1.16.0
 *
 * AVACloud API specification
 *
 * API version: 1.16.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package avaclient

// CatalogueDto This class describes an external catalogue. Catalogues, or collections, hold information to categorize and describe items. For example, the German DIN 276 cost group standards describe different types of costs for building projects. When referencing the DIN 276 catalogue and providing an item key or identifier, it is possible to reference data in this catalogue.
type CatalogueDto struct {
	// Elements GUID identifier.
	Id string `json:"id"`
	// This is used to store the GAEB XML Id within this Catalogue. This data is not used for any calculations or evaluations but only for GAEB serialization and deserialization.
	GaebXmlId string `json:"gaebXmlId,omitempty"`
	// The name that is given for this catalogue.
	Name string `json:"name,omitempty"`
	// Additional information about this catalogue.
	Description   string           `json:"description,omitempty"`
	CatalogueType CatalogueTypeDto `json:"catalogueType"`
}
