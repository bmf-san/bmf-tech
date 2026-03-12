---
title: Notes on Building a Notification Platform
description: Research notes and a structured overview of Notes on Building a Notification Platform, summarizing key concepts and findings.
slug: notification-infrastructure-notes
date: 2023-08-28T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Notification
translation_key: notification-infrastructure-notes
---



# Overview
This post summarizes my thoughts and research on building a notification platform.

# What is a Notification Platform?
A system platform for sending notifications (email, push, SMS, voice, etc.) to users.

It receives requests from clients (systems requesting notifications) and handles processes related to notifications, such as the destination and content.

# Considerations in Designing and Implementing a Notification Platform
There seems to be a lot to consider, so I wrote down my thoughts in no particular order. It's not organized.

- Nature of Notification Messages
  - Nature
    - Important Announcements
      - ex. Terms changes, payment-related
    - Marketing
      - ex. Sale information
  - Depending on the nature, the architecture might prioritize different characteristics
    - For example, for emails, announcements might use SES, while marketing might use Amazon Pinpoint, optimizing performance and cost based on characteristics
- Notification Channels
  - Email, push, SMS, voice, etc.
- Sending Patterns
  - Individual sending
  - Group sending
  - Cancellation
    - Mechanism to cancel sent notifications
    - Also serves as a fail-safe
      - Prevents human errors like setting 1000 instead of 1, which can directly impact costs
  - Notification Priority
  - Conditional Specification
    - Mechanism to conditionally branch notifications, like not sending SMS if email is possible
    - Time restrictions for sending
- Cost
  - Mechanisms and ideas to reduce costs
- Scalability
  - Pay attention to internal systems that integrate with the notification platform
    - For example, if user-related information is needed, consider the performance of that system
      - Fewer dependencies on external systems (services outside the notification platform) are better
- Messaging
  - Queuing
  - Scheduling
- Management of Recipient User Information
  - Recipient information
  - Opt-in/Opt-out
- Authentication
  - Necessary when accessing personal data or integrating with authentication-related APIs?
- Notification Message Templates
  - Templates for embedding static/dynamic data
  - Whether multilingual support is needed
- Error Handling and Retry
  - How to catch errors like sending failures and retry
  - Bounce processing for emails
- Monitoring
  - Not just system metrics, but also notification-related data (ex. reception rate, click rate)
  - Ability to track notifications
    - Need to issue something like a trace ID?
- Extensibility
  - Adding notification channels, customizing templates
- Integration with External Services
  - Design to be replaceable and not overly dependent
  - Integration with analytics platforms
    - Optimizing notifications can reduce manual operation
    - Desire to automate where possible
- Operation
  - Obviously consider operation (from both developer and business perspectives)
    - Current operations and future operations
- Testing
  - Ease of testing and debugging
    - High difficulty due to potential integration of multiple systems

# Solutions
## SaaS
There are various options like Salesforce, Braze, Airship, OneSignal, SendGrid.

These services not only send notifications across multiple channels but also integrate with customer management and marketing tools.

There are many services supporting multichannel notifications.

## PaaS
AWS Pinpoint is a platform supporting multichannel notifications.

# Case Studies
I explored domestic and international case studies.

