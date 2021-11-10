package fvAEPg

import ( 
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"terraform-provider-aci/aci/internal/utils"
	"terraform-provider-aci/aci/internal/meta"
	"terraform-provider-aci/aci/mo/fv"
	"fmt"
)

/*
ClassName is the ACI Managed Object Abstract Class Name
DnPrefix is the ACI Managed Object Relative Name Prefix
ParentDnName is the Terraform Provider Schema Name Used To Identify The Parent Object, internally this is the ACI MO Parent object but we may 
			 call it something different in the schema to make it clear to TF users
ObjectNameAttribute is the Schema Name that is the relative name that ACI uses primarly for the naming of an object relativly.
OverWriteExisting is used when an object exists by default such as a rsctx when a BD is created but we need to manage the object, so 
			 this flag stop the check of existing mo and takes management of it and makes the create function like a update and take management func.
*/

func getObjectConfig() *meta.ManagedObjectMeta {

	return &meta.ManagedObjectMeta {

		ObjectClassName:        "fvAEPg",
		ObjectNamePrefix:       "epg",
		ObjectNameFieldName:	"name",
		ObjectParentFieldName:  "ap_id",
		OverWriteExisting:      false,
		IsReln:					false,
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
	return new(fv.AEPg)
}

// specific formatting for object name i.e. subnet has subnet-[..ip..] brackets around ip for name
func formatObjectDnName(d *schema.ResourceData, ObjectNamePrefix string, ObjectName string) string {
	return fmt.Sprintf("%s-%s", ObjectNamePrefix, ObjectName )
}

func getObjectSchema() map[string]*schema.Schema {

	return map[string]*schema.Schema{

		"ap_id": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,						
			Description: "ACI EPG Parent Application Profile ID.",
			//ValidateFunc:	validation.StringLenBetween(1, 255),
		},
		"desc": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Description: "ACI EPG Description.",
			ValidateFunc:	validation.StringLenBetween(1, 128),
		},
		"dn": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
			Description: "ACI EPG DN.",
			//ValidateFunc: util.ValidateMaxLength(256),
		},
		"flood_on_encap": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI EPG - Specify whether unknown traffic sourced from this EPG is flooded only with the EPG or throughout the Bridge Domain (Default=disabled=flood_bd).",
			Default:  		"disabled",
			ValidateFunc:	utils.ValidateList([]string{"enabled","disabled"}),
		},
		"forward_control": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Description: "ACI EPG Forwarding Control - Only used if Intra-EPG Isolation is Enforced. Allows the enabling of proxy-arp for Intra EPG Isolation. ",
			Default:  	 "",
			ValidateFunc:	utils.ValidateList([]string{"","proxy-arp"}),
			DiffSuppressFunc:	utils.SuppressDiffCheck("none", ""),
		},
		"is_attr_based_epg": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Description: "ACI EPG Is Attribute Based EPG",
			Default:  "no",
			ValidateFunc:	utils.ValidateList([]string{"yes","no"}),
		},
		"label_match_criteria": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Description: "ACI EPG - The provider label match criteria.",
			Default:  "AtleastOne",
			ValidateFunc:	utils.ValidateList([]string{"AtleastOne","All", "AtmostOne", "None"}),
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
			Description: "ACI EPG Name.",
			ValidateFunc:	validation.StringLenBetween(1, 64),
		},
		"intra_epg_isolation": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI EPG The preferred policy control - Intra-EPG endpoint isolation policies provide full isolation for virtual or physical endpoints;\nno communication is allowed between endpoints in an EPG that is operating with isolation enforced.",
			Default:  		"unenforced",
			ValidateFunc:	utils.ValidateList([]string{"unenforced","enforced"}),
		},
		"preferred_group_member": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Description: "ACI EPG - Prefered Group Member - Including this EPG in the Preferred Group Membership means this and other EPGs with the same membership setting and in the same VRF can communicate without contracts.",
			Default:  "exclude",
			ValidateFunc:	utils.ValidateList([]string{"exclude","include"}),
		},
		"priority_level": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Description: "ACI EPG QoS Priority Level",
			Default:  "unspecified",
			ValidateFunc:	utils.ValidateList([]string{"unspecified","level1", "level2", "leve3"}),
		},
	}
}