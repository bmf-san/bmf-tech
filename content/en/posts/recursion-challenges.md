---
title: Challenges with Recursion
description: 'Understand recursion challenges involving return statements and call stack evaluation with practice strategies.'
slug: recursion-challenges
date: 2023-07-08T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Recursion
translation_key: recursion-challenges
---


Recursion is a programmer's art! It's expected to be written elegantly! ...is what I'd like to confidently say, but honestly, I'm not very good at it.

The only time I usually write recursive processes is during coding quizzes, and in reality, I don't have many opportunities to write them.

Depending on the case, recursion can become computationally expensive or just memory-consuming if not done elegantly with memoization or tail optimization. However, it has the advantage of resulting in simple code for certain algorithms.

Yet, I feel the cognitive load remains high.

I've thought about what increases this cognitive load and what makes it challenging, and I've realized two things, which I'll write about here.

# return
It's hard to grasp what gets returned and when.

For example, the following case is simple and easy to understand, but when there are multiple recursive cases, the cognitive load increases.

```go
package main

import "fmt"

func fact(n int) int {
        // Base case
	if n < 2 {
		return 1
	}
        // Recursive case
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(5)) // 120
}
```

When confused by returns, it might be good to write out the process straightforwardly. Is there a better way?

# Stack
The evaluation of recursive functions is pushed onto the call stack. Since it's a stack, it's LIFO.

If you don't understand this, you might get lost even when tracing the code with a debugger.

For example, the following simple code is easy to think about.

```go
package main

import "fmt"

func proc(n int) {
	if n == 0 {
		return
	} else {
		fmt.Printf("%d", n)
		proc(n - 1)
		fmt.Printf("%d", n)
		return
	}
}

func main() {
	proc(5) // 5432112345
}
```

It might be very basic, but the more complex the code becomes, the more confusing it gets, so I want to return to these basics when that happens.

# Impressions
If you can derive recurrence relations and think mathematically, and if you're used to it, the difficulty might feel reduced, but I still can't shake off the difficulty with recursion, so I need practice...

I think an approach where you first write it as a for loop and then rewrite it as a recursive process might work, but I believe there are cases where it would still be challenging depending on the problem.

Once I become proficient with recursion through practice, I'll look back at this article again.