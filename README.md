# CodingChallenges-Go

In this branch we are implementing the basic `ccwc` challenge in GoLang.

### Pre-Requisite:

1. Download the test file to run the tests on. [Download Link](https://www.dropbox.com/scl/fi/d4zs6aoq6hr3oew2b6a9v/test.txt?rlkey=20c9d257pxd5emjjzd1gcbn03&dl=0)

### Steps:

- [ ] Return the number of bytes in the file.

  ```sh
  ccwc -c test.txt
  342190 test.txt
  ```

- [ ] Return the number of lines in the file.

  ```sh
  ccwc -l test.txt
  7145 test.txt
  ```

- [ ] Return the word counts the file.

  ```sh
  ccwc -w test.txt
  58164 test.txt
  ```

- [ ] Return the multibyte characters count from the file

  ```sh
  >wc -m test.txt
  339292 test.txt

  >ccwc -m test.txt
  339292 test.txt
  ```

- [ ] Support default option, when no option is passed -c, -l and -w options

  ```sh
  >ccwc test.txt
  7145   58164  342190 test.txt
  ```

- [ ] Support being able to read from standard input if no filename is specified

  ```sh
  >cat test.txt | ccwc -l
  7145
  ```
