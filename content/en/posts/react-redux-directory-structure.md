---
title: Considering Directory Structure for React+Redux
slug: react-redux-directory-structure
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - ES6
  - React
  - Redux
description: Exploring directory structures for integrating React+Redux into applications.
translation_key: react-redux-directory-structure
---



When trying to introduce React+Redux into an application, I was pondering over the directory structure. I came across a promising article and decided to summarize it.

[A Better File Structure For React/Redux Applications](http://marmelab.com/blog/2015/12/17/react-directory-structure.html)

# Pattern Following Redux Structure
This is a simple and straightforward pattern. It might be the standard for SPA development.

```
actions/
    CommandActions.js
    UserActions.js
components/
    Header.js
    Sidebar.js
    Command.js
    CommandList.js
    CommandItem.js
    CommandHelper.js
    User.js
    UserProfile.js
    UserAvatar.js
containers/
    App.js
    Command.js
    User.js
reducers/
    index.js
    command.js
    user.js
routes.js
```

# Pattern with Domain in Basic Structure
This is a pattern that might come to mind first when there are multiple domains. It's neat, but it might become cumbersome to compile. Is this okay for SPA?
You could also group by domain within each directory.

```
actions/
    CommandActions.js
    ProductActions.js  
    UserActions.js
components/
    Header.js
    Sidebar.js
    Command.js
    CommandList.js
    CommandItem.js
    CommandHelper.js
    Product.js        
    ProductList.js     
    ProductItem.js     
    ProductImage.js    
    User.js
    UserProfile.js
    UserAvatar.js
containers/
    App.js
    Command.js
    Product.js          
    User.js
reducers/
    index.js
    foo.js
    bar.js
    product.js         
routes.js
```

# Pattern of Dividing by Domain
This pattern involves dividing the top directory by domain and managing actions, containers, reducers, etc., with suffixes. This pattern might be familiar when introducing it to server-side MVC frameworks.

```
app/
    Header.js
    Sidebar.js
    App.js
    reducers.js
    routes.js
command/
    Command.js
    CommandContainer.js
    CommandActions.js
    CommandList.js
    CommandItem.js
    CommandHelper.js
    commandReducer.js
product/
    Product.js
    ProductContainer.js
    ProductActions.js
    ProductList.js
    ProductItem.js
    ProductImage.js
    productReducer.js
user/
    User.js
    UserContainer.js
    UserActions.js
    UserProfile.js
    UserAvatar.js
    userReducer.js
```

# Things to Consider for Directory Structure
* Is import cumbersome?
* Handling of containers and components (e.g., grouping them together)
* Testing
* Others...

# Impressions
After reviewing various articles and repositories, it seems to vary depending on the environment...
