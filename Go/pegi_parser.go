// Code generated from ./PegiParser.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // PegiParser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 57, 390,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 3, 2, 3, 2, 3, 2, 5, 2, 96, 10, 2, 3, 2, 3,
	2, 3, 2, 7, 2, 101, 10, 2, 12, 2, 14, 2, 104, 11, 2, 3, 2, 5, 2, 107, 10,
	2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7,
	4, 120, 10, 4, 12, 4, 14, 4, 123, 11, 4, 3, 4, 5, 4, 126, 10, 4, 3, 5,
	5, 5, 129, 10, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8,
	6, 8, 140, 10, 8, 13, 8, 14, 8, 141, 3, 9, 3, 9, 3, 9, 5, 9, 147, 10, 9,
	3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 153, 10, 10, 3, 11, 3, 11, 3, 11, 7,
	11, 158, 10, 11, 12, 11, 14, 11, 161, 11, 11, 3, 12, 3, 12, 3, 12, 7, 12,
	166, 10, 12, 12, 12, 14, 12, 169, 11, 12, 3, 13, 3, 13, 6, 13, 173, 10,
	13, 13, 13, 14, 13, 174, 3, 14, 3, 14, 5, 14, 179, 10, 14, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 5, 17, 192,
	10, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 7, 17, 209, 10, 17, 12, 17, 14,
	17, 212, 11, 17, 3, 18, 3, 18, 5, 18, 216, 10, 18, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 5, 19, 224, 10, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3,
	21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 236, 10, 22, 5, 22, 238,
	10, 22, 3, 22, 5, 22, 241, 10, 22, 3, 22, 5, 22, 244, 10, 22, 5, 22, 246,
	10, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23,
	256, 10, 23, 3, 24, 3, 24, 3, 24, 5, 24, 261, 10, 24, 3, 25, 3, 25, 3,
	25, 5, 25, 266, 10, 25, 3, 26, 3, 26, 5, 26, 270, 10, 26, 3, 27, 3, 27,
	3, 27, 3, 27, 7, 27, 276, 10, 27, 12, 27, 14, 27, 279, 11, 27, 3, 27, 3,
	27, 3, 28, 3, 28, 3, 28, 7, 28, 286, 10, 28, 12, 28, 14, 28, 289, 11, 28,
	3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 5, 30, 299, 10,
	30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 5, 33, 309,
	10, 33, 3, 34, 3, 34, 5, 34, 313, 10, 34, 3, 34, 3, 34, 7, 34, 317, 10,
	34, 12, 34, 14, 34, 320, 11, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 6,
	35, 327, 10, 35, 13, 35, 14, 35, 328, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36,
	3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 5, 38, 342, 10, 38, 3, 39, 3,
	39, 3, 39, 3, 39, 3, 40, 3, 40, 5, 40, 350, 10, 40, 3, 41, 3, 41, 3, 41,
	7, 41, 355, 10, 41, 12, 41, 14, 41, 358, 11, 41, 3, 42, 3, 42, 3, 43, 3,
	43, 3, 43, 3, 43, 3, 43, 5, 43, 367, 10, 43, 3, 43, 3, 43, 7, 43, 371,
	10, 43, 12, 43, 14, 43, 374, 11, 43, 3, 43, 3, 43, 3, 44, 3, 44, 3, 44,
	3, 44, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 46, 5, 46, 388, 10, 46, 3,
	46, 2, 3, 32, 47, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
	32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66,
	68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 2, 9, 4, 2, 7, 7, 18, 18,
	3, 2, 52, 53, 4, 2, 32, 32, 38, 38, 4, 2, 34, 34, 36, 36, 3, 2, 25, 30,
	4, 2, 33, 33, 35, 35, 4, 2, 39, 42, 46, 46, 2, 396, 2, 95, 3, 2, 2, 2,
	4, 110, 3, 2, 2, 2, 6, 113, 3, 2, 2, 2, 8, 128, 3, 2, 2, 2, 10, 132, 3,
	2, 2, 2, 12, 134, 3, 2, 2, 2, 14, 139, 3, 2, 2, 2, 16, 146, 3, 2, 2, 2,
	18, 148, 3, 2, 2, 2, 20, 154, 3, 2, 2, 2, 22, 162, 3, 2, 2, 2, 24, 170,
	3, 2, 2, 2, 26, 178, 3, 2, 2, 2, 28, 180, 3, 2, 2, 2, 30, 184, 3, 2, 2,
	2, 32, 191, 3, 2, 2, 2, 34, 215, 3, 2, 2, 2, 36, 223, 3, 2, 2, 2, 38, 225,
	3, 2, 2, 2, 40, 227, 3, 2, 2, 2, 42, 230, 3, 2, 2, 2, 44, 255, 3, 2, 2,
	2, 46, 257, 3, 2, 2, 2, 48, 265, 3, 2, 2, 2, 50, 269, 3, 2, 2, 2, 52, 271,
	3, 2, 2, 2, 54, 282, 3, 2, 2, 2, 56, 292, 3, 2, 2, 2, 58, 298, 3, 2, 2,
	2, 60, 300, 3, 2, 2, 2, 62, 304, 3, 2, 2, 2, 64, 308, 3, 2, 2, 2, 66, 310,
	3, 2, 2, 2, 68, 323, 3, 2, 2, 2, 70, 332, 3, 2, 2, 2, 72, 336, 3, 2, 2,
	2, 74, 341, 3, 2, 2, 2, 76, 343, 3, 2, 2, 2, 78, 349, 3, 2, 2, 2, 80, 351,
	3, 2, 2, 2, 82, 359, 3, 2, 2, 2, 84, 361, 3, 2, 2, 2, 86, 377, 3, 2, 2,
	2, 88, 381, 3, 2, 2, 2, 90, 387, 3, 2, 2, 2, 92, 93, 5, 4, 3, 2, 93, 94,
	5, 90, 46, 2, 94, 96, 3, 2, 2, 2, 95, 92, 3, 2, 2, 2, 95, 96, 3, 2, 2,
	2, 96, 102, 3, 2, 2, 2, 97, 98, 5, 6, 4, 2, 98, 99, 5, 90, 46, 2, 99, 101,
	3, 2, 2, 2, 100, 97, 3, 2, 2, 2, 101, 104, 3, 2, 2, 2, 102, 100, 3, 2,
	2, 2, 102, 103, 3, 2, 2, 2, 103, 106, 3, 2, 2, 2, 104, 102, 3, 2, 2, 2,
	105, 107, 5, 14, 8, 2, 106, 105, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107,
	108, 3, 2, 2, 2, 108, 109, 7, 2, 2, 3, 109, 3, 3, 2, 2, 2, 110, 111, 7,
	5, 2, 2, 111, 112, 5, 12, 7, 2, 112, 5, 3, 2, 2, 2, 113, 125, 7, 6, 2,
	2, 114, 126, 5, 8, 5, 2, 115, 121, 7, 8, 2, 2, 116, 117, 5, 8, 5, 2, 117,
	118, 5, 90, 46, 2, 118, 120, 3, 2, 2, 2, 119, 116, 3, 2, 2, 2, 120, 123,
	3, 2, 2, 2, 121, 119, 3, 2, 2, 2, 121, 122, 3, 2, 2, 2, 122, 124, 3, 2,
	2, 2, 123, 121, 3, 2, 2, 2, 124, 126, 7, 9, 2, 2, 125, 114, 3, 2, 2, 2,
	125, 115, 3, 2, 2, 2, 126, 7, 3, 2, 2, 2, 127, 129, 9, 2, 2, 2, 128, 127,
	3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 131, 5, 10,
	6, 2, 131, 9, 3, 2, 2, 2, 132, 133, 5, 12, 7, 2, 133, 11, 3, 2, 2, 2, 134,
	135, 9, 3, 2, 2, 135, 13, 3, 2, 2, 2, 136, 137, 5, 16, 9, 2, 137, 138,
	5, 90, 46, 2, 138, 140, 3, 2, 2, 2, 139, 136, 3, 2, 2, 2, 140, 141, 3,
	2, 2, 2, 141, 139, 3, 2, 2, 2, 141, 142, 3, 2, 2, 2, 142, 15, 3, 2, 2,
	2, 143, 147, 5, 18, 10, 2, 144, 147, 5, 40, 21, 2, 145, 147, 5, 84, 43,
	2, 146, 143, 3, 2, 2, 2, 146, 144, 3, 2, 2, 2, 146, 145, 3, 2, 2, 2, 147,
	17, 3, 2, 2, 2, 148, 149, 5, 22, 12, 2, 149, 152, 7, 21, 2, 2, 150, 153,
	5, 20, 11, 2, 151, 153, 5, 24, 13, 2, 152, 150, 3, 2, 2, 2, 152, 151, 3,
	2, 2, 2, 153, 19, 3, 2, 2, 2, 154, 159, 5, 32, 17, 2, 155, 156, 7, 15,
	2, 2, 156, 158, 5, 32, 17, 2, 157, 155, 3, 2, 2, 2, 158, 161, 3, 2, 2,
	2, 159, 157, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160, 21, 3, 2, 2, 2, 161,
	159, 3, 2, 2, 2, 162, 167, 7, 7, 2, 2, 163, 164, 7, 15, 2, 2, 164, 166,
	7, 7, 2, 2, 165, 163, 3, 2, 2, 2, 166, 169, 3, 2, 2, 2, 167, 165, 3, 2,
	2, 2, 167, 168, 3, 2, 2, 2, 168, 23, 3, 2, 2, 2, 169, 167, 3, 2, 2, 2,
	170, 172, 5, 26, 14, 2, 171, 173, 5, 26, 14, 2, 172, 171, 3, 2, 2, 2, 173,
	174, 3, 2, 2, 2, 174, 172, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 25, 3,
	2, 2, 2, 176, 179, 5, 28, 15, 2, 177, 179, 5, 30, 16, 2, 178, 176, 3, 2,
	2, 2, 178, 177, 3, 2, 2, 2, 179, 27, 3, 2, 2, 2, 180, 181, 5, 90, 46, 2,
	181, 182, 7, 31, 2, 2, 182, 183, 5, 22, 12, 2, 183, 29, 3, 2, 2, 2, 184,
	185, 5, 90, 46, 2, 185, 186, 7, 31, 2, 2, 186, 187, 5, 32, 17, 2, 187,
	31, 3, 2, 2, 2, 188, 189, 8, 17, 1, 2, 189, 192, 5, 34, 18, 2, 190, 192,
	5, 48, 25, 2, 191, 188, 3, 2, 2, 2, 191, 190, 3, 2, 2, 2, 192, 210, 3,
	2, 2, 2, 193, 194, 12, 7, 2, 2, 194, 195, 9, 4, 2, 2, 195, 209, 5, 32,
	17, 8, 196, 197, 12, 6, 2, 2, 197, 198, 9, 5, 2, 2, 198, 209, 5, 32, 17,
	7, 199, 200, 12, 5, 2, 2, 200, 201, 9, 6, 2, 2, 201, 209, 5, 32, 17, 6,
	202, 203, 12, 4, 2, 2, 203, 204, 7, 24, 2, 2, 204, 209, 5, 32, 17, 5, 205,
	206, 12, 3, 2, 2, 206, 207, 7, 23, 2, 2, 207, 209, 5, 32, 17, 4, 208, 193,
	3, 2, 2, 2, 208, 196, 3, 2, 2, 2, 208, 199, 3, 2, 2, 2, 208, 202, 3, 2,
	2, 2, 208, 205, 3, 2, 2, 2, 209, 212, 3, 2, 2, 2, 210, 208, 3, 2, 2, 2,
	210, 211, 3, 2, 2, 2, 211, 33, 3, 2, 2, 2, 212, 210, 3, 2, 2, 2, 213, 216,
	5, 36, 19, 2, 214, 216, 5, 40, 21, 2, 215, 213, 3, 2, 2, 2, 215, 214, 3,
	2, 2, 2, 216, 35, 3, 2, 2, 2, 217, 224, 5, 38, 20, 2, 218, 224, 5, 46,
	24, 2, 219, 220, 7, 8, 2, 2, 220, 221, 5, 32, 17, 2, 221, 222, 7, 9, 2,
	2, 222, 224, 3, 2, 2, 2, 223, 217, 3, 2, 2, 2, 223, 218, 3, 2, 2, 2, 223,
	219, 3, 2, 2, 2, 224, 37, 3, 2, 2, 2, 225, 226, 5, 44, 23, 2, 226, 39,
	3, 2, 2, 2, 227, 228, 5, 80, 41, 2, 228, 229, 5, 42, 22, 2, 229, 41, 3,
	2, 2, 2, 230, 245, 7, 8, 2, 2, 231, 238, 5, 20, 11, 2, 232, 235, 5, 78,
	40, 2, 233, 234, 7, 15, 2, 2, 234, 236, 5, 20, 11, 2, 235, 233, 3, 2, 2,
	2, 235, 236, 3, 2, 2, 2, 236, 238, 3, 2, 2, 2, 237, 231, 3, 2, 2, 2, 237,
	232, 3, 2, 2, 2, 238, 240, 3, 2, 2, 2, 239, 241, 7, 22, 2, 2, 240, 239,
	3, 2, 2, 2, 240, 241, 3, 2, 2, 2, 241, 243, 3, 2, 2, 2, 242, 244, 7, 15,
	2, 2, 243, 242, 3, 2, 2, 2, 243, 244, 3, 2, 2, 2, 244, 246, 3, 2, 2, 2,
	245, 237, 3, 2, 2, 2, 245, 246, 3, 2, 2, 2, 246, 247, 3, 2, 2, 2, 247,
	248, 7, 9, 2, 2, 248, 43, 3, 2, 2, 2, 249, 256, 5, 88, 45, 2, 250, 256,
	5, 58, 30, 2, 251, 256, 5, 50, 26, 2, 252, 256, 5, 12, 7, 2, 253, 256,
	7, 43, 2, 2, 254, 256, 7, 46, 2, 2, 255, 249, 3, 2, 2, 2, 255, 250, 3,
	2, 2, 2, 255, 251, 3, 2, 2, 2, 255, 252, 3, 2, 2, 2, 255, 253, 3, 2, 2,
	2, 255, 254, 3, 2, 2, 2, 256, 45, 3, 2, 2, 2, 257, 260, 7, 7, 2, 2, 258,
	259, 7, 18, 2, 2, 259, 261, 7, 7, 2, 2, 260, 258, 3, 2, 2, 2, 260, 261,
	3, 2, 2, 2, 261, 47, 3, 2, 2, 2, 262, 266, 5, 34, 18, 2, 263, 264, 9, 7,
	2, 2, 264, 266, 5, 32, 17, 2, 265, 262, 3, 2, 2, 2, 265, 263, 3, 2, 2,
	2, 266, 49, 3, 2, 2, 2, 267, 270, 5, 52, 27, 2, 268, 270, 5, 54, 28, 2,
	269, 267, 3, 2, 2, 2, 269, 268, 3, 2, 2, 2, 270, 51, 3, 2, 2, 2, 271, 272,
	7, 12, 2, 2, 272, 277, 5, 32, 17, 2, 273, 274, 7, 15, 2, 2, 274, 276, 5,
	32, 17, 2, 275, 273, 3, 2, 2, 2, 276, 279, 3, 2, 2, 2, 277, 275, 3, 2,
	2, 2, 277, 278, 3, 2, 2, 2, 278, 280, 3, 2, 2, 2, 279, 277, 3, 2, 2, 2,
	280, 281, 7, 13, 2, 2, 281, 53, 3, 2, 2, 2, 282, 283, 7, 12, 2, 2, 283,
	287, 5, 90, 46, 2, 284, 286, 5, 56, 29, 2, 285, 284, 3, 2, 2, 2, 286, 289,
	3, 2, 2, 2, 287, 285, 3, 2, 2, 2, 287, 288, 3, 2, 2, 2, 288, 290, 3, 2,
	2, 2, 289, 287, 3, 2, 2, 2, 290, 291, 7, 13, 2, 2, 291, 55, 3, 2, 2, 2,
	292, 293, 5, 32, 17, 2, 293, 294, 7, 15, 2, 2, 294, 295, 5, 90, 46, 2,
	295, 57, 3, 2, 2, 2, 296, 299, 5, 66, 34, 2, 297, 299, 5, 68, 35, 2, 298,
	296, 3, 2, 2, 2, 298, 297, 3, 2, 2, 2, 299, 59, 3, 2, 2, 2, 300, 301, 5,
	72, 37, 2, 301, 302, 7, 17, 2, 2, 302, 303, 5, 74, 38, 2, 303, 61, 3, 2,
	2, 2, 304, 305, 7, 7, 2, 2, 305, 63, 3, 2, 2, 2, 306, 309, 5, 62, 32, 2,
	307, 309, 5, 60, 31, 2, 308, 306, 3, 2, 2, 2, 308, 307, 3, 2, 2, 2, 309,
	65, 3, 2, 2, 2, 310, 312, 7, 10, 2, 2, 311, 313, 5, 64, 33, 2, 312, 311,
	3, 2, 2, 2, 312, 313, 3, 2, 2, 2, 313, 318, 3, 2, 2, 2, 314, 315, 7, 15,
	2, 2, 315, 317, 5, 64, 33, 2, 316, 314, 3, 2, 2, 2, 317, 320, 3, 2, 2,
	2, 318, 316, 3, 2, 2, 2, 318, 319, 3, 2, 2, 2, 319, 321, 3, 2, 2, 2, 320,
	318, 3, 2, 2, 2, 321, 322, 7, 11, 2, 2, 322, 67, 3, 2, 2, 2, 323, 324,
	7, 10, 2, 2, 324, 326, 5, 90, 46, 2, 325, 327, 5, 70, 36, 2, 326, 325,
	3, 2, 2, 2, 327, 328, 3, 2, 2, 2, 328, 326, 3, 2, 2, 2, 328, 329, 3, 2,
	2, 2, 329, 330, 3, 2, 2, 2, 330, 331, 7, 11, 2, 2, 331, 69, 3, 2, 2, 2,
	332, 333, 5, 64, 33, 2, 333, 334, 7, 15, 2, 2, 334, 335, 5, 90, 46, 2,
	335, 71, 3, 2, 2, 2, 336, 337, 7, 7, 2, 2, 337, 73, 3, 2, 2, 2, 338, 342,
	5, 32, 17, 2, 339, 342, 5, 58, 30, 2, 340, 342, 5, 50, 26, 2, 341, 338,
	3, 2, 2, 2, 341, 339, 3, 2, 2, 2, 341, 340, 3, 2, 2, 2, 342, 75, 3, 2,
	2, 2, 343, 344, 7, 12, 2, 2, 344, 345, 7, 13, 2, 2, 345, 346, 5, 78, 40,
	2, 346, 77, 3, 2, 2, 2, 347, 350, 5, 80, 41, 2, 348, 350, 5, 82, 42, 2,
	349, 347, 3, 2, 2, 2, 349, 348, 3, 2, 2, 2, 350, 79, 3, 2, 2, 2, 351, 356,
	7, 7, 2, 2, 352, 353, 7, 18, 2, 2, 353, 355, 7, 7, 2, 2, 354, 352, 3, 2,
	2, 2, 355, 358, 3, 2, 2, 2, 356, 354, 3, 2, 2, 2, 356, 357, 3, 2, 2, 2,
	357, 81, 3, 2, 2, 2, 358, 356, 3, 2, 2, 2, 359, 360, 5, 84, 43, 2, 360,
	83, 3, 2, 2, 2, 361, 362, 7, 3, 2, 2, 362, 363, 7, 7, 2, 2, 363, 372, 7,
	10, 2, 2, 364, 367, 5, 86, 44, 2, 365, 367, 5, 80, 41, 2, 366, 364, 3,
	2, 2, 2, 366, 365, 3, 2, 2, 2, 367, 368, 3, 2, 2, 2, 368, 369, 5, 90, 46,
	2, 369, 371, 3, 2, 2, 2, 370, 366, 3, 2, 2, 2, 371, 374, 3, 2, 2, 2, 372,
	370, 3, 2, 2, 2, 372, 373, 3, 2, 2, 2, 373, 375, 3, 2, 2, 2, 374, 372,
	3, 2, 2, 2, 375, 376, 7, 11, 2, 2, 376, 85, 3, 2, 2, 2, 377, 378, 6, 44,
	7, 2, 378, 379, 7, 7, 2, 2, 379, 380, 5, 78, 40, 2, 380, 87, 3, 2, 2, 2,
	381, 382, 9, 8, 2, 2, 382, 89, 3, 2, 2, 2, 383, 388, 7, 16, 2, 2, 384,
	388, 7, 2, 2, 3, 385, 388, 6, 46, 8, 2, 386, 388, 6, 46, 9, 2, 387, 383,
	3, 2, 2, 2, 387, 384, 3, 2, 2, 2, 387, 385, 3, 2, 2, 2, 387, 386, 3, 2,
	2, 2, 388, 91, 3, 2, 2, 2, 42, 95, 102, 106, 121, 125, 128, 141, 146, 152,
	159, 167, 174, 178, 191, 208, 210, 215, 223, 235, 237, 240, 243, 245, 255,
	260, 265, 269, 277, 287, 298, 308, 312, 318, 328, 341, 349, 356, 366, 372,
	387,
}
var literalNames = []string{
	"", "'interface'", "'struct'", "'parser'", "'import'", "", "'('", "')'",
	"'{'", "'}'", "'['", "']'", "'='", "','", "';'", "':'", "'.'", "'++'",
	"'--'", "':='", "'...'", "'||'", "'&&'", "'=='", "'!='", "'<'", "'<='",
	"'>'", "'>='", "'|'", "' / '", "'!'", "' + '", "' -'", "' - '", "'-'",
	"' * '",
}
var symbolicNames = []string{
	"", "INTERFACE", "STRUCT", "PARSER", "IMPORT", "IDENTIFIER", "L_PAREN",
	"R_PAREN", "L_CURLY", "R_CURLY", "L_BRACKET", "R_BRACKET", "ASSIGN", "COMMA",
	"SEMI", "COLON", "DOT", "PLUS_PLUS", "MINUS_MINUS", "DECLARE_ASSIGN", "ELLIPSIS",
	"LOGICAL_OR", "LOGICAL_AND", "EQUALS", "NOT_EQUALS", "LESS", "LESS_OR_EQUALS",
	"GREATER", "GREATER_OR_EQUALS", "OR", "DIV", "EXCLAMATION", "PLUS", "UNARY_MINUS",
	"MINUS", "MINUS_IN_NAME", "STAR", "DECIMAL_LIT", "BINARY_LIT", "OCTAL_LIT",
	"HEX_LIT", "FLOAT_LIT", "DECIMAL_FLOAT_LIT", "HEX_FLOAT_LIT", "RUNE_LIT",
	"BYTE_VALUE", "OCTAL_BYTE_VALUE", "HEX_BYTE_VALUE", "LITTLE_U_VALUE", "BIG_U_VALUE",
	"RAW_STRING_LIT", "INTERPRETED_STRING_LIT", "WS", "COMMENT", "TERMINATOR",
	"LINE_COMMENT",
}

