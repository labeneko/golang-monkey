workflow "Test" {
  on = "push"
  resolves = ["test"]
}

action "test" {
  uses = "./.github/actions/golang"
}
