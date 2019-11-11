package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

type Paragraph struct {
	Index    int
	Group    int
	Text     string
	Genre    string
	GenreIdx string
}

var grupoJornalistico = map[int]int{
	//1: 1, 2: 1, 3: 6, 4: 3, 5: 4, 6: 5, 8: 3, 9: 6, 10: 1, 11: 0, 12: 2, 13: 5, 14: 5, 15: 1, 16: 0, 17: 0, 18: 1, 19: 2, 20: 1, 21: 4, 22: 1, 23: 4, 24: 5, 25: 1, 26: 5, 28: 5, 29: 4, 39: 1, 40: 5, 41: 0, 43: 1, 47: 4, 61: 1, 62: 1, 63: 1, 64: 1, 65: 1, 66: 3, 67: 6, 69: 5, 70: 5, 71: 5, 72: 0,
	// 1: 0, 2: 0, 4: 4, 5: 5, 6: 3, 8: 1, 9: 4, 10: 0, 12: 6, 13: 3, 14: 3, 15: 0, 16: 6, 17: 6, 18: 0, 19: 2, 20: 0, 21: 5, 22: 0, 23: 5, 24: 3, 25: 0, 26: 3, 28: 3, 29: 5, 40: 0, 41: 6, 43: 0, 47: 5, 61: 0, 63: 0, 65: 0, 66: 1, 67: 4, 69: 3, 72: 3,
	//61: 3, 62: 1, 63: 1, 64: 0, 65: 5, 66: 0, 67: 1, 68: 1, 69: 5, 70: 5, 71: 3, 72: 1, 73: 5, 74: 5, 75: 1, 76: 2, 77: 1, 78: 0, 79: 3, 80: 0, 81: 3, 82: 1, 83: 3, 84: 3, 85: 4, 86: 1, 87: 5, 88: 1, 89: 4, 90: 1, 91: 1, 92: 1, 93: 0, 94: 1, 95: 5, 96: 3, 97: 1, 98: 4, 99: 4, 100: 1, 101: 1, 102: 1, 103: 1, 104: 3, 105: 5, 106: 5, 107: 3, 108: 3, 109: 3, 110: 2, 111: 3, 112: 4, 113: 3, 114: 3, 115: 1, 116: 1, 117: 2, 118: 4, 119: 1, 120: 0, 121: 3,
	61: 1, 62: 3, 63: 5, 64: 7, 65: 1, 66: 7, 67: 5, 68: 5, 69: 4, 70: 4, 71: 4, 72: 1, 73: 4, 74: 4, 75: 3, 76: 6, 77: 1, 78: 6, 79: 1, 80: 2, 81: 1, 82: 1, 83: 1, 84: 4, 85: 7, 86: 5, 87: 4, 88: 3, 89: 6, 90: 3, 91: 3, 92: 3, 93: 2, 94: 5, 95: 1, 96: 3, 97: 3, 98: 2, 99: 7, 100: 5, 101: 3, 102: 5, 103: 5, 104: 1, 105: 2, 106: 4, 107: 4, 108: 1, 109: 1, 110: 0, 111: 1, 112: 4, 113: 4, 114: 1, 115: 1, 116: 3, 117: 0, 118: 6, 119: 5, 120: 6, 121: 1,
}

var grupoLiterario = map[int]int{
	//	30: 3, 31: 1, 32: 2, 33: 3, 34: 2, 35: 0, 45: 1, 51: 1, 68: 1,
	// 30: 1, 31: 2, 33: 0, 34: 0, 35: 3, 45: 1, 51: 2, 68: 2,
	//122: 4, 123: 3, 124: 4, 125: 1, 126: 1, 127: 3, 128: 3, 129: 3, 130: 0, 131: 0, 132: 0, 133: 0, 134: 3, 135: 1, 136: 4, 137: 4, 138: 3, 139: 3, 140: 4, 141: 0, 142: 1, 143: 1, 144: 1, 145: 1, 146: 2, 147: 4, 148: 3, 149: 0, 150: 0, 151: 3,
	122: 6, 123: 1, 124: 3, 125: 3, 126: 2, 127: 1, 128: 3, 129: 1, 130: 4, 131: 4, 132: 4, 133: 6, 134: 1, 135: 2, 136: 3, 137: 2, 138: 1, 139: 1, 140: 5, 141: 4, 142: 2, 143: 5, 144: 2, 145: 0, 146: 0, 147: 3, 148: 1, 149: 4, 150: 4, 151: 3,
}

