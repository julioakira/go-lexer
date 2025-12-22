package main

import (
  "fmt"
  "time"
  "unicode"
  "strings"
)

type CustomError struct {
  timestamp time.Time
  position string
  error string
}

const ( 
  // Structural tokens 
  TKN_BRACE_OPEN rune = '{'
  NAME_BRACE_OPEN string = "BRACE_OPEN"
  TKN_BRACE_CLOSE rune = '}'
  NAME_BRACE_CLOSE string = "BRACE_CLOSE"
  TKN_BRACKET_OPEN rune = '['
  NAME_BRACKET_OPEN string = "BRACKET_OPEN"
  TKN_BRACKET_CLOSE rune = ']'
  NAME_BRACKET_CLOSE string = "BRACKET_CLOSE"
  TKN_COLON rune = ':'
  NAME_COLON string = "COLON"
  TKN_COMMA rune = ','
  NAME_COMMA string = "COMMA"
  TKN_QUOTE rune = '"'
  NAME_QUOTE string = "QUOTE"

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
  value string
  name string
  start_position int
  end_position int
  width int
}



func tokenizer(input string) ([]Token, *CustomError) {
  tokens := []Token{}
  inputLen := len(input)
  idx := 0

  if (inputLen == 0) {
    return tokens, &CustomError{
      timestamp: time.Now(),
      position: "0",
      error: "No input to parse",
    }
  }

  for idx < inputLen {
    char := rune(input[idx])
    if (unicode.IsSpace(char)) {
      idx++;
      continue
    }

    switch char {
    case TKN_BRACE_OPEN:
      tokens = append(tokens, Token{
        value: string(char),
        name: NAME_BRACE_OPEN,
        start_position: idx,
        end_position: idx + 1,
        width: 1,
      })
      idx++;

    case TKN_BRACE_CLOSE:
      tokens = append(tokens, Token{
        value: string(char),
        name: NAME_BRACE_CLOSE,
        start_position: idx,
        end_position: idx + 1,
        width: 1,
      })
      idx++;

    case TKN_BRACKET_OPEN:
      tokens = append(tokens, Token{
        value: string(char),
        name: NAME_BRACKET_OPEN,
        start_position: idx,
        end_position: idx + 1,
        width: 1,
      })
      idx++;

    case TKN_BRACKET_CLOSE:
      tokens = append(tokens, Token{
        value: string(char),
        name: NAME_BRACKET_CLOSE,
        start_position: idx,
        end_position: idx + 1,
        width: 1,
      })
      idx++;

    case TKN_COLON:
      tokens = append(tokens, Token{
        value: string(char),
        name: NAME_COLON,
        start_position: idx,
        end_position: idx + 1,
        width: 1,
      })
      idx++;

    case TKN_COMMA:
      tokens = append(tokens, Token{
        value: string(char),
        name: NAME_COMMA,
        start_position: idx,
        end_position: idx + 1,
        width: 1,
      })
      idx++;

    case TKN_QUOTE:
      tokens = append(tokens, Token{
        value: string(char),
        name: NAME_QUOTE,
        start_position: idx,
        end_position: idx + 1,
      })
      idx++;
      strStart := idx;

      for (idx < inputLen && input[idx] != '"') {
        if (char == '\\') {
          idx++;
        }
        idx++;
      }

      currentString := input[strStart:idx];
      if (idx >= inputLen) {
        return nil, &CustomError{
          timestamp: time.Now(),
          position: "0",
          error: "Error: Unterminated string",
        }
      }
      tokens = append(tokens, Token{
        value: string(currentString),
        name: STRING,
        start_position: idx,
        end_position: idx + len(currentString),
        width: len(currentString),
      })
      idx++;

    default:
      remainingString := input[idx:]

      if (strings.HasPrefix(remainingString, "true")) {
        tokens = append(tokens, Token{
          value: "true",
          name: TRUE,
          start_position: idx,
          end_position: idx + 4,
          width: 4,
        })
        idx = idx + 4
      }
      if (strings.HasPrefix(remainingString, "false")) {
        tokens = append(tokens, Token{
          value: "false",
          name: TRUE,
          start_position: idx,
          end_position: idx + 5,
          width: 5,
        })
        idx = idx + 5
      }
      if (strings.HasPrefix(remainingString, "null")) {
        tokens = append(tokens, Token{
          value: "null",
          name: NULL,
          start_position: idx,
          end_position: idx + 4,
          width: 4,
        })
        idx = idx + 4
      }
      
  }
    idx++
  }
  return tokens, nil
}


func main() {
  raw := `
    {
  "array": [
    1,
    2,
    3
  ],
  "boolean": true,
  "color": "gold",
  "null": null,
  "number": 123,
  "object": {
    "a": "b",
    "c": "d"
  },
  "string": "Hello World"
}
  `
  fmt.Printf("input Len: %d\n", len(raw))
  tokenized, err := tokenizer(raw)
  if err != nil  {
    fmt.Println(err)
    return
  }
  fmt.Printf("%+v\n", tokenized)
}








