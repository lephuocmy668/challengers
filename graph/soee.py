# https://leetcode.com/problems/satisfiability-of-equality-equations
# Runtime: 28 ms, faster than 100.00% of Python3 online submissions for Satisfiability of Equality Equations.
# Memory Usage: 14.6 MB, less than 19.05% of Python3 online submissions for Satisfiability of Equality Equations.
class Solution:
    def equationsPossible(self, equations: List[str]) -> bool:
        # init graph
        graph = collections.defaultdict(list)
        for equation in equations:
            if "==" in equation:
                v1, v2 = equation[0], equation[3]
                graph[v1].append(v2)
                graph[v2].append(v1)

        # graph coloring
        vertex_colors = {val: None for val in "abcdefghijklmnopqrstuvwxyz"}
        color = 0
        for start_node in graph.keys():
            if vertex_colors[start_node] is None:
                color += 1
                stack = [start_node]
                while stack:
                    node = stack.pop()
                    for nei in graph[node]:
                        if vertex_colors[nei] is None:
                            vertex_colors[nei] = color
                            stack.append(nei)

        # check inequal
        for equation in equations:
            if "!=" in equation:
                v1, v2 = equation[0], equation[3]
                if v1 == v2:
                    return False
                if vertex_colors[v1] and vertex_colors[v2]:
                    if vertex_colors[v1] == vertex_colors[v2]:
                        return False
                    
        return True