package generate

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"strings"
	"sync"
	"text/template"

	"gorm.io/gorm/schema"
)

type Column struct {
	Name       string
	SqlName    string
	Types      string
	ImportPath string
}

type Param struct {
	StructName string
	TableName  string
	ImportMap  map[string]struct{}
	Cols       []Column
}

var importList = []string{"github.com/Ranxy/gormfields/query", "gorm.io/gorm"}

type Generate struct {
	PackageName  string
	OutPath      string
	UseZeroCheck bool
	Param        Param
}

func (g *Generate) Gen(val interface{}) {
	g.parserStruct(val)
	g.generateFromParam()
}

func (g *Generate) generateFromParam() {
	for _, pkg := range importList {
		g.Param.ImportMap[pkg] = struct{}{}
	}

	tmp := `// Code generated by gorm_column_gen. DO NOT EDIT.
package {{ .PackageName}}
	
import (
	{{range $key, $val := .Param.ImportMap}}"{{$key}}"
	{{end}}
) 
{{$head := .}}

{{range .Param.Cols}}
func {{ .Name }}({{ .Name }} {{ .Types}}, opts ...query.WithOpt) query.GormQueryReq {
	res := h{{ .Name }}{
		{{ .Name }}: {{ .Name }},
	}
	for _, opt := range opts {
		opt(&res.opt)
	}
	return res

}

type h{{ .Name }} struct {
	{{ .Name }}  {{ .Types}}
	opt query.Opt
}

func (i h{{ .Name }}) Do(db *gorm.DB) *gorm.DB {
	{{if $head.UseZeroCheck}}
	var zero {{ .Types}}
	if i.opt.CheckZero.CheckZero() && i.{{ .Name }} == zero {
		return db
	}
	{{end}}
	return i.opt.Or.Do(db)("{{ .SqlName}} = ?", i.{{ .Name }})
}
{{end}}

	`

	res := bytes.NewBuffer(make([]byte, 0))

	temp, err := template.New("query.template").Parse(tmp)
	if err != nil {
		panic(err)
	}

	err = temp.Execute(res, g)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(buildTypeGormFileName(g.OutPath, g.Param.StructName))
	if err != nil {
		panic(err)
	}

	f.Write(res.Bytes())
	f.Close()
}
func buildTypeGormFileName(outPath string, structName string) string {
	return outPath + strings.ToLower(structName) + "_gormfields.go"
}

func (g *Generate) parserStruct(val interface{}) {

	g.Param = Param{ImportMap: map[string]struct{}{}}

	var cacheStore = sync.Map{}
	schemaVal, err := schema.Parse(val, &cacheStore, &schema.NamingStrategy{})
	if err != nil {
		panic(err)
	}

	fmt.Println(schemaVal.Name)
	g.Param.StructName = schemaVal.Name
	g.Param.TableName = schemaVal.Table

	for Name, fields := range schemaVal.FieldsByName {
		//simple type
		if fields.FieldType.Kind() <= reflect.Complex128 {
			g.Param.Cols = append(g.Param.Cols, Column{
				Name:    Name,
				SqlName: fields.DBName,
				Types:   fields.FieldType.Kind().String(),
			})

		} else {
			g.Param.Cols = append(g.Param.Cols, Column{
				Name:       Name,
				SqlName:    fields.DBName,
				Types:      fields.FieldType.String(),
				ImportPath: fields.FieldType.PkgPath(),
			})
			if fields.FieldType.PkgPath() != "" {
				g.Param.ImportMap[fields.FieldType.PkgPath()] = struct{}{}
			}

		}

	}
}
