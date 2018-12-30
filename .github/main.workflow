workflow "Test" {
  on = "push"
  resolves = ["test"]
}

action "test" {
  uses = "./say_hello.sh"
  args = "\"Hello World\""
}
