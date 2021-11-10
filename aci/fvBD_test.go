//
package aci // or general total testing ?

import (
    "fmt"
    "log"
    _"github.com/hashicorp/terraform/helper/acctest"
    "github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
	"time"
)

func TestAccAciBridgeDomain(t *testing.T) {

    resource.Test(t, resource.TestCase{

        PreCheck:     func() { testAccPreCheck(t) },

        Providers:    testAccProviders,

        //CheckDestroy: testCheckBridgeDomainConfigDestroy_basic,

        Steps: []resource.TestStep{

            // step one - Bridge Domains
            resource.TestStep{

                Config: testBridgeDomainTestConfig_basic(),

                Check: resource.ComposeTestCheckFunc(

					
					// check correct tenant DN assigned
					testDelay("1 - 1500ms", 1500),
					resource.TestCheckResourceAttr("aci_fvBD.bd_c", "tenant_id", "uni/tn-TEN_ENGINEERING"),					
					
					// check the BD was created and with expected DN
					//testDelay("2"),
					resource.TestCheckResourceAttr("aci_fvBD.bd_c", "dn", "uni/tn-TEN_ENGINEERING/BD-BD_ENGINEERING_IC"),
					
                    // check the BD was created with expected defaults values 
					// as provided by APIC as we leave to APIC default in testing for these attributes
					//testDelay("3"),
					resource.TestCheckResourceAttr("aci_fvBD.bd_c", "type", "regular"),
					
					//testDelay("4"),
					resource.TestCheckResourceAttr("aci_fvBD.bd_c", "unicastroute", "yes"),
					
					//testDelay("5"),
					resource.TestCheckResourceAttr("aci_fvBD.bd_c", "unkmacucastact", "proxy"),
					
					//testDelay("6"),
                    resource.TestCheckResourceAttr("aci_fvBD.bd_c", "unkmcastact", "flood"),

					// detailed check
					//testDelay("7"),					
                    testCheckBridgeDomainConfigCreate_basic("aci_fvBD.bd_c"),

                ),

            },
            // step two - 

        },
        
    })
    
}

func testDelay(num string, delay int64) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		log.Printf("[DEBUG] Delay Number: %s", num)
		time.Sleep(time.Duration(delay) * time.Millisecond)
		return nil
	}
}

func testCheckBridgeDomainConfigCreate_basic(name string) resource.TestCheckFunc {

    return func(s *terraform.State) error {

        // Ensure we have enough information in state to look up in API
        rs, ok := s.RootModule().Resources[name]
        if !ok {
            return fmt.Errorf("Not found: %s", name)
        }

        name := rs.Primary.Attributes["name"]
        dn := rs.Primary.Attributes["dn"]
        iplearning := rs.Primary.Attributes["iplearning"]   
        
        log.Printf("[INFO] Bridge Domain: name:%s dn:%s iplearning:%s", name, dn, iplearning)
        tenantId, hasTenantId := rs.Primary.Attributes["tenant_id"]
        if !hasTenantId {
            return fmt.Errorf("[DEBUG] Bad: no tenant found in state for bridge domain: %s", name)
        }
        log.Printf("\n\n[INFO]\n\n\n Bridge Domain: Tenant ID:%s %s %s %s\n\n\n\n", tenantId, name, dn, iplearning)
        
		return nil

    }
}

/*
func testCheckBridgeDomainConfigDestroy_basic(s *terraform.State) error {

    conn := testAccProvider.Meta().(*ArmClient).publicIPClient

    for _, rs := range s.RootModule().Resources {
        if rs.Type != "azurerm_public_ip" {
            continue
        }

        name := rs.Primary.Attributes["name"]
        resourceGroup := rs.Primary.Attributes["resource_group_name"]

        resp, err := conn.Get(resourceGroup, name, "")

        if err != nil {
            return nil
        }

        if resp.StatusCode != http.StatusNotFound {
            return fmt.Errorf("Public IP still exists:\n%#v", resp.Properties)
        }
    }

    return nil
}
*/

// do we make this func public so for example we can chain configs to merge 
// as they will grow as we add more and more MO's
func testBridgeDomainTestConfig_basic() string {
    return fmt.Sprintf(`
        
        variable "tenant" {
            default = "TEN_ENGINEERING"
        }

        variable "bridge_domain" {
            default = "BD_ENGINEERING_IC"
        }

        data "aci_fvTenant" "tenant" {
            name    = "${var.tenant}"
        }

        resource "aci_fvBD" "bd_c" {
            tenant_id               = "${data.aci_fvTenant.tenant.id}"
            name                    = "${var.bridge_domain}"
            descr                   = "Testing BD For Terraform Test Pack"
            iplearning              = "yes"
            arpflood                = "yes"
            limitiplearntosubnets   = "no"
        }`)
}
