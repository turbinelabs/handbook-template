package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func main() {
	var (
		overridesDir = flag.String("overrides-dir", "overrides", "directory to pull override files from")
		outDir = flag.String("out", "out", "where to write the handbook")
		templateDir = flag.String("template-dir", "templates", "where to read template files from")
		varsFile = flag.String("vars", "vars.json", "JSON file to read vars from")
	)

	flag.Parse()

	absOverridesDir, err := filepath.Abs(*overridesDir)
	if err != nil {
		log.Fatal(err)
	}
	absOutDir, err := filepath.Abs(*outDir)
	if err != nil {
		log.Fatal(err)
	}	
	absTemplateDir, err := filepath.Abs(*templateDir)
	if err != nil {
		log.Fatal(err)
	}
	absVarsPath, err := filepath.Abs(*varsFile)
	if err != nil {
		log.Fatal(err)
	}
	
	log.Printf("reading templates from %s", absTemplateDir)
	log.Printf("reading overrides from %s", absOverridesDir)
	log.Printf("reading variables from %s", absVarsPath)
	log.Printf("writing a new handbook to %s", absOutDir)

	data := readVars(absVarsPath)
	err = filepath.Walk(absTemplateDir, generate(absTemplateDir, absOverridesDir, absOutDir, data))
	if err != nil {
		log.Fatal(err)
	}
}

/**
* read vars into an interface{}
* this will end up being a map[string]string, which works just fine in go templates
*/
func readVars(varsPath string) interface{} {
	file, err := ioutil.ReadFile(varsPath)
	if err != nil {
		log.Fatal(err)
	}
	var data interface{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

/**
 * generate a callback function that picks a template for each file, processes it and writes it to the output directory
 */
func generate(templatesDir, overridesDir string, outDir string, data interface{}) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		relPath, err := filepath.Rel(templatesDir, path)
		if err != nil {
			return err
		}
		if strings.Contains(path, "/.") {
			log.Printf("skipping file %s", path)
			return nil
		} else {
			log.Printf("processing %s", relPath)
			if info.IsDir() {
				err = os.Mkdir(filepath.Join(outDir, relPath), 0770)
				if !os.IsExist(err) {
					return err
				}
			} else {
				
				// check for overrides. If we can find a matching relative file
				// in overridesDir, use it as our template instead
				templatePath := filepath.Join(overridesDir, relPath)
				_, err := os.Open(templatePath)
				if os.IsNotExist(err) {
					templatePath = path
				} else if err != nil {
					return err
				} else {
					log.Printf("overriding template %s with replacement at %s", path, templatePath)
				}
				
				log.Printf("parsing template file %s", templatePath)
				tmpl, err := template.ParseFiles(templatePath)
				if err != nil {
					return err
				}
				absFilePath := filepath.Join(outDir, relPath)
				log.Printf("writing handbook page to %s", absFilePath)
				file, err := os.Create(absFilePath)
				if err != nil {
					return err
				}
				err = tmpl.Execute(file, data)
				if err != nil {
					return err
				}
				return file.Close()
			}
		}
		return nil
	}
}
	
