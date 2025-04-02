# binary-tree

The small library that does not want to use recursion to compare binary trees.

## Biased results

The traditional recursive "Equals":
```
Equals (10      nodes): 299ns
Equals (100     nodes): 941ns
Equals (1000    nodes): 12.836µs
Equals (10000   nodes): 161.821µs
Equals (100000  nodes): 1.481126ms
Equals (1000000 nodes): 14.844285ms
```

The non recursive "FastEquals":
```
FastEquals (10      nodes): 185ns
FastEquals (100     nodes): 191ns
FastEquals (1000    nodes): 584ns
FastEquals (10000   nodes): 1.256µs
FastEquals (100000  nodes): 4.819µs
FastEquals (1000000 nodes): 17.75µs
```
