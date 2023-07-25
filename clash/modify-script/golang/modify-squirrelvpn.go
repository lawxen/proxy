package main

import (
  "gopkg.in/yaml.v3"
  "log"
  "os"
)

func loadYAML(path string, v interface{}) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(data, v)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	var data map[string]interface{}

	// Load original YAML
	loadYAML("../../squirrelvpn.yaml", &data)

	// Load proxy YAML
	var proxyConfigs []interface{}
	loadYAML("../config/proxy.yaml", &proxyConfigs)

	// Append proxies
	proxies := data["proxies"].([]interface{})
	data["proxies"] = append(proxies, proxyConfigs...)

	// Load rules YAML
	var rulesConfigs []interface{}
	loadYAML("../config/rules.yaml", &rulesConfigs)

	// Prepend rules
	rules := data["rules"].([]interface{})
	data["rules"] = append(rulesConfigs, rules...)

	// Add proxy to groups
	proxyGroups := data["proxy-groups"].([]interface{})
	proxyGroup := proxyGroups[0].(map[string]interface{})

	groupProxies := proxyGroup["proxies"].([]interface{})
	proxyGroup["proxies"] = append(groupProxies, "shadowsocks")

	// Write output YAML
	outFile, _ := os.Create("../../squirrelvpn-modified.yaml")
	defer outFile.Close()

	yaml.NewEncoder(outFile).Encode(data)
}
