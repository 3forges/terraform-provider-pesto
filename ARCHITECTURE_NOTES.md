# Architecture notes

## About complex types for attributes

### frontmatter_defintion to be a map?

One of the attributes of the `pesto_content_type` resource, is implemented:

* As a string,
* And its value must be a valid Typescript interface, complying with certain constraints: every typescript interaface attribute must be of one of the following Typescript language types: `string`, `boolean`, or `number`.

I want to improve the teraform provider, such that the terraform provider user, does not need anytypescript language knowledge.

Instead of writing the definition of a valid TypeScript interface, compliant with the above described constraints, I want the terraform provider user, to only specify a Map:

For example, instead of:

```Hcl

resource "pesto_content_type" "car_contenttype_with_tofu" {
  project_id = pesto_project.mothra_project.id // "${pesto_project.gidhora_project.id}"
  name       = "car"
  frontmatter_definition = <<EOF
export interface nameOftestCtxt1_Frontmatter {
  number_of_doors: number
  weight_in_kilograms: number
  horsepower: number
  max_speed: number
  zero_sixty_time: number
  color: string
  trademark: string
  second_hand: boolean
  year: string
}
EOF
  description            = "A pesto content type representing a car for sale, created by terraformation"

}
```

I would like that my terraform provider user writes:

```Hcl

resource "pesto_content_type" "car_contenttype_with_tofu" {
  project_id = pesto_project.mothra_project.id // "${pesto_project.gidhora_project.id}"
  name       = "car"
  frontmatter_definition = {
    "number_of_doors" = "number"
    "weight_in_kilograms" = "number"
    "horsepower" = "number"
    "max_speed" = "number"
    "zero_sixty_time" = "number"
    "color" = "string"
    "trademark" = "string"
    "second_hand" = "boolean"
    "year" = "string"
  }
  description            = "A pesto content type representing a car for sale, created by terraformation"

}
```

And the TypeScript interface which would be defined will automatically generated as:

```TypeScript
// export interface ${NAME_OF_THE_TF_RESOURCE}_Frontmatter {
export interface car_contenttype_with_tofu_Frontmatter {
  number_of_doors: number
  weight_in_kilograms: number
  horsepower: number
  max_speed: number
  zero_sixty_time: number
  color: string
  trademark: string
  second_hand: boolean
  year: string
}
```


* Question: how do I access the value of each element of the golang Map defined in the `*.tf` files?

* Question: How can I define terraform attribute value validation constraints?

## References

* https://developer.hashicorp.com/terraform/plugin/framework/handling-data/attributes/map
* https://developer.hashicorp.com/terraform/plugin/framework/handling-data/types/map#accessing-values
* https://developer.hashicorp.com/terraform/plugin/framework/handling-data/types/custom
* Examples of using a Map Attribute inside a terraform provider implementation: https://github.com/jianyuan/terraform-provider-sentry/blob/23fe95b74864d619108b48365451902d539f17d3/internal/provider/resource_issue_alert.go#L1006

<!--

donc déjà comment je peux définir des contraintes de validation de chaque valeur renseignée ?
plus important encore: ok comment j'accède aux valeurs fixées dans les fichiers *.tf , dans le code du provider ? réponse: https://developer.hashicorp.com/terraform/plugin/framework/handling-data/types/map#accessing-values 
est-ce que je peux définir un nouveau type du genre je veux que chaque élément de la Map prenne ces valeurs spécifiques (un peu comme une enum ? réponse oui, cf. https://developer.hashicorp.com/terraform/plugin/framework/handling-data/types/custom

-->
