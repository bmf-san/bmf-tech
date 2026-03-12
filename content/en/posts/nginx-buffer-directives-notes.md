---
title: "NGINX Buffer Directives Explained: Optimizing Proxy Performance"
slug: nginx-buffer-directives
date: 2024-03-02T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Nginx
description: Notes taken while researching buffer sizes in NGINX.
translation_key: nginx-buffer-directives-notes
---

I took notes while researching buffer sizes.

- [large_client_header_buffers](http://mogile.web.fc2.com/nginx/http/ngx_http_core_module.html#large_client_header_buffers)
    - Directive included in the ngx_http_core_module
    - Syntax: large_client_header_buffers number size;
    - Default: large_client_header_buffers 4 8k;
    - Context: http, server
    - Specifies the number and size of buffers used to read request headers from the client
    - Increasing the number increases the buffer count and memory usage (probably)
    - The maximum buffer size depends solely on size, not in the form of number×size...
- [fastcgi_buffers](http://mogile.web.fc2.com/nginx/http/ngx_http_fastcgi_module.html#fastcgi_buffers)
    - Directive included in the ngx_http_fastcgi_module
    - Syntax: fastcgi_buffers number size;
    - Default: fastcgi_buffers 8 4k|8k;
    - Context: http, server, location
    - Sets the number and size of buffers used to read responses from the FastCGI server per connection
    - size ≒ fastcgi_buffer_size is not true. It is believed to be the memory size used by the buffer...
- [fastcgi_buffer_size](http://mogile.web.fc2.com/nginx/http/ngx_http_fastcgi_module.html#fastcgi_buffer_size)
    - Directive included in the ngx_http_fastcgi_module
    - Syntax: fastcgi_buffer_size size;
    - Default: fastcgi_buffer_size 4k|8k;
    - Context: http, server, location
    - Sets the size of the buffer used to read the first part of the response from the FastCGI server