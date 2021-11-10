package fvSubnet

import ( 
	"github.com/hashicorp/terraform/helper/schema"
	_"github.com/hashicorp/terraform/helper/validation"
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

		ObjectClassName:        "fvSubnet",
		ObjectNamePrefix:       "subnet",
		ObjectNameFieldName:	"ip",
		ObjectParentFieldName:  "parent_id",
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
	return new(fv.Subnet)
}

// specific formatting for object name i.e. subnet has [..] brackets around ip for name
func formatObjectDnName(d *schema.ResourceData, ObjectNamePrefix string, ObjectName string) string {
	return fmt.Sprintf("%s-[%s]", ObjectNamePrefix, ObjectName)
}

func getObjectSchema() map[string]*schema.Schema {

	return map[string]*schema.Schema{

		"parent_id": &schema.Schema{
				Type:        	schema.TypeString,
				Required:    	true,
				ForceNew:		true,
				Description: 	"ACI Subnet - Containing MO Parent ID (DN)",
				//ValidateFunc: util.ValidateMaxLength(256),
			},
			"ctrl": &schema.Schema{
				Type:        		schema.TypeString,
				Optional:    		true,
				Description: 		"ACI Subnet - The subnet control state. The control can be specific protocols applied to the subnet such as IGMP Snooping.",
				Default:  			"unspecified",
				ValidateFunc:		utils.ValidateList([]string{"nd","unspecified","querier","no-default-gateway"}),
				DiffSuppressFunc:	utils.SuppressDiffCheck("unspecified", ""),

			},
			"desc": &schema.Schema{
				Type:        	schema.TypeString,
				Optional:    	true,
				Description: 	"ACI Subnet - Description",
				//ValidateFunc:   utils.ValidateMaxLength(128),
			},
			"dn": &schema.Schema{
				Type:        	schema.TypeString,
				Computed:    	true,
				Description: 	"ACI Subnet - DN",
			},
			"ip": &schema.Schema{
				Type:        	schema.TypeString,
				Required:    	true,
				ForceNew: 	 	true,
				Description: 	"ACI Subnet - The IP address and mask of the default gateway in d.d.d.d/dd format.",
				//ValidateFunc: util.ValidateMaxLength(256),
			},
			"name": &schema.Schema{
				Type:        	schema.TypeString,
				Optional:    	true,
				Description: 	"ACI Subnet - Subnet Name",
				//ValidateFunc: util.ValidateMaxLength(256), (no, yes)
			},
			"ownerkey": &schema.Schema{
				Type:        	schema.TypeString,
				Optional:    	true,
				Description: 	"ACI Subnet Owner Key - Arbitary key for enabling clients to own their data for entity correlation.",
			},
			"ownertag": &schema.Schema{
				Type:        	schema.TypeString,
				Optional:    	true,
				Description: 	"ACI Subnet Owner Tag - A tag for enabling clients to add their own data. For example, to indicate who created this object.",
			},
			"preferred": &schema.Schema{
				Type:        	schema.TypeString,
				Optional:    	true,
				Description: 	"ACI Subnet - Indicates if the subnet is preferred (primary) over the available alternatives. Only one preferred subnet is allowed.",
				Default:  		"no",
				ValidateFunc:	utils.ValidateList([]string{"yes","no"}),
			},
			"scope": &schema.Schema{
				Type:        	schema.TypeString,
				Optional:    	true,
				Description: 	"ACI Subnet - Control whether BUM traffic is allowed between sites (Private to VRF (private), Advertised Externally (public), Shared Between VRFs (shared))",
				Default:  		"private",
				ValidateFunc:	utils.ValidateList([]string{"private","public","shared"}),
			},
			"virtual": &schema.Schema{
				Type:        	schema.TypeString,
				Optional:    	true,
				Description: 	"ACI Subnet - Treated as virtual IP address. Used in case of BD extended to multiple sites.",
				Default:  		"no",
				ValidateFunc:	utils.ValidateList([]string{"yes","no"}),
			},
		}
}