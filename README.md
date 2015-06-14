# go-get-git-ignore

Fetch and save .gitignore file from github/gitignore quickly.

## Usage

```sh
$ get-git-ignore --lang=Go
$ ls -a
.gitignore
```

* `--lang`: Language for .gitignore.
* `--out`: Path to save .gitignore file. Default is `.gitignore`.

## Installation

```sh
$ go get github.com/handlename/go-get-git-ignore/cmd/get-git-ignore
```

## Example

Using with [peco](https://github.com/peco/peco),
you can filter candidates interactively.
Install peco and put like below to your `.zshrc` file.

```sh
function _peco_ggi_list () {
    lang=$(get-git-ignore | peco)

    if [ -n "$lang" ]; then
        get-git-ignore --lang=$lang
    fi
}
alias ggi=_peco_ggi_list
```

## Licence

[MIT](https://github.com/handlename/go-get-git-ignore/blob/master/LICENCE)

## Author

[handlename](https://github.com/handlename)
