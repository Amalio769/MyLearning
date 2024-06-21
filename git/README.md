# MyGitLearning
My Git Learning

* Upload a repository from laptop to github
First you create the repo in github --> <repo-name>
```bash
cd <dir-repo>
git init
git remote add origin https://github.com/<github-user-name>/<repo-name>.git
git add .
git commit -m "Initial commit"
git push --set-upstream origin master
```
* Create a `tag` for the repo
```bash
git tag "v1.0.0"
git push --tags
```