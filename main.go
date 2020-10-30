package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func tree(root string) error {
	//fmt.Println(path);
	err := filepath.Walk(root, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fi.Name()[0] == '.' {
			return filepath.SkipDir
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			log.Fatal(err.Error())
		}
		depth := len(strings.Split(rel, string(filepath.Separator)))

		fmt.Printf("%s%s\n", strings.Repeat("  ", depth), fi.Name())
		return nil
	})
	return err
}

func trees(root, indent string) error {
	fi, err := os.Stat(root)
	if err != nil {
		return fmt.Errorf("could not stat %s: %v", root, err)
	}

	fmt.Println(fi.Name())
	if !fi.IsDir() {
		return nil
	}

	fis, err := ioutil.ReadDir(root)
	if err != nil {
		return fmt.Errorf("could not read dir %s: %v", root, err)
	}

	var names []string
	for _, fi := range fis {
		if fi.Name()[0] != '.' {
			names = append(names, fi.Name())
		}
	}

	for i, name := range names {
		add := "│  "
		if i == len(names)-1 {
			fmt.Printf(indent + "└──")
			add = "   "
		} else {
			fmt.Printf(indent + "├──")
		}

		if err := trees(filepath.Join(root, name), indent+add); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	// Taking input as the argument in the commanline and then storing
	// in an array , so that we can loop through it
	args := []string{"."}
	//If argument is passed only then check otherwise fuck off
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}
	//fmt.Print(args);

	// I love the looping statement of go thus looping through all
	//Arguments in the array
	for _, arg := range args {
		//fmt.Println(arg)
		err := trees(arg, "")
		if err != nil {
			log.Printf("Tree %s: %v\n", arg, err)
		}
	}
}
