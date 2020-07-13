/*
 * AVACloud API 1.16.0
 *
 * AVACloud API specification
 *
 * API version: 1.16.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package avaclient

// NoteTextDto struct for NoteTextDto
type NoteTextDto struct {
	// Elements GUID identifier.
	Id string `json:"id"`
	// This is used to store the GAEB XML Id within this IElement. This data is not used for any calculations or evaluations but only for GAEB serialization and deserialization.
	GaebXmlId                string `json:"gaebXmlId,omitempty"`
	ElementTypeDiscriminator string `json:"elementTypeDiscriminator"`
	// If this is set to true, this text is meant to not be seen as part of the regular elements hierarchy but as a special opening text at the beginning of the project. For example, in GAEB XML, this would map to the GAEB.Award.AddText. Typically, such texts describe project-wide contractual definitions. If this is set to true, this NoteText should be placed at the top of the elements hierarchy directly in the ServiceSpecification.Elements group, otherwise it will likely not be treated correctly when exporting to GAEB. You can only set IsOpeningText or IsClosingText to true.
	IsOpeningText bool `json:"isOpeningText"`
	// If this is set to true, this text is meant to not be seen as part of the regular elements hierarchy but as a special closing text at the end of the project. For Example, in GAEB XML, this would map to the GAEB.AddText. Typically, such texts are used to describe project wide finishing descriptions. If this is set to true, this NoteText should be placed at the top of the elements hierarchy directly in the ServiceSpecification.Elements group, otherwise it will likely not be treated correctly when exporting to GAEB. You can only set IsOpeningText or IsClosingText to true.
	IsClosingText bool `json:"isClosingText"`
	// Short description for this DescriptionBase element.
	ShortText    string          `json:"shortText,omitempty"`
	AdditionType AdditionTypeDto `json:"additionType"`
	// Detailed description for this DescriptionBase element. When the HtmlLongText is set, this is automatically overwritten and filled with the appropriate plain text representation of the Html text. Vice versa, setting this property overrides the HtmlLongText.
	LongText string `json:"longText,omitempty"`
	// This contains the Html representation of the Longtext. When the LongText is set, this is automatically overwritten and filled with the appropriate Html representation of the plaintext. Vice versa, setting this property overrides the LongText. GAEB 90 and GAEB 2000 exports do not support any image functionality. In GAEB XML, only images that use an embedded Base64 data uri are exported, regular url references are cleared before written out.
	HtmlLongText            string                     `json:"htmlLongText,omitempty"`
	StandardizedDescription StandardizedDescriptionDto `json:"standardizedDescription,omitempty"`
	ElementType             string                     `json:"elementType,omitempty"`
}
