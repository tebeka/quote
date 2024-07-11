# quotes

Use the source, Luke! ☺

## Usage

To print a random quote, use the following command:

```
$ quote
Quiet people have the loudest minds.
    - Richard Feynman
```

### Options

- `-n`: Number of results to return (default is 1)
- `--json`: Output in JSON format

### JSON Example


```
$ quote --json
```

Example JSON output:
```json
[
  {
    "text": "Quiet people have the loudest minds.",
    "author": "Richard Feynman"
  }
]
```
