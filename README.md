# sha256check

Simple tool to check sha256 sums match on files.

# Description

sha256check is a simple tool that will read a file and verify its sha256 sum to the one supplied.

It over comes the bash sha256sum limitations as it actually verifies the sum rather than just displaying it.
It can also read from the STDIN pipe so you can use it with inline echo and cat calls.

Examples:

Using STDIN:

```sh
$> ./sha256check -verbose -s 86b0c5a1e2b73b08fd54c727f4458649ed9fe3ad1b6e8ac9460c070113509a1e -f - <<< echo "Toaster" 
$> OK
$> echo $?
$> 0
```

Using a file:

```sh
$> ./sha256check -verbose -s 86b0c5a1e2b73b08fd54c727f4458649ed9fe3ad1b6e8ac9460c070113509a1e -f ./testfile.txt
$> FAIL
$> echo $?
$> 1
```

Displaying only:

```sh
$> ./sha256check -d -f ./testfile.txt
$> b25a793d1c50f41f9cfddca4c4dadd4f63cdb8e61ff0d1f2d52b449417fd4f0d
```

## Help menu

```text
  -d    Display the SHA256 sum value of the passed in file path
  -f string
        The file that you want to check. Use '-' for STDIN
  -h    Shows the help menu
  -s string
        SHA256 value that you want to assert against
  -v    Show the version of the application
  -verbose
        Enables verbose logging
```