var ruleNames = []string{
	"sourceFile", "parserClause", "importDecl", "importSpec", "importPath",
	"string_", "statementList", "statement", "shortVarDecl", "expressionList",
	"identifierList", "pipe", "pipeUnit", "pipeUnitIdentifers", "pipeUnitExpression",
	"expression", "primaryExpr", "operand", "literal", "functionCall", "arguments",
	"basicLit", "operandName", "unaryExpr", "sliceValue", "sliceValueSingleLine",
	"sliceValueMultiLine", "sliceValueMultiLineLine", "messageValue", "keyedElementKV",
	"keyedElementIdentifer", "keyedElement", "messageValueSingleLine", "messageValueMultiLine",
	"lineKeyedElement", "key", "element", "sliceType", "typeT", "typeName",
	"typeLit", "interfaceType", "fieldDecl", "integer", "eos",
}

type PegiParser struct {
	PegiParserBase
}

// NewPegiParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *PegiParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewPegiParser(input antlr.TokenStream) *PegiParser {
	this := new(PegiParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "PegiParser.g4"

	return this
}

// PegiParser tokens.
const (
	PegiParserEOF                    = antlr.TokenEOF
	PegiParserINTERFACE              = 1
	PegiParserSTRUCT                 = 2
	PegiParserPARSER                 = 3
	PegiParserIMPORT                 = 4
	PegiParserIDENTIFIER             = 5
	PegiParserL_PAREN                = 6
	PegiParserR_PAREN                = 7
	PegiParserL_CURLY                = 8
	PegiParserR_CURLY                = 9
	PegiParserL_BRACKET              = 10
	PegiParserR_BRACKET              = 11
	PegiParserASSIGN                 = 12
	PegiParserCOMMA                  = 13
	PegiParserSEMI                   = 14
	PegiParserCOLON                  = 15
	PegiParserDOT                    = 16
	PegiParserPLUS_PLUS              = 17
	PegiParserMINUS_MINUS            = 18
	PegiParserDECLARE_ASSIGN         = 19
	PegiParserELLIPSIS               = 20
	PegiParserLOGICAL_OR             = 21
	PegiParserLOGICAL_AND            = 22
	PegiParserEQUALS                 = 23
	PegiParserNOT_EQUALS             = 24
	PegiParserLESS                   = 25
	PegiParserLESS_OR_EQUALS         = 26
	PegiParserGREATER                = 27
	PegiParserGREATER_OR_EQUALS      = 28
	PegiParserOR                     = 29
	PegiParserDIV                    = 30
	PegiParserEXCLAMATION            = 31
	PegiParserPLUS                   = 32
	PegiParserUNARY_MINUS            = 33
	PegiParserMINUS                  = 34
	PegiParserMINUS_IN_NAME          = 35
	PegiParserSTAR                   = 36
	PegiParserDECIMAL_LIT            = 37
	PegiParserBINARY_LIT             = 38
	PegiParserOCTAL_LIT              = 39
	PegiParserHEX_LIT                = 40
	PegiParserFLOAT_LIT              = 41
	PegiParserDECIMAL_FLOAT_LIT      = 42
	PegiParserHEX_FLOAT_LIT          = 43
	PegiParserRUNE_LIT               = 44
	PegiParserBYTE_VALUE             = 45
	PegiParserOCTAL_BYTE_VALUE       = 46
	PegiParserHEX_BYTE_VALUE         = 47
	PegiParserLITTLE_U_VALUE         = 48
	PegiParserBIG_U_VALUE            = 49
	PegiParserRAW_STRING_LIT         = 50
	PegiParserINTERPRETED_STRING_LIT = 51
	PegiParserWS                     = 52
	PegiParserCOMMENT                = 53
	PegiParserTERMINATOR             = 54
	PegiParserLINE_COMMENT           = 55
)

