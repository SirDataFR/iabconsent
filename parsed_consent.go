/*

Package iabconsent provides structs and methods for parsing
Vendor Consent Strings as defined by the IAB Consent String 1.1 Spec.
More info on the spec here:
https://github.com/InteractiveAdvertisingBureau/GDPR-Transparency-and-Consent-Framework/blob/master/Consent%20string%20and%20vendor%20list%20formats%20v1.1%20Final.md#vendor-consent-string-format-.

Copyright (c) 2018 LiveRamp. All rights reserved.

Written by Andy Day, Software Engineer @ LiveRamp
for use in the LiveRamp Pixel Server.

*/
package iabconsent

import (
	"time"
	"fmt"
)

// ParsedConsent represents data extracted from an IAB Consent String, v1.1.
type ParsedConsent struct {
	Version           int
	Created           time.Time
	LastUpdated       time.Time
	CMPID             int
	CMPVersion        int
	ConsentScreen     int
	ConsentLanguage   string
	VendorListVersion int
	PurposesAllowed   map[int]bool
	MaxVendorID       int
	IsRangeEncoding   bool
	ConsentedVendors  map[int]bool
	DefaultConsent    bool
	NumEntries        int
	RangeEntries      []*RangeEntry
}

// EveryPurposeAllowed returns true iff every purpose number in ps exists in
// the ParsedConsent, otherwise false.
func (p *ParsedConsent) EveryPurposeAllowed(ps []int) bool {
	for _, rp := range ps {
		if !p.PurposesAllowed[rp] {
			return false
		}
	}
	return true
}

// PurposeAllowed returns true if the purpose number exists in the ParsedConsent and is true,
// otherwise false
func (p *ParsedConsent) PurposeAllowed(v int) bool {
	return p.PurposesAllowed[v]
}

// VendorAllowed returns true if the ParsedConsent contains affirmative consent
// for VendorID v.
func (p *ParsedConsent) VendorAllowed(v int) bool {
	if p.IsRangeEncoding {
		for _, re := range p.RangeEntries {
			if re.StartVendorID <= v && v <= re.EndVendorID {
				return !p.DefaultConsent
			}
		}
		return p.DefaultConsent
	}

	return p.ConsentedVendors[v]
}

func (p *ParsedConsent) ToConsentString() string {
	return Format(p)
}

func (p *ParsedConsent) ToString() string {
	return fmt.Sprintf("Version=%v, Created=%v, LastUpdated=%v, CMPID=%v, CMPVersion=%v," +
		" ConsentScreen=%v, ConsentLanguage=%v, VendorListVersion=%v, PurposesAllowed=%v," +
		" MaxVendorID=%v, IsRangeEncoding=%v, ConsentedVendors=%v, RangeEntries=%v, DefaultConsent=%v",
		p.Version, p.Created, p.LastUpdated, p.CMPID, p.CMPVersion,
		p.ConsentScreen, p.ConsentLanguage, p.VendorListVersion, p.PurposesAllowed,
		p.MaxVendorID, p.IsRangeEncoding, p.ConsentedVendors, p.RangeEntries, p.DefaultConsent)
}

// RangeEntry defines an inclusive range of vendor IDs from StartVendorID to
// EndVendorID.
type RangeEntry struct {
	StartVendorID int
	EndVendorID   int
}
