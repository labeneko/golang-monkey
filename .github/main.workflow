workflow "Test" {
  on = "push"
  resolves = ["test"]
}

action "test" {
  uses = "docker://golang:latest"
  runs = "go version"
}
