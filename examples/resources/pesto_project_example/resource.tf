resource "pesto_project" "godzilla_project" {
  name                 = "godzillaRulesDemo"
  description          = "first example project to test creating a project with OpenTOFU the terraformation king. It also has been updated using the OpenTOFU King. A third test updating a project, to test if [stringplanmodifier.UseStateForUnknown()] works."
  git_service_provider = "giteaJapan"
  git_ssh_uri          = "git@github.com:3forges/godzillaRulesDemo.git"
}

