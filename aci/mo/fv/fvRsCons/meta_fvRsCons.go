package fvRsCons
//
// The filter for the subject of a service contract. 
// A subject represents a sub-application running behind an endpoint group, such as an exchange server. 
// A subject is parented by the contract. 
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

		ObjectClassName:        "fvRsCons",
		ObjectNamePrefix:       "rscons",
		ObjectNameFieldName:	"contract_name",
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
	return new(fv.RsCons)
}

// specific formatting for object name i.e. subnet has [..] brackets around ip for name
func formatObjectDnName(d *schema.ResourceData, ObjectNamePrefix string, ObjectName string) string {
	return fmt.Sprintf("%s-%s", ObjectNamePrefix, ObjectName )
}

func getObjectSchema() map[string]*schema.Schema {

	return map[string]*schema.Schema{

		"epg_id": &schema.Schema{
			Type:        	schema.TypeString,
			Required:    	true,
			ForceNew: 		true,
			Description: 	"ACI EPG Consumer Contract Reln - Containing EPG ID (DN)",
			ValidateFunc: 	validation.StringLenBetween(1, 255),
		},
		"dn": &schema.Schema{
			Type:        	schema.TypeString,
			Computed:    	true,
			Description: 	"ACI EPG Consumer Contract Reln - DN",
		},		
		"priority_level": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Description: "ACI EPG Consumer Contract Reln - QoS Priority Level [prio]",
			Default:  "unspecified",
			ValidateFunc:	utils.ValidateList([]string{"unspecified","level1", "level2", "leve3"}),
		},
		"contract_name": &schema.Schema{
			Type:        	schema.TypeString,
			Required:    	true,
			ForceNew: 		true,
			Description: 	"ACI EPG Consumer Contract Reln - Contract relative name [tnVzBrCPName]",
			ValidateFunc: 	validation.StringLenBetween(1, 64),
		},
	}
}