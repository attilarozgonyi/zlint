package lint

import (
	"io"
	"os"
	"strings"

	"github.com/pelletier/go-toml"
)

type Context struct {
	tree *toml.Tree
}

func (c Context) Configure(lint interface{}, namespace string) error {
	target := c.tree.Get(namespace)
	if target == nil {
		return nil
	}
	return target.(*toml.Tree).Unmarshal(lint)
}

func NewContextFromFile(file string) (Context, error) {
	if file == "" {
		return NewEmptyContext(), nil
	}
	f, err := os.Open(file)
	if err != nil {
		return Context{}, err
	}
	defer f.Close()
	return NewContext(f)
}

func NewContextFromString(ctx string) (Context, error) {
	return NewContext(strings.NewReader(ctx))
}

func NewContext(r io.Reader) (Context, error) {
	tree, err := toml.LoadReader(r)
	if err != nil {
		return Context{}, nil
	}
	return Context{tree}, nil
}

func NewEmptyContext() Context {
	ctx, _ := NewContextFromString("")
	return ctx
}
