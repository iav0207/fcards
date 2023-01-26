# CLI flashcards

Command line local application that helps me practice languages using flashcards.

Key features inlcude:

- practice with a random sample taken from a given set of flashcards
- practice cards in the default direction, or inverse (questions swapped with answers), or random, which is my favorite
- at the end of a round reiterate over the cards that were given incorrect answers

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

The program will give a random sample of cards for you to practice, and will let you repeat the ones you gave wrong answers to.

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

### My flashcards collection

If you're curious about or willing to use my collection of flashcards, you can find them in [my flashcard decks repo](https://github.com/iav0207/my-flashcards-decks).
I link the collection to the TSV folder of `fcards` tool to always have the decks in sync, like this

```bash
mv ~/.fcards/tsv ~/.fcards/tsv.bak
ln -sfF "${path_to_decks_repo}/flashcards/tsv" ~/.fcards/tsv
```

You can use that or the script

```bash
make use_cards_from path="${path_you_want_to_use_for_decks}"
```

