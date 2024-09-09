# CodingChallenges-Go

In this challenge we are going to implement JSON parser in Golang.

### Pre-Requisite:

Download the test data from the [Link](https://www.dropbox.com/s/vthtr4897fkuhw8/tests.zip?dl=0)

### Steps:

NOTE: It should take standard input from CLI and either return the parsed JSON or error code.
EX:

```sh
cat <fileName>.json | ccjp
```

- [x] Identify Valid JSON to Invalid JSON.
      When a proper JSON is passed, it should exit with status code 0 and simply print the Input, othewise with status code 1.

- [x] Extend the JSON Parser to simple String Key and String Value

  ```json
  { "key": "value" }
  ```

- [x] In this step, the goal is to extend the parser for the simple value where key will be string and the value can be string, number, null, float and boolean.

  ```json
  {
    "key1": true,
    "key2": false,
    "key3": null,
    "key4": "value",
    "key5": 101
  }
  ```

- [x] In this step, we want to parse arrays and objects in the JSON:

  ```json
  {
    "key": "value",
    "key-n": 101,
    "key-o": {},
    "key-l": []
  }
  ```

- [x] In this step we are going to thoroughly test our implementation based on the numerouse test cases.
