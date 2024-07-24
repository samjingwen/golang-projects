package qns4_7

import "errors"

type Pair struct {
	first, second int
}

type dependencies map[int]struct{}

var void = struct{}{}

func BuildOrder(projects []int, depList []Pair) ([]int, error) {
	// map project to its dependencies
	adjList := make(map[int]dependencies)
	for _, project := range projects {
		adjList[project] = make(dependencies)
	}
	for _, pair := range depList {
		deps := adjList[pair.second]
		deps[pair.first] = void
	}

	var res []int
	var isCycle bool

	visited := make(map[int]struct{})
	visiting := make(map[int]struct{})

	var dfs func(project int)
	dfs = func(project int) {
		if _, exists := visiting[project]; exists {
			isCycle = true
			return
		}
		visiting[project] = void

		deps := adjList[project]
		for dep := range deps {
			if _, exists := visited[dep]; exists {
				delete(deps, dep)
				continue
			}
			dfs(dep)
			delete(deps, dep)
		}
		if _, exists := visited[project]; !exists {
			res = append(res, project)
			visited[project] = void
		}
		delete(visiting, project)
	}

	for _, project := range projects {
		dfs(project)
	}
	if isCycle {
		return nil, errors.New("cycle detected")
	}
	return res, nil
}
