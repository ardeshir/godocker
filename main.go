package main

import (
	"fmt"
//	"io/ioutil"
	"os"
	"os/exec"
//	"path/filepath"
//	"strconv"
	"syscall"
)

// go run main.go run <cmd> <args>
func main() {

   switch os.Args[1] {
   case "run":
    run()
   case "child":
    child()
   default:
   panic("help")
   }

}

func run() {
   fmt.Printf("Running %v\n", os.Args[2:])

   cmd       := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
   cmd.Stdin  = os.Stdin
   cmd.Stdout = os.Stdout
   cmd.Stderr = os.Stderr
   // adding name spaces
   cmd.SysProcAttr = &syscall.SysProcAttr {
     Cloneflags: syscall.CLONE_NEWUTS,
   }


   must(cmd.Run())


}

func child() {
   fmt.Printf("Running %v\n", os.Args[2:])

   cmd       := exec.Command(os.Args[2], os.Args[3:]...)
   cmd.Stdin  = os.Stdin
   cmd.Stdout = os.Stdout
   cmd.Stderr = os.Stderr
   
   must(syscall.Sethostname([]byte("incontainer")))
   must(cmd.Run())


}

func must(err error) {
  if err != nil {
	panic(err)
  }
}
