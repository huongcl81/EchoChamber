# Threat Model

## Assumptions
- Attacker has some level of system access (initial compromise)
- Defender wants to detect lateral movement or hidden C2

## Goals
- Detect unusual access patterns via system logs
- Trigger actions without active connections or sockets

## Limitations
- Requires access to system logs
- Can be bypassed if attacker avoids logging channels
