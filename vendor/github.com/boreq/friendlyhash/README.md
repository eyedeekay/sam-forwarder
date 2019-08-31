# Friendly hash [![GoDoc](https://godoc.org/github.com/boreq/friendlyhash?status.svg)](https://godoc.org/github.com/boreq/friendlyhash) [![Build Status](https://travis-ci.org/boreq/friendlyhash.svg?branch=master)](https://travis-ci.org/boreq/friendlyhash) [![codecov](https://codecov.io/gh/boreq/friendlyhash/branch/master/graph/badge.svg)](https://codecov.io/gh/boreq/friendlyhash)

Friendly hash is a Go library which implements human-readable and reversible
representation of known-length byte slices. It can be used to represent hashes
in a human-readable way.

## The core idea

This library aims to make hashes friendlier to humans. As an example the following hash:

    9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08

Can be instead represented in the following way to make it more visually recognizable:

    swiss.laboratory.mostly.parks.inches.therapy.homes.preferred.victory.applicant.making.leading.documentation.ownership.every.models.expense.targets.picture.series.return.signature

Or even displayed using emoji:

    🏘 📠 🥓 😶 😡 🤶 📋 🚶‍♀️ ⌨ 👨‍👩‍👧‍👦 🍷 😆 🍑 🐀 🍣 🌈 🚍 🍏 🐬 💺 🚝 🏯 👏 🧀 🍸 🚇 🙏 😝 🏗


## Length of the output

The table below presents the number of dictionary elements which have to be
used to represent a hash of a specific size using the dictionary containing the
given number of elements.

Example: 11 elements are needed to encode a hash if it is 128 bits (16 bytes)
long and a dictionary of 5000 elements is used.

| Dictionary size | 128 bits | 256 bits | 512 bits |
|-----------------|----------|----------|----------|
| 5000            | 11       | 21       | 42       |
| 10000           | 10       | 20       | 39       |
| 15000           | 10       | 19       | 37       |
| 20000           | 9        | 18       | 36       |
| 30000           | 9        | 18       | 35       |
| 40000           | 9        | 17       | 34       |
| 50000           | 9        | 17       | 33       |
| 60000           | 9        | 17       | 33       |
| 70000           | 8        | 16       | 32       |
| 80000           | 8        | 16       | 32       |
| 90000           | 8        | 16       | 32       |
| 100000          | 8        | 16       | 31       |


## Word list

This package doesn't provide a dictionary of words used for encoding. I can
recommend using one of the following lists available on Github:

- https://github.com/first20hours/google-10000-english
- https://github.com/dwyl/english-words

If you prefer to use emoji the following library has a list of those:

- https://github.com/hackebrot/turtle

## Code example

    dictionary := []string{"word1", "word2", "word3", "word4", "word5", "word6"}
    hashSize := 2

    // Create
    h, err := New(dictionary, hashSize)
    if err != nil {
        panic(err)
    }

    // Humanize
    humanized, err := h.Humanize([]byte{'a', 'b'})
    if err != nil {
        panic(err)
    }
    fmt.Println(strings.Join(humanized, "-"))

    // Dehumanize
    dehumanized, err := h.Dehumanize(humanized)
    if err != nil {
        panic(err)
    }
    fmt.Printf("%q\n", dehumanized)

    // Output:
    // word2-word3-word1-word2-word2-word3-word1-word3
    // "ab"
