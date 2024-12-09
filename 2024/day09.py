

def part1():
    with open("input/day09.txt", "r") as reader:
        input = reader.readline().strip()

    files_as_ids = []
    for file_id, block_size in enumerate(input):
        if file_id % 2 == 0:
            file_id_int = int(file_id/2)
            to_append = [file_id_int for x in range(int(block_size))]
            files_as_ids += to_append
        else:
            to_append = ["." for x in range(int(block_size))]
            files_as_ids += to_append

    new_list = files_as_ids.copy()

    left_pointer = 0

    with open('input/day09_out.txt', 'w') as the_file:
        the_file.write("".join(str(x) for x in files_as_ids)+"\n")

    with open('input/day09_out.txt', 'a') as the_file:
        for j, v in reversed(list(enumerate(files_as_ids))):
            if v == ".":
                continue
            
            while j > left_pointer:
                if new_list[left_pointer] == ".":
                    new_list[left_pointer] = v
                    new_list[j] = "."
                    break
                else:
                    left_pointer += 1
        
        total = 0
        for k, d in enumerate(new_list):
            if d == ".":
                continue
            total += (k * d)
        print(new_list[:100])
        print(total)

def part2():
    with open("input/day09.txt", "r") as reader:
        input = reader.readline().strip()

    files_as_ids = []
    for file_id, block_size in enumerate(input):
        if file_id % 2 == 0:
            file_id_int = int(file_id/2)
            to_append = [file_id_int for x in range(int(block_size))]
            files_as_ids += to_append
        else:
            to_append = ["." for x in range(int(block_size))]
            files_as_ids += to_append

    new_list = files_as_ids.copy()
    print(new_list)

    with open('input/day09_out.txt', 'w') as the_file:
        the_file.write("".join(str(x) for x in files_as_ids)+"\n")

    left_pointer = 0

    with open('input/day09_out.txt', 'w') as the_file:
        the_file.write("".join(str(x) for x in files_as_ids)+"\n")

    val = None
    val_count = 1
    with open('input/day09_out.txt', 'a') as the_file:
        for j in range(len(files_as_ids) - 1, -1, -1):
            v = files_as_ids[j]
            
            if v == ".":
                continue
            
            # one value, but we need the contiguous block
            next_in_line = files_as_ids[j - 1]
            if next_in_line == v:
                val_count += 1
            else:
                # look for somewhere to put val_count lots of v
                space_count = 0
                for i, vi in enumerate(files_as_ids):
                    if i >= j: # must be to the left
                        break
                    if vi == '.':
                        space_count += 1
                        if space_count >= val_count:
                            files_as_ids[i-val_count+1:i+1] = [v for i in range(val_count)]
                            files_as_ids[j:j+val_count] = ["." for i in range(val_count)]
                            break
                    else:
                        space_count = 0
                
                val_count = 1

        total = 0
        for k, d in enumerate(files_as_ids):
            if d == ".":
                continue
            total += (k * d)
            
        print(total)

part2()