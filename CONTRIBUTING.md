
# 1. How to Contribute


## Developing 

#### 1. <b>Clone from `develop` branch.</b>
Every new feature should clone from `develop` branch, this branch will deploy to `staging` env.
```shell
git clone -b develop git@github.com:chatbotgang/cantata.git
```

#### 2. <b>Create your new branch and switch to it.</b>
```shell
git checkout -b <New Branch Name>
```

#### 3. <b>Making some progress.</b>
See [Feature Development Guidelines](./guideline_feature_developement.md)

#### 4. <b>Add your updates to staging area</b>
```shell
git add <Files You Updated>
```

#### 5. <b>Commit to local git repository</b>
```shell
git commit -m "<type>[optional scope]: <description>"
```
Please check [commit message convention](https://www.conventionalcommits.org/en/v1.0.0/)

#### 6. <b>Use `rebase` to merge `develop` branch</b>
```shell
git rebase develop
```
Please check [why use `rebase`?](https://www.conventionalcommits.org/en/v1.0.0/)

#### 7. <b>Push to remote repository</b>
```shell
git push
```
You may not able to push commits because your branch has diverged from the published branch in the remote. If nobody is working on the same branch with you, you can use `git push -f` to push the rebased feature branch.[check the article](https://gitbook.tw/chapters/github/using-force-push)

#### 8. <b>Open a PR, and assign at least one reviewer</b>
In this stage, CI will automatically trigger some check, include unit tests.......


## Deploy to `staging`

#### 1. <b>After reviewers approve your PR, merge your update to `develop` branch with `rebase merge`</b>
In this stage, CD will automatically deploy your updates to `staging` env.

#### 2. <b>Check if your updates works fine in `staging` env</b>


## Deploy to `production`

#### 1. <b>Open a release PR from `develop` into `main` branch.</b>
The release PR page will show the diff commits for this PR, select all commit's authors as reviewer. This PR can't be merged untill all reviewers approved.
 
#### 2. <b>Annouce release time in Slack channel `team-eng-backend`.</b>

#### 3. <b>At launch time, if all reviewers already approved, merge this PR with `merge commit`.</b>
In this stage, CD will automatically deploy all updates to `production` env.