// PegiParser rules.
const (
	PegiParserRULE_sourceFile              = 0
	PegiParserRULE_parserClause            = 1
	PegiParserRULE_importDecl              = 2
	PegiParserRULE_importSpec              = 3
	PegiParserRULE_importPath              = 4
	PegiParserRULE_string_                 = 5
	PegiParserRULE_statementList           = 6
	PegiParserRULE_statement               = 7
	PegiParserRULE_shortVarDecl            = 8
	PegiParserRULE_expressionList          = 9
	PegiParserRULE_identifierList          = 10
	PegiParserRULE_pipe                    = 11
	PegiParserRULE_pipeUnit                = 12
	PegiParserRULE_pipeUnitIdentifers      = 13
	PegiParserRULE_pipeUnitExpression      = 14
	PegiParserRULE_expression              = 15
	PegiParserRULE_primaryExpr             = 16
	PegiParserRULE_operand                 = 17
	PegiParserRULE_literal                 = 18
	PegiParserRULE_functionCall            = 19
	PegiParserRULE_arguments               = 20
	PegiParserRULE_basicLit                = 21
	PegiParserRULE_operandName             = 22
	PegiParserRULE_unaryExpr               = 23
	PegiParserRULE_sliceValue              = 24
	PegiParserRULE_sliceValueSingleLine    = 25
	PegiParserRULE_sliceValueMultiLine     = 26
	PegiParserRULE_sliceValueMultiLineLine = 27
	PegiParserRULE_messageValue            = 28
	PegiParserRULE_keyedElementKV          = 29
	PegiParserRULE_keyedElementIdentifer   = 30
	PegiParserRULE_keyedElement            = 31
	PegiParserRULE_messageValueSingleLine  = 32
	PegiParserRULE_messageValueMultiLine   = 33
	PegiParserRULE_lineKeyedElement        = 34
	PegiParserRULE_key                     = 35
	PegiParserRULE_element                 = 36
	PegiParserRULE_sliceType               = 37
	PegiParserRULE_typeT                   = 38
	PegiParserRULE_typeName                = 39
	PegiParserRULE_typeLit                 = 40
	PegiParserRULE_interfaceType           = 41
	PegiParserRULE_fieldDecl               = 42
	PegiParserRULE_integer                 = 43
	PegiParserRULE_eos                     = 44
)

// ISourceFileContext is an interface to support dynamic dispatch.
type ISourceFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourceFileContext differentiates from other interfaces.
	IsSourceFileContext()
}

type SourceFileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceFileContext() *SourceFileContext {
	var p = new(SourceFileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_sourceFile
	return p
}

func (*SourceFileContext) IsSourceFileContext() {}

func NewSourceFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceFileContext {
	var p = new(SourceFileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_sourceFile

	return p
}

func (s *SourceFileContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceFileContext) EOF() antlr.TerminalNode {
	return s.GetToken(PegiParserEOF, 0)
}

func (s *SourceFileContext) ParserClause() IParserClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParserClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParserClauseContext)
}

func (s *SourceFileContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *SourceFileContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *SourceFileContext) AllImportDecl() []IImportDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImportDeclContext)(nil)).Elem())
	var tst = make([]IImportDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImportDeclContext)
		}
	}

	return tst
}

func (s *SourceFileContext) ImportDecl(i int) IImportDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImportDeclContext)
}

func (s *SourceFileContext) StatementList() IStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *SourceFileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceFileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceFileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterSourceFile(s)
	}
}

func (s *SourceFileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitSourceFile(s)
	}
}

func (s *SourceFileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitSourceFile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) SourceFile() (localctx ISourceFileContext) {
	localctx = NewSourceFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PegiParserRULE_sourceFile)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PegiParserPARSER {
		{
			p.SetState(90)
			p.ParserClause()
		}
		{
			p.SetState(91)
			p.Eos()
		}

	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PegiParserIMPORT {
		{
			p.SetState(95)
			p.ImportDecl()
		}
		{
			p.SetState(96)
			p.Eos()
		}

		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PegiParserINTERFACE || _la == PegiParserIDENTIFIER {
		{
			p.SetState(103)
			p.StatementList()
		}

	}
	{
		p.SetState(106)
		p.Match(PegiParserEOF)
	}

	return localctx
}

// IParserClauseContext is an interface to support dynamic dispatch.
type IParserClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetParserPath returns the parserPath rule contexts.
	GetParserPath() IString_Context

	// SetParserPath sets the parserPath rule contexts.
	SetParserPath(IString_Context)

	// IsParserClauseContext differentiates from other interfaces.
	IsParserClauseContext()
}

type ParserClauseContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	parserPath IString_Context
}

func NewEmptyParserClauseContext() *ParserClauseContext {
	var p = new(ParserClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_parserClause
	return p
}

func (*ParserClauseContext) IsParserClauseContext() {}

func NewParserClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParserClauseContext {
	var p = new(ParserClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_parserClause

	return p
}

func (s *ParserClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParserClauseContext) GetParserPath() IString_Context { return s.parserPath }

func (s *ParserClauseContext) SetParserPath(v IString_Context) { s.parserPath = v }

func (s *ParserClauseContext) PARSER() antlr.TerminalNode {
	return s.GetToken(PegiParserPARSER, 0)
}

func (s *ParserClauseContext) String_() IString_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_Context)
}

func (s *ParserClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParserClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParserClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterParserClause(s)
	}
}

func (s *ParserClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitParserClause(s)
	}
}

func (s *ParserClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitParserClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) ParserClause() (localctx IParserClauseContext) {
	localctx = NewParserClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PegiParserRULE_parserClause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(PegiParserPARSER)
	}
	{
		p.SetState(109)

		var _x = p.String_()

		localctx.(*ParserClauseContext).parserPath = _x
	}

	return localctx
}

// IImportDeclContext is an interface to support dynamic dispatch.
type IImportDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportDeclContext differentiates from other interfaces.
	IsImportDeclContext()
}

type ImportDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportDeclContext() *ImportDeclContext {
	var p = new(ImportDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_importDecl
	return p
}

func (*ImportDeclContext) IsImportDeclContext() {}

func NewImportDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportDeclContext {
	var p = new(ImportDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_importDecl

	return p
}

func (s *ImportDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportDeclContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(PegiParserIMPORT, 0)
}

func (s *ImportDeclContext) AllImportSpec() []IImportSpecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImportSpecContext)(nil)).Elem())
	var tst = make([]IImportSpecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImportSpecContext)
		}
	}

	return tst
}

func (s *ImportDeclContext) ImportSpec(i int) IImportSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportSpecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImportSpecContext)
}

func (s *ImportDeclContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(PegiParserL_PAREN, 0)
}

func (s *ImportDeclContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(PegiParserR_PAREN, 0)
}

func (s *ImportDeclContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *ImportDeclContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ImportDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterImportDecl(s)
	}
}

func (s *ImportDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitImportDecl(s)
	}
}

func (s *ImportDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitImportDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) ImportDecl() (localctx IImportDeclContext) {
	localctx = NewImportDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PegiParserRULE_importDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Match(PegiParserIMPORT)
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PegiParserIDENTIFIER, PegiParserDOT, PegiParserRAW_STRING_LIT, PegiParserINTERPRETED_STRING_LIT:
		{
			p.SetState(112)
			p.ImportSpec()
		}

	case PegiParserL_PAREN:
		{
			p.SetState(113)
			p.Match(PegiParserL_PAREN)
		}
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == PegiParserIDENTIFIER || _la == PegiParserDOT || _la == PegiParserRAW_STRING_LIT || _la == PegiParserINTERPRETED_STRING_LIT {
			{
				p.SetState(114)
				p.ImportSpec()
			}
			{
				p.SetState(115)
				p.Eos()
			}

			p.SetState(121)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(122)
			p.Match(PegiParserR_PAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IImportSpecContext is an interface to support dynamic dispatch.
type IImportSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAlias returns the alias token.
	GetAlias() antlr.Token

	// SetAlias sets the alias token.
	SetAlias(antlr.Token)

	// IsImportSpecContext differentiates from other interfaces.
	IsImportSpecContext()
}

type ImportSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	alias  antlr.Token
}

func NewEmptyImportSpecContext() *ImportSpecContext {
	var p = new(ImportSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_importSpec
	return p
}

func (*ImportSpecContext) IsImportSpecContext() {}

func NewImportSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportSpecContext {
	var p = new(ImportSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_importSpec

	return p
}

func (s *ImportSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportSpecContext) GetAlias() antlr.Token { return s.alias }

func (s *ImportSpecContext) SetAlias(v antlr.Token) { s.alias = v }

func (s *ImportSpecContext) ImportPath() IImportPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImportPathContext)
}

func (s *ImportSpecContext) DOT() antlr.TerminalNode {
	return s.GetToken(PegiParserDOT, 0)
}

func (s *ImportSpecContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PegiParserIDENTIFIER, 0)
}

