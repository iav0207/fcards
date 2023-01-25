# CLI flashcards

Command line local application that helps me study using flashcards.

## Installation

```bash
make install
```

## Usage

To start playing with cards from a specific set of files, pass the files as arguments to `fcards play` like this:

```bash
fcards play file1.tsv file2.tsv mydecks/*
```

## Input format

A file represents a deck of flashcards, one per non-empty line.
Every line is expected to be a tab-separated pair of a question and an answer.

