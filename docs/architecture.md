# EchoChamber Architecture

## Overview
EchoChamber is composed of:

- A Go daemon that watches system logs for steganographic patterns.
- Configurable YAML-based rules for C2 commands.
- Optional C++ binaries for executing low-level payloads.

## Components
- `echod`: Main log parser and pattern matcher
- `patterns.yaml`: Rule configuration
- `rev.cpp`: Example payload for reverse shell
- `logger.cpp`: Utility to write logs
