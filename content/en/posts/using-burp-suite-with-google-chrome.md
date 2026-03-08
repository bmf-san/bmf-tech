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
  - Vulnerabilities
translation_key: using-burp-suite-with-google-chrome
---

# Overview
This post covers various settings for using Burp Suite with Chrome. I wanted to be able to use Burp in Chrome for vulnerability assessments and responses.

# Environment
Mac OS

# Preparation
- [burpsuite](https://support.portswigger.net/customer/portal/topics/718317-installing-and-configuring-burp/articles)
- [chrome extension - proxy switchsharp](https://chrome.google.com/webstore/detail/proxy-switchysharp/dpplabbmogkhghncfbfdeeokoefdjegm?hl=ja)

# Steps
## Set Up Proxy
Set up Proxy Profiles using the [chrome extension - proxy switchsharp](https://chrome.google.com/webstore/detail/proxy-switchysharp/dpplabbmogkhghncfbfdeeokoefdjegm?hl=ja).

Set the Profile Name to Burp (it can be anything, but since it's for Burp's proxy settings, let's keep it as Burp).
Choose Manual Configuration, set the HTTP Proxy to `127.0.0.1`, and configure the Port to avoid conflicts in your environment.
Press Save to save the Profile.

Using proxy switchsharp is convenient because it allows you to quickly change proxy settings without having to modify them each time.

You can change the proxy settings by selecting proxy switchsharp from the extensions in the upper right corner of Chrome and choosing the desired Profile.

When you don't need to adjust the proxy settings, select Direct Connection.

## Set Up Certificate
Launch Burp.

Check if the Chrome proxy settings are set to the Profile you configured above. (Just saving the proxy settings with proxy switchsharp does not activate them; you need to select the Profile from the proxy switchsharp extension in the upper right corner of the Chrome browser to enable it.)

If you launched Burp with the default settings, access `http://127.0.0.1:8080`.

Click on the CA Certificate in the upper right corner and download the certificate.

Open the downloaded certificate in Keychain Access and set it to `Always Trust`.
The certificate name should be `Port Swigger CA`.

With these steps, you should now be able to use Burp in Chrome.

# Additional Notes
To intercept localhost in Chrome, you need to add `<-loopback>` to No Proxy For in the Profile Detail of proxy switchsharp.

[Burp Interception does not work for localhost in Chrome](https://stackoverflow.com/questions/55616614/burp-interception-does-not-work-for-localhost-in-chrome)

# References
- [How to configure Google Chrome to use Burp Suite](https://taiyakon.com/2018/05/burp-suite-macosgoogle-chrome.html)