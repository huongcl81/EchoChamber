#!/bin/bash
echo "[*] Setting up EchoChamber environment..."
mkdir -p /opt/echochamber/logs
cp ./configs/patterns.yaml /opt/echochamber/
echo "[+] Done. Run 'go run cmd/echod/main.go' to start."