package api

import (
	"fmt"
	"errors"
	"log"
	//"encoding/json"
	"github.com/hashicorp/terraform/helper/schema"
	"terraform-provider-aci/aci/internal/utils"
	"aci"
)

func exists(dn string, m interface{}, class_name string) (bool, error) {

	log.Printf("[DEBUG] Checking existance of ACI Resource DN: %s", dn)

	data, err := Read(m, dn)
	if err != nil {
		return false, err
	}

	attrs, err := utils.ExtractMOAttributes(&data, class_name)
	if err != nil {
		return false, err
	}

	// exists if attrs dn matches
	if attrs["dn"] == dn {
		log.Printf("[DEBUG] ACI Resource EXISTS DN: %s ", dn)
		return true, nil
	}

	log.Printf("[DEBUG] ACI Resource DOES NOT EXIST DN: %s ", dn)	
	return false, nil
}

func delete(m interface{}, dn string) error {
	
	apicCookie, ok := m.(map[string]interface{})["APIC-cookie"].(string) 
	if !ok{
		return errors.New("Unable to extract apic-cookie from ResourceData")
	}

	apics, ok := m.(map[string]interface{})["APIC"].([]string)
	if !ok{
		return errors.New("Unable to extract apic list from ResourceData")
	}

	path := fmt.Sprintf("mo/%s", dn)

	var info = new(aci.ApicDeleteInfo)
	info.ApicClient.Cookie = apicCookie
	info.ApicClient.ApicHosts = apics
	info.Path = path
	info.Delay = m.(map[string]interface{})["sync_delay"].(int)
	err := aci.Delete(info)
	if err != nil {
		return err
	}

	return nil
}

func post(m interface{}, parent_dn string, class_name string, attributes []byte) ([]byte, error) {

	apicCookie, ok := m.(map[string]interface{})["APIC-cookie"].(string) 
	if !ok{
		return nil, errors.New("Unable to extract apic-cookie from ResourceData")
	}

	apics, ok := m.(map[string]interface{})["APIC"].([]string)
	if !ok{
		return nil, errors.New("Unable to extract apic list from ResourceData")
	}
		
	log.Printf("[DEBUG] Posting attributes: %s", attributes )
	var postinfo = new(aci.ApicPostInfo)
	postinfo.ApicClient.ApicHosts =  apics
	postinfo.ApicClient.Cookie = apicCookie
	postinfo.Path = fmt.Sprintf("mo/%s.json", parent_dn ) 
	postinfo.Filter.Rsp_subtree = "full" // was modified
	//postinfo.Filter.Rsp_prop_include = "all"
	postinfo.Delay = m.(map[string]interface{})["sync_delay"].(int)
	postinfo.Payload = []byte(fmt.Sprintf(`
	{
		"%s": {
			"attributes":%s
		}
	}`,
	class_name, attributes ) )
		
    log.Printf("[DEBUG] - %s Payload:%s", class_name, fmt.Sprintf("%s", postinfo.Payload) )
    
    response, err := aci.Post(postinfo)
	if err != nil {
		return response, err
	}

	return response, nil
}

//
// returns raw payload attributes for a given DN
// Use utils.GetMOAttributes for map[string] return of attributes
//
func get(m interface{}, dn string) ([]byte, error) {

	apicCookie, ok := m.(map[string]interface{})["APIC-cookie"].(string) 
	if !ok{
		return nil, errors.New("Unable to extract apic-cookie from ResourceData")
	}

	apics, ok := m.(map[string]interface{})["APIC"].([]string)
	if !ok{
		return nil, errors.New("Unable to extract apic list from ResourceData")
	}

	path := fmt.Sprintf("mo/%s", dn)

	var info = new(aci.ApicGetInfo)
	info.ApicClient.Cookie = apicCookie
	info.ApicClient.ApicHosts = apics
	info.Path = path
	info.Delay = 0 //m.(map[string]interface{})["get_delay"].(int)

	log.Printf("[DEBUG] Get ACI Resource - %s", info.Path)
	
	response, err := aci.Get(info)
	if err != nil {
		log.Printf("[DEBUG] Get ACI Resource Failed - %v\nRESPONSE: %v", err, response)
		return response, err
	}

	return response, nil
}

func Create(
		parent_dn string,						// parent MO DN uni/...
		class_name string,						// ACI Class Name of object to be created i.e. fvTenant
		rel_name string,						// this is the mo objects relative name i.e. tn-TEN_TENANTNAME
		attributes []byte,						// this is a byte array of MO attributes in payload form
		overwriteexisting bool, 				// boolean to check if MO exists and to take managment if so
		d *schema.ResourceData, 				// data provided from TF run
		m interface{} ) (string, error) {		// APIC login credentials

	log.Printf("[DEBUG] Creating ACI Resource - %s", class_name)
	
	if !overwriteexisting {
		// check to ensure MO does not exist as we are creating
		generated_dn := fmt.Sprintf("%s/%s", parent_dn, rel_name) 
		
		exists, err := exists(generated_dn, m, class_name)
		if err != nil {
			return "", err
		}
		if exists {
			return "", errors.New( fmt.Sprintf("[DEBUG] Cannot create new MO: [%s] - already exists.", generated_dn ) )
		}
	}
	
	// POST to APIC
	log.Printf("[DEBUG] Create attributes: %s", attributes )
	response, err := post(m, parent_dn, class_name, attributes)
	if err != nil {
		return "", err
	}
	
	// POST reponse
	attrs, err := utils.ExtractMOAttributes(&response, class_name)
	if err != nil {
		return "", err
	}
	
	// GET new MO DN
	dn, ok := attrs["dn"].(string)
	if !ok {
		return "", errors.New( fmt.Sprintf("[DEBUG] Failed to extract DN from new MO [%s]", attrs ) )
	}

	return dn, nil
}

func Read(m interface{}, dn string) ([]byte, error) {
	log.Printf("[DEBUG] Reading ACI Resource DN - %s", dn)
	return get(m, dn)
}

func Update(
		parent_dn string,						// MO DN uni/...
		class_name string,						// ACI Class Name of object to be created i.e. fvTenant
		attributes []byte,						// this is a byte array of MO attributes in payload form
		d *schema.ResourceData, 				// data provided from TF run
		m interface{} ) (error) {		// APIC login credentials

	log.Printf("[DEBUG] Updating ACI Resource DN - %s", d.Id())

	exists, err := exists(d.Id(), m, class_name)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New( fmt.Sprintf("[DEBUG] Cannot update non-existent DN: [%s] ", d.Id() ) )
	}
	
	// POST to APIC
	log.Printf("[DEBUG] MO attributes: %s", attributes )
	_, err = post(m, parent_dn, class_name, attributes)
	if err != nil {
		return err
	}
	
	/*
	// POST reponse
	attrs, err := utils.ExtractMOAttributes(&response, class_name)
	if err != nil {
		return "", err
	}
	
	// Return DN of modified object
	dn, ok := attrs["dn"].(string)
	if !ok {
		return "", errors.New( fmt.Sprintf("[DEBUG] Failed to extract DN from updated MO [%s]", attrs ) )
	}
	*/
	return nil
}

func Delete(m interface{}, dn string) error {

	log.Printf("[DEBUG] Deleting ACI Resource DN - %s", dn)

	err := delete(m, dn)
	if err != nil {
		return errors.New(fmt.Sprintf("ACI Object Delete Failed DN: %s, Error %v", dn, err) )
	}

	return nil
}