# Serabi - Struct Validator
Validation for struct, with 3 type of validation based on how the validation works:
- Basic validation, checks the value of the field accroding to the given rule
- Field comparison validattion, checks the value of the field according to the given rule relative to other field
- Regex validation, checks the value of the field accroding to the given rule simply using regex

The code is made as simple as possible for the beginner to understand. A more complex validator can be found here
https://github.com/go-playground/validator 

## Installation and Usage
Just use go get command

`go get github.com/karincake/serabi`

Import in the package

```go
import "github.com/karincake/serabi"
```

Call the function based on the needs

```go
myData := myStruct{}
err := serabi.Validate(data)
if err != nil {
    // do something with err
}
```

## Available Functions
### Main Function
- `Validate(input any)` to fill validate a struct.

### Wrapper Function 
The following functions are to help filling a struct (see semprit, https://github.com/karincake/semprit) before validating them:
- `ValidateFormData(container any, *http.Request)` to fill data with content of HTTP Form Data and then validate the struct
- `ValidateURL(container any, url.URL)` to fill data with content of URL and then validate the struct
- `ValidateIoReader(container any, io.Reader)` to fill data with content of IoReader (with content of JSON format) and then validate the struct

Due to its function to fill an object, make sure to use pointer of object for the first parameter.

### Helper Function
- `AddTag(tag string, f FvFunc)`, to register a field-validator
- `AddTagForField(tag string, f FvFunc)`, to register a field-validator that is comparing with another field
- `AddTagForRegex(tag string, regexString string, message string)`, to register field-validator that uses regex string
- `RemoveTag(tag string)`, to remove a field-validator

## Types
- `FvFunc func(reflect.Value, string) error`, function type that defines a validation. First parameter is for field value, second parameter is for validation rule

## Tag Validation Format
How to write validation tag:
- Tag name is `validate`.
- The value can have multiple validation rules, separated by semicolon (;).
- A validation rule is a key-value pair separated by equals sign (=)
- A field comparison validator uses rule's value to determine which field to compare

Examples:
- `validate="required"`, one rule
- `validate="required;minLength=10"`, multiple rules
- `validate="required;eqField=Password"`, multiple rules, with one of it is a field comparison validation

## Available Validation
The Basic Validation (included)
|Code|Description|
|---|---|
|required|required|
|gt=x|greater than x|
|gte=x|greater than or equal to x|
|lt=x|less than x|
|lte=x|less than or equal to x|
|length=x|length is x|
|minLength=x|minimum length is x|
|maxLength=x|maximum length is x|

The Field Comparison (included)
|Code|Description|
|---|---|
|eqField=x|equal to field x|
|gtField=x|greater than field x|
|gteField=x|greater than or equal to field x|
|ltField=x|less than field x|
|lteField=x|less than or equal to x|

The Regex (included)
|Code|Description|
|---|---|
|alpha|Alphabet characters|
|alphaSpace|Alphabet characters with space within it|
|alphaNumeric|Alphabet and numeric characters|
|alphaUnder|Alphabet characters and underscore|
|alphaNumericUnder|Alphabet characters, numeric characters, and underscore|
|numeric|Numeric characters|
|numval|Number value|
|email|Valid email format|

The main package includes only very basic and common validations. Some validations are separated for the user to use as an additional by importing the side effect manually, i.e

`import _ github.com/karincake/serabi/encodingregex`

The Cryptography Regex (cryptographyregex, needs to import the side effect manually)
|Code|Description|
|---|---|
|md4|MD4 hash|
|md5|MD5 hash|
|sha256|SHA256 hash|
|sha384|SHA384 hash|
|sha512|SHA512 hash|
|ripemd128|RIPEMD-128 hash|
|ripemd160|RIPEMD-160 hash|
|tiger128|TIGER128 hash|
|tiger160|TIGER160 hash|
|tiger192|TIGER192 hash|

The Encoding (encodingregex, needs to import the side effect manually)
|Code|Description|
|---|---|
|base64|Base64 String|
|base64URL|Base64URL String|
|base64RawURL|Base64RawURL String|
|url|URL String|
|html|HTML Encoded|

The Identifier Regex (identifierregex, needs to import the side effectmanually)
|Code|Description|
|---|---|
|uuid|UUID format|
|uuid3|UUID v3 format|
|uuid4|UUID v4 format|
|uuid5|UUID v5 format|
|uuidRfc4122|UUID RFC4122|
|uuid3Rfc4122|UUID v3 RFC4122|
|uuid4Rfc4122|UUID v4 RFC4122|
|uuid5Rfc4122|UUID v5 RFC4122|

## Options
- `CacheEnabled bool`, to cache the field identification results. The default value is `false`. When set to `true`, the feature can boost the performance up to 2-3 times faster.
- `CacheMaxCount int`, to limit the number of strcuts that can be cached. The default value is `5`.