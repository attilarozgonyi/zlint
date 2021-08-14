package lint

import (
	"fmt"
	"reflect"
	"testing"
)

type Wat struct {
	Inner       inner
	Exp         int
	notExported int
}

type inner struct {
	VeryInner veryInner
	String    string
}

type veryInner struct {
	Exported RFC5280Context
}

func TestAsdasd(t *testing.T) {
	e := Wat{}
	findAll(e)
	//value := reflect.ValueOf(e)
	//fields := value.NumField()
	//for field := 0; field < fields; field++ {
	//	field := value.Field(field)
	//	if !field.CanInterface() {
	//		continue
	//	}
	//	switch field.Interface().(type) {
	//	case RFC5280Context:
	//		t.Log("yoooo")
	//	default:
	//		t.Log("like whatever")
	//	}
	//}
}

func findAll(i interface{}) {
	value := reflect.ValueOf(i)
	if value.Type().Kind() != reflect.Struct {
		fmt.Println("not a struct")
		return
	}
	for field := 0; field < value.NumField(); field++ {
		field := value.Field(field)
		if !field.CanInterface() {
			continue
		}
		switch t := field.Interface().(type) {
		case RFC5280Context:
			fmt.Println("yoooo")
		case *RFC5280Context:
			fmt.Println("yoooo ptr")
		case RFC5480Context:
			fmt.Println("yoooo")
		case *RFC5480Context:
			fmt.Println("yoooo ptr")
		case RFC5891Context:
			fmt.Println("yoooo")
		case *RFC5891Context:
			fmt.Println("yoooo ptr")
		case CABFBaselineRequirementsContext:
			fmt.Println("yoooo")
		case *CABFBaselineRequirementsContext:
			fmt.Println("yoooo ptr")
		case CABFEVGuidelinesContext:
			fmt.Println("yoooo")
		case *CABFEVGuidelinesContext:
			fmt.Println("yoooo ptr")
		case MozillaRootStorePolicyContext:
			fmt.Println("yoooo")
		case *MozillaRootStorePolicyContext:
			fmt.Println("yoooo ptr")
		case AppleRootStorePolicyContext:
			fmt.Println("yoooo")
		case *AppleRootStorePolicyContext:
			fmt.Println("yoooo ptr")
		case CommunityContext:
			fmt.Println("yoooo")
		case *CommunityContext:
			fmt.Println("yoooo ptr")
		case EtsiEsiContext:
			fmt.Println("yoooo")
		case *EtsiEsiContext:
			fmt.Println("yoooo ptr")
		default:
			findAll(t)
		}
	}
}

func TestHate(t *testing.T) {
	a := struct {
		A string
		B string
	}{}
	asdased(&a)
	t.Log(a.B)
}

func asd(i interface{}) {
	asdased(i)
}

func asdased(i interface{}) {
	v := reflect.Indirect(reflect.ValueOf(i))
	v.Field(1).SetString("yaaas")
}

func TestTopLevel(t *testing.T) {
	ctx := `
[CABF_BR]
DoesItWork = "yes, yes it does"

[e_ca_common_name_missing2]
BeerHall = "liedershousenssss"
`
	type caCommonNameMissing struct {
		BeerHall string
		Working  *CABFBaselineRequirementsContext
	}
	a := &caCommonNameMissing{}
	c, err := NewContextFromString(ctx)
	if err != nil {
		t.Fatal(err)
	}
	err = c.Configure(a, "e_ca_common_name_missing2")
	if err != nil {
		t.Fatal(err)
	}
	if a.Working.DoesItWork != "yes, yes it does" {
		t.Fatal("scream")
	}
}

func TestTopLevel2(t *testing.T) {
	ctx := `
[CABF_BR]
DoesItWork = "yes, yes it does"

[e_ca_common_name_missing2]
BeerHall = "liedershousenssss"
`
	type caCommonNameMissing struct {
		BeerHall string
		Working  CABFBaselineRequirementsContext
	}
	a := &caCommonNameMissing{}
	c, err := NewContextFromString(ctx)
	if err != nil {
		t.Fatal(err)
	}
	err = c.Configure(a, "e_ca_common_name_missing2")
	if err != nil {
		t.Fatal(err)
	}
	if a.Working.DoesItWork != "yes, yes it does" {
		t.Fatal("scream")
	}
}

