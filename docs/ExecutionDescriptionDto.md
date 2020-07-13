# ExecutionDescriptionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Elements GUID identifier. | 
**GaebXmlId** | **string** | This is used to store the GAEB XML Id within this IElement. This data is not used for any calculations or evaluations but only for GAEB serialization and deserialization. | [optional] 
**ElementTypeDiscriminator** | **string** |  | 
**Blocks** | [**[]NoteTextDto**](NoteTextDto.md) | Blocks within an ExecutionDescription contain the actual information. | [optional] 
**Label** | **string** | Labels this ExecutionDescription. | [optional] 
**Identifier** | **string** | Uniquely identifies this ExecutionDescription. | [optional] 
**ElementType** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


