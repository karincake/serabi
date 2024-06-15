# Testing

## Data Type
There are several structs available for testing. Struct name format is used to indicate their characteristics with following format:

`Data(Size)(Complexity)[OptionalInfo]`

Size is for how many fields the struct has with the following categories:
- Small for 1 field
- Medium for 3-5 fields
- Large for 7-10 fields, 5 primary fields and 3-5 sub fields from other structs
- Huge for 13-15 fields, 5 primary fields and 7-10 sub fields from other structs

 Complexity is for how many tags are being used with the following categories:
 - Simple for 1 field
 - Complex for more than 1 field

 OptionalInfo is for any info needed to differs one struct from another.

## Testing Mode
There are two modes of testing: successful and failing tests for each data type