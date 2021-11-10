// go test terraform-provider-aci/aci -v testacc
//
//
package aci

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"os"
	"testing"
	"log"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	log.Printf("[TESTING] - Init")

	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"aci": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	log.Printf("[TESTING] - TestProvider")

	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}

}

func TestProvider_impl(t *testing.T) {
	log.Printf("[TESTING] - TestProvider_impl")

	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	log.Printf("[TESTING] - testAccPreCheck")
	if v := os.Getenv("ACI_APIC_USERNAME"); v == "" {
		t.Fatal("ACI_APIC_USERNAME must be set for acceptance tests")
	}

	if v := os.Getenv("ACI_APIC_PASSWORD"); v == "" {
		t.Fatal("ACI_APIC_PASSWORD must be set for acceptance tests")
	}

	if v := os.Getenv("ACI_APIC"); v == "" {
		t.Fatal("ACI_APIC must be set for acceptance tests")
	}
}
