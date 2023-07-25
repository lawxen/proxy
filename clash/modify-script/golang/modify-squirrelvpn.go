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

	var data, proxyConfig, rulesConfigs map[string]interface{}

	// Load original YAML
	loadYAML("../../squirrelvpn.yaml", &data)

	// Load additional roxies and append them
	loadYAML("../config/proxies.yaml", &proxyConfig)
	proxies := data["proxies"].([]interface{})
	proxiesAdditional := proxyConfig["proxies"].([]interface{})
	data["proxies"] = append(proxies, proxiesAdditional...)

	// Load additional rules and prepend them
	loadYAML("../config/rules.yaml", &rulesConfigs)
	rules := data["rules"].([]interface{})
	rulesAdditional := rulesConfigs["rules"].([]interface{})
	data["rules"] = append(rulesAdditional, rules...)

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
