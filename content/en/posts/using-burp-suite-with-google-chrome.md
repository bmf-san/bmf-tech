---
title: How to Use Burp Suite with Google Chrome
slug: using-burp-suite-with-google-chrome
date: 2019-03-22T00:00:00Z
author: bmf-san
categories:
  - Tools
tags:
  - Burp Suite
  - Security
  - Vulnerability
description: Steps to configure Burp Suite for use with Chrome.
translation_key: using-burp-suite-with-google-chrome
---

# Overview
This post covers various settings for using Burp Suite with Chrome. I wanted to use Burp with Chrome during vulnerability assessments and responses.

# Environment
Mac OS

# Preparation
- [burpsuite](https://support.portswigger.net/customer/portal/topics/718317-installing-and-configuring-burp/articles)
- [chrome extension - proxy switchsharp](https://chrome.google.com/webstore/detail/proxy-switchysharp/dpplabbmogkhghncfbfdeeokoefdjegm?hl=ja)

# Steps
## Configure the Proxy
Set up Proxy Profiles using [chrome extension - proxy switchsharp](https://chrome.google.com/webstore/detail/proxy-switchysharp/dpplabbmogkhghncfbfdeeokoefdjegm?hl=ja).

Name the profile Burp (or anything else, but since it's for Burp, name it Burp). Select Manual Configuration, set `127.0.0.1` for HTTP Proxy, and configure the Port to avoid conflicts in your environment. Save the profile.

Using proxy switchsharp allows you to quickly change proxy settings via the extension, avoiding the hassle of manually changing them each time.

To change proxy settings, select proxy switchsharp from the extensions in the top right of Chrome and choose any Profile.

When you don't need to fiddle with proxy settings, select Direct Connection.

## Configure the Certificate
Start Burp.

Ensure Chrome's proxy settings are set to the Profile configured above. (Simply saving the proxy settings with proxy switchsharp doesn't activate them, so you need to select the Profile from the extensions in the top right of the Chrome browser to activate it.)

If Burp is started with default settings, access `http://127.0.0.1:8080`.

Click the CA Certificate in the top right and download the certificate.

Open the downloaded certificate with Keychain Access and set it to `Always Trust`. The certificate name should be `Port Swigger CA`.

With these steps, Burp should now be usable with Chrome.

# Additional Notes
To intercept localhost with Chrome, add `<-loopback>` in proxy switchsharp's Profile Detail > No Proxy For.

[Burp Interception does not work for localhost in Chrome](https://stackoverflow.com/questions/55616614/burp-interception-does-not-work-for-localhost-in-chrome)

# References
- [How to configure Google Chrome to use Burp Suite](https://taiyakon.com/2018/05/burp-suite-macosgoogle-chrome.html)