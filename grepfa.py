#!/usr/bin/env python3

import re
import sys

def regexit(regex, header, seq):
    matches = regex.finditer(seq)
    if matches:
        for i in matches:
            outlist = [header]
            if i.lastindex:
                for j in range(i.lastindex + 1):
                    outlist.append(i.start(j))
                    outlist.append(i.end(j))
                    outlist.append(i.group(j))
            else:
                outlist.append(i.start())
                outlist.append(i.end())
                outlist.append(i.group())
            print("\t".join(map(str,outlist)))

if __name__ == "__main__":
    if len(sys.argv) <= 1:
        exit("At least 1 argument (regex) required!")
    regexstr = sys.argv[1]
    if len(sys.argv) <= 2:
        inconn = sys.stdin
    elif sys.argv[2] == "-":
        inconn = sys.stdin
    else:
        inconn = open(sys.argv[2],"r")

    regex = re.compile(regexstr)

    header = ""
    seq = ""
    for l in inconn:
        l=l.rstrip('\n')
        if len(l) <= 0:
            pass
        elif l[0]==">":
            if len(header) > 0 and len(seq) > 0:
                regexit(regex, header, seq)
            header = l
            seq = ""
        else:
            seq = seq + l

    regexit(regex, header, seq)

    inconn.close()
