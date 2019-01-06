# Distributed-Networks
A project to compute the existence of a fault tolerant path between the source and destination

## Motivation: 
Earlier only 1 fault in a 2D mesh could be tolerated and upto 2 faults in a 3D mesh.
We realized that it is not just these nodes that are faulty and hence can't be in the path between source and destination, but also those that are rendered useless and unreachable due to these faulty nodes.

![cnn](https://github.com/nikki30/Distributed-Networks/blob/master/img/1.png)
![cnn](https://github.com/nikki30/Distributed-Networks/blob/master/img/2.png)

## APC Method

![cnn](https://github.com/nikki30/Distributed-Networks/blob/master/img/3.png)

## FAPC and RAPC Method

![cnn](https://github.com/nikki30/Distributed-Networks/blob/master/img/4.png)

## Example - Mesh network with faulty nodes

![cnn](https://github.com/nikki30/Distributed-Networks/blob/master/img/5.png)

## FAPC values

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




