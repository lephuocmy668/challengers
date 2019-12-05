function getNode(val) {
    return {
        val, next: null, prev: null
    }
}

function LinkList() {
    this.head = getNode(0)
    this.tail = getNode(0)
    this.head.next = this.tail
    this.tail.prev = this.head
}

LinkList.prototype.add = function (node) {
    const prev = this.tail.prev
    const next = this.tail

    prev.next = node
    node.prev = prev

    node.next = next
    next.prev = node
}

LinkList.prototype.delete = function (node) {
    const prev = node.prev
    const next = node.next

    prev.next = next
    next.prev = prev
}

LinkList.prototype.values = function () {
    let next = this.head.next
    const res = []
    while (next !== this.tail) {
        res.push(next.val)
        next = next.next
    }
    return res
}
var topKFrequent = function (nums, k) {
    const freq = [new LinkList()]
    const pos = {}

    for (let i of nums) {
        if (i in pos) {
            const node = pos[i].node
            let freqPos = pos[i].freqPos
            freq[freqPos].delete(node)
            pos[i].freqPos++
            freqPos = pos[i].freqPos
            if (freqPos >= freq.length) {
                freq.push(new LinkList())
            }
            freq[freqPos].add(node)
        } else {
            const node = getNode(i)
            pos[i] = { node, freqPos: 0 }
            freq[0].add(node)
        }
    }


    const res = [];
    while (res.length < k && freq.length > 0) {
        const ls = freq.pop()
        const tmp = ls.values()
        for (let i of tmp) {
            res.push(i)
            if (res.length === k) {
                return res
            }
        }
    }

    return res
};

// use a double link list to store the node of same frequences
// use a hash map to locate the node by its value
// use a list to store those double link list
// move a node from one double link list to another, once it occurs again
// pop k most frequent nodes