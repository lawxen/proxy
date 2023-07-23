package main

import (
  "io/ioutil"
  "log"

  "gopkg.in/yaml.v3"
)

func main() {

  // Load original YAML
  dataBytes, err := ioutil.ReadFile("../squirrelvpn.yaml")
  if err != nil {
    log.Fatal(err)
  }

  var data map[string]interface{}
  err = yaml.Unmarshal(dataBytes, &data)
  if err != nil {
    log.Fatal(err)
  }

  // Load proxy YAML
  proxyBytes, err := ioutil.ReadFile("proxy.yaml")
  if err != nil {
    log.Fatal(err)
  }

  var proxyConfigs []interface{}
  err = yaml.Unmarshal(proxyBytes, &proxyConfigs)
  if err != nil {
    log.Fatal(err)
  }

  // Append proxies
  proxies, ok := data["proxies"].([]interface{})
  if !ok {
    log.Fatal("proxies not found")
  }

  data["proxies"] = append(proxies, proxyConfigs...)

  // Load rules YAML
  rulesBytes, err := ioutil.ReadFile("rules.yaml")
  if err != nil {
    log.Fatal(err)
  }

  var rulesConfigs []interface{}
  err = yaml.Unmarshal(rulesBytes, &rulesConfigs)
  if err != nil {
    log.Fatal(err)
  }

  // Prepend rules
  rules, ok := data["rules"].([]interface{})
  if !ok {
    log.Fatal("rules not found")
  }

  data["rules"] = append(rulesConfigs, rules...)

  // Add proxy to groups
  proxyGroups, ok := data["proxy-groups"].([]interface{})
  if !ok {
    log.Fatal("proxy-groups not found")
  }

  proxyGroup := proxyGroups[0].(map[string]interface{})
  groupProxies := proxyGroup["proxies"].([]interface{})
  proxyGroup["proxies"] = append(groupProxies, "shadowsocks")

  // Marshal and write output
  output, err := yaml.Marshal(data)
  if err != nil {
    log.Fatal(err)
  }

  err = ioutil.WriteFile("../squirrelvpn-modified.yaml", output, 0644)
  if err != nil {
    log.Fatal(err)
  }
}