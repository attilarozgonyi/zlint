package cabf_br

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
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

type MyLint struct {
	BottlesOfBeerOnTheWall int
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_ca_too_few_beers",
		Description:   "CA Certificates MUST have at least 99 beers.",
		Citation:      "BRs: 7.1.4.3.1",
		Source:        lint.CABFBaselineRequirements,
		EffectiveDate: util.CABV148Date,
		Lint:          NewMyLint,
	})
}

// This is the new thing. You give us a struct to deserialize
// into and we will get the appropriate context into it.
//
// In this case, the lint itself holds the data in quesiton.
func (l *MyLint) Configure() interface{} {
	return l
}

func NewMyLint() lint.LintInterface {
	return &MyLint{}
}

func (l *MyLint) CheckApplies(c *x509.Certificate) bool {
	return util.IsCACert(c)
}

func (l *MyLint) Execute(c *x509.Certificate) *lint.LintResult {
	if l.BottlesOfBeerOnTheWall < 99 {
		return &lint.LintResult{Status: lint.Error, Details: "Time for a beer run!"}
	} else {
		return &lint.LintResult{Status: lint.Pass}
	}
}
