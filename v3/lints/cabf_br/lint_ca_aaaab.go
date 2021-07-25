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

type LintEmbedsAConfiguration struct {
	configuration                        embeddedConfiguration
	SomeOtherFieldThatWeDontWantToExpose int
}

type embeddedConfiguration struct {
	IsWebPKI bool `toml:"is_web_pki" comment:"Indicates that the certificate is intended for the Web PKI."`
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "w_web_pki_cert",
		Description:   "CA Certificates SHOULD....something....about the web pki",
		Citation:      "BRs: 7.1.4.3.1",
		Source:        lint.CABFBaselineRequirements,
		EffectiveDate: util.CABV148Date,
		Lint:          NewLintEmbedsAConfiguration,
	})
}

// A pointer to an embedded struct may be passed to the framework
// if the author does not wish to expose certain fields in their primary struct.
func (l *LintEmbedsAConfiguration) Configure() interface{} {
	return &l.configuration
}

func NewLintEmbedsAConfiguration() lint.LintInterface {
	return &LintEmbedsAConfiguration{configuration: embeddedConfiguration{}}
}

func (l *LintEmbedsAConfiguration) CheckApplies(c *x509.Certificate) bool {
	return util.IsCACert(c)
}

func (l *LintEmbedsAConfiguration) Execute(c *x509.Certificate) *lint.LintResult {
	if l.configuration.IsWebPKI {
		return &lint.LintResult{Status: lint.Warn, Details: "Time for a beer run!"}
	} else {
		return &lint.LintResult{Status: lint.Pass}
	}
}
