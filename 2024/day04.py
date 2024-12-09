def get_neighbours(loc):
    neighbours = []
    offsets = [(-1, -1), (-1, 1), (1, -1), (1, 1)]
    for i,j in offsets:
        neighbours.append((loc[0]+i, loc[1]+j))
    # for i in range(-1, 2):
    #     for j in range(-1, 2):
    #         neighbours.append((loc[0]+i, loc[1]+j))

    return neighbours

def get_delta(x_loc, m_loc):
    r_delta = m_loc[0] - x_loc[0]
    c_delta = m_loc[1] - x_loc[1]
    return r_delta, c_delta
    
def get_next_loc(grid, loc, char, r_delta, c_delta):
    if loc is not None:
        #max_x = len(grid[0])
        #max_y = len(grid)
        next_loc = (loc[0]+r_delta, loc[1]+c_delta)
        #if next_loc[0] in range(0, max_x) and next_loc[1] in range(0, max_y):
        if next_loc[0] < 0 or next_loc[1] < 0:
            return
        try:
            if grid[next_loc[0]][next_loc[1]] == char:
                return next_loc
        except IndexError:
            pass


def get_a_loc(grid, neighbours):
    max_x = len(grid[0])
    max_y = len(grid)

    all_m = []
    for (r,c) in neighbours:
        #if c in range(0, max_x) and r in range(0, max_y):
        try:
            if grid[r][c] == "A":
                all_m.append((r,c))
        except IndexError:
            pass
    return all_m


def part2(grid):
    count = {}
    for r, row in enumerate(grid):
        for c, char in enumerate(row):
            # print(x, char)
            if char == "M":
                #print(r,c)
                # x_loc = (x,y)
                neighbours = get_neighbours((r,c))
                #print(neighbours)
                a_locs = get_a_loc(grid, neighbours)
                #print(m_locs)
                
                for a_loc in a_locs:
                    # get direction of travel
                    r_delta, c_delta = get_delta((r, c), a_loc)
                    
                    #a_loc = get_next_loc(grid, m_loc, "A", r_delta, c_delta)
                    s_loc = get_next_loc(grid, a_loc, "S", r_delta, c_delta)
                    
                    if a_loc is not None and s_loc is not None:
                        if a_loc in count.keys():
                            count[a_loc] += 1
                        else:
                            count[a_loc] = 1
                        print(r,c)
                        print(a_loc)
                        print(s_loc)
    count = {k: v for k,v in count.items() if v == 2}
    return len(count)

with open("input/day04_example.txt", "r") as reader:
    input = reader.readlines()
    clean = [x.strip() for x in input]
    print(part2(clean))


def get_neighbours(loc):
    neighbours = []
    for i in range(-1, 2):
        for j in range(-1, 2):
            neighbours.append((loc[0]+i, loc[1]+j))

    return neighbours

def get_delta(x_loc, m_loc):
    r_delta = m_loc[0] - x_loc[0]
    c_delta = m_loc[1] - x_loc[1]
    return r_delta, c_delta
    
def get_next_loc(grid, loc, char, r_delta, c_delta):
    if loc is not None:
        #max_x = len(grid[0])
        #max_y = len(grid)
        next_loc = (loc[0]+r_delta, loc[1]+c_delta)
        #if next_loc[0] in range(0, max_x) and next_loc[1] in range(0, max_y):
        if next_loc[0] < 0 or next_loc[1] < 0:
            return
        try:
            if grid[next_loc[0]][next_loc[1]] == char:
                return next_loc
        except IndexError:
            pass


def get_m_loc(grid, neighbours):
    max_x = len(grid[0])
    max_y = len(grid)

    all_m = []
    for (r,c) in neighbours:
        #if c in range(0, max_x) and r in range(0, max_y):
        try:
            if grid[r][c] == "M":
                all_m.append((r,c))
        except IndexError:
            pass
    return all_m


def part1(grid):
    count = 0
    for r, row in enumerate(grid):
        for c, char in enumerate(row):
            # print(x, char)
            if char == "X":
                #print(r,c)
                # x_loc = (x,y)
                neighbours = get_neighbours((r,c))
                #print(neighbours)
                m_locs = get_m_loc(grid, neighbours)
                #print(m_locs)
                
                for m_loc in m_locs:
                    # get direction of travel
                    r_delta, c_delta = get_delta((r, c), m_loc)
                    
                    a_loc = get_next_loc(grid, m_loc, "A", r_delta, c_delta)
                    s_loc = get_next_loc(grid, a_loc, "S", r_delta, c_delta)
                    
                    if a_loc is not None and s_loc is not None:
                        count += 1
                        print(r,c)
                        print(a_loc)
                        print(s_loc)
    return count

with open("input/day04_example.txt", "r") as reader:
    input = reader.readlines()
    clean = [x.strip() for x in input]
    print(part1(clean))