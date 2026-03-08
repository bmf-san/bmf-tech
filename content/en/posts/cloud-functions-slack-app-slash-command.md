---
title: Implementing Slash Command for Slack App Using Cloud Functions
slug: cloud-functions-slack-app-slash-command
image: /assets/images/posts/cloud-functions-slack-app-slash-command/188304723-637b0b8a-6253-45db-86c9-17b33131444b.png
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
Implementing a Slash Command for a Slack App using Cloud Functions.

Here is the boilerplate I created.

[go-slack-app-on-cloud-functions-boilerplate](https://github.com/bmf-san/go-slack-app-on-cloud-functions-boilerplate)

There are various ways to create a Slack App with a Slash Command, but I decided to try using Cloud Functions because it's inexpensive, easy, and serverless.

# Prerequisites
- Access to Google Cloud Platform
- gcloud command setup completed
  - You need to be able to use the gcloud command to deploy the application.
- Cloud Build API enabled
  - This is necessary for building the function to deploy to Cloud Functions.

# Creating a Function with Cloud Functions
Create a function in the Cloud Functions console.
Select **HTTP** as the trigger type, allow **unauthenticated invocations**, and check **HTTPS required**.

Make sure to note the trigger URL listed under the function details in Cloud Functions after deploying the function later.

# Preparing the Slack App
## Create a Slack App
Press **From scratch** at [Create an app](https://api.slack.com/apps?new_app=1).

<img width="721" alt="Screenshot 2022-09-04 17 29 50" src="https://user-images.githubusercontent.com/13291041/188304723-637b0b8a-6253-45db-86c9-17b33131444b.png">

Enter the **App Name**.

Select a workspace under **Pick a workspace to develop your app in:** and press **Create App**.

<img width="714" alt="Screenshot 2022-09-04 17 32 07" src="https://user-images.githubusercontent.com/13291041/188304790-225cfb1c-2a31-4627-8f57-35b856b4aed8.png">

## Configure Slash Command
In the settings screen (ex. https://api.slack.com/app/****), select Slash Commands.

<img width="720" alt="Screenshot 2022-09-04 17 33 21" src="https://user-images.githubusercontent.com/13291041/188304841-db000433-2a20-4e4a-b303-9a2fac7e3e7b.png">

Press **Create New Command** and enter **Command**, **Short Description**, **Usage Hint**, and **Escape channels, users, and links sent to your app** as desired.

Enter the trigger URL you noted earlier in the **Request URL** field.
The trigger URL will be in the format **https://REGION-NAME-PROJECT-ID.cloudfunctions.net/FUNCTION_NAME**.

<img width="568" alt="Screenshot 2022-09-04 17 35 59" src="https://user-images.githubusercontent.com/13291041/188304929-a17ccbf4-3194-490e-ad65-12c77c5f324a.png">

Once entered, press **Save**.

## Install the Slack App
In the settings screen (ex. https://api.slack.com/apps/****), press **Install App**.

<img width="738" alt="Screenshot 2022-09-04 17 37 16" src="https://user-images.githubusercontent.com/13291041/188304972-f35057c1-7392-429e-90f0-b1ec02e096b0.png">

Press **Install to workspace** to install the app in your chosen workspace.

## Obtain Signing Secret
In the settings screen (ex. https://api.slack.com/apps/****), press **Basic Information**.

In the App Credentials section, note the **Signing Secret** value.

# Implementing the Function
Implement the function to be deployed to Cloud Functions.

There were some tricky parts (like doing **go mod vendor**), but I will skip the implementation details.

Refer to the source code below.

[go-slack-app-on-cloud-functions-boilerplate](https://github.com/bmf-san/go-slack-app-on-cloud-functions-boilerplate)

# Deploying the Function
Follow the README of [go-slack-app-on-cloud-functions-boilerplate](https://github.com/bmf-san/go-slack-app-on-cloud-functions-boilerplate) to prepare environment variables and deploy.

# Verification
Try using the created Slash Command in Slack.

ex. 
**/hello Bob**

<img width="489" alt="Screenshot 2022-09-04 17 47 11" src="https://user-images.githubusercontent.com/13291041/188305315-5fe063c2-971b-4b18-978a-719596c2bd87.png">

# Thoughts
I would be happy if I could code the part of creating the Slack App.

# Postscript
I wanted to create a Slack Command for attendance management at work, so based on [go-slack-app-on-cloud-functions-boilerplate](https://github.com/bmf-san/go-slack-app-on-cloud-functions-boilerplate), I created [akashi-slack-slash-command](https://github.com/bmf-san/akashi-slack-slash-command).

In this implementation, the storage is set to Spreadsheet, but there is a management issue where the sharing settings of the Spreadsheet cannot be flexibly adjusted due to permissions when using Google Workspace, so at my workplace, we have adjusted the implementation to use Cloud Storage instead of Spreadsheet.

If there are organizations using Akashi for attendance management and Slack as a chat tool, this Slack Command should be easy to use.

The operational cost is not high, but scalability might be a bit questionable.

I think it should work fine for organizations with a few thousand people or less. Probably.