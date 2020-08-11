#!/usr/bin/env python3

import re
import sys

def regexit(header_regex, regex, sub, header, seq):
    if header_regex.search(header):
        seq = regex.sub(sub, seq)
    print(header)
    print(seq)

if __name__ == "__main__":
    if len(sys.argv) <= 2:
        exit("At least 3 argument (regex) required!")
    header_regexstr = sys.argv[1]
    regexstr = sys.argv[2]
    sub = sys.argv[3]
    if len(sys.argv) <= 4:
        inconn = sys.stdin
    elif sys.argv[4] == "-":
        inconn = sys.stdin
    else:
        inconn = open(sys.argv[4],"r")

    header_regex = re.compile(header_regexstr)
    regex = re.compile(regexstr)

    header = ""
    seq = ""
    for l in inconn:
        l=l.rstrip('\n')
        if len(l) <= 0:
            pass
        elif l[0]==">":
            if len(header) > 0 and len(seq) > 0:
                regexit(header_regex, regex, sub, header, seq)
            header = l
            seq = ""
        else:
            seq = seq + l

    regexit(header_regex, regex, sub, header, seq)

    inconn.close()
