<!-- DO NOT REMOVE - contributor_list:data:start:["Matt-Gleich"]:end -->

# new_release

📦 go package to check for a new GitHub release

![build](https://github.com/Matt-Gleich/new_release/workflows/build/badge.svg)
![test](https://github.com/Matt-Gleich/new_release/workflows/test/badge.svg)
![lint](https://github.com/Matt-Gleich/new_release/workflows/lint/badge.svg)
![release](https://github.com/Matt-Gleich/new_release/workflows/release/badge.svg)

## 🚀 Install

Simply run the following in the root of your project directory:

```txt
go get -u github.com/Matt-Gleich/new_release
```

## 📄 Documentation [![GoDoc](https://godoc.org/github.com/Matt-Gleich/new_release?status.svg)](https://godoc.org/github.com/Matt-Gleich/new_release)

### `func Check`

```go
func Check(currentVersion string, repoURL string) (bool, string, error)
```

Check for an update. Takes in the current version and GitHub repo URL. Returns true or false if there is an update or not as well as the version value. Will return false if there is no network connection.

#### Example

```go
package main

import "github.com/Matt-Gleich/new_release"

func main() {
    isOutdated, version, err := new_release.Check("v1.0.0", "https://github.com/Matt-Gleich/nuke/")
}
```

## 🙌 Contributing

Before contributing please read the [CONTRIBUTING.md file](https://github.com/Matt-Gleich/new_release/blob/master/CONTRIBUTING.md)

<!-- DO NOT REMOVE - contributor_list:start -->

## 👥 Contributors

- **[@Matt-Gleich](https://github.com/Matt-Gleich)**

<!-- DO NOT REMOVE - contributor_list:end -->
