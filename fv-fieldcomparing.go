// built-in field checker
package serabi

// register the field checkers
func init() {
	AddTagForField("gtField", gtTagValidator)
	AddTagForField("gteField", gteTagValidator)
	AddTagForField("ltField", ltTagValidator)
	AddTagForField("lteField", lteTagValidator)
	AddTagForField("lengthField", minLengthTagValidator)
}
