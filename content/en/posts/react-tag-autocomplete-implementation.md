---
title: Implementing a Tag Feature with Suggestions Using React Tag Autocomplete
description: 'Implement tag autocomplete feature in React using React Tag Autocomplete library with API integration.'
slug: react-tag-autocomplete-implementation
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - React
translation_key: react-tag-autocomplete-implementation
---

I was thinking of developing a tag feature from scratch, but I found many convenient React Components, so I decided to use them.

# Environment
* React
* Babel
* Browserify
* [React Tag Autocomplete](https://github.com/i-like-robots/react-tags)
* npm

# Introduction
Install React Tag Autocomplete using npm.

`npm install --save react-tag-autocomplete`

There are various ways to include it, but in this environment, we will use require.

```hoge.js
var ReactTags = require('react-tag-autocomplete');
```

Now we are ready.

# Implementation

```hoge.html
// Various parts omitted
<div id="react-tag-autocomplete"></div>
```

There is a Usage section on [github](https://github.com/i-like-robots/react-tags), but let's consider a case where we modify it a bit to fetch data by hitting an API. (Here, we use superagent.)

The API returns a JSON response like this:

```
[{"id":1,"name":"Programming"},{"id":2,"name":"Housework"},{"id":3,"name":"Home Security"},{"id":4,"name":"Early to Bed, Early to Rise"},{"id":5,"name":"Three-Day Monk"}]
```

Debugging res.body.skills
![Screenshot 2016-09-28 3.04.10.png](/assets/images/posts/react-tag-autocomplete-implementation/819ec751-a2d9-df73-d353-76847710c4b7.png)

```hoge.js
var ReactTags = require('react-tag-autocomplete');

var App = React.createClass({
  getInitialState: function () {
    return {
      tags: [],
      suggestions: []
    }
  },

  componentDidMount: function () {
    request
      .get('/api/v1/user/config')
      .end(function(err, res){
        if (err) {
          alert('Communication error. Please reload.');
        }
        this.setState({
          suggestions: res.body.skills
        });
      }.bind(this));
  },

  handleDelete: function (i) {
    var tags = this.state.tags
    tags.splice(i, 1)
    this.setState({ tags: tags })
  },![tags.gif](/assets/images/posts/react-tag-autocomplete-implementation/173c6de9-b87a-6200-65ed-506e181f565e.gif)
![tags.gif](/assets/images/posts/react-tag-autocomplete-implementation/a3372702-2a85-9b80-0b53-ede2c9c3c486.gif)


  handleAddition: function (tag) {
    var tags = this.state.tags
    tags.push(tag)
    this.setState({ tags: tags })
  },

  render: function () {
    return (
      <ReactTags
        tags={this.state.tags}
        suggestions={this.state.suggestions}
        handleDelete={this.handleDelete}
        handleAddition={this.handleAddition} />
    )
  }
})

ReactDOM.render(
  <App />, 
  document.getElementById('react-tag-autocomplete')
);
```

Operation check (there are some unnecessary things shown...)
![tags.gif](/assets/images/posts/react-tag-autocomplete-implementation/c0c75e77-d212-4cd4-5e1b-c59a8a412c61.gif)

The CSS is not set, so it looks quite plain lol

Other options and CSS class names are clearly stated on [github](https://github.com/i-like-robots/react-tags).

# Impressions
It's such a convenient era.ヽ(´ー｀)ノ