var grupoDivulgacao = map[int]int{
	//	7: 1, 27: 3, 36: 6, 37: 3, 38: 6, 42: 0, 44: 3, 46: 0, 48: 6, 49: 0, 50: 0, 52: 3, 53: 3, 54: 6, 55: 6, 56: 6, 57: 0, 58: 3, 59: 3, 60: 3, 73: 6, 74: 0, 75: 0, 76: 4, 77: 1, 78: 1, 79: 1, 80: 1, 81: 2, 82: 0, 83: 1, 84: 1, 85: 4, 86: 1, 87: 1, 88: 1, 89: 2, 90: 2, 91: 1, 92: 5, 93: 1, 94: 2, 95: 5, 96: 1, 97: 1, 98: 5, 99: 0, 100: 2,
	//	7: 2, 27: 7, 36: 5, 37: 2, 38: 5, 42: 0, 44: 2, 46: 0, 48: 5, 49: 0, 50: 0, 52: 7, 53: 7, 54: 5, 55: 5, 56: 5, 57: 0, 58: 7, 59: 7, 60: 2, 73: 7, 74: 0, 75: 0, 76: 4, 77: 3, 78: 3, 79: 0, 80: 3, 81: 0, 82: 3, 83: 3, 84: 3, 85: 2, 86: 0, 87: 3, 88: 3, 89: 1, 90: 3, 91: 3, 92: 7, 93: 3, 94: 1, 95: 3, 96: 3, 97: 0, 98: 6, 99: 0, 100: 3,
	//	7: 0, 27: 0, 36: 3, 37: 7, 38: 3, 42: 2, 44: 7, 46: 2, 48: 3, 49: 0, 50: 3, 52: 0, 53: 0, 54: 3, 55: 3, 56: 3, 57: 2, 58: 0, 59: 0, 60: 7, 73: 0, 74: 2, 75: 2, 76: 6, 77: 2, 78: 2, 79: 0, 80: 2, 81: 2, 82: 2, 83: 2, 84: 2, 85: 7, 86: 2, 87: 2, 88: 2, 89: 5, 90: 2, 91: 2, 92: 0, 93: 7, 94: 5, 95: 2, 96: 2, 97: 2, 98: 5, 99: 3, 100: 2, 101: 7, 102: 0, 103: 3, 104: 3, 105: 1, 106: 0, 107: 0, 108: 0, 109: 0, 110: 0, 111: 0, 112: 1, 113: 1, 114: 3, 115: 4, 116: 0, 117: 0, 118: 4, 119: 3, 120: 0, 121: 0, 122: 3, 123: 1, 124: 0, 125: 3, 126: 0, 127: 4, 128: 0, 129: 4, 130: 0, 131: 0,
	//101: 7, 102: 0, 103: 5, 104: 5, 105: 6, 106: 7, 107: 7, 108: 0, 109: 7, 110: 7, 111: 0, 112: 1, 113: 1, 114: 5, 115: 2, 116: 7, 117: 0, 118: 2, 119: 5, 120: 4, 121: 7, 122: 5, 123: 1, 124: 0, 125: 5, 126: 4, 127: 2, 128: 7, 129: 3, 130: 0, 131: 0,
	// 7: 1, 27: 5, 37: 4, 46: 4, 52: 4, 53: 5, 54: 4, 55: 4, 57: 1, 59: 1, 78: 0, 91: 1, 101: 5, 102: 1, 103: 4, 104: 4, 105: 2, 106: 1, 107: 1, 109: 1, 110: 1, 111: 1, 114: 2, 115: 6, 116: 1, 117: 1, 118: 6, 119: 4, 120: 5, 122: 1, 123: 2, 124: 1, 126: 4, 127: 6, 128: 1, 129: 3, 130: 1, 131: 1,
	//	1: 6, 2: 4, 3: 6, 4: 6, 5: 6, 6: 4, 7: 6, 8: 6, 9: 0, 10: 4, 11: 2, 12: 4, 13: 4, 14: 0, 15: 6, 16: 6, 17: 1, 18: 4, 19: 4, 20: 0, 21: 0, 22: 0, 23: 1, 24: 5, 25: 0, 26: 0, 27: 5, 28: 6, 29: 4, 30: 0, 31: 1, 32: 0, 33: 6, 34: 5, 35: 4, 36: 3, 37: 0, 38: 0, 39: 4, 40: 2, 41: 2, 42: 6, 43: 0, 44: 4, 45: 6, 46: 2, 47: 4, 48: 0, 49: 2, 50: 4, 51: 2, 52: 0, 53: 4, 54: 4, 55: 0, 56: 0, 57: 6, 58: 4, 59: 0, 60: 4,
	1: 4, 2: 6, 3: 4, 4: 8, 5: 8, 6: 5, 7: 8, 8: 4, 9: 4, 10: 6, 11: 1, 12: 5, 13: 5, 14: 8, 15: 4, 16: 4, 17: 1, 18: 0, 19: 6, 20: 0, 21: 3, 22: 0, 23: 7, 24: 2, 25: 6, 26: 8, 27: 7, 28: 4, 29: 0, 30: 0, 31: 2, 32: 8, 33: 4, 34: 1, 35: 6, 36: 1, 37: 8, 38: 8, 39: 0, 40: 3, 41: 7, 42: 0, 43: 8, 44: 6, 45: 4, 46: 3, 47: 5, 48: 0, 49: 3, 50: 5, 51: 7, 52: 0, 53: 6, 54: 6, 55: 0, 56: 0, 57: 4, 58: 0, 59: 0, 60: 6,
}

