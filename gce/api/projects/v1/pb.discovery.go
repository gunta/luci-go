// Code generated by cproto. DO NOT EDIT.

package projects

import discovery "go.chromium.org/luci/grpc/discovery"

import "github.com/golang/protobuf/protoc-gen-go/descriptor"

func init() {
	discovery.RegisterDescriptorSetCompressed(
		[]string{
			"projects.Projects",
		},
		[]byte{31, 139,
			8, 0, 0, 0, 0, 0, 0, 255, 220, 88, 75, 115, 27, 199,
			181, 198, 60, 4, 129, 45, 138, 228, 109, 82, 36, 53, 34, 165,
			99, 74, 178, 68, 25, 4, 172, 135, 109, 89, 146, 93, 30, 2,
			67, 114, 100, 16, 192, 157, 25, 136, 214, 221, 88, 195, 153, 6,
			48, 247, 2, 211, 112, 207, 128, 50, 175, 203, 190, 117, 215, 89,
			228, 39, 164, 42, 75, 239, 83, 149, 109, 42, 155, 44, 242, 35,
			242, 15, 82, 217, 165, 42, 85, 169, 84, 119, 207, 128, 32, 37,
			58, 177, 83, 217, 132, 27, 246, 233, 199, 233, 239, 124, 125, 230,
			60, 128, 126, 55, 131, 62, 232, 209, 74, 208, 103, 116, 24, 141,
			135, 21, 202, 122, 213, 193, 56, 136, 170, 189, 128, 84, 253, 81,
			84, 29, 49, 250, 223, 36, 72, 147, 234, 209, 253, 106, 64, 227,
			110, 212, 171, 140, 24, 77, 41, 46, 229, 43, 27, 14, 42, 214,
			196, 10, 94, 70, 197, 33, 73, 89, 20, 172, 42, 160, 221, 157,
			113, 50, 9, 175, 162, 139, 217, 238, 85, 21, 148, 187, 51, 78,
			46, 242, 19, 140, 244, 34, 26, 175, 106, 242, 132, 148, 54, 62,
			64, 23, 165, 206, 4, 223, 59, 57, 204, 181, 94, 122, 176, 80,
			201, 175, 174, 200, 61, 19, 117, 207, 255, 116, 17, 21, 177, 174,
			23, 12, 5, 253, 74, 65, 202, 44, 214, 244, 2, 126, 240, 189,
			2, 53, 58, 58, 102, 81, 175, 159, 194, 131, 247, 239, 127, 12,
			94, 159, 64, 163, 83, 179, 193, 28, 167, 125, 202, 146, 10, 152,
			131, 1, 136, 13, 9, 48, 146, 16, 118, 68, 194, 10, 130, 78,
			66, 128, 118, 33, 237, 71, 9, 36, 116, 204, 2, 2, 1, 13,
			9, 68, 9, 244, 232, 17, 97, 49, 9, 97, 28, 135, 132, 65,
			218, 39, 96, 142, 252, 128, 43, 142, 2, 18, 39, 164, 12, 47,
			8, 75, 34, 26, 195, 131, 202, 251, 8, 210, 190, 159, 66, 224,
			199, 112, 72, 160, 75, 199, 113, 8, 81, 44, 78, 53, 236, 154,
			213, 116, 45, 232, 70, 3, 82, 65, 168, 132, 20, 21, 107, 197,
			210, 2, 218, 71, 170, 94, 192, 250, 76, 97, 73, 49, 76, 48,
			65, 210, 15, 93, 202, 192, 135, 221, 90, 27, 50, 171, 33, 165,
			192, 200, 136, 178, 20, 118, 107, 22, 124, 53, 166, 169, 15, 227,
			52, 26, 68, 255, 235, 167, 252, 254, 46, 101, 21, 132, 16, 210,
			244, 130, 130, 181, 153, 210, 28, 250, 63, 164, 235, 5, 181, 128,
			181, 89, 117, 221, 96, 130, 142, 216, 31, 146, 132, 27, 123, 162,
			67, 190, 94, 114, 162, 191, 130, 160, 159, 166, 163, 228, 73, 181,
			26, 12, 232, 56, 172, 244, 40, 237, 13, 72, 37, 160, 195, 106,
			64, 135, 163, 113, 74, 170, 33, 13, 146, 42, 35, 93, 194, 72,
			28, 144, 42, 35, 73, 202, 125, 71, 190, 107, 82, 29, 68, 73,
			90, 65, 104, 22, 93, 224, 0, 116, 142, 96, 34, 93, 192, 218,
			236, 165, 255, 200, 37, 5, 107, 179, 120, 53, 151, 52, 172, 205,
			94, 91, 67, 251, 2, 184, 130, 181, 121, 245, 138, 241, 217, 4,
			56, 199, 253, 83, 56, 185, 44, 148, 43, 58, 214, 231, 213, 217,
			245, 236, 46, 229, 2, 87, 95, 202, 37, 126, 217, 204, 66, 46,
			105, 88, 155, 95, 92, 66, 199, 2, 135, 138, 181, 69, 117, 221,
			24, 188, 73, 96, 102, 238, 191, 142, 58, 85, 231, 119, 79, 164,
			11, 88, 91, 156, 80, 199, 249, 89, 156, 80, 167, 106, 88, 91,
			188, 182, 134, 218, 72, 213, 21, 172, 175, 20, 12, 197, 168, 159,
			118, 167, 132, 28, 17, 230, 15, 114, 242, 146, 31, 225, 81, 156,
			158, 149, 210, 60, 250, 12, 233, 186, 194, 61, 234, 170, 122, 221,
			120, 40, 8, 249, 145, 218, 56, 90, 69, 184, 196, 85, 117, 34,
			21, 177, 118, 53, 179, 75, 17, 46, 113, 21, 95, 205, 37, 13,
			107, 87, 215, 214, 15, 139, 34, 30, 61, 68, 191, 198, 232, 154,
			36, 181, 42, 102, 14, 199, 221, 42, 25, 142, 210, 227, 44, 96,
			205, 103, 140, 231, 139, 27, 23, 209, 5, 139, 175, 111, 31, 161,
			197, 128, 14, 43, 103, 214, 183, 145, 88, 109, 115, 177, 173, 252,
			215, 157, 94, 148, 246, 199, 135, 226, 193, 122, 116, 224, 199, 189,
			147, 107, 70, 233, 241, 136, 36, 242, 182, 63, 43, 202, 47, 85,
			109, 183, 189, 253, 189, 122, 125, 87, 106, 108, 103, 251, 42, 7,
			100, 48, 248, 60, 166, 175, 99, 143, 239, 127, 254, 215, 5, 30,
			173, 174, 23, 30, 46, 160, 223, 207, 138, 104, 117, 189, 128, 31,
			252, 102, 22, 196, 129, 128, 14, 96, 123, 220, 237, 18, 150, 192,
			22, 72, 85, 119, 18, 8, 253, 212, 135, 40, 78, 9, 11, 250,
			126, 220, 227, 193, 132, 13, 253, 20, 157, 10, 113, 239, 63, 206,
			14, 128, 29, 7, 21, 56, 39, 186, 229, 14, 25, 146, 35, 50,
			160, 35, 194, 146, 105, 175, 28, 101, 32, 182, 14, 37, 136, 42,
			66, 224, 144, 48, 74, 82, 22, 29, 142, 197, 195, 249, 113, 8,
			227, 132, 240, 72, 150, 69, 71, 62, 115, 24, 197, 62, 59, 22,
			184, 146, 50, 188, 142, 210, 62, 80, 38, 254, 211, 113, 138, 96,
			72, 195, 168, 27, 5, 226, 233, 203, 224, 51, 2, 35, 194, 134,
			81, 154, 146, 144, 187, 204, 81, 20, 146, 80, 70, 75, 30, 29,
			187, 116, 48, 160, 175, 163, 184, 199, 253, 53, 140, 82, 241, 85,
			249, 140, 32, 30, 157, 158, 32, 4, 252, 239, 222, 25, 96, 226,
			51, 156, 142, 215, 195, 113, 146, 2, 35, 169, 159, 197, 92, 255,
			144, 30, 241, 165, 140, 49, 4, 49, 77, 163, 128, 148, 101, 164,
			231, 31, 26, 215, 48, 125, 99, 28, 158, 129, 19, 70, 73, 48,
			240, 163, 33, 97, 149, 243, 64, 68, 241, 52, 23, 57, 136, 17,
			163, 225, 56, 32, 39, 56, 208, 9, 144, 127, 10, 7, 202, 51,
			74, 72, 131, 241, 144, 196, 169, 159, 63, 82, 149, 50, 160, 105,
			159, 48, 24, 250, 41, 97, 145, 63, 72, 78, 168, 22, 15, 148,
			246, 9, 130, 105, 244, 19, 163, 154, 36, 18, 39, 211, 169, 80,
			59, 237, 91, 49, 61, 89, 19, 188, 71, 105, 194, 45, 138, 165,
			42, 202, 18, 24, 250, 199, 60, 235, 141, 19, 254, 176, 20, 72,
			28, 82, 198, 19, 43, 227, 32, 134, 52, 21, 161, 34, 28, 243,
			80, 17, 18, 22, 29, 145, 16, 186, 140, 14, 81, 158, 119, 187,
			233, 107, 238, 38, 153, 7, 65, 50, 34, 1, 247, 32, 24, 177,
			136, 59, 22, 227, 190, 19, 75, 47, 74, 18, 129, 29, 129, 183,
			103, 187, 224, 182, 118, 188, 3, 211, 177, 192, 118, 161, 237, 180,
			94, 216, 117, 171, 14, 219, 47, 193, 219, 179, 160, 214, 106, 191,
			116, 236, 221, 61, 15, 246, 90, 141, 186, 229, 184, 96, 54, 235,
			80, 107, 53, 61, 199, 222, 238, 120, 45, 199, 69, 176, 97, 186,
			96, 187, 27, 98, 197, 108, 190, 4, 235, 139, 182, 99, 185, 46,
			180, 28, 176, 247, 219, 13, 219, 170, 195, 129, 233, 56, 102, 211,
			179, 45, 183, 12, 118, 179, 214, 232, 212, 237, 230, 110, 25, 182,
			59, 30, 52, 91, 30, 130, 134, 189, 111, 123, 86, 29, 188, 86,
			89, 92, 251, 230, 57, 104, 237, 192, 190, 229, 212, 246, 204, 166,
			103, 110, 219, 13, 219, 123, 41, 46, 220, 177, 189, 38, 191, 108,
			167, 229, 32, 48, 161, 109, 58, 158, 93, 235, 52, 76, 7, 218,
			29, 167, 221, 114, 45, 224, 150, 213, 109, 183, 214, 48, 237, 125,
			171, 94, 1, 187, 9, 205, 22, 88, 47, 172, 166, 7, 238, 158,
			217, 104, 156, 54, 20, 65, 235, 160, 105, 57, 28, 253, 180, 153,
			176, 109, 65, 195, 54, 183, 27, 22, 191, 74, 216, 89, 183, 29,
			171, 230, 113, 131, 78, 70, 53, 187, 110, 53, 61, 179, 81, 70,
			224, 182, 173, 154, 109, 54, 202, 96, 125, 97, 237, 183, 27, 166,
			243, 178, 156, 41, 117, 173, 255, 236, 88, 77, 207, 54, 27, 80,
			55, 247, 205, 93, 203, 133, 187, 127, 143, 149, 182, 211, 170, 117,
			28, 107, 159, 163, 110, 237, 128, 219, 217, 118, 61, 219, 235, 120,
			22, 236, 182, 90, 117, 65, 182, 107, 57, 47, 236, 154, 229, 62,
			133, 70, 203, 21, 132, 117, 92, 171, 140, 160, 110, 122, 166, 184,
			186, 237, 180, 118, 108, 207, 125, 202, 199, 219, 29, 215, 22, 196,
			217, 77, 207, 114, 156, 78, 219, 179, 91, 205, 77, 216, 107, 29,
			88, 47, 44, 7, 106, 102, 199, 181, 234, 130, 225, 86, 147, 91,
			203, 125, 197, 106, 57, 47, 185, 90, 206, 131, 120, 129, 50, 28,
			236, 89, 222, 158, 229, 112, 82, 5, 91, 38, 167, 193, 245, 28,
			187, 230, 77, 111, 107, 57, 224, 181, 28, 15, 77, 217, 9, 77,
			107, 183, 97, 239, 90, 205, 154, 197, 151, 91, 92, 205, 129, 237,
			90, 155, 96, 58, 182, 203, 55, 216, 226, 98, 56, 48, 95, 66,
			171, 35, 172, 230, 15, 213, 113, 45, 36, 199, 83, 174, 91, 22,
			239, 9, 246, 14, 152, 245, 23, 54, 71, 158, 237, 110, 183, 92,
			215, 206, 220, 69, 208, 86, 219, 203, 56, 159, 20, 145, 80, 90,
			225, 163, 18, 214, 54, 10, 79, 209, 12, 82, 75, 183, 229, 80,
			78, 222, 44, 124, 42, 38, 47, 201, 161, 156, 188, 85, 40, 139,
			73, 69, 14, 229, 228, 237, 194, 123, 98, 50, 27, 202, 201, 119,
			11, 27, 98, 18, 201, 161, 156, 188, 83, 120, 71, 76, 222, 146,
			67, 57, 121, 183, 112, 67, 76, 222, 144, 195, 191, 168, 162, 176,
			213, 30, 22, 22, 140, 63, 170, 96, 66, 143, 196, 132, 69, 1,
			136, 252, 9, 67, 146, 36, 126, 143, 200, 20, 112, 76, 199, 162,
			104, 102, 100, 139, 39, 154, 148, 130, 127, 68, 163, 16, 66, 210,
			141, 98, 17, 254, 198, 163, 1, 79, 38, 36, 68, 167, 207, 139,
			240, 123, 76, 199, 12, 204, 182, 205, 11, 124, 72, 143, 71, 81,
			224, 15, 128, 124, 237, 15, 71, 3, 81, 196, 167, 84, 230, 175,
			20, 252, 68, 68, 49, 70, 190, 26, 147, 36, 69, 144, 69, 53,
			70, 146, 17, 141, 249, 205, 199, 35, 89, 101, 198, 92, 31, 79,
			62, 125, 26, 86, 96, 135, 50, 136, 226, 36, 245, 227, 128, 228,
			217, 136, 231, 215, 40, 32, 176, 67, 41, 124, 35, 167, 0, 216,
			40, 128, 109, 159, 221, 61, 83, 100, 84, 68, 141, 177, 201, 115,
			211, 152, 197, 9, 156, 179, 254, 84, 170, 249, 150, 7, 182, 62,
			129, 231, 110, 171, 41, 50, 9, 73, 38, 97, 158, 215, 113, 175,
			196, 238, 87, 220, 50, 201, 133, 216, 72, 15, 69, 69, 252, 234,
			155, 111, 95, 77, 181, 1, 15, 75, 151, 39, 165, 211, 111, 151,
			208, 135, 255, 112, 23, 152, 153, 119, 182, 13, 52, 126, 168, 248,
			50, 126, 90, 147, 185, 113, 3, 93, 174, 147, 1, 73, 137, 35,
			223, 5, 207, 33, 53, 10, 87, 21, 209, 68, 170, 81, 184, 241,
			57, 186, 108, 197, 201, 152, 157, 183, 97, 186, 123, 228, 173, 231,
			15, 117, 143, 27, 107, 8, 237, 146, 244, 188, 171, 202, 232, 82,
			35, 74, 38, 203, 235, 8, 141, 252, 30, 249, 50, 165, 255, 67,
			226, 108, 219, 12, 159, 241, 248, 196, 70, 136, 102, 229, 110, 233,
			63, 184, 140, 38, 76, 157, 219, 198, 78, 118, 224, 119, 209, 124,
			76, 190, 78, 191, 156, 186, 65, 54, 206, 151, 249, 116, 59, 191,
			229, 193, 31, 20, 84, 106, 231, 135, 62, 70, 69, 73, 22, 94,
			57, 81, 125, 138, 62, 99, 249, 108, 145, 43, 253, 11, 63, 68,
			69, 73, 227, 244, 209, 83, 196, 26, 111, 192, 197, 91, 72, 219,
			37, 41, 94, 58, 89, 56, 97, 239, 45, 219, 63, 64, 58, 103,
			4, 95, 57, 89, 153, 226, 211, 88, 62, 59, 45, 137, 123, 254,
			139, 121, 217, 210, 127, 242, 239, 208, 210, 207, 32, 85, 43, 96,
			173, 116, 241, 150, 24, 242, 110, 252, 226, 167, 168, 44, 27, 253,
			217, 194, 188, 98, 0, 152, 121, 16, 226, 209, 41, 20, 207, 7,
			126, 222, 76, 77, 125, 192, 179, 165, 43, 168, 154, 247, 241, 115,
			234, 130, 177, 33, 56, 136, 66, 105, 30, 153, 110, 133, 165, 154,
			73, 195, 91, 208, 177, 62, 167, 206, 174, 76, 53, 222, 115, 147,
			134, 151, 235, 158, 155, 185, 52, 213, 120, 207, 205, 205, 163, 15,
			101, 247, 136, 11, 43, 138, 113, 239, 52, 198, 128, 17, 63, 21,
			229, 220, 120, 20, 250, 111, 65, 203, 123, 68, 156, 161, 21, 61,
			226, 210, 15, 163, 37, 194, 241, 50, 180, 162, 37, 212, 151, 84,
			188, 146, 247, 125, 23, 184, 130, 210, 84, 79, 184, 148, 161, 149,
			61, 225, 210, 220, 60, 186, 38, 110, 82, 176, 182, 172, 94, 49,
			230, 166, 187, 209, 137, 86, 222, 244, 47, 171, 75, 11, 217, 73,
			165, 200, 55, 231, 90, 57, 226, 229, 153, 201, 154, 134, 181, 229,
			197, 37, 244, 16, 169, 186, 138, 117, 163, 176, 174, 24, 119, 78,
			115, 208, 35, 41, 79, 13, 228, 235, 40, 73, 121, 90, 58, 77,
			0, 71, 98, 148, 48, 42, 35, 93, 87, 57, 1, 107, 234, 130,
			113, 227, 124, 2, 122, 36, 199, 169, 10, 235, 215, 84, 99, 73,
			96, 81, 133, 245, 107, 25, 78, 85, 88, 191, 150, 89, 175, 10,
			235, 215, 230, 230, 209, 3, 164, 234, 26, 214, 161, 112, 75, 49,
			222, 61, 141, 83, 180, 20, 103, 81, 38, 25, 76, 238, 143, 80,
			90, 68, 63, 87, 144, 174, 107, 28, 231, 77, 117, 213, 248, 127,
			69, 0, 61, 242, 7, 99, 241, 229, 156, 9, 76, 192, 72, 64,
			68, 205, 30, 197, 224, 195, 244, 183, 91, 225, 31, 91, 152, 211,
			35, 90, 12, 126, 24, 248, 97, 174, 105, 114, 61, 216, 93, 153,
			176, 202, 124, 167, 204, 196, 221, 136, 37, 114, 107, 70, 133, 38,
			168, 184, 169, 194, 21, 97, 174, 38, 168, 184, 153, 81, 161, 9,
			42, 110, 206, 44, 230, 146, 134, 181, 155, 203, 43, 232, 35, 164,
			234, 58, 214, 239, 20, 42, 138, 241, 158, 160, 34, 207, 231, 20,
			252, 55, 136, 57, 195, 135, 174, 96, 237, 78, 105, 9, 173, 33,
			93, 215, 57, 29, 155, 234, 13, 99, 254, 212, 111, 27, 217, 239,
			22, 186, 248, 221, 98, 83, 157, 72, 69, 172, 109, 102, 191, 91,
			232, 2, 218, 38, 54, 114, 73, 195, 218, 230, 250, 117, 244, 51,
			69, 168, 85, 176, 182, 165, 174, 27, 223, 77, 145, 156, 85, 37,
			89, 73, 50, 69, 245, 20, 195, 167, 60, 79, 244, 95, 83, 204,
			162, 183, 82, 203, 219, 56, 34, 26, 238, 152, 194, 144, 178, 83,
			86, 92, 22, 216, 248, 71, 177, 165, 110, 222, 200, 160, 42, 23,
			56, 184, 82, 46, 113, 168, 51, 171, 185, 164, 97, 109, 235, 218,
			26, 119, 182, 98, 1, 235, 247, 11, 159, 72, 103, 203, 171, 31,
			94, 143, 12, 253, 56, 26, 141, 7, 254, 91, 156, 173, 200, 73,
			185, 95, 90, 64, 31, 33, 189, 40, 66, 216, 35, 245, 153, 113,
			15, 100, 190, 202, 2, 86, 50, 29, 74, 192, 142, 83, 194, 98,
			127, 192, 11, 48, 201, 123, 81, 70, 171, 71, 197, 217, 92, 82,
			177, 246, 232, 178, 145, 75, 26, 214, 30, 221, 126, 130, 190, 19,
			87, 40, 88, 123, 172, 110, 25, 95, 129, 204, 107, 89, 148, 153,
			186, 66, 126, 23, 73, 5, 65, 77, 68, 52, 190, 20, 147, 215,
			147, 229, 73, 128, 75, 78, 125, 234, 52, 22, 79, 21, 147, 128,
			23, 158, 236, 248, 60, 164, 156, 190, 199, 19, 164, 60, 35, 60,
			158, 32, 229, 100, 62, 190, 253, 30, 143, 144, 69, 241, 179, 226,
			19, 245, 142, 177, 1, 187, 36, 157, 84, 134, 111, 143, 46, 242,
			56, 183, 237, 73, 113, 38, 151, 248, 121, 180, 148, 75, 26, 214,
			158, 220, 184, 141, 42, 66, 181, 134, 181, 103, 106, 213, 120, 71,
			120, 209, 68, 247, 219, 2, 130, 60, 205, 67, 194, 179, 226, 68,
			82, 177, 246, 236, 210, 114, 46, 113, 101, 239, 108, 229, 85, 228,
			223, 2, 0, 0, 255, 255, 148, 14, 173, 138, 76, 24, 0, 0,
		},
	)
}

// FileDescriptorSet returns a descriptor set for this proto package, which
// includes all defined services, and all transitive dependencies.
//
// Will not return nil.
//
// Do NOT modify the returned descriptor.
func FileDescriptorSet() *descriptor.FileDescriptorSet {
	// We just need ONE of the service names to look up the FileDescriptorSet.
	ret, err := discovery.GetDescriptorSet("projects.Projects")
	if err != nil {
		panic(err)
	}
	return ret
}
