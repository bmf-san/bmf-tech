---
title: Took the Google Cloud Certified Associate Cloud Engineer Exam
slug: google-cloud-associate-engineer-exam
date: 2023-06-07T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Google Cloud Platform
description: Reflections on studying for the Google Cloud Certified Associate Cloud Engineer exam.
translation_key: google-cloud-associate-engineer-exam
---



# Overview
I took and passed the Google Cloud Certified Associate Cloud Engineer exam, so I wanted to reflect on my study process for future retakes or other exams.

# Background
I am a software engineer with about 7-8 years of experience. I have around 2 years of experience with GCP, though I have more experience with AWS.

# Motivation
I had opportunities to work with both AWS and GCP in my job. I had already obtained the Associate level certification for AWS but not for GCP, so I decided to study for it.

I obtained the AWS certification about two years ago. 
cf. [Took the AWS Certified Solutions Architect Associate Exam](https://bmf-tech.com/posts/AWS%e8%aa%8d%e5%ae%9a%e3%82%bd%e3%83%aa%e3%83%a5%e3%83%bc%e3%82%b7%e3%83%a7%e3%83%b3%e3%82%a2%e3%83%bc%e3%82%ad%e3%83%86%e3%82%af%e3%83%88%e3%82%a2%e3%82%bd%e3%82%b7%e3%82%a8%e3%82%a4%e3%83%88%e3%82%92%e5%8f%97%e9%a8%93%e3%81%97%e3%81%9f)

Originally, I planned to take the GCP exam after the AWS one, but various things happened, and it ended up being at this timing.

This year, I wanted to properly catch up with Kubernetes, and I felt that GKE might be better than EKS for Kubernetes operations (personal opinion). This motivation to learn about GKE and other GCP services and to be able to design architecture on GCP also pushed me to take the exam.

The AWS certification is valid for 3 years, while the GCP one is slightly shorter at 2 years, but this did not particularly affect my motivation.

# Exam Experience
I took the exam at a test center in Yokohama.

Like AWS, the exam can be taken online, but I chose onsite because there are fewer things to worry about regarding the environment.

It was a weekday, so there were hardly any people, and I could concentrate fully.

I can't say much about the exam content, but I felt that I could fully utilize what I studied.

The official exam results were delayed beyond the expected date, but it seems that others who took the exam around the same time experienced the same delay.
cf. https://note.com/aiue408/n/n8d5587f7362a

# Study Period
The study period was about 2-3 months.

I studied during gaps in my parental leave, but since I had been studying a bit beforehand, it felt like it was actually less than 2 months.

I planned to pass within 3 months, but I was able to move up the schedule by about 2 weeks, which was good.

# What I Studied
I summarized my study notes in [Notes on GCP](https://bmf-tech.com/posts/GCP%e3%81%ab%e3%81%a4%e3%81%84%e3%81%a6%e3%81%ae%e8%a6%9a%e3%81%88%e6%9b%b8%e3%81%8d). I wrote them roughly to input key points into my head, and they were quite helpful.

The official information is comprehensive, but there are fewer practice test collections (reference books) compared to AWS, which made me anxious about the fewer practice opportunities.

By understanding the exam scope and trends and studying based on official information, I managed to get by.

## Books
- GCP Textbook
- GCP Textbook 2

GCP Textbook 3 is also published, but since there didn't seem to be deep questions about Cloud AI products, I skipped reading it.

If you want to start with an introduction, it might be good to read, but if you're aiming for the shortest exam study, you might not need to read it.

## Documentation
- [Documentation](https://cloud.google.com/docs?hl=ja)
  - I skimmed through it.
  - Like the AWS certification, if you want to pass the exam quickly, you don't need to read everything. Just refer to it to fill in gaps identified by practice tests. I read it thoroughly because I was interested.
- [Cloud Architecture Guidance](https://cloud.google.com/architecture?hl=ja)
  - Various materials proposing references, guidance, and best practices
  - It's tough to read everything, so I only read what interested me
  - Regardless of the exam, if you're using GCP, reading the [Google Cloud Architecture Framework](https://cloud.google.com/architecture/framework?hl=ja) might be beneficial.
- [gcloud CLI Quick Reference](https://cloud.google.com/sdk/docs/cheatsheet?hl=ja)
  - I didn't look at everything. I only referred to commonly used commands related to practice tests and exam scope.
  - Based on the trends of Udemy practice tests, I thoroughly reviewed commands around iam, compute, container, config, and app.

## Blogs
- [GCPSketchnote](https://github.com/priyankavergadia/GCPSketchnote)
  - Articles summarizing key points and tips about services concisely
  - I read through them, but if you're aiming for the shortest exam study, you might not need to read them.
- [Google Cloud Japan Advent Calendar 2022 - Starting Google Cloud Now](https://zenn.dev/google_cloud_jp/articles/12bd83cd5b3370#%E4%BB%8A%E3%81%8B%E3%82%89%E5%A7%8B%E3%82%81%E3%82%8B-google-cloud)
  - You might not need to read it for the shortest exam study, but the Google Cloud Japan team writes easy-to-understand articles, so it's educational.
- [Associate Cloud Engineer Exam Preparation Manual: Exam Trends and Study Methods](https://blog.g-gen.co.jp/entry/associate-cloud-engineer)
  - I used it to check for any gaps in my learning after various studies
  - It covers key points of the exam
  - g-gen writes various easy-to-understand articles about GCP, so I read many articles besides this one
  - Since the official information is the primary source, make sure to check it before using it as a reference

# YouTube
- [【GCP】Google Cloud Platform Certified Associate Cloud Engineer Explanation Video【Google Cloud】](https://www.youtube.com/watch?app=desktop&v=7-IZv9o15t8)
  - I watched it at double speed to quickly grasp the key points

# Coursera
- [Architecting with Google Compute Engine Japanese Specialization](https://www.coursera.org/specializations/gcp-architecture-jp?action=enroll&authMode=signup#courses)
  - I didn't do this
  - It seemed to overlap with the Google Skills Boost content mentioned later, so I judged it unnecessary to take

# Google Skills Boost
- [Preparing for Your Associate Cloud Engineer Journey | Google Cloud Skills Boost](https://www.cloudskillsboost.google/course_templates/77?utm_source=gcp_training)
  - A free learning plan for obtaining GCP ACE is provided
  - If you're unsure where to start, following the [Workbook](https://www.cloudskillsboost.google/course_sessions/2870921/documents/375737) flow seems good
    - I should have followed this from the start
    - Related materials are provided for each theme, making it easy to learn

The following courses were also listed as related materials in the [Workbook](https://www.cloudskillsboost.google/course_sessions/2870921/documents/375737).

I completed several courses introduced in [Preparing for Your Associate Cloud Engineer Journey | Google Cloud Skills Boost](https://www.cloudskillsboost.google/course_templates/77?utm_source=gcp_training).

- [Essential Google Cloud Infrastructure: Foundation | Google Cloud Skills Boost](https://www.cloudskillsboost.google/course_templates/50)
- [Essential Google Cloud Infrastructure: Core Services | Google Cloud Skills Boost](https://www.cloudskillsboost.google/course_templates/49)
- [Elastic Google Cloud Infrastructure: Scaling and Automation | Google Cloud Skills Boost](https://www.cloudskillsboost.google/course_templates/178)
- [Getting Started with Google Kubernetes Engine | Google Cloud Skills Boost](https://www.cloudskillsboost.google/course_templates/2)
- [Google Cloud Fundamentals: Core Infrastructure | Google Cloud Skills Boost](https://www.cloudskillsboost.google/course_templates/60)

While learning from the above courses, I used Qwicklabs to cover areas I was weak in.

# Google Cloud Certified Jumpstart Program
- [2nd Google Cloud Certification Jumpstart Program](https://cloudonair.withgoogle.com/events/jp-cert-2)
  - An official program
  - I only glanced at it since I found out about it a few days before the exam, but it seemed efficient for learning

# Practice Tests
- [【GCP Certification】Google Cloud Platform Associate Cloud Engineer Practice Test](https://www.udemy.com/course/gcp-ace-mogi/)
  - There are four practice tests, and this is probably the most comprehensive as a test collection
  - I repeated answering and reviewing until I reached the target pass rate, though I only did two rounds...
  - It felt harder than the Google-provided practice test mentioned later, but studying based on this was meaningful
- [Associate Cloud Engineer Practice Test](https://docs.google.com/forms/d/e/1FAIpQLSc7bkUHpDbFShBI5xE4u8OO2vl99DrP0htnswa-la9DQynToA/viewform?hl=ja&hl=ja)
  - A free practice test provided by Google in a Google Form format
  - I tackled it as a final touch-up a couple of days before the exam

# Others
I tackled several GCP-related tasks on [Qwiklabs](https://nvlabs.qwiklabs.com/journeys).

Hands-on practice helps retain information and is practical. It was helpful during the AWS exam as well.

Since there are questions about gcloud commands in the exam, getting familiar with commands on Qwicklabs is meaningful.

It also helped deepen understanding of services I hadn't used before.

I found that a subscription was more cost-effective than purchasing individual credits, so I subscribed. I've subscribed about four times in my life, but I don't subscribe to an annual plan since I don't use it year-round, so I pay as needed...

# Thoughts
I'm glad I passed on the first try.

Unlike AWS certification exams, there is a retake policy for GCP, which imposes penalties like a maximum number of attempts and a cooldown period before retaking. You also have to pay the exam fee again for retakes, which adds pressure each time.

Even if I failed, I tried to think positively that understanding the exam trends and difficulty would increase my pass rate if I had to retake it, so I didn't focus on failing. However, my hands shook a bit when submitting the exam answers.

I believe AWS and GCP certification exams are not just memorization tests but exams that test practical knowledge. I think it's worth spending the cost (time and money) to acquire and demonstrate foundational knowledge, so I want to continue taking exams in the future.