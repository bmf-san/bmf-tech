---
title: Laravel5.3で自作artisanコマンド.md
description: Laravel5.3で自作artisanコマンド.md
slug: laravel-5-3-custom-artisan-command
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Laravel
  - リポジトリーパターン
  - artisan
translation_key: laravel-5-3-custom-artisan-command
---


今携わっているプロジェクトでRepositoryパターンを導入しているのですが、Repository関連のファイルを自動で生成するコマンドがあったら便利だなと思い、作ってみました。

# コマンドを生成
`php artisan make:command Repository`

/CommandsにRespository.phpというコマンド用のファイルが生成されます。


# コマンドファイルを編集
Repository.phpを編集します。
handleメソッド部分は[Creating file using Artisan Command in Laravel 5.1](http://stackoverflow.com/questions/32798132/creating-file-using-artisan-command-in-laravel-5-1)のコードをお借りして、少しカスタマイズしました。（偶然同じことをやろうとしている方がいたので・・）

```Repository.php
<?php

namespace App\Console\Commands;

use Illuminate\Console\Command;

class Repository extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'make:repository {modelName : The name of the model}';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Create respository files.';

    /**
     * Create a new command instance.
     *
     * @return void
     */
    public function __construct()
    {
        parent::__construct();
    }

    /**
     * Execute the console command.
     *
     * @return mixed
     */
    public function handle()
    {
        $modelName = $this->argument('modelName');

             if ($modelName === '' || is_null($modelName) || empty($modelName)) {
                 $this->error('Model name invalid..!');
             }

             if (! file_exists('app/Http/Repositories/Contracts') && ! file_exists('app/Http/Repositories/Eloquent')) {

                 mkdir('app/Http/Repositories/Contracts', 0775, true);
                 mkdir('app/Http/Repositories/Eloquent', 0775, true);

                 $contractFileName = 'app/Http/Repositories/Contracts/' . $modelName . 'RepositoryContract.php';
                 $eloquentFileName = 'app/Http/Repositories/Eloquent/' . $modelName . 'Repository.php';

                 if(! file_exists($contractFileName) && ! file_exists($eloquentFileName)) {
                     $contractFileContent = "<?php\n\nnamespace App\\Http\\Repositories\\Contracts;\n\ninterface " . $modelName . "RepositoryContract\n{\n}";

                     file_put_contents($contractFileName, $contractFileContent);

                     $eloquentFileContent = "<?php\n\nnamespace App\\Http\\Repositories\\Eloquent;\n\nuse App\\Repositories\\Contracts\\".$modelName."RepositoryContract;\n\nclass " . $modelName . "Repository implements " . $modelName . "RepositoryContract\n{\n}";

                     file_put_contents($eloquentFileName, $eloquentFileContent);

                     $this->info('Repository files created successfully.');

                 } else {
                     $this->error('Repository files already exists.');
                 }
             }
    }
}
```

# Kernel.phpにコマンドを登録
RepositoryコマンドをKernel.phpに登録します。

```Kernel.php
<?php

namespace App\Console\Commands;

use Illuminate\Console\Command;

class Repository extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'make:repository {modelName : The name of the model}';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Create repository files.';

    /**
     * Create a new command instance.
     *
     * @return void
     */
    public function __construct()
    {
        parent::__construct();
    }

    /**
     * Execute the console command.
     *
     * @return mixed
     */
    public function handle()
    {
        $modelName = $this->argument('modelName');
        $contractFileName = 'app/Repositories/Contracts/' . $modelName . 'RepositoryContract.php';
        $eloquentFileName = 'app/Repositories/Eloquent/' . $modelName . 'Repository.php';

        if ($modelName === '' || is_null($modelName) || empty($modelName)) {
            $this->error('Model name invalid..!');
        }

        if (! file_exists('app/Repositories/Contracts') && ! file_exists('app/Repositories/Eloquent')) {
            mkdir('app/Repositories/Contracts', 0775, true);
            mkdir('app/Repositories/Eloquent', 0775, true);

            $this->createFiles($modelName, $contractFileName, $eloquentFileName);
        } else {
            $this->createFiles($modelName, $contractFileName, $eloquentFileName);
        }
    }

    public function createFiles($modelName, $contractFileName, $eloquentFileName)
    {
        if(! file_exists($contractFileName) && ! file_exists($eloquentFileName)) {
            $contractFileContent = "<?php\n\nnamespace App\\Repositories\\Contracts;\n\ninterface " . $modelName . "RepositoryContract\n{\n}";

            file_put_contents($contractFileName, $contractFileContent);

            $eloquentFileContent = "<?php\n\nnamespace App\\Repositories\\Eloquent;\n\nuse App\\Repositories\\Contracts\\".$modelName."RepositoryContract;\n\nclass " . $modelName . "Repository implements " . $modelName . "RepositoryContract\n{\n}";

            file_put_contents($eloquentFileName, $eloquentFileContent);

            $this->info('Repository files created successfully.');
        } else {
            $this->error('Repository files already exists.');
        }
    }
}
```

# コマンドを実行してみる
`php artisan make:repository Hoge`

```
Repositories
├── Contracts
│   └── HogeRepositoryContract.php
└── Eloquent
    └── HogeRepository.php
```

こんな感じでファイルが生成されるかと思います。

#　所感
今回はRepositoryパターンの実装用のコマンドを作成しましたが、これを応用して色々なコマンドが作れそうですね。
今回つくったコマンドは、現在開発している[プロダクト](https://github.com/bmf-san/laravel-react-redux-blog-boilerplate)に導入しています。

