# SmartCommit

SmartCommit is a blazing-fast CLI tool that uses AI to generate meaningful Git commit messages from your staged changes — powered by local or remote LLMs (Ollama, OpenAI, Gemini or any HTTP-based LLM).

---

## 🚀 Features

* 🔍 Reads your staged Git diff
* 🤖 Supports multiple LLM providers: Ollama, OpenAI, HTTP endpoints (Claude, Gemini, etc.)
* ✏️ Customizable system prompt (edit in Vim)
* ⚡ Interactive flow with options to commit, edit, regenerate, or quit
* ⚡ One‑shot `--yes` flag for auto‑commit
* 📦 Single static binaries for Linux, macOS (Intel/ARM), and Windows

---

## 📥 Quick Installation

### 1. Prebuilt Binary (Recommended)
(will soon add direct installs from brew/scoops soon)

1. Go to [Releases](https://github.com/manyfacedqod/smartcommit/releases).
2. Download the `.tar.gz` for your OS.
3. Extract and move into your `PATH`:

   * **Linux**

     ```bash
     tar -xzf smartcommit-linux.tar.gz
     sudo mv smartcommit-linux /usr/local/bin/smartcommit
     chmod +x /usr/local/bin/smartcommit
     ```
   * **macOS**

     ```bash
     tar -xzf smartcommit-macos.tar.gz
     sudo mv smartcommit-macos /usr/local/bin/smartcommit
     chmod +x /usr/local/bin/smartcommit
     ```
   * **Windows (Git Bash)**

     ```bash
     tar -xzf smartcommit-windows.tar.gz
     mv smartcommit-windows.exe /usr/local/bin/smartcommit.exe
     ```

### 2. Go Install

```bash
go install github.com/manyfacedqod/smartcommit@v1.1.1
```

Ensure `$GOPATH/bin` or `~/go/bin` is in your `PATH`.

### 3. Build from Source

```bash
git clone https://github.com/manyfacedqod/smartcommit
cd smartcommit
go build -o smartcommit
sudo mv smartcommit /usr/local/bin/
```

---

## ⚙️ First‑Time Setup

```bash
smartcommit setup
```

You’ll be prompted to configure:

1. **Provider**: `ollama`, `openai`, or `http`
2. **Model name**: e.g., `llama3`, `gpt-4`, `gemini`
3. **API Key** (HTTP/OpenAI) or **Base URL** (HTTP/Ollama)
4. **System prompt** is saved to `~/.config/smartcommit/config.yaml`

---

## 💡 Usage Examples

### Interactive Commit Flow

```bash
git add .
smartcommit generate
```

```
💡 Generated Commit Message:
fix: handle empty username in login flow

Choose [c]ommit, [e]dit, [r]egenerate, [q]uit:
```

* **c**: commit
* **e**: inline edit via PromptUI
* **r**: regenerate
* **q**: quit

### Auto‑Commit Without Prompt

```bash
smartcommit generate --yes
```

### Config Commands

* **Show current config**:

  ```bash
  smartcommit config show
  ```
* **Edit system prompt (Vim)**:

  ```bash
  smartcommit config edit
  ```

---

## 📄 Command Reference

| Command                      | Description                                 |
| ---------------------------- | ------------------------------------------- |
| `smartcommit setup`          | Configure your LLM provider and settings    |
| `smartcommit generate`       | Generate commit message from staged diff    |
| `smartcommit generate --yes` | Auto‑generate & commit without prompt       |
| `smartcommit config show`    | Display current configuration               |
| `smartcommit config edit`    | Edit the system prompt using your `$EDITOR` |
| `smartcommit --version`      | Show version (v1.1.1)                       |

---

## 👨‍💻 Open Source Contribution

We welcome contributions! Here’s how to get started:

1. **Fork** the repo and **clone** your fork:

   ```bash
   git clone https://github.com/<your-user>/smartcommit
   cd smartcommit
   ```
2. **Create** a feature branch:

   ```bash
   git checkout -b feat/your-feature
   ```
3. **Implement** your changes, **add tests**, and **update docs** as needed.
4. **Run** existing tests & build:

   ```bash
   ```

go test ./...
go build -o smartcommit

```
5. **Commit** and **push**, then open a **Pull Request** against `main`.

### Contribution Guidelines

- Follow existing code style.
- Write clear commit messages (`feat:`, `fix:`, `docs:`, etc.).
- Include tests for new features.
- Document changes in this `README.md` or code comments.

Thank you for making SmartCommit better! 🙏

---

## 📄 License

MIT License © manyfacedqod

```
