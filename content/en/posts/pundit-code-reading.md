---
title: Code Reading of Pundit
slug: pundit-code-reading
date: 2024-10-22T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - pundit
  - Ruby
description: Exploring the code of Pundit.
translation_key: pundit-code-reading
---

# Overview
Conducting a code reading of Pundit.

# Preparation
1. Clone the Pundit repository
   - `git clone git@github.com:varvet/pundit.git`

# Code Reading
Examining the `authorize` method used for applying permissions.

1. authorize
- [varvet/pundit/blob/main/lib/pundit.rb#L75](https://github.com/varvet/pundit/blob/main/lib/pundit.rb#L75)
  - Defined as a class method in the module
- [varvet/pundit/blob/main/lib/pundit/context.rb#L55](https://github.com/varvet/pundit/blob/main/lib/pundit/context.rb#L55)
  - The actual processing is handled by the `authorize` method in the `Pundit::Context` class
  - This is where the policy is checked and permissions are determined

The implementation was surprisingly simple.