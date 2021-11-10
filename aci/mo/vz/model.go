// vz: virtual zones (former name of the policy controls) i.e. Contracts
package vz

// Contract
type BrCP struct {
	Descr						string `json:"descr,omitempty" schema:"desc"`
	Dn							string `json:"dn,omitempty" schema:"dn"`
	Name						string `json:"name,omitempty" schema:"name"`
	NameAlias					string `json:"nameAlias,omitempty" schema:"name_alias"`
	OwnerKey					string `json:"ownerKey,omitempty" schema:"ownerkey"`
	OwnerTag					string `json:"ownerTag,omitempty" schema:"ownertag"`	
	Prio						string `json:"prio,omitempty" schema:"priority_level"`	
	Scope						string `json:"scope,omitempty" schema:"contract_scope"`
	TargetDscp					string `json:"targetDscp,omitempty" schema:"target_dscp"`
}

// Contract Subject
type Subj struct {
	ConsMatchT					string `json:"consMatchT,omitempty" schema:"consumer_subject_match"`
	Descr						string `json:"descr,omitempty" schema:"desc"`
	Dn							string `json:"dn,omitempty" schema:"dn"`
	Name						string `json:"name,omitempty" schema:"name"`
	NameAlias					string `json:"nameAlias,omitempty" schema:"name_alias"`
	Prio						string `json:"prio,omitempty" schema:"priority_level"`
	ProvMatchT					string `json:"provMatchT,omitempty" schema:"provider_subject_match"`	
	RevFltPorts					string `json:"revFltPorts,omitempty" schema:"reverse_filter_ports"`	
	TargetDscp					string `json:"targetDscp,omitempty" schema:"target_dscp"`
}

// Contract Subject Reln to vzFilters
type RsSubjFiltAtt struct {
	Directives					string `json:"directives,omitempty" schema:"directives"`
	Dn							string `json:"dn,omitempty" schema:"dn"`
	TnVzFilterName				string `json:"tnVzFilterName,omitempty" schema:"filter_name"`	
}

// Filter Group
type Filter struct {
	Descr						string `json:"descr,omitempty" schema:"desc"`
	Dn							string `json:"dn,omitempty" schema:"dn"`
	Name						string `json:"name,omitempty" schema:"name"`	
	NameAlias					string `json:"nameAlias,omitempty" schema:"name_alias"`	
	OwnerKey					string `json:"ownerKey,omitempty" schema:"ownerkey"`	
	OwnerTag					string `json:"ownerTag,omitempty" schema:"ownertag"`	
}

// Filter Entry - Child of vzFilter.
type Entry struct {
	ApplyToFrag					string `json:"applyToFrag,omitempty" schema:"apply_to_frag"`
	ArpOpc						string `json:"arpOpc,omitempty" schema:"arp_operation"`
	DFromPort					string `json:"dFromPort,omitempty" schema:"dest_from_port"`
	DToPort						string `json:"dToPort,omitempty" schema:"dest_to_port"`
	Descr						string `json:"descr,omitempty" schema:"desc"`
	Dn							string `json:"dn,omitempty" schema:"dn"`
	EtherT						string `json:"etherT,omitempty" schema:"ethernet_type"`
	Icmpv4T						string `json:"icmpv4T,omitempty" schema:"icmpv4_type"`
	Icmpv6T						string `json:"icmpv6T,omitempty" schema:"icmpv6_type"`
	MatchDscp					string `json:"matchDscp,omitempty" schema:"match_dscp"`
	Name						string `json:"name,omitempty" schema:"name"`
	NameAlias					string `json:"nameAlias,omitempty" schema:"name_alias"`
	Prot						string `json:"prot,omitempty" schema:"l3_ip_protocol"`
	SFromPort					string `json:"sFromPort,omitempty" schema:"source_from_port"`
	SToPort						string `json:"sToPort,omitempty" schema:"source_to_port"`
	Stateful					string `json:"stateful,omitempty" schema:"stateful"`
	TcpRules					string `json:"tcpRules,omitempty" schema:"tcp_session_flags"`
}

// Consumer Contract Label
type ConsLbl struct {
	Descr						string `json:"descr,omitempty" schema:"desc"`
	Dn							string `json:"dn,omitempty" schema:"dn"`
	Name						string `json:"name,omitempty" schema:"name"`
	NameAlias					string `json:"nameAlias,omitempty" schema:"name_alias"`	
	OwnerKey					string `json:"ownerKey,omitempty" schema:"ownerkey"`	
	OwnerTag					string `json:"ownerTag,omitempty" schema:"ownertag"`	
	Tag							string `json:"tag,omitempty" schema:"tag"`
}

// Consumer Contract Label
type ProvLbl struct {
	Descr						string `json:"descr,omitempty" schema:"desc"`
	Dn							string `json:"dn,omitempty" schema:"dn"`
	IsComplement				string `json:"isComplement,omitempty" schema:"is_complement"`
	Name						string `json:"name,omitempty" schema:"name"`
	NameAlias					string `json:"nameAlias,omitempty" schema:"name_alias"`	
	OwnerKey					string `json:"ownerKey,omitempty" schema:"ownerkey"`	
	OwnerTag					string `json:"ownerTag,omitempty" schema:"ownertag"`	
	Tag							string `json:"tag,omitempty" schema:"tag"`
}

// Consumer Subject Label
type ConsSubjLbl struct {
	Descr						string `json:"descr,omitempty" schema:"desc"`
	Dn							string `json:"dn,omitempty" schema:"dn"`
	IsComplement				string `json:"isComplement,omitempty" schema:"is_complement"`
	Name						string `json:"name,omitempty" schema:"name"`
	NameAlias					string `json:"nameAlias,omitempty" schema:"name_alias"`	
	OwnerKey					string `json:"ownerKey,omitempty" schema:"ownerkey"`	
	OwnerTag					string `json:"ownerTag,omitempty" schema:"ownertag"`	
	Tag							string `json:"tag,omitempty" schema:"tag"`
}

// Provider Subject Label
type ProvSubjLbl struct {
	Descr						string `json:"descr,omitempty" schema:"desc"`
	Dn							string `json:"dn,omitempty" schema:"dn"`
	IsComplement				string `json:"isComplement,omitempty" schema:"is_complement"`
	Name						string `json:"name,omitempty" schema:"name"`
	NameAlias					string `json:"nameAlias,omitempty" schema:"name_alias"`	
	OwnerKey					string `json:"ownerKey,omitempty" schema:"ownerkey"`	
	OwnerTag					string `json:"ownerTag,omitempty" schema:"ownertag"`	
	Tag							string `json:"tag,omitempty" schema:"tag"`
}
