# Importing other packages in the same module
Credit:
https://stackoverflow.com/questions/55442878/organize-local-code-in-packages-using-go-modules


When importing another package with modules, you always use the full path including the module path. This is true even when importing another package in the same module. For example, if a module declared its identity in its go.mod as module `github.com/my/repo`, and you had this organization:

```
repo/
├── go.mod      <<<<< Note go.mod is located in repo root
├── pkg1
│   └── pkg1.go
└── pkg2
    └── pkg2.go
```
Then `pkg1` would import its peer package as import `"github.com/my/repo/pkg2"`. Note that you cannot use relative import paths like import `"../pkg2"` or import `"./subpkg"`. (This is part of what OP hit above with import `"./stuff"`).