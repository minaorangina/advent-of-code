def part1(filepath):
    with open(filepath, "r") as reader:
        line = reader.readline()

        sum = 0
        while line != "":
            card_nums = line.split(":")[1]
            halves = [x.strip() for x in card_nums.split("|")]

            winning = set(halves[0].split())
            have = set(halves[1].split())

            inters = winning.intersection(have)

            matches = len(inters)

            if matches > 0:
                sum += pow(2, matches-1)

            line = reader.readline()

        return sum


assert part1("input/day04-1-test.txt") == 13
print(part1("input/day04-1.txt"))
