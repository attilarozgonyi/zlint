package lint

import (
	"fmt"
	"testing"
)

func TestBasicConfiguration(t *testing.T) {
	ctx, err := NewContextFromString(`
[security]
good = true
`)
	if err != nil {
		panic(err)
	}
	type lol struct {
		Good bool
	}
	sec := lol{}
	err = ctx.Configure(&sec, "security")
	if err != nil {
		panic(err)
	}
	fmt.Println(sec)
}

func TestBasicConfigurationNotPresent(t *testing.T) {
	ctx, err := NewContextFromString(`
[security]
good = true
`)
	if err != nil {
		panic(err)
	}
	type lol struct {
		Good bool
	}
	sec := lol{}
	err = ctx.Configure(&sec, "securitys")
	if err != nil {
		panic(err)
	}
	fmt.Println(sec)
}

func TestBasicConfigurationBad(t *testing.T) {
	ctx, err := NewContextFromString(`
[security]
good = true
`)
	if err != nil {
		panic(err)
	}
	type lol struct {
		Goods bool
		Boots int
	}
	sec := lol{}
	err = ctx.Configure(&sec, "security")
	if err != nil {
		panic(err)
	}
	fmt.Println(sec)
}
