---
title: Introducing Bookstacks — An iPhone Bookshelf App with ISBN Barcode Scanning
description: 'A deep dive into Bookstacks, an iPhone app built with Flutter and Riverpod that lets you register books via ISBN barcode scan and organise them with customisable labels.'
slug: introducing-bookstacks
date: 2026-03-20T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Flutter
  - iOS
  - Riverpod
  - Dart
translation_key: introducing-bookstacks
---

# Introducing Bookstacks — An iPhone Bookshelf App with ISBN Barcode Scanning

## Why I Built It

Managing a personal book collection sounds trivial — until you are standing in a bookshop wondering whether you already own a particular title, or you have accumulated a stack of books you intended to read but keep forgetting about. Existing apps either require manual entry or rely on services that need accounts and internet connections.

Bookstacks takes a different approach: scan the barcode, get the book. No account, no subscription. The app uses the ISBN barcode to look up metadata from the OpenBD API (a free Japanese book data service), makes duplicate detection automatic, and organises books with flexible labels.

## Architecture

The app follows Clean Architecture — Domain, Application, Infrastructure, and Presentation — wired together by Riverpod providers. Hive handles local persistence; the app runs fully offline once it caches the book data. The OpenBD API is the only external dependency.

![Home screen](/assets/images/posts/introducing-bookstacks/01_home.png)

## ISBN Barcode Scan × OpenBD API — The Async Flow

The registration flow from scan to persistence is the most interesting technical piece in the app. The sequence is:

1. **Scan** — `mobile_scanner` delivers a raw `BarcodeCapture` event with detected values. The app accepts both ISBN-13 (978/979 prefix, 13 digits) and ISBN-10 (10 digits); all other barcode formats are ignored.

2. **Metadata fetch** — `OpenBdDatasource` calls the OpenBD endpoint `https://api.openbd.jp/v1/get?isbn={isbn}` with an `http.Client`. The response is a JSON array; the first element is either a full book object or `null` (not found). `OpenBdDatasource` parses the response and returns `OpenBdBookData` with `isbn`, `title`, `author`, `category`, and `coverImageUrl`.

3. **Duplicate check** — The confirm screen watches the `booksNotifierProvider` and checks whether a book with the same ISBN already exists. On a match, the screen disables the register button and alerts the user.

4. **Persist** — The app passes the validated `Book` entity through the `AddBook` use case to the `BookRepository` interface, writing it to Hive. `BooksNotifier` reloads the list, triggering a rebuild of the bookshelf grid.

The scan screen calls `OpenBdDatasource` directly for metadata retrieval. Persistence flows through the `AddBook` use case, so the presentation layer never touches Hive directly.

![Book detail](/assets/images/posts/introducing-bookstacks/02_detail.png)

## Key Features

### Bookshelf View — Grid and List

The app displays books in a three-column grid by default, with cover images loaded lazily via `cached_network_image`. A toggle switches to a list view for denser browsing.

![Grid view](/assets/images/posts/introducing-bookstacks/01_home.png)
![List view](/assets/images/posts/introducing-bookstacks/03_list.png)

### Label System

Bookstacks ships with five preset labels: “Read”, “Reading”, “Tsundoku” (to-read pile), “Read Later”, and “Want to Buy”. Users can create custom labels with a free-form name. Hive stores labels and links them to books via a many-to-one association; tapping a label chip filters the list instantly.

![Labels](/assets/images/posts/introducing-bookstacks/05_labels.png)

### Settings

The settings screen provides label management, version info, privacy policy, and bibliographic data provider details. The app theme follows the system setting automatically. The home screen displays an AdMob banner ad with an auto-retry mechanism (up to three attempts, five seconds apart) to handle temporary load failures.

![Settings](/assets/images/posts/introducing-bookstacks/04_settings.png)

## Summary

Bookstacks is available on the App Store.

- **App Store**: [Bookstacks](https://apps.apple.com/jp/app/bookstacks-%E6%9C%AC%E6%A3%9A%E7%AE%A1%E7%90%86/id6760252143)
