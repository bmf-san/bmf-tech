---
title: Created a Template to Manage Resumes and Work Histories with Markdown
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

Previously, I wrote an article about [managing resumes on Github](https://bmf-tech.com/posts/Github%E3%81%A7%E3%83%AC%E3%82%B8%E3%83%A5%E3%83%A1%E3%82%92%E7%AE%A1%E7%90%86%E3%81%99%E3%82%8B%E3%82%88%E3%81%86%E3%81%AB%E3%81%97%E3%81%9F).

I prepared a repository for resume management as a template, making it available for anyone to use.

[bmf-san/resume-manager](https://github.com/bmf-san/resume-manager)

With the recent advancements in AI, managing resumes has become significantly easier, so I highly recommend it to those who haven't been managing or regularly updating their resumes.

## Main Features

### 1. Centralized Management with Markdown

By managing all documents in Markdown, the following benefits are achieved:

- Version control with Git
- Efficient editing with text editors
- Easy visualization of differences
- Automatic proofreading with CI

### 2. Automation of PDF Generation

With the command `npm run pdf:all`, all documents can be converted to PDF. The following processes are performed automatically:

- Merging multiple Markdown files
- Automatic insertion of page breaks
- Safe replacement of personal information
- Standardization of PDF format

### 3. Safe Management of Personal Information

Personal information is centrally managed in `secrets.env`, which is excluded from Git management. By using placeholders in Markdown files:

- Reduces the risk of personal information leakage
- Ensures consistency of information across multiple documents
- Facilitates template sharing

### 4. Automation of Text Proofreading

Using textlint, the following proofreading tasks are automated:

- Detection and correction of inconsistencies in notation
- Checking the usage of particles
- Suggestions for improving readability

## Advanced Usage

This template can be used not only for basic management of resumes and work histories but also for managing various career-related documents such as:

- Career portfolios
- Management history
- Technology stack lists
- Records of extracurricular activities
- Career summaries
- Qualification information

Additionally, by combining it with recent AI tools (like ChatGPT and GitHub Copilot), the following tasks can be streamlined:

- Improving text structure
- Clarifying technical explanations
- Suggesting quantitative expressions of achievements
- Checking for consistency in writing

## Technology Stack

- Node.js
- md-to-pdf (PDF generation)
- textlint (text proofreading)
- dotenv (environment variable management)

## Conclusion

By using this template, I hope that managing career documents becomes more efficient, allowing for a greater focus on job hunting and career reflection.

The source code is [available on GitHub](https://github.com/bmf-san/resume-manager), so I encourage anyone interested to give it a try.