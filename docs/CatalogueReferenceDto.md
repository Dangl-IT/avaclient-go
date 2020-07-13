# CatalogueReferenceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Elements GUID identifier. | 
**CatalogueReferenceKey** | **string** | This points to the item in the catalogue itself. This means that, for example when this quantity assignment references a \&quot;DIN 276\&quot; catalogue, this property indicates the number / identifier / key in DIN 276 that is referenced. | [optional] 
**CatalogueReferenceId** | **string** | The Id of the CatalogueReference that is targeted by this item. Set this property to set the referenced catalogue. | 
**ProjectCatalogues** | [**[]CatalogueDto**](CatalogueDto.md) | These are CatalogueReference that are used within this ServiceSpecification. Catalogue references are used to describe catalogues, or collections, that can be used to describe elements with commonly known properties. For example, QuantityAssignments use these to categorize themselves. | [optional] 
**Catalogue** | [**CatalogueDto**](CatalogueDto.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


