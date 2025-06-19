#!/bin/bash
echo "[*] Simulating log entries..."
logger "sshd[4242]: Failed password for invalid user neo from 192.168.1.66 port 22 ssh2"
logger "sudo: pam_unix(sudo:session): session opened for user root by user(uid=0)"
logger "kernel: [123456.789] USB device detected: /dev/hax0r"
echo "[+] Done."