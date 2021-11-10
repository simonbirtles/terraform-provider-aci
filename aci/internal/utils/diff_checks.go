package utils

import (
	"github.com/hashicorp/terraform/helper/schema"
	"strings"
)

// Compares the TF and APIC read value against a known default value to determine
// if the TF or APIC values are actually defaults. - For some attributes the APIC 
// has two defaults, you can configure the APIC with either default - usually a 
// actual value or an empty string, but the APIC read will only return one of 
// the defaults, so we have this function to compare the TF value & APIC Value and 
// what we know the real default on and APIC read would be and supress if they 
// match.
//
func SuppressDiffCheck(compare_a, compare_b string) schema.SchemaDiffSuppressFunc {

	return func(k, old, new string, d *schema.ResourceData) bool {

		if ( (strings.ToLower(old) == strings.ToLower(compare_a) || old == compare_b) &&
			 (strings.ToLower(new) == strings.ToLower(compare_a) || new == compare_b) ) {
			return true 
		}
		return false
	}
}

// Checks the value from the APIC read with a given attribute to determine if
// they match and if so then supress the diff
//
func SuppressDiffCheckDefault(apic_default string) schema.SchemaDiffSuppressFunc {

	// old is from TF state file, new is from the APIC Read for this MO
	return func(k, old, new string, d *schema.ResourceData) bool {
		if strings.ToLower(new) == strings.ToLower(apic_default) {
			return true
		}
		return false
	}
}