import re
from collections import Counter

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

def part2(path):
    input = parse(path)
    max_depth = 25
    count = 0
    count += len(algo_two(input, max_depth))
    return count

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

                # TODO - deal with zeros
                output.append(left)
                output.append(right)
            
            else:
                output.append(str(int(v)*2024))
            
        got = " ".join(str(x) for x in output)
        # print(f"blink {blink+1}: {got}")
        src = output.copy()
    print(f"number of stones: {len(src)}")
    return len(src)

def do(src, depth=1, max_depth=1):
    if depth > max_depth:
        return src
    depth=depth+1

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

            # TODO - deal with zeros
            output.append(left)
            output.append(right)
        
        else:
            output.append(str(int(v)*2024))
    
    return do(output, depth, max_depth)

def algo_two(input=input, max_depth=0):
    src = input.copy()
    for blink in range (0, max_depth):
        src = do(src, max_depth).copy()
        
    got = " ".join(str(x) for x in src)
    # print(f"blink {blink+1}: {got}")
    # src = output.copy()
    print(f"number of stones: {len(src)}")
    return src


def part2_again(path):
    input = parse(path)
    input = [1]#, 17]
    quants = {1: 1}#, 17: 1}
    max_depth = 8
    
    for iteration in range(max_depth):
        new_dict = {}
        new_list = []
        for i,xi in enumerate(input):
            s = str(xi)
            q = quants[xi]
            if len(s) % 2 == 0: # even digits
                l = int(len(s) / 2)
                new_dict[int(s[:l])] = q
                new_dict[int(s[l:])] = q
                new_list.append(int(s[:l]))
                new_list.append(int(s[l:]))
            elif xi == 0:
                new_dict[1] = q
                new_list.append(1)
            else:
                new_dict[xi * 2024] = q
                new_list.append(xi * 2024)
        #print("---")
        #print(new_list)
        counts = Counter(new_list)
        #print(counts)
        input = counts.keys()
        quants = {k: v * counts[k] for k, v in new_dict.items()}
        print(quants)
        print(sum(quants.values()))
    return sum(quants.values())

print(part2_again("input/day11_example.txt"))
