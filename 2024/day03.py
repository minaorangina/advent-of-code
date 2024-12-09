import re

part1_pattern = "(?<=.mul\()[0-9]{0,3},[0-9]{0,3}(?=\))"
part2_pattern = "((?<=don't\(\)).*?(?=do\(\))|(?<=don't\(\)).*(?=$))"


def do(pattern, input):

    total = 0

    for row in input:
        # print("row", row, "pattern", pattern)
        got = re.findall(pattern, row)

        pairs = [x.split(",") for x in got]

        for pair in pairs:
            vals = [int(x) for x in pair]
            total += (vals[0] * vals[1])

    return total

def part1():
    with open("input/day03_example.txt", "r") as reader:
        input = reader.readlines()
        print(do(input, part1_pattern))


def part2():
    with open("input/day03.txt", "r") as reader:
        input = reader.readlines()

        count = 0
        # print("input rows ", len(input))
        # TODO - join rows
        for row in input:

            to_replace = re.findall(part2_pattern, row)

            out = row
            for rep in to_replace:
                out = out.replace(rep, "")
            print(out,"\n")
            res = do(part1_pattern, [out])
            count += res
            print("ending in", row[len(row)-20:])
            print(res)
        
        
        print(count)
    

part2()
