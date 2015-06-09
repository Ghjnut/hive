// A service that manages interactions between other services
package hive

import (
  "fmt"

  "github.com/yosida95/golang-jenkins"
)

func main() {
  fmt.Printf("Hello, world.\n")
  auth := &gojenkins.Auth{"Ghjnut", "Token123"}
  jenkins := gojenkins.NewJenkins(auth, "http://127.0.0.1")
  jobs, err := jenkins.GetJobs()

  fmt.Printf("%+v", jobs)

  if err != nil {
    fmt.Printf("error %v\n", err)
  }

  if len(jobs) == 1 {
    fmt.Printf("return no jobs\n")
  }
}
