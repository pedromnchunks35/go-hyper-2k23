# Sorting Algorithms
###  1. Bubble Sort [Link](BubbleSort/bubblesort.go)
- It is meant to swap 2 elements that are in the wrong order
- Time complexity is O(n^2)
###   2. Selection Sort [Link](SelectionSort/selectionSort.go)
- It is used to stort the elements by repeatedly finding the minimum element and placing it at the first in the unsorted part of the array
- It is the contrast of the bubble, it puts the minimum member at the begginining instead of putting the biggest to the right
- Time complexity is O(n^2)
###   3. Insertion Sort [Link](InsertionSort/insertionSort.go)
- Also timecomplexity of O(n^2)
###   4. Merge Sort [Link](MergeSort/mergeSort.go)
- It uses recursion to divide the array, making it smaller and smaller comparing after only the first members of the array
- Bad side is that it uses to much pointers to memory
- It is O(n*logn)
###   5. Quick Sort [Link](QuickSort/quickSort.go)
- Is one of the most efficient, it divides the array in multiple parts and by changing the temp array of positions is also changing the main array
- It is O(n*logn)
# Search Algorithms
###   1. Linear Search [Link](LinearSearch/linearSearch.go)
- This is the normal Search
- O(n)
###   2. Binary Search [Link](BinarySearch/binarySearch.go)
- It only works with ordered lists
- It dvides the list until finding the value which will be between the low and high
# Graph Algorithms
###  1. Depth-First Search (DFS) [Link](DepthFirstSearch/dfs.go)
- This is a graph search
- It uses recursive methods to iterate the array of dests associated with the source
- It will dive in to the leafs first and not into the pair (this is the difference against the Breadth)
- O(V+E), V is number of Vertices and E the number of edges
###  2. Breadth-First Search (BFS) [Link](BreathFirstSearch/bfs.go)
- In this algorithm you create a queue, to place in orderer which elements you will iter. By itering a node you place his childs inside of the key right ahead, so by order you will check them one by one
###  3. Dijkstr's Algorithm [Link](DijkstraAlgorithm/dijks.go)
- It was conceived to solve the minimum time problem between two places
- Its representation is a table where we place the shortest distance to the point we want to discover the shortest time and the respective previous point in relation to the point we current are that leads to the point we wanted to discover the shortest path
- By steps it is something like this
  ```
  (ALL POINTS START WITH MINIMUM DISTANCE AS INFINITE)
  Loop (UNTIL ALL THE POINTS ARE VISITED)
    GET EDGES FROM THE POINT
    MARK THAT CURRENT LOCATION AS VISITED
    GO CHECK INTO THE TABLE IF THE MINIMUM COST
  Loop
  ```
###  4. Minimum Spanning Tree (MST) [Link](Prims/prims.go)
- This is a algorithm to find the shortest path that visits every point
- There are 2 algorithms for this, but the most well known is prims
- Prims firstly starts in a arbitary point
- Selects everytime the smallest cost node
- When the current nodes does not have more nodes, it will then go to the not yet visited nodes
- After that it continues with the same logic until every node gets visited
# Dynamic Programming
###  1. Memoization [Link](Memoization/memo.go)
- Creating a cache for a given resource in order to not having to make the same solution once again
###  2. Recursion [Link](Memozation/memo.go)
- Recursion is using the result as a input of the same function
- We did recursion in memoization as well.. so you can check that program as well
# Tree Algorithms 
###  1. Binary Trees [LINK](BinaryTrees/binaryT.go)
- This is the secound time we implement the binary tree algorithm, you can check more of that in the first lesson exercises
###  2. Binary Search Trees [Link](BinaryTrees/binaryT.go)
- Searching all over the leafs using recursive algorithm
###  3. AVL Trees [LINK](AVLTrees/AVLT.go)
- It is a algorithm to by using recursion becoming able to balance the trees
- This algorithm makes the search over the tree more fast after it
- It shares 4 scenarios where you need to rotate the tree: Left Dig, Left->Right, Right Dig, Right left
- We should Rotate according to the balanceFactor. Balance factor is BF = L - R
- Also denote that every insertion we also make a loop back to update the heigth of every node
- Check own impl to see the comments
###  4. Red-Black Trees
###  5. Pre-orderer
###  6. In-order
###  6. Post-order
# Hashing
###  1. Hash Functions
###  2. Collision Resolution(Separate Chaining and Open Addressing)
###  3. SHA-1
###  4. SHA-256
###  5. MD5
# Greedy Algorithms
###  1. Activity Selection Problem
###  2. Knapsack problem
# String Algorithms
###  1. Naive
###  2. Rabin-karp
###  3. Knuth-Morris-Pratt
###  4. Radix Sort
###  5. Run-Length Encoding