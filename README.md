# EchoChamber 🧬

## Overview

**EchoChamber** is a hybrid C++/Go research framework that explores _steganographic C2 communication using system logs_.

Imagine this: A system that interprets certain log lines as hidden instructions — like Morse code embedded in dmesg. No sockets, no signals — just subtle patterns, triggered by natural-looking log entries.

> 🧠 Built for hackers, by hackers. Inspired by spycraft, hardened by paranoia.

---

## 🔧 Project Structure

```
cmd/
  echod/          # Go daemon that parses logs and detects C2 patterns
  simulate/       # CLI to simulate log injection

cpp/
  agent/          # C++ agent that reacts to commands (exec, reverse shell, etc)
  utils/          # Reusable utility code (hashing, pattern match)

internal/
  core/           # Log scanning engine
  decoder/        # NLP-inspired pattern decoder

configs/
  patterns.yaml   # Customizable linguistic patterns for hidden commands

scripts/
  setup.sh        # Installer / environment prep
  test_logs.sh    # Fakes log events to test C2

docs/
  architecture.md
  threat-model.md
```

---

## ⚠️ Disclaimer

This project is **for educational purposes only**. Do not deploy EchoChamber on production systems or networks you do not own.

---

## 🚀 Goals

- Covert communication using noisy channels (logs)
- Real-time log monitoring (Go)
- Response mechanism via syscall or C++ binaries
- Supports linguistic steganography in `journalctl`, `auth.log`, `dmesg`

---

> Made with spite & caffeine by Minh ☕️
