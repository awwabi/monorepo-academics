# ACADEMICS APP - Monorepo
This application was created with the aim of being a thesis material, the code in this repository has been modified from the original application used in myITS Academics.

## Installation
1. Clone this repository
2. Run `git clone git@bitbucket.org:dptsi/go-auth-module.git auth` in the root directory
3. Add .env file in each module
4. Run `go mod tidy` in the root directory
5. Run each module with `go run main.go` in the module directory
6. Enable go workspace by running `go work int` in the root directory
7. Add each module to the workspace by running `go work add <module-name>` in the root directory
