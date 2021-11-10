package utils

import (
	"fmt"
	"errors"
	"encoding/json"
	"github.com/hashicorp/terraform/helper/schema"
	"terraform-provider-aci/aci/internal/meta"
	"strings"
	"log"
	"reflect"
)

// GetMOParentDn
// Strips the MO name from the DN returning the MO parent DN
//
func GetMOParentDn(dn string) (string, error) {

	in_embedded := false
	last_found_pos := -1
	for pos, char := range dn {
		if char == '/' && in_embedded == false {
			last_found_pos = pos
		} else if char == '[' {
			in_embedded = true
		} else if char == ']' {
			in_embedded = false
		}
	}
		
	if last_found_pos == -1 {
		return "", errors.New(fmt.Sprintf("Invalid DN to obtain parent DN. DN: %s", dn) )
	}

	log.Printf("Sep is %d", last_found_pos)
	log.Printf("Len is %d of %s", len(dn), dn)
	parent := dn[:last_found_pos]
	log.Printf("Parent DN is %s of %s", parent, dn)
	return parent, nil	
}

// GetMOSchemaKeys
// Returns a []string of the schema keys (MO struct field names) from the given MO struct
//
func GetMOSchemaKeys(mo_struct interface{}) []string {

	// returns a new Value initialized to the concrete value stored in the interface ApicQueryFilter
	s := reflect.ValueOf(mo_struct).Elem()

	// Value is the reflection interface to a Go value.
	typeOfT := s.Type()
	
	var keys []string
	// loop through fields in the struct, grab the schema name
	// this schema name is from the schema tag if present or from the 
	// field name in lowercase if schema tag is not present.
	for i := 0; i < s.NumField(); i++ {
		field := typeOfT.Field(i)
		field_key := field.Name

		tag, ok := field.Tag.Lookup("schema")
		log.Printf("[DEBUG] tag=%s, ok=%v", tag, ok)

		if tag, ok := field.Tag.Lookup("schema"); ok {
			log.Printf("[DEBUG] Using Schema Tag")
			if len(tag) == 0 {
				panic(fmt.Sprintf("PANIC: Struct %s has an empty schema tag for attribute %s - This MUST be fixed", typeOfT, field_key))
			}
			keys = append(keys, strings.ToLower(tag))
		} else {
			log.Printf("[DEBUG] Using Struct Field Name")
			keys = append(keys, strings.ToLower(field_key))
		}

	}
	return keys
}

/*
* Implements: Takes a pointer to a struct and maps the schema.ResourceData to the Struct elements
* Expects mo_struct as input e.g : bds := new(aci.BridgeDomain)
* This is basically a custom json.Marshal
*
* Returns: 
* Nothing: but sets the passed managed object struct fields with the matched attributes
* from the schema.ResourceData
*/
func SetMOStructValues(d *schema.ResourceData, mo_struct interface{}, config *meta.ManagedObjectMeta) { 

	log.Printf("[DEBUG] SetMOStructValues(..)")
	// returns a new Value initialized to the concrete value stored in the interface 
	s := reflect.ValueOf(mo_struct).Elem()
	log.Printf("[DEBUG] *** 's' Is: %s", s )

	typeOfT := s.Type()
	log.Printf("[DEBUG] *** Struct Type Is: %s", typeOfT )

	// loop through fields in the struct
	for i := 0; i < s.NumField(); i++ {

		field := typeOfT.Field(i)
		field_key := field.Name
		//json_tag, json_tag_ok := field.Tag.Lookup("json")
		schema_tag, schema_tag_ok := field.Tag.Lookup("schema")
		access_tag, _ := field.Tag.Lookup("access")

		// access tag=implicit means APIC read only attribute so dont set in struct
		// as this func set the struct values ready for marshall to payload for update
		// or create (APIC POSTs)
		if access_tag=="implicit" {
			log.Printf("[DEBUG] SetMOStructValues: Skipping %s 'implicit' attribute: %s", typeOfT, field_key)
			continue
		}

		//var schema_attribute_value string
		var schema_attribute_key string

		// TF schema key from schema tag
		if schema_tag_ok {

			schema_attribute_key = strings.Split(schema_tag, ",")[0]
			log.Printf("[DEBUG] 'schema_tag_ok' schema_tag=%s, schema_attribute_key=%s", schema_tag, schema_attribute_key)
			
			if len(schema_attribute_key) == 0 {
				panic(fmt.Sprintf("PANIC: Struct %s has an empty schema tag for attribute %s - This MUST be fixed", typeOfT, field_key))
			}
	
		} else {
			log.Printf("[DEBUG] Error in struct 'schema' tag, missing schema tag - using field name for schema attr key.")
			schema_attribute_key = strings.ToLower(field_key)
		}
		
		log.Printf("[DEBUG] *** Field Name Is: %s", schema_attribute_key)

		// we want to set a value in the struct only if either:
		// 1. The user has set a value in the TF file (resource data)
		// 2. An explicit default is set for the attribute in this provider
		// We DONT want to set a type default.

		schema_attribute_value, ok := getResourceDataValue(d, schema_attribute_key, config)
		if ok {
			log.Printf("[DEBUG] *** ValueIs: %v",schema_attribute_value )
			s.Field(i).SetString( (schema_attribute_value).(string) )
		}

		//schema_attribute_value, _ := d.GetOk(strings.ToLower( schema_attribute_key ))

		//if exists {
		//	log.Printf("[DEBUG] *** ValueIs: %s",schema_attribute_value )
			// TODO Needs switch statement for type if we use anything 
			// other than string values for APIC params
		//	s.Field(i).SetString( (schema_attribute_value).(string) )
		//}
	}
}

