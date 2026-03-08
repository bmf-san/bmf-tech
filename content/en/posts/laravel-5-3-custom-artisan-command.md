---
title: Creating Custom Artisan Command in Laravel 5.3
slug: laravel-5-3-custom-artisan-command
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Laravel
  - Repository Pattern
  - artisan
translation_key: laravel-5-3-custom-artisan-command
---

Currently, I am working on a project that implements the Repository pattern, and I thought it would be convenient to have a command that automatically generates Repository-related files, so I created one.

# Generate Command
`php artisan make:command Repository`

A command file named Repository.php will be generated in the /Commands directory.

# Edit Command File
Edit the Repository.php file. I borrowed the code from [Creating file using Artisan Command in Laravel 5.1](http://stackoverflow.com/questions/32798132/creating-file-using-artisan-command-in-laravel-5-1) for the handle method and made some customizations. (Coincidentally, someone else was trying to do the same thing...)

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

                     $eloquentFileContent = "<?php\n\nnamespace App\\Http\\Repositories\\Eloquent;\n\nuse App\\Repositories\\Contracts\".$modelName."RepositoryContract;\n\nclass " . $modelName . "Repository implements " . $modelName . "RepositoryContract\n{\n}";

                     file_put_contents($eloquentFileName, $eloquentFileContent);

                     $this->info('Repository files created successfully.');

                 } else {
                     $this->error('Repository files already exists.');
                 }
             }
    }
}
```

# Register Command in Kernel.php
Register the Repository command in Kernel.php.

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

            $eloquentFileContent = "<?php\n\nnamespace App\\Repositories\\Eloquent;\n\nuse App\\Repositories\\Contracts\".$modelName."RepositoryContract;\n\nclass " . $modelName . "Repository implements " . $modelName . "RepositoryContract\n{\n}";

            file_put_contents($eloquentFileName, $eloquentFileContent);

            $this->info('Repository files created successfully.');
        } else {
            $this->error('Repository files already exists.');
        }
    }
}
```

# Execute the Command
`php artisan make:repository Hoge`

```
Repositories
├── Contracts
│   └── HogeRepositoryContract.php
└── Eloquent
    └── HogeRepository.php
```

I believe files will be generated like this.

# Thoughts
This time, I created a command for implementing the Repository pattern, but it seems like I could create various commands by applying this. The command I created has been integrated into the [product](https://github.com/bmf-san/laravel-react-redux-blog-boilerplate) I am currently developing.