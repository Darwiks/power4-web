package power4

func resetGrille(grille interface{}) interface{} {
	switch grille.(type) {
	case [6][7]int:
		return [6][7]int{}
	case [6][9]int:
		return [6][9]int{}
	case [7][8]int:
		return [7][8]int{}
	}
	return grille
}

func placerPion(grille interface{}, col int, joueur *int) interface{} {
	switch g := grille.(type) {
	case [6][7]int:
		for i := 5; i >= 0; i-- {
			if g[i][col] == 0 {
				g[i][col] = *joueur
				*joueur = 3 - *joueur
				break
			}
		}
		return g
	case [6][9]int:
		for i := 5; i >= 0; i-- {
			if g[i][col] == 0 {
				g[i][col] = *joueur
				*joueur = 3 - *joueur
				break
			}
		}
		return g
	case [7][8]int:
		for i := 6; i >= 0; i-- {
			if g[i][col] == 0 {
				g[i][col] = *joueur
				*joueur = 3 - *joueur
				break
			}
		}
		return g
	}
	return grille
}

func checkVictoire(grille interface{}, joueur int) bool {
	switch g := grille.(type) {
	case [6][7]int:
		s := make([][]int, 6)
		for i := 0; i < 6; i++ {
			s[i] = make([]int, 7)
			for j := 0; j < 7; j++ {
				s[i][j] = g[i][j]
			}
		}
		return checkVictoireTable(s, joueur)
	case [6][9]int:
		s := make([][]int, 6)
		for i := 0; i < 6; i++ {
			s[i] = make([]int, 9)
			for j := 0; j < 9; j++ {
				s[i][j] = g[i][j]
			}
		}
		return checkVictoireTable(s, joueur)
	case [7][8]int:
		s := make([][]int, 7)
		for i := 0; i < 7; i++ {
			s[i] = make([]int, 8)
			for j := 0; j < 8; j++ {
				s[i][j] = g[i][j]
			}
		}
		return checkVictoireTable(s, joueur)
	}
	return false
}

func checkVictoireTable(g [][]int, joueur int) bool {
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[0]); j++ {
			if g[i][j] != joueur {
				continue
			}
			if j+3 < len(g[0]) && g[i][j+1] == joueur && g[i][j+2] == joueur && g[i][j+3] == joueur {
				return true
			}
			if i+3 < len(g) && g[i+1][j] == joueur && g[i+2][j] == joueur && g[i+3][j] == joueur {
				return true
			}
			if i+3 < len(g) && j+3 < len(g[0]) && g[i+1][j+1] == joueur && g[i+2][j+2] == joueur && g[i+3][j+3] == joueur {
				return true
			}
			if i+3 < len(g) && j-3 >= 0 && g[i+1][j-1] == joueur && g[i+2][j-2] == joueur && g[i+3][j-3] == joueur {
				return true
			}
		}
	}
	return false
}
