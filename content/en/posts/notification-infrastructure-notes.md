---
title: Notes on Building a Notification Infrastructure
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
This post summarizes my thoughts and research on building a notification infrastructure.

# What is a Notification Infrastructure?
A system infrastructure for sending notifications (email, push, SMS, voice, etc.) to users.

It processes requests from clients (systems requesting notifications) regarding the destination and content of the notifications.

# Considerations for Designing and Implementing Notification Infrastructure
I thought there would be a lot to consider, so I wrote down my thoughts in a rough order. It's not organized.

- Nature of Notification Messages
  - Types
    - Important Announcements
      - e.g., changes to terms, payment-related notifications
    - Marketing
      - e.g., sale information
  - Depending on the type, the characteristics emphasized by the architecture may change.
    - For example, for emails, announcement types might use SES, while marketing-related notifications might use Amazon Pinpoint, suggesting that optimal forms may vary based on characteristics such as performance and cost.
- Notification Channels
  - Email, push, SMS, voice, etc.
- Sending Patterns
  - Individual sending
  - Group sending
  - Cancellation
    - A mechanism to retract sent notifications
    - Also serves as a fail-safe
      - A mechanism to prevent human errors like setting 1000 instead of 1 for sending. Such mistakes are likely to directly impact costs.
  - Notification Priority
  - Conditional Specifications
    - If notifications can be sent via email, then SMS notifications may not be sent, etc.
    - Time specifications for when not to send notifications.
- Cost
  - Mechanisms and ideas to keep costs down.
- Scalability
  - While the notification infrastructure itself is important, care must also be taken with the internal systems it interacts with.
    - For example, if user-related information needs to be requested, performance considerations for that system will also be necessary.
      - The fewer external systems (services outside the notification infrastructure) that are dependent, the better.
- Messaging
  - Queuing
  - Scheduling
- Management of Destination User Information
  - Destination information
  - Opt-in/Opt-out
- Authentication
  - Necessary when accessing personal data or integrating with authentication-related APIs?
- Notification Message Templates
  - Templates for embedding static and dynamic data
  - Consideration of whether multilingual support is needed.
- Error Handling and Retry
  - How to catch errors such as sending failures and how to resend.
  - Bounce handling for emails, etc.
- Monitoring
  - Not only system metrics but also notification-related data (e.g., delivery rates, click rates, etc.)
  - Ability to track notifications.
    - Is issuing an ID like a trace ID necessary?
- Extensibility
  - Adding notification channels, customization of templates, etc.
- Integration with External Services
  - Design to be replaceable and not overly dependent.
  - Integration with analytics infrastructure, etc.
    - If optimization of notifications can be achieved, it would reduce the manual operation of notifications.
    - A desire to automate as much as possible.
- Operations
  - It's obvious, but consider operations (from both developer and business perspectives).
    - What operations are currently in place, and what operations will be in the future?
- Testing
  - Ease of testing and debugging.
    - It could be complex due to the involvement of multiple systems.

# Solutions
## SaaS
There are various options like Salesforce, Braze, Airship, OneSignal, SendGrid, etc.

These not only send notifications across multiple channels but also integrate with customer management and marketing tools.

Many services that support multi-channel notifications can be found.

## PaaS
AWS Pinpoint is a platform that supports notifications across multiple channels.

# Case Studies
I looked into case studies from both domestic and international sources.

