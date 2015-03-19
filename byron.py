# example usage: python byron.py -t e805bd8086eebade820ac8368ec2fbf4

import json
import os
import sys
import urllib2

from optparse import OptionParser
from subprocess import call

parser = OptionParser()
parser.add_option("-t", "--token", dest="token")
(options, args) = parser.parse_args()

try:
  resp  = urllib2.urlopen("http://169.254.169.254/latest/meta-data/instance-id")
except Exception as e:
  print "Unable to retrieve provider ID: " + e
  sys.exit(1)

id = resp.read()
print "OBTAINED PROVIDER ID: " + id

try:
  resp = urllib2.urlopen("https://api.serverdensity.io/inventory/devices?token=" + options.token)
except Exception as e:
  print "Unable to retrieve JSON payload"
  sys.exit(1)

try:
  json = json.load(resp)
except Exception as e:
  print "Unable to parse JSON payload"
  sys.exit(1)

for i in json:
  print "FOUND A PROVIDER ID: "  + i["providerId"]
  if i["providerId"] == id:
    agent_key = i["agentKey"]
    break

print "HERE IS YOUR AGENT KEY: " + agent_key

try:
  call(["wget", "https://www.serverdensity.com/downloads/agent-install.sh"])
except Exception as e:
  print "Unable to retrive installation script"
  sys.exit(1)

os.chmod("agent-install.sh", 0755)

call(["sed", "-i", "s/www.example.com/localhost/", "agent-install.sh"])

try:
  call(["./agent-install.sh", "-a", "https://economist.serverdensity.io", "-g", "Drupal-Web", "-k", agent_key])
except Exception as e:
  print "Unable to run installation script"
  sys.exit(1)