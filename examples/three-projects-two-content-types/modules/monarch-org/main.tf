resource "pesto_project" "godzilla_project" {
  name                 = "godzillaRulesDemo"
  description          = "first example project to test creating a project with OpenTOFU the terraformation king. It also has been updated using the OpenTOFU King. A third test updating a project, to test if [stringplanmodifier.UseStateForUnknown()] works."
  git_service_provider = "giteaJapan"
  git_ssh_uri          = "git@github.com:3forges/godzillaRulesDemo.git"
}

resource "pesto_project" "mothra_project" {
  name                 = "mothraDemoCedric"
  description          = "second example project to test creating a project with OpenTOFU the terraformation king. It also has been updated using the OpenTOFU King. A fourth test updating a project, to test if [stringplanmodifier.UseStateForUnknown()] works."
  git_service_provider = "giteaJapan"
  git_ssh_uri          = "git@github.com:3forges/mothra.git"

}

resource "pesto_project" "gidhora_project" {
  name                 = "gidhoraDemoCedric"
  description          = "third example project to test creating a project with OpenTOFU the terraformation king"
  git_service_provider = "giteaJapan"
  git_ssh_uri          = "git@github.com:3forges/gidhora.git"
}

resource "pesto_content_type" "contenttype1_with_tofu" {
  project_id = pesto_project.gidhora_project.id // "${pesto_project.gidhora_project.id}"
  name       = "socks"
  // frontmatter_definition = "rubbish_Frontmatter_Def"
  // frontmatter_definition = "export interface nameOftestCtxt1_Frontmatter;"
  // frontmatter_definition = "export interface nameOftestCtxt1_Frontmatter {}"
  // frontmatter_definition = "export interface nameOftestCtxt1_Frontmatter { \n  addedByTerraformationByUpdatingFMdef: boolean \n  includeInJumbo: boolean \n  size: number \n  price: number \n  color: string \n  trademark: string \n  isFromNewClothesCollection: boolean \n}"
  frontmatter_definition = {
    "available_in_stock" : "boolean",
    "european_size" : "number",
    "american_size" : "number",
    "price" : "number",
    "color" : "string",
    "trademark" : "string",
    "isInLatestClothesCollection" : "boolean",
    "addedToTestUpdatingWithTerraformProvider" : "boolean"
  }
  description = "A pesto content type representing a pair of socks for sale, created by terraformation"

}

resource "pesto_content_type" "car_contenttype_with_tofu" {
  project_id = pesto_project.mothra_project.id // "${pesto_project.gidhora_project.id}"
  name       = "car"
  // frontmatter_definition = "rubbish_Frontmatter_Def"
  // frontmatter_definition = "export interface nameOftestCtxt1_Frontmatter;"
  // frontmatter_definition = "export interface nameOftestCtxt1_Frontmatter {}"
  // frontmatter_definition = "export interface nameOftestCtxt1_Frontmatter { \n  addedByTerraformationByUpdatingFMdef: boolean \n  includeInJumbo: boolean \n  size: number \n  price: number \n  color: string \n  trademark: string \n  isFromNewClothesCollection: boolean \n}"
  frontmatter_definition = {
    "number_of_doors" : "number"
    "weight_in_kilograms" : "number"
    "horsepower" : "number"
    "max_speed" : "number"
    "zero_sixty_time" : "number"
    "color" : "string"
    "trademark" : "string"
    "second_hand" : "boolean"
    "year" : "string"
  }
  description = "A pesto content type representing a car for sale, created by terraformation"

}

