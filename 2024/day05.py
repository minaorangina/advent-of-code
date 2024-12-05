
from collections import defaultdict

with open("input/day05_example.txt", "r") as reader:
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
    
    print(count)