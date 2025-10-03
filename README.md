<h1 align="center">umm</h1>
<p align="center" style="font-style: italic;">umm, how do I do that?</p>

> [!NOTE]
> Work in progress.

## 📖 Description

A command-line AI assistant that converts natural language questions into executable commands with explanations. Ask what you need in plain English and get practical solutions.

## ✨ Features

- Ask questions in plain English
- Built-in safety checks for dangerous commands
- Build on previous queries with context
- Search, view, and manage past interactions
- Choose between GPT-4o, GPT-4o-mini, and GPT-4

## 🚀 Quick Start

```bash
# Ask a question
umm "list all files in tree format"

# Follow up on your last query
umm + "but only show directories"

# Run the suggested command
umm --run

# View your history
umm history
```

## 🛠️ Development

Requires Go 1.24+

```bash
make build
make install
```

Configure your OpenAI API key:

```bash
umm config setup
```
