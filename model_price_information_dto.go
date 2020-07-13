/*
 * AVACloud API 1.16.0
 *
 * AVACloud API specification
 *
 * API version: 1.16.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package avaclient

// PriceInformationDto Holds global price information for a ServiceSpecification
type PriceInformationDto struct {
	// Elements GUID identifier.
	Id string `json:"id"`
	// The amount of currency per one hour of manual labour work in this project.
	HourlyWage float32 `json:"hourlyWage"`
	// The final, total price will be deducted by this rate.
	DeductionFactor float32 `json:"deductionFactor"`
	// This is given when there is only one flat price for the whole service specification.
	FlatSum float32 `json:"flatSum"`
	// Global tax rate for the project. Note that certain elements may have a different, specific tax rate.
	TaxRate float32 `json:"taxRate"`
	// Trade discounts for offered in this ServiceSpecification.
	TradeDiscounts []TradeDiscountDto `json:"tradeDiscounts,omitempty"`
}
