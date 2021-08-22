package lint

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"

	"github.com/pelletier/go-toml"
)

type Context struct {
	tree *toml.Tree
}

func (c Context) Configure(lint interface{}, namespace string) error {
	if err := c.deser(lint, namespace); err != nil {
		return err
	}
	return c.findAll(lint)
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

func (c Context) findAll(i interface{}) error {
	value := reflect.Indirect(reflect.ValueOf(i))
	if value.Kind() != reflect.Struct {
		return nil
	}
	kind := value.Kind()
	fmt.Println(kind)
	for field := 0; field < value.NumField(); field++ {
		field := value.Field(field)
		if !field.CanInterface() {
			continue
		}
		var val reflect.Value
		switch t := field.Interface().(type) {
		case Global:
			err := c.deser(&t, "")
			if err != nil {
				return err
			}
			val = reflect.ValueOf(t)
		case *Global:
			if t == nil {
				t = &Global{}
			}
			err := c.deser(t, "")
			if err != nil {
				return err
			}
			val = reflect.ValueOf(t)
		case RFC5280Context:
			err := c.deser(&t, string(RFC5280))
			if err != nil {
				return err
			}
			val = reflect.ValueOf(t)
		case *RFC5280Context:
			if t == nil {
				t = &RFC5280Context{}
			}
			err := c.deser(t, string(RFC5280))
			if err != nil {
				return err
			}
			val = reflect.ValueOf(t)
		case RFC5480Context:
			err := c.deser(&t, string(RFC5480))
			if err != nil {
				return err
			}
			val = reflect.ValueOf(t)
		case *RFC5480Context:
			if t == nil {
				t = &RFC5480Context{}
			}
			err := c.deser(t, string(RFC5480))
			if err != nil {
				return err
			}
			val = reflect.ValueOf(t)
		case RFC5891Context:
			err := c.deser(&t, string(RFC5891))
			if err != nil {
				return err
			}
			val = reflect.ValueOf(t)
		case *RFC5891Context:
			if t == nil {
				t = &RFC5891Context{}
			}
			err := c.deser(t, string(RFC5891))
			if err != nil {
				return err
			}
			val = reflect.ValueOf(t)
		case CABFBaselineRequirementsContext:
			err := c.deser(&t, string(CABFBaselineRequirements))
			if err != nil {
				return err
			}
			val = reflect.ValueOf(t)
		case *CABFBaselineRequirementsContext:
			addr := field.CanAddr()
			fmt.Println(addr)
			if t == nil {
				t = &CABFBaselineRequirementsContext{}
			}
			err := c.deser(t, string(CABFBaselineRequirements))
			if err != nil {
				return err
			}
			val = reflect.ValueOf(t)
		case CABFEVGuidelinesContext:
			err := c.deser(&t, string(CABFEVGuidelines))
			if err != nil {
				return err
			}
			val = reflect.ValueOf(t)
		case *CABFEVGuidelinesContext:
			if t == nil {
				t = &CABFEVGuidelinesContext{}
			}
			err := c.deser(t, string(CABFEVGuidelines))
			if err != nil {
				return err
			}
			val = reflect.ValueOf(t)
		case MozillaRootStorePolicyContext:
			err := c.deser(&t, string(MozillaRootStorePolicy))
			if err != nil {
				return err
			}
			val = reflect.ValueOf(t)
		case *MozillaRootStorePolicyContext:
			if t == nil {
				t = &MozillaRootStorePolicyContext{}
			}
			err := c.deser(t, string(MozillaRootStorePolicy))
			if err != nil {
				return err
			}
			val = reflect.ValueOf(t)
		case AppleRootStorePolicyContext:
			err := c.deser(&t, string(AppleRootStorePolicy))
			if err != nil {
				return err
			}
			val = reflect.ValueOf(t)
		case *AppleRootStorePolicyContext:
			if t == nil {
				t = &AppleRootStorePolicyContext{}
			}
			err := c.deser(t, string(AppleRootStorePolicy))
			if err != nil {
				return err
			}
			val = reflect.ValueOf(t)
		case CommunityContext:
			err := c.deser(&t, string(Community))
			if err != nil {
				return err
			}
			val = reflect.ValueOf(t)
		case *CommunityContext:
			if t == nil {
				t = &CommunityContext{}
			}
			err := c.deser(t, string(Community))
			if err != nil {
				return err
			}
			val = reflect.ValueOf(t)
		case EtsiEsiContext:
			err := c.deser(&t, string(EtsiEsi))
			if err != nil {
				return err
			}
			val = reflect.ValueOf(t)
		case *EtsiEsiContext:
			if t == nil {
				t = &EtsiEsiContext{}
			}
			err := c.deser(t, string(EtsiEsi))
			if err != nil {
				return err
			}
			val = reflect.ValueOf(t)
		default:
			if !field.CanAddr() {
				continue
			}
			field = field.Addr()
			switch t2 := field.Interface().(type) {
			default:
				err := c.findAll(t2)
				if err != nil {
					return err
				}
				continue
			}
		}
		field.Set(val)
	}
	return nil
}

func (c Context) deser(i interface{}, namespace string) error {
	target := c.tree.Get(namespace)
	if target == nil {
		return nil
	}
	return target.(*toml.Tree).Unmarshal(i)
}

type Global struct {
	YarHar string
}
type RFC5280Context struct{}
type RFC5480Context struct{}
type RFC5891Context struct{}
type CABFBaselineRequirementsContext struct {
	DoesItWork string
}
type CABFEVGuidelinesContext struct{}
type MozillaRootStorePolicyContext struct{}
type AppleRootStorePolicyContext struct{}
type CommunityContext struct{}
type EtsiEsiContext struct{}

func stripAll(i interface{}, namespace string) map[string]interface{} {
	return map[string]interface{}{namespace: _stripAll(i)}
}

func _stripAll(i interface{}) interface{} {
	value := reflect.Indirect(reflect.ValueOf(i))
	if value.Kind() != reflect.Struct {
		return i
	}
	m := make(map[string]interface{})
	for field := 0; field < value.NumField(); field++ {
		name := value.Type().Field(field).Name
		field := value.Field(field)
		if !field.CanInterface() {
			continue
		}
		switch t := field.Interface().(type) {
		case Global:
		case *Global:
		case RFC5280Context:
		case *RFC5280Context:
		case RFC5480Context:
		case *RFC5480Context:
		case RFC5891Context:
		case *RFC5891Context:
		case CABFBaselineRequirementsContext:
		case *CABFBaselineRequirementsContext:
		case CABFEVGuidelinesContext:
		case *CABFEVGuidelinesContext:
		case MozillaRootStorePolicyContext:
		case *MozillaRootStorePolicyContext:
		case AppleRootStorePolicyContext:
		case *AppleRootStorePolicyContext:
		case CommunityContext:
		case *CommunityContext:
		case EtsiEsiContext:
		case *EtsiEsiContext:
		default:
			m[name] = _stripAll(t)
		}
	}
	return m
}
