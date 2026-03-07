---
title: Considerations for React+Redux Directory Structure
slug: react-redux-directory-structure
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - ES6
  - React
  - Redux
translation_key: react-redux-directory-structure
---

When trying to introduce React+Redux into an application, I was struggling with the directory structure, but I came across a good article and decided to summarize it.

[A Better File Structure For React/Redux Applications](http://marmelab.com/blog/2015/12/17/react-directory-structure.html)

# Pattern Following Redux Structure
A simple and straightforward pattern. This is probably the standard for SPA development.

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

# Basic Structure with Domain
A pattern that comes to mind when there are multiple domains. It looks neat, but I have a feeling it might get cumbersome with compilation. Is this okay for SPA?? It might be good to group by domain within each directory.

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

# Pattern Divided by Domain
A pattern where the top directory is divided by domain, managing actions, containers, reducers, etc., with suffixes. This pattern seems familiar when introducing it into MVC server-side frameworks.

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

# Thoughts
I have seen various articles and repositories, but it seems to vary depending on the environment...