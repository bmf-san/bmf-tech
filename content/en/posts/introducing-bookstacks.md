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

Managing a personal book collection sounds trivial — until you are standing in a bookshop wondering whether you already own a particular title, or you have accumulated a stack of unread books you keep forgetting about. Existing apps often require an account or tedious manual entry.

Bookstacks takes a different approach: scan the barcode, get the book. No account, no subscription. The app uses the ISBN barcode to look up metadata from the OpenBD API (a free Japanese book data service), makes duplicate detection automatic, and lets you organise books with flexible labels.

The app is on the [App Store](https://apps.apple.com/jp/app/bookstacks-%E6%9C%AC%E6%A3%9A%E7%AE%A1%E7%90%86/id6760252143) — give it a try.

## Use Cases

- **Duplicate check** — Scan a book in the shop to see whether you already own it before buying
- **Tsundoku management** — Tag unread books with "Tsundoku" and pick the next one from the list
- **Reading log** — Mark finished books as "Read" to keep your shelves organised
- **Wish list** — Collect candidates under a "Want to Buy" label

## Key Features

![Home screen](/assets/images/posts/introducing-bookstacks/en/01_home.png)

### Register Books by Barcode

Scanning an ISBN barcode fetches title, author, category, and cover image automatically from the OpenBD API. The app accepts both ISBN-13 (978/979 prefix) and ISBN-10. If a book with the same ISBN already exists, the app detects the duplicate and blocks re-registration.

![Book detail](/assets/images/posts/introducing-bookstacks/en/02_detail.png)

### Organise with Labels

Five preset labels ship with the app: "Read", "Reading", "Tsundoku", "Read Later", and "Want to Buy". Users can create custom labels with any name. Tapping a label chip filters the list instantly.

![Labels](/assets/images/posts/introducing-bookstacks/en/05_labels.png)

### Bookshelf View — Grid and List

Books appear in a three-column grid with cover images by default. A toggle switches to a list view for denser browsing.

![Grid view](/assets/images/posts/introducing-bookstacks/en/01_home.png)
![List view](/assets/images/posts/introducing-bookstacks/en/03_list.png)

### Settings

The settings screen provides label management, version info, privacy policy, and bibliographic data provider details. The theme follows the system setting automatically.

![Settings](/assets/images/posts/introducing-bookstacks/en/04_settings.png)

## Tech Stack

| Layer | Technology |
|---|---|
| UI | Flutter (iOS) |
| State management / DI | Riverpod + riverpod_generator |
| Persistence | Hive |
| Barcode scanning | mobile_scanner |
| Book metadata | OpenBD API (http) |
| Image caching | cached_network_image |
| Testing | flutter_test / mocktail |

The app uses a four-layer Clean Architecture: Domain, Application, Infrastructure, and Presentation. Riverpod handles dependency injection and state management. Hive stores all data locally, so the app runs fully offline once the book data is in the local store.

## How ISBN Scanning and Book Registration Work

The registration flow from scan to persistence involves four steps.

1. **Scan** — `mobile_scanner` reads the barcode. Only ISBN-13 (978/979 prefix, 13 digits) and ISBN-10 (10 digits) pass through; all other formats get discarded.

2. **Metadata fetch** — `OpenBdDatasource` calls `https://api.openbd.jp/v1/get?isbn={isbn}`. The response is a JSON array; the first element is either a full book object or `null` (not found). The app extracts `isbn`, `title`, `author`, `category`, and `coverImageUrl`.

3. **Duplicate check** — The confirm screen watches the `booksNotifierProvider` and checks whether a book with the same ISBN exists. On a match, the register button becomes disabled and the app alerts the user.

4. **Persist** — The validated `Book` entity passes through the `AddBook` use case to the `BookRepository` interface and gets written to Hive. `BooksNotifier` then reloads the list and rebuilds the bookshelf grid.

## Summary

Bookstacks is an app I built to make managing a book collection as frictionless as possible. Scanning a barcode to register a book is something I find genuinely useful in everyday use.

Feel free to download it.

- **App Store**: [Bookstacks](https://apps.apple.com/jp/app/bookstacks-%E6%9C%AC%E6%A3%9A%E7%AE%A1%E7%90%86/id6760252143)
