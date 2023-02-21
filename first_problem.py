def make_zigzag(arr):
    n = len(arr)
    replacements = 0
    for i in range(1, n, 2):
        if arr[i] >= arr[i-1]:
            replacements += arr[i] - arr[i-1] + 1
            arr[i] = arr[i-1] - 1
        if i < n - 1 and arr[i] >= arr[i+1]:
            replacements += arr[i] - arr[i+1] + 1
            arr[i] = arr[i+1] - 1
    for i in range(0, n, 2):
        if i > 0 and arr[i] <= arr[i-1]:
            replacements += arr[i-1] - arr[i] + 1
            arr[i] = arr[i-1] + 1
        if i < n - 1 and arr[i] <= arr[i+1]:
            replacements += arr[i+1] - arr[i] + 1
            arr[i] = arr[i+1] + 1
    return replacements


print(make_zigzag([2,1,2,3,4,5,2,9]))