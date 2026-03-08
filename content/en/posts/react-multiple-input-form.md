---
title: Implementing a Form with Multiple Inputs in React
slug: react-multiple-input-form
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - React
translation_key: react-multiple-input-form
---

Implementing forms in React can be a bit quirky.

First, it might be helpful to understand that there are two patterns for forms in React: controlled forms and uncontrolled forms. Reference: [React Documentation](https://facebook.github.io/react/docs/forms-ja-JP.html)

I still have a lot to learn, but since there are few implementation examples for React, I want to publish this article in hopes of providing some inspiration. 
(I would appreciate any suggestions for easier methods or better approaches.)

While researching for this implementation, I found that most examples only dealt with a single input, which was quite challenging...

# What to Do
* Implement a form that handles multiple inputs

# What Not to Do
* Form submission... I will only check if the data is being retrieved correctly.
* Implementation using bind... I don't quite understand bind...

# Implementation

Most examples that handle multiple inputs seem to use bind. Many of the examples found in the documentation or through searches often used bind.

However, I still don't quite understand this bind...

There were many examples using simple, non-nested object states, but I wasn't sure how to handle nested objects (cry).

After a lot of searching, I found this article: [Stack Overflow - Best practice for ReactJS form components](http://stackoverflow.com/questions/26626454/best-practice-for-reactjs-form-components)

So, I will implement it while referring to the above.

The HTML looks like this:

```form.html
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8" />
  <title>React Multi Input Form</title>
</head>
    <body>

    <div id="multi-input-form"></div>

    <!-- scripts -->
    <script src="path/to/react.js"></script>
    <script src="path/to/react-dom.js"></script>
    <script src="path/to/browser.min.js"></script>

    </body>
</html>
```

And the React code looks like this:

```form.js
var FormApp = React.createClass({
    getInitialState: function () {
        // Prepare a nested object
        return {
            data: {
                name: '',
                email: ''
            }
        };
    },

    handleChange: function (event) {
        // Access the nested object's data
        var data = this.state.data;

        // Process values based on the name attribute of the event
        switch (event.target.name) {
            case 'name':
                data.name = event.target.value;
                break;
            case 'email':
                data.email = event.target.value;
                break;
        }

        // Update the state
        this.setState({
            data: data
        });
    },

    handleSubmit: function () {
        console.log(this.state.data.name);
        console.log(this.state.data.email);
    },

    render: function () {
        return (
            <form action="javascript:void(0)" onSubmit={this.handleSubmit}>
                {/* Name */}
                <label htmlFor="name">Name</label>
                <input type="text" name="name" value={this.state.name} onChange={this.handleChange} />

                {/* Email */}
                <label htmlFor="email">Email Address</label>
                <input type="email" name="email" value={this.state.email} onChange={this.handleChange} />

                 {/* Submit Button */}
                <button type="submit">Submit</button>
            </form>
        );
    }
});

ReactDOM.render(
    <FormApp />,
    document.getElementById('multi-input-form')
);
```

As I mentioned at the beginning, I thought it was a hassle, but it seems surprisingly simple. Haha.
There doesn't seem to be anything particularly noteworthy.

The operation looks like this:

![react-form.gif](https://qiita-image-store.s3.amazonaws.com/0/124495/9f1e3bb6-d5ed-eec8-7605-cfdeb54ae4b2.gif)

# Thoughts

While researching various React articles, I felt that there is a significant polarization between articles that only scratch the surface of React and those that dive deep into using React on the front lines.

This may sound presumptuous coming from someone with a lack of foundational knowledge in JavaScript, but I wish there were more articles focused on the process.

Even within React, some people use ES something or strict mode, and the build environments differ... I feel lost in the quagmire of front-end development (*_*).

Every time I write about React, I feel a sense of despair, but I enjoy writing about React. (｀･ω･´)ゞ

# References
* [Stack Overflow - Best practice for ReactJS form components](http://stackoverflow.com/questions/26626454/best-practice-for-reactjs-form-components) ... This might be the most referenced article for this implementation.
* [Stack Overflow - How do I edit multiple input controlled components in React?](http://stackoverflow.com/questions/35965275/how-do-i-edit-multiple-input-controlled-components-in-react)
* [If you're writing React with ES2015 or later, using onChange in form components to setState might be the way to go](http://bps-tomoya.hateblo.jp/entry/2016/05/25/154401)
* [Handling Forms in React.js](http://qiita.com/koba04/items/40cc217ab925ef651113)