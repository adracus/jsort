# jsort

Comparing two JSON snippets is hard, if they're not sorted.
Luckily, the Go JSON encoder automatically sorts JSON correctly,
so this implementation just de- and encodes to achieve sorting.

## Installation

```bash
go get github.com/adracus/jsort
```

## Usage

```bash
cat foo.json | jsort
```

