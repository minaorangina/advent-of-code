def parse(path):
    with open(path, "r") as reader:
        lines = reader.readlines()
        clean = [x.strip() for x in lines]
        return clean

def find_guard(grid):
    orientations = ["^",">","v","<"]
    for r, row in enumerate(grid):
        for c, char in enumerate(row):
            if char in orientations:
                return (r, c), char
    raise Exception("guard missing")

class Coords:
    def __init__(self, r, c):
        self.r = r
        self.c = c
    
    def __add__(self, other):
        return Coords(self.r + other.r, self.c + other.c)

    def __repr__(self):
        if self == UP:
            return "UP"
        elif self == DOWN:
            return "DOWN"
        elif self == LEFT:
            return "LEFT"
        elif self == RIGHT:
            return "RIGHT"
        else:
            return f"({self.r},{self.c})"

UP = Coords(-1, 0)
DOWN = Coords(1, 0)
LEFT = Coords(0, -1)
RIGHT = Coords(0, 1)

def turn(c):
    if c == UP:
        return RIGHT
    elif c == RIGHT:
        return DOWN
    elif c == DOWN:
        return LEFT
    elif c == LEFT:
        return UP
    else:
        raise ValueError(c)

def in_bounds(grid, c):
    row_ok = c.r >= 0 and c.r < len(grid)
    col_ok = c.c >= 0 and c.c < len(grid[1])
    return row_ok and col_ok

store = {}
visited = {}

def leads_to_exit(grid, c, d):
    next = c + d # if you don't turn
    visited[str(c) + str(d)] = True

    if str(next) + str(d) in visited:
        return False # loop
    
    if not in_bounds(grid, next):
        store[str(c)+str(d)] = True
        return True # exits (no loop)
    else:
        if grid[next.r][next.c] == '#': # obstacle
            new_d = turn(d)
            result = leads_to_exit(grid, c, new_d)
            store[str(c)+ str(new_d)] = True
            return result
        else: # in bounds and not an obstacle so go in a straight line
            result = leads_to_exit(grid, c + d, d)
            store[str(c + d) + str(d)] = True
            return result

grid = parse("input/day06_example.txt")
#print(leads_to_exit(grid, Coords(6,4), UP))
#grid[6] = f"{grid[6][:3]}#{grid[6][4:]}"

# for r in range(10):
#     for c in range(10):
#         for d in [UP, DOWN, LEFT, RIGHT]:
#             if grid[r][c] not in ["^",">","v","<","#"]:
#                 leads_to_exit(grid, Coords(r, c), d)

print(leads_to_exit(grid, Coords(6, 4), UP))
print(store)
