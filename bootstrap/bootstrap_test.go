package bootstrap

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerator_writeMain(t *testing.T) {
	g := Generator{
		PkgPath:       "abc.com/abc/def",
		PkgName:       "def",
		Types:         []string{"Typea", "UserList"},
		OutPath:       "",
		BuildTags:     "",
		GenBuildFlags: "",
		NoFormat:      false,
	}

	t.Cleanup(func() {
		err := os.RemoveAll(filepath.Join("./", g.PkgName+"boot.go"))
		require.NoError(t, err)
	})

	g.writeMain()
}
