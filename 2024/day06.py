def parse(path):
    with open(path, "r") as reader:
        lines = reader.readlines()
        clean = [x.strip() for x in lines]
        return clean


def valid(position):
    r, c = get_row_col(position)
    if r < 0 or r >= max_row_len:
        return False
    if c < 0 or c >= max_col_len:
        return False
    return True
    
def get_next_position(position, guard):
    r, c = get_row_col(position)
    row_delta, col_delta = 0,0
    if guard == "^":
        row_delta = -1
    elif guard == ">":
        col_delta = 1
    elif guard == "v":
        row_delta = 1
    elif guard == "<":
        col_delta = -1
    else:
        raise Exception("invalid guard")

    new_row = r + row_delta
    new_col = c + col_delta
    return (new_row,new_col)
    
def get_row_col(position):
    return position[0], position[1]

def rotate_guard(guard):
    return next_guards[guard]

def find_guard(grid):
    orientations = ["^",">","v","<"]
    for r, row in enumerate(grid):
        for c, char in enumerate(row):
            if char in orientations:
                return (r, c), char
    raise Exception("guard missing")

def part1(grid):
    position, guard = find_guard(grid)
    visited = set()
    while valid(position):
        r, c = get_row_col(position)
        visited.add(f"{r},{c}")
        # am i at the end?
        next_pos = get_next_position(position, guard)
        if not valid(next_pos):
            break
        
        # can i progress?
        r, c = get_row_col(next_pos)
        next_occupant = grid[r][c]

        if next_occupant == "#":
            guard = rotate_guard(guard)
            continue
        else:
            position = next_pos
    return len(visited)

def part2(grid):
    position, guard = find_guard(grid)
    visited = set()
    while valid(position):
        r, c = get_row_col(position)
        visited.add(f"{r},{c}")
        # am i at the end?
        next_pos = get_next_position(position, guard)
        if not valid(next_pos):
            break
        
        # can i progress?
        r, c = get_row_col(next_pos)
        next_occupant = grid[r][c]

        if next_occupant == "#":
            guard = rotate_guard(guard)
            continue
        else:
            position = next_pos
    return len(visited)
    
grid = parse("input/day06.txt")
max_row_len = len(grid[0])
max_col_len = len(grid)
next_guards = {
    "^":">",
    ">":"v",
    "v":"<",
    "<":"^",
}

print(part1(grid))