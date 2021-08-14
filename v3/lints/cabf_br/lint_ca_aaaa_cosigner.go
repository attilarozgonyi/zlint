package cabf_br

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_sub_ca_eku_incompatible_values",
		Description:   "Subordinate CA extkeyUsage: if serverAuth is present, then emailProtection, codeSigning, timeStamping, and anyExtendedKeyUsage MUST NOT be present.",
		Citation:      "BRs: 7.1.2.2",
		Source:        lint.CABFBaselineRequirements,
		EffectiveDate: util.CABFBRs_1_7_1_Date,
		Lint:          NewSubCAEkuIncompatibleValues,
	})
}

type subCAEkuIncompatibleValues struct {
	CrossSigner string `comment:"Section 7.1.2.2:  For Cross Certificates that share a Subject Distinguished Name and Subject Public Key with a Root Certificate operated in accordance with these Requirements, this extension MAY be present. If present, this extension SHOULD NOT be marked critical. This extension MUST only contain usages for which the issuing CA has verified the Cross Certificate is authorized to assert. This extension MAY contain the anyExtendedKeyUsage [RFC5280] usage, if the Root Certificate(s) associated with this Cross Certificate are operated by the same organization as the issuing Root Certificate."`
}

func (l *subCAEkuIncompatibleValues) Configure() interface{} {
	return l
}

func NewSubCAEkuIncompatibleValues() lint.LintInterface {
	return &subCAEkuIncompatibleValues{}
}

func (l *subCAEkuIncompatibleValues) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubCA(c) && util.IsExtInCert(c, util.EkuSynOid)
}

func (l *subCAEkuIncompatibleValues) Execute(c *x509.Certificate) *lint.LintResult {
	// Check if l.CrossSigner was provided.
	// Parse to cert.
	// Do appropriate checks against section 7.1.2.2
	//...
	return nil
}
