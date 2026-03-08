---
title: Notes on NGINX Buffer-Related Directives
slug: nginx-buffer-directives-notes
date: 2024-03-02T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Nginx
translation_key: nginx-buffer-directives-notes
---

I wanted to record the buffer sizes, so here are my notes.

- [large_client_header_buffers](http://mogile.web.fc2.com/nginx/http/ngx_http_core_module.html#large_client_header_buffers)
    - Directive included in the ngx_http_core_module module
    - Syntax: large_client_header_buffers number size;
    - Default: large_client_header_buffers 4 8k;
    - Context: http, server
    - Specifies the number and size of buffers used to read request headers from clients.
    - Increasing the number increases the number of buffers and likely the memory used.
    - The maximum buffer size should depend solely on size, not in the form of number × size...
- [fastcgi_buffers](http://mogile.web.fc2.com/nginx/http/ngx_http_fastcgi_module.html#fastcgi_buffers)
    - Directive included in the ngx_http_fastcgi_module module
    - Syntax: fastcgi_buffers number size;
    - Default: fastcgi_buffers 8 4k|8k;
    - Context: http, server, location
    - Sets the number and size of buffers used to read responses from the FastCGI server for each connection.
    - size is not equivalent to fastcgi_buffer_size. It is thought to be the memory size used by the buffer...
- [fastcgi_buffer_size](http://mogile.web.fc2.com/nginx/http/ngx_http_fastcgi_module.html#fastcgi_buffer_size)
    - Directive included in the ngx_http_fastcgi_module module
    - Syntax: fastcgi_buffer_size size;
    - Default: fastcgi_buffer_size 4k|8k;
    - Context: http, server, location
    - Sets the size of the buffer used to read the first part of the response from the FastCGI server.