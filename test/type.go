package test

// data(size)(complexity)[optional]
// size is by field count
// complexity is by tag count
// opional, is any additional info

type DataSmallSimpleString struct {
	Name string `validate:"required"`
}

type DataSmallSimpleNumber struct {
	Age uint8 `validate:"gt=15"`
}

type DataSmallSimpleBoolean struct {
	Married bool `validate:"required"`
}

type DataSmallComplexString struct {
	Name string `validate:"required;alphaSpace"`
}

type DataSmallComplexArrayofString struct {
	Data []string `validate:"required;minLength=5"`
}

type DataSmallComplexNumber struct {
	Age uint8 `validate:"required;gt=5;lt=15"`
	Exp uint8 `validate:"required;gtField=Age"`
}

type DataMediumSimple struct {
	Name    string `validate:"alphaSpace"`
	Address string `validate:"minLength=10"`
	Age     uint8  `validate:"gt=15"`
	Email   string `validate:"required"`
	Phone   string `validate:"numeric"`
}

type DataMediumSimpleArrayofStruct struct {
	Data []DataMediumSimple `validate:"required"`
}

type DataMediumSimpleField struct {
	Name        string `validate:"alpha"`
	Password    string `validate:"minLength=8"`
	RePassword  string `validate:"eqField=Password"`
	Age         int    `validate:"gt=10"`
	YearsActive int    `validate:"ltField=Age"`
}

type DataMediumComplex struct {
	Name    string `validate:"required;alphaSpace;minLength=10;maxlength=50"`
	Address string `validate:"required;minLength=10;maxlength=50"`
	Age     uint8  `validate:"required;numeric;gt=15;lt=40"`
	Email   string `validate:"required;email;minLength=10;maxlength=100"`
	Phone   string `validate:"required;numeric;minLength=10;maxlength=20"`
}

type DataLargeSimple struct {
	Name      string `validate:"alphaSpace"`
	Address   string `validate:"minlength=10"`
	Age       uint8  `validate:"gt=15"`
	Email     string `validate:"required"`
	Phone     string `validate:"numeric"`
	workStats workStatsSimple
}

type DataLargeComplex struct {
	Name      string `validate:"required;alphaSpace"`
	Address   string `validate:"required;minLength=10;maxlength=50"`
	Age       uint8  `validate:"required;numeric;gt=15;lt=40"`
	Email     string `validate:"required"`
	Phone     string `validate:"required;numeric;minLength=10;maxlength=20"`
	workStats workStatsComplex
}

type DataHugeSimple struct {
	Name      string `validate:"alphaSpace"`
	Address   string `validate:"minlength=10"`
	Age       uint8  `validate:"gt=5"`
	Email     string `validate:"required"`
	Phone     string `validate:"numeric"`
	workStats workStatsSimple
	GameStats gameStatsSimple
}

type DataHugeComplex struct {
	Name      string `validate:"required;alphaSpace"`
	Address   string `validate:"required;minLength=10;maxlength=50"`
	Age       uint8  `validate:"required;numeric;gt=15;lt=40"`
	Email     string `validate:"required;email;minLength=10;maxlength=100"`
	Phone     string `validate:"required;numeric;minLength=10;maxlength=20"`
	WorkStats workStatsComplex
	GameStats gameStatsComplex
}

type workStatsSimple struct {
	Office      string  `validate:"required"`
	Rate        float32 `validate:"gt=-10"`
	Expertise   string  `validate:"minLength=10"`
	YearsActive uint8   `validate:"gt=2"`
	NetWorth    float64 `validate:"gt=10000"`
}

type workStatsComplex struct {
	Office      string  `validate:"required;minLength=10;maxlength=50"`
	Rate        float32 `validate:"required;gt=-10"`
	Expertise   string  `validate:"required;minLength=10;maxlength=50"`
	YearsActive uint8   `validate:"required;gt=2;"`
	NetWorth    float64 `validate:"required;gt=10000"`
}

type gameStatsSimple struct {
	CharName    string  `validate:"minLength=5"`
	Score       int     `validate:"gt=-10"`
	Speciality  string  `validate:"minLength=10"`
	YearsActive uint8   `validate:"gt=2"`
	NetWorth    float64 `validate:"gt=10000"`
}

type gameStatsComplex struct {
	CharName    string  `validate:"required;minLength=5;maxlength=5;alphaNumericUnder"`
	Score       int     `validate:"required;gt=-10"`
	Speciality  string  `validate:"required;minLength=10;maxlength=50"`
	YearsActive uint8   `validate:"required;gt=2;"`
	NetWorth    float64 `validate:"required;gt=10000"`
}
