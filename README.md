<div align="center">
<img src="https://capsule-render.vercel.app/api?type=waving&color=0:101827,100:0f766e&height=170&section=header&text=Go%20Build%20Constraints%20Lab&fontSize=32&fontColor=ffffff&animation=fadeIn" width="100%" alt="Waving header"/>
<img src="https://readme-typing-svg.demolab.com?font=Fira+Code&weight=600&size=18&pause=900&color=14B8A6&center=true&vCenter=true&width=760&lines=A%20Go%20CLI%20that%20previews%20which%20local%20.go%20files%20match%20a%20selected%20GOOS%2C%20GOARCH%2C" alt="Typing title"/>

[![License](https://img.shields.io/github/license/CyberLoop-dev/go-build-constraints-lab)](LICENSE)
![Language](https://img.shields.io/github/languages/top/CyberLoop-dev/go-build-constraints-lab)
![Stars](https://img.shields.io/github/stars/CyberLoop-dev/go-build-constraints-lab?style=flat)
![Last commit](https://img.shields.io/github/last-commit/CyberLoop-dev/go-build-constraints-lab)
<img src="assets/mark.svg" width="112" alt="Go Build Constraints Lab illustration"/>
</div>

## 🧭 Overview

A Go CLI that previews which local .go files match a selected GOOS, GOARCH, and build-tag set without compiling or modifying the project.

## ✨ Features

- 🧭 Evaluate //go:build expressions and filename suffixes.
- 🖥️ Choose target GOOS and GOARCH values.
- 🏷️ Supply additional build tags.
- 🔒 Read source directories only; never execute files.

## 🎬 Demo

No hosted demo is provided. Run the local examples with the command below.

## ⚡ Quick Start

~~~bash
go run . -dir ./examples -goos windows -goarch amd64
~~~

## 🧰 Tech Stack

- Go 1.22+
- Standard library; no runtime service or external API is required

## 🗂️ Folder Structure

~~~text
go-build-constraints-lab/
├── main.go
├── main_test.go
├── examples/feature_windows.go
├── examples/feature_unix.go
├── examples/feature_default.go
├── go.mod
├── assets/mark.svg
├── .gitignore
├── LICENSE
└── README.md
~~~

## 🗺️ Roadmap

- [x] Core behavior and local examples
- [x] Custom SVG and setup documentation
- [ ] Add more guided exercises from community feedback

## 🤝 Contributing

Issues and focused pull requests are welcome. Keep examples deterministic and within the documented scope.

## 📄 License

MIT License. See [LICENSE](LICENSE).

<div align="center"><img src="https://capsule-render.vercel.app/api?type=waving&color=0:0f766e,100:101827&height=100&section=footer" width="100%" alt="Waving footer"/></div>
