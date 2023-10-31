# PassGen
A secure offline random password generator. It is a command line tool to generate passwords by randomly combining letters, digits and special characters. It runs by Go locally and requires no third-party dependency. Multiple passwords can be generated rapidly, and exported to the terminal or files. Passwords generated are strong, unique and complex to secure your accounts.

## Feature
- Random generation, different character combinations could be set according to various use cases
- Fast generation of multiple passwords concurrently
- Running locally and offline
- No third-party dependency (Only required Go and its standard library)
- No installation required
- Free and open source under GNU general public license

## How to Use It
```
git clone https://github.com/wdrbdev/PassGen.git
cd PassGen
go run PassGen.go
```

## Usage
It generates a random alphanumeric password with a length of 16 shown on the terminal by default.

```
  -char
    	Include special characters (#@!...)
  -delimiter string
    	Delimiter of passwords generated. (default "\n")
  -digit
    	Include digits (123...)
  -h	Print usage (shorthand)
  -help
    	Print usage
  -length int
    	Length of password generated. (default 16)
  -lower
    	Include lower case letters (a-z)
  -n int
    	Total count of passwords generated.(shorthand) (default 1)
  -number int
    	Total count of passwords generated. (default 1)
  -output string
    	The file path as output destination. The default outputs to stdout.
  -unique
    	Exclude duplicate characters
  -unsimilar
    	Exclude similar characters (0oO1lI...)
  -upper
    	Include upper case letters (A-Z)
```

## Example
- Generate a password of a length of 32 with uppercase letters, lowercase letters, numbers, and special characters.

```
go run PassGen.go --length 32 --upper --lower --digit --char
```

- Generate 10 password separated by comma in to a file
```
go run PassGen.go -n 10 --output <file path> --delimiter ,
```

- Show usage information
```
go run PassGen.go -h
```
```
go run PassGen.go --help
```
