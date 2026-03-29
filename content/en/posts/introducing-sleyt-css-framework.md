---
title: Introducing sleyt — A Minimal CSS Framework with Glass Morphism
description: 'An introduction to sleyt, a pure-CSS minimal framework featuring glass morphism design, built-in data visualisation components (bar, line, donut charts), a comprehensive component library, and full dark-mode support — with zero JavaScript runtime dependency.'
slug: introducing-sleyt-css-framework
date: 2026-03-20T00:00:00Z
author: bmf-san
categories:
  - Tools
tags:
  - CSS
  - Frontend
translation_key: introducing-sleyt-css-framework
---

# Introducing sleyt — A Minimal CSS Framework with Glass Morphism

## Why I Built It

I wanted a lightweight, simple CSS framework for my own projects, so I built one.

It focuses on being small, pure CSS with no JavaScript runtime dependency, and a modern visual design.

Some components and flexibility gaps remain, but it works well enough for practical use.

![Documentation site](/assets/images/posts/introducing-sleyt-css-framework/01_docs_home.png)

## Component Library

sleyt ships 20+ ready-to-use components split across four directories:

**Components** (`src/components/`): accordion, alerts, badges, buttons, cards, charts, code, forms, modals, navbar, navigation, progress, prose, sidebar, spinners, showcase, swatch, tables, tabs, tooltip.

**Layout** (`src/layout/`): container, dashboard, flexbox, grid.

**Base** (`src/base/`): reset, themes, variables.

**Utilities** (`src/utilities/`): spacing, colours, typography, borders, effects, glass, display, position, transforms, transitions.

![Component library — Buttons](/assets/images/posts/introducing-sleyt-css-framework/06_demo_components.png)

![Blog demo](/assets/images/posts/introducing-sleyt-css-framework/03_demo_blog.png)

### Data Visualisation

`charts.css` provides CSS-only bar charts, line charts, and donut charts — no JavaScript chart library required. Bar charts are pure CSS, using element height to represent data values. Line charts use SVG `<path>` elements styled with CSS. Donut chart segments are SVG `<circle>` elements drawn with `stroke-dasharray` and `stroke-dashoffset`. Data values flow in via CSS custom properties.

![Dashboard demo](/assets/images/posts/introducing-sleyt-css-framework/05_demo_dashboard.png)

![Dashboard charts demo](/assets/images/posts/introducing-sleyt-css-framework/08_demo_charts.png)

### Glass Morphism

Use `.glass`, `.glass-light`, `.glass-heavy`, and `.frosted` utility classes to add transparency and backdrop blur to any element. Stepped `backdrop-blur` utilities give fine-grained control over blur intensity.

![Glass morphism utilities](/assets/images/posts/introducing-sleyt-css-framework/07_docs_glass.png)

### Dark Mode

sleyt includes dark mode built-in, activating automatically via `@media (prefers-color-scheme: dark)`. Component colours use CSS custom properties (variables), so overriding a palette takes a single-file change.

### Accessibility

Component markup follows semantic HTML5 patterns. Buttons use real `<button>` elements, navigation uses `<nav>`, colour contrast ratios target WCAG AA.

## Installation

```bash
npm install sleyt
```

Then import in your CSS:

```css
@import "sleyt/dist/css/index.css";
```

Or use the CDN — add this to your HTML `<head>`:

```html
<link rel="stylesheet" href="https://unpkg.com/sleyt@latest/dist/css/index.css">
```

## Demo and Documentation

The [demo page](https://bmf-san.github.io/sleyt/demo.html) shows every component in context, including dark-mode variants and chart types. Three full-page demos — a blog listing, an article detail, and an admin dashboard — show how sleyt looks in realistic UI patterns.

![Blog article detail demo](/assets/images/posts/introducing-sleyt-css-framework/04_demo_blog_detail.png)

The [documentation](https://bmf-san.github.io/sleyt/) covers installation, customisation via CSS variables, and component usage examples.

## Summary

sleyt is a lightweight CSS-only framework focused on readability and modern visual style with no JavaScript runtime cost.

- **npm**: [sleyt](https://www.npmjs.com/package/sleyt)
- **GitHub**: [bmf-san/sleyt](https://github.com/bmf-san/sleyt)
- **Demo**: [bmf-san.github.io/sleyt/demo.html](https://bmf-san.github.io/sleyt/demo.html)
