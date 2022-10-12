
## 1. How to Contribute

1. Clone from `develop` branch
```shell
git clone -b develop git@github.com:chatbotgang/cantata.git
```

2. Create your new branch and switch to it
```shell
git checkout -b <new branch name>
```

3. making some progress

4. Add your contributions to staging area
```shell
git add <file you updated>
```
5. Commit to local git repository
```shell
git commit -m "<type>[optional scope]: <description>"
```
please check the [commit message convention](https://www.conventionalcommits.org/en/v1.0.0/)

6. Use `rebase` to merge branch
```shell
git rebase develop
```
[why use `rebase`?](https://www.conventionalcommits.org/en/v1.0.0/)

7. Push to remote repository
```shell
git push
```
