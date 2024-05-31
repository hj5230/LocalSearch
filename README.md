# LocalSearch

## Introduction

A lightweight, CLI based application for finding matching strings in speccified directory

## Develop Note

### Bugs to be fixed

- Incorrect recognition of Chinese character length causes highlight mark misalignment
	- utf8 3 -> 2
	- utf8mb4 4 -> 2
	- ...
- Heuristic algorithm unable to search some sort of text file, e.g. nohup.out

### Features to be added

- Help page with command `help` or flag `-h`
- Search current directory if directory `-d` is not specified
- Use `config.yaml` replace constant define
- Ignore file or directory with flag `-i`
- Hide excessive results
	- Initially displays `and ... more matched results`
	- switch to `cat` styled after pressing arrow down key
	- Allow using arrow keys to switch pages, i.e. `cat` styled
- Allows printing of all results at once with flag `-a` (save to file with `>`)