- [techblog.zozo.com - Introduction to Real-Time Marketing System and Its Replacement Plan](https://techblog.zozo.com/entry/real-time-marketing-system)
  - A case that builds a foundation including MA elements while supporting multi-channel like AWS Pinpoint.
- [techblog.zozo.com - Replacing the Real-Time Data Integration Infrastructure Supporting Marketing Automation System](https://techblog.zozo.com/entry/ma-realtime-data-infrastructure-replacement)
  - A follow-up on the above, discussing the subsequent replacement.
- [techblog.zozo.com - Rule-Based Optimization Improvements in Personalized Delivery](https://techblog.zozo.com/entry/improving-optimization-for-personalized-marketing)
  - Amazing. I thought this is the essence of notifications.
  - To reach this point, the impact of notifications on business must have been sufficiently estimated, and it seems that clever notifications can significantly impact business.
  - "There is still much room for improvement to achieve the true goal of delivering only the notifications that users truly want."
- [leandrofranchi.medium.com - How to design a Notification System](https://leandrofranchi.medium.com/how-to-design-a-notification-system-23f381cdeb00)
  - An example of system architecture design for a multi-channel notification infrastructure.
  - Access to user-related data is not particularly considered, focusing solely on notification delivery.
  - My impression is that the configuration would look something like this.
  - It may not be necessary to prepare a common interface for notifications from the start; it might be reasonable to discover the interface as each notification channel is created and settled down.
- [www.notificationapi.com - Notification Service Design - with diagrams](https://www.notificationapi.com/blog/notification-service-design-with-architectural-diagrams)
  - Similar to the above, a design example.
  - The overall structure is similar to the above. This one considers user settings, while the above design considers logging.
- [cloudificationzone.com - Notification System Design](https://cloudificationzone.com/2021/08/13/notification-system-design/)
  - A more concretized design example.
  - It was unclear what kind of notifications were specifically inbound.
- [atmarkit.itmedia.co.jp - Basic Knowledge of Push Notifications & Architecture and Mechanism of a Push Notification Infrastructure Exceeding 10,000 per Second](https://atmarkit.itmedia.co.jp/ait/articles/1412/18/news022.html)
  - A push notification infrastructure utilizing DynamoDB and Node.js.
- [zenn.dev - Handling 100,000 requests per minute, Email/Push Notification Mass Distribution AWS Architecture](https://zenn.dev/coconala/articles/a3a5e33cd1d984)
  - A story about switching to a more performant architecture while keeping costs down.
  - By distributing requests for the delivery API, which takes the most time, scalability improves.
    - As the number of delivery workers increases, scalability rises linearly.
- [www.slideshare.net - Architecture of Push Notification Infrastructure for System Acceleration Forum](https://www.slideshare.net/recruitcojp/ss-42921628)

# AWS Pinpoint
I personally researched AWS Pinpoint, which I am interested in.

- An AWS service for messages (notifications) that supports multiple channels. Released in 2016.
  - Supports push notifications, emails, SMS, and voice messages.
  - Additionally, using the [Custom Channel](https://docs.aws.amazon.com/ja_jp/pinpoint/latest/userguide/channels-custom.html) feature allows for extending notification channels.
    - For example, adding Facebook Messenger.
- Pay-as-you-go
  - It seems not as expensive as I thought.
  - Emails cost $1.00 per 10,000, push notifications are free for the first million, then $1.00 per million thereafter.
  - If I roughly estimate for a user base of several million, it might become a considerable amount.
    - Of course, it's necessary to properly verify the cost-effectiveness. Quite strictly.
- Notification analysis is also possible.
  - It seems to be able to integrate with marketing initiatives.
    - However, since the capabilities seem limited, compatibility should be carefully considered.
- Scalability
  - There is a limit to the number of notifications that can be sent per second.
    - Quotas can be increased upon request.
  - The scaling of the part that will likely be the bottleneck in sending notifications can almost entirely be delegated to AWS.

## References
- [docs.aws.amazon.com - What is Amazon Pinpoint](https://docs.aws.amazon.com/ja_jp/pinpoint/latest/developerguide/welcome.html)
- [pages.awscloud.com - Capture Users with Amazon Pinpoint
~Amazon Digital User Engagement~](https://pages.awscloud.com/rs/112-TZM-766/images/A3-01.pdf)
- [www.slideshare.net - Amazon Pinpoint × Growth Hack Use Case Collection](https://www.slideshare.net/AmazonWebServicesJapan/amazon-pinpoint-x)
- [www.acrovision.jp - What is Amazon Pinpoint? A Simple Explanation! Achieve Low-Cost Push Notifications!](https://www.acrovision.jp/service/aws/?p=1421)
- [qiita.com - First Time with Amazon Pinpoint①~Overview~](https://qiita.com/mottie/items/ebd3ed7a1a1d78ac0e76)
- [qiita.com - First Time with Amazon Pinpoint②~Implementation~](https://qiita.com/mottie/items/662f8c2938f5046471d9)
- [onetech.jp - What is AWS PINPOINT? Thorough Explanation of Features, Pricing, and Benefits!](https://onetech.jp/blog/what-is-aws-pinpoint-15773)
- [tec.tecotec.co.jp - What is Amazon Pinpoint (Push Notification Edition)](https://tec.tecotec.co.jp/entry/2021/01/28/090000)
- [coffee-tech-blog.com - Building an MA Infrastructure with Amazon Pinpoint](https://coffee-tech-blog.com/email-newsletter-automation-aws/)
- [www.ragate.co.jp - AWS Experts Explain Amazon Pinpoint: Simplifying and Streamlining Marketing with AWS](https://www.ragate.co.jp/blog/articles/11830)

# Thoughts
I think it is necessary to first organize who (operators, administrators, marketers, developers, etc.), what (message content), to whom, through which notification channel, when (by when) they want to notify, and what the total volume of notifications is. (It seems obvious, but...)