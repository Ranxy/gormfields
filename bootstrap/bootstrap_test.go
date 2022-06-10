package bootstrap

import "testing"

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

	g.writeMain()
}
