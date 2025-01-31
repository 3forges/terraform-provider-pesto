data "pesto_projects" "all_three_pesto_projects" {
  depends_on = [
    pesto_project.godzilla_project,
    pesto_project.mothra_project,
    pesto_project.gidhora_project,
  ]
}
