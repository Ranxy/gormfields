package generate

import (
	"database/sql"
	"os"
	"testing"

	"github.com/Ranxy/gormfields/generate/example"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestGenerate(t *testing.T) {

	g := Generate{
		OutPath:      "./generate_fields",
		UseZeroCheck: true,
		PackageName:  "generate_fields",
		Param: Param{
			StructName: "TestP",
			TableName:  "Testps",
			Cols: []Column{
				{
					Name:    "UserID",
					SqlName: "user_id",
					Types:   "uint",
				},
				{
					Name:    "UserName",
					SqlName: "user_name",
					Types:   "string",
				},
			},
			ImportMap: map[string]struct{}{},
		},
	}

	t.Cleanup(func() {
		err := os.RemoveAll(g.OutPath)
		require.NoError(t, err)
	})

	err := os.Mkdir(g.OutPath, os.ModePerm)
	require.NoError(t, err)

	g.generateFromParam()
}

type StatusType uint
type Testp struct {
	gorm.Model
	UserName    string  `json:"user_name" gorm:"user_name"`
	UserDisplay *string `json:"user_display" gorm:"user_display"`
	UserDp      sql.NullString
	Stp         float64
	Status      StatusType
}

func Test_parserStruct(t *testing.T) {

	g := &Generate{}

	g.parserStruct(Testp{})

	require.Equal(t, "Testp", g.Param.StructName)
	require.Equal(t, "testps", g.Param.TableName)

	m := make(map[string]Column)
	for _, col := range g.Param.Cols {
		m[col.Name] = col
	}

	require.Equal(t, "UserName", m["UserName"].Name)
	require.Equal(t, "user_name", m["UserName"].SqlName)
	require.Equal(t, "string", m["UserName"].Types)
	require.Equal(t, "", m["UserName"].ImportPath)

	require.Equal(t, "UserDisplay", m["UserDisplay"].Name)
	require.Equal(t, "user_display", m["UserDisplay"].SqlName)
	require.Equal(t, "*string", m["UserDisplay"].Types)
	require.Equal(t, "", m["UserDisplay"].ImportPath)

	require.Equal(t, "UserDp", m["UserDp"].Name)
	require.Equal(t, "user_dp", m["UserDp"].SqlName)
	require.Equal(t, "sql.NullString", m["UserDp"].Types)
	require.Equal(t, "database/sql", m["UserDp"].ImportPath)

	require.Equal(t, "Stp", m["Stp"].Name)
	require.Equal(t, "stp", m["Stp"].SqlName)
	require.Equal(t, "float64", m["Stp"].Types)
	require.Equal(t, "", m["Stp"].ImportPath)

	require.Equal(t, "Status", m["Status"].Name)
	require.Equal(t, "status", m["Status"].SqlName)
	require.Equal(t, "uint", m["Status"].Types)
	require.Equal(t, "", m["Stp"].ImportPath)
}

func Test_generate(t *testing.T) {
	g := Generate{
		SelfPackage:  "github.com/Ranxy/gormfields/generate/example",
		OutPath:      "./example/example_fields",
		PackageName:  "example_fields",
		UseZeroCheck: true,
	}

	err := os.MkdirAll(g.OutPath, os.ModePerm)
	require.NoError(t, err)
	g.Gen(example.User{})

	t.Cleanup(func() {
		err := os.RemoveAll(g.OutPath)
		require.NoError(t, err)
	})
}
