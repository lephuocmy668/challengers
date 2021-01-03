#!/bin/python3
# https://www.hackerrank.com/challenges/dijkstrashortreach/problem
import math
import os
import random
import re
import sys
from collections import defaultdict
from heapq import heappop, heappush

class Graph:
    def __init__(self):
        self.nodes = set()
        self.edges = defaultdict(list)
        self.distances = defaultdict()
    
    def add_node(self, node):
        self.nodes.add(node)
    
    def add_edge(self, start, end, distance):
        self.nodes.add(start)
        self.nodes.add(end)
        self.edges[start].append((end, distance))
        self.edges[end].append((start, distance))        

# Complete the shortestReach function below.
def shortestReach(n, edges, s):    
    graph = Graph()
    for edge in edges:
        graph.add_edge(edge[0], edge[1], edge[2])
    
    visited = {s: 0}   
    path = {}
    nodes = graph.nodes.copy()
    while(nodes):
        min_node = None
        for node in nodes:
            if node in visited:
                if min_node is None:
                    min_node = node
                elif visited[node] < visited[min_node]:
                    min_node = node
        if min_node is None:
            break
    
        nodes.remove(min_node)
        current_weight = visited[min_node]
        
        for edge, distance in graph.edges[min_node]:
            weight = current_weight + distance
            if edge not in visited or weight < visited[edge]:
                visited[edge] = weight
                path[edge] = min_node
    
    result = []
    for node in [n + 1 for n in range(n)]:
        if node == s:
            continue
        if node in visited:
            result.append(visited[node])
        else:
            result.append(-1)
    return result

def optimizeShortestReach(n, edges, start):
    graph = Graph()
    for edge in edges:
        graph.add_edge(edge[0], edge[1], edge[2])

    visited = [False for _ in range(n + 1)]
    path = {}
    distances = [float("inf") for _ in range(n + 1)]
    distances[start] = 0
    min_heap = [(distances[start], start)]
    
    while min_heap:
        current_distance, min_node = heappop(min_heap)
        if visited[min_node]:
            continue
        visited[min_node] = True
        for edge, weight in graph.edges[min_node]:
            distance = current_distance + weight
            if not visited[edge] and distances[edge] > distance:
                path[edge] = min_node
                distances[edge] = distance
                heappush(min_heap, (distances[edge], edge))
    del distances[start]
    del distances[0]
    return [-1 if d == float("inf") else d for d in distances]

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    t = int(input())

    for t_itr in range(t):
        nm = input().split()

        n = int(nm[0])

        m = int(nm[1])

        edges = []

        for _ in range(m):
            edges.append(list(map(int, input().rstrip().split())))

        s = int(input())

        result = shortestReach(n, edges, s)

        fptr.write(' '.join(map(str, result)))
        fptr.write('\n')

    fptr.close()
