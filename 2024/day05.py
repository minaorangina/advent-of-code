
from collections import defaultdict

def parse(path):
    with open(path, "r") as reader:
        input = reader.readlines()

        rules = defaultdict(list)

        linebrk = 0
        for i, line in enumerate(input):
            l = line.split("|")
            if l[0] == "\n":
                linebrk = i
                break
            rules[l[0]].append(l[1].strip())
        
        pages_seq = input[linebrk+1:]
        pages_seq = [x.strip().split(",") for x in pages_seq]

        return rules, pages_seq

def part1():
    rules, pages_seq = parse("input/day05_example.txt")

    count = 0
    for seq in pages_seq:
        mid_idx = int(len(seq)/2)
        mid = int(seq[mid_idx])
        valid = True

        for i, page in enumerate(seq):
            after_pages = seq[i+1:]
            for a in after_pages:
                if a not in rules[page]:
                    valid = False
                    break
        if valid:
            count += mid
    
    return count

print(part1()) 

def mergesort(rules, arr):
    if len(arr) <= 1:
        return arr
    
    mid_idx = int(len(arr)/2)
    a = mergesort(rules, arr[0:mid_idx])
    b = mergesort(rules, arr[mid_idx:])

    sorted = []

    while len(a) > 0 and len(b) > 0:
        if b[0] not in rules[a[0]]:
            sorted.append(b[0])
            b = b[1:]
        else:
            sorted.append(a[0])
            a = a[1:]
    
    sorted = sorted + a
    sorted = sorted + b
    
    return sorted


def part2():
    rules, pages_seq = parse("input/day05.txt")
    count = 0

    for seq in pages_seq:
        mid_idx = int(len(seq)/2)
        dodgy = False

        for i, page in enumerate(seq):
            after_pages = seq[i+1:]
            for a in after_pages:
                if a not in rules[page]:
                    dodgy = True
                    break
            if dodgy:
                break

        if dodgy:
            sorted = mergesort(rules,seq)
            mid = int(sorted[mid_idx])
            count += mid

    return count

print(part2())