# random_account
A commandline tool which generate a random account username and password


#### 1. Usage
```text
USAGE:
    random_account [FLAGS] [OPTIONS]

FLAGS:
    -h, --help       Prints help information
        --save       save account data to file
    -V, --version    Prints version information

OPTIONS:
    -l, --length <length>                      set account name length [default: 24]
    -L, --password_length <password_length>    set account password length [default: 32]
    -p, --prefix <prefix>                      set account username prefix [default: ]
    -s, --site <site>                          set account site url [default: ]
```

#### 2. examples
- create an account for logging www.example.com
```bash
random_account --site www.example.com
```
```text
---
- site: www.example.com
  username: skhEzdMvPaDf0ASyyyjPOxmn
  password: tICZkFJ5DSS9NZI9wI2PIodhNMAtOJKH
  time: "2021-12-08T10:54:19.611+08:00"
```

- create an account for logging www.example.com with prefix
```bash
random_account --site www.example.com --prefix example-
```
```text
- site: www.example.com
  username: example-qJZoL6bQxWeOwJvS
  password: y4Nu6cE9sSgSFoLllbDyGE1SP5Q1i3Uf
  time: 2021-12-08T10:55:09.311+08:00
```

- - create an account for logging www.example.com with user-defined account name
```bash
random_account --site www.example.com --prefix example --length 0
```
```text
- site: www.example.com
  username: example
  password: TGUPlSfUM2RSPTUlaSHurwkLWqQda81E
  time: 2021-12-08T10:55:36.876+08:00
```