func getResourceDataValue(d *schema.ResourceData, schema_attribute_key string, config *meta.ManagedObjectMeta) (interface{}, bool) {
	
	log.Printf("[DEBUG] getResourceDataValue - Looking for a value for attribute: %s", schema_attribute_key)

	schema_attribute_value, ok := d.GetOk(strings.ToLower( schema_attribute_key ))
	if ok {
		// TF file has explicit value set
		log.Printf("[DEBUG] getResourceDataValue - Returning TF Set Value: %v", schema_attribute_value)
		return schema_attribute_value, true
	}

	// no explicit value set in TF file
	// if we have a explicit default set, then use this
	schema := config.SchemaF()
	attribute_schama, ok := schema[schema_attribute_key]
	if !ok {
		panic(fmt.Sprintf("\n\n[PANIC] getResourceDataValue - Key attribute [%s] does not exist in Schema %s. This must be fixed.\n\n", schema_attribute_key, config.ObjectClassName) )
	}

	default_value, _ := attribute_schama.DefaultValue()

	if default_value != nil {
		// we have an explicit default we can use
		log.Printf("[DEBUG] getResourceDataValue - Returning Explicit Default Value: %v", default_value)
		return default_value, true
	}

	// no value set by TF user and no explict default so return last error
	log.Printf("[DEBUG] getResourceDataValue - No usable value")
	return nil, false
}

/*
* Implements: Extracts the direct MO attributes from the given returned payload from the APIC
* from a GET/POST for an MO DN. 
* Generally called from GetMOAttributes
* Expects the correct full response payload from APIC GET MO DN
* Takes the raw APIC JSON response payload, converts to JSON object then checks the APIC response,
* then converts the given Managed Object attributes to a map[string]interface{}
*
* Returns: 
* map[string]interface{} as k,v for each of the MO direct attributes: { ... }
* Empty map[string]interface{} if no records returned with no error
*/
func ExtractMOAttributes(byte_data *[]byte, class_name string) (map[string]interface{}, error) {

	var json_data interface{}

	// from raw APIC response []byte string to json 
	// so json output has attributes names/case/etc as APIC definition
	err := json.Unmarshal(*byte_data, &json_data)
	if err != nil {
		return nil, err
	}

	record_count, ok := (json_data).(map[string]interface{})["totalCount"]
	if !ok {
		return nil, errors.New(fmt.Sprintf("[ExtractMOAttributes %s] APIC Context Request failed with with malformed response payload", class_name) )
	}

	if record_count == "0" {
		return map[string]interface{} {}, nil //, errors.New(fmt.Sprintf("[ExtractMOAttributes %s] Empty Response Payload For APIC MO. If this error occured on a create then the object may already exist on the APIC and not managed by this Terraform instance.", mo_name) )
	}

	// imdata is an array
	imdata, ok := (json_data).(map[string]interface{})["imdata"]
	if !ok {
		return nil, errors.New(fmt.Sprintf("[ExtractMOAttributes %s] APIC Context Request failed with with malformed response payload", class_name) )
	}

	// the only member of the imdata array is the classname of the DN returned.
	mo_attrs, ok := imdata.([]interface{})[0].(map[string]interface{})[class_name]
	if !ok {
		return nil, errors.New(fmt.Sprintf("[ExtractMOAttributes %s] APIC Context Request failed with with malformed response payload", class_name) )
	}

	attributes, ok := mo_attrs.(map[string]interface{})["attributes"]
	if !ok {
		return nil, errors.New(fmt.Sprintf("[ExtractMOAttributes %s] APIC Context Request failed with with malformed response payload", class_name) )
	}

	kv, ok  := attributes.(map[string]interface{})
	if !ok {
		return nil, errors.New(fmt.Sprintf("[ExtractMOAttributes %s] APIC Context Request failed with with malformed response payload", class_name) )
	}

	return kv, nil
}

// return the APIC MO object attributes as a []byte string from the schemadata and using the specified struct
// this []byte string is used as raw payload for the APIC POST.
func CreateMOAttributesPayload(d *schema.ResourceData, mo_struct interface{}, config *meta.ManagedObjectMeta) ([]byte, error) {

	// Uses the pointer to the passed struct to add the TF schema values to the struct
	// maps the keys from both sides (TF and APIC) as case is different
	SetMOStructValues(d, mo_struct, config)

	// create payload
	attributes, err := json.Marshal(mo_struct)
	if err != nil {
		return nil, err
	}

	return attributes, nil
}

