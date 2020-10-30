package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func tree(root string) error {
  //fmt.Println(path);
  err:=filepath.Walk(root,func(path string, fi os.FileInfo, err error) error {
    if(err!=nil) {
      return err;
    }
    if fi.Name()[0]=='.'{
      return filepath.SkipDir
    }
      fmt.Println(fi.Name());
      return nil;
  })
  return err;
}

func main(){
  // Taking input as the argument in the commanline and then storing 
  // in an array , so that we can loop through it 
  args:= []string{"."}
  //If argument is passed only then check otherwise fuck off
  if len(os.Args) > 1 {
    args = os.Args[1:]
  }
  //fmt.Print(args);

  // I love the looping statement of go thus looping through all 
  //Arguments in the array 
  for _,arg := range args{
    //fmt.Println(arg)
    err := tree(arg);
    if err!= nil {
      log.Printf("Tree %s: %v\n",arg,err);
    }
  }
}
