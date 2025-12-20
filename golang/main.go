package main

import (
	"fmt"
	"time"
)

type CustomError struct {
  timestamp time.Time
  error string
}

const ( 
  // Structural tokens 
  BRACE_OPEN string = "{"
  BRACE_CLOSE string = "}"
  BRACKET_OPEN string = "["
  BRACKET_CLOSE string = "]"
  COLON string = ":"
  COMMA string = ","

  // Literals tokens
  STRING = "STRING"
  NUMBER  = "NUMBER"

  // Primitives/Values tokens
  TRUE = "TRUE"
  FALSE = "FALSE"
  NULL = "NULL"

  // EOF Token
  EOF = "EOF"
)

type Token struct {
  token_value string
  token_name string
  start_position int64
  end_position int64
  width int64
  line int64
}



func tokenizer(input string) ([]Token, *CustomError) {
  tokens := []Token{}
  if (len(input) == 0) {
    return tokens, &CustomError{
      time.Now(),
      "No input to parse",
    }
  }
  return tokens, nil
}

func main() {
  json, err := tokenizer("")
  if err != nil  {
    fmt.Println(err)
    return
  }
  fmt.Println(json)
}








