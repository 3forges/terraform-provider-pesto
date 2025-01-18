/// I found a direct test for the utils func:
// TEST IT ON: https://go.dev/play/
// You can edit this code!
// Click here and start typing.

// / package main
package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ////////
// / inspired by:
// /  > MapAttribute: https://github.com/bpg/terraform-provider-proxmox/blob/6f657892c0a29d6677ef6d72690dbfb991a67ad1/fwprovider/ha/resource_hagroup.go#L88
// /  > utility function: https://github.com/bpg/terraform-provider-proxmox/blob/6f657892c0a29d6677ef6d72690dbfb991a67ad1/fwprovider/ha/resource_hagroup.go#L325
// /  > Create method: https://github.com/bpg/terraform-provider-proxmox/blob/6f657892c0a29d6677ef6d72690dbfb991a67ad1/fwprovider/ha/resource_hagroup.go#L160
// bakeFrontmatterDefFieldsToStrTsInterface converts the map of frontmatter_definition fields into a string, which is a TypeScript Interface.

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

	// return strings.Join(fmFieldsArray, ",")
	tsInterfaceFields := strings.Join(fmFieldsArray, ",")
	return fmt.Sprintf(`export interface `+contentTypeName+`_frontmatter_def { 
%s }`, tsInterfaceFields)
}

func main() {
	fmt.Println("Hello, 世界")

	/// frontmatter_definition := map[string]string{"number_of_doors": "number", "weight_in_kilograms": "number", "horsepower": "number", "max_speed": "number", "zero_sixty_time": "number", "color": "string", "trademark": "string", "second_hand": "boolean", "year": "string"}
	// elements := map[string]attr.Value{
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

// and it works
