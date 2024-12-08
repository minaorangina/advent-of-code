
from collections import defaultdict

def parse(path):
    with open(path, "r") as reader:
        grid = [x.strip() for x in reader.readlines()]
        return grid

def get_antinode_loc(origin_coord, match_coord):
    origin_r, origin_c = origin_coord[0], origin_coord[1]
    match_r, match_c = match_coord[0], match_coord[1]
    delta_r = origin_r - match_r
    delta_c = origin_c - match_c
    antinode_r = origin_r + delta_r
    antinode_c = origin_c + delta_c

    return (antinode_r, antinode_c)

def is_valid(max_row, max_col, loc):
    loc_r, loc_c = loc[0], loc[1]
    if loc_r < 0 or loc_c < 0:
        return False
    if loc_r > max_row:
        return False
    if loc_c > max_col:
        return False
    return True

def part1(path):
    grid = parse(path)

    max_row = len(grid)-1
    max_col = len(grid[0])-1

    antennae = defaultdict(list)
    antinodes = set()

    for r, row in enumerate(grid):
        for c, char in enumerate(row):
            if char == ".":
                continue
            antennae[char].append((r,c))
    
    print(antennae)

    for char, locations in antennae.items():
        for origin_loc in locations:
            for match_loc in locations:
                if match_loc == origin_loc:
                    continue
                anti_loc = get_antinode_loc(origin_loc, match_loc)
 
                if is_valid(max_row, max_col, anti_loc):
                    antinodes.add(anti_loc)
    
    return len(antinodes)


print(part1("input/day08.txt"))