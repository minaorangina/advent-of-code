from collections import defaultdict

class Coords:
    def __init__(self, r, c, grid=None):
        self.r = r
        self.c = c
        self.grid = grid
    
    def __add__(self, other):
        return Coords(self.r + other.r, self.c + other.c, self.grid)

    def __eq__(self, value):
        return self.r == value.r and self.c == value.c

    def __hash__(self):
        return self.r * len(self.grid[0]) + self.c

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
            return f"({self.get_char()}@{self.r},{self.c})"
        
    def get_char(self):
        if not self.is_valid():
            return ""
        return self.grid[self.r][self.c]
    
    def is_valid(self):
        return self.r >= 0 and self.r < len(self.grid) and self.c >= 0 and self.c < len(self.grid[0]) 

    def get_neighbours(self):
        return [self + UP, self + DOWN, self + LEFT, self + RIGHT]

    def get_valid_neighbours(self):
        return [i for i in self.get_neighbours() if i.is_valid()]

    
UP = Coords(-1, 0)
DOWN = Coords(1, 0)
LEFT = Coords(0, -1)
RIGHT = Coords(0, 1)


def parse(path):
    with open(path, "r") as reader:
        input = reader.readlines()
        input = [x.strip() for x in input]
        return input



    # return locations

def do(curr, visited, regions, key):
    if curr in visited:
        return
    
    regions[key].append(curr)
    visited.add(curr)
    
    all_neighbours = curr.get_valid_neighbours()
    # print("all neighbours", all_neighbours)

    for n in all_neighbours:
        if n.get_char() == curr.get_char():
            do(n, visited, regions, key)



def part1(input):
    regions = defaultdict(list)

    visited = set()
    
    # for loop across whole grid
    for ri, row in enumerate(input):
        for ci, char in enumerate(row):
        # find all neighbours for the first non-seen coord I encounter
            # print(f"now looking at {char} at loc({ri},{ci})")

            curr = Coords(ri, ci, input)
            do(curr, visited, regions, str(curr))

    print(regions)




input = parse("input/day12_example.txt")
n_rows = len(input)
n_cols = len(input[0])

part1(input)