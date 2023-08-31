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

`import "github.com/karincake/serabi"`

Call the function based on the needs

```
myData := myStruct{}
err := serabi.Validate(data)
if err != nil {
    // do something with err
}
```

Usage of the Validation Tag

In case the data is empty and want to be filled from some sources, there are helper function than can be used with the task.
- `serabi.ValidateFormData(&any, *http.Request)` to fill data with content of HTTP Form Data
- `serabi.ValidateURL(&any, url.URL)` to fill data with content of URL
- `serabi.ValidateURL(&any, io.Reader)` to fill data with content of IoReader (with content of JSON format)

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

The Cryptography Regex (cryptographyregex, needs to import manually)
|Code|Description|
|---|---|
|md4||
|md5||
|sha256||
|sha384||
|sha512||
|ripemd128||
|ripemd160||
|tiger128||
|tiger160||
|tiger192||

The Encoding (encodingregex, needs to import manually)
|---|---|
|base64|Base64 String|
|base64URL|Base64URL String|
|base64RawURL|Base64RawURL String|
|url|URL String|
|html|HTML Encoded|

The Identifier Regex (identifierregex, needs to import manually)
|---|---|
|uuid|UUID format|
|uuid3|UUID v3 format|
|uuid4|UUID v4 format|
|uuid5|UUID v5 format|
|uuidRfc4122|UUID RFC4122|
|uuid3Rfc4122|UUID v3 RFC4122|
|uuid4Rfc4122|UUID v4 RFC4122|
|uuid5Rfc4122|UUID v5 RFC4122|
