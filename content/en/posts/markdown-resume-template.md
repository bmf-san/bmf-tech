---
title: Created a Template to Manage Resumes and CVs with Markdown
description: A step-by-step guide on Created a Template to Manage Resumes and CVs with Markdown, with practical examples and configuration tips.
slug: markdown-resume-template
date: 2025-07-16T00:00:00Z
author: bmf-san
categories:
  - Career
tags:
  - Resume
  - Job Change
translation_key: markdown-resume-template
---

## Introduction

Previously, I wrote an article about [managing resumes on GitHub](https://bmf-tech.com/posts/Github%E3%81%A7%E3%83%AC%E3%82%B8%E3%83%A5%E3%83%A1%E3%82%92%E7%AE%A1%E7%90%86%E3%81%99%E3%82%8B%E3%82%88%E3%81%86%E3%81%AB%E3%81%97%E3%81%9F).

I have prepared a repository for resume management as a template so that anyone can use it.

[bmf-san/resume-manager](https://github.com/bmf-san/resume-manager)

With the power of AI making resume management significantly easier these days, I highly recommend it to those who do not regularly update their resumes.

## Main Features

### 1. Centralized Management with Markdown

By managing all documents with Markdown, the following benefits are realized:

- Version control with Git
- Efficient editing with text editors
- Easy visualization of differences
- Automatic proofreading with CI

### 2. Automated PDF Generation

With a single `npm run pdf:all` command, all documents can be converted to PDF. The following processes are automated:

- Merging multiple Markdown files
- Automatic insertion of page breaks
- Secure replacement of personal information
- Unified PDF formatting

### 3. Secure Management of Personal Information

Personal information is centrally managed in `secrets.env` and is excluded from Git management. By using placeholders in Markdown files:

- Reduce the risk of personal information leakage
- Ensure consistency of information across multiple documents
- Easy sharing of templates

### 4. Automated Text Proofreading

Using textlint, the following proofreading tasks are automated:

- Detection and correction of inconsistencies
- Checking the use of particles
- Suggestions for improving readability

## Advanced Usage

In addition to basic resume and CV management, this template can also be used for managing the following career-related documents:

- Career portfolio
- Management history
- Technical stack list
- Records of extracurricular activities
- Career summary
- Certification information

Furthermore, by combining with recent AI tools (ChatGPT and GitHub Copilot), the following tasks can be streamlined:

- Improving document structure
- Clarifying technical explanations
- Proposing quantitative expressions of achievements
- Checking document consistency

## Tech Stack

- Node.js
- md-to-pdf (PDF generation)
- textlint (text proofreading)
- dotenv (environment variable management)

## Conclusion

By using this template, I hope that managing career documents becomes more efficient, allowing you to focus on job hunting and career reflection.

The source code is [available on GitHub](https://github.com/bmf-san/resume-manager), so if you're interested, please give it a try.