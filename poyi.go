package main

import(
  "strings"
)

func poyiReturn(str string) string {
    var result string
    if strings.Contains(str, "G仔"){
      if strings.Contains(str, "早安"){result = "poyi哥早安啊"}
    }
    return result
}
