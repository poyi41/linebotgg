package main

import(
  "strings"
)

func poyiReturn(str string) string {
    var result string
    if strings.Contains(str, "G仔"){
      result = "是poyi哥";
    }
    return result
}
