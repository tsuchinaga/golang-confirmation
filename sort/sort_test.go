package sort

import (
	"reflect"
	"testing"
)

func TestTest1(t *testing.T) {
	testTable := []struct {
		nums   []int
		expect []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2, 3, 1}, []int{1, 2, 3}},
		{[]int{3, 3, 1}, []int{1, 3, 3}},
		{[]int{1, 1, 1}, []int{1, 1, 1}},
	}

	for i, test := range testTable {
		Test1(test.nums)

		if !reflect.DeepEqual(test.expect, test.nums) {
			t.Fatalf("%s No.%02d 失敗\n期待: %+v\n実際: %+v\n", t.Name(), i+1, test.expect, test.nums)
		}
	}
}

func TestTest2(t *testing.T) {
	testTable := []struct {
		nums   []int
		expect []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2, 3, 1}, []int{3, 2, 1}},
		{[]int{3, 3, 1}, []int{3, 3, 1}},
		{[]int{1, 1, 1}, []int{1, 1, 1}},
	}

	for i, test := range testTable {
		Test2(test.nums)

		if !reflect.DeepEqual(test.expect, test.nums) {
			t.Fatalf("%s No.%02d 失敗\n期待: %+v\n実際: %+v\n", t.Name(), i+1, test.expect, test.nums)
		}
	}
}

