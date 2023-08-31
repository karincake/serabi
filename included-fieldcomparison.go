package serabi

//  included default tag validation for field comparison validation

// register the field checkers
func init() {
	AddTagForField("eqField", eqTagValidator)
	AddTagForField("gtField", gtTagValidator)
	AddTagForField("gteField", gteTagValidator)
	AddTagForField("ltField", ltTagValidator)
	AddTagForField("lteField", lteTagValidator)
}
