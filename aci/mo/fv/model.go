// fv: fabric virtualization
package fv

// Application Profile
type Ap struct {
	Descr						string `json:"descr,omitempty" schema:"desc"`
	Dn							string `json:"dn,omitempty" schema:"dn"`
	Name						string `json:"name,omitempty" schema:"name"`
	OwnerKey					string `json:"ownerKey,omitempty" schema:"ownerkey"`
	OwnerTag					string `json:"ownerTag,omitempty" schema:"ownerkey"`	
	Prio						string `json:"prio,omitempty" schema:"priority_level"`	
}

// Bridge Domain
type BD struct {
	ArpFlood					string `json:"arpFlood,omitempty" schema:"arp_flooding" access:"admin"`
	BcastP						string `json:"bcastP,omitempty" schema:"multicast_group_ipv4" access:"implicit"`
	Descr						string `json:"descr,omitempty" schema:"desc" access:"admin"`
	Dn							string `json:"dn,omitempty" schema:"dn"`
	EpClear						string `json:"epClear,omitempty" schema:"endpoint_clear" access:"admin"`
	EpMoveDetectMode			string `json:"epMoveDetectMode,omitempty" schema:"garp_move_detection" access:"admin"`
	HostBasedRouting			string `json:"hostBasedRouting,omitempty" schema:"host_based_routing" access:"admin"`
	IntersiteBumTrafficAllow	string `json:"intersiteBumTrafficAllow,omitempty" schema:"intersite_bum_enabled" access:"admin"`
	IntersiteL2Stretch			string `json:"intersiteL2Stretch,omitempty" schema:"intersite_l2stretch" access:"admin"`
	IpLearning					string `json:"ipLearning,omitempty" schema:"ip_learning" access:"admin"`
	LimitIpLearnToSubnets		string `json:"limitIpLearnToSubnets,omitempty" schema:"limit_ip_learn_to_subnet" access:"admin"`
	LlAddr						string `json:"llAddr,omitempty" schema:"ipv6_local_link_addr" access:"admin"`
	Mac							string `json:"mac,omitempty" schema:"bridge_domain_mac" access:"admin"`
	McastAllow					string `json:"mcastAllow,omitempty" schema:"allow_multicast" access:"admin"`
	Mtu							string `json:"mtu,omitempty" schema:"max_l2_mtu" access:"implicit"`
	MultiDstPktAct				string `json:"multiDstPktAct,omitempty" schema:"l2_forwarding_method" access:"admin"`
	Name						string `json:"name,omitempty" schema:"name" access:"naming"`
	OptimizeWanBandwidth		string `json:"OptimizeWanBandwidth,omitempty" schema:"intersite_bw_optimize" access:"admin"`
	OwnerKey					string `json:"ownerKey,omitempty" schema:"ownerkey" access:"admin"`
	OwnerTag					string `json:"ownerTag,omitempty" schema:"ownertag" access:"admin"`
	Type						string `json:"type,omitempty" schema:"bridge_domain_type" access:"admin"`
	UnicastRoute				string `json:"unicastRoute,omitempty" schema:"unicast_routing_enabled" access:"admin"`
	UnkMacUcastAct				string `json:"unkMacUcastAct,omitempty" schema:"unknown_l2_forward_method" access:"admin"`
	UnkMcastAct					string `json:"unkMcastAct,omitempty" schema:"unknown_mcast_forward_method" access:"admin"`
	Vmac						string `json:"vmac,omitempty" schema:"l2_out_virtual_mac" access:"admin"`
}

// Context
type Ctx struct {
	Descr						string `json:"descr,omitempty" schema:"desc"`
	Dn							string `json:"dn,omitempty" schema:"dn" schema:"dn"`
	BdEnforcedEnable			string `json:"bdEnforcedEnable,omitempty" schema:"bd_enforced_mode"`
	KnwMcastAct					string `json:"knwMcastAct,omitempty" schema:"known_multicast_active"`
	Name						string `json:"name,omitempty" schema:"name"`
	OwnerKey					string `json:"ownerKey,omitempty" schema:"ownerkey"`
	OwnerTag					string `json:"ownerTag,omitempty" schema:"ownertag"`
	PcEnfDir					string `json:"pcEnfDir,omitempty" schema:"policy_control_direction"`
	PcEnfPref					string `json:"pcEnfPref,omitempty" schema:"policy_control_preference"`
}

// EPG
type AEPg struct {
	Descr						string `json:"descr,omitempty" schema:"desc"`
	Dn							string `json:"dn,omitempty" schema:"dn"`
	FloodOnEncap				string `json:"floodOnEncap,omitempty" schema:"flood_on_encap"`
	FwdCtrl						string `json:"fwdCtrl,omitempty" schema:"forward_control"`
	IsAttrBasedEPg				string `json:"isAttrBasedEPg,omitempty" schema:"is_attr_based_epg"`
	MatchT						string `json:"matchT,omitempty" schema:"label_match_criteria"`
	Name						string `json:"name,omitempty" schema:"name"`
	PcEnfPref					string `json:"pcEnfPref,omitempty" schema:"intra_epg_isolation"`
	PrefGrMemb					string `json:"prefGrMemb,omitempty" schema:"preferred_group_member"`
	Prio						string `json:"prio,omitempty" schema:"priority_level"`
}