// SetResourceDataValues
//
// Takes the APIC MO attributes in map[string]string format and sets the schema.ResourceData
// attribute values from this map.
//
// Input:  	d *schema.ResourceData 	- The TF ResourceData to be set with values from src_key_val_map
//			src_key_val_map    		- map[string]string - key:attribute name, value:attribite value to set in Schema
//			src_schema_key_map 		- map[string]string - key: key name used in the src_key_val_map key, value: the TF schema key to use to set with src_key_val_map value
//
// Output: schema.ResourceData.attributes.value set as input key value
func SetResourceDataValues(d *schema.ResourceData, src_key_val_map map[string]interface{}, src_schema_key_map map[string]string) {

	// iterate through the provided key/values attributes
	// if we have a match of the keys in src_key_val_map and src_schema_key_map
	// then;
	// Set the schema key (key defined in src_schema_key_map VALUE)
	// with value from src_key_val_map value.
	for key, value := range src_key_val_map {

		var schema_key string
		var ok bool
		if schema_key, ok = src_schema_key_map[key]; !ok {
			schema_key = key
		}
		if len(schema_key) == 0 {
			schema_key = key
		}


		log.Printf("[DEBUG] Reading k,v k=%s, v=%s", schema_key, value)
		// expects the attributes maps keys to be the same name as the schema names but poss. different case
		// convert the keys in the attributes map to lowercase to match the TF schema name
		// as TF schema attribute case is always lowercase.
		// check the attribute key exists in the schema.ResourceData
		////_, exists := d.GetOk( strings.ToLower(schema_key) )			
		// if a match between schema.ResourceData.attribute.key and attributes.key 
		// then set the schema.ResourceData.attribute.value
		////if exists {
			log.Printf("[DEBUG] Setting schema.ResourceData value: %v on schema key: %s from source key:%s", value, schema_key, key)
			d.Set(strings.ToLower(schema_key), value)
		////}
	}
}

//
// The MO Struct has the following attr keys
// field name: Struct Field Name, capitilised first letter but otherwise same case as APIC attr key
// tag=>json: Has the key name of the APIC attr key is same case etc as APIC
// tag=>schema: Has the name that is used in the TF schema 
//
// Input: MO Struct
// Output: map[string]string = APIC_MO_ATTR_KEY: SCHEMA_MO_ATTR_KEY
func MapApicKeyToSchemaKey(mo_struct interface{}) map[string]string {

	// returns a new Value initialized to the concrete value stored in the interface ApicQueryFilter
	s := reflect.ValueOf(mo_struct).Elem()

	// Value is the reflection interface to a Go value.
	typeOfT := s.Type()
	log.Printf("[DEBUG] *** Struct Type Is: %s", typeOfT )

	keys := make(map[string]string)
	// loop through fields in the struct, grab the schema name
	// this schema name is from the schema tag if present or from the 
	// field name in lowercase if schema tag is not present.
	for i := 0; i < s.NumField(); i++ {

		field := typeOfT.Field(i)
		field_key := field.Name
		log.Printf("[DEBUG] field_key=%s", field_key)
		
		var map_apic_key string
		var map_schema_key string


		// apic key from json tag
		if apic_key, apic_ok := field.Tag.Lookup("json"); apic_ok {
			log.Printf("[DEBUG] apic_key=%s, ok=%v", apic_key, apic_ok)

			map_apic_key = strings.Split(apic_key, ",")[0]
			//log.Printf("[DEBUG] json apic_key=%s idx[0]=%s", strings.Split(apic_key, ","), map_apic_key)
			if len(map_apic_key) == 0 {
				// missing first value in json tag so use the field name with lower case first char
				log.Printf("[DEBUG] Error in struct 'json' tag, missing value - using field name with lower case first char as apic key.")
				map_apic_key = strings.Join(  []string{ strings.ToLower(field_key[:1]), field_key[1:len(field_key)] } , "" )
			}
	
		} else {
			// missing json tag so use the field name with lower case first char
			log.Printf("[DEBUG] Error in struct 'json' tag, missing json tag - using field name with lower case first char as apic key.")
			map_apic_key = strings.Join(  []string{ strings.ToLower(field_key[:1]), field_key[1:len(field_key)] } , "" )
		}


		// TF schema key from schema tag
		if schema_key, schema_ok := field.Tag.Lookup("schema"); schema_ok {
			log.Printf("[DEBUG] schema_key=%s, ok=%v", schema_key, schema_ok)

			map_schema_key = strings.Split(schema_key, ",")[0]
			if len(map_schema_key) == 0 {
				panic(fmt.Sprintf("PANIC: Struct %s has an empty schema tag for attribute %s - This MUST be fixed", typeOfT, field_key))
				map_schema_key = strings.ToLower(field_key)
			}
	
		} else {
			log.Printf("[DEBUG] Error in struct 'schema' tag, missing schema tag - using field name for schema attr key.")
			map_schema_key = strings.ToLower(field_key)
		}

		// add to map
		log.Printf("[DEBUG] Setting Key[map_apic_key]: %s to value[map_schema_key] %s", map_apic_key, map_schema_key)
		keys[map_apic_key] = map_schema_key
	}
	return keys
}