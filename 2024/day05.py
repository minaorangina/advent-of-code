
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
    # with open("input/day05.txt", "r") as reader:
    #     input = reader.readlines()

    #     rules = defaultdict(list)

    #     linebrk = 0
    #     for i, line in enumerate(input):
    #         l = line.split("|")
    #         if l[0] == "\n":
    #             linebrk = i
    #             break
    #         rules[l[0]].append(l[1].strip())
        
    #     pages_seq = input[linebrk+1:]
    #     pages_seq = [x.strip().split(",") for x in pages_seq]
    rules, pages_seq = parse("input/day05_example.txt")

    count = 0
    for seq in pages_seq:
        print(seq)
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
            print("valid", mid, mid_idx)
            count += mid
    
    return count
# print(part1()) 


def mergesort(rules, arr):
    if len(arr) <= 1:
        return arr
    
    mid_idx = int(len(arr)/2)
    a = mergesort(rules, arr[0:mid_idx])
    b = mergesort(rules, arr[mid_idx:])

    sorted = []

    print("a",a, "b",b)

    print("looping for", a, b)
    for i in range(min(len(a), len(b))):
        print("round", i)
        # if b[i] in rules[a[i]]:
        if int(b[0]) < int(a[0]):
            print("b is less than a")
            print(b, "bbbb")
            sorted.append(b[0])
            b = b[1:]
            # sorted.append(a[0])
            
        else:
            print("a is less than b")
            print(a, "???")
            sorted.append(a[0])
            a = a[1:]
            # sorted.append(b[0])
        print("sorted thus far", sorted)
        # a = a[1:]
        # b = b[1:]
    
    sorted = sorted + a
    sorted = sorted + b
    print("sorted?",sorted)
    
    return sorted


def part2():
    rules, pages_seq = parse("input/day05_example.txt")
    print(rules)

    count = 0
    for seq in pages_seq:
        print(seq)
        mid_idx = int(len(seq)/2)
        
        valid = True
        mid = 0

        for i, page in enumerate(seq):
            after_pages = seq[i+1:]
            for a in after_pages:
                if a not in rules[page]:
                    valid = False
                    print("invalid", mid_idx)
                    sorted = mergesort(rules,seq)
                    print("sorted", sorted)
                    mid = int(sorted[mid_idx])

        if not valid:
            print("valid", mid, mid_idx)
            count += mid

# part2()
print(mergesort(None,["7","9","2","3"]))