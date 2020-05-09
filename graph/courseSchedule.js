// https://leetcode.com/problems/course-schedule/submissions/
// Runtime: 72 ms, faster than 66.19% of JavaScript online submissions for Course Schedule.
// Memory Usage: 38 MB, less than 100.00% of JavaScript online submissions for Course Schedule.
/**
 * @param {number} numCourses
 * @param {number[][]} prerequisites
 * @return {boolean}
 */
var canFinish = function(numCourses, prerequisites) {
    const visited = {};
    const directionMap = new Map();
    
    for(let i=0; i<prerequisites.length; i++){
        const [target, prerequisite] = prerequisites[i];

        // check course is dependent on itself.
        if(target === prerequisite) {
         return false;
        }
        // check cyclic dependency
        if(visited[prerequisite]) {
            if(directionMap.has(prerequisite) && (directionMap.get(prerequisite)[0] === target)) {
                return false;
            }
        }
        
        if(visited[target]){
            // check cyclic dependency
            if(visited[prerequisite] && directionMap.has(prerequisite)) {
                if(!dfs(directionMap, target, prerequisite)) {
                    return false;
                }
            }
            
            if(!visited[prerequisite]) {
                visited[prerequisite] = true;
            }
            if(!directionMap.has(target)){
                directionMap.set(target, [prerequisite]);
                continue;
            }
            
            const courses = directionMap.get(target);
            courses.push(prerequisite);
            directionMap.set(target, courses);
            continue;
        }
        visited[target] = true;
        visited[prerequisite] = true;
        directionMap.set(target, [prerequisite]);
    }    

    return (numCourses >= Object.keys(visited).length) ? true: false;
};

const dfs = (directionMap, target, prerequisite) => {
    const stack = [];
    stack.push(...directionMap.get(prerequisite));
    while(stack.length){
        const course = stack.pop();
        if(course === target) {
            return false;
        }
        if(directionMap.has(course)){
            stack.push(...directionMap.get(course));
        }
    }
    return true;
}