# CLI flashcards

Command line local application that helps me practice languages using flashcards.

## Installation

Assuming you have Go installed and `$GOPATH` configured, you just need to run

```bash
make install
```

and `fcards` command should become available in your terminal.

## Usage

To start playing with cards from a specific set of files, pass the files as arguments to `fcards play` like this:

```bash
fcards play file1.tsv file2.tsv mydecks/*
```

The program will give a random sample of cards for you to practice.

When ran with no arguments supplied, the command will default to importing all TSV-files from `~/.fcards/tsv/` directory.

To learn more about `play` command consult `fcards play --help`.

## Input format

A file represents a deck of flashcards, one per non-empty line.
Every line is expected to be a tab-separated pair of a question and an answer.

Here's an example fragment of one of my decks:
```tsv
das Gefälle	the gradient
Lange war Deutschland in Kleinstaaten zersplittert	Germany was split in small states for a long time
garen	to ferment
der Käfig	the cage
die Zapfsäule	the petrol pump
abklingen	to fade away
die Rohkost	the raw food
verschlucken	to swallow
benetzen	to moisten
der Schleim	the mucus
```

