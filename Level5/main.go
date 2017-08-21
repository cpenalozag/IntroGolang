package main

import(
  "Level5/lib"
)

func main() {
  notifiers := lib.GetAllConnections()
  for _, c := range notifiers {
    c.Notify()
  }
}
