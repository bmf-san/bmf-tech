---
title: Managing Resumes with Github
slug: github-resume-management
date: 2023-06-07T00:00:00Z
author: bmf-san
categories:
  - Career
tags:
  - Git
  - GitHub
  - Resume
translation_key: github-resume-management
---

I have started managing my resume with Github instead of Google Drive.

I found this to be a personally beneficial initiative, so I wanted to write an article about it.

In this context, a resume refers to a text-based summary of various aspects of my career.

# Motivation
I was uploading and managing files on Google Drive, but it was somewhat difficult to update and I often lacked the motivation to do so. I thought it would be more desirable to have a method that makes differences clear and allows for easy publication.

There are two main reasons for managing my resume.

The first is "preparation for career uncertainties." I want to regularly take stock of my career, reflect on my experiences, and use it as material for self-analysis regarding what I want to pursue in my future career. I also want to have material ready to explain what I am capable of. This leads to the second reason.

The second reason is "preparation for job changes." Given the nature of being an engineer and the state of society and the economy, I believe it is necessary to always be prepared for a job change at any time.

# How am I managing it on Github?
I created a private repository on Github to manage my resume. (I want to make it a public repository, but I have just created it and have not yet sorted out what information can be made public and what should remain private.)

The base is referenced from [github.com - kawamataryo/resume](https://github.com/kawamataryo/resume).

## Directory Structure
It looks like this:

```sh
.
├── README.md
├── docs
│   ├── certification
│   ├── md
│   ├── pages
│   └── pdf
├── package-lock.json
└── package.json
```

The `certification` directory contains files related to certifications.

The `md` directory contains various resumes written in markdown format. I write my resumes in the following categories:

- Management Experience
- Technical Stack
- Extracurricular Activities
- Career Summary
- Work History

These categories and the format of each file are largely inspired by the resumes from Job Draft. I believe they sufficiently cover the necessary information for a resume.

For the work history, I refer to a CSV file prepared using my own [tool](https://github.com/bmf-san/go-github-pull-request) that downloads related PRs in CSV format from Github.

Using PRs makes the output clear and makes it easy to understand what I have done, which simplifies writing the resume. Of course, there may be outputs other than PRs, but I rely on my memory to fill in those gaps.

The various resumes written in markdown format are converted to PDF format (md-to-pdf), and files are generated under the `pdf` directory.

The `pages` directory contains resumes in pages format. This is something I want to change to markdown format due to its inertia.

## Operation
- Create a branch
- Update the resume
- Text structure `npm run textlint`
- Generate PDF `npm run md-to-pdf`
- Merge into main

This is roughly how it goes.

# Thoughts
When doing various tasks and personal activities daily, it can become cumbersome to reflect that in my resume, but I feel that the operation has improved, making it much easier.