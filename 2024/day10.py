def part2():
    with open("input/day10.txt", "r") as reader:
        data = reader.readlines()
        grid_str = [x.strip() for x in data]
        grid = [list(r) for r in grid_str]
        grid = [[int(v) for v in r] for r in grid]

        max_row = len(grid)-1
        max_col = len(grid[0])-1

    def is_valid(r, c):
        # print("is valid?", r,c)
        return r >= 0 and r <= max_row and c >= 0 and c <= max_col

    def make_coord(r,c):
        # print("making for", r, c)
        return {
            "loc": f"({r},{c})",
            "r": r,
            "c": c,
            "value": grid[r][c],
        }

    def get_neighbours(coord):
        neighbours = []

        for (r,c) in [(0, -1), (-1, 0), (0, 1), (1, 0)]:
            new_r = coord["r"] + r
            new_c = coord["c"] + c
            # print("potential new coord", new_r, new_c)
            if is_valid(new_r,new_c):
                neighbour = make_coord(new_r, new_c)
                if neighbour != coord:
                    neighbours.append(neighbour)
            else:
                pass
                # print("invalid potential new coord", new_r, new_c)

        return neighbours


    def search(seen, coord, route):
        route.append(coord["loc"])
        if coord["value"] == 9:
            loc = coord["loc"]
            # print(f"9 found at {loc}")
            
            seen.add(str(route))
            # print("sending back", seen)
            return seen

        neighbours = get_neighbours(coord)

        for n in neighbours:
            if n["value"] == coord["value"]+1:
                seen = seen.union(search(seen, n, route))
                # break
        
        return seen

    trailheads = []
    for r, row in enumerate(grid):
        for c, v in enumerate(row):
            if v == 0:
                coord = make_coord(r,c)
                trailheads.append(coord)

    # print(trailheads)            
    total = 0
    for coord in trailheads:
        #print("starting with", coord)
        routes = search(set(), coord, [coord])
        #print("for coord ", coord["loc"], "9s found at", nines_from_coord)
        total += len(routes)

    print(total)

def part1():
    with open("input/day10_example.txt", "r") as reader:
        data = reader.readlines()
        grid_str = [x.strip() for x in data]
        grid = [list(r) for r in grid_str]
        grid = [[int(v) for v in r] for r in grid]

        max_row = len(grid)-1
        max_col = len(grid[0])-1
        
    def is_valid(r, c):
        # print("is valid?", r,c)
        return r >= 0 and r <= max_row and c >= 0 and c <= max_col

    def make_coord(r,c):
        # print("making for", r, c)
        return {
            "loc": f"({r},{c})",
            "r": r,
            "c": c,
            "value": grid[r][c],
        }
        
    def get_neighbours(coord):
        neighbours = []

        for (r,c) in [(0, -1), (-1, 0), (0, 1), (1, 0)]:
            new_r = coord["r"] + r
            new_c = coord["c"] + c
            # print("potential new coord", new_r, new_c)
            if is_valid(new_r,new_c):
                neighbour = make_coord(new_r, new_c)
                if neighbour != coord:
                    neighbours.append(neighbour)
            else:
                pass
                # print("invalid potential new coord", new_r, new_c)

        return neighbours
        
        
    def search_nines(seen, coord):
        if coord["value"] == 9:
            loc = coord["loc"]
            
            seen.add(str(loc))
            return seen

        neighbours = get_neighbours(coord)

        for n in neighbours:
            if n["value"] == coord["value"]+1:
                seen = seen.union(search_nines(seen, n))
                # break
        
        return seen

    trailheads = []
    for r, row in enumerate(grid):
        for c, v in enumerate(row):
            if v == 0:
                coord = make_coord(r,c)
                trailheads.append(coord)

    # print(trailheads)            
    total = 0
    for coord in trailheads:
        nines_from_coord = search_nines(set(), coord)
        total += len(nines_from_coord)

    print(total)

part1()
part2()