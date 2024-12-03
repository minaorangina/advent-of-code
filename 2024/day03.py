import re



def do(pattern, input):

    total = 0

    for row in input:

        got = re.findall(pattern, row)

        pairs = [x.split(",") for x in got]

        for pair in pairs:
            vals = [int(x) for x in pair]
            total += (vals[0] * vals[1])

    return total

with open("input/day03_example1.txt", "r") as reader:
    input = reader.readlines()

    part1_pattern = "(?<=.mul\()[0-9]{0,3},[0-9]{0,3}(?=\))"
    print(do(part1_pattern, input))


with open("input/day03_example2.txt", "r") as reader:
    input = reader.readlines()

    part2_pattern = "(?<=.mul\()[0-9]{0,3},[0-9]{0,3}(?=\))"
    print(do(part2_pattern, input))
