import re


def part1(filepath):
    with open(filepath, "r") as reader:
        line = reader.readline()

        sum = 0
        while line != "":
            card_nums = line.split(":")[1]
            halves = [x.strip() for x in card_nums.split("|")]

            winning = halves[0].split()
            have = halves[1].split()

            points = 0

            for w in winning:
                if w in have:
                    if points == 0:
                        points = 1
                    else:
                        points *= 2

            sum += points
            line = reader.readline()

        return sum


assert part1("input/day04-1-test.txt") == 13
print(part1("input/day04-1.txt"))
