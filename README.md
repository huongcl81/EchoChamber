🧬 EchoChamber - Covert C2 via System Logs

"Logs never lie. But sometimes… they whisper." 👀
A sneaky C2 framework that talks through system logs like it's on mute 🔇💻

---

⚡ TL;DR

EchoChamber is a cursed little experiment that turns your system logs into a command-and-control channel.
No sockets. No HTTP. No weird ports. Just… logs.
It’s like steganography but for journalctl.
Yeah, it’s weird. Yeah, it’s fun. Yeah, it might get you a TOKEN 👽

---

🚨 IMPORTANT NOTICE

This project is done by the PROJD team for a Stanford University course project.
Any form of copying or misuse is a violation and strictly prohibited.
This is an open project only for students to learn and practice — don’t go wild with it outside of that.

---

🔥 What Can It Do?

- 🪄 Watch system logs in real-time (`auth.log`, `dmesg`, etc.)
- 🧬 Match sus patterns using regex + vibes (linguistic stego-style)
- 🦾 Launch low-level payloads (written in C++) when magic logs appear
- 🎭 Simulate fake attacker logs to mess with blue teams
- 📜 Config your own cursed patterns with `patterns.yaml`
- 🧼 Stay clean — no noisy connections, no alerts (hopefully)

---

🧠 How It Works

1. You drop a "normal-looking" log line like:
   sshd[1337]: Failed password for invalid user neo from 192.168.1.66 port 22 ssh2

2. EchoChamber daemon scans logs every few secs like:
   tail -f but make it paranoid 🫣

3. If a line matches a trigger in `patterns.yaml`... BOOM 💥

4. It launches your C++ agent, could be anything — reverse shell, file dump, you name it.

5. Everything’s offline stealth mode, no open ports, no alerts.

---

📂 Project Structure Breakdown

cmd/
  echod/          - Go daemon for real-time log watching
  simulate/       - CLI for injecting fake attacker logs

cpp/
  agent/          - Payload executables in C++
  utils/          - Helpers for logging & obfuscation

internal/
  core/           - Go shared logic modules
  decoder/        - Regex and pattern matching engine

configs/
  patterns.yaml   - Your secret sauce: regex-action pairs

scripts/
  setup.sh        - Quick installer, sets up environment
  test_logs.sh    - Generates fake logs to test detection

docs/
  architecture.md - How this beast ticks under the hood
  threat-model.md - Who we’re fighting, and what we assume

---

📜 Example Patterns

- Match SSH brute force log entries to trigger reverse shell  
- Detect sudo root session opens to log or alert  
- Customize your own regex commands in patterns.yaml

---

💡 Why Logs? Why Not Just Use Network?

Because modern defenders watch networks HARD 🔥  
But logs? Logs are everywhere, and most ignore them or think they’re boring.  
This project exploits that blind spot, turning logs into a covert whisper channel.  
Silent, subtle, and oh so hacker.

---

🙏 TOKEN REQUEST

Yo SysCops, Minh here 👋.  
This project is fresh, original, and built with pure passion for the segfault community.  
If you vibe with this, hook me up with a TOKEN please 🙏  
I swear on my keyboard, it’s for research and helping the community, not warez or mining scams.  
Much love, peace out 🫡

---

Made with ☕, 💻, and borderline madness  
By Minh & the PROJD crew — Stanford University course project  
Any copy or misuse = violation 🚫. Keep it real, keep it clean.

---

# End of README - Catch you in the shadows 👻
