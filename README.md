# Serabi - Struct Validator
Validation for struct, with 3 type of validation based on how the validation works:
- Basic validation, checks the value of the field accroding to the given rule
- Field comparison validattion, checks the value of the field according to the given rule relative to other field
- Regex validation, checks the value of the field accroding to the given rule simply using regex

The code is made as simple as possible for the beginner to understand

## Usage
serabi.Validate(data)

Where data is an instance of struct

## Notes
We included only very basic and common validations. Some validations are separated for the user to use as an additional. 