- [techblog.zozo.com - Introduction and Replacement Plan of Real-time Marketing System](https://techblog.zozo.com/entry/real-time-marketing-system)
  - A case of building a platform supporting multichannel like AWS Pinpoint, including MA elements
- [techblog.zozo.com - Story of Replacing Real-time Data Integration Platform Supporting Marketing Automation System](https://techblog.zozo.com/entry/ma-realtime-data-infrastructure-replacement)
  - Continuation of the above, discussing the replacement
- [techblog.zozo.com - Rule-based Optimization Improvement in Personalized Distribution](https://techblog.zozo.com/entry/improving-optimization-for-personalized-marketing)
  - Impressive. This is the essence of notifications.
  - To reach this level, the business impact of notifications must be well estimated, and notifications can significantly impact business
  - "There is still much room for improvement to achieve the true goal of delivering only the notifications users truly want."
- [leandrofranchi.medium.com - How to design a Notification System](https://leandrofranchi.medium.com/how-to-design-a-notification-system-23f381cdeb00)
  - Example of system architecture design for a multichannel notification platform
  - Does not particularly consider access to user-related data, focusing solely on notification delivery
  - The configuration would likely look like this
  - It's not always necessary to prepare a common interface for notifications from the start; finding an interface as notification channels are developed is also a valid approach
- [www.notificationapi.com - Notification Service Design - with diagrams](https://www.notificationapi.com/blog/notification-service-design-with-architectural-diagrams)
  - Similar to the above, an example design
  - The overall structure is similar. This one considers user settings, while the above considers logging
- [cloudificationzone.com - Notification System Design](https://cloudificationzone.com/2021/08/13/notification-system-design/)
  - A more concrete design example
  - Unclear what exactly "Inbound" notifications are
- [atmarkit.itmedia.co.jp - Basics of Push Notifications & Architecture and Mechanism of Push Notification Platform Handling Over 10,000 per Second](https://atmarkit.itmedia.co.jp/ait/articles/1412/18/news022.html)
  - Push notification platform utilizing DynamoDB and Node.js
- [zenn.dev - AWS Architecture for Mass Email/Push Notification Handling 100,000 Requests per Minute](https://zenn.dev/coconala/articles/a3a5e33cd1d984)
  - Story of switching to a more cost-effective and performant architecture
  - Improved scalability by distributing the most time-consuming delivery API requests
    - Scalability increases linearly with the number of delivery workers
- [www.slideshare.net - Architecture of Push Notification Platform for System Speedup Forum](https://www.slideshare.net/recruitcojp/ss-42921628)

# AWS Pinpoint
I personally researched AWS Pinpoint.

- An AWS service for messages (notifications) across multiple channels, released in 2016.
  - Supports push notifications, email, SMS, and voice messages
  - Additionally, the [Custom Channels](https://docs.aws.amazon.com/ja_jp/pinpoint/latest/userguide/channels-custom.html) feature allows expanding notification channels
    - For example, adding Facebook Messenger
- Pay-as-you-go
  - Seems not as expensive as expected
  - Email costs $1.00 per 10,000, push notifications are free for the first million, then $1.00 per million
  - With a rough estimate for millions of users, it could become a significant amount
    - Obviously, need to rigorously verify cost-effectiveness
- Can analyze notifications
  - Could integrate with marketing initiatives
    - Limited capabilities, so compatibility needs consideration
- Scalability
  - Limit on the number of notifications sent per second
    - Can request quota increase
  - AWS handles scaling for the most time-consuming part of sending notifications

## References
- [docs.aws.amazon.com - What is Amazon Pinpoint](https://docs.aws.amazon.com/ja_jp/pinpoint/latest/developerguide/welcome.html)
- [pages.awscloud.com - Engage Users with Amazon Pinpoint
〜Amazon Digital User Engagement〜](https://pages.awscloud.com/rs/112-TZM-766/images/A3-01.pdf)
- [www.slideshare.net - Amazon Pinpoint × Growth Hack Use Cases](https://www.slideshare.net/AmazonWebServicesJapan/amazon-pinpoint-x)
- [www.acrovision.jp - What is Amazon Pinpoint? A Clear Explanation! Achieve Push Notifications at Low Cost!](https://www.acrovision.jp/service/aws/?p=1421)
- [qiita.com - First Amazon Pinpoint①~Overview~](https://qiita.com/mottie/items/ebd3ed7a1a1d78ac0e76)
- [qiita.com - First Amazon Pinpoint②~Implementation~](https://qiita.com/mottie/items/662f8c2938f5046471d9)
- [onetech.jp - What is AWS PINPOINT? Thorough Explanation of Features, Pricing, and Benefits!](https://onetech.jp/blog/what-is-aws-pinpoint-15773)
- [tec.tecotec.co.jp - What is Amazon Pinpoint (Push Notification Edition)](https://tec.tecotec.co.jp/entry/2021/01/28/090000)
- [coffee-tech-blog.com - Building an MA Platform with Amazon Pinpoint](https://coffee-tech-blog.com/email-newsletter-automation-aws/)
- [www.ragate.co.jp - Amazon Pinpoint Explained by AWS Experts: Simplify and Streamline Marketing with AWS](https://www.ragate.co.jp/blog/articles/11830)

# Impressions
I thought it is necessary to first organize who (operators, administrators, marketers, developers, etc.), what (message content), to whom, through which notification channel, when (by when) they want to notify, and the total volume of notifications. (It's obvious, but...)
