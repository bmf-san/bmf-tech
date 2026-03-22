---
title: Introducing Done Log — A Daily Routine Tracker iPhone App
description: 'An introduction to Done Log, an iPhone app built with Flutter and Riverpod for tracking daily routine tasks with a flexible recurrence rule engine — daily, every N days, specific weekdays, or once.'
slug: introducing-donelog
date: 2026-03-20T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Flutter
  - iOS
  - Riverpod
  - Dart
translation_key: introducing-donelog
---

# Introducing Done Log — A Daily Routine Tracker iPhone App

## Why I Built It

General-purpose task managers like a to-do app are great for projects, but they are overkill for checking whether you took your medication this morning or locked the front door tonight. Done Log focuses on a single job: record daily routine tasks with one tap, and automatically reset them according to a recurrence rule so there is nothing to manage.

## Architecture

Done Log follows Clean Architecture — Domain, Application, Infrastructure, and Presentation — connected by Riverpod providers. Hive stores tasks and completion history locally; the app works entirely offline. The app persists completion history separately from task definitions so that past records are never lost when a recurrence rule changes.

![Today's task list](/assets/images/posts/introducing-donelog/01_today.png)

## The Recurrence Rule Engine

The most technically interesting part of the app is the `RecurrenceRule` domain entity and its `shouldShowToday()` method. This pure-Dart function determines whether a task should appear on today's list — no framework involvement, no side effects. It receives the current timestamp and the time the task was last completed, and returns a boolean.

By keeping this logic in the domain layer as a plain function, testing it in isolation requires no infrastructure mocks. All four recurrence types — `daily`, `everyNDays`, `weekdays`, and `once` — resolve within a single `switch` expression with no shared mutable state.

Reset happens at app startup via the `CheckAndResetTasks` use case: it iterates every task, calls `shouldShowToday()`, and updates completion state in Hive if the rule says the task is ready again.

## Key Features

### Today View

Tasks that match the current date and recurrence rule appear under "Today". The app shows completed tasks with a strikethrough and a checkmark, giving an instant visual progress indicator.

![Completed task](/assets/images/posts/introducing-donelog/02_task_completed.png)

### Task Registration Form

The task form exposes all four recurrence types. For `everyNDays`, a slider lets users pick any interval from 2 to 30 days. For `weekdays`, day checkboxes allow arbitrary week patterns (e.g. Mon, Wed, Fri only).

![Task form](/assets/images/posts/introducing-donelog/03_task_form.png)

### Calendar View

Past completions appear on a monthly calendar. Tapping any date shows which tasks the user finished that day, making it easy to review habit streaks or spot missed days.

![Calendar view](/assets/images/posts/introducing-donelog/04_calendar.png)

### Settings and Dark Mode

Language (Japanese / English) and light/dark/system theme are configurable. Schedule notification reminders for any time of day.

![Settings](/assets/images/posts/introducing-donelog/05_settings.png)
![Dark mode](/assets/images/posts/introducing-donelog/06_dark_mode.png)

## Summary

Done Log is available on the App Store.

- **App Store**: [Done Log](https://apps.apple.com/jp/app/done-log/id6759606196)