func (s *ImportSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterImportSpec(s)
	}
}

func (s *ImportSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitImportSpec(s)
	}
}

func (s *ImportSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitImportSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) ImportSpec() (localctx IImportSpecContext) {
	localctx = NewImportSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PegiParserRULE_importSpec)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PegiParserIDENTIFIER || _la == PegiParserDOT {
		{
			p.SetState(125)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ImportSpecContext).alias = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == PegiParserIDENTIFIER || _la == PegiParserDOT) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ImportSpecContext).alias = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(128)
		p.ImportPath()
	}

	return localctx
}

// IImportPathContext is an interface to support dynamic dispatch.
type IImportPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportPathContext differentiates from other interfaces.
	IsImportPathContext()
}

type ImportPathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportPathContext() *ImportPathContext {
	var p = new(ImportPathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_importPath
	return p
}

func (*ImportPathContext) IsImportPathContext() {}

func NewImportPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportPathContext {
	var p = new(ImportPathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_importPath

	return p
}

func (s *ImportPathContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportPathContext) String_() IString_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_Context)
}

func (s *ImportPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportPathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterImportPath(s)
	}
}

func (s *ImportPathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitImportPath(s)
	}
}

func (s *ImportPathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitImportPath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) ImportPath() (localctx IImportPathContext) {
	localctx = NewImportPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PegiParserRULE_importPath)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)
		p.String_()
	}

	return localctx
}

// IString_Context is an interface to support dynamic dispatch.
type IString_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsString_Context differentiates from other interfaces.
	IsString_Context()
}

type String_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_Context() *String_Context {
	var p = new(String_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_string_
	return p
}

func (*String_Context) IsString_Context() {}

func NewString_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_Context {
	var p = new(String_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_string_

	return p
}

func (s *String_Context) GetParser() antlr.Parser { return s.parser }

func (s *String_Context) RAW_STRING_LIT() antlr.TerminalNode {
	return s.GetToken(PegiParserRAW_STRING_LIT, 0)
}

func (s *String_Context) INTERPRETED_STRING_LIT() antlr.TerminalNode {
	return s.GetToken(PegiParserINTERPRETED_STRING_LIT, 0)
}

func (s *String_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterString_(s)
	}
}

func (s *String_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitString_(s)
	}
}

func (s *String_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitString_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) String_() (localctx IString_Context) {
	localctx = NewString_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PegiParserRULE_string_)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PegiParserRAW_STRING_LIT || _la == PegiParserINTERPRETED_STRING_LIT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IStatementListContext is an interface to support dynamic dispatch.
type IStatementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementListContext differentiates from other interfaces.
	IsStatementListContext()
}

type StatementListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementListContext() *StatementListContext {
	var p = new(StatementListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_statementList
	return p
}

func (*StatementListContext) IsStatementListContext() {}

func NewStatementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementListContext {
	var p = new(StatementListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_statementList

	return p
}

func (s *StatementListContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementListContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StatementListContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementListContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *StatementListContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *StatementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterStatementList(s)
	}
}

func (s *StatementListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitStatementList(s)
	}
}

func (s *StatementListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitStatementList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) StatementList() (localctx IStatementListContext) {
	localctx = NewStatementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PegiParserRULE_statementList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PegiParserINTERFACE || _la == PegiParserIDENTIFIER {
		{
			p.SetState(134)
			p.Statement()
		}
		{
			p.SetState(135)
			p.Eos()
		}

		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) ShortVarDecl() IShortVarDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShortVarDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShortVarDeclContext)
}

func (s *StatementContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *StatementContext) InterfaceType() IInterfaceTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInterfaceTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInterfaceTypeContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PegiParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(141)
			p.ShortVarDecl()
		}

	case 2:
		{
			p.SetState(142)
			p.FunctionCall()
		}

	case 3:
		{
			p.SetState(143)
			p.InterfaceType()
		}

	}

	return localctx
}

// IShortVarDeclContext is an interface to support dynamic dispatch.
type IShortVarDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShortVarDeclContext differentiates from other interfaces.
	IsShortVarDeclContext()
}

type ShortVarDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShortVarDeclContext() *ShortVarDeclContext {
	var p = new(ShortVarDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_shortVarDecl
	return p
}

func (*ShortVarDeclContext) IsShortVarDeclContext() {}

func NewShortVarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShortVarDeclContext {
	var p = new(ShortVarDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_shortVarDecl

	return p
}

func (s *ShortVarDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ShortVarDeclContext) IdentifierList() IIdentifierListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *ShortVarDeclContext) DECLARE_ASSIGN() antlr.TerminalNode {
	return s.GetToken(PegiParserDECLARE_ASSIGN, 0)
}

func (s *ShortVarDeclContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ShortVarDeclContext) Pipe() IPipeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPipeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPipeContext)
}

func (s *ShortVarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShortVarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShortVarDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterShortVarDecl(s)
	}
}

func (s *ShortVarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitShortVarDecl(s)
	}
}

func (s *ShortVarDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitShortVarDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) ShortVarDecl() (localctx IShortVarDeclContext) {
	localctx = NewShortVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PegiParserRULE_shortVarDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(146)
		p.IdentifierList()
	}
	{
		p.SetState(147)
		p.Match(PegiParserDECLARE_ASSIGN)
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(148)
			p.ExpressionList()
		}

	case 2:
		{
			p.SetState(149)
			p.Pipe()
		}

	}

	return localctx
}

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_expressionList
	return p
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionListContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PegiParserCOMMA)
}

func (s *ExpressionListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PegiParserCOMMA, i)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterExpressionList(s)
	}
}

func (s *ExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitExpressionList(s)
	}
}

func (s *ExpressionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitExpressionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PegiParserRULE_expressionList)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		p.expression(0)
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(153)
				p.Match(PegiParserCOMMA)
			}
			{
				p.SetState(154)
				p.expression(0)
			}

		}
		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}

	return localctx
}

// IIdentifierListContext is an interface to support dynamic dispatch.
type IIdentifierListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierListContext differentiates from other interfaces.
	IsIdentifierListContext()
}

type IdentifierListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierListContext() *IdentifierListContext {
	var p = new(IdentifierListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_identifierList
	return p
}

func (*IdentifierListContext) IsIdentifierListContext() {}

func NewIdentifierListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierListContext {
	var p = new(IdentifierListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_identifierList

	return p
}

func (s *IdentifierListContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierListContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(PegiParserIDENTIFIER)
}

func (s *IdentifierListContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(PegiParserIDENTIFIER, i)
}

func (s *IdentifierListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PegiParserCOMMA)
}

func (s *IdentifierListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PegiParserCOMMA, i)
}

func (s *IdentifierListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterIdentifierList(s)
	}
}

func (s *IdentifierListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitIdentifierList(s)
	}
}

func (s *IdentifierListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitIdentifierList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) IdentifierList() (localctx IIdentifierListContext) {
	localctx = NewIdentifierListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, PegiParserRULE_identifierList)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.Match(PegiParserIDENTIFIER)
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(161)
				p.Match(PegiParserCOMMA)
			}
			{
				p.SetState(162)
				p.Match(PegiParserIDENTIFIER)
			}

		}
		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}

	return localctx
}

// IPipeContext is an interface to support dynamic dispatch.
type IPipeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPipeContext differentiates from other interfaces.
	IsPipeContext()
}

type PipeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipeContext() *PipeContext {
	var p = new(PipeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_pipe
	return p
}

func (*PipeContext) IsPipeContext() {}

func NewPipeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipeContext {
	var p = new(PipeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_pipe

	return p
}

func (s *PipeContext) GetParser() antlr.Parser { return s.parser }

func (s *PipeContext) AllPipeUnit() []IPipeUnitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPipeUnitContext)(nil)).Elem())
	var tst = make([]IPipeUnitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPipeUnitContext)
		}
	}

	return tst
}

func (s *PipeContext) PipeUnit(i int) IPipeUnitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPipeUnitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPipeUnitContext)
}

func (s *PipeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterPipe(s)
	}
}

func (s *PipeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitPipe(s)
	}
}

func (s *PipeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitPipe(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) Pipe() (localctx IPipeContext) {
	localctx = NewPipeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, PegiParserRULE_pipe)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.PipeUnit()
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(169)
				p.PipeUnit()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}

	return localctx
}

// IPipeUnitContext is an interface to support dynamic dispatch.
type IPipeUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPipeUnitContext differentiates from other interfaces.
	IsPipeUnitContext()
}

type PipeUnitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipeUnitContext() *PipeUnitContext {
	var p = new(PipeUnitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_pipeUnit
	return p
}

func (*PipeUnitContext) IsPipeUnitContext() {}

func NewPipeUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipeUnitContext {
	var p = new(PipeUnitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_pipeUnit

	return p
}

func (s *PipeUnitContext) GetParser() antlr.Parser { return s.parser }

func (s *PipeUnitContext) PipeUnitIdentifers() IPipeUnitIdentifersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPipeUnitIdentifersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPipeUnitIdentifersContext)
}

func (s *PipeUnitContext) PipeUnitExpression() IPipeUnitExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPipeUnitExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPipeUnitExpressionContext)
}

func (s *PipeUnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipeUnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipeUnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterPipeUnit(s)
	}
}

func (s *PipeUnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitPipeUnit(s)
	}
}

func (s *PipeUnitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitPipeUnit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) PipeUnit() (localctx IPipeUnitContext) {
	localctx = NewPipeUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, PegiParserRULE_pipeUnit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(174)
			p.PipeUnitIdentifers()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(175)
			p.PipeUnitExpression()
		}

	}

	return localctx
}

// IPipeUnitIdentifersContext is an interface to support dynamic dispatch.
type IPipeUnitIdentifersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPipeUnitIdentifersContext differentiates from other interfaces.
	IsPipeUnitIdentifersContext()
}

type PipeUnitIdentifersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipeUnitIdentifersContext() *PipeUnitIdentifersContext {
	var p = new(PipeUnitIdentifersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_pipeUnitIdentifers
	return p
}

func (*PipeUnitIdentifersContext) IsPipeUnitIdentifersContext() {}

func NewPipeUnitIdentifersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipeUnitIdentifersContext {
	var p = new(PipeUnitIdentifersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_pipeUnitIdentifers

	return p
}

func (s *PipeUnitIdentifersContext) GetParser() antlr.Parser { return s.parser }

func (s *PipeUnitIdentifersContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *PipeUnitIdentifersContext) OR() antlr.TerminalNode {
	return s.GetToken(PegiParserOR, 0)
}

func (s *PipeUnitIdentifersContext) IdentifierList() IIdentifierListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *PipeUnitIdentifersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipeUnitIdentifersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipeUnitIdentifersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterPipeUnitIdentifers(s)
	}
}

func (s *PipeUnitIdentifersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitPipeUnitIdentifers(s)
	}
}

func (s *PipeUnitIdentifersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitPipeUnitIdentifers(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) PipeUnitIdentifers() (localctx IPipeUnitIdentifersContext) {
	localctx = NewPipeUnitIdentifersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, PegiParserRULE_pipeUnitIdentifers)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(178)
		p.Eos()
	}
	{
		p.SetState(179)
		p.Match(PegiParserOR)
	}
	{
		p.SetState(180)
		p.IdentifierList()
	}

	return localctx
}

// IPipeUnitExpressionContext is an interface to support dynamic dispatch.
type IPipeUnitExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPipeUnitExpressionContext differentiates from other interfaces.
	IsPipeUnitExpressionContext()
}

type PipeUnitExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipeUnitExpressionContext() *PipeUnitExpressionContext {
	var p = new(PipeUnitExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_pipeUnitExpression
	return p
}

func (*PipeUnitExpressionContext) IsPipeUnitExpressionContext() {}

func NewPipeUnitExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipeUnitExpressionContext {
	var p = new(PipeUnitExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_pipeUnitExpression

	return p
}

func (s *PipeUnitExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PipeUnitExpressionContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *PipeUnitExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(PegiParserOR, 0)
}

func (s *PipeUnitExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PipeUnitExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipeUnitExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipeUnitExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterPipeUnitExpression(s)
	}
}

func (s *PipeUnitExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitPipeUnitExpression(s)
	}
}

func (s *PipeUnitExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitPipeUnitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) PipeUnitExpression() (localctx IPipeUnitExpressionContext) {
	localctx = NewPipeUnitExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, PegiParserRULE_pipeUnitExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		p.Eos()
	}
	{
		p.SetState(183)
		p.Match(PegiParserOR)
	}
	{
		p.SetState(184)
		p.expression(0)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMul_op returns the mul_op token.
	GetMul_op() antlr.Token

	// GetAdd_op returns the add_op token.
	GetAdd_op() antlr.Token

	// GetRel_op returns the rel_op token.
	GetRel_op() antlr.Token

	// SetMul_op sets the mul_op token.
	SetMul_op(antlr.Token)

	// SetAdd_op sets the add_op token.
	SetAdd_op(antlr.Token)

	// SetRel_op sets the rel_op token.
	SetRel_op(antlr.Token)

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	mul_op antlr.Token
	add_op antlr.Token
	rel_op antlr.Token
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) GetMul_op() antlr.Token { return s.mul_op }

func (s *ExpressionContext) GetAdd_op() antlr.Token { return s.add_op }

func (s *ExpressionContext) GetRel_op() antlr.Token { return s.rel_op }

func (s *ExpressionContext) SetMul_op(v antlr.Token) { s.mul_op = v }

func (s *ExpressionContext) SetAdd_op(v antlr.Token) { s.add_op = v }

func (s *ExpressionContext) SetRel_op(v antlr.Token) { s.rel_op = v }

func (s *ExpressionContext) PrimaryExpr() IPrimaryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *ExpressionContext) UnaryExpr() IUnaryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryExprContext)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) STAR() antlr.TerminalNode {
	return s.GetToken(PegiParserSTAR, 0)
}

func (s *ExpressionContext) DIV() antlr.TerminalNode {
	return s.GetToken(PegiParserDIV, 0)
}

func (s *ExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(PegiParserPLUS, 0)
}

func (s *ExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(PegiParserMINUS, 0)
}

func (s *ExpressionContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(PegiParserEQUALS, 0)
}

func (s *ExpressionContext) NOT_EQUALS() antlr.TerminalNode {
	return s.GetToken(PegiParserNOT_EQUALS, 0)
}

func (s *ExpressionContext) LESS() antlr.TerminalNode {
	return s.GetToken(PegiParserLESS, 0)
}

func (s *ExpressionContext) LESS_OR_EQUALS() antlr.TerminalNode {
	return s.GetToken(PegiParserLESS_OR_EQUALS, 0)
}

func (s *ExpressionContext) GREATER() antlr.TerminalNode {
	return s.GetToken(PegiParserGREATER, 0)
}

func (s *ExpressionContext) GREATER_OR_EQUALS() antlr.TerminalNode {
	return s.GetToken(PegiParserGREATER_OR_EQUALS, 0)
}

func (s *ExpressionContext) LOGICAL_AND() antlr.TerminalNode {
	return s.GetToken(PegiParserLOGICAL_AND, 0)
}

func (s *ExpressionContext) LOGICAL_OR() antlr.TerminalNode {
	return s.GetToken(PegiParserLOGICAL_OR, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *PegiParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 30
	p.EnterRecursionRule(localctx, 30, PegiParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(187)
			p.PrimaryExpr()
		}

	case 2:
		{
			p.SetState(188)
			p.UnaryExpr()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(206)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PegiParserRULE_expression)
				p.SetState(191)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(192)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).mul_op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == PegiParserDIV || _la == PegiParserSTAR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).mul_op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(193)
					p.expression(6)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PegiParserRULE_expression)
				p.SetState(194)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(195)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).add_op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == PegiParserPLUS || _la == PegiParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).add_op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(196)
					p.expression(5)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PegiParserRULE_expression)
				p.SetState(197)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(198)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).rel_op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PegiParserEQUALS)|(1<<PegiParserNOT_EQUALS)|(1<<PegiParserLESS)|(1<<PegiParserLESS_OR_EQUALS)|(1<<PegiParserGREATER)|(1<<PegiParserGREATER_OR_EQUALS))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).rel_op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(199)
					p.expression(4)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PegiParserRULE_expression)
				p.SetState(200)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(201)
					p.Match(PegiParserLOGICAL_AND)
				}
				{
					p.SetState(202)
					p.expression(3)
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PegiParserRULE_expression)
				p.SetState(203)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(204)
					p.Match(PegiParserLOGICAL_OR)
				}
				{
					p.SetState(205)
					p.expression(2)
				}

			}

		}
		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimaryExprContext is an interface to support dynamic dispatch.
type IPrimaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryExprContext differentiates from other interfaces.
	IsPrimaryExprContext()
}

type PrimaryExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExprContext() *PrimaryExprContext {
	var p = new(PrimaryExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_primaryExpr
	return p
}

func (*PrimaryExprContext) IsPrimaryExprContext() {}

func NewPrimaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExprContext {
	var p = new(PrimaryExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_primaryExpr

	return p
}

func (s *PrimaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExprContext) Operand() IOperandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *PrimaryExprContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *PrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterPrimaryExpr(s)
	}
}

func (s *PrimaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitPrimaryExpr(s)
	}
}

func (s *PrimaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitPrimaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) PrimaryExpr() (localctx IPrimaryExprContext) {
	localctx = NewPrimaryExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, PegiParserRULE_primaryExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(211)
			p.Operand()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(212)
			p.FunctionCall()
		}

	}

	return localctx
}

// IOperandContext is an interface to support dynamic dispatch.
type IOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperandContext differentiates from other interfaces.
	IsOperandContext()
}

type OperandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandContext() *OperandContext {
	var p = new(OperandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_operand
	return p
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_operand

	return p
}

func (s *OperandContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *OperandContext) OperandName() IOperandNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperandNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperandNameContext)
}

func (s *OperandContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(PegiParserL_PAREN, 0)
}

func (s *OperandContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OperandContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(PegiParserR_PAREN, 0)
}

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterOperand(s)
	}
}

func (s *OperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitOperand(s)
	}
}

func (s *OperandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitOperand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) Operand() (localctx IOperandContext) {
	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, PegiParserRULE_operand)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(221)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PegiParserL_CURLY, PegiParserL_BRACKET, PegiParserDECIMAL_LIT, PegiParserBINARY_LIT, PegiParserOCTAL_LIT, PegiParserHEX_LIT, PegiParserFLOAT_LIT, PegiParserRUNE_LIT, PegiParserRAW_STRING_LIT, PegiParserINTERPRETED_STRING_LIT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(215)
			p.Literal()
		}

	case PegiParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(216)
			p.OperandName()
		}

	case PegiParserL_PAREN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(217)
			p.Match(PegiParserL_PAREN)
		}
		{
			p.SetState(218)
			p.expression(0)
		}
		{
			p.SetState(219)
			p.Match(PegiParserR_PAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) BasicLit() IBasicLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBasicLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBasicLitContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, PegiParserRULE_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.BasicLit()
	}

	return localctx
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *FunctionCallContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, PegiParserRULE_functionCall)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(225)
		p.TypeName()
	}
	{
		p.SetState(226)
		p.Arguments()
	}

	return localctx
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(PegiParserL_PAREN, 0)
}

func (s *ArgumentsContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(PegiParserR_PAREN, 0)
}

func (s *ArgumentsContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ArgumentsContext) TypeT() ITypeTContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeTContext)
}

func (s *ArgumentsContext) ELLIPSIS() antlr.TerminalNode {
	return s.GetToken(PegiParserELLIPSIS, 0)
}

func (s *ArgumentsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PegiParserCOMMA)
}

func (s *ArgumentsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PegiParserCOMMA, i)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (s *ArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, PegiParserRULE_arguments)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
		p.Match(PegiParserL_PAREN)
	}
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PegiParserINTERFACE)|(1<<PegiParserIDENTIFIER)|(1<<PegiParserL_PAREN)|(1<<PegiParserL_CURLY)|(1<<PegiParserL_BRACKET)|(1<<PegiParserEXCLAMATION))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(PegiParserUNARY_MINUS-33))|(1<<(PegiParserDECIMAL_LIT-33))|(1<<(PegiParserBINARY_LIT-33))|(1<<(PegiParserOCTAL_LIT-33))|(1<<(PegiParserHEX_LIT-33))|(1<<(PegiParserFLOAT_LIT-33))|(1<<(PegiParserRUNE_LIT-33))|(1<<(PegiParserRAW_STRING_LIT-33))|(1<<(PegiParserINTERPRETED_STRING_LIT-33)))) != 0) {
		p.SetState(235)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(229)
				p.ExpressionList()
			}

		case 2:
			{
				p.SetState(230)
				p.TypeT()
			}
			p.SetState(233)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(231)
					p.Match(PegiParserCOMMA)
				}
				{
					p.SetState(232)
					p.ExpressionList()
				}

			}

		}
		p.SetState(238)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PegiParserELLIPSIS {
			{
				p.SetState(237)
				p.Match(PegiParserELLIPSIS)
			}

		}
		p.SetState(241)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PegiParserCOMMA {
			{
				p.SetState(240)
				p.Match(PegiParserCOMMA)
			}

		}

	}
	{
		p.SetState(245)
		p.Match(PegiParserR_PAREN)
	}

	return localctx
}

// IBasicLitContext is an interface to support dynamic dispatch.
type IBasicLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBasicLitContext differentiates from other interfaces.
	IsBasicLitContext()
}

type BasicLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBasicLitContext() *BasicLitContext {
	var p = new(BasicLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_basicLit
	return p
}

func (*BasicLitContext) IsBasicLitContext() {}

func NewBasicLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BasicLitContext {
	var p = new(BasicLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_basicLit

	return p
}

func (s *BasicLitContext) GetParser() antlr.Parser { return s.parser }

func (s *BasicLitContext) Integer() IIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *BasicLitContext) MessageValue() IMessageValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageValueContext)
}

func (s *BasicLitContext) SliceValue() ISliceValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISliceValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISliceValueContext)
}

func (s *BasicLitContext) String_() IString_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_Context)
}

func (s *BasicLitContext) FLOAT_LIT() antlr.TerminalNode {
	return s.GetToken(PegiParserFLOAT_LIT, 0)
}

func (s *BasicLitContext) RUNE_LIT() antlr.TerminalNode {
	return s.GetToken(PegiParserRUNE_LIT, 0)
}

func (s *BasicLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BasicLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BasicLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterBasicLit(s)
	}
}

func (s *BasicLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitBasicLit(s)
	}
}

func (s *BasicLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitBasicLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) BasicLit() (localctx IBasicLitContext) {
	localctx = NewBasicLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, PegiParserRULE_basicLit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(247)
			p.Integer()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(248)
			p.MessageValue()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(249)
			p.SliceValue()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(250)
			p.String_()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(251)
			p.Match(PegiParserFLOAT_LIT)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(252)
			p.Match(PegiParserRUNE_LIT)
		}

	}

	return localctx
}

// IOperandNameContext is an interface to support dynamic dispatch.
type IOperandNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperandNameContext differentiates from other interfaces.
	IsOperandNameContext()
}

type OperandNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandNameContext() *OperandNameContext {
	var p = new(OperandNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_operandName
	return p
}

func (*OperandNameContext) IsOperandNameContext() {}

func NewOperandNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandNameContext {
	var p = new(OperandNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_operandName

	return p
}

func (s *OperandNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandNameContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(PegiParserIDENTIFIER)
}

func (s *OperandNameContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(PegiParserIDENTIFIER, i)
}

func (s *OperandNameContext) DOT() antlr.TerminalNode {
	return s.GetToken(PegiParserDOT, 0)
}

func (s *OperandNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterOperandName(s)
	}
}

func (s *OperandNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitOperandName(s)
	}
}

func (s *OperandNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitOperandName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) OperandName() (localctx IOperandNameContext) {
	localctx = NewOperandNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, PegiParserRULE_operandName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(255)
		p.Match(PegiParserIDENTIFIER)
	}
	p.SetState(258)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(256)
			p.Match(PegiParserDOT)
		}
		{
			p.SetState(257)
			p.Match(PegiParserIDENTIFIER)
		}

	}

	return localctx
}

// IUnaryExprContext is an interface to support dynamic dispatch.
type IUnaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetUnary_op returns the unary_op token.
	GetUnary_op() antlr.Token

	// SetUnary_op sets the unary_op token.
	SetUnary_op(antlr.Token)

	// IsUnaryExprContext differentiates from other interfaces.
	IsUnaryExprContext()
}

type UnaryExprContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	unary_op antlr.Token
}

func NewEmptyUnaryExprContext() *UnaryExprContext {
	var p = new(UnaryExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_unaryExpr
	return p
}

func (*UnaryExprContext) IsUnaryExprContext() {}

func NewUnaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExprContext {
	var p = new(UnaryExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_unaryExpr

	return p
}

func (s *UnaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExprContext) GetUnary_op() antlr.Token { return s.unary_op }

func (s *UnaryExprContext) SetUnary_op(v antlr.Token) { s.unary_op = v }

func (s *UnaryExprContext) PrimaryExpr() IPrimaryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *UnaryExprContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryExprContext) UNARY_MINUS() antlr.TerminalNode {
	return s.GetToken(PegiParserUNARY_MINUS, 0)
}

func (s *UnaryExprContext) EXCLAMATION() antlr.TerminalNode {
	return s.GetToken(PegiParserEXCLAMATION, 0)
}

func (s *UnaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterUnaryExpr(s)
	}
}

func (s *UnaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitUnaryExpr(s)
	}
}

func (s *UnaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitUnaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) UnaryExpr() (localctx IUnaryExprContext) {
	localctx = NewUnaryExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, PegiParserRULE_unaryExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(263)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PegiParserIDENTIFIER, PegiParserL_PAREN, PegiParserL_CURLY, PegiParserL_BRACKET, PegiParserDECIMAL_LIT, PegiParserBINARY_LIT, PegiParserOCTAL_LIT, PegiParserHEX_LIT, PegiParserFLOAT_LIT, PegiParserRUNE_LIT, PegiParserRAW_STRING_LIT, PegiParserINTERPRETED_STRING_LIT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(260)
			p.PrimaryExpr()
		}

	case PegiParserEXCLAMATION, PegiParserUNARY_MINUS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(261)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnaryExprContext).unary_op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == PegiParserEXCLAMATION || _la == PegiParserUNARY_MINUS) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnaryExprContext).unary_op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(262)
			p.expression(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISliceValueContext is an interface to support dynamic dispatch.
type ISliceValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSliceValueContext differentiates from other interfaces.
	IsSliceValueContext()
}

type SliceValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySliceValueContext() *SliceValueContext {
	var p = new(SliceValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_sliceValue
	return p
}

func (*SliceValueContext) IsSliceValueContext() {}

func NewSliceValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceValueContext {
	var p = new(SliceValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_sliceValue

	return p
}

func (s *SliceValueContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceValueContext) SliceValueSingleLine() ISliceValueSingleLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISliceValueSingleLineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISliceValueSingleLineContext)
}

func (s *SliceValueContext) SliceValueMultiLine() ISliceValueMultiLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISliceValueMultiLineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISliceValueMultiLineContext)
}

func (s *SliceValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SliceValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterSliceValue(s)
	}
}

func (s *SliceValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitSliceValue(s)
	}
}

func (s *SliceValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitSliceValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) SliceValue() (localctx ISliceValueContext) {
	localctx = NewSliceValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, PegiParserRULE_sliceValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(265)
			p.SliceValueSingleLine()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(266)
			p.SliceValueMultiLine()
		}

	}

	return localctx
}

// ISliceValueSingleLineContext is an interface to support dynamic dispatch.
type ISliceValueSingleLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSliceValueSingleLineContext differentiates from other interfaces.
	IsSliceValueSingleLineContext()
}

type SliceValueSingleLineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySliceValueSingleLineContext() *SliceValueSingleLineContext {
	var p = new(SliceValueSingleLineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_sliceValueSingleLine
	return p
}

func (*SliceValueSingleLineContext) IsSliceValueSingleLineContext() {}

func NewSliceValueSingleLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceValueSingleLineContext {
	var p = new(SliceValueSingleLineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_sliceValueSingleLine

	return p
}

func (s *SliceValueSingleLineContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceValueSingleLineContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(PegiParserL_BRACKET, 0)
}

func (s *SliceValueSingleLineContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *SliceValueSingleLineContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SliceValueSingleLineContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(PegiParserR_BRACKET, 0)
}

func (s *SliceValueSingleLineContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PegiParserCOMMA)
}

func (s *SliceValueSingleLineContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PegiParserCOMMA, i)
}

func (s *SliceValueSingleLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceValueSingleLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SliceValueSingleLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterSliceValueSingleLine(s)
	}
}

func (s *SliceValueSingleLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitSliceValueSingleLine(s)
	}
}

func (s *SliceValueSingleLineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitSliceValueSingleLine(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) SliceValueSingleLine() (localctx ISliceValueSingleLineContext) {
	localctx = NewSliceValueSingleLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, PegiParserRULE_sliceValueSingleLine)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(269)
		p.Match(PegiParserL_BRACKET)
	}
	{
		p.SetState(270)
		p.expression(0)
	}
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PegiParserCOMMA {
		{
			p.SetState(271)
			p.Match(PegiParserCOMMA)
		}
		{
			p.SetState(272)
			p.expression(0)
		}

		p.SetState(277)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(278)
		p.Match(PegiParserR_BRACKET)
	}

	return localctx
}

// ISliceValueMultiLineContext is an interface to support dynamic dispatch.
type ISliceValueMultiLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSliceValueMultiLineContext differentiates from other interfaces.
	IsSliceValueMultiLineContext()
}

type SliceValueMultiLineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySliceValueMultiLineContext() *SliceValueMultiLineContext {
	var p = new(SliceValueMultiLineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_sliceValueMultiLine
	return p
}

func (*SliceValueMultiLineContext) IsSliceValueMultiLineContext() {}

func NewSliceValueMultiLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceValueMultiLineContext {
	var p = new(SliceValueMultiLineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_sliceValueMultiLine

	return p
}

func (s *SliceValueMultiLineContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceValueMultiLineContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(PegiParserL_BRACKET, 0)
}

func (s *SliceValueMultiLineContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *SliceValueMultiLineContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(PegiParserR_BRACKET, 0)
}

func (s *SliceValueMultiLineContext) AllSliceValueMultiLineLine() []ISliceValueMultiLineLineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISliceValueMultiLineLineContext)(nil)).Elem())
	var tst = make([]ISliceValueMultiLineLineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISliceValueMultiLineLineContext)
		}
	}

	return tst
}

func (s *SliceValueMultiLineContext) SliceValueMultiLineLine(i int) ISliceValueMultiLineLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISliceValueMultiLineLineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISliceValueMultiLineLineContext)
}

func (s *SliceValueMultiLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceValueMultiLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SliceValueMultiLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterSliceValueMultiLine(s)
	}
}

func (s *SliceValueMultiLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitSliceValueMultiLine(s)
	}
}

func (s *SliceValueMultiLineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitSliceValueMultiLine(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) SliceValueMultiLine() (localctx ISliceValueMultiLineContext) {
	localctx = NewSliceValueMultiLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, PegiParserRULE_sliceValueMultiLine)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(280)
		p.Match(PegiParserL_BRACKET)
	}
	{
		p.SetState(281)
		p.Eos()
	}
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PegiParserIDENTIFIER)|(1<<PegiParserL_PAREN)|(1<<PegiParserL_CURLY)|(1<<PegiParserL_BRACKET)|(1<<PegiParserEXCLAMATION))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(PegiParserUNARY_MINUS-33))|(1<<(PegiParserDECIMAL_LIT-33))|(1<<(PegiParserBINARY_LIT-33))|(1<<(PegiParserOCTAL_LIT-33))|(1<<(PegiParserHEX_LIT-33))|(1<<(PegiParserFLOAT_LIT-33))|(1<<(PegiParserRUNE_LIT-33))|(1<<(PegiParserRAW_STRING_LIT-33))|(1<<(PegiParserINTERPRETED_STRING_LIT-33)))) != 0) {
		{
			p.SetState(282)
			p.SliceValueMultiLineLine()
		}

		p.SetState(287)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(288)
		p.Match(PegiParserR_BRACKET)
	}

	return localctx
}

// ISliceValueMultiLineLineContext is an interface to support dynamic dispatch.
type ISliceValueMultiLineLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSliceValueMultiLineLineContext differentiates from other interfaces.
	IsSliceValueMultiLineLineContext()
}

type SliceValueMultiLineLineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySliceValueMultiLineLineContext() *SliceValueMultiLineLineContext {
	var p = new(SliceValueMultiLineLineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_sliceValueMultiLineLine
	return p
}

func (*SliceValueMultiLineLineContext) IsSliceValueMultiLineLineContext() {}

func NewSliceValueMultiLineLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceValueMultiLineLineContext {
	var p = new(SliceValueMultiLineLineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_sliceValueMultiLineLine

	return p
}

func (s *SliceValueMultiLineLineContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceValueMultiLineLineContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SliceValueMultiLineLineContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PegiParserCOMMA, 0)
}

func (s *SliceValueMultiLineLineContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *SliceValueMultiLineLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceValueMultiLineLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SliceValueMultiLineLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterSliceValueMultiLineLine(s)
	}
}

func (s *SliceValueMultiLineLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitSliceValueMultiLineLine(s)
	}
}

func (s *SliceValueMultiLineLineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitSliceValueMultiLineLine(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) SliceValueMultiLineLine() (localctx ISliceValueMultiLineLineContext) {
	localctx = NewSliceValueMultiLineLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, PegiParserRULE_sliceValueMultiLineLine)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(290)
		p.expression(0)
	}
	{
		p.SetState(291)
		p.Match(PegiParserCOMMA)
	}
	{
		p.SetState(292)
		p.Eos()
	}

	return localctx
}

// IMessageValueContext is an interface to support dynamic dispatch.
type IMessageValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageValueContext differentiates from other interfaces.
	IsMessageValueContext()
}

type MessageValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageValueContext() *MessageValueContext {
	var p = new(MessageValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_messageValue
	return p
}

func (*MessageValueContext) IsMessageValueContext() {}

func NewMessageValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageValueContext {
	var p = new(MessageValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_messageValue

	return p
}

func (s *MessageValueContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageValueContext) MessageValueSingleLine() IMessageValueSingleLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageValueSingleLineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageValueSingleLineContext)
}

func (s *MessageValueContext) MessageValueMultiLine() IMessageValueMultiLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageValueMultiLineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageValueMultiLineContext)
}

func (s *MessageValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterMessageValue(s)
	}
}

func (s *MessageValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitMessageValue(s)
	}
}

func (s *MessageValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitMessageValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) MessageValue() (localctx IMessageValueContext) {
	localctx = NewMessageValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, PegiParserRULE_messageValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(294)
			p.MessageValueSingleLine()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(295)
			p.MessageValueMultiLine()
		}

	}

	return localctx
}

