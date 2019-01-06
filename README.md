# Distributed-Networks
Heuristic to compute the existence of a fault tolerant path between the source and destination

## Motivation: 
It is not just faulty nodes in the path between a source and destination that hinder communication, it is also the nodes near the faulty nodes such as unreachable nodes (can't be reached from source due to a faulty node) and the useless nodes(can't forward to the destination due to a faulty node)

## Related works:

![cnn](https://github.com/nikki30/Distributed-Networks/blob/master/img/1.png)
![cnn](https://github.com/nikki30/Distributed-Networks/blob/master/img/2.png)

## Allowed Path Counter Method (the base paper)

Assign value to each node starting with 1 for a non-faulty source. Each node then gets a value that is the addition of the value on its left and bottom (only considering XY routing), which represents the number of ways to get to that node from the source. 
![cnn](https://github.com/nikki30/Distributed-Networks/blob/master/img/3.png)

## Forward-APC and Reverse-APC Method

![cnn](https://github.com/nikki30/Distributed-Networks/blob/master/img/4.png)

## Example - Consider Mesh network with faulty nodes

![cnn](https://github.com/nikki30/Distributed-Networks/blob/master/img/5.png)

## The corresponding FAPC values for the example

![cnn](https://github.com/nikki30/Distributed-Networks/blob/master/img/6.png)

## Our algorithm
1.N clusters are formed with sources S1, S2, ...Sn

2.S1 sends (w, h) to S2, S3, ...Sn where n is the number of clusters, w is the width of the subRSD, h is the height of the subRSD

3.Each Si gets a list of the faulty nodes within its subRSD

4.Each Si runs the APC algorithms in parallel 

5.We keep merging subRSDs pairwise until we get the FAPC values for the entire RSD

## Splitting the mesh network

![cnn](https://github.com/nikki30/Distributed-Networks/blob/master/img/7.png)

![cnn](https://github.com/nikki30/Distributed-Networks/blob/master/img/8.png)

## Comparison

![cnn](https://github.com/nikki30/Distributed-Networks/blob/master/img/9.png)

## Results:

![cnn](https://github.com/nikki30/Distributed-Networks/blob/master/img/10.png)

## Computation to merge subRSD’s FAPC values

![cnn](https://github.com/nikki30/Distributed-Networks/blob/master/img/11.png)

## FAPC values after the merge

![cnn](https://github.com/nikki30/Distributed-Networks/blob/master/img/12.png)

## Final Merge

![cnn](https://github.com/nikki30/Distributed-Networks/blob/master/img/13.png)

## Heuristic Approach

![cnn](https://github.com/nikki30/Distributed-Networks/blob/master/img/14.png)

## Results

![cnn](https://github.com/nikki30/Distributed-Networks/blob/master/img/15.png)




