# Grepfa

A command line tool for regular expression searches of sequence in .fasta files.

## Introduction

I often find myself trying to quckly identify motifs in genomic data, especially in .fasta files.
When I'm in a hurry, I grep for the things I need, but I always feel a little bad about it, since grep won't
find motifs that cross line boundaries, and chromosome-length lines aren't really understandable through grep.
So, I wrote this utility to search in a useful way for motifs in .fasta files.

## Usage

The usage is as much like grep as possible.

### Usage
```sh
python3 grepfa.py SEARCH_REGEX QUERY_STRING
```

`grepfa` will do a perl-compatible regular expression search for all instances of SEARCH_REGEX in QUERY_STRING.
Here are some examples of usage:

### Identifying stop sites in a gene sequence
```
grepfa '[Aa][Tt][Gg]' in.fa
```

If you run the above command on the following file:
```
>one
cagtcagtcagtcagtcagtcagtcagtcagtcagtcagtcagtcagtcagtcagtgtcagtcagtcagtcagtcagtca
gtcagtcagtcagtcagtc
>two
aaaacgtgtgtgtcgtcatgcgggtcttttgtagcgtcacacgtcgatcaccccgtagtacgtattgcgccgtacgcagt
gtacgtccagtcagtcagtcagtcagtcagtcagcagcggctaatatgcgc
acgcgcat
```

You will get this output:
```
>two	17	20	atg
>two	125	128	atg
```


## Output format
The output is tab-separated, with the following columns
Sequence name | Start | End | Full match sequence | First capture group start | First capture group end | First capture group sequence | ...
