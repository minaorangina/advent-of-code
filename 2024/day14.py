import re
max_row = 7
max_col = 11
# max_row = 103
# max_col = 101

class Coords:
    def __init__(self, r, c):
        self.r = r
        self.c = c
    
    def __add__(self, other):
        new_coord = Coords(self.r + other.r, self.c + other.c)
        new_coord.r = new_coord.r % max_row
        new_coord.c = new_coord.c % max_col
        return new_coord

    def __mul__(self, n):
        # end = start + (velocity * 100)
        return Coords(self.r * n, self.c * n)

    def __eq__(self, value):
        return self.r == value.r and self.c == value.c

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
        return f"({self.r},{self.c})"
        
    
    def is_valid(self):
        return self.r >= 0 and self.r < max_row and self.c >= 0 and self.c < max_col
    
UP = Coords(-1, 0)
DOWN = Coords(1, 0)
LEFT = Coords(0, -1)
RIGHT = Coords(0, 1)

with open("input/day14_example.txt", "r") as reader:
    lines = reader.readlines()

    row_boundary = (max_row-1)/2
    col_boundary = (max_col-1)/2

    upper_l, upper_r, bottom_l, bottom_r = 0,0,0,0
    for line in lines:
        # get position
        position = [int(x) for x in (re.findall(r"(?<=p=)\d+,\d+", line)[0].split(","))]

        # get velocity
        vel = [int(x) for x in (re.findall(r"(?<=v=)-?\d+,-?\d+", line)[0].split(","))]

        robot = Coords(*position)
        velocity = Coords(*vel)

        end = robot + (velocity * 100)
        
        if end.r < row_boundary and end.c < col_boundary:
            upper_l += 1
        elif end.r < row_boundary and end.c > col_boundary:
            upper_r += 1
        elif end.r > row_boundary and end.c < col_boundary:
            bottom_l += 1
        elif end.r > row_boundary and end.c > col_boundary:
            bottom_r += 1

    print(upper_l * upper_r * bottom_l * bottom_r)