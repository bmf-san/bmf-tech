---
title: Managing My Resume with GitHub
slug: github-resume-management
date: 2023-06-07T00:00:00Z
author: bmf-san
categories:
  - Career
tags:
  - Git
  - GitHub
  - Resume
description: I switched from Google Drive to GitHub for managing my resume.
translation_key: github-resume-management
---


I switched from managing my resume on Google Drive to using GitHub.

This has been a personally beneficial change, so I decided to write about it.

The resume here refers to a text-based summary of various aspects of my career.

# Motivation
I used to upload and manage files on Google Drive, but it was somewhat cumbersome to update, and I often lacked the motivation to do so. I thought it would be more desirable to have a method where differences are easily noticeable and can be shared publicly without much hassle.

There are two reasons why I manage my resume in the first place.

The first is "preparation for career anxiety." I want to periodically review my career, reflect on my experiences, and use it as material for self-analysis to determine what I aim for in my future career. Additionally, I want to have material ready to explain what I am capable of, which ties into the second reason.

The second reason is "preparation for job changes." Given the nature of being a specialist as an engineer and the state of society and the economy, I believe it's necessary to be prepared to change jobs at any time.

# How am I managing it on GitHub?
I created a private repository on GitHub to manage my resume. (I am considering making it a public repository, but I haven't yet sorted out what information can be public and what should remain private.)

The main reference is [github.com - kawamataryo/resume](https://github.com/kawamataryo/resume).

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

`certification` contains files related to certifications.

`md` is where I write various resumes in markdown format. I categorize my resumes as follows:

- Management History
- Tech Stack
- Activities Outside Work
- Career Summary
- Work History

I have greatly referenced the resume format from job draft sites for these categories and file formats, as I believe they sufficiently cover the necessary information for a resume.

For the work history, I use a [self-made tool](https://github.com/bmf-san/go-github-pull-request) to download CSV files of PRs I am involved in on GitHub, which I use as a reference when writing.

With PRs, the output is clear, making it easy to understand what I have done, thus making it easier to write the resume. Of course, there are outputs other than PRs, but I rely on memory for those.

The various resumes written in markdown format are converted to PDF (md-to-pdf), and files are generated under `pdf`.

`pages` stores resumes in pages format. This is a bit of a hassle, so I plan to change it to markdown format.

## Operation
- Create a branch
- Update the resume
- Text structure `npm run textlint`
- Generate PDF `npm run md-to-pdf`
- Merge into main

That's roughly how it goes.

# Thoughts
When you're doing various tasks and miscellaneous personal activities daily, it becomes cumbersome to reflect them in your resume, but with improved operations, it feels much easier now.