// Test3とTest4の回数比較
func TestTest34(t *testing.T) {
	testTable := []struct {
		nums1 []int
		nums2 []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 1, 2, 3, 4}, []int{5, 1, 2, 3, 4}},
		{[]int{269, 2, 874, 173, 10, 6, 87, 611, 150, 104, 31, 852, 13, 593, 216, 16, 17, 38, 798, 152, 77, 76, 23, 35, 167, 63, 27, 283, 65, 30, 46, 343, 564, 734, 301, 36, 715, 37, 845, 785, 41, 42, 64, 44, 455, 18, 599, 60, 394, 50, 51, 441, 53, 634, 74, 148, 786, 456, 252, 121, 54, 62, 245, 55, 47, 648, 187, 68, 69, 70, 71, 673, 73, 377, 168, 22, 21, 3, 79, 226, 176, 154, 83, 859, 85, 86, 58, 589, 272, 90, 91, 207, 713, 105, 95, 96, 97, 562, 836, 159, 101, 208, 103, 174, 94, 106, 107, 108, 109, 502, 111, 288, 743, 746, 165, 116, 425, 118, 119, 828, 5, 149, 129, 560, 139, 339, 794, 675, 123, 192, 577, 132, 704, 752, 277, 136, 640, 869, 125, 140, 141, 142, 143, 144, 145, 585, 147, 56, 122, 137, 204, 20, 153, 82, 88, 49, 157, 158, 225, 160, 421, 162, 784, 537, 230, 166, 222, 131, 177, 170, 357, 172, 4, 179, 175, 112, 169, 178, 48, 511, 433, 182, 183, 184, 185, 315, 773, 280, 424, 190, 191, 450, 203, 194, 195, 196, 534, 790, 199, 200, 333, 342, 193, 151, 210, 206, 580, 102, 351, 367, 211, 212, 771, 214, 215, 25, 217, 218, 477, 220, 221, 1, 223, 224, 100, 80, 213, 161, 432, 520, 231, 232, 233, 234, 486, 667, 52, 788, 239, 240, 241, 723, 355, 631, 189, 544, 664, 248, 249, 628, 757, 59, 253, 254, 255, 479, 257, 851, 259, 722, 541, 262, 263, 264, 209, 566, 267, 871, 15, 266, 271, 89, 61, 350, 285, 782, 8, 278, 279, 627, 281, 719, 117, 781, 275, 328, 287, 14, 289, 290, 291, 292, 293, 294, 637, 296, 297, 298, 299, 303, 24, 302, 803, 800, 606, 317, 307, 820, 309, 310, 311, 312, 438, 314, 459, 316, 354, 458, 649, 335, 321, 434, 246, 691, 325, 482, 327, 286, 300, 330, 156, 406, 201, 331, 320, 701, 337, 385, 134, 340, 341, 202, 43, 344, 345, 346, 347, 171, 386, 274, 265, 352, 750, 696, 655, 356, 348, 358, 359, 360, 525, 680, 363, 364, 365, 366, 205, 368, 369, 370, 219, 372, 373, 374, 375, 393, 197, 378, 379, 475, 28, 382, 32, 384, 338, 349, 387, 633, 389, 390, 391, 392, 376, 591, 395, 396, 397, 398, 399, 400, 464, 402, 403, 404, 405, 332, 454, 546, 409, 410, 411, 806, 868, 414, 415, 416, 417, 418, 419, 420, 779, 422, 758, 524, 381, 229, 427, 428, 429, 430, 431, 426, 181, 19, 435, 135, 437, 313, 647, 440, 237, 448, 443, 243, 445, 446, 447, 442, 449, 130, 451, 452, 453, 407, 45, 11, 457, 318, 284, 460, 461, 462, 463, 401, 465, 466, 467, 468, 469, 470, 471, 488, 473, 474, 532, 476, 371, 478, 256, 480, 481, 326, 483, 484, 412, 235, 487, 472, 489, 810, 875, 492, 552, 494, 495, 496, 497, 498, 769, 720, 501, 110, 81, 507, 505, 506, 504, 361, 590, 510, 180, 512, 513, 514, 515, 509, 247, 518, 519, 115, 521, 522, 523, 26, 508, 545, 527, 612, 529, 530, 531, 380, 716, 383, 535, 536, 164, 538, 539, 540, 261, 542, 543, 238, 526, 408, 670, 548, 799, 550, 636, 493, 553, 554, 555, 679, 557, 741, 559, 124, 561, 98, 563, 34, 565, 270, 567, 568, 569, 570, 571, 572, 573, 574, 575, 576, 75, 880, 579, 92, 581, 582, 583, 584, 120, 645, 587, 876, 155, 516, 334, 592, 503, 766, 595, 596, 597, 598, 29, 600, 601, 602, 603, 604, 605, 305, 607, 608, 533, 610, 436, 528, 846, 614, 629, 616, 617, 877, 619, 620, 621, 622, 623, 624, 625, 626, 188, 250, 615, 630, 244, 632, 388, 273, 748, 551, 236, 638, 639, 99, 641, 642, 643, 644, 586, 646, 439, 72, 319, 650, 651, 652, 653, 656, 444, 654, 657, 658, 659, 660, 807, 662, 663, 517, 665, 666, 295, 668, 669, 547, 671, 672, 66, 674, 128, 676, 677, 678, 556, 362, 681, 864, 282, 684, 685, 686, 687, 688, 689, 690, 324, 692, 693, 694, 695, 306, 697, 698, 725, 700, 780, 804, 703, 133, 705, 682, 707, 708, 709, 710, 711, 712, 93, 714, 7, 609, 717, 718, 683, 500, 721, 260, 242, 724, 793, 726, 727, 728, 738, 730, 731, 732, 733, 33, 735, 736, 737, 729, 739, 740, 558, 742, 113, 423, 792, 114, 747, 635, 749, 353, 751, 126, 753, 754, 578, 756, 251, 744, 759, 760, 761, 762, 763, 764, 765, 594, 770, 768, 499, 767, 227, 772, 67, 774, 775, 776, 777, 778, 228, 336, 186, 276, 783, 163, 40, 57, 787, 323, 789, 198, 791, 745, 699, 9, 795, 796, 797, 322, 549, 304, 801, 802, 329, 702, 805, 485, 661, 808, 809, 490, 811, 812, 813, 814, 815, 816, 817, 818, 819, 308, 821, 822, 823, 824, 825, 826, 827, 146, 829, 830, 831, 832, 833, 834, 835, 881, 837, 838, 839, 840, 841, 842, 843, 844, 39, 613, 847, 848, 849, 850, 258, 12, 853, 854, 855, 856, 857, 858, 84, 860, 861, 862, 863, 706, 865, 866, 867, 413, 138, 870, 268, 872, 873, 78, 491, 588, 618, 878, 879, 755, 127, 882, 883}, []int{269, 2, 874, 173, 10, 6, 87, 611, 150, 104, 31, 852, 13, 593, 216, 16, 17, 38, 798, 152, 77, 76, 23, 35, 167, 63, 27, 283, 65, 30, 46, 343, 564, 734, 301, 36, 715, 37, 845, 785, 41, 42, 64, 44, 455, 18, 599, 60, 394, 50, 51, 441, 53, 634, 74, 148, 786, 456, 252, 121, 54, 62, 245, 55, 47, 648, 187, 68, 69, 70, 71, 673, 73, 377, 168, 22, 21, 3, 79, 226, 176, 154, 83, 859, 85, 86, 58, 589, 272, 90, 91, 207, 713, 105, 95, 96, 97, 562, 836, 159, 101, 208, 103, 174, 94, 106, 107, 108, 109, 502, 111, 288, 743, 746, 165, 116, 425, 118, 119, 828, 5, 149, 129, 560, 139, 339, 794, 675, 123, 192, 577, 132, 704, 752, 277, 136, 640, 869, 125, 140, 141, 142, 143, 144, 145, 585, 147, 56, 122, 137, 204, 20, 153, 82, 88, 49, 157, 158, 225, 160, 421, 162, 784, 537, 230, 166, 222, 131, 177, 170, 357, 172, 4, 179, 175, 112, 169, 178, 48, 511, 433, 182, 183, 184, 185, 315, 773, 280, 424, 190, 191, 450, 203, 194, 195, 196, 534, 790, 199, 200, 333, 342, 193, 151, 210, 206, 580, 102, 351, 367, 211, 212, 771, 214, 215, 25, 217, 218, 477, 220, 221, 1, 223, 224, 100, 80, 213, 161, 432, 520, 231, 232, 233, 234, 486, 667, 52, 788, 239, 240, 241, 723, 355, 631, 189, 544, 664, 248, 249, 628, 757, 59, 253, 254, 255, 479, 257, 851, 259, 722, 541, 262, 263, 264, 209, 566, 267, 871, 15, 266, 271, 89, 61, 350, 285, 782, 8, 278, 279, 627, 281, 719, 117, 781, 275, 328, 287, 14, 289, 290, 291, 292, 293, 294, 637, 296, 297, 298, 299, 303, 24, 302, 803, 800, 606, 317, 307, 820, 309, 310, 311, 312, 438, 314, 459, 316, 354, 458, 649, 335, 321, 434, 246, 691, 325, 482, 327, 286, 300, 330, 156, 406, 201, 331, 320, 701, 337, 385, 134, 340, 341, 202, 43, 344, 345, 346, 347, 171, 386, 274, 265, 352, 750, 696, 655, 356, 348, 358, 359, 360, 525, 680, 363, 364, 365, 366, 205, 368, 369, 370, 219, 372, 373, 374, 375, 393, 197, 378, 379, 475, 28, 382, 32, 384, 338, 349, 387, 633, 389, 390, 391, 392, 376, 591, 395, 396, 397, 398, 399, 400, 464, 402, 403, 404, 405, 332, 454, 546, 409, 410, 411, 806, 868, 414, 415, 416, 417, 418, 419, 420, 779, 422, 758, 524, 381, 229, 427, 428, 429, 430, 431, 426, 181, 19, 435, 135, 437, 313, 647, 440, 237, 448, 443, 243, 445, 446, 447, 442, 449, 130, 451, 452, 453, 407, 45, 11, 457, 318, 284, 460, 461, 462, 463, 401, 465, 466, 467, 468, 469, 470, 471, 488, 473, 474, 532, 476, 371, 478, 256, 480, 481, 326, 483, 484, 412, 235, 487, 472, 489, 810, 875, 492, 552, 494, 495, 496, 497, 498, 769, 720, 501, 110, 81, 507, 505, 506, 504, 361, 590, 510, 180, 512, 513, 514, 515, 509, 247, 518, 519, 115, 521, 522, 523, 26, 508, 545, 527, 612, 529, 530, 531, 380, 716, 383, 535, 536, 164, 538, 539, 540, 261, 542, 543, 238, 526, 408, 670, 548, 799, 550, 636, 493, 553, 554, 555, 679, 557, 741, 559, 124, 561, 98, 563, 34, 565, 270, 567, 568, 569, 570, 571, 572, 573, 574, 575, 576, 75, 880, 579, 92, 581, 582, 583, 584, 120, 645, 587, 876, 155, 516, 334, 592, 503, 766, 595, 596, 597, 598, 29, 600, 601, 602, 603, 604, 605, 305, 607, 608, 533, 610, 436, 528, 846, 614, 629, 616, 617, 877, 619, 620, 621, 622, 623, 624, 625, 626, 188, 250, 615, 630, 244, 632, 388, 273, 748, 551, 236, 638, 639, 99, 641, 642, 643, 644, 586, 646, 439, 72, 319, 650, 651, 652, 653, 656, 444, 654, 657, 658, 659, 660, 807, 662, 663, 517, 665, 666, 295, 668, 669, 547, 671, 672, 66, 674, 128, 676, 677, 678, 556, 362, 681, 864, 282, 684, 685, 686, 687, 688, 689, 690, 324, 692, 693, 694, 695, 306, 697, 698, 725, 700, 780, 804, 703, 133, 705, 682, 707, 708, 709, 710, 711, 712, 93, 714, 7, 609, 717, 718, 683, 500, 721, 260, 242, 724, 793, 726, 727, 728, 738, 730, 731, 732, 733, 33, 735, 736, 737, 729, 739, 740, 558, 742, 113, 423, 792, 114, 747, 635, 749, 353, 751, 126, 753, 754, 578, 756, 251, 744, 759, 760, 761, 762, 763, 764, 765, 594, 770, 768, 499, 767, 227, 772, 67, 774, 775, 776, 777, 778, 228, 336, 186, 276, 783, 163, 40, 57, 787, 323, 789, 198, 791, 745, 699, 9, 795, 796, 797, 322, 549, 304, 801, 802, 329, 702, 805, 485, 661, 808, 809, 490, 811, 812, 813, 814, 815, 816, 817, 818, 819, 308, 821, 822, 823, 824, 825, 826, 827, 146, 829, 830, 831, 832, 833, 834, 835, 881, 837, 838, 839, 840, 841, 842, 843, 844, 39, 613, 847, 848, 849, 850, 258, 12, 853, 854, 855, 856, 857, 858, 84, 860, 861, 862, 863, 706, 865, 866, 867, 413, 138, 870, 268, 872, 873, 78, 491, 588, 618, 878, 879, 755, 127, 882, 883}},
		{[]int{96, 2, 671, 75, 140, 6, 7, 53, 9, 14, 782, 67, 19, 639, 60, 203, 17, 18, 896, 20, 21, 22, 403, 677, 811, 26, 71, 120, 29, 30, 132, 28, 33, 718, 206, 36, 135, 99, 516, 511, 41, 865, 476, 50, 45, 46, 5, 48, 470, 546, 410, 536, 558, 54, 523, 56, 57, 58, 149, 37, 136, 697, 531, 64, 65, 66, 173, 595, 153, 70, 230, 72, 187, 74, 220, 76, 77, 233, 79, 80, 473, 180, 83, 177, 250, 86, 303, 88, 89, 90, 91, 376, 585, 94, 765, 1, 349, 98, 433, 191, 459, 102, 103, 659, 105, 798, 714, 164, 835, 195, 498, 112, 668, 114, 469, 244, 117, 169, 307, 68, 121, 672, 627, 124, 134, 126, 446, 128, 368, 130, 131, 611, 133, 125, 184, 152, 157, 209, 139, 200, 141, 251, 143, 144, 145, 146, 147, 158, 59, 150, 816, 279, 69, 154, 155, 226, 101, 148, 159, 661, 630, 198, 163, 616, 573, 166, 468, 118, 168, 423, 171, 172, 713, 287, 175, 176, 84, 178, 104, 82, 873, 352, 367, 15, 185, 186, 73, 605, 189, 358, 502, 192, 34, 229, 345, 43, 197, 370, 421, 729, 721, 695, 818, 253, 205, 35, 207, 663, 212, 381, 211, 648, 851, 214, 215, 216, 217, 344, 219, 213, 332, 222, 223, 55, 225, 314, 227, 228, 194, 759, 231, 623, 179, 315, 174, 236, 237, 238, 239, 575, 241, 242, 393, 116, 615, 586, 85, 193, 357, 656, 308, 431, 204, 254, 426, 256, 257, 382, 259, 260, 261, 262, 263, 264, 882, 716, 267, 268, 269, 270, 571, 201, 429, 274, 275, 276, 277, 278, 488, 280, 487, 282, 353, 284, 378, 493, 235, 589, 734, 666, 481, 292, 293, 294, 777, 296, 248, 298, 594, 300, 301, 319, 707, 304, 305, 306, 119, 182, 309, 310, 311, 312, 313, 156, 122, 316, 317, 364, 396, 320, 321, 322, 323, 581, 821, 326, 327, 578, 329, 330, 631, 221, 333, 457, 545, 336, 337, 338, 778, 365, 757, 342, 458, 16, 110, 346, 347, 348, 453, 495, 351, 142, 283, 354, 355, 356, 249, 758, 359, 360, 547, 362, 363, 318, 730, 422, 183, 129, 369, 162, 521, 379, 392, 770, 375, 743, 188, 285, 372, 106, 271, 328, 383, 384, 81, 386, 387, 388, 389, 416, 52, 373, 243, 394, 842, 635, 397, 398, 399, 400, 401, 402, 23, 404, 556, 837, 407, 897, 341, 380, 411, 412, 445, 414, 415, 390, 417, 418, 419, 420, 199, 366, 115, 424, 290, 406, 427, 428, 273, 430, 245, 44, 651, 434, 435, 436, 437, 438, 439, 440, 441, 572, 769, 444, 413, 208, 395, 12, 449, 450, 451, 452, 97, 687, 455, 456, 32, 343, 137, 460, 652, 462, 463, 464, 465, 466, 862, 10, 499, 3, 471, 472, 385, 474, 95, 8, 741, 478, 479, 480, 775, 482, 903, 568, 726, 486, 694, 724, 489, 490, 491, 492, 286, 494, 361, 496, 497, 111, 234, 500, 501, 100, 503, 504, 505, 612, 507, 815, 509, 510, 40, 884, 513, 519, 515, 39, 517, 518, 514, 520, 266, 522, 224, 650, 525, 526, 527, 550, 529, 590, 63, 532, 533, 534, 535, 391, 537, 538, 539, 299, 541, 542, 543, 544, 335, 432, 350, 548, 549, 528, 551, 552, 570, 554, 555, 405, 61, 196, 559, 560, 773, 562, 563, 564, 704, 566, 806, 484, 853, 553, 210, 442, 557, 574, 810, 576, 577, 258, 579, 580, 324, 582, 583, 167, 93, 691, 587, 787, 288, 530, 647, 540, 593, 592, 247, 596, 597, 598, 599, 600, 601, 506, 603, 604, 377, 736, 607, 608, 609, 911, 31, 602, 850, 614, 751, 108, 617, 618, 619, 620, 621, 622, 232, 624, 161, 626, 123, 628, 629, 625, 331, 632, 633, 634, 302, 636, 637, 638, 584, 640, 641, 642, 643, 644, 645, 646, 591, 138, 649, 776, 38, 461, 653, 654, 655, 334, 657, 658, 78, 689, 160, 662, 127, 664, 665, 425, 667, 113, 669, 670, 49, 170, 673, 674, 675, 676, 24, 678, 679, 680, 681, 682, 683, 871, 685, 686, 807, 688, 660, 690, 246, 692, 693, 747, 202, 696, 62, 698, 699, 700, 701, 702, 703, 565, 705, 706, 796, 845, 709, 754, 711, 712, 448, 107, 715, 409, 717, 297, 719, 720, 722, 272, 723, 165, 725, 485, 727, 728, 47, 340, 731, 732, 733, 289, 735, 606, 737, 738, 739, 740, 477, 742, 92, 744, 745, 746, 281, 748, 749, 750, 454, 752, 753, 710, 755, 756, 371, 109, 27, 760, 854, 762, 763, 764, 475, 766, 767, 768, 443, 374, 771, 772, 561, 774, 291, 524, 295, 339, 779, 780, 781, 11, 783, 784, 785, 786, 588, 788, 789, 790, 791, 792, 793, 794, 795, 87, 797, 51, 878, 800, 801, 802, 803, 804, 805, 567, 252, 808, 809, 240, 13, 812, 813, 814, 508, 151, 817, 218, 819, 820, 325, 822, 895, 824, 825, 826, 827, 828, 829, 830, 831, 832, 833, 834, 190, 836, 255, 838, 839, 840, 841, 447, 843, 844, 708, 846, 847, 848, 849, 613, 4, 852, 569, 761, 855, 856, 857, 858, 859, 860, 861, 467, 863, 864, 879, 866, 867, 868, 869, 870, 684, 872, 181, 874, 875, 876, 877, 799, 42, 880, 881, 265, 883, 512, 885, 886, 887, 888, 889, 890, 891, 892, 893, 894, 823, 25, 408, 898, 899, 900, 901, 902, 483, 904, 905, 906, 907, 908, 909, 910, 610, 912, 913, 914, 915, 916}, []int{96, 2, 671, 75, 140, 6, 7, 53, 9, 14, 782, 67, 19, 639, 60, 203, 17, 18, 896, 20, 21, 22, 403, 677, 811, 26, 71, 120, 29, 30, 132, 28, 33, 718, 206, 36, 135, 99, 516, 511, 41, 865, 476, 50, 45, 46, 5, 48, 470, 546, 410, 536, 558, 54, 523, 56, 57, 58, 149, 37, 136, 697, 531, 64, 65, 66, 173, 595, 153, 70, 230, 72, 187, 74, 220, 76, 77, 233, 79, 80, 473, 180, 83, 177, 250, 86, 303, 88, 89, 90, 91, 376, 585, 94, 765, 1, 349, 98, 433, 191, 459, 102, 103, 659, 105, 798, 714, 164, 835, 195, 498, 112, 668, 114, 469, 244, 117, 169, 307, 68, 121, 672, 627, 124, 134, 126, 446, 128, 368, 130, 131, 611, 133, 125, 184, 152, 157, 209, 139, 200, 141, 251, 143, 144, 145, 146, 147, 158, 59, 150, 816, 279, 69, 154, 155, 226, 101, 148, 159, 661, 630, 198, 163, 616, 573, 166, 468, 118, 168, 423, 171, 172, 713, 287, 175, 176, 84, 178, 104, 82, 873, 352, 367, 15, 185, 186, 73, 605, 189, 358, 502, 192, 34, 229, 345, 43, 197, 370, 421, 729, 721, 695, 818, 253, 205, 35, 207, 663, 212, 381, 211, 648, 851, 214, 215, 216, 217, 344, 219, 213, 332, 222, 223, 55, 225, 314, 227, 228, 194, 759, 231, 623, 179, 315, 174, 236, 237, 238, 239, 575, 241, 242, 393, 116, 615, 586, 85, 193, 357, 656, 308, 431, 204, 254, 426, 256, 257, 382, 259, 260, 261, 262, 263, 264, 882, 716, 267, 268, 269, 270, 571, 201, 429, 274, 275, 276, 277, 278, 488, 280, 487, 282, 353, 284, 378, 493, 235, 589, 734, 666, 481, 292, 293, 294, 777, 296, 248, 298, 594, 300, 301, 319, 707, 304, 305, 306, 119, 182, 309, 310, 311, 312, 313, 156, 122, 316, 317, 364, 396, 320, 321, 322, 323, 581, 821, 326, 327, 578, 329, 330, 631, 221, 333, 457, 545, 336, 337, 338, 778, 365, 757, 342, 458, 16, 110, 346, 347, 348, 453, 495, 351, 142, 283, 354, 355, 356, 249, 758, 359, 360, 547, 362, 363, 318, 730, 422, 183, 129, 369, 162, 521, 379, 392, 770, 375, 743, 188, 285, 372, 106, 271, 328, 383, 384, 81, 386, 387, 388, 389, 416, 52, 373, 243, 394, 842, 635, 397, 398, 399, 400, 401, 402, 23, 404, 556, 837, 407, 897, 341, 380, 411, 412, 445, 414, 415, 390, 417, 418, 419, 420, 199, 366, 115, 424, 290, 406, 427, 428, 273, 430, 245, 44, 651, 434, 435, 436, 437, 438, 439, 440, 441, 572, 769, 444, 413, 208, 395, 12, 449, 450, 451, 452, 97, 687, 455, 456, 32, 343, 137, 460, 652, 462, 463, 464, 465, 466, 862, 10, 499, 3, 471, 472, 385, 474, 95, 8, 741, 478, 479, 480, 775, 482, 903, 568, 726, 486, 694, 724, 489, 490, 491, 492, 286, 494, 361, 496, 497, 111, 234, 500, 501, 100, 503, 504, 505, 612, 507, 815, 509, 510, 40, 884, 513, 519, 515, 39, 517, 518, 514, 520, 266, 522, 224, 650, 525, 526, 527, 550, 529, 590, 63, 532, 533, 534, 535, 391, 537, 538, 539, 299, 541, 542, 543, 544, 335, 432, 350, 548, 549, 528, 551, 552, 570, 554, 555, 405, 61, 196, 559, 560, 773, 562, 563, 564, 704, 566, 806, 484, 853, 553, 210, 442, 557, 574, 810, 576, 577, 258, 579, 580, 324, 582, 583, 167, 93, 691, 587, 787, 288, 530, 647, 540, 593, 592, 247, 596, 597, 598, 599, 600, 601, 506, 603, 604, 377, 736, 607, 608, 609, 911, 31, 602, 850, 614, 751, 108, 617, 618, 619, 620, 621, 622, 232, 624, 161, 626, 123, 628, 629, 625, 331, 632, 633, 634, 302, 636, 637, 638, 584, 640, 641, 642, 643, 644, 645, 646, 591, 138, 649, 776, 38, 461, 653, 654, 655, 334, 657, 658, 78, 689, 160, 662, 127, 664, 665, 425, 667, 113, 669, 670, 49, 170, 673, 674, 675, 676, 24, 678, 679, 680, 681, 682, 683, 871, 685, 686, 807, 688, 660, 690, 246, 692, 693, 747, 202, 696, 62, 698, 699, 700, 701, 702, 703, 565, 705, 706, 796, 845, 709, 754, 711, 712, 448, 107, 715, 409, 717, 297, 719, 720, 722, 272, 723, 165, 725, 485, 727, 728, 47, 340, 731, 732, 733, 289, 735, 606, 737, 738, 739, 740, 477, 742, 92, 744, 745, 746, 281, 748, 749, 750, 454, 752, 753, 710, 755, 756, 371, 109, 27, 760, 854, 762, 763, 764, 475, 766, 767, 768, 443, 374, 771, 772, 561, 774, 291, 524, 295, 339, 779, 780, 781, 11, 783, 784, 785, 786, 588, 788, 789, 790, 791, 792, 793, 794, 795, 87, 797, 51, 878, 800, 801, 802, 803, 804, 805, 567, 252, 808, 809, 240, 13, 812, 813, 814, 508, 151, 817, 218, 819, 820, 325, 822, 895, 824, 825, 826, 827, 828, 829, 830, 831, 832, 833, 834, 190, 836, 255, 838, 839, 840, 841, 447, 843, 844, 708, 846, 847, 848, 849, 613, 4, 852, 569, 761, 855, 856, 857, 858, 859, 860, 861, 467, 863, 864, 879, 866, 867, 868, 869, 870, 684, 872, 181, 874, 875, 876, 877, 799, 42, 880, 881, 265, 883, 512, 885, 886, 887, 888, 889, 890, 891, 892, 893, 894, 823, 25, 408, 898, 899, 900, 901, 902, 483, 904, 905, 906, 907, 908, 909, 910, 610, 912, 913, 914, 915, 916}},
		{[]int{14, 47, 75, 4, 42, 6, 61, 62, 16, 23, 5, 12, 91, 15, 10, 27, 17, 18, 19, 20, 21, 22, 88, 37, 35, 3, 9, 79, 29, 30, 40, 32, 33, 34, 25, 36, 24, 53, 39, 31, 95, 11, 69, 44, 45, 77, 2, 48, 49, 50, 134, 52, 1, 54, 55, 65, 57, 58, 59, 60, 7, 101, 99, 64, 56, 78, 67, 13, 43, 118, 107, 114, 116, 108, 96, 70, 46, 66, 28, 80, 81, 82, 83, 84, 85, 86, 87, 26, 89, 90, 68, 92, 93, 94, 41, 38, 97, 98, 63, 100, 8, 102, 103, 104, 105, 106, 71, 74, 109, 110, 111, 112, 113, 72, 115, 128, 117, 76, 119, 120, 121, 122, 123, 124, 125, 126, 127, 73, 129, 130, 131, 132, 133, 51}, []int{14, 47, 75, 4, 42, 6, 61, 62, 16, 23, 5, 12, 91, 15, 10, 27, 17, 18, 19, 20, 21, 22, 88, 37, 35, 3, 9, 79, 29, 30, 40, 32, 33, 34, 25, 36, 24, 53, 39, 31, 95, 11, 69, 44, 45, 77, 2, 48, 49, 50, 134, 52, 1, 54, 55, 65, 57, 58, 59, 60, 7, 101, 99, 64, 56, 78, 67, 13, 43, 118, 107, 114, 116, 108, 96, 70, 46, 66, 28, 80, 81, 82, 83, 84, 85, 86, 87, 26, 89, 90, 68, 92, 93, 94, 41, 38, 97, 98, 63, 100, 8, 102, 103, 104, 105, 106, 71, 74, 109, 110, 111, 112, 113, 72, 115, 128, 117, 76, 119, 120, 121, 122, 123, 124, 125, 126, 127, 73, 129, 130, 131, 132, 133, 51}},
	}

	for i, test := range testTable {
		t.Logf("%s No.%02d 結果\nTest3: %d\nTest4: %d", t.Name(), i+1, Test3(test.nums1), Test4(test.nums2))
	}
}
