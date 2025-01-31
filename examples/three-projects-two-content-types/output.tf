/**

output "pesto_projects" {
  description = "Outputs all the Pesto Projects data."
  value       = data.pesto_projects.all_three_pesto_projects
}

*/
output "pesto_projects" {
  description = "Outputs all the Pesto Projects data."
  value       = module.monarch.pesto_projects // data.pesto_projects.all_three_pesto_projects
}