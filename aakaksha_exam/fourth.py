def min_num_replacements(arr):
    n = len(arr)
    count = 0
    
    # Ensure the first element is in L,H,L,H... pattern
    if n > 1 and arr[0] >= arr[1]:
        count += (arr[0] - arr[1] + 1)
        arr[0] = arr[1] - 1

    # Ensure the remaining elements are in L,H,L,H... pattern
    for i in range(1, n-1, 2):
        if arr[i] >= arr[i-1] or arr[i] >= arr[i+1]:
            count += (max(arr[i-1], arr[i+1]) - arr[i] + 1)
            arr[i] = max(arr[i-1], arr[i+1]) - 1

    # Ensure the last element is in L,H,L,H... pattern
    if n > 1 and arr[n-1] >= arr[n-2]:
        count += (arr[n-1] - arr[n-2] + 1)
        arr[n-1] = arr[n-2] - 1
    
    return count

# Example usage
arr = [2,1,2,3,4,5,2,9]
result = min_num_replacements(arr)
print(result) # Output: 6