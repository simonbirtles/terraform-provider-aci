package aci

import (

	// terraform 	
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	
	// aci rest module 
	"aci"


	"terraform-provider-aci/aci/mo/query"
	// aci managed objects

	// fv
	"terraform-provider-aci/aci/mo/fv/fvAEPg"
	"terraform-provider-aci/aci/mo/fv/fvAp"
	"terraform-provider-aci/aci/mo/fv/fvBD"
	"terraform-provider-aci/aci/mo/fv/fvCtx"
	"terraform-provider-aci/aci/mo/fv/fvRsBd"
	"terraform-provider-aci/aci/mo/fv/fvRsCons"
	"terraform-provider-aci/aci/mo/fv/fvRsCtx"
	"terraform-provider-aci/aci/mo/fv/fvRsDomAtt"
	"terraform-provider-aci/aci/mo/fv/fvRsProv"
	"terraform-provider-aci/aci/mo/fv/fvSubnet"
	"terraform-provider-aci/aci/mo/fv/fvTenant"

	// phys
	"terraform-provider-aci/aci/mo/phys/physDomP"

	// tag
	"terraform-provider-aci/aci/mo/tag/tagInst"

	// vmm
	"terraform-provider-aci/aci/mo/vmm/vmmDomP"

	// vz
	"terraform-provider-aci/aci/mo/vz/vzBrCP"
	"terraform-provider-aci/aci/mo/vz/vzConsLbl"
	"terraform-provider-aci/aci/mo/vz/vzConsSubjLbl"
	"terraform-provider-aci/aci/mo/vz/vzEntry"
	"terraform-provider-aci/aci/mo/vz/vzFilter"
	"terraform-provider-aci/aci/mo/vz/vzProvLbl"
	"terraform-provider-aci/aci/mo/vz/vzProvSubjLbl"
	"terraform-provider-aci/aci/mo/vz/vzSubj"
	"terraform-provider-aci/aci/mo/vz/vzRsSubjFiltAtt"

	// helpers
	//"terraform-provider-aci/aci/helpers/VMwareEpgPg"

)

// Provider : ACI terraform provider
func Provider() terraform.ResourceProvider {

	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				DefaultFunc: schema.EnvDefaultFunc("ACI_APIC_USERNAME", nil),
				Description: "Username to authenticate with ACI APIC",
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				DefaultFunc: schema.EnvDefaultFunc("ACI_APIC_PASSWORD", nil),
				Description: "Password to authenticate with ACI APIC",
			},
			"apic": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				DefaultFunc: schema.EnvDefaultFunc("ACI_APIC", nil),
				Description: "APIC IP Address",		// ** CHANGE TO LIST - dont need list as APIC cluster in readonly mode if one fails.
			},
			"allow_unverified_ssl": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				DefaultFunc: schema.EnvDefaultFunc("ACI_APIC_ALLOW_UNVERIFIED_SSL", true),
				Description: "If set, APIC client will permit unverifiable SSL certificates.",
			},
			"sync_delay": &schema.Schema{
				Type:     		schema.TypeInt,
				Optional: 		true,
				Default:		500,
				Description: 	"Time (ms) to wait for APIC DB cluster sync after a POST/DELETE call to the APIC - Default 500ms",
			},
		},

		// Map Resource to Function
		ResourcesMap: map[string]*schema.Resource{

			// fv
			"aci_fvAp":				fvAp.Resource(),
			"aci_fvAEPg":			fvAEPg.Resource(),
			"aci_fvBD":				fvBD.Resource(),
			"aci_fvCtx":			fvCtx.Resource(),
			"aci_fvRsBd":			fvRsBd.Resource(),
			"aci_fvRsCons":			fvRsCons.Resource(),
			"aci_fvRsCtx":			fvRsCtx.Resource(),
			"aci_fvRsDomAtt":		fvRsDomAtt.Resource(),
			"aci_fvRsProv":			fvRsProv.Resource(),
			"aci_fvSubnet":   		fvSubnet.Resource(),
			//"aci_fvTenant":			fvTenant.Resource(),

			// tag
			"aci_tagInst":   		tagInst.Resource(),

			// vz
			"aci_vzBrCP":			vzBrCP.Resource(),
			"aci_vzConsLbl":		vzConsLbl.Resource(),
			"aci_vzConsSubjLbl":	vzConsSubjLbl.Resource(),
			"aci_vzEntry":			vzEntry.Resource(),
			"aci_vzFilter":			vzFilter.Resource(),
			"aci_vzProvLbl":		vzProvLbl.Resource(),
			"aci_vzProvSubjLbl":	vzProvSubjLbl.Resource(),
			"aci_vzRsSubjFiltAtt":	vzRsSubjFiltAtt.Resource(),
			"aci_vzSubj":			vzSubj.Resource(),
	
		},

		// Map Data to Functions
		DataSourcesMap: map[string]*schema.Resource{

			// managed objects
			// query
			"aci_raw_query_mo":		query.Data(),

			// fv
			"aci_fvAp":				fvAp.Data(),
			"aci_fvAEPg":			fvAEPg.Data(),
			"aci_fvBD":				fvBD.Data(),
			"aci_fvCtx":			fvCtx.Data(),
			"aci_fvRsBd":			fvRsBd.Data(),
			"aci_fvRsCons":			fvRsCons.Data(),
			"aci_fvRsCtx":			fvRsCtx.Data(),
			"aci_fvRsDomAtt":		fvRsDomAtt.Data(),
			"aci_fvRsProv":			fvRsProv.Data(),
			"aci_fvSubnet":   		fvSubnet.Data(),
			"aci_fvTenant": 		fvTenant.Data(),

			// phys
			"aci_phys_dom":   		physDomP.Data(),

			// tag
			"aci_tagInst":   		tagInst.Data(),

			// vmm
			"aci_vmm_dom":   		vmmDomP.Data(),

			// vz
			"aci_vzBrCP":			vzBrCP.Data(),
			"aci_vzConsLbl":		vzConsLbl.Data(),
			"aci_vzConsSubjLbl":	vzConsSubjLbl.Data(),
			"aci_vzEntry":			vzEntry.Data(),
			"aci_vzFilter":			vzFilter.Data(),
			"aci_vzProvLbl":		vzProvLbl.Data(),
			"aci_vzProvSubjLbl":	vzProvSubjLbl.Data(),
			"aci_vzRsSubjFiltAtt":	vzRsSubjFiltAtt.Data(),
			"aci_vzSubj":			vzSubj.Data(),
			
			// helpers
			//"aci_vmware_portgroup":	helpers.Data_VMware_PortGroup(),
			
		},

		// Configure Client REST 
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	cookie, err := aci.Aci_login(d.Get("apic").(string), d.Get("username").(string), d.Get("password").(string))
	if err != nil {
		return nil, err
	}

	params := make(map[string]interface{})
	params["APIC-cookie"] = cookie
	params["sync_delay"] = d.Get("sync_delay").(int)
	params["allow_unverified_ssl"] = d.Get("allow_unverified_ssl").(bool)
	params["APIC"] = []string{ d.Get("apic").(string) }
	return params, nil
}
