package main

import "flag"

var tokenFlag = flag.String("token", "", "Creates Token File For DO API Connection.")
var createFlag = flag.String("create", "", "Create a Droplet")

func init() {
        flag.StringVar(createFlag, "s", "", "Add")
}

func main() {
        flag.Parse()
}
