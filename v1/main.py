"""
Understanding the Problem
You are given an integer array arr of size n.
Assume a sliding window of size k starting from index 0.
In each iteration, the sliding window moves to the right by one position till n-k.
Write a program to return an array representing the maximum number in all sliding windows.

Problem Note

The first element of the resultant array is max(arr[0...k]) , then the second element is max(arr[1...k+1]) and so on.
The size of the resultant array will be n-k+1 .
You are expected to solve this question in O(n) time complexity
Example 1

Input: arr[] = [4, 3, 8, 9, 0, 1], k = 3
Output: [8, 9, 9, 9]
Explanation: The window size is 3 and the maximum at different iterations are as follows:
max(4, 3, 8) = 8
max(3, 8, 9) = 9
max(8, 9, 0) = 9
max(9, 0, 1) = 9
Hence, we get arr = [8, 9, 9, 9] as output.
Example 2

Input: arr[] = [9, 8, 6, 4, 3, 1], k = 4
Output: [9, 8, 6]
Explanation: The window size is 4 and the maximum at different iterations are as follows:
max(9, 8, 6, 4) = 9
max(8, 6, 4, 3) = 8
max(6, 4, 3, 1) = 6
Hence, we get arr = [9, 8, 6] as output.
Example 3

Input: arr[] = [1, 2, 3, 4, 10, 6, 9, 8, 7, 5], k = 3
Output: [3, 4, 10, 10, 10, 9, 9, 8]
Explanation: The window size is 3 and the maximum at different iterations are as follows:
max(1, 2, 3) = 3
max(2, 3, 4) = 4
max(3, 4, 10) = 10
max(4, 10, 6) = 10
max(10, 6, 9) = 10
max(6, 9, 8) = 9
max(9, 8, 7) = 9
max(8, 7, 5) = 8
Hence, we get arr = [3, 4, 10, 10, 10, 9, 9, 8] as output.


We can use a double-ended queue to keep only the indices of those elements which are useful.
The use of deque is to add and drop elements from both the ends of a queue.
We will slide the window of K elements by “dropping” the first element and “adding”
the next element after the window to move it forward.

The deque will keep the index of the maximum element at the front and also at a time,
it will delete all the unnecessary elements from the window.
You can look at the solution steps for more explanation

1. Create a dequeue to store elements.
2. Iterate through the array, insert the first K elements in the array.
   While insertion we will take care of the window such that there are no unnecessary indices.
   To remove these indices, we will remove all the elements from the back of the queue
   that is smaller than the current array element.
3. After the iteration for the first K element, the maximum element's index is at the front of the queue.
4. Now, Iterate through the remaining part of the array and remove the element from the front
   if they are out of the current window.
5. Again, insert the element in the dequeue and before inserting delete those unnecessary indices
   which are smaller than the current array element.
6. Now, after each iteration, you will get the maximum element of the current window.
"""

from collections import deque


def get_max_on_windows(k: int, a: list[int]) -> list[int]:
    d = deque[int]()
    r = list[int]()
    # for First K elements
    for i in range(k):
        while len(d) > 0 and a[i] < a[d[len(d)-1]]:
            # Remove the indices of elements that are smaller than the current elements
            d.pop()
        d.append(i)
        print("i: ", i, ", d: ", d)
    # the element at the front has index of the highest element in the window
    r.append(a[d[0]])
    print("r: ", r)
    # for rest elements
    for i in range(k, len(a)-1):
        # drop the elements that are out of window
        while len(d) > 0 and d[0] > i-k:
            d.popleft()
        # remove those elements smaller than the current element from back
        while len(d) > 0 and a[i] < a[d[len(d)-1]]:
            d.pop()
        d.append(i)
        r.append(a[d[0]])
    return r


def main():
    get_max_on_windows(3, [3, 4, 1, 9, 4, 7, 4, 5, 6, 4])
    pass


if __name__ == "__main__":
    main()
    pass