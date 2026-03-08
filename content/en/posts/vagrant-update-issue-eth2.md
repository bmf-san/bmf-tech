---
title: 'Updated Vagrant from 1.7.4 to 1.9.1 and Encountered ''Bringing up interface eth2: Device eth2 does not seem to be present, delaying initialization.'''
slug: vagrant-update-issue-eth2
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - apache
  - Vagrant
  - VirtualBox
  - Tips
description: A story about encountering network errors after updating Vagrant from an older version.
translation_key: vagrant-update-issue-eth2
---


Vagrant was at the slightly old version of 1.7.4, so I decided to update it and got stuck with a network-related error.


# Error Details
```
"/etc/udev/rules.d/70-persistent-net.rules" is not a file
==> default: Configuring and enabling network interfaces...
The following SSH command responded with a non-zero exit status.
Vagrant assumes that this means the command failed!

# Down the interface before munging the config file. This might
# fail if the interface is not actually set up yet so ignore
# errors.
/sbin/ifdown 'eth1'
# Move new config into place
mv -f '/tmp/vagrant-network-entry-eth1-1485326655-0' '/etc/sysconfig/network-scripts/ifcfg-eth1'
# attempt to force network manager to reload configurations
nmcli c reload || true

# Restart network (through NetworkManager if running)
if service NetworkManager status 2>&1 | grep -q running; then
  service NetworkManager restart
else
  service network restart
fi


Stdout from the command:

Shutting down interface eth0:  [  OK  ]
Shutting down loopback interface:  [  OK  ]
Bringing up loopback interface:  [  OK  ]
Bringing up interface eth0:
Determining IP information for eth0... done.
[  OK  ]
Bringing up interface eth1:  Determining if ip address 192.168.33.10 is already in use for device eth1...
[  OK  ]
Bringing up interface eth2:  Device eth2 does not seem to be present, delaying initialization.
[FAILED]


Stderr from the command:

bash: line 10: nmcli: command not found
```

# Solution
After some research, it seems to be stuck on something like a network configuration file.

I couldn't find a solution, so I went with my gut.


`cd /etc/sysconfig/network-scripts`

`mv ifcfg-eth2 eth2-ifcfg` Temporarily rename it to something random
`vagrant reload`

If there are no issues, delete the previous file and `vagrant reload`
`rm -rf eth2-ifcfg` (edited)


In version 1.7.4, there were only eth0 and eth1, and eth2 didn't seem to exist.
When I checked the contents of eth2, it was overlapping with eth1, so I thought "we don't need this" and deleted it, which fixed the issue.


# Thoughts
Even though it was a gut feeling, I came up with this solution based on hints from reference sites, but I'm a bit worried if this is okay.


# References
* [[VirtualBox 4.3] Solution when a cloned guest OS (CentOS) cannot connect to the network: Device eth0 does not seem to be present, delaying initialization](http://qiita.com/satomyumi/items/964182390a08b678d576)
* [Various versions of Vagrant are available](https://releases.hashicorp.com/vagrant/) - Please proceed at your own risk.
