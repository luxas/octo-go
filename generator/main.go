package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/willabides/octo-go/generator/internal/model/openapi"
)

func main() {
	var schemaPath string
	var outputPath string
	var pkgPath string
	var pkgName string
	flag.StringVar(&schemaPath, "schema", "", "path to openapi schema")
	flag.StringVar(&outputPath, "out", "", "directory to write all these files")
	flag.StringVar(&pkgPath, "pkgpath", "", "path for output package")
	flag.StringVar(&pkgName, "pkg", "", "name for output package")
	flag.Parse()
	err := run(schemaPath, outputPath, pkgPath, pkgName)
	if err != nil {
		log.Fatal(err)
	}
}

func run(schemaPath, outputPath, pkgPath, pkgName string) error {
	schemaFile, err := os.Open(schemaPath)
	if err != nil {
		return err
	}
	mdl, err := openapi.Openapi2Model(schemaFile)
	if err != nil {
		return err
	}
	endpoints := mdl.Endpoints

	sort.Slice(endpoints, func(i, j int) bool {
		return endpoints[i].ID < endpoints[j].ID
	})

	concernFiles := map[string]*jen.File{}

	for _, endpoint := range endpoints {
		if endpoint.Legacy {
			continue
		}
		if concernFiles[endpoint.Concern] == nil {
			cf := jen.NewFilePathName(pkgPath, pkgName)
			cf.HeaderComment("Code generated by octo-go; DO NOT EDIT.")
			concernFiles[endpoint.Concern] = cf
		}
		file := concernFiles[endpoint.Concern]
		addClientMethod(file, endpoint)
		addRequestStruct(file, endpoint)
		addRequestBody(file, endpoint)
		addResponseBody(file, endpoint)
		addResponse(file, endpoint)
	}

	for concern, concernFile := range concernFiles {
		name := fmt.Sprintf("zz_%s_gen.go", strings.ReplaceAll(concern, "-", "_"))
		var f *os.File
		f, err = os.Create(filepath.Join(outputPath, name))
		if err != nil {
			return err
		}
		err = concernFile.Render(f)
		if err != nil {
			return err
		}
	}

	componentSchemas := mdl.ComponentSchemas
	if len(componentSchemas) > 0 {
		filename := "schemas_gen.go"
		var f *os.File
		f, err = os.Create(filepath.Join(outputPath, "components", filename))
		if err != nil {
			return err
		}
		cf := jen.NewFilePath(path.Join(pkgPath, "components"))
		cf.HeaderComment("Code generated by octo-go; DO NOT EDIT.")
		addComponentSchemas(cf, componentSchemas)
		err = cf.Render(f)
		if err != nil {
			return err
		}
	}

	endpointTypesOut, err := os.Create(filepath.Join(outputPath, "zz_endpoint_types_gen.go"))
	if err != nil {
		return err
	}
	endpointTypesFile := jen.NewFilePathName(pkgPath, pkgName)
	endpointTypesFile.HeaderComment("Code generated by octo-go; DO NOT EDIT.")
	addEndpointTypes(endpointTypesFile)
	err = endpointTypesFile.Render(endpointTypesOut)
	if err != nil {
		return err
	}
	return nil
}
