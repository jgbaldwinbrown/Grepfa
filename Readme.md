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

## Output format
The output is tab-separated, with the following columns:

Sequence name | Start | End | Full match sequence | First capture group start | First capture group end | First capture group sequence | ...

## Examples

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

### Counting the lengths of all chromosomes

The following will count the lengths of all chromosomes, and will produce a valid `.bed` file containing
all of the bases contained in all chromosomes:

```
grepfa '.+' in.fa | cut -d '	' -f 1,2,3
```

### Identifying gaps in a genome assembly

The following will identify the length, content, and location of all gaps in a genome assembly:

```
grepfa '[Nn]+' in.fa
```

### Identifying restriction cut sites

To identify the following (HindIII) cut sites in a sequence:

```
A|A G C T T
T T C G A|A
```

Use the following invocation:

```
grepfa 'aagctt|ttcgaa' in.fa
```

### Adding context around searches

To add basepairs of context around any search matches, simply add matches for any character to the ends:

```
grepfa '{0,50}(aagctt|ttcgaa){0,50}' in.fa
```

## Other included tools

### Grepfah

`grepfah` works just like `grepfa`, but performs its searches on headers instead of sequences.

### Sedfa

`sedfa` takes three or four arguments as follows:

```
sedfa HEADER_REGEX SEQUENCE_REGEX REPLACE_STRING [INPUT_FILE]
```

If `INPUT_FILE` is excluded, `stdin` is used for input. For each sequence that matches `HEADER_REGEX`, `SEQUENCE_REGEX` will be replaced
with `REPLACE_STRING`. `REPLACE_STRING` can include capture groups using the format `\[0-9]+` to specify the capture group.

#### Removing adapters from the beginning of sequence reads:

Assuming the adapter sequence in question is `aggtctcc`, the following will remove adapters from the beginning of all reads in a .fasta file:

```
sedfa '.*' '.*aggtctcc' '' in.fa
```

#### Simulated mutations

The following will simulate a deletion in a uniquely specified sequence:

```
sedfa '.*' 'atagccgggcta' 'atagcgggcta' in.fa
```
