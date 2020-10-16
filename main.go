package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/ellemouton/genmake"
)

const envDir = "GENSCRATCHPATH"

var (
	cmdPath = flag.String("path", "", "scratch environment path")
	dirName = flag.String("name", "", "scratch environment directory name")
)

func getDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	if os.Getenv(envDir) != "" {
		dir = os.Getenv(envDir)
	}

	if *cmdPath != "" {
		dir = *cmdPath
	}

	if *dirName != "" {
		return dir + "/" + *dirName, nil
	}

	num, err := getNextFolderNum(dir, "gs_")
	if err != nil {
		return "", err
	}

	return dir + "/gs_" + strconv.Itoa(num), nil
}

func main() {
	flag.Parse()

	scratchDir, err := getDir()
	if err != nil {
		log.Fatal(err)
	}

	err = os.Mkdir(scratchDir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(scratchDir + "/main.cpp")
	if err != nil {
		log.Fatal(err)
	}

	t := template.New("scratch")
	_, err = t.Parse(mainTmpl)
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(f, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = genmake.Generate(scratchDir, "play")
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("vim", scratchDir+"/main.cpp")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Your playground directory:", scratchDir)
}

func getNextFolderNum(dir, pref string) (int, error) {
	max := 0

	err := filepath.Walk(dir, func(p string, f os.FileInfo, _ error) error {
		if f.IsDir() {
			s := strings.Split(p[len(dir):], "/")
			if len(s) == 2 && s[1] == f.Name() && f.Name()[:len(pref)] == pref {
				i, err := strconv.Atoi(s[1][len(pref):])
				if err != nil {
					return err
				}

				if i > max {
					max = i
				}
			}

		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return max + 1, nil
}
