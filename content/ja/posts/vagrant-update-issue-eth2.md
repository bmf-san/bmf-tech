---
title: 'Vagrantを1.7.4から1.9.1に一気にアップデートしたら「Bringing up interface eth2:  Device eth2 does not seem to be present, delaying initialization.」'
description: 'Vagrantを1.7.4から1.9.1に一気にアップデートしたら「Bringing up interface eth2:  Device eth2 does not seem to be present, delaying initialization.」について、基本的な概念から実践的な知見まで詳しく解説します。'
slug: vagrant-update-issue-eth2
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - apache
  - Vagrant
  - VirtualBox
  - Tips
translation_key: vagrant-update-issue-eth2
---


Vagrantのバージョンが1.7.4というちょっと古いバージョンだったのでアップデートしてみたらネットワーク周りでエラーがでてハマった話です。


# エラー詳細
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

# 対応
色々調べたところ、ネットワーク周りの設定ファイルみたいなやつで引っかかっているらしいです。

解決に至る対応策は見当たらなかったので勘で対応しました（）


`cd /etc/sysconfig/network-scripts`

`mv ifcfg-eth2 eth2-ifcfg` 一旦テキトーな名前に変更しておく
`vagrant reload`


問題なければ先程のファイルを削除アンド`vagrant reload`
`rm -rf eth2-ifcfg` (edited)


1.7.4のときはeth0とeth1だけで、eth2は存在していなかった気がします。
eth2の中身を確認するとeth1とダブっていたので、「これいらんやろ」と消したら直ったというわけです。


# 所感
勘とはいえ参考サイトをヒントに考えた結果なのですが、この対応で問題ないのかちょっと不安ですｗ
 　

# 参考
* [[VirtualBox 4.3] 複製したゲストOS (CentOS) がネットワークに繋がらない: Device eth0 does not seem to be present, delaying initialization が表示された際の対応](http://qiita.com/satomyumi/items/964182390a08b678d576)
* [Vagrantのバージョンが沢山おいてある](https://releases.hashicorp.com/vagrant/)ー自己責任でお願いします。

