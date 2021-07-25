package cabf_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"

	"github.com/pelletier/go-toml"
)

func TestMarhsal(t *testing.T) {
	l := NewMyLint().(lint.Configurable)
	//tree, err := toml.TreeFromMap(map[string]interface{}{
	//	"ca_beer": l.Configure(),
	//})
	//if err != nil {
	//	t.Fatal(err)
	//}
	b, err := toml.Marshal(map[string]interface{}{
		"ca_beer": l.Configure(),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(b))
	//b, err := tree.Marshal()
	//if err != nil {
	//	t.Fatal(err)
	//}
	//t.Log(string(b))
	//b, err := toml.Marshal(l.Configure())
	//if err != nil {
	//	panic(err)
	//}
	//t.Log(string(b))
}
