# filepath = "input/day11-test.txt"
filepath = "input/day11.txt"


def do(offset):
    with open(filepath, "r") as reader:
        file = reader.readlines()
        file = [x.strip() for x in file]

    # find all empty cols
    stuff = {
        "rows": [],
        "cols": [],
    }

    for i, row in enumerate(file):
        all_row = True
        for char in row:
            if char != ".":
                all_row = False
                break

        if all_row:
            stuff["rows"].append(i)

    for i in range(len(file[0])):
        all_col = True
        for j in range(len(file)):
            if file[j][i] != ".":
                all_col = False
                break

        if all_col:
            stuff["cols"].append(i)

    locations = []
    new_locations = []
    for i, row in enumerate(file):
        for j, char in enumerate(row):
            if char == "#":
                locations.append([j, i])
                new_locations.append([j, i])

    for y in stuff["rows"]:
        for i, coord in enumerate(locations):
            if coord[1] > y:
                new_locations[i][1] += offset

    for x in stuff["cols"]:
        for i, coord in enumerate(locations):
            if coord[0] > x:
                new_locations[i][0] += offset


    sum = 0
    pair_count = 0
    for i, coord in enumerate(new_locations):
        for coord1 in new_locations[:i]:
            pair_count += 1
            sum += abs(coord[0] - coord1[0]) + abs(coord[1] - coord1[1])
   
    print(sum, pair_count)


def part1():
    return do(1)


def part2():
    return do(1e6-1)


part1()
part2()
