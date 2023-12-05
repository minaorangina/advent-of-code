import numpy as np


def is_digit(rows, loc):
    candidate = getAtCoord(rows, loc)
    return candidate.isdigit()


def is_symbol_neighbour(hotspots, loc):
    return loc in hotspots


def is_symbol_neighbour2(hotspots, loc):
    if loc in hotspots:
        return loc


def getAtCoord(rows, loc):
    r = loc[0]
    c = loc[1]
    return rows[r][c]


def print_neighbours(rows, n):
    print("int neighbours: %s" % str(n))
    for dr, loc in n.items():
        print("neighbour at (%s, %s) %s" % (loc[0], loc[1], getAtCoord(rows, loc)))


def part1():
    # filepath = "input/day03-1.txt"
    filepath = "input/day03-1-test.txt"
    sum = 0
    with open(filepath, "r") as reader:
        rows = reader.readlines()

        max_s = len(rows)
        max_e = len(rows[0])

        int_range = np.arange(48, 58, 1)
        symbols = [46, 10]  # dot, new line

        exclusions = np.append(int_range, symbols)

        symbol_locs = []

        for i, r in enumerate(rows):
            for j, char in enumerate(r):
                if ord(char) not in exclusions:
                    symbol_locs.append((i, j))

        stuff = {}
        hotspots = []
        # for each symbol, get valid neighbours
        for sym_loc in symbol_locs:
            stuff[str(sym_loc)] = {"symbol": getAtCoord(rows, sym_loc)}
            r = sym_loc[0]
            c = sym_loc[1]
            # first, find all neighbours
            neighbours = {}
            neighbours["w"] = (r, c - 1)
            neighbours["e"] = (r, c + 1)
            neighbours["n"] = (r - 1, c)
            neighbours["s"] = (r + 1, c)

            neighbours["se"] = (r + 1, c + 1)
            neighbours["sw"] = (r + 1, c - 1)
            neighbours["ne"] = (r - 1, c + 1)
            neighbours["nw"] = (r - 1, c - 1)

            # remove any impossible locations (or should I bother?)
            indexable_neighbours = {}
            for dr, n_loc in neighbours.items():
                x = n_loc[0]
                y = n_loc[1]

                if (0 <= x < max_s) and (0 <= y < max_e):
                    indexable_neighbours[dr] = n_loc

            stuff[str(sym_loc)] = stuff[str(sym_loc)] | indexable_neighbours
            hotspots = hotspots + list(indexable_neighbours.values())
            print(stuff)

        # look through row for digits
        for i, r in enumerate(rows):
            got = "0"
            curr = "0"
            curr_locs = []
            for j, char in enumerate(r):
                loc = (i, j)

                if is_digit(rows, loc):
                    curr += char
                    curr_locs.append(loc)

                # if punctuation, do the check
                else:
                    if curr == "0":
                        # nothing to examine
                        continue

                    valid = False

                    for cl in curr_locs:
                        # curr_char = rows[cl[0]][cl[1]]

                        if is_symbol_neighbour(hotspots, cl):
                            valid = True
                            break

                    if valid:
                        sum += int(float(curr))

                    curr = "0"
                    curr_locs = []
                    continue

    print(sum)


def part2():
    # filepath = "input/day03-1.txt"
    filepath = "input/day03-1-test.txt"
    sum = 0
    with open(filepath, "r") as reader:
        rows = reader.readlines()

        max_s = len(rows)
        max_e = len(rows[0])

        symbol_locs = []

        for i, r in enumerate(rows):
            for j, char in enumerate(r):
                if char == "*":
                    symbol_locs.append((i, j))

        stuff = {}
        hotspots = []
        # for each symbol, get valid neighbours
        for sym_loc in symbol_locs:
            stuff[str(sym_loc)] = {"symbol": getAtCoord(rows, sym_loc)}
            r = sym_loc[0]
            c = sym_loc[1]
            # first, find all neighbours
            neighbours = {}
            neighbours["w"] = (r, c - 1)
            neighbours["e"] = (r, c + 1)
            neighbours["n"] = (r - 1, c)
            neighbours["s"] = (r + 1, c)

            neighbours["se"] = (r + 1, c + 1)
            neighbours["sw"] = (r + 1, c - 1)
            neighbours["ne"] = (r - 1, c + 1)
            neighbours["nw"] = (r - 1, c - 1)

            # remove any impossible locations (or should I bother?)
            indexable_neighbours = {}
            for dr, n_loc in neighbours.items():
                x = n_loc[0]
                y = n_loc[1]

                if (0 <= x < max_s) and (0 <= y < max_e):
                    indexable_neighbours[dr] = n_loc

            stuff[str(sym_loc)] = stuff[str(sym_loc)] | indexable_neighbours
            hotspots = hotspots + list(indexable_neighbours.values())
        print(stuff)

        # look through row for digits
        for i, r in enumerate(rows):
            curr = "0"
            curr_locs = []
            star_locs = {}

            for j, char in enumerate(r):
                loc = (i, j)

                if is_digit(rows, loc):
                    curr += char
                    curr_locs.append(loc)

                # if punctuation, do the check
                else:
                    if curr == "0":
                        # nothing to examine
                        continue

                    for cl in curr_locs:
                        # curr_char = rows[cl[0]][cl[1]]

                        star_loc = is_symbol_neighbour2(hotspots, cl)
                        if star_loc != None:
                            star_locs[str(star_loc)] = star_locs.get(str(star_loc), 0) + 1

                    for k, v in star_locs.items():
                        if v > 1:
                            print("do something")
                            # sum += int(float(curr))

                    curr = "0"
                    curr_locs = []
                    continue

    print(sum)


part2()
