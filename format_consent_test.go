package iabconsent

import (
	"github.com/go-check/check"
	"time"
)

type FormatConsentSuite struct{}

func (p *FormatConsentSuite) TestFormatConsents(c *check.C) {
	var cases = []struct {
		Consent  *ParsedConsent
		Expected string
	}{
		{
			Consent: &ParsedConsent{
				Version:           1,
				Created:           timeFromDs(14924661858),
				LastUpdated:       timeFromDs(15240021858),
				CMPID:             14,
				CMPVersion:        22,
				ConsentScreen:     30,
				ConsentLanguage:   "FR",
				VendorListVersion: 0,
				PurposesAllowed: map[int]bool{
					2:  true,
					3:  true,
					20: true,
					21: true,
					23: true,
				},
				MaxVendorID:     10,
				IsRangeEncoding: false,
				ConsentedVendors: map[int]bool{
					1: true,
					2: true,
					4: true,
					5: true,
					7: true,
					9: true,
				},
			},
			Expected: "BN5lERiOMYEdiAOAWeFRAAYAAaAAptQ",
		},
		{
			Consent: &ParsedConsent{
				Version:           1,
				Created:           timeFromDs(14924661858),
				LastUpdated:       timeFromDs(15240021858),
				CMPID:             10,
				CMPVersion:        22,
				ConsentScreen:     23,
				ConsentLanguage:   "EN",
				VendorListVersion: 245,
				PurposesAllowed: map[int]bool{
					4:  true,
					5:  true,
					6:  true,
					7:  true,
					9:  true,
					14: true,
					17: true,
					24: true,
				},
				MaxVendorID:     5024,
				IsRangeEncoding: true,
				DefaultConsent:  false,
				NumEntries:      5,
				RangeEntries: []*RangeEntry{
					{
						StartVendorID: 20,
						EndVendorID:   20,
					},
					{
						StartVendorID: 200,
						EndVendorID:   400,
					},
					{
						StartVendorID: 401,
						EndVendorID:   410,
					},
					{
						StartVendorID: 515,
						EndVendorID:   515,
					},
					{
						StartVendorID: 5000,
						EndVendorID:   5024,
					},
				},
			},
			Expected: "BN5lERiOMYEdiAKAWXEND1HoSBE6CAFAApAMgBkIDIgM0AgOJxAnQA",
		},
		{
			Consent: &ParsedConsent{
				Version:           1,
				Created:           timeFromDs(15257231285),
				LastUpdated:       timeFromDs(15257231285),
				CMPID:             7,
				CMPVersion:        1,
				ConsentScreen:     1,
				ConsentLanguage:   "EN",
				VendorListVersion: 14,
				PurposesAllowed: map[int]bool{
					1: true,
					2: true,
					3: true,
					4: true,
					5: true,
				},
				MaxVendorID:     112,
				IsRangeEncoding: true,
				DefaultConsent:  false,
				NumEntries:      4,
				RangeEntries: []*RangeEntry{
					{
						StartVendorID: 9,
						EndVendorID:   9,
					},
					{
						StartVendorID: 25,
						EndVendorID:   25,
					},
					{
						StartVendorID: 27,
						EndVendorID:   28,
					},
					{
						StartVendorID: 30,
						EndVendorID:   30,
					},
				},
			},
			Expected: "BONZt-1ONZt-1AHABBENAO-AAAAHCAEAASABmADYAOAAeA",
		},
	}

	for _, tc := range cases {
		c.Log(tc)
		cs := tc.Consent.ToConsentString()

		c.Assert(cs, check.Equals, tc.Expected)
	}
}

func (p *FormatConsentSuite) TestParse(c *check.C) {
	_, err := Parse("BOOd8eCOPpnYYAKABCFRBCAAAAAcQAAAgAAYEBAUKgCAwAA0KAAIABABAiAAgQ1AxAbIeGiiAAQugCFAYABAAAAADAECAAAAQFBiA6OGgA")
	c.Check(err, check.IsNil)
}

func timeFromDs(deciseconds int64) time.Time {
	return time.Unix(deciseconds/10, (deciseconds%10)*int64(time.Millisecond*100)).UTC()
}

var _ = check.Suite(&FormatConsentSuite{})
