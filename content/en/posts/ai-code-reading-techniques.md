---
title: Steps for Code Reading Using AI
slug: ai-code-reading-techniques
date: 2025-11-02T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Code Reading
  - AI
description: A miscellaneous note on the process of code reading utilizing AI.
translation_key: ai-code-reading-techniques
---

A miscellaneous note on the process of code reading utilizing AI.

Suggestions for prompt crafting would be helpful, but for now, only the process is summarized.

## Steps
### 1. Review the overall structure and key parts of the code
Use prompts to identify the overall structure of the code and the key parts that need to be read.

During this step, leverage README files, other documentation, and tools (such as call graphs) to ensure the understanding of the overall structure is reasonably accurate.

It’s also beneficial to check data structures and sequences.

### 2. Create a procedure document
Request AI to create a procedure document for code reading via prompts.

Tailor the document to the purpose of the code reading.

At a minimum, it should include the reading order and a checklist. If the amount of code can be read in a short time, a checklist may not be necessary. However, for time-consuming tasks, preparing a simple checklist for progress management can help offload working memory and make the process more convenient.

### 3. Read the code following the procedure document
Read the code according to the procedure document.

While reading the code, it’s helpful to take notes on the key points in a suitable format. Prepare a file for taking these notes. Summarize the information obtained from AI in these notes to save effort and focus on reading the code.

Additionally, if you find areas for improvement during the code reading, make sure to note them down as well.

## Summary
Use AI to efficiently perform the following:

- Grasp the overall structure
- Create various documents:
  - Code reading procedure document (reading order, checklist)
  - Code reading notes
  - Improvement notes

If deeper understanding is required, it’s also beneficial to create documents like [System Specification Structure](https://bmf-tech.com/posts/%e8%87%aa%e5%88%86%e7%9a%84%e3%82%b7%e3%82%b9%e3%83%86%e3%83%a0%e3%81%ae%e4%bb%95%e6%a7%98%e6%9b%b8%e3%81%ae%e6%a7%8b%e6%88%90).