# gitai

gitai is a simple CLI tool that takes a git diff and generates either a code review or a commit message using a local LLM (in this case, Ollama).

I built this project while working as an apprentice at Arribatec.
The main reason was that I often find it hard to come up with good commit titles — and sometimes I’m not even sure if my code is good or bad.
So I wanted something that could help me reflect on my changes, get feedback, and write better commits automatically.

## 💭 Motivation

I got the idea for gitai after seeing a similar tool that generates commits automatically.
However, I noticed that it sends data like your OS, username, and repository name to a closed-source API, which didn’t feel very safe or transparent.
I wanted something open, local, and privacy-friendly, while also adding a code review feature that could actually help me learn from my mistakes.

As an apprentice, I’m still learning every day — and sometimes I’m unsure whether my code is clean or efficient.
gitai helps me look at my work from another perspective, making it easier to understand what I’m improving and what I can do better next time.

## 🧠 How it works

- You stage your changes with Git.

- gitai runs a git diff on your staged files.

- The diff is sent to a local Ollama instance.

- Depending on your command, gitai uses a predefined prompt to:

  - Create a commit message (codeCommit), or

  - Generate a code review (codeReview).

Everything happens locally — no data leaves your machine.

## 🚀 Commands

```bash
gitai commit # generates a commit title from your git diff
gitai review # generates a code review with ai
gitai timeline # shows commits 
```

<img width="536" height="184" alt="Skjermbilde 2025-10-22 kl  12 31 44" src="https://github.com/user-attachments/assets/7667d1d3-2ba3-4623-ad19-69c5dc11fbb4" />
<img width="1244" height="416" alt="Skjermbilde 2025-10-22 kl  12 32 03" src="https://github.com/user-attachments/assets/0afd1bc5-935b-4a98-a032-203e41327284" />
