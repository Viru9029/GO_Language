def pairs(k, arr):
    # Write your code here
    arr = sorted(arr)
    counter = 0
    stack = set()

    for i in range(len(arr)):
        if arr[i] in stack:
            counter +=1
        stack.add(k + arr[i])

    return counter


print(pairs(4,[1,3,2,0]))


# if __name__ == '__main__':
#     fptr = open(os.environ['OUTPUT_PATH'], 'w')

#     first_multiple_input = input().rstrip().split()

#     n = int(first_multiple_input[0])

#     k = int(first_multiple_input[1])

#     arr = list(map(int, input().rstrip().split()))

#     result = pairs(k, arr)

#     fptr.write(str(result) + '\n')

#     fptr.close()