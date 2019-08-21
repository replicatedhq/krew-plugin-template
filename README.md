# kubectl krew template repo

There's a lot of scaffolding needed to set up a good kubectl plugin. This repo is a GitHub Template Repo to make it easy to set all of this scaffolding up for a new repo.

The assumptions made are:
1. You'll write your plugin in go
2. You want client-go to interact with the cluster
3. You want all of the kubectl pflags available to your plugin
4. You will store your code on GitHub.com

## Create your repo

On the top of this [repo](https://github.com/replicatedhq/krew-plugin-template), there's a button to create a new repo from this template. This is not a fork, it will make a copy of this repo into your own organization or GitHub account. 

Click that, and create your own version of this repo. Clone it locally. The rest of the steps you will be performing on your local copy.

## Make it yours

Once you have your own repo created locally, change to the directory and run:

```shell
make setup
```

This will prompt you for a few things, such as your GitHub org, repo name and plugin name. The setup application will then update the import paths and code with the data you provided.

Commit and check it in to your repo!

```shell
git add .
git commit -m "Updating from template"
git push -u origin master
```

## Write your Plugin

Next, open the pkg/plugin/plugin.go file. This is where you can start writing your plugin.

For an example, take a look at the [outdated](https://github.com/replicatedhq/outdated) plugin that inspired this template.
