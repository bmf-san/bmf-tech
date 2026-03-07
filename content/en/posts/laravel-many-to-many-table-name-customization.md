---
title: When Excluding Table Names from Conventions in Laravel Many-to-Many Relationships
slug: laravel-many-to-many-table-name-customization
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Laravel
description: A story about a small misunderstanding when designing many-to-many relationships in Laravel.
translation_key: laravel-many-to-many-table-name-customization
---

When designing many-to-many relationships, I thought I was following the documentation correctly, but I made a small misunderstanding.

# Here are three tables, right?

The tables this time:
* events
* event_tags
* event_tag_event ← pivot table

Normally, the tables would be:
* events
* tags
* tag_event

If you follow the default conventions, you can create relationships like this. However, when you use slightly unconventional names, there are some things to watch out for.

# Checking the documentation
[Laravel 5.1 Eloquent: Relationships](https://readouble.com/laravel/5.1/ja/eloquent-relationships.html#many-to-many)

Oh, I see, you just need to provide the second argument.

```php
Event.php
public function eventTags()
{
  // The second argument is the pivot table!
  return $this->belongstoMany('App\Modles\EventTag', 'event_tag_event')->withTimestamps();
}
```

```php
EventTag.php
public function events()
{
  return $this->belongsToMany('App\Models\Events');
}
```

When I checked using tinker...

```
SQLSTATE[42000]: Syntax error or access violation: 1066 Not unique table/alias on relationship
```

It throws an error. ヽ(´ー｀)ノ

# The second argument is the pivot table name!

Could it be that I need to specify the pivot table name?

[SQLSTATE[42000]: Syntax error or access violation: 1066 Not unique table/alias on relationship](http://stackoverflow.com/questions/31059595/sqlstate42000-syntax-error-or-access-violation-1066-not-unique-table-alias-o)

```php
Event.php
public function eventTags()
{
  // The second argument is the pivot table!
  return $this->belongstoMany('App\Modles\EventTag', 'event_tag_event')->withTimestamps();
}
```

```php
EventTag.php
public function events()
{
  return $this->belongsToMany('App\Models\Events');
}
```

No more errors. ヽ(´ー｀)ノ

# Thoughts
Recently, I’ve been more interested in frontend than backend, and I can’t sleep at night. ヽ(´ー｀)ノ