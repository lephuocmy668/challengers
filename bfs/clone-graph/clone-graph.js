// https://leetcode.com/problems/clone-graph/submissions/
// Runtime: 64 ms, faster than 81.09% of JavaScript online submissions for Clone Graph.
// Memory Usage: 35.6 MB, less than 100.00% of JavaScript online submissions for Clone Graph.
/**
 * // Definition for a Node.
 * function Node(val, neighbors) {
 *    this.val = val === undefined ? 0 : val;
 *    this.neighbors = neighbors === undefined ? [] : neighbors;
 * };
 */
/**
 * @param {Node} node
 * @return {Node}
 */
var cloneGraph = function (node) {
    if (!node) return node;

    // map[vistedNode]clonedNode
    const vistedNodes = new Map();
    const queue = [node];

    vistedNodes.set(node, new Node(node.val));

    while (queue.length > 0) {
        const currentNode = queue.pop();
        const clonedNode = vistedNodes.get(currentNode);

        currentNode.neighbors.forEach(neighbor => {
            if (!vistedNodes.has(neighbor)) {
                vistedNodes.set(neighbor, new Node(neighbor.val));
                queue.unshift(neighbor);
            }
            const clonedNeighbor = vistedNodes.get(neighbor);
            clonedNode.neighbors.push(clonedNeighbor);
        })
    }

    return vistedNodes.get(node);
};