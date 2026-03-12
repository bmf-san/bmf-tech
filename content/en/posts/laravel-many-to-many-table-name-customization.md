---
title: Excluding Table Names from Laravel's Many-to-Many Relations
description: 'An in-depth look at Excluding Table Names from Laravel''s Many-to-Many Relations, covering key concepts and practical insights.'
slug: laravel-many-to-many-table-name-customization
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Laravel
translation_key: laravel-many-to-many-table-name-customization
---

When designing many-to-many relationships, I thought I was following the documentation, but I had a slight misunderstanding.

# Here are three tables, right?

The tables in this case:
* events
* event_tags
* event_tag_event ← pivot table

For normal tables:
* events
* tags
* tag_event

You would typically set up the relationship according to the default conventions, but when using slightly unconventional names, there are some things to be cautious about.

# Let's take a look at the documentation
[Laravel 5.1 Eloquent: Relationships](https://readouble.com/laravel/5.1/ja/eloquent-relationships.html#many-to-many)

Ah, so I just need to provide a second argument!

```php
public function eventTags()
{
  // The second argument is the Pivot table!
  return $this->belongstoMany('App\Modles\EventTag', 'event_tag_event')->withTimestamps();
}
```

```php
public function events()
{
  return $this->belongsToMany('App\Models\Events');
}
```

When I launched tinker to check...

```
SQLSTATE[42000]: Syntax error or access violation: 1066 Not unique table/alias on relationship
```

I got an error. ヽ(´ー｀)ノ

# The second argument is the Pivot table name!

Could it be that I need to specify the Pivot table name?

[SQLSTATE[42000]: Syntax error or access violation: 1066 Not unique table/alias on relationship](http://stackoverflow.com/questions/31059595/sqlstate42000-syntax-error-or-access-violation-1066-not-unique-table-alias-o)

```php
public function eventTags()
{
  // The second argument is the Pivot table!
  return $this->belongstoMany('App\Modles\EventTag', 'event_tag_event')->withTimestamps();
}
```

```php
public function events()
{
  return $this->belongsToMany('App\Models\Events');
}
```

I didn't get an error this time. ヽ(´ー｀)ノ

# Thoughts
Recently, I've been more interested in the front end than the back end, and I can't sleep at night. ヽ(´ー｀)ノ