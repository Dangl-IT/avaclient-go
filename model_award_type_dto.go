/*
 * AVACloud API 1.16.0
 *
 * AVACloud API specification
 *
 * API version: 1.16.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package avaclient

// AwardTypeDto This enumeration describes the type of the award / procurement process. If this is used in a GAEB context, more information about award types can be found in the German VOB/A rules and the GAEB standard
type AwardTypeDto string

// List of AwardTypeDto
const (
	AWARDTYPEDTO_UNSPECIFIED                                         AwardTypeDto = "Unspecified"
	AWARDTYPEDTO_OPEN_PROCUREMENT                                    AwardTypeDto = "OpenProcurement"
	AWARDTYPEDTO_CLOSED_PRODUCREMENT                                 AwardTypeDto = "ClosedProducrement"
	AWARDTYPEDTO_NEGOTIATION_PROCUREMENT_WITHOUT_PUBLIC_ANNOUNCEMENT AwardTypeDto = "NegotiationProcurementWithoutPublicAnnouncement"
	AWARDTYPEDTO_NEGOTIATION_PROCUREMENT                             AwardTypeDto = "NegotiationProcurement"
	AWARDTYPEDTO_OPEN_CALL                                           AwardTypeDto = "OpenCall"
	AWARDTYPEDTO_SELECTED_CALL_WITHOUT_PUBLIC_COMPETITION            AwardTypeDto = "SelectedCallWithoutPublicCompetition"
	AWARDTYPEDTO_SELECTED_CALL                                       AwardTypeDto = "SelectedCall"
	AWARDTYPEDTO_DIRECT_AWARD                                        AwardTypeDto = "DirectAward"
	AWARDTYPEDTO_INTERNATIONAL_NATO_PROCUREMENT                      AwardTypeDto = "InternationalNATOProcurement"
	AWARDTYPEDTO_COMPETITIVE_DIALOG                                  AwardTypeDto = "CompetitiveDialog"
)