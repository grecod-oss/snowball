Snowball Ukrainian
================

This package implements the
[Ukrainian language Snowball stemmer](http://snowball.tartarus.org/algorithms/ukrainian/stemmer.html).

## Ukrainian overview

Russian has 32 letters, 10 Vowels, 20 consonants
and 2 unpronounced signs.  The capital letters 
look the same as the lower case letters, with
the exception of cursive capital letter and
lower case.

## Implementation

The Ukrainian language stemmer comprises preprocessing, a number of steps.
Each of these is defined in a separate file in this
package.  All of the steps operate on a `SnowballWord` from the
`snowballword` package and *modify the word in place*.
