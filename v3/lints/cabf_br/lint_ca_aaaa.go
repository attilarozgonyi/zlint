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

type LintItselfIsConfigurable struct {
	BottlesOfBeerOnTheWall int `comment:"This MUST be set to the number of bottles of beer that were on the signing CAs wall at the time of issuance."`
	ISpyWithMyLittleEye    struct {
		Descriptions []string `comment:"Any number of descriptions that are valid for the subject."`
		Subject      string
	} `toml:"i_spy_with_my_little_eye"`
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_ca_too_few_beers",
		Description:   "CA Certificates MUST have at least 99 beers.",
		Citation:      "BRs: 7.1.4.3.1",
		Source:        lint.CABFBaselineRequirements,
		EffectiveDate: util.CABV148Date,
		Lint:          NewLintItselfIsConfigurable,
	})
}

func NewLintItselfIsConfigurable() lint.LintInterface {
	// Go code encourages that the zero value for a struct be valid, however
	// if you wish for your Lint to have non-zero defaults then your constructor is
	// the place to do so. The value initialized in this constructor WILL appear
	// in the example configuration printed through the `-exampleConfig flag`.
	return &LintItselfIsConfigurable{ISpyWithMyLittleEye: struct {
		Descriptions []string `comment:"Any number of descriptions that are valid for the subject."`
		Subject      string
	}{Descriptions: []string{"larger than a bread box", "smaller than a barn"}, Subject: "A car"}}
}

// In this case, the struct that is the lint itself is the target for configuration.
func (l *LintItselfIsConfigurable) Configure() interface{} {
	return l
}

func (l *LintItselfIsConfigurable) CheckApplies(c *x509.Certificate) bool {
	return util.IsCACert(c)
}

func (l *LintItselfIsConfigurable) Execute(c *x509.Certificate) *lint.LintResult {
	if l.BottlesOfBeerOnTheWall < 99 {
		return &lint.LintResult{Status: lint.Error}
	} else {
		return &lint.LintResult{Status: lint.Pass}
	}
}
