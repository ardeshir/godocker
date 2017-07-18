package main

import (
  "os"
  "os/exec"
)

// go run main.go run <cmd> <args>
func main() {

}

func run() {
   
   cmd       := exec.Command()
   cmd.Stdin  = os.Stdin
   cmd.Stdout = os.Stdout
   cmd.Stderr = os.Stderr

   must(cmd.Run())


}

func must(err error) {
  if err != nil {
	panic(err)
  }
}
