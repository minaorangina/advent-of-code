
from collections import defaultdict

def parse(path):
    with open(path, "r") as reader:
        grid = [x.strip() for x in reader.readlines()]
        return grid

def get_antinode_loc(origin_coord, match_coord, amplitude):
    origin_r, origin_c = origin_coord[0], origin_coord[1]
    match_r, match_c = match_coord[0], match_coord[1]
    delta_r = (origin_r - match_r) * amplitude
    delta_c = (origin_c - match_c) * amplitude
    antinode_r = origin_r + delta_r
    antinode_c = origin_c + delta_c

    # do something with antinode here
    

    return (antinode_r, antinode_c)

def loc_exists(max_row, max_col, loc):
    loc_r, loc_c = loc[0], loc[1]
    if loc_r < 0 or loc_c < 0:
        return False
    if loc_r > max_row:
        return False
    if loc_c > max_col:
        return False
    return True

def loc_available(grid, loc):
    loc_r, loc_c = loc[0],loc[1]
    return grid[loc_r][loc_c] == "."

# take coords, a delta, the grid
# as long as new coords are valid and new coords is unoccupied,
# add to list of antinodes
def write_grid(grid):
    with open('input/day08_out.txt', 'w') as the_file:
        for r in grid:
            the_file.write(r+"\n")

def update_grid(grid, loc):
    loc_r, loc_c = loc[0],loc[1]

    char = "#"
    if not loc_available(grid, loc):
        # occupied by an antenna
        char = grid[loc_r][loc_c]

    new_row = grid[loc_r]
    new_row = new_row[0:loc_c] + char + new_row[loc_c+1:]
  
    grid[loc_r] = new_row

def do2(grid, curr_loc, next_loc, max_row, max_col, antinodes, amplitude):
    print(f"---- amplitude {amplitude}")

    anti_loc = get_antinode_loc(curr_loc, next_loc,amplitude)
    
    if not loc_exists(max_row, max_col, anti_loc):
        print(f"    out of bounds for {curr_loc} -> {anti_loc}")
        return
    
    anti_r, anti_c = anti_loc[0], anti_loc[1]
    print(f"    antinode for {curr_loc} -> {anti_loc}\n    currently occupied by {grid[anti_r][anti_c]}")
    antinodes.add(anti_loc)

    update_grid(grid, anti_loc)

    do2(grid, curr_loc, next_loc, max_row, max_col, antinodes, amplitude+1)


def part2(path):
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
                print(f"for char {char} @ {origin_loc}, match at {match_loc}")
                do2(grid, origin_loc, match_loc, max_row, max_col, antinodes, 1)

    
    write_grid(grid)
    print(antinodes)
    return len(antinodes)

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
 
                if loc_exists(max_row, max_col, anti_loc):
                    antinodes.add(anti_loc)
    
    return len(antinodes)


# print(part1("input/day08_example.txt"))
print(part2("input/day08_example.txt"))