# Contributing to ShardingSphere Operator

Thanks for your interest in contributing to ShardingSphere Operator! This document outlines some of the conventions on building, running, and testing ShardingSphere Operator, the development workflow, commit message formatting, contact points and other resources.



# First ShardingSphere Operator Pull Request

## Prerequisites

To build ShardingSphere Operator from scratch you will need to install the following tools:

* Git
* [Golang 1.17](https://golang.org/dl/)
* make
* Kubernetes 1.20+ 

## Pull Requests

### Submit a PR
1. Fork the ShardingSphere-Operator repo and set remote repo.
      ```
      # git clone https://github.com/<yourname>/shardingsphere-on-cloud.git
      # cd shardingsphere-on-cloud

      # git remote add upstream https://github.com/SphereEx/shardingsphere-on-cloud.git

      # git remote -v
      origin	https://github.com/<your name>/shardingsphere-on-cloud.git (fetch)
      origin	https://github.com/<your name>/shardingsphere-on-cloud.git (push)
      upstream	https://github.com/database-mesh/shardingsphere-on-cloud.git (fetch)
      upstream	https://github.com/database-mesh/shardingsphere-on-cloud.git (push)
      ```
2. Open a regular issue for binding the pull request.
3. Submit a Draft Pull Requests, tag your work in progress.
4. Create your own branch and develop with it.Before developing, it's recommend to pull from remote repo to keep your repo latest. Now,you could develop at new branch.
      ```
      git checkout master
      git fetch upstream
      git rebase upstream/master
      git checkout -b futures-0.1.0-dev
      ```
5. If you have added code that should be tested, add unit tests.
6. Verify and ensure that the test suites passes, make test.
7. Make sure your code passes both linters, make lint.
8. Submit and push your code to the remote repo.
      ```
      git add <file>
      git commit -m 'commit log'
      git push origin futures-0.1.0-dev
      ```
9.  Change the status to “Ready for review”.

### PR Template

```
<!-- Please answer these questions before submitting a pull request -->

### Type of change:

<!-- Please delete options that are not relevant. -->

- [ ] Bugfix
- [ ] New feature provided
- [ ] Improve performance
- [ ] Backport patches

### What this PR does / why we need it:
<!--- Why is this change required? What problem does it solve? -->
<!--- If it fixes an open issue, please link to the issue here. -->

### Pre-submission checklist:

<!--
Please follow the requirements:
1. Test is required for the feat/fix PR, unless you have a good reason
2. Doc is required for the feat PR
3. Use "request review" to notify the reviewer once you have resolved the review
-->

* [ ] Did you explain what problem does this PR solve? Or what new features have been added?
* [ ] Have you added corresponding test cases?
* [ ] Have you modified the corresponding document?


```

### PR Commit Message

Format: `<type>(<scope>): <subject>`

`<scope>` is optional

```
fix(functions): fix group by string bug
^--^  ^------------^
|     |
|     +-> Summary in present tense.
|
+-------> Type: chore, docs, feat, fix, refactor, style, or test.
```

More types:

* `feat`: new feature for the user.
* `fix`: bug fix for the user.
* `docs`: changes to the documentation.
* `style`: formatting, missing semi colons, etc; no production code change.
* `refactor`: refactoring production code, eg. renaming a variable.
* `test`: adding missing tests, refactoring tests; no production code change.
* `chore`: updating grunt tasks etc; no production code change.

## Issues
ShardingSphere Operator uses [GitHub issues](https://github.com/SphereEx/shardingsphere-on-cloud/issues) to track bugs. Please include necessary information and instructions to reproduce your issue.

## Code of Conduct
Please refer to the [CNCF Code of Conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md), which describes the expectations for interactions within the community. 
