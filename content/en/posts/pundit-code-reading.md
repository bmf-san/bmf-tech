---
title: Code Reading of Pundit
description: 'Understand Ruby Pundit authorization framework through code reading, exploring permission policies and authorization mechanisms.'
slug: pundit-code-reading
date: 2024-10-22T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - pundit
  - Ruby
translation_key: pundit-code-reading
---

# Overview
Reading the code of Pundit.

# Preparation
1. Clone the Pundit repository
   - `git clone git@github.com:varvet/pundit.git`

# Code Reading
Let's take a look at the `authorize` method used when applying permissions.

1. authorize
- [varvet/pundit/blob/main/lib/pundit.rb#L75](https://github.com/varvet/pundit/blob/main/lib/pundit.rb#L75)
  - Defined as a class method of the module
- [varvet/pundit/blob/main/lib/pundit/context.rb#L55](https://github.com/varvet/pundit/blob/main/lib/pundit/context.rb#L55)
  - The actual processing is done in the `authorize` method of the `Pundit::Context` class
  - Here, the policy is checked, and permissions are determined.

The implementation was surprisingly simple.