func TestNested(t *testing.T) {
	ctx := `
[CABF_BR]
DoesItWork = "yes, yes it does"

[e_ca_common_name_missing2]
BeerHall = "liedershousenssss"
`
	type caCommonNameMissing struct {
		BeerHall string
		Inner    struct {
			Working *CABFBaselineRequirementsContext
		}
	}
	b := struct {
		Working *CABFBaselineRequirementsContext
	}{}
	a := &caCommonNameMissing{Inner: b}
	c, err := NewContextFromString(ctx)
	if err != nil {
		t.Fatal(err)
	}
	err = c.Configure(a, "e_ca_common_name_missing2")
	if err != nil {
		t.Fatal(err)
	}
	if a.Inner.Working.DoesItWork != "yes, yes it does" {
		t.Fatal("scream")
	}
}

func TestNested2(t *testing.T) {
	ctx := `
[CABF_BR]
DoesItWork = "yes, yes it does"

[e_ca_common_name_missing2]
BeerHall = "liedershousenssss"
`
	type caCommonNameMissing struct {
		BeerHall string
		Inner    struct {
			Working CABFBaselineRequirementsContext
		}
	}
	b := struct {
		Working CABFBaselineRequirementsContext
	}{}
	a := &caCommonNameMissing{Inner: b}
	c, err := NewContextFromString(ctx)
	if err != nil {
		t.Fatal(err)
	}
	err = c.Configure(a, "e_ca_common_name_missing2")
	if err != nil {
		t.Fatal(err)
	}
	if a.Inner.Working.DoesItWork != "yes, yes it does" {
		t.Fatal("scream")
	}
}

func TestNestedNone(t *testing.T) {
	ctx := `
[e_ca_common_name_missing2]
BeerHall = "liedershousenssss"
`
	type caCommonNameMissing struct {
		BeerHall string
		Inner    struct {
			Working CABFBaselineRequirementsContext
		}
	}
	b := struct {
		Working CABFBaselineRequirementsContext
	}{}
	a := &caCommonNameMissing{Inner: b}
	c, err := NewContextFromString(ctx)
	if err != nil {
		t.Fatal(err)
	}
	err = c.Configure(a, "e_ca_common_name_missing2")
	if err != nil {
		t.Fatal(err)
	}
	if a.Inner.Working.DoesItWork != "" {
		t.Fatal("scream")
	}
}

func TestGlobal(t *testing.T) {
	ctx := `
YarHar = "This is the bomb"
[e_ca_common_name_missing2]
BeerHall = "liedershousenssss"
`
	type caCommonNameMissing struct {
		BeerHall string
		Inner    struct {
			Working *Global
		}
	}
	a := &caCommonNameMissing{}
	c, err := NewContextFromString(ctx)
	if err != nil {
		t.Fatal(err)
	}
	err = c.Configure(a, "e_ca_common_name_missing2")
	if err != nil {
		t.Fatal(err)
	}
	if a.Inner.Working.YarHar != "This is the bomb" {
		t.Fatal(a.Inner.Working.YarHar)
	}
}

func TestMixed(t *testing.T) {
	ctx := `
YarHar = "This is the bomb"
[e_ca_common_name_missing2]
BeerHall = "liedershousenssss"

[CABF_BR]
DoesItWork = "yes, yes it does"
`
	type caCommonNameMissing struct {
		BeerHall string
		Inner    struct {
			Working *Global
		}
		Berp CABFBaselineRequirementsContext
	}
	a := &caCommonNameMissing{}
	c, err := NewContextFromString(ctx)
	if err != nil {
		t.Fatal(err)
	}
	err = c.Configure(a, "e_ca_common_name_missing2")
	if err != nil {
		t.Fatal(err)
	}
	if a.Inner.Working.YarHar != "This is the bomb" {
		t.Fatal(a.Inner.Working.YarHar)
	}
	if a.Berp.DoesItWork != "yes, yes it does" {
		t.Fatal(a.Berp.DoesItWork)
	}
}

func TestIsCopy(t *testing.T) {
	ctx := `
YarHar = "This is the bomb"
[e_ca_common_name_missing2]
BeerHall = "liedershousenssss"

[CABF_BR]
DoesItWork = "yes, yes it does"
`
	type caCommonNameMissing struct {
		BeerHall string
		Berp     *CABFBaselineRequirementsContext
	}
	a := &caCommonNameMissing{}
	c, err := NewContextFromString(ctx)
	if err != nil {
		t.Fatal(err)
	}
	err = c.Configure(a, "e_ca_common_name_missing2")
	if err != nil {
		t.Fatal(err)
	}
	if a.Berp.DoesItWork != "yes, yes it does" {
		t.Fatal(a.Berp.DoesItWork)
	}
	a.Berp.DoesItWork = "something else"
	t.Log(c.tree.Get("CABF_BR"))
}
