---
title: "React+Reduxのディレクトリ構成検討"
slug: "react-plus-redux"
date: 2017-10-01
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "ES6"
  - "React"
  - "Redux"
draft: false
---

React+Reduxをアプリケーションに導入しようとするときに、そのディレクトリ構成について悩んでいたのですが、良さ気な記事を見かけたのでまとめてみました。

[A Better File Structure For React/Redux Applications](http://marmelab.com/blog/2015/12/17/react-directory-structure.html)

# Reduxの構成に従うパターン
シンプルかつベタなパターンです。SPA開発ならこれが定番でしょうか。

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

# 基本構成にドメインが入ってきたパターン
ドメインが複数ある時に、真っ先に思い浮かびそうなパターン。
スッキリしていますが、コンパイルとか面倒くさくなりそうな予感。SPAならこれでもOK？？
各ディレクトリ内でドメインでグルーピングしてディレクトリきっても良さそう。

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

# せや、ドメインで分割したろパターン
トップのディレクトリをドメインできって、action,container,reducerやらを接尾辞で管理していくパターン。
MVCのサーバーサイドフレームワークに導入してする際は、このパターンが馴染みそう。

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

# ディレクトリ構成で検討したほうがいいかもしれないこと
* importは煩わしくないか？
* containerとcomponentの扱い（一纏めにするなど）
* テスト
* その他・・・

# 所感
色々な記事やリポジトリを拝見しましたが、環境によってバラバラなようです・・・

