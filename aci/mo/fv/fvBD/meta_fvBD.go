package fvBD

import ( 
	"regexp"
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

		ObjectClassName:        "fvBD",
		ObjectNamePrefix:       "BD",
		ObjectNameFieldName:	"name",
		ObjectParentFieldName:  "tenant_id",
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
	return new(fv.BD)
}

// specific formatting for object name i.e. subnet has [..] brackets around ip for name
func formatObjectDnName(d *schema.ResourceData, ObjectNamePrefix string, ObjectName string) string {
	return fmt.Sprintf("%s-%s", ObjectNamePrefix, ObjectName )
}

func getObjectSchema() map[string]*schema.Schema {

	return map[string]*schema.Schema{

		"tenant_id": &schema.Schema{
			Type:        	schema.TypeString,
			Required:    	true,
			ForceNew: 		true,
			Description: 	"ACI Bridge Domain - Containing Tenant ID (DN)",
			ValidateFunc: 	validation.StringLenBetween(1, 255),
		},
		"desc": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Bridge Domain - Description",
		},
		"dn": &schema.Schema{
			Type:        	schema.TypeString,
			Computed:    	true,
			Description: 	"ACI Bridge Domain - DN",
		},				
		"arp_flooding": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Bridge Domain - ARP Flooding ",
			//Default:		"no",
			DiffSuppressFunc:	utils.SuppressDiffCheck("no", ""),
		},
		"multicast_group_ipv4": &schema.Schema{
			Type:        	schema.TypeString,
			Computed:    	true,
			Description: 	"ACI Bridge Domain - Outer multicast group IP address",
		},
		"endpoint_clear": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Bridge Domain - Clear all EPs in all leaves for this BD",
			DiffSuppressFunc:	utils.SuppressDiffCheck("no", ""),
		},
		"garp_move_detection": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Bridge Domain - End Point move detection option uses the Gratuitous Address Resolution Protocol (GARP).",
		},
		"host_based_routing": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Bridge Domain - Enables advertising host routes (/32 prefixes) out of the L3OUT(s) that are associated to this BD.",
			//Default:		"no",
			ValidateFunc:	utils.ValidateList([]string{"yes","no"}),
		},
		"intersite_bum_enabled": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Bridge Domain - Control whether BUM traffic is allowed between sites",
			Default:		"no",
		},
		"intersite_l2stretch": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Bridge Domain - l2Stretch flag is enabled between sites",
			Default:		"no",
		},
		"ip_learning": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Bridge Domain - IP Learning",
			Default:		"yes",
		},
		"limit_ip_learn_to_subnet": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:   	 true,
			Description: 	"ACI Bridge Domain - Limits IP address learning to the bridge domain subnets only. By default, all IPs are learned",
			Default:		"yes",
		},
		"ipv6_local_link_addr": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Bridge Domain - The override of the system generated IPv6 link-local address",
			DiffSuppressFunc:	utils.SuppressDiffCheck("::", ""),
		},
		"bridge_domain_mac": &schema.Schema{
			Type:        		schema.TypeString,
			Optional:    		true,
			Description: 		"ACI Bridge Domain - The MAC address of the bridge domain (BD) or switched virtual interface (SVI).",
			DiffSuppressFunc:	utils.SuppressDiffCheck("00:22:BD:F8:19:FF", ""),
		},
		"allow_multicast": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Bridge Domain - Flag to indicate if multicast is enabled",
			Default:		"no",
		},
		"max_l2_mtu": &schema.Schema{
			Type:        	schema.TypeString,
			Computed:    	true,
			Description: 	"ACI Bridge Domain - The layer 2 maximum transmit unit (MTU) size.",
		},
		"l2_forwarding_method": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Bridge Domain - The multiple destination forwarding method for L2 Multicast, Broadcast, and Link Layer traffic types.",
			Default:		"bd-flood",
		},
		"intersite_bw_optimize": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Bridge Domain - OptimizeWanBandwidth flag is enabled between sites",
			Default:		"no",
		},
		"name": &schema.Schema{
			Type:        	schema.TypeString,
			Required:    	true,
			ForceNew: 		true,
			ValidateFunc: 	validation.StringMatch(regexp.MustCompile("^[a-zA-Z0-9_.-]+$"), "Valid characters are: [a-zA-Z0-9_.-]+"),
			Description: 	"ACI Bridge Domain - Name",
		},
		"ownerkey": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Bridge Domain - Arbitary key for enabling clients to own their data for entity correlation.",
		},
		"ownertag": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Bridge Domain - A tag for enabling clients to add their own data. For example, to indicate who created this object.",
		},
		"bridge_domain_type": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Bridge Domain - The Bridge Domain type is regular or fc.",
			Default:		"regular",
		},
		"unicast_routing_enabled": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Bridge Domain - The forwarding method based on predefined forwarding criteria (IP or MAC address). (Unicast Routing Enable)",
			Default:		"yes",
		},
		"unknown_l2_forward_method": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Bridge Domain - The forwarding method for unknown layer 2 destinations.",
			Default:		"proxy",
		},
		"unknown_mcast_forward_method": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Bridge Domain - Method for forwarding data for an unknown multicast destination.",
			Default:		"flood",
		},
		"l2_out_virtual_mac": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Description: "ACI Bridge Domain - Virtual MAC address of the BD/SVI. Used when the BD is extended to multiple sites using l2 Outside",
			DiffSuppressFunc:	utils.SuppressDiffCheck("not-applicable", ""),
		},		
	}
}

