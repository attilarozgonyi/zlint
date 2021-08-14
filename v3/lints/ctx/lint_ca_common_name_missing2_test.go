package ctx

/*
 * ZLint Copyright 2021 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestCaCommonNameMissing(t *testing.T) {
	inputPath := "caCommonNameMissing.pem"
	expected := lint.Error
	out := test.TestLint("e_ca_common_name_missing2", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCaCommonNameNotMissing(t *testing.T) {
	inputPath := "caCommonNameNotMissing.pem"
	expected := lint.Pass
	out := test.TestLint("e_ca_common_name_missing2", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCaCommonNameNotMissing2(t *testing.T) {
	inputPath := "caCommonNameNotMissing.pem"
	expected := lint.Pass
	ctx := `
[e_ca_common_name_missing2]
BeerHall = "liedershousen"
`
	out := test.TestLintWithCtx("e_ca_common_name_missing2", inputPath, ctx)
	if out.Details != "liedershousen" {
		panic("noooooooooooo")
	}
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCaCommonNameNotMissing3(t *testing.T) {
	inputPath := "caCommonNameNotMissing.pem"
	expected := lint.Pass
	ctx := `
[e_ca_common_name_missing2]
BeerHall = "liedershousenssss"
`
	out := test.TestLintWithCtx("e_ca_common_name_missing2", inputPath, ctx)
	if out.Details != "liedershousenssss" {
		panic("noooooooooooo")
	}
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestConcurrency(t *testing.T) {
	inputPath := "caCommonNameNotMissing.pem"
	expected := lint.Pass
	wg := sync.WaitGroup{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			num := strconv.Itoa(rand.Intn(9999))
			ctx := fmt.Sprintf(`
[e_ca_common_name_missing2]
BeerHall = "%s"
`, num)
			out := test.TestLintWithCtx("e_ca_common_name_missing2", inputPath, ctx)
			if out.Details != num {
				t.Errorf("wanted %s got %s", num, out.Details)
			}
			if out.Status != expected {
				t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
			}
		}()
	}
	wg.Wait()
}

func TestCaCommonNameNotMissing4(t *testing.T) {
	inputPath := "caCommonNameNotMissing.pem"
	expected := lint.Pass
	ctx := `
[CABF_BR]
DoesItWork = "yes, yes it does"

[e_ca_common_name_missing2]
BeerHall = "liedershousenssss"
`
	out := test.TestLintWithCtx("e_ca_common_name_missing2", inputPath, ctx)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
	if out.Details != "liedershousenssss" {
		panic("noooooooooooo")
	}
}
