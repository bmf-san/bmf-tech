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
translation_key: vagrant-update-issue-eth2
---

I was using an older version of Vagrant (1.7.4) and decided to update it, but I ran into network-related errors.

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

# Response
After researching, it seems that there was an issue with the network configuration files.

I couldn't find a definitive solution, so I improvised.

`cd /etc/sysconfig/network-scripts`

`mv ifcfg-eth2 eth2-ifcfg` Temporarily rename it to something arbitrary.
`vagrant reload`

If there are no issues, delete the previous file and `vagrant reload` again.
`rm -rf eth2-ifcfg`

I believe there were only eth0 and eth1 when I was on 1.7.4, and eth2 didn't exist. Upon checking the contents of eth2, it was a duplicate of eth1, so I thought, "This isn't needed," and deleted it, which resolved the issue.

# Thoughts
Even though it was an improvised solution based on hints from reference sites, I am a bit anxious about whether this approach is correct.

# References
* [[VirtualBox 4.3] When the cloned guest OS (CentOS) cannot connect to the network: Response when 'Device eth0 does not seem to be present, delaying initialization' is displayed](http://qiita.com/satomyumi/items/964182390a08b678d576)
* [Various versions of Vagrant available](https://releases.hashicorp.com/vagrant/) - Use at your own risk.