---
title: Creating My Own Technology Radar
slug: create-technology-radar
date: 2022-12-17T00:00:00Z
author: bmf-san
categories:
  - Tools
tags:
  - Technology Radar
description: How to create your own Technology Radar.
translation_key: create-technology-radar
---

[Makuake Advent Calendar 2022](https://adventar.org/calendars/8496) - Day 7 article!

# Overview
This post explains how to create your own Technology Radar.

# What is a Technology Radar?
A Technology Radar is an analytical report on technology trends in software development, published by ThoughtWorks (a global software development and consulting company where Martin Fowler is a member).

[www.thoughtworks.com/radar](https://www.thoughtworks.com/radar)

The report is updated about twice a year and provides insights into technology trends.

You can also view the [Archive](https://www.thoughtworks.com/radar/archive).

The Technology Radar is structured into the following four quadrants:

- Techniques
  - e.g., design, development processes, etc.
- Tools
  - e.g., databases, version control systems, etc.
- Platforms
  - e.g., development environments, cloud platforms, etc.
- Languages and Frameworks
  - e.g., programming languages, application frameworks, etc.

Each quadrant is further categorized into the following four rings:

- Hold
  - On hold. Handle with caution. Requires investigation.
- Trial
  - Worth adopting. In trial phase.
- Assess
  - Worth investing resources to evaluate its value.
- Adopt
  - Strongly recommended for adoption.

On the [www.thoughtworks.com/radar](https://www.thoughtworks.com/radar) website, you can also view evaluation comments and the history of blips (movement between rings) for each technology.

# Creating Your Own Technology Radar
There are tools available to easily create your own Technology Radar.

[Build your own Radar](https://www.thoughtworks.com/radar/byor)

Here are two methods to create one:

## Using radar.thoughtworks.com
By entering a Google Spreadsheet link on [radar.thoughtworks.com](https://radar.thoughtworks.com/), you can generate a Technology Radar.

Note that radars created using this method will be public.

## Hosting it Yourself
You can also self-host your Radar using the repository available at [github.com - thoughtworks/build-your-own-radar](https://github.com/thoughtworks/build-your-own-radar).

A Docker image is provided, so here's how you can try it out using Docker:

### Clone the Sample Repository
You can fork or clone the original repository, but for convenience, I’ve prepared a sample repository.

[github.com - bmf-san/technology-radar-boilerplate](https://github.com/bmf-san/technology-radar-boilerplate)

#### 1. Clone the Repository
Clone [github.com - bmf-san/technology-radar-boilerplate](https://github.com/bmf-san/technology-radar-boilerplate).

#### 2. Start the Container
Run the following command:

`make run`

#### 3. Generate the Radar
Access `http://localhost:8080`, enter `http://localhost:8080/files/radar.json`, and click `Build My Radar`.

![form](https://user-images.githubusercontent.com/13291041/205632529-e528abd5-9013-458e-a540-6fc1251867e9.png)

#### 4. Play Around with the Generated Radar
Once the Radar is generated, you’ll be redirected to a link like this:

`http://localhost:8080/?sheetId=http%3A%2F%2Flocalhost%3A8080%2Ffiles%2Fradar.json`

![radar](https://user-images.githubusercontent.com/13291041/205632536-d39195f1-2570-4645-bfb4-869bc7f77454.png)

You can adjust the content displayed on the Radar by editing `./files/radar.json`.

(Originally, I wanted to make it possible to provision the JSON file, but due to frontend build constraints, it seemed difficult...)

You can also print the Radar from the `Print this radar` button in the top-right corner.

# Thoughts
Creating a Radar like this to publicly share the technology stack adopted or evaluated by an organization or team seems like a great initiative. It can be used to define the reasons for technology selection, the selection and evaluation process, and to clarify what to invest in as part of a technology portfolio.

As mentioned in [www.oreilly.co.jp - Fundamentals of Software Architecture](https://www.oreilly.co.jp/books/9784873119823/), creating and regularly updating such a Radar as an individual also seems like a good practice. Mapping out the technologies you are following might make you realize how narrow your perspective is, though...