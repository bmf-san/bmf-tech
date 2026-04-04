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

I had been managing routine tasks in Google Keep, but there was no way to track daily completion, so I built a task management app focused specifically on routine tasks.

It records task completion with one tap and automatically resets tasks according to a recurrence rule. The focus was on reducing management overhead and making logging as easy as possible.

The app is on the [App Store](https://apps.apple.com/jp/app/done-log/id6759606196) — give it a try.

## Use Cases

- **Daily medication check** — Record whether you took your morning pills and see your progress for the day
- **Weekly routine management** — Manage repeating tasks tied to specific weekdays, such as exercise on Mon, Wed, and Fri
- **Habit review** — Look back at past completions on the calendar view to track how well habits are sticking
- **Every-N-days routines** — Register tasks that repeat at a fixed interval, like cleaning every two weeks

## Key Features

![Today's task list](/assets/images/posts/introducing-donelog/en/01_today.png)

### Today's Task View

Tasks that match the current date and recurrence rule appear under "Today". The app shows completed tasks with a strikethrough and a checkmark, giving an instant visual progress indicator.

![Completed task](/assets/images/posts/introducing-donelog/en/02_task_completed.png)

### Task Registration Form

The task form exposes all four recurrence types. For `everyNDays`, a slider lets users pick any interval from 2 to 30 days. For `weekdays`, day checkboxes allow arbitrary week patterns (e.g. Mon, Wed, Fri only).

![Task form](/assets/images/posts/introducing-donelog/en/03_task_form.png)

### Calendar View

Past completions appear on a monthly calendar. Tapping any date shows which tasks the user finished that day, making it easy to review habit streaks or spot missed days.

![Calendar view](/assets/images/posts/introducing-donelog/en/04_calendar.png)

### Settings and Dark Mode

Language (Japanese / English) and light/dark/system theme are configurable. Schedule notification reminders for any time of day.

![Settings](/assets/images/posts/introducing-donelog/en/05_settings.png)
![Dark mode](/assets/images/posts/introducing-donelog/en/06_dark_mode.png)

## Tech Stack

| Layer | Technology |
|---|---|
| UI | Flutter (iOS) |
| State management / DI | Riverpod + riverpod_generator |
| Persistence | Hive |
| Calendar | table_calendar |
| Notifications | flutter_local_notifications |
| Testing | flutter_test / mockito |

The app uses a four-layer Clean Architecture: Domain, Application, Infrastructure, and Presentation. Riverpod handles dependency injection and state management. Hive stores all data locally, so the app works fully offline. Task definitions and completion history live in separate tables, so past records are never lost when a recurrence rule changes.

## The Recurrence Rule Engine

The recurrence rule engine decides whether a given task should appear on today's list. It supports four rule types — daily, every N days, specific weekdays, and once — so each task can follow its own rhythm and reset automatically.

The rule logic lives entirely in the domain layer, separate from the database and UI. This makes it straightforward to change a rule or add a new type without affecting the rest of the app.

Reset runs automatically at app startup. The next morning, tasks are already reset and ready to go — no manual intervention needed.

## Summary

Done Log is an app I built to manage daily routine tasks with as little friction as possible. The one-tap recording and automatic resets from the recurrence rule engine are the features I find most useful in everyday use.

Feel free to download it.

- **App Store**: [Done Log](https://apps.apple.com/jp/app/done-log/id6759606196)
