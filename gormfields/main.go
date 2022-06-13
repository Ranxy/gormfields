// Content of this file was copied and modify from the package github.com/mailru/easyjson
// Under the MIT License licence:
// github.com/mailru/easyjson/blob/v0.7.7/LICENSE

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Ranxy/gormfields/bootstrap"
	"github.com/Ranxy/gormfields/parser"
)

var buildTags = flag.String("build_tags", "", "build tags to add to generated file")
var genBuildFlags = flag.String("gen_build_flags", "", "build flags when running the generator while bootstrapping")
var zeroCheck = flag.Bool("zero_check", false, "use zero check for query")
var allStructs = flag.Bool("all", false, "generate marshaler/unmarshalers for all structs in a file")
var specifiedName = flag.String("output_filename", "", "specify the filename of the output")
var processPkg = flag.Bool("pkg", false, "process the whole package instead of just the given file")

func generate(fname string) (err error) {
	fInfo, err := os.Stat(fname)
	if err != nil {
		return err
	}

	p := parser.Parser{AllStructs: *allStructs}
	if err := p.Parse(fname, fInfo.IsDir()); err != nil {
		return fmt.Errorf("Error parsing %v: %v", fname, err)
	}

	fmt.Println("PARSER VALUE ", p)

	p.PkgName += "_fields"

	var outPath string
	outPath = filepath.Join(fname, p.PkgName)
	tempPath := filepath.Join(fname, p.PkgName+"_temp")

	var success bool

	outExist, err := PathExist(outPath)
	if err != nil {
		return err
	}

	if outExist {
		err = os.Rename(outPath, tempPath)
		if err != nil {
			return err
		}

		defer func() {
			if success {
				err := os.RemoveAll(tempPath)
				if err != nil {
					panic(err)
				}
			} else {
				err = os.Rename(tempPath, outPath)
				if err != nil {
					panic(err)
				}

			}
		}()
	}

	var trimmedBuildTags string
	if *buildTags != "" {
		trimmedBuildTags = strings.TrimSpace(*buildTags)
	}

	var trimmedGenBuildFlags string
	if *genBuildFlags != "" {
		trimmedGenBuildFlags = strings.TrimSpace(*genBuildFlags)
	}

	g := bootstrap.Generator{
		BuildTags:     trimmedBuildTags,
		GenBuildFlags: trimmedGenBuildFlags,
		PkgPath:       p.PkgPath,
		PkgName:       p.PkgName,
		Types:         p.StructNames,
		OutPath:       outPath,
		UseZeroCheck:  *zeroCheck,
	}

	if err := g.Run(); err != nil {
		return fmt.Errorf("Bootstrap failed: %v", err)
	}
	success = true
	return nil
}
func main() {
	flag.Parse()

	files := flag.Args()

	gofile := os.Getenv("GOFILE")
	if *processPkg {
		gofile = filepath.Dir(gofile)
	}

	if len(files) == 0 && gofile != "" {
		files = []string{gofile}
	} else if len(files) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	for _, fname := range files {
		if err := generate(fname); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
