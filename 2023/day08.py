# filepath = "input/day08-test.txt"
filepath = "input/day08.txt"

def clean(row):
    row = row.split("=")
    key = row[0].strip()
    val = row[1].replace("(", "").replace(")", "")
    val = val.split(",")
    val = [x.strip() for x in val]
    return key, val

def part1():



    with open(filepath, "r") as reader:
        data = reader.readlines()

    direction = data[0].strip()
    network = data[2:]


    net = {}
    dirs = {"L": 0, "R": 1}

    for r in network:
        row = clean(r)
        net[row[0]] = row[1]

    pos = "AAA"
    i = 0
    sum = 0

    while pos != "ZZZ":
        dir_idx = dirs[direction[i % len(direction)]]
        pos = net[pos][dir_idx]
        i += 1
        sum += 1

    print(sum)

part1()