type ParagraphOrder []Paragraph

func (a ParagraphOrder) Len() int      { return len(a) }
func (a ParagraphOrder) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ParagraphOrder) Less(i, j int) bool {
	if a[i].Genre == a[j].Genre {
		return a[i].Group < a[j].Group
	} else {
		return a[i].Genre < a[j].Genre
	}
}

var paragraphList = []Paragraph{}

func main() {

	log.Println("starting")

	raw := readFile("/home/sidleal/sid/usp/rastros/maisum/rastros151.txt")
	lines := strings.Split(raw, "\n")

	for i, line := range lines {
		if line == "" {
			break
		}
		idx := i + 1

		group := -1
		genre := ""
		genreIdx := ""
		if _, found := grupoJornalistico[idx]; found {
			group = grupoJornalistico[idx]
			genre = "Jornalístico"
			genreIdx = "JN"
		} else if _, found := grupoLiterario[idx]; found {
			group = grupoLiterario[idx]
			genre = "Literário"
			genreIdx = "LT"
		} else if _, found := grupoDivulgacao[idx]; found {
			group = grupoDivulgacao[idx]
			genre = "Divulgação Científica"
			genreIdx = "DC"
		}
		paragraphList = append(paragraphList, Paragraph{idx, group, line, genre, genreIdx})
	}

	sort.Sort(ParagraphOrder(paragraphList))

	lastGroup := -1
	for _, p := range paragraphList {
		if lastGroup != p.Group {
			log.Println("\n=============================================", "Gênero:", p.Genre, "Grupo", p.Group, "=============================================")
		} else {
			log.Println("---------------------")
		}
		log.Println(p.Index, fmt.Sprintf("%v-%v", p.GenreIdx, p.Group), p.Text)
		lastGroup = p.Group

	}

}