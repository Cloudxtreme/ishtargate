package ishtargate

import (
	"log"
)

type IshtarGate struct {
	// Servers
	// [ServerName]*ServerObject
	Servers map[string]*Server `json:"servers,omitempty"`
	// Global Firewall Rules
	// list of types can be found in `firewall.go`
	// 1 can be used for iptables
	FirewallType int `json:"firewall-type,omitempty"`
	// Before Server.FirewallRulesBefore
	FirewallRulesBefore []*Firewall_Rule `json:"firewall-rules-before,omitempty"`
	// After Server.FirewallRulesAfter
	FirewallRulesAfter []*Firewall_Rule `json:"firewall-rules-after,omitempty"`
	// Global Variables
	Vars map[string]interface{} `json:"vars,omitempty"`
}

func (self *IshtarGate) IsValid() bool {
	if self == nil {
		log.Println("IshtarGate nil")
		return false
	}
	if self.FirewallType != FIREWALL_IPTABLES {
		log.Printf("IshtarGate.FirewallType: %d INVALID\n", self.FirewallType)
		return false
	}
	if len(self.Servers) == 0 {
		log.Println("IshtarGate.Servers empty")
		return false
	}
	for name, server := range self.Servers {
		if name == "" {
			log.Println("IshtarGate.Servers[] name empty")
			return false
		}
		if !server.IsValid() {
			log.Printf("IshtarGate.Servers[%s] server invalid\n", name)
			return false
		}
	}
	// FirewallRulesBefore can be empty
	for _, rule := range self.FirewallRulesBefore {
		if !rule.IsValid() {
			log.Println("IshtarGate.FirewallRulesBefore rule invalid")
			return false
		}
	}
	// FirewallRulesAfter can be empty
	for _, rule := range self.FirewallRulesAfter {
		if !rule.IsValid() {
			log.Println("IshtarGate.FirewallRulesAfter rule invalid")
			return false
		}
	}
	return true
}
