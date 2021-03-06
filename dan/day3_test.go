package dan

import "testing"

var day3Lines = []string{
	"...#...#....#....##...###....#.",
	"#.#...#...#....#.........#..#..",
	".#....##..#.#..##..##..........",
	".....#.#.............#..#......",
	".......#...#.##.#......#..#.#..",
	"#.#....#......##........##.....",
	".....##.#....#..#...#...##...#.",
	"...#...#..#.......#..#...##...#",
	"..........#...........##.......",
	"..#..#..#...................#..",
	"#..#....#.....##.#.#..........#",
	".#.##.......###.....#.#...#....",
	".#..##....##....#.......#...###",
	"#.#..##...#.#..#...............",
	".........#....#.......##.#.#...",
	"...###...##....##...#..##.#..#.",
	"....#.........#..#...#.......#.",
	"....................#..#.#.#...",
	"..#....#..........#...........#",
	".#.....#..#.....##........##..#",
	"#..##..#...##............#..##.",
	".#..##....#..........#..#.##.#.",
	"..#####..#.#............##.....",
	"...###.#....##..#.#....#.....#.",
	".#.......##....#...#.#.....##..",
	"...#....#...##.#...#..#........",
	".####.....#....#.#.#...#.......",
	"...#....#.....#.......#........",
	"#..#.#.......#...#............#",
	"...#.....###.##....#.#.###.#...",
	".#.........#.......#.#....##...",
	"#.#..#...#.#...##......#..#....",
	".....#...#..#.#...#..###..#....",
	"......#.........#...###........",
	".....#..##...#..........#.....#",
	"..#..#.#.##.#...#....#....##..#",
	"##....#.##...#.##.#..##....#...",
	".....#.#.#.#..#....##.#...#.#..",
	".....##.......#........#.......",
	"...#.#.....#...#...##.#......##",
	"........#..#.#...#.#.....#.#..#",
	"#..##...#.#...##..##...#.#...##",
	".##.#.#..#...#.....#.#.##.#...#",
	".#.####.........##.........#..#",
	".##..............#....#...#...#",
	"......#...#..#...#..#..###.#...",
	".......##...#.#.#..##..#......#",
	".....#....#..##..#.........#...",
	".....#..#.#.#........#.#.####..",
	"#..#.......###....##...........",
	"#..##..........#.#......#.#....",
	".....##........#...#..##.......",
	"###...#.##.#.#.#.#.##...##.....",
	"....#...#........##.#.##..##...",
	".#..#.#.#......#.......##..#..#",
	".#...#.................#....#..",
	".##..#..........#..##.......#..",
	".#.#.#.....#..#.#.........##..#",
	"...#......##...#.......#...##..",
	"##...###....#.###.............#",
	"#.....#.#..#.#..#........#.#.#.",
	".....#.#......##..#.#.....#.##.",
	".......#...........#..#.......#",
	"..#....#.#.#......#.....#...#..",
	".....##........#..##..#..##....",
	"#.#........#...##....#.#..##...",
	"#......#......#....#..#...#.##.",
	"....#.#.......#.#.#............",
	"......####.#.##...#.#.##.....##",
	"..###.#.#..#.........#.####....",
	".#.......#..#.#....#.#..#.#.##.",
	"#....#....#............##...##.",
	"....#....#............#....#..#",
	"..#........#..#....#..#..#...#.",
	".#......##....#..........#....#",
	"#.##.....#..........#.###.#....",
	"....##...#.....#.#......#.##...",
	"#.#.....#.......###.###..#..#.#",
	"..###..##.............#.####.##",
	"#....#.....#....#..##.......#..",
	".....#....#...#.#.#.#..#...#.##",
	"...#.....#..#....###......#.#.#",
	"##.........#.#..#..#.#..#.....#",
	".#.....#.#....#.........##..#.#",
	".#.#..#.###..#..#..........#...",
	".##....#.#.#...#......##.....#.",
	"#.#....#....#...#...##...#..#..",
	"#...#........#....#....#......#",
	"#......#...#..#.#.##.....##..#.",
	"....#...#......##...#..#....#..",
	".#......##.##.......#.......#..",
	".#...#..####...........#.#.#...",
	".........#...#.#.........#.....",
	"#.##.....#.#..#.#.###...###..#.",
	"#...##.###......#.###..##.#.##.",
	"...##.#.....#....#..#......#...",
	"#....###.#..#...##.....#......#",
	"........###...#...#............",
	"........#....#...#...#....#...#",
	"#....#..#..#....#.#........#.#.",
	"##...#.....#.#..........#..#..#",
	"#.#...##.....#........#...#...#",
	"##.#.#.......#...#..#.###....#.",
	".#.......#....##..##...#.....#.",
	"#....#....#.....#.......#......",
	".##.##.##...##...#.#.#..#..#...",
	"#..#..#.##....#......##....###.",
	".......#.#.........#..##.#...##",
	".#..##...#....#.....#..........",
	"..#.#...#......#.#..#..........",
	".##....#.#.#.##.......###...#..",
	"..##.#...#.#.#.#.......#..#....",
	"#..#.......#...#........#.....#",
	".....#.......#......###..#.....",
	"...##.#.......#.....##.....##..",
	"##..#.......#.#.....#....#.....",
	"..#....#.##.##...#...#......#..",
	".#..#.###.#....###........#...#",
	"....##.##...##..#..#.#....#....",
	"..###...##.....##..............",
	"#....#...##...#....#..........#",
	".##........#......##...##...#.#",
	"..#.#.##..........#......#.....",
	"...#...#.........#.##........##",
	"..#.#..#.#..#...#....#...#.....",
	"...##...#..#.###.#..#.#...#....",
	"....###........#..#..##...#....",
	"#.#....##.......#.#........#...",
	".###...#..#.#.#.#..#...#..##.##",
	"..#.........#####.#......#..#..",
	"#.....#.....##..#....#...#...#.",
	"...#..#....##....##.....##.#...",
	".........#............#.##.....",
	"....##.#..#....#.##.......#..##",
	".###....#.#..#......#.#.......#",
	".###...###.#.........#.#..#...#",
	".....#........#..#.#..#.#..##.#",
	".###..#....##.........#..##....",
	"..#.......#..#..##...#.###.#...",
	"#.......#...........#.#...#.###",
	"#.##.##...##.#...##..#.....#...",
	"..#..#........###.#.....##.....",
	"#.....##....#...##...####..#..#",
	"....#........#...#...#.........",
	"......#.#.#.#.......#..#.....##",
	"..#..#....#.....#.#...##......#",
	"..#....#...#.###.........#.###.",
	"...#......##..#.#.....#...#....",
	"...#.......#...#...#........##.",
	"............#...#..#....#.....#",
	"....##......................#..",
	"#.#.#....#....#..........##....",
	"#.#.....#.#.##..#...#.##....##.",
	"...#...#..#...#..#.#.#.......#.",
	"#.....#..........#.........##.#",
	"#...##..#..#.#.......###....#..",
	".#...#..##....#.....##.......#.",
	"....#.##.....#.........#.#....#",
	"........#.#...####..#.......#.#",
	".####...#.#......####.....#.##.",
	"###..#....#..#.......#.#..##..#",
	"#......#.#....##..#.##.#....#.#",
	"...###...#...#..##.#..#..#.#...",
	"...##..##....#..#.....#........",
	".....#..............#......#..#",
	"......#....#......#..#.........",
	"#..#.....#.##...........##.....",
	".#..#.#..................##....",
	"#.#..#..##...#....#.#......#...",
	".##.#.##......#.##...#...#...#.",
	"..#...#.........#.#..#.#....#..",
	".#.####.#..#.#......##.#..#....",
	"#..#.......#....#..............",
	"....#............#..#..........",
	".....#####.....#.....#..##...##",
	"#.#....#.#...............#..##.",
	".#.#..#...#......#.....#.#.#...",
	".#....#.#.#......#.....##....#.",
	"....#....#.##..#.......###...##",
	".....#..#.##...#...#...#..#.#..",
	"##..#........#.#..#..##......#.",
	".#..#..##.......#..#.....#.....",
	".#.#.....###..##.#.#...........",
	"..##..##.####..........#..#....",
	"..##..#..#...#....#......#.#...",
	"#...#.#......##.....##.#..###..",
	"#..#..............#........##.#",
	".........#.##..#.#..#..##.##.#.",
	"#....##....#.#..#.#...##..#....",
	".#....#.......#............##..",
	".......#.#.......#...#.#......#",
	"......##...#.......#.#........#",
	"..###..#.#.##....##...#....##..",
	"..##.##..........##..###.......",
	".#.#.#..#..#.#.......#.#...##..",
	"..#..##.........#.###..#......#",
	"....#.#.#...##.#...#...##..###.",
	"..###..##.........##...#...#..#",
	".#..##...#.......#.......#..#.#",
	"........##....##....#.#.###.#.#",
	"#.....#.#.................#.#..",
	"....#.#.#.....##.####.#......#.",
	"....#.......#.#.##.##..........",
	"...#...........#...#.##...#.###",
	"#....#....#..........#.##......",
	"##..#...........##.....##.##...",
	".#.##...##..##....#..#.....####",
	"#...#...#.##..........##..##...",
	"....##..#....#.....#.#...#....#",
	"..#....#..##...###.#.#.........",
	"#......#.#.#...#...#.........#.",
	"#............###.#.#.#..##...#.",
	".##.....####...##..##..#..##.#.",
	"#..#........#.....#.#.....#...#",
	"#............#....#.#.#........",
	"......##...##.#....#.....#...#.",
	"..#........##......#.#.....##..",
	".#..#..#.....##.......#..#.#..#",
	".#....#..#....##.#.#.#..#..#.##",
	".####.#..........#...#..##.....",
	"...###..###...##..#............",
	"#..#.....##.#...#..##..#.......",
	".....##....#...###.##...#......",
	"...##..#...#..#..##....##....#.",
	"...###....#.###.#.#.##....#....",
	"##.#.#.....#....#.#....#..#....",
	".......##.....#.#..##...##...#.",
	".#....#.#...##.#..#....#.....#.",
	"..#...#..#...#.##........#...#.",
	"#....#......##.#....##...#.#..#",
	".....#..#..#..#......#...#.#.#.",
	"..###....#........#...#.......#",
	"###...#.......#.#.......##.##..",
	"......##.....#.#........#....#.",
	"#.##..#.#.#.#..#....#.##.....#.",
	"..........#.##.#...#...#..#..#.",
	"..#...##.#..........#..##.###..",
	"..###..##.##..#.#...##.####..#.",
	"#.#.#...............##....###.#",
	"....#.........#.#....#.#....#.#",
	"..#...#.###...#....###.....#...",
	"..#..#....#...#............#...",
	".#..#....#..##.....##..........",
	"..#....#.#...#.#.#.#.......##.#",
	".........#....##........#.#....",
	"...#..##.#..#.##...#...#.#....#",
	"....####...#...####.#....###..#",
	"......##...#.##.#.......#..#...",
	"#.#...#.#...#.#...#....#.#.#...",
	".#.....##...#.....###.#....#...",
	"......##.....###...#.#...#.#...",
	"#..#..##.#.#......#....#..#..#.",
	"....#.###.....#..#...#.##.....#",
	"##.##........#......#....#..##.",
	"##.....##.#.....#.....##.....#.",
	".....#.##...#.#..#.#.#.....#...",
	".#.##..#...#.#..#.....#.#......",
	".....##.......#..#...##..#..#..",
	"#.....#..#.####......#........#",
	".#..#..##.#..##............#..#",
	".##..#.#....##.##.....#......#.",
	".......##.........#..#.........",
	".#...#.......................#.",
	"#......#.#....##.#.......#..#..",
	"..##..##......#.......#....#.#.",
	"##......#......##...##.........",
	"..#....####....#.#.....##.#.#..",
	"..........#..#.#.#.....#..#.#..",
	"##..##...........##.......#....",
	"##....#.#....#..#......###....#",
	"...#.#.#..#.......##.......#...",
	"#....#.......#.......#.........",
	"...##......##....#...#......#.#",
	"#......#####.#.........#.....#.",
	"#..#.............#..#....#...#.",
	".......#.##..#..#..#..#....####",
	"......#.##..##..........###...#",
	".#.##....###..#........#....##.",
	"#......#..#...###.#...#.....#..",
	".#.#.......#....##.......#.#...",
	"..#.##..#..##.....#.........#.#",
	"#.#...#..#.##....#.......##....",
	".#.....###....#.#..#...#.....#.",
	"#...#..#.......#.#.....##...#.#",
	"#.#####.........#....##.....#..",
	"#....#..##...#....#.##.......#.",
	".#.#.........##....##....#.....",
	"...#..##.......#....#.#.#......",
	"#.###.##...###....#.....#.####.",
	".#...#.#.#..##.#..........#....",
	"#.#.....#.##.#..####.....##.#..",
	"...###.##..####.......#......##",
	".##..#.........#...#.#.....#.##",
	"..#.....##....###.....#.#...##.",
	"#....#....#..#....#.##.........",
	"......###....#.#..#..#....##...",
	".#.#................#.......##.",
	"...#.......#.........#.#.......",
	"...#..........#...##.....###...",
	"....#......#...#...............",
	".##...#....#.....#.##......#...",
	".#.....###...##..##...#.#......",
	"....##........#.....#...#....#.",
	"#.........#.#...##...#.#..#....",
	"...#.#.....#.#........#.#....#.",
	".#........#.....#.#.#.#.#..#...",
	"....#...#.....#.#....#........#",
	"..###.#....#.#....##...##..#.##",
	".#....#.#.####.#.#.....#.......",
	".#...#...#.................##.#",
	"..................##..#..#.#.#.",
	".#..#............##....###.....",
	".......#....#...........#......",
	"....#.#.#.....###.........#..##",
	"...#.#....#.#.##.#.##.....##..#",
	".#.##.#...##...#.......#.....##",
	".#............#...#..##...#.#.#",
	"#.##..#.##..#..##.###.#........",
	"..............##....#...#..#.#.",
	".#.#...#.#....#....###........#",
	".#....#.#....#......###........",
	"..#.......##......#.##.....#...",
	".....#......#..#...#.#.....#...",
}

func trees(right, down int) int {
	var trees, column int
	for i := 0; i < len(day3Lines); i += down {
		line := day3Lines[i]
		if column > len(line)-1 {
			column -= len(line)
		}
		if line[column:column+1] == "#" {
			trees++
		}
		column += right
	}
	return trees
}

func TestDay3a(t *testing.T) {
	t.Logf("Trees: %d", trees(3, 1))
}

func TestDay3b(t *testing.T) {
	trees := []int{
		trees(1, 1),
		trees(3, 1),
		trees(5, 1),
		trees(7, 1),
		trees(1, 2),
	}
	total := 1
	for _, v := range trees {
		total *= v
	}
	t.Logf("Total: %d", total)
}
