
## 1. How to Contribute

1. <b>Clone from `develop` branch.</b>
Every new feature should clone from `develop` branch, this branch will deploy to staging server.
```shell
git clone -b develop git@github.com:chatbotgang/cantata.git
```

2. <b>Create your new branch and switch to it.</b>
```shell
git checkout -b <New Branch Name>
```

3. <b>Making some progress.</b>

4. <b>Add your contributions to staging area</b>
```shell
git add <Files You Updated>
```

5. <b>Commit to local git repository</b>
```shell
git commit -m "<type>[optional scope]: <description>"
```
Please check [commit message convention](https://www.conventionalcommits.org/en/v1.0.0/)

6. <b>Use `rebase` to merge `develop` branch</b>
```shell
git rebase develop
```
Please check [why use `rebase`?](https://www.conventionalcommits.org/en/v1.0.0/)

7. <b>Push to remote repository</b>
```shell
git push
```
You may not able to push commits because your branch has diverged from the published branch in the remote. If nobody is working on the same branch with you, you can use `git push -f` to push the rebased feature branch.[check the article](https://gitbook.tw/chapters/github/using-force-push)

8. <b>Launch a pull request, and assign at least one reviewer</b>
