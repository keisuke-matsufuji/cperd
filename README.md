# GitHub Action: PHP error-reporting directive visualize with reviewdog 🐶

This action makes visible the PHP error-reporting directives in the php.ini file and in .htaccess as predefined PHP error levels.

In repositories that have this Actions in place, if PHP error-reporting directives are included in the php.ini file or .htaccess when creating a PullRequest, a valid PHP error level will be set in the Pull Request's Add to the review comments.

## Example usage
```yaml
name: CI
on:
  - pull_request
jobs:
  steep:
    name: php error-reporting-directive visualize
    runs-on: ubuntu-latest
    steps:
      - name: Check out code
        uses: actions/checkout@v3
      - name: check
        uses: keisuke-matsufuji/php-error-reporting-directive-visualize@main
        env:
          # A token with repo scope is OK.
          # In GitHub Actions, there is a default Secret called `GITHUB_TOKEN`
          # You can use this if you like
          REVIEWDOG_GITHUB_API_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

## Pull Request Comment Example

<img width="925" alt="image" src="https://github.com/keisuke-matsufuji/php-error-reporting-directive-visualize/assets/49528505/339af229-c5e5-45f6-9f82-b3201c5278a7">

## For Debug
```
$ make docker/build
$ make docker/run
# bash cperd/entrypoint.sh
```

## License
[MIT](https://choosealicense.com/licenses/mit/)
