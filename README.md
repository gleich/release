<!-- DO NOT REMOVE - contributor_list:data:start:["gleich"]:end -->

# release

📦 go package to check for a new GitHub release

![build](https://github.com/gleich/release/workflows/build/badge.svg)
![test](https://github.com/gleich/release/workflows/test/badge.svg)
![lint](https://github.com/gleich/release/workflows/lint/badge.svg)
![release](https://github.com/gleich/release/workflows/release/badge.svg)

## 🚀 Install

Simply run the following in the root of your project directory:

```txt
go get -u github.com/gleich/release
```

## 📄 Documentation [![GoDoc](https://godoc.org/github.com/gleich/release?status.svg)](https://godoc.org/github.com/gleich/release)

### `func Check`

```go
func Check(currentVersion string, repoURL string) (bool, string, error)
```

Check for an update. Takes in the current version and GitHub repo URL. Returns true or false if there is an update or not as well as the version value. Will return false if there is no network connection.

#### Example

```go
package main

import "github.com/gleich/release"

func main() {
    isOutdated, version, err := release.Check("v1.0.0", "https://github.com/gleich/nuke/")
}
```

## 🙌 Contributing

Before contributing please read the [CONTRIBUTING.md file](https://github.com/gleich/release/blob/master/CONTRIBUTING.md)

<!-- DO NOT REMOVE - contributor_list:start -->

## 👥 Contributors

- **[@gleich](https://github.com/gleich)**

<!-- DO NOT REMOVE - contributor_list:end -->
