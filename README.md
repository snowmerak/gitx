# gitx

gitx is a simple command line tool to help you manage your git repositories.

## Installation

```bash
go install github.com/snowmerak/gitx@latest
```

## Usage

```bash
gitx [command]
```

## Commands

gitx supports the following commands:

### config

#### Initialize configuration

```bash
gitx config init [<ssh-key-name>]
```

The ssh key pair is in the `./.gitx`.  
The file names are `./.gitx/[<ssh-key-name>].prv.pem` and `./.gitx/[<ssh-key-name>].pub.pem`.

#### Initialize git ignore configuration

```bash
gitx config ignore init
```

This command will create a `.gitignore` file in the current directory.  
And add some rules to it.

### ssh

#### Generate SSH key

```bash
gitx ssh generate <name>
```

The ssh key pair will be generated in the `./.gitx/[<ssh-key-name>].prv.pem` and `./.gitx/[<ssh-key-name>].pub.pem`.

### fork

#### Fork a feature branch

```bash
gitx fork feature <name>
```

Switch to the `feature/<name>` branch.

#### Fork a proposal branch

```bash
gitx fork proposal <name>
```

Switch to the `proposal/<name>` branch.

#### Fork a hotfix branch

```bash
gitx fork hotfix <name>
```

Switch to the `hotfix/<name>` branch.

#### Fork a bugfix branch

```bash
gitx fork bugfix <name>
```

Switch to the `bugfix/<name>` branch.

#### Fork a daily branch

```bash
gitx fork daily <name>
```

Switch to the `daily/<name>` branch.

#### Revert a branch

```bash
gitx fork revert
```

Revert to the previous branch.

### Pull

```bash
gitx pull
```

Pull changes from remote.

### Push

```bash
gitx push <message>
```

Push changes to remote.

### Changes

```bash
gitx changes
```

Show changes

Example:

```bash
11:39PM INF changes
Created files:
  pages/contents.md

Changed files:
  .DS_Store
  journals/2024_01_26.md
```
