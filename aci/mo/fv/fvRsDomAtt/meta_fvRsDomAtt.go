package fvRsDomAtt
//
// An EPG can be linked to a domain profile via the Associated Domain Profiles. 
// The domain profiles attached can be VMM, Physical, L2 External, or L3 External Domains.  
//
import ( 
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"terraform-provider-aci/aci/internal/utils"
	"terraform-provider-aci/aci/internal/meta"
	"terraform-provider-aci/aci/mo/fv"
	"fmt"
)

func getObjectConfig() *meta.ManagedObjectMeta {

	return &meta.ManagedObjectMeta {

		ObjectClassName:        "fvRsDomAtt",
		ObjectNamePrefix:       "rsdomAtt",
		ObjectNameFieldName:	"domain_profile_id",
		ObjectParentFieldName:  "epg_id",
		OverWriteExisting:      false,
		IsReln:					true,
		IsMandatoryReln:		false,
		CanCreate:				true,
		CanRead:				true,
		CanUpdate:				true,
		CanDelete:				true,
	
		ObjectModelF:           getObjectModel,
		SchemaF:                getObjectSchema,
		FormatObjectNameF:      formatObjectDnName,		
	}
}

func getObjectModel(d *schema.ResourceData) interface{} {
	return new(fv.RsDomAtt)
}

// specific formatting for object name i.e. fvRsDomAtt has [..] brackets around tDn for name
func formatObjectDnName(d *schema.ResourceData, ObjectNamePrefix string, ObjectName string) string {
	return fmt.Sprintf("%s-[%s]", ObjectNamePrefix, ObjectName )
}

func getObjectSchema() map[string]*schema.Schema {

	return map[string]*schema.Schema{

		"epg_id": &schema.Schema{
			Type:        	schema.TypeString,
			Required:    	true,
			ForceNew: 		true,
			Description: 	"ACI EPG Domain Profile Reln - Containing EPG ID/DN",
		},
		"class_pref": &schema.Schema{
			Type:        		schema.TypeString,
			Optional:    		true,
			Description: 		"ACI EPG Domain Profile Reln - Class Pref. (encap, useg)",
			Default: 			"encap",
			ValidateFunc:		utils.ValidateList([]string{"encap","useg"}),
		},
		"delimiter": &schema.Schema{
			Type:        		schema.TypeString,
			Optional:    		true,
			Description: 		"ACI EPG Domain Profile Reln - VMM Network Seperator Character",
			Default: 			"|",
			ValidateFunc:		validation.StringLenBetween(0, 1),
			//DiffSuppressFunc:	utils.SuppressDiffCheck("|", ""),
			StateFunc: 			func(val interface{}) string {
									if val.(string) == "" {
										return "|"
									}
									return val.(string)
								},
		},
		"dn": &schema.Schema{
			Type:        	schema.TypeString,
			Computed:    	true,
			Description: 	"ACI EPG Domain Profile Reln - DN",
		},	
		"port_encapsulation": &schema.Schema{
			Type:        		schema.TypeString,
			Optional:    		true,
			Description: 		"ACI EPG Domain Profile Reln - Encapsulation",
			Default: 			"unknown",
		},
		"encapsulation_mode": &schema.Schema{
			Type:        		schema.TypeString,
			Optional:    		true,
			Description: 		"ACI EPG Domain Profile Reln - Enable Logging",
			Default: 			"auto",
			ValidateFunc:		utils.ValidateList([]string{"auto","vxlan","vlan"}),
		},
		"epg_cos_value": &schema.Schema{
			Type:        		schema.TypeString,
			Optional:    		true,
			Description: 		"ACI EPG Domain Profile Reln - EPG CoS Value",
			Default: 			"Cos0",
			ValidateFunc:		utils.ValidateList([]string{"Cos0","Cos1","Cos2","Cos3","Cos4","Cos5","Cos6","Cos7"}),
		},
		"epg_cos_enable": &schema.Schema{
			Type:        		schema.TypeString,
			Optional:    		true,
			Description: 		"ACI EPG Domain Profile Reln - Enable Logging",
			Default: 			"disabled",
			ValidateFunc:		utils.ValidateList([]string{"enabled","disabled"}),
		},
		"policy_deployment_mode": &schema.Schema{
			Type:        		schema.TypeString,
			Optional:    		true,
			Description: 		"ACI EPG Domain Profile Reln - Policy Deployment Mode (Deployment Immmediacy)",
			Default: 			"lazy",
			ValidateFunc:		utils.ValidateList([]string{"immediate","lazy"}),
		},
		"netflow_direction": &schema.Schema{
			Type:        		schema.TypeString,
			Optional:    		true,
			Description: 		"ACI EPG Domain Profile Reln - Netflow Direction",
			Default: 			"both",
			ValidateFunc:		utils.ValidateList([]string{"ingress","egress","both"}),
		},
		"netflow_enable": &schema.Schema{
			Type:        		schema.TypeString,
			Optional:    		true,
			Description: 		"ACI EPG Domain Profile Reln - Enable Logging",
			Default: 			"disabled",
			ValidateFunc:		utils.ValidateList([]string{"enabled","disabled"}),
		},
		"primary_encap": &schema.Schema{
			Type:        		schema.TypeString,
			Optional:    		true,
			Description: 		"ACI EPG Domain Profile Reln - Enable Logging",
			Default: 			"",
			ValidateFunc:		utils.ValidateList([]string{"nd","unspecified"}),
			DiffSuppressFunc:	utils.SuppressDiffCheck("unknown", ""),
		},
		"primary_encap_inner": &schema.Schema{
			Type:        		schema.TypeString,
			Optional:    		true,
			Description: 		"ACI EPG Domain Profile Reln - Primary Encapsulation",
			Default: 			"unknown",
			//ValidateFunc:		utils.ValidateList([]string{"nd","unspecified"}),
			DiffSuppressFunc:	utils.SuppressDiffCheck("unknown", ""),
		},
		"policy_resolution_mode": &schema.Schema{
			Type:        		schema.TypeString,
			Optional:    		true,
			Description: 		"ACI EPG Domain Profile Reln - Enable Logging",
			Default: 			"immediate",
			ValidateFunc:		utils.ValidateList([]string{"immediate","lazy","pre-provision"}),
		},
		"second_encap_inner": &schema.Schema{
			Type:        		schema.TypeString,
			Optional:    		true,
			Description: 		"ACI EPG Domain Profile Reln - Secondary Encapsulation Inner",
			Default: 			"unknown",
			//ValidateFunc:		utils.ValidateList([]string{"nd","unspecified"}),
			DiffSuppressFunc:	utils.SuppressDiffCheck("unknown", ""),
		},
		"switching_mode": &schema.Schema{
			Type:        		schema.TypeString,
			Optional:    		true,
			Description: 		"ACI EPG Domain Profile Reln - Enable Logging",
			Default: 			"native",
			ValidateFunc:		utils.ValidateList([]string{"native","AVE"}),
		},
		"domain_profile_id": &schema.Schema{
			Type:        	schema.TypeString,
			Required:    	true,
			ForceNew: 		true,
			Description: 	"ACI EPG Domain Profile Reln - Domain Profile ID to Attach To The EPG (tDn)",
		},
		
	}
}