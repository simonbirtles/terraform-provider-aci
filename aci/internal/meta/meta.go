package meta

import (
    "github.com/hashicorp/terraform/helper/schema"
)

type GetObjectModelFunc     func(d *schema.ResourceData) interface{}
type GetSchemaFunc          func() map[string]*schema.Schema
type FormatObjectNameFunc   func(d *schema.ResourceData, ObjectNamePrefix string, ObjectName string) string

// ManagedObjectMeta holds additional data about the APIC managed object
// that this schema applies to. This additional data is required and no 
// fields are optional.
// Refer to the APIC MIM https://developer.cisco.com/site/apic-mim-ref-api/
// for additional infomation to complete this struct.
//
type ManagedObjectMeta struct {

	// The APIC managed object concrete class name as per the MIM
	// https://developer.cisco.com/site/apic-mim-ref-api/
	// Such as fvBD, fvRsCtx, fvTenant.
	ObjectClassName         string
	
	// The APIC managed object naming prefix as per the MIM 
	// https://developer.cisco.com/site/apic-mim-ref-api/
	// Such as ctx for fvCtx, tn for fvTenant.
	ObjectNamePrefix        string

	// The name of the schema attribute that is used to name the APIC
	// managed object. The APIC generally uses the "name" attribute but
	// cases like the fvSubnet, the "ip" attribute is used to name the 
	// attribute. This field name specified the schema field to generate
	// the APIC managed object.
	ObjectNameFieldName		string	

	// The name of the schema attribute that has the APIC managed objects
	// parent DN. This is the APIC managed object DN that this APIC 
	// managed object will be created under.
	ObjectParentFieldName   string

	// There are cases where an APIC managed object is created and children
	// of that object are automatically created. These usually cannot be 
	// manually created or deleted and therefore cant be created by terraform.
	// So to allow terraform to ignore the fact the object exists and to run 
	// the create function to update the object and take ownership, this flag
	// allows the create function to bypass the existence check and just update
	// the object.
	OverWriteExisting       bool

	// True if this APIC managed object is a relationship object, usually prefixed 
	// with 'rs'. 'rt' objects are not created by this Terraform provider.
	IsReln					bool

	// Refering back to the 'OverWriteExisting' flag comments, this flag states
	// that this object will be created by the APIC automatically as its a 
	// mandatory child of the parent managed object.
	IsMandatoryReln			bool

	// True if the object can be created via the APIC REST API.
	// Some managed objects are created automatically by the APIC and cannot be 
	// manually created. Some cases as with 'IsMandatoryReln' and 'OverWriteExisting'
	// do not allow objects to be created. Refer to the APIC MIM 'Creatable/Deletable'
	// section for each object. https://developer.cisco.com/site/apic-mim-ref-api/
	CanCreate				bool

	// True if the object can be read via the APIC REST API.
	// No cases known that this cannot happen but for consistency.
	CanRead					bool

	// True if the object can be updates via the APIC REST API.
	// No cases known that this cannot happen but for consistency.
	CanUpdate				bool

	// True if the object can be deleted via the APIC REST API.
	// Some managed objects are created automatically by the APIC and cannot be 
	// manually created or deleted. Some cases as with 'IsMandatoryReln' and 
	// 'OverWriteExisting' do not allow objects to be created. Refer to the APIC 
	// MIM 'Creatable/Deletable' section for each object. 
	// https://developer.cisco.com/site/apic-mim-ref-api/
	CanDelete				bool

	// Function returns a pointer to an instantiated struct representing the 
	// APIC managed object model.
	ObjectModelF            GetObjectModelFunc
	
	// Function to return the Terraform Schema for this object.
	SchemaF                 GetSchemaFunc
	
	// Function to return the formatted APIC managed object name.
	// Naming conventions vary, mostly are 'prefix-name' but can 
	// be different as per fvSubnet 'subnet-[x.x.x.x/y]' so 
	// allows the specific formatting to be applied per object.
    FormatObjectNameF       FormatObjectNameFunc
}
