---
title: Implementing Slack App Slash Command Using Cloud Functions
slug: cloud-functions-slack-app-slash-command
date: 2022-09-19T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - slack-bot
  - Slack
  - Golang
  - Google Cloud Platform
  - Cloud Functions
translation_key: cloud-functions-slack-app-slash-command
---

# Overview
Implement a Slash Command for a Slack App using Cloud Functions.

Here is the boilerplate I created this time.

[go-slack-app-on-cloud-functions-boilerplate](https://github.com/bmf-san/go-slack-app-on-cloud-functions-boilerplate)

There are various ways to create a Slack App with Slash Commands, but I tried using CloudFront Functions because it is cheap, easy, and serverless.

# Prerequisites
- Google Cloud Platform is available
- The gcloud command is set up
  - The gcloud command must be available to deploy the application
- Cloud Build API is enabled
  - It is necessary to build the function to deploy functions to Cloud Functions.

# Create a Function with Cloud Functions
Create a function in the Cloud Functions console.
Select **HTTP** as the trigger type, allow **unauthenticated invocation**, and check **HTTPS required**.

After deploying the function, note the trigger URL listed under the function details > Trigger in Cloud Functions, as it will be used later.

# Prepare the Slack App
## Create a Slack App
At [Create an app](https://api.slack.com/apps?new_app=1), click **From scratch**.

<img width="721" alt="Screenshot 2022-09-04 17 29 50" src="/assets/images/posts/cloud-functions-slack-app-slash-command/188304723-637b0b8a-6253-45db-86c9-17b33131444b.png">

Enter the **App Name**.

Select a workspace under **Pick a workspace to develop your app in:** and click **Create App**.

<img width="714" alt="Screenshot 2022-09-04 17 32 07" src="/assets/images/posts/cloud-functions-slack-app-slash-command/188304790-225cfb1c-2a31-4627-8f57-35b856b4aed8.png">

## Configure Slash Command
In the settings screen (e.g., https://api.slack.com/app/****), select Slash Commands.

<img width="720" alt="Screenshot 2022-09-04 17 33 21" src="/assets/images/posts/cloud-functions-slack-app-slash-command/188304841-db000433-2a20-4e4a-b303-9a2fac7e3e7b.png">

Click **Create New Command** and enter **Command**, **Short Description**, **Usage Hint**, and **Escape channels, users, and links sent to your app** as desired.

Enter the trigger URL noted earlier as the **Request URL**.
The trigger URL is in the format **https://REGION-NAME-PROJECT-ID.cloudfunctions.net/FUNCTION_NAME**.

<img width="568" alt="Screenshot 2022-09-04 17 35 59" src="/assets/images/posts/cloud-functions-slack-app-slash-command/188304929-a17ccbf4-3194-490e-ad65-12c77c5f324a.png">

Once entered, click **Save**.

## Install the Slack App
In the settings screen (e.g., https://api.slack.com/apps/****), click **Install App**.

<img width="738" alt="Screenshot 2022-09-04 17 37 16" src="/assets/images/posts/cloud-functions-slack-app-slash-command/188304972-f35057c1-7392-429e-90f0-b1ec02e096b0.png">

Click **Install to workspace** to install the app in any workspace.

## Obtain Signing Secret
In the settings screen (e.g., https://api.slack.com/apps/****), click **Basic Information**.

Under App Credentials, there is a **Signing Secret**, so note the value.

# Implement the Function
Implement the function to be deployed to Cloud Functions.

There are some tricky parts (like doing **go mod vendor**), but the implementation details are omitted.

Refer to the source code below.

[go-slack-app-on-cloud-functions-boilerplate](https://github.com/bmf-san/go-slack-app-on-cloud-functions-boilerplate)

# Deploy the Function
Follow the README of [go-slack-app-on-cloud-functions-boilerplate](https://github.com/bmf-san/go-slack-app-on-cloud-functions-boilerplate) to prepare environment variables and deploy.

# Verify Operation
Try using the created Slash Command in Slack.

ex. 
**/hello Bob**

<img width="489" alt="Screenshot 2022-09-04 17 47 11" src="/assets/images/posts/cloud-functions-slack-app-slash-command/188305315-5fe063c2-971b-4b18-978a-719596c2bd87.png">

# Thoughts
It would be nice to be able to code the part of creating a Slack App.

# Additional Notes
I wanted to create a Slack Command for attendance management at work, so I created [akashi-slack-slash-command](https://github.com/bmf-san/akashi-slack-slash-command) based on [go-slack-app-on-cloud-functions-boilerplate](https://github.com/bmf-san/go-slack-app-on-cloud-functions-boilerplate).

In this implementation, storage is set to Spreadsheet, but when using Google Workspace, there is a management issue where the sharing settings of the Spreadsheet cannot be flexibly adjusted due to permissions, so at work, we replaced the storage from Spreadsheet to Cloud Storage and adjusted the implementation.

If an organization uses Akashi for attendance management and Slack as a chat tool, it should be a Slack Command that can be easily used.

The operational cost is not significant, but scalability might be a bit questionable.

It should probably work fine unless the organization exceeds several thousand people. Probably.