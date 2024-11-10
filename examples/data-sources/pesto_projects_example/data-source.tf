data "pesto_projects" "example" {
  depends_on = [
    pesto_project.godzilla_project,
    pesto_project.mothra_project,
    pesto_project.gidhora_project,
  ]
}