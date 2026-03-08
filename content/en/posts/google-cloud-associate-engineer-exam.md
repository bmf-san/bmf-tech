---
title: Took the Google Cloud Certified Associate Cloud Engineer Exam
slug: google-cloud-associate-engineer-exam
date: 2023-06-07T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Google Cloud Platform
translation_key: google-cloud-associate-engineer-exam
---

# Overview
I passed the Google Cloud Certified Associate Cloud Engineer exam, so I want to reflect on my study process for future retakes or other exams.

# Background
I have about 7-8 years of experience as a software engineer. I have around 2 years of experience with GCP, but I have worked more with AWS.

# Motivation
I had the opportunity to work with both AWS and GCP in my job. I had already obtained the associate-level certification for AWS, but I hadn’t done so for GCP, so I decided to study for it.

I obtained the AWS certification about 2 years ago.
cf. [Took the AWS Certified Solutions Architect Associate Exam](https://bmf-tech.com/posts/AWS%e8%aa%8d%e5%ae%9a%e3%82%bd%e3%83%aa%e3%83%a5%e3%82%b8%e3%83%a7%e3%83%b3%e3%82%a2%e3%83%bc%e3%82%ad%e3%83%86%e3%82%af%e3%82%bf%e3%82%bd%e3%82%b7%e3%82%a8%e3%82%a4%e3%83%88%e3%82%92%e5%8f%97%e9%a8%93%e3%81%97%e3%81%9f)

Originally, I planned to take the GCP exam after the AWS certification, but due to various circumstances, it ended up happening at this timing.

This year, I wanted to properly catch up on Kubernetes, and I thought that for K8S operations, GKE might be better than EKS (just my personal opinion). This desire to learn about GCP services, including GKE, and to be able to design architectures on GCP also motivated me to take the exam.

The AWS certification is valid for 3 years, while GCP's is a bit shorter at 2 years, but this didn't particularly affect my motivation.

# Exam Experience
I took the exam at a test center in Yokohama.

Like AWS, it can also be taken online, but I chose the onsite option because there are fewer environmental concerns and distractions.

Since it was a weekday, there were hardly any people, allowing me to concentrate fully.

I can't say much about the exam content, but to sum it up in one word, I felt that I was able to demonstrate what I had studied sufficiently.

The official exam results were delayed more than expected, but it seems that others who took the exam around the same time also experienced delays.
cf. https://note.com/aiue408/n/n8d5587f7362a

# Study Period
The study period was about 2-3 months.

I studied during my maternity leave, but since I had already started studying a bit before, it felt like it was actually just under 2 months.

I had planned to pass within 3 months, but I was able to move my schedule up by about 2 weeks, which was great.

# What I Studied
I summarized my study notes in [Notes on GCP](https://bmf-tech.com/posts/GCP%e3%81%ab%e3%81%a4%e3%81%84%e3%81%a6%e3%81%ae%e8%a6%9a%e3%81%88%e6%9b%b8%e3%81%8d). I wrote it roughly to input the key points into my head, and it was quite helpful.

While the official information is comprehensive, there are fewer mock exam question collections (reference books) compared to AWS, which made me anxious about having fewer practice opportunities.

By understanding the exam scope and trends, and building my studies on the official information, I managed to get through.

## Books
- GCP Textbook
- GCP Textbook 2

There is also a GCP Textbook 3 published, but I skipped reading it because it didn't seem to have many deep questions about Cloud AI products.

If you want to get started, it might be good to read, but if you're aiming for the shortest path to exam preparation, you might not need to read it.

## Documentation
- [Documentation](https://cloud.google.com/docs?hl=ja) 
  - I went through it to some extent.
  - Similar to the AWS certification, if you want to pass the exam as quickly as possible, it's sufficient to not read everything but to solve mock exams and refer to the areas where you need improvement. I read it all because I was interested.
- [Cloud Architecture Guidance](https://cloud.google.com/architecture?hl=ja)
  - Various materials proposing references, guidance, and best practices
  - Since reading everything is tough, I only read what I was interested in.
  - If you're going to use GCP regardless of the exam, I think it wouldn't hurt to read the [Google Cloud Architecture Framework](https://cloud.google.com/architecture/framework?hl=ja), so this alone might be sufficient.
- [gcloud CLI Quick Reference](https://cloud.google.com/sdk/docs/cheatsheet?hl=ja)
  - I didn't look at everything. I only referred to commonly used commands related to the mock exams and exam scope.
  - Based on the trends from Udemy's mock exams, I carefully reviewed commands around iam, compute, container, config, and app.

## Blogs
- [GCPSketchnote](https://github.com/priyankavergadia/GCPSketchnote)
  - A collection of articles that concisely summarize key points and tips about services.
  - I read through it, but if you're preparing for the exam as quickly as possible, you might not need to read it.
- [Google Cloud Japan Advent Calendar 2022 - Starting Google Cloud Now](https://zenn.dev/google_cloud_jp/articles/12bd83cd5b3370#%E4%BB%8A%E3%81%8B%E3%82%89%E5%A7%8B%E3%82%81%E3%82%8B-google-cloud)
  - You might not need to read it if you're preparing for the exam quickly, but the people from Google Cloud Japan write understandable articles, so it's educational.
- [Associate Cloud Engineer Exam Preparation Manual: Exam Trends and Study Methods](https://blog.g-gen.co.jp/entry/associate-cloud-engineer)
  - I used this as a reference to check for any gaps in my learning after studying various materials.
  - It captures the key points of the exam.
  - G-gen has written various easy-to-understand articles about GCP, so I read several of them.
  - Since the official information is the primary source, it's essential to check that before using this as a reference.

# YouTube
- [【GCP】Google Cloud Platform Certified Associate Cloud Engineer Explanation Video【Google Cloud】](https://www.youtube.com/watch?app=desktop&v=7-IZv9o15t8)
  - I wanted to quickly grasp the key points, so I watched it at double speed.

# Coursera
- [Architecting with Google Compute Engine Japanese Version Specialized Course](https://www.coursera.org/specializations/gcp-architecture-jp?action=enroll&authMode=signup#courses)
  - I didn't take this.
  - I judged that it likely overlaps with the content of Google Skills Boost mentioned later, so I felt it wasn't necessary to take.

# Google Skills Boost
- [Preparing for Your Associate Cloud Engineer Journey | Google Cloud Skills Boost](https://www.cloudskillsboost.google/course_templates/77?utm_source=gcp_training)
  - A free learning plan is provided for obtaining GCP ACE.
  - If you're unsure where to start, it seems good to follow the [Workbook](https://www.cloudskillsboost.google/course_sessions/2870921/documents/375737) flow for studying.
    - I might have been better off following this from the start.
    - Related materials are provided by theme, making it easier to study.

The following are courses listed as related materials in the [Workbook](https://www.cloudskillsboost.google/course_sessions/2870921/documents/375737).

I also completed several courses introduced in [Preparing for Your Associate Cloud Engineer Journey | Google Cloud Skills Boost](https://www.cloudskillsboost.google/course_templates/77?utm_source=gcp_training).
- [Essential Google Cloud Infrastructure: Foundation | Google Cloud Skills Boost](https://www.cloudskillsboost.google/course_templates/50)
- [Essential Google Cloud Infrastructure: Core Services | Google Cloud Skills Boost](https://www.cloudskillsboost.google/course_templates/49)
- [Elastic Google Cloud Infrastructure: Scaling and Automation | Google Cloud Skills Boost](https://www.cloudskillsboost.google/course_templates/178)
- [Getting Started with Google Kubernetes Engine | Google Cloud Skills Boost](https://www.cloudskillsboost.google/course_templates/2)
- [Google Cloud Fundamentals: Core Infrastructure | Google Cloud Skills Boost](https://www.cloudskillsboost.google/course_templates/60)

While learning from the above courses, I also covered my weak areas with Qwiklabs.

# Google Cloud Certified Jumpstart Program
- [2nd Google Cloud Certification Jumpstart Program](https://cloudonair.withgoogle.com/events/jp-cert-2)
  - An official program provided by Google.
  - I learned about it a few days before the exam, so I only glanced at it, but it seems like a good way to study efficiently.

# Question Collections
- [【GCP Certification】Google Cloud Platform Associate Cloud Engineer Mock Question Collection](https://www.udemy.com/course/gcp-ace-mogi/)
  - There are 4 practice questions, and this is probably the most comprehensive question collection.
  - I repeated answering and reviewing until I reached the target pass rate. Although I only did it twice...
  - It seems a bit more difficult than the mock exams provided by Google, but I think studying based on this was meaningful.
- [Associate Cloud Engineer Mock Exam](https://docs.google.com/forms/d/e/1FAIpQLSc7bkUHpDbFShBI5xE4u8OO2vl99DrP0htnswa-la9DQynToA/viewform?hl=ja&hl=ja)
  - A mock exam in the form of a Google Form provided for free by Google.
  - I worked on this as a final touch about two days before the exam.

# Others
I worked on several GCP-related tasks on [Qwiklabs](https://nvlabs.qwiklabs.com/journeys).

Hands-on practice helps retention and is practical, which is great. I also relied on it during the AWS exam.

Since there are questions in the exam that ask about gcloud commands, I think it's beneficial to get familiar with commands through Qwiklabs.

It also helped deepen my understanding of services I hadn't interacted with before.

I found that a subscription was more cost-effective than purchasing individual credits, so I opted for that. I've charged for it about four times in my life, but I haven't subscribed to an annual plan, so I charge each time... (since I don't use it year-round...)

# Thoughts
I'm glad I passed on the first try.

Unlike the AWS certification exam, there is a re-examination policy in place, so if you fail, you face penalties like maximum attempts and a cooldown period before retaking. This adds pressure each time.

Even if I had failed, I thought positively that I could understand the trends and difficulty of the questions, so if I had to retake it, my chances of passing would increase. I tried not to focus on failing, but my hands did shake a bit when submitting my answers, haha.

I believe that AWS and GCP certification exams are not just memorization tests but assess practical knowledge, so I think it's worth the cost (time and money) to acquire and demonstrate foundational knowledge. I want to continue taking exams in the future.