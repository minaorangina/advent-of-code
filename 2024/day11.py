import re
from collections import Counter, defaultdict

def parse(path):
    with open(path, "r") as reader:
        input = reader.readlines()
        input[0] = input[0].strip()
        input[0] = input[0].split(" ")
        input = input[0]
        return input

def part1(path):
    input = parse(path)
    return algo(input)

def algo(input):
    src = input.copy()
    for blink in range (0, 25):
        output = []
        for i, v in enumerate(src):
            idx = len(output) -1

            if v == "0":
                output.append("1")
            
            elif v == "1":
                output.append(str(int(v) * 2024))
            
            elif len(v) % 2 == 0:
                left = v[0:int(len(v)/2)]
                right = v[int(len(v)/2):]

                pattern = r"^0+$"
                if re.match(pattern, left):
                    left = "0"
                if re.match(pattern, right):
                    right = "0"

                if len(left) > 1:
                    left = left.lstrip("0")
                if len(right) > 1:
                    right = right.lstrip("0")

                output.append(left)
                output.append(right)
            
            else:
                output.append(str(int(v)*2024))
            
        got = " ".join(str(x) for x in output)
        # print(f"blink {blink+1}: {got}")
        src = output.copy()
    print(f"number of stones: {len(src)}")
    return len(src)

def do2(value_counts):

    next_level = []
    for k, count in value_counts.items():
        s = str(k)
        if k == 0:
            # should be 1
            # append to next_level
            next_level.append((1, count))
        elif len(s) % 2 == 0:
            # should split in two
            # append to next_level
            l = int(len(s) / 2)
            next_level.append((int(s[:l]), count))
            next_level.append((int(s[l:]), count))
        else:
            # should * 2024
            # append to next_level
            next_level.append((k * 2024, count))

    new_counts = defaultdict(int)
    for (k, v) in next_level:
        new_counts[k] += v
    
    return new_counts


def part2_again_again(path, n_iters):
    input = parse(path)
    input = [int(x) for x in input]

    value_counts = Counter(input)

    for _ in range(n_iters):
        value_counts = do2(value_counts)
    return sum(value_counts.values())


print(part2_again_again("input/day11_example.txt", 75))
