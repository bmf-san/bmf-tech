---
title: Struggles with Recursion
slug: recursion-challenges
date: 2023-07-08T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Recursion
translation_key: recursion-challenges
---

Recursion is a programmer's delight! It should be elegantly written and is expected to be straightforward! ...I wish I could say that with confidence, but honestly, I struggle with it.

The only time I really write recursive code is during coding quizzes, and in reality, I don't have many opportunities to write it.

In some cases, recursion can become complex without memoization or tail optimization, leading to high computational costs or memory-heavy code. However, depending on the algorithm, it can also result in simpler code.

That said, I feel the cognitive load remains high.

I thought about what increases this cognitive load and makes it difficult, and I realized two things that I want to note down.

# Return
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

When confused about returns, I think it's good to write out the process straightforwardly. I wonder if there are better methods.

# Stack
The evaluation of a recursive function is stacked in the call stack. Since it's a stack, it follows LIFO.

Once the evaluation is complete, data is popped from the stack for processing.

If you don't understand this, you can easily get lost while debugging the code.

For example, consider the following simple code, which is easier to think about.

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

It may seem very basic, but as the code becomes more complex, it can become quite confusing, so I want to return to this foundation in such cases.

# Thoughts
If I can derive recurrence relations and think mathematically, I feel that the difficulty might decrease with familiarity, but I still can't shake off my struggles with recursion, so practice is necessary...

I think an approach of first writing it with a for loop and then rewriting it recursively could work, but I believe there are cases where it can be challenging depending on the problem.

Once I become proficient in recursion through practice, I will revisit this article.