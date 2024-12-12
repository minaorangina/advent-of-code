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
        # if self == UP:
        #     return "UP"
        # elif self == DOWN:
        #     return "DOWN"
        # elif self == LEFT:
        #     return "LEFT"
        # elif self == RIGHT:
        #     return "RIGHT"
        # else:
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

    def get_region_neighbours(self):
        valid = self.get_valid_neighbours()
        return [n for n in valid if n.get_char() == self.get_char()]

    
UP = Coords(-1, 0)
DOWN = Coords(1, 0)
LEFT = Coords(0, -1)
RIGHT = Coords(0, 1)


def parse(path):
    with open(path, "r") as reader:
        input = reader.readlines()
        input = [x.strip() for x in input]
        return input

def do(curr,key, visited=[], regions={},perimeters={}):
    if curr in visited:
        return
    
    regions[key].append(curr)
    visited.add(curr)

    region_neighbours = curr.get_region_neighbours()
    perimeter = 4 - len(region_neighbours)
    perimeters[key] += perimeter
    

    for n in region_neighbours:
        do(n, key, visited=visited, regions=regions, perimeters=perimeters)



def part1(input):
    regions = defaultdict(list)
    perimeters = defaultdict(int)

    visited = set()
    
    # for loop across whole grid
    for ri, row in enumerate(input):
        for ci, char in enumerate(row):
        # find all neighbours for the first non-seen coord I encounter
            # print(f"now looking at {char} at loc({ri},{ci})")

            curr = Coords(ri, ci, input)
            do(curr,str(curr), visited, regions, perimeters)

    print(regions)
    print(perimeters)

    count = 0
    for k in regions.keys():
        count += len(regions[k]) * perimeters[k]
    print(count)


input = parse("input/day12.txt")
n_rows = len(input)
n_cols = len(input[0])

part1(input)