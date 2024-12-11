# al

🎭 Create local aliases, like package.json but everywhere

Version : 0.0.1

## 🚀 Installation

### Requirements

### Via Homebrew

Wip 🚧

### Manual

You can paste the binary in your `bin` directory (e.g., on mac it's `/usr/bin/local`). \
Don't forget to grant execution permissions to the binary.

```bash
chmox +x al
```

## 💻 Usage

`al` needs a `al.json` file where you define your custom aliases, generally at the root of your project. Since the config file exists, you can use `al` followed by your local alias.

```bash
# al.json
{
    "test": "echo "This is a test""
}
```

```bash
# you can now use al to run your command
al test
```

You can use `--init` / `-i` to create a new config file.

```bash
al --init
```

## 🧑‍🤝‍🧑 Contributing

To contribute, fork the repository and open a pull request detailling your changes.

Create a branch with a [conventionnal name](https://tilburgsciencehub.com/building-blocks/collaborate-and-share-your-work/use-github/naming-git-branches/).

- fix: `bugfix/the-bug-fixed`
- features: `feature/the-amazing-feature`
- test: `test/the-famous-test`
- hotfix `hotfix/oh-my-god-bro`
- wip `wip/the-work-name-in-progress`

## 📌 Roadmap

- [ ] Create a new option to add a new alias from command
- [ ] Improve UI
- [ ] Improve user feedback when using an alias
- [ ] Installation via Homebrew
- [ ] Installation via apt

## 📑 License

This project is under MIT license. For more information, please see the file [LICENSE](./LICENSE).
