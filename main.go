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
  fmt.Printf("Parent run %v\n", os.Args[2:])

   cmd       := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
   cmd.Stdin  = os.Stdin
   cmd.Stdout = os.Stdout
   cmd.Stderr = os.Stderr
   // adding name spaces
   cmd.SysProcAttr = &syscall.SysProcAttr {
     Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
   }


   must(cmd.Run())


}

func child() {
   fmt.Printf("Child %v\n", os.Args[2:])

   cmd       := exec.Command(os.Args[2], os.Args[3:]...)
   cmd.Stdin  = os.Stdin
   cmd.Stdout = os.Stdout
   cmd.Stderr = os.Stderr
   
   must(syscall.Sethostname([]byte("incontainer")))
   // must(syscall.Chroot("/home/ec2-user/newfs"))
   // must(os.Chdir("/"))
   // must(syscall.Mount("proc","proc","proc", 0, ""))

   must(cmd.Run())
   // clean up 
   // must(syscall.Unmount("proc",0))

}

func must(err error) {
  if err != nil {
	panic(err)
  }
}
