package qns4_1

import "testing"

func TestRouteBetweenNodes(t *testing.T) {
	type args struct {
		graph Graph
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{[][]int{{2, 4, 6}, {4, 5}, {0, 4, 3}, {2, 8}, {0, 1, 2}, {1, 4, 6}, {0, 5}, {}, {3}, {}}, 8, 6}, true},
		{"", args{[][]int{{2, 4, 6}, {4, 5}, {0, 4, 3}, {2, 8}, {0, 1, 2}, {1, 4, 6}, {0, 5}, {}, {3}, {}}, 0, 8}, true},
		{"", args{[][]int{{2, 4, 6}, {4, 5}, {0, 4, 3}, {2, 8}, {0, 1, 2}, {1, 4, 6}, {0, 5}, {}, {3}, {}}, 8, 9}, false},
		{"", args{[][]int{{2, 4, 6}, {4, 5}, {0, 4, 3}, {2, 8}, {0, 1, 2}, {1, 4, 6}, {0, 5}, {}, {3}, {}}, 8, 7}, false},
		{"", args{[][]int{{9}, {5}, {3}, {2, 8}, {}, {1, 6}, {5, 7}, {6}, {3}, {0}}, 8, 7}, false},
		{"", args{[][]int{{9}, {5}, {3}, {2, 8}, {}, {1, 6}, {5, 7}, {6}, {3}, {0}}, 7, 8}, false},
		{"", args{[][]int{{9}, {5}, {3}, {2, 8}, {}, {1, 6}, {5, 7}, {6}, {3}, {0}}, 0, 8}, false},
		{"", args{[][]int{{9}, {5}, {3}, {2, 8}, {}, {1, 6}, {5, 7}, {6}, {3}, {0}}, 0, 1}, false},
		{"", args{[][]int{{9}, {5}, {3}, {2, 8}, {}, {1, 6}, {5, 7}, {6}, {3}, {0}}, 2, 8}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RouteBetweenNodes(tt.args.graph, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("RouteBetweenNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
