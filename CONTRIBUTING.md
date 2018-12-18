# How to contribute
We welcome contributions from the community and are pleased to have them. <br/>
Please follow this guide when logging issues or making code changes.

## Logging Issues
All issues should be created using the [new GitHub Issue form](https://github.com/siesgstarena/epicentre/issues/new). <br/>
Use our [Issue Template](.github/ISSUE_TEMPLATE.md). <br/>
Clearly describe the issue including steps to reproduce if there are any. Also, make sure to indicate the earliest version that has the issue being reported.

## Patching Code
Code changes are welcome and should follow the guidelines below.

* Fork the repository on GitHub.
* Fix the issue ensuring that your code follows this project's [tslint configuration](tslint.json).
    * Run `npm run tslint` to generate a report of lint issues
    * Run `tslint --fix --project .` to auto fix most of the lint issues
* Add tests for your new code ensuring that you have 100% code coverage (we can help you reach 100% but will not merge without it).
    * Run `npm run test` to generate a report of test coverage
* Use our [Pull Request Template](.github/PULL_REQUEST_TEMPLATE.md) for making a new pull request.
* [Pull requests](http://help.github.com/send-pull-requests/) should be made to the [features branch](https://github.com/siesgstarena/epicentre/tree/features).