// IKeyedElementKVContext is an interface to support dynamic dispatch.
type IKeyedElementKVContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyedElementKVContext differentiates from other interfaces.
	IsKeyedElementKVContext()
}

type KeyedElementKVContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyedElementKVContext() *KeyedElementKVContext {
	var p = new(KeyedElementKVContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_keyedElementKV
	return p
}

func (*KeyedElementKVContext) IsKeyedElementKVContext() {}

func NewKeyedElementKVContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyedElementKVContext {
	var p = new(KeyedElementKVContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_keyedElementKV

	return p
}

func (s *KeyedElementKVContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyedElementKVContext) Key() IKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *KeyedElementKVContext) COLON() antlr.TerminalNode {
	return s.GetToken(PegiParserCOLON, 0)
}

func (s *KeyedElementKVContext) Element() IElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementContext)
}

func (s *KeyedElementKVContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyedElementKVContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyedElementKVContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterKeyedElementKV(s)
	}
}

func (s *KeyedElementKVContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitKeyedElementKV(s)
	}
}

func (s *KeyedElementKVContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitKeyedElementKV(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) KeyedElementKV() (localctx IKeyedElementKVContext) {
	localctx = NewKeyedElementKVContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, PegiParserRULE_keyedElementKV)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(298)
		p.Key()
	}
	{
		p.SetState(299)
		p.Match(PegiParserCOLON)
	}
	{
		p.SetState(300)
		p.Element()
	}

	return localctx
}

// IKeyedElementIdentiferContext is an interface to support dynamic dispatch.
type IKeyedElementIdentiferContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyedElementIdentiferContext differentiates from other interfaces.
	IsKeyedElementIdentiferContext()
}

type KeyedElementIdentiferContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyedElementIdentiferContext() *KeyedElementIdentiferContext {
	var p = new(KeyedElementIdentiferContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_keyedElementIdentifer
	return p
}

func (*KeyedElementIdentiferContext) IsKeyedElementIdentiferContext() {}

func NewKeyedElementIdentiferContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyedElementIdentiferContext {
	var p = new(KeyedElementIdentiferContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_keyedElementIdentifer

	return p
}

func (s *KeyedElementIdentiferContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyedElementIdentiferContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PegiParserIDENTIFIER, 0)
}

func (s *KeyedElementIdentiferContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyedElementIdentiferContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyedElementIdentiferContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterKeyedElementIdentifer(s)
	}
}

func (s *KeyedElementIdentiferContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitKeyedElementIdentifer(s)
	}
}

func (s *KeyedElementIdentiferContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitKeyedElementIdentifer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) KeyedElementIdentifer() (localctx IKeyedElementIdentiferContext) {
	localctx = NewKeyedElementIdentiferContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, PegiParserRULE_keyedElementIdentifer)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(302)
		p.Match(PegiParserIDENTIFIER)
	}

	return localctx
}

// IKeyedElementContext is an interface to support dynamic dispatch.
type IKeyedElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyedElementContext differentiates from other interfaces.
	IsKeyedElementContext()
}

type KeyedElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyedElementContext() *KeyedElementContext {
	var p = new(KeyedElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_keyedElement
	return p
}

func (*KeyedElementContext) IsKeyedElementContext() {}

func NewKeyedElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyedElementContext {
	var p = new(KeyedElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_keyedElement

	return p
}

func (s *KeyedElementContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyedElementContext) KeyedElementIdentifer() IKeyedElementIdentiferContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyedElementIdentiferContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyedElementIdentiferContext)
}

func (s *KeyedElementContext) KeyedElementKV() IKeyedElementKVContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyedElementKVContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyedElementKVContext)
}

func (s *KeyedElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyedElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyedElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterKeyedElement(s)
	}
}

func (s *KeyedElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitKeyedElement(s)
	}
}

func (s *KeyedElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitKeyedElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) KeyedElement() (localctx IKeyedElementContext) {
	localctx = NewKeyedElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, PegiParserRULE_keyedElement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(304)
			p.KeyedElementIdentifer()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(305)
			p.KeyedElementKV()
		}

	}

	return localctx
}

// IMessageValueSingleLineContext is an interface to support dynamic dispatch.
type IMessageValueSingleLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageValueSingleLineContext differentiates from other interfaces.
	IsMessageValueSingleLineContext()
}

type MessageValueSingleLineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageValueSingleLineContext() *MessageValueSingleLineContext {
	var p = new(MessageValueSingleLineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_messageValueSingleLine
	return p
}

func (*MessageValueSingleLineContext) IsMessageValueSingleLineContext() {}

func NewMessageValueSingleLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageValueSingleLineContext {
	var p = new(MessageValueSingleLineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_messageValueSingleLine

	return p
}

func (s *MessageValueSingleLineContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageValueSingleLineContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(PegiParserL_CURLY, 0)
}

func (s *MessageValueSingleLineContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(PegiParserR_CURLY, 0)
}

func (s *MessageValueSingleLineContext) AllKeyedElement() []IKeyedElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKeyedElementContext)(nil)).Elem())
	var tst = make([]IKeyedElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKeyedElementContext)
		}
	}

	return tst
}

func (s *MessageValueSingleLineContext) KeyedElement(i int) IKeyedElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyedElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKeyedElementContext)
}

func (s *MessageValueSingleLineContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PegiParserCOMMA)
}

func (s *MessageValueSingleLineContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PegiParserCOMMA, i)
}

func (s *MessageValueSingleLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageValueSingleLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageValueSingleLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterMessageValueSingleLine(s)
	}
}

func (s *MessageValueSingleLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitMessageValueSingleLine(s)
	}
}

func (s *MessageValueSingleLineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitMessageValueSingleLine(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) MessageValueSingleLine() (localctx IMessageValueSingleLineContext) {
	localctx = NewMessageValueSingleLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, PegiParserRULE_messageValueSingleLine)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(308)
		p.Match(PegiParserL_CURLY)
	}
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PegiParserIDENTIFIER {
		{
			p.SetState(309)
			p.KeyedElement()
		}

	}
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PegiParserCOMMA {
		{
			p.SetState(312)
			p.Match(PegiParserCOMMA)
		}
		{
			p.SetState(313)
			p.KeyedElement()
		}

		p.SetState(318)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(319)
		p.Match(PegiParserR_CURLY)
	}

	return localctx
}

// IMessageValueMultiLineContext is an interface to support dynamic dispatch.
type IMessageValueMultiLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageValueMultiLineContext differentiates from other interfaces.
	IsMessageValueMultiLineContext()
}

type MessageValueMultiLineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageValueMultiLineContext() *MessageValueMultiLineContext {
	var p = new(MessageValueMultiLineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_messageValueMultiLine
	return p
}

func (*MessageValueMultiLineContext) IsMessageValueMultiLineContext() {}

func NewMessageValueMultiLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageValueMultiLineContext {
	var p = new(MessageValueMultiLineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_messageValueMultiLine

	return p
}

func (s *MessageValueMultiLineContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageValueMultiLineContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(PegiParserL_CURLY, 0)
}

func (s *MessageValueMultiLineContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *MessageValueMultiLineContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(PegiParserR_CURLY, 0)
}

func (s *MessageValueMultiLineContext) AllLineKeyedElement() []ILineKeyedElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineKeyedElementContext)(nil)).Elem())
	var tst = make([]ILineKeyedElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineKeyedElementContext)
		}
	}

	return tst
}

func (s *MessageValueMultiLineContext) LineKeyedElement(i int) ILineKeyedElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineKeyedElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineKeyedElementContext)
}

func (s *MessageValueMultiLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageValueMultiLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageValueMultiLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterMessageValueMultiLine(s)
	}
}

func (s *MessageValueMultiLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitMessageValueMultiLine(s)
	}
}

func (s *MessageValueMultiLineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitMessageValueMultiLine(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) MessageValueMultiLine() (localctx IMessageValueMultiLineContext) {
	localctx = NewMessageValueMultiLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, PegiParserRULE_messageValueMultiLine)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(321)
		p.Match(PegiParserL_CURLY)
	}
	{
		p.SetState(322)
		p.Eos()
	}
	p.SetState(324)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PegiParserIDENTIFIER {
		{
			p.SetState(323)
			p.LineKeyedElement()
		}

		p.SetState(326)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(328)
		p.Match(PegiParserR_CURLY)
	}

	return localctx
}

// ILineKeyedElementContext is an interface to support dynamic dispatch.
type ILineKeyedElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLineKeyedElementContext differentiates from other interfaces.
	IsLineKeyedElementContext()
}

type LineKeyedElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineKeyedElementContext() *LineKeyedElementContext {
	var p = new(LineKeyedElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_lineKeyedElement
	return p
}

func (*LineKeyedElementContext) IsLineKeyedElementContext() {}

func NewLineKeyedElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineKeyedElementContext {
	var p = new(LineKeyedElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_lineKeyedElement

	return p
}

func (s *LineKeyedElementContext) GetParser() antlr.Parser { return s.parser }

func (s *LineKeyedElementContext) KeyedElement() IKeyedElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyedElementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyedElementContext)
}

func (s *LineKeyedElementContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PegiParserCOMMA, 0)
}

func (s *LineKeyedElementContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *LineKeyedElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineKeyedElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineKeyedElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterLineKeyedElement(s)
	}
}

func (s *LineKeyedElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitLineKeyedElement(s)
	}
}

func (s *LineKeyedElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitLineKeyedElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) LineKeyedElement() (localctx ILineKeyedElementContext) {
	localctx = NewLineKeyedElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, PegiParserRULE_lineKeyedElement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(330)
		p.KeyedElement()
	}
	{
		p.SetState(331)
		p.Match(PegiParserCOMMA)
	}
	{
		p.SetState(332)
		p.Eos()
	}

	return localctx
}

// IKeyContext is an interface to support dynamic dispatch.
type IKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyContext differentiates from other interfaces.
	IsKeyContext()
}

type KeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyContext() *KeyContext {
	var p = new(KeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_key
	return p
}

func (*KeyContext) IsKeyContext() {}

func NewKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyContext {
	var p = new(KeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_key

	return p
}

func (s *KeyContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PegiParserIDENTIFIER, 0)
}

func (s *KeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterKey(s)
	}
}

func (s *KeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitKey(s)
	}
}

func (s *KeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) Key() (localctx IKeyContext) {
	localctx = NewKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, PegiParserRULE_key)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(334)
		p.Match(PegiParserIDENTIFIER)
	}

	return localctx
}

// IElementContext is an interface to support dynamic dispatch.
type IElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElementContext differentiates from other interfaces.
	IsElementContext()
}

type ElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementContext() *ElementContext {
	var p = new(ElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_element
	return p
}

func (*ElementContext) IsElementContext() {}

func NewElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementContext {
	var p = new(ElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_element

	return p
}

func (s *ElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ElementContext) MessageValue() IMessageValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageValueContext)
}

func (s *ElementContext) SliceValue() ISliceValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISliceValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISliceValueContext)
}

func (s *ElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterElement(s)
	}
}

func (s *ElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitElement(s)
	}
}

func (s *ElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) Element() (localctx IElementContext) {
	localctx = NewElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, PegiParserRULE_element)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(339)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(336)
			p.expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(337)
			p.MessageValue()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(338)
			p.SliceValue()
		}

	}

	return localctx
}

// ISliceTypeContext is an interface to support dynamic dispatch.
type ISliceTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSliceTypeContext differentiates from other interfaces.
	IsSliceTypeContext()
}

type SliceTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySliceTypeContext() *SliceTypeContext {
	var p = new(SliceTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_sliceType
	return p
}

func (*SliceTypeContext) IsSliceTypeContext() {}

func NewSliceTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceTypeContext {
	var p = new(SliceTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_sliceType

	return p
}

func (s *SliceTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceTypeContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(PegiParserL_BRACKET, 0)
}

func (s *SliceTypeContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(PegiParserR_BRACKET, 0)
}

func (s *SliceTypeContext) TypeT() ITypeTContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeTContext)
}

func (s *SliceTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SliceTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterSliceType(s)
	}
}

func (s *SliceTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitSliceType(s)
	}
}

func (s *SliceTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitSliceType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) SliceType() (localctx ISliceTypeContext) {
	localctx = NewSliceTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, PegiParserRULE_sliceType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(341)
		p.Match(PegiParserL_BRACKET)
	}
	{
		p.SetState(342)
		p.Match(PegiParserR_BRACKET)
	}
	{
		p.SetState(343)
		p.TypeT()
	}

	return localctx
}

// ITypeTContext is an interface to support dynamic dispatch.
type ITypeTContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeTContext differentiates from other interfaces.
	IsTypeTContext()
}

type TypeTContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeTContext() *TypeTContext {
	var p = new(TypeTContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_typeT
	return p
}

func (*TypeTContext) IsTypeTContext() {}

func NewTypeTContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeTContext {
	var p = new(TypeTContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_typeT

	return p
}

func (s *TypeTContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeTContext) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *TypeTContext) TypeLit() ITypeLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeLitContext)
}

func (s *TypeTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeTContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterTypeT(s)
	}
}

func (s *TypeTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitTypeT(s)
	}
}

func (s *TypeTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitTypeT(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) TypeT() (localctx ITypeTContext) {
	localctx = NewTypeTContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, PegiParserRULE_typeT)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(347)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PegiParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(345)
			p.TypeName()
		}

	case PegiParserINTERFACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(346)
			p.TypeLit()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITypeNameContext is an interface to support dynamic dispatch.
type ITypeNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeNameContext differentiates from other interfaces.
	IsTypeNameContext()
}

type TypeNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeNameContext() *TypeNameContext {
	var p = new(TypeNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_typeName
	return p
}

func (*TypeNameContext) IsTypeNameContext() {}

func NewTypeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeNameContext {
	var p = new(TypeNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_typeName

	return p
}

func (s *TypeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeNameContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(PegiParserIDENTIFIER)
}

func (s *TypeNameContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(PegiParserIDENTIFIER, i)
}

func (s *TypeNameContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(PegiParserDOT)
}

func (s *TypeNameContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(PegiParserDOT, i)
}

func (s *TypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterTypeName(s)
	}
}

func (s *TypeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitTypeName(s)
	}
}

func (s *TypeNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitTypeName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) TypeName() (localctx ITypeNameContext) {
	localctx = NewTypeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, PegiParserRULE_typeName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(349)
		p.Match(PegiParserIDENTIFIER)
	}
	p.SetState(354)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(350)
				p.Match(PegiParserDOT)
			}
			{
				p.SetState(351)
				p.Match(PegiParserIDENTIFIER)
			}

		}
		p.SetState(356)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())
	}

	return localctx
}

// ITypeLitContext is an interface to support dynamic dispatch.
type ITypeLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeLitContext differentiates from other interfaces.
	IsTypeLitContext()
}

type TypeLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeLitContext() *TypeLitContext {
	var p = new(TypeLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_typeLit
	return p
}

func (*TypeLitContext) IsTypeLitContext() {}

func NewTypeLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeLitContext {
	var p = new(TypeLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_typeLit

	return p
}

func (s *TypeLitContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeLitContext) InterfaceType() IInterfaceTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInterfaceTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInterfaceTypeContext)
}

func (s *TypeLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterTypeLit(s)
	}
}

func (s *TypeLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitTypeLit(s)
	}
}

func (s *TypeLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitTypeLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) TypeLit() (localctx ITypeLitContext) {
	localctx = NewTypeLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, PegiParserRULE_typeLit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(357)
		p.InterfaceType()
	}

	return localctx
}

// IInterfaceTypeContext is an interface to support dynamic dispatch.
type IInterfaceTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInterfaceName returns the interfaceName token.
	GetInterfaceName() antlr.Token

	// SetInterfaceName sets the interfaceName token.
	SetInterfaceName(antlr.Token)

	// IsInterfaceTypeContext differentiates from other interfaces.
	IsInterfaceTypeContext()
}

type InterfaceTypeContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	interfaceName antlr.Token
}

func NewEmptyInterfaceTypeContext() *InterfaceTypeContext {
	var p = new(InterfaceTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_interfaceType
	return p
}

func (*InterfaceTypeContext) IsInterfaceTypeContext() {}

func NewInterfaceTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterfaceTypeContext {
	var p = new(InterfaceTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_interfaceType

	return p
}

func (s *InterfaceTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *InterfaceTypeContext) GetInterfaceName() antlr.Token { return s.interfaceName }

func (s *InterfaceTypeContext) SetInterfaceName(v antlr.Token) { s.interfaceName = v }

func (s *InterfaceTypeContext) INTERFACE() antlr.TerminalNode {
	return s.GetToken(PegiParserINTERFACE, 0)
}

func (s *InterfaceTypeContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(PegiParserL_CURLY, 0)
}

func (s *InterfaceTypeContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(PegiParserR_CURLY, 0)
}

func (s *InterfaceTypeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PegiParserIDENTIFIER, 0)
}

func (s *InterfaceTypeContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *InterfaceTypeContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *InterfaceTypeContext) AllFieldDecl() []IFieldDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldDeclContext)(nil)).Elem())
	var tst = make([]IFieldDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldDeclContext)
		}
	}

	return tst
}

func (s *InterfaceTypeContext) FieldDecl(i int) IFieldDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldDeclContext)
}

func (s *InterfaceTypeContext) AllTypeName() []ITypeNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeNameContext)(nil)).Elem())
	var tst = make([]ITypeNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeNameContext)
		}
	}

	return tst
}

func (s *InterfaceTypeContext) TypeName(i int) ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *InterfaceTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterfaceTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InterfaceTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterInterfaceType(s)
	}
}

func (s *InterfaceTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitInterfaceType(s)
	}
}

func (s *InterfaceTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitInterfaceType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) InterfaceType() (localctx IInterfaceTypeContext) {
	localctx = NewInterfaceTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, PegiParserRULE_interfaceType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(359)
		p.Match(PegiParserINTERFACE)
	}
	{
		p.SetState(360)

		var _m = p.Match(PegiParserIDENTIFIER)

		localctx.(*InterfaceTypeContext).interfaceName = _m
	}
	{
		p.SetState(361)
		p.Match(PegiParserL_CURLY)
	}
	p.SetState(370)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(364)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(362)
					p.FieldDecl()
				}

			case 2:
				{
					p.SetState(363)
					p.TypeName()
				}

			}
			{
				p.SetState(366)
				p.Eos()
			}

		}
		p.SetState(372)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext())
	}
	{
		p.SetState(373)
		p.Match(PegiParserR_CURLY)
	}

	return localctx
}

// IFieldDeclContext is an interface to support dynamic dispatch.
type IFieldDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldDeclContext differentiates from other interfaces.
	IsFieldDeclContext()
}

type FieldDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldDeclContext() *FieldDeclContext {
	var p = new(FieldDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_fieldDecl
	return p
}

func (*FieldDeclContext) IsFieldDeclContext() {}

func NewFieldDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldDeclContext {
	var p = new(FieldDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_fieldDecl

	return p
}

func (s *FieldDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PegiParserIDENTIFIER, 0)
}

func (s *FieldDeclContext) TypeT() ITypeTContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeTContext)
}

func (s *FieldDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterFieldDecl(s)
	}
}

func (s *FieldDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitFieldDecl(s)
	}
}

func (s *FieldDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitFieldDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) FieldDecl() (localctx IFieldDeclContext) {
	localctx = NewFieldDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, PegiParserRULE_fieldDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(375)

	if !(p.noTerminatorBetween(2)) {
		panic(antlr.NewFailedPredicateException(p, "p.noTerminatorBetween(2)", ""))
	}
	{
		p.SetState(376)
		p.Match(PegiParserIDENTIFIER)
	}
	{
		p.SetState(377)
		p.TypeT()
	}

	return localctx
}

// IIntegerContext is an interface to support dynamic dispatch.
type IIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerContext differentiates from other interfaces.
	IsIntegerContext()
}

type IntegerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerContext() *IntegerContext {
	var p = new(IntegerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_integer
	return p
}

func (*IntegerContext) IsIntegerContext() {}

func NewIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerContext {
	var p = new(IntegerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_integer

	return p
}

func (s *IntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerContext) DECIMAL_LIT() antlr.TerminalNode {
	return s.GetToken(PegiParserDECIMAL_LIT, 0)
}

func (s *IntegerContext) BINARY_LIT() antlr.TerminalNode {
	return s.GetToken(PegiParserBINARY_LIT, 0)
}

func (s *IntegerContext) OCTAL_LIT() antlr.TerminalNode {
	return s.GetToken(PegiParserOCTAL_LIT, 0)
}

func (s *IntegerContext) HEX_LIT() antlr.TerminalNode {
	return s.GetToken(PegiParserHEX_LIT, 0)
}

func (s *IntegerContext) RUNE_LIT() antlr.TerminalNode {
	return s.GetToken(PegiParserRUNE_LIT, 0)
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitInteger(s)
	}
}

func (s *IntegerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitInteger(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) Integer() (localctx IIntegerContext) {
	localctx = NewIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, PegiParserRULE_integer)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(379)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(PegiParserDECIMAL_LIT-37))|(1<<(PegiParserBINARY_LIT-37))|(1<<(PegiParserOCTAL_LIT-37))|(1<<(PegiParserHEX_LIT-37))|(1<<(PegiParserRUNE_LIT-37)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IEosContext is an interface to support dynamic dispatch.
type IEosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEosContext differentiates from other interfaces.
	IsEosContext()
}

type EosContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEosContext() *EosContext {
	var p = new(EosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PegiParserRULE_eos
	return p
}

func (*EosContext) IsEosContext() {}

func NewEosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EosContext {
	var p = new(EosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PegiParserRULE_eos

	return p
}

func (s *EosContext) GetParser() antlr.Parser { return s.parser }

func (s *EosContext) SEMI() antlr.TerminalNode {
	return s.GetToken(PegiParserSEMI, 0)
}

func (s *EosContext) EOF() antlr.TerminalNode {
	return s.GetToken(PegiParserEOF, 0)
}

func (s *EosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.EnterEos(s)
	}
}

func (s *EosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PegiParserListener); ok {
		listenerT.ExitEos(s)
	}
}

func (s *EosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PegiParserVisitor:
		return t.VisitEos(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PegiParser) Eos() (localctx IEosContext) {
	localctx = NewEosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, PegiParserRULE_eos)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(385)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(381)
			p.Match(PegiParserSEMI)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(382)
			p.Match(PegiParserEOF)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(383)

		if !(p.lineTerminatorAhead()) {
			panic(antlr.NewFailedPredicateException(p, "p.lineTerminatorAhead()", ""))
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(384)

		if !(p.checkPreviousTokenText("}")) {
			panic(antlr.NewFailedPredicateException(p, "p.checkPreviousTokenText(\"}\")", ""))
		}

	}

	return localctx
}

func (p *PegiParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 15:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 42:
		var t *FieldDeclContext = nil
		if localctx != nil {
			t = localctx.(*FieldDeclContext)
		}
		return p.FieldDecl_Sempred(t, predIndex)

	case 44:
		var t *EosContext = nil
		if localctx != nil {
			t = localctx.(*EosContext)
		}
		return p.Eos_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *PegiParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PegiParser) FieldDecl_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.noTerminatorBetween(2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PegiParser) Eos_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return p.lineTerminatorAhead()

	case 7:
		return p.checkPreviousTokenText("}")

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
