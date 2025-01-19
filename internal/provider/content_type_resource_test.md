
* One test I ran to check impl of the function:

```Golang
// / I found a direct test for the utils func:
// TEST IT ON: https://go.dev/play/
// You can edit this code!
// Click here and start typing.

/**/

package main

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func bakeFrontmatterDefFieldsToStrTsInterface(frontmatter_definition types.Map, contentTypeName string) string {
	fmFields := frontmatter_definition.Elements()
	fmFieldsArray := make([]string, len(fmFields))
	i := 0

	for name, value := range fmFields {
		if value.IsNull() {
			fmFieldsArray[i] = name
		} else {
			fmFieldsArray[i] = fmt.Sprintf("\n %s : %s ", name, value.(types.String).ValueString())
		}

		i++
	}
	tsInterfaceFields := strings.Join(fmFieldsArray, ",")
	return fmt.Sprintf(`export interface `+contentTypeName+`_frontmatter_def {
%s }`, tsInterfaceFields)
}

func main() {
	fmt.Println("Hello, 世界")
	elements := map[string]attr.Value{
		"key1":                types.StringValue("value1"),
		"key2":                types.StringValue("value2"),
		"number_of_doors":     types.StringValue("number"),
		"weight_in_kilograms": types.StringValue("number"),
		"horsepower":          types.StringValue("number"),
		"max_speed":           types.StringValue("number"),
		"zero_sixty_time":     types.StringValue("number"),
		"color":               types.StringValue("string"),
		"trademark":           types.StringValue("string"),
		"second_hand":         types.StringValue("boolean"), "year": types.StringValue("string")}
	frontmatter_definition, diags := types.MapValue(types.StringType, elements)
	result := bakeFrontmatterDefFieldsToStrTsInterface(frontmatter_definition, "sayMyName")

	fmt.Println(diags)
	fmt.Println(result)
}

```
