package main

func main() {
  var blockedIPs []string

  blockedIPs = append(blockedIPs, "192.168.0.16")
  blockedIPs = append(blockedIPs, "192.168.0.17")
  blockedIPs = append(blockedIPs, "192.168.0.18")
  blockedIPs = append(blockedIPs, "192.168.0.19")
  blockedIPs = append(blockedIPs, "192.168.0.20")

  blockedIPs := []string {"192.168.0.19","192.168.0.20"}

  addToBlockedList(blockedIPs)

  for _, element := range langs {
    fmt.Println(element)
  }
  for i := range list {
    fmt.Println(list[i])
  }
}

func addToBlockedList(ips []string) {
  util.SaveBlockedIPs(ips)
}

func isIPBlocked(ip string) bool {
  blockedIPs := getBlockedIPs()
  for _,blockedIP := range blockedIPs {
    if ip==blockedIP{
      return true
    }
  }
  return false
}
