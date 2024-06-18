# Build Your Own wc Tool

Solution to the [wc Tool](https://codingchallenges.fyi/challenges/challenge-wc) challenge.

## TODO

- [ ] Add tests

## Requirements

go@1.22

> [!NOTE]
> Refer to the [go.mod](./go.mod) file for correct version to use.

## How to build?

Use following command to build the tool.

```sh
> go build
```

## How to run?

Run following command to get usage information as shown below.

```sh
> ./ccwc -h
Usage: ./ccwc [OPTIONS] [file ...]
  -c    Print the number of bytes in the input
  -l    Print the number of lines in the input
  -m    Print the number of chars in the input
  -w    Print the number of words in the input
```

### Examples with test data

1. Default output - lines, words and bytes

  ```sh
  > ./ccwc test/testdata/test.txt
   7145   58164  342190 test/testdata/test.txt
  ```

2. Count multi-byte characters

  ```sh
  > ./ccwc -m test/testdata/test.txt
  ```

3. Count lines, words and multi-byte characters

  ```sh
  > ./ccwc -l -w -m test/testdata/test.txt
  7145   58164  339292 test/testdata/test.txt
  ```

  > [!WARNING]
  > Go [flag](https://pkg.go.dev/flag) package does not support combination of flags.
  > Each flag must be specified separately.
  >
  > Following command will not work.
  >
  > ```sh
  > > ./ccwc -lwm test/testdata/test.txt
  > flag provided but not defined: -lwm
  > Usage: ./ccwc [OPTIONS] [file ...]
  > -c    Print the number of bytes in the input
  > -l    Print the number of lines in the input
  > -m    Print the number of chars in the input
  > -w    Print the number of words in the input
  > ```

4. Read from stdin

  ```sh
  > cat test/testdata/test.txt | ./ccwc -w
    58164
  ```

5. Count lines, words and bytes characters in multiple files

  ```sh
  > ./ccwc test/testdata/test.txt test/testdata/test.txt
    7145   58164  342190 test/testdata/test.txt
    7145   58164  342190 test/testdata/test.txt
   14290  116328  684380 total
  ```
