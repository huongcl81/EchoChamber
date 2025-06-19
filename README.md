EchoChamber — Covert C2 via System Logs

> "The quieter you become, the more you are able to hear." — Ram Dass
> EchoChamber listens where no one else is paying attention.

What is EchoChamber?
--------------------
EchoChamber is a research project in stealthy command-and-control (C2) techniques.
It leverages system logs as a covert communication channel — interpreting carefully
crafted log entries as hidden instructions.

Instead of relying on network traffic, signals, or traditional backdoors, this tool
watches auth.log, journalctl, dmesg, and similar sources for linguistically disguised
patterns that signal an action.

Core Features
-------------
- Real-time log surveillance (auth.log, dmesg, etc.)
- Regex + NLP pattern matching to detect hidden instructions
- Linguistic steganography: “normal” log messages act as triggers
- C++ agents for executing actions (reverse shell, file access, etc.)
- Simulation scripts to mimic attacker behavior
- Configurable patterns (patterns.yaml) to craft your own covert codes
- No sockets, no noise — pure observational stealth

How It Works
------------
1. You write a log message that looks totally normal:
   sshd[4242]: Failed password for invalid user neo from 192.168.1.66 port 22 ssh2

2. EchoChamber's daemon (written in Go) scans new logs every few seconds.

3. It uses regex patterns from configs/patterns.yaml to detect encoded triggers.

4. When a match is found, it runs a C++ agent (like spawning a reverse shell).

5. Everything happens passively. No inbound connections needed.

Project Structure
-----------------
cmd/
  echod/          - Go daemon for real-time log monitoring
  simulate/       - CLI to inject fake log entries
cpp/
  agent/          - C++ payload executables
  utils/          - Logging, obfuscation helpers
internal/
  core/           - Shared Go logic
  decoder/        - Regex & linguistic pattern matcher
configs/
  patterns.yaml   - Triggers defined as regex-action pairs
scripts/
  setup.sh        - Installer script
  test_logs.sh    - Log simulation test suite
docs/
  architecture.md - How EchoChamber is built
  threat-model.md - Adversary simulation assumptions

Request for TOKEN
-----------------
This project is experimental, original, and purpose-built to align with the values
of the segfault community — exploring new frontiers in system-based C2, and sharing
tools for Red/Blue team simulation.

If the SysCops find this project aligns with the community spirit, I humbly request
consideration for a TOKEN — or at least your attention. 🫡

- Minh
