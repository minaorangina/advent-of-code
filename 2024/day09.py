

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

new_list = files_as_ids.copy()#list(files_as_ids)

left_pointer = 0

with open('input/day09_out.txt', 'w') as the_file:
    the_file.write("".join(str(x) for x in files_as_ids)+"\n")


# print(files_as_ids)
with open('input/day09_out.txt', 'a') as the_file:
    for j, v in reversed(list(enumerate(files_as_ids))):
        if v == ".":
            continue
        # print(f"first idx from back with int: idx {j} with val {v}")

        # if I'm pointing to a dot, do work
        while j > left_pointer:
            if new_list[left_pointer] == ".":
                new_list[left_pointer] = v
                new_list[j] = "."
                # as_str = "".join(str(x) for x in new_list) + "\n"
                # the_file.write(as_str)
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