// RS to Bridge Domain
type RsBd struct {
	Dn							string `json:"dn,omitempty" schema:"dn"`
	OwnerKey					string `json:"ownerKey,omitempty" schema:"ownerkey"`
	OwnerTag					string `json:"ownerTag,omitempty" schema:"ownertag"`
	TnFvBDName					string `json:"tnFvBDName,omitempty" schema:"bridge_domain_name"`
}

// RS to Consumed Contract
type RsCons struct {
	Dn							string `json:"dn,omitempty" schema:"dn"`
	Prio						string `json:"prio,omitempty" schema:"priority_level"`
	TnVzBrCPName				string `json:"tnVzBrCPName,omitempty" schema:"contract_name"`
}

// RS to Context
type RsCtx struct {
	Dn							string `json:"dn,omitempty" schema:"dn"`
	OwnerKey					string `json:"ownerKey,omitempty" schema:"ownerkey"`
	OwnerTag					string `json:"ownerTag,omitempty" schema:"ownertag"`
	TnFvCtxName					string `json:"tnFvCtxName,omitempty" schema:"ctx_name"`
}

// RS to Context
type RsDomAtt struct {
	ClassPref					string `json:"classPref,omitempty" schema:"class_pref" access:"admin"`
	Delimiter					string `json:"delimiter,omitempty" schema:"delimiter" access:"admin"`
	Dn							string `json:"dn,omitempty" schema:"dn" access:"admin"`
	Encap						string `json:"encap,omitempty" schema:"port_encapsulation" access:"admin"`
	EncapMode					string `json:"encapMode,omitempty" schema:"encapsulation_mode" access:"admin"`
	EpgCos						string `json:"epgCos,omitempty" schema:"epg_cos_value" access:"admin"`
	EpgCosPref					string `json:"epgCosPref,omitempty" schema:"epg_cos_enable" access:"admin"`
	InstrImedcy					string `json:"instrImedcy,omitempty" schema:"policy_deployment_mode" access:"admin"`
	NetflowDir					string `json:"netflowDir,omitempty" schema:"netflow_direction" access:"admin"`
	NetflowPref					string `json:"netflowPref,omitempty" schema:"netflow_enable" access:"admin"`
	PrimaryEncap				string `json:"primaryEncap,omitempty" schema:"primary_encap" access:"admin"`
	PrimaryEncapInner			string `json:"primaryEncapInner,omitempty" schema:"primary_encap_inner" access:"admin"`
	ResImedcy					string `json:"resImedcy,omitempty" schema:"policy_resolution_mode" access:"admin"`
	SecondaryEncapInner			string `json:"secondaryEncapInner,omitempty" schema:"second_encap_inner" access:"admin"`
	SwitchingMode				string `json:"switchingMode,omitempty" schema:"switching_mode" access:"admin"`
	TDn							string `json:"tDn,omitempty" schema:"domain_profile_id" access:"admin"`
}

// RS to Provider Contract
type RsProv struct {
	Dn							string `json:"dn,omitempty" schema:"dn" access:"admin"`
	Prio						string `json:"prio,omitempty" schema:"priority_level" access:"admin"`
	TnVzBrCPName				string `json:"tnVzBrCPName,omitempty" schema:"contract_name" access:"admin"`
}

// Tenant
type Tenant struct {
	Descr						string `json:"descr,omitempty"schema:"desc"`
	Dn							string `json:"dn,omitempty" schema:"dn"`
	Name						string `json:"name,omitempty"schema:"name"`
	OwnerKey					string `json:"ownerKey,omitempty"schema:"ownerkey"`
	OwnerTag					string `json:"ownerTag,omitempty"schema:"ownertag"`
}

// Subnet
type Subnet struct {
	Ctrl						string `json:"ctrl,omitempty" schema:"ctrl"`
	Descr						string `json:"descr,omitempty" schema:"desc"`
	Dn							string `json:"dn,omitempty" schema:"dn"`
	Ip							string `json:"ip,omitempty" schema:"ip"`
	Name						string `json:"name,omitempty" schema:"name"`
	OwnerKey					string `json:"ownerKey,omitempty" schema:"ownerkey"`
	OwnerTag					string `json:"ownerTag,omitempty" schema:"ownertag"`
	Preferred					string `json:"preferred,omitempty" schema:"preferred"`
	Scope						string `json:"scope,omitempty" schema:"scope"`
	Virtual						string `json:"virtual,omitempty" schema:"virtual"`
}
