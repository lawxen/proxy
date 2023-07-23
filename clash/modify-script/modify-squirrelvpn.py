# Modify the config file of SquirrelVPN to add shadowsocks proxy and bypass some site like openai.com and claude.ai\
# How to use:
# pip install pyyaml
# python modify-squirrelvpn.py
import yaml 

class MyDumper(yaml.Dumper):

    def increase_indent(self, flow=False, indentless=False):
        return super(MyDumper, self).increase_indent(flow, False)

# Load original squirrelvpn.yaml
with open('../squirrelvpn.yaml',encoding='utf8') as f:
    data = yaml.safe_load(f)

# Load proxy.yaml containing shadowsocks proxy config    
with open('proxy.yaml',encoding='utf8') as f:
    proxy = yaml.safe_load(f)

# Append imported proxy to proxies list    
data['proxies'].append(proxy[0])

# Load rules.yaml containing routing rules
with open('rules.yaml',encoding='utf8') as f:
    rules = yaml.safe_load(f)

# Insert rules to beginning of rules list
data['rules'][0:0] = rules
data["proxy-groups"][0]["proxies"].append("shadowsocks")

# Write out modified data to new yaml file 
with open('../squirrelvpn-modified.yaml', 'w',encoding='utf8') as f:
    yaml.dump(data, f, Dumper=MyDumper, default_flow_style=False, allow_unicode=True)