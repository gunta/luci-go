// Code generated by cproto. DO NOT EDIT.

package instances

import discovery "go.chromium.org/luci/grpc/discovery"

import "github.com/golang/protobuf/protoc-gen-go/descriptor"

func init() {
	discovery.RegisterDescriptorSetCompressed(
		[]string{
			"instances.Instances",
		},
		[]byte{31, 139,
			8, 0, 0, 0, 0, 0, 0, 255, 236, 122, 207, 115, 27, 71,
			118, 63, 102, 6, 252, 213, 150, 245, 99, 68, 201, 44, 104, 87,
			126, 11, 155, 18, 160, 5, 7, 252, 33, 217, 38, 233, 253, 250,
			11, 130, 67, 114, 28, 16, 96, 0, 80, 92, 201, 229, 178, 134,
			51, 13, 96, 214, 192, 52, 118, 186, 65, 138, 82, 20, 39, 169,
			218, 164, 42, 135, 173, 84, 206, 57, 229, 144, 195, 94, 114, 72,
			182, 42, 215, 220, 115, 206, 127, 144, 115, 174, 169, 74, 85, 42,
			245, 186, 103, 6, 32, 69, 90, 155, 77, 85, 78, 102, 89, 229,
			233, 158, 238, 247, 227, 211, 111, 222, 251, 116, 55, 200, 63, 153,
			228, 94, 151, 177, 110, 159, 150, 135, 17, 19, 236, 120, 212, 41,
			211, 193, 80, 156, 89, 178, 105, 222, 80, 47, 173, 228, 101, 126,
			134, 76, 217, 248, 126, 235, 132, 220, 246, 216, 192, 186, 240, 126,
			139, 200, 183, 7, 216, 60, 208, 158, 63, 236, 6, 162, 55, 58,
			182, 60, 54, 40, 119, 89, 223, 13, 187, 99, 53, 67, 113, 54,
			164, 92, 105, 251, 15, 77, 251, 91, 221, 216, 61, 216, 250, 141,
			126, 127, 87, 73, 60, 136, 199, 89, 71, 180, 223, 255, 131, 144,
			157, 134, 109, 28, 255, 229, 127, 221, 36, 211, 102, 246, 126, 102,
			237, 38, 249, 151, 107, 68, 187, 102, 26, 247, 51, 230, 234, 63,
			95, 3, 57, 193, 99, 125, 216, 26, 117, 58, 52, 226, 176, 4,
			74, 212, 67, 14, 190, 43, 92, 8, 66, 65, 35, 175, 231, 134,
			93, 10, 29, 22, 13, 92, 65, 160, 202, 134, 103, 81, 208, 237,
			9, 88, 93, 94, 254, 44, 158, 0, 78, 232, 89, 0, 149, 126,
			31, 228, 59, 14, 17, 229, 52, 58, 161, 190, 69, 160, 39, 196,
			144, 111, 148, 203, 62, 61, 161, 125, 54, 164, 17, 79, 48, 64,
			39, 135, 177, 17, 75, 199, 202, 136, 50, 33, 208, 164, 126, 192,
			69, 20, 28, 143, 68, 192, 66, 112, 67, 31, 70, 156, 66, 16,
			2, 103, 163, 200, 163, 178, 231, 56, 8, 221, 232, 76, 218, 197,
			75, 112, 26, 136, 30, 176, 72, 254, 159, 141, 4, 129, 1, 243,
			131, 78, 224, 185, 40, 161, 4, 110, 68, 97, 72, 163, 65, 32,
			4, 245, 97, 24, 177, 147, 192, 167, 62, 136, 158, 43, 64, 244,
			208, 187, 126, 159, 157, 6, 97, 23, 60, 22, 250, 1, 78, 226,
			56, 137, 192, 128, 138, 13, 66, 0, 255, 30, 93, 48, 140, 3,
			235, 36, 22, 121, 204, 167, 48, 24, 113, 1, 17, 21, 110, 16,
			74, 169, 238, 49, 59, 193, 87, 49, 98, 4, 66, 38, 2, 143,
			150, 64, 244, 2, 14, 253, 128, 11, 148, 48, 169, 49, 244, 47,
			152, 227, 7, 220, 235, 187, 193, 128, 70, 214, 85, 70, 4, 225,
			36, 22, 137, 17, 195, 136, 249, 35, 143, 142, 237, 32, 99, 67,
			254, 87, 118, 16, 136, 189, 243, 153, 55, 26, 208, 80, 184, 201,
			34, 149, 89, 4, 76, 244, 104, 4, 3, 87, 208, 40, 112, 251,
			124, 12, 181, 92, 32, 209, 163, 4, 38, 173, 79, 157, 170, 211,
			64, 206, 68, 193, 161, 59, 160, 104, 208, 100, 108, 133, 108, 252,
			78, 226, 30, 8, 142, 30, 133, 74, 20, 139, 56, 12, 220, 51,
			56, 166, 24, 41, 62, 8, 6, 52, 244, 89, 196, 41, 6, 197,
			48, 98, 3, 38, 40, 40, 76, 4, 7, 159, 70, 193, 9, 245,
			161, 19, 177, 1, 81, 40, 112, 214, 17, 167, 24, 38, 113, 4,
			1, 31, 82, 15, 35, 8, 134, 81, 128, 129, 21, 97, 236, 132,
			42, 138, 56, 151, 182, 19, 104, 239, 57, 45, 104, 53, 118, 218,
			71, 149, 166, 13, 78, 11, 14, 154, 141, 167, 206, 182, 189, 13,
			91, 207, 160, 189, 103, 67, 181, 113, 240, 172, 233, 236, 238, 181,
			97, 175, 81, 219, 182, 155, 45, 168, 212, 183, 161, 218, 168, 183,
			155, 206, 214, 97, 187, 209, 108, 17, 200, 87, 90, 224, 180, 242,
			242, 77, 165, 254, 12, 236, 159, 31, 52, 237, 86, 11, 26, 77,
			112, 246, 15, 106, 142, 189, 13, 71, 149, 102, 179, 82, 111, 59,
			118, 171, 4, 78, 189, 90, 59, 220, 118, 234, 187, 37, 216, 58,
			108, 67, 189, 209, 38, 80, 115, 246, 157, 182, 189, 13, 237, 70,
			73, 170, 125, 123, 30, 52, 118, 96, 223, 110, 86, 247, 42, 245,
			118, 101, 203, 169, 57, 237, 103, 82, 225, 142, 211, 174, 163, 178,
			157, 70, 147, 64, 5, 14, 42, 205, 182, 83, 61, 172, 85, 154,
			112, 112, 216, 60, 104, 180, 108, 64, 207, 182, 157, 86, 181, 86,
			113, 246, 237, 109, 11, 156, 58, 212, 27, 96, 63, 181, 235, 109,
			104, 237, 85, 106, 181, 243, 142, 18, 104, 28, 213, 237, 38, 90,
			63, 233, 38, 108, 217, 80, 115, 42, 91, 53, 27, 85, 73, 63,
			183, 157, 166, 93, 109, 163, 67, 227, 167, 170, 179, 109, 215, 219,
			149, 90, 137, 64, 235, 192, 174, 58, 149, 90, 9, 236, 159, 219,
			251, 7, 181, 74, 243, 89, 41, 22, 218, 178, 255, 240, 208, 174,
			183, 157, 74, 13, 182, 43, 251, 149, 93, 187, 5, 133, 119, 161,
			114, 208, 108, 84, 15, 155, 246, 62, 90, 221, 216, 129, 214, 225,
			86, 171, 237, 180, 15, 219, 54, 236, 54, 26, 219, 18, 236, 150,
			221, 124, 234, 84, 237, 214, 38, 212, 26, 45, 9, 216, 97, 203,
			46, 17, 216, 174, 180, 43, 82, 245, 65, 179, 177, 227, 180, 91,
			155, 248, 188, 117, 216, 114, 36, 112, 78, 189, 109, 55, 155, 135,
			7, 109, 167, 81, 47, 194, 94, 227, 200, 126, 106, 55, 161, 90,
			57, 108, 217, 219, 18, 225, 70, 29, 189, 197, 88, 177, 27, 205,
			103, 40, 22, 113, 144, 43, 80, 130, 163, 61, 187, 189, 103, 55,
			17, 84, 137, 86, 5, 97, 104, 181, 155, 78, 181, 61, 57, 172,
			209, 132, 118, 163, 217, 38, 19, 126, 66, 221, 222, 173, 57, 187,
			118, 189, 106, 227, 235, 6, 138, 57, 114, 90, 118, 17, 42, 77,
			167, 133, 3, 28, 169, 24, 142, 42, 207, 160, 113, 40, 189, 198,
			133, 58, 108, 217, 68, 61, 79, 132, 110, 73, 174, 39, 56, 59,
			80, 217, 126, 234, 160, 229, 241, 232, 131, 70, 171, 229, 196, 225,
			34, 97, 171, 238, 197, 152, 91, 132, 204, 18, 77, 55, 13, 152,
			253, 0, 159, 102, 77, 35, 159, 217, 36, 115, 68, 159, 93, 84,
			143, 170, 243, 163, 204, 255, 147, 157, 239, 169, 71, 213, 249, 113,
			166, 36, 59, 53, 245, 168, 58, 23, 51, 63, 149, 157, 241, 163,
			234, 124, 144, 201, 203, 78, 162, 30, 85, 231, 195, 204, 79, 100,
			231, 199, 234, 81, 117, 22, 50, 31, 202, 206, 15, 213, 227, 127,
			234, 68, 207, 102, 76, 99, 45, 115, 51, 247, 239, 58, 84, 160,
			75, 67, 26, 5, 30, 200, 250, 9, 3, 202, 185, 219, 165, 170,
			4, 156, 177, 17, 120, 110, 8, 17, 93, 194, 66, 35, 24, 184,
			39, 44, 240, 193, 167, 157, 32, 148, 233, 111, 52, 236, 99, 49,
			161, 62, 57, 63, 95, 166, 223, 51, 54, 138, 160, 114, 224, 112,
			11, 42, 32, 206, 134, 129, 231, 246, 129, 190, 116, 7, 195, 62,
			133, 128, 163, 60, 89, 191, 4, 184, 92, 102, 177, 136, 254, 114,
			68, 185, 32, 16, 103, 181, 136, 242, 33, 11, 81, 243, 217, 80,
			166, 62, 55, 68, 121, 88, 124, 122, 204, 183, 96, 135, 69, 16,
			132, 92, 184, 161, 71, 147, 106, 132, 245, 53, 240, 40, 236, 48,
			6, 175, 85, 23, 64, 52, 244, 96, 203, 141, 10, 23, 72, 134,
			37, 57, 70, 17, 107, 211, 40, 10, 57, 92, 241, 126, 83, 137,
			121, 131, 137, 173, 71, 225, 203, 86, 163, 46, 43, 9, 229, 105,
			154, 239, 176, 8, 94, 200, 209, 47, 208, 51, 133, 133, 28, 200,
			142, 127, 65, 61, 1, 47, 94, 191, 121, 97, 17, 66, 136, 145,
			205, 104, 166, 177, 54, 251, 254, 241, 180, 84, 179, 70, 254, 190,
			76, 62, 188, 72, 157, 68, 48, 160, 92, 184, 131, 225, 85, 244,
			105, 147, 204, 181, 147, 49, 230, 2, 153, 225, 20, 235, 20, 95,
			208, 64, 43, 24, 205, 164, 105, 206, 147, 169, 208, 13, 25, 95,
			208, 65, 43, 76, 53, 85, 99, 235, 143, 47, 167, 92, 215, 83,
			137, 9, 237, 250, 233, 187, 105, 87, 106, 233, 255, 128, 122, 253,
			118, 137, 204, 152, 83, 247, 51, 127, 174, 105, 63, 112, 175, 31,
			184, 215, 15, 220, 235, 7, 238, 245, 3, 247, 250, 129, 123, 253,
			31, 114, 175, 148, 18, 225, 99, 194, 189, 182, 18, 66, 134, 143,
			9, 247, 74, 9, 217, 98, 74, 200, 30, 100, 202, 9, 33, 195,
			199, 132, 123, 165, 132, 236, 97, 74, 200, 10, 99, 66, 134, 143,
			255, 120, 79, 114, 175, 169, 87, 88, 249, 114, 127, 119, 15, 42,
			144, 150, 220, 49, 163, 224, 224, 194, 144, 5, 161, 144, 89, 45,
			24, 96, 149, 241, 233, 144, 134, 62, 13, 133, 98, 65, 103, 170,
			255, 21, 11, 169, 36, 75, 158, 219, 167, 161, 239, 70, 165, 177,
			20, 234, 35, 171, 138, 121, 128, 204, 158, 157, 200, 245, 198, 53,
			34, 121, 129, 37, 0, 73, 129, 108, 99, 141, 100, 125, 85, 226,
			130, 16, 14, 219, 85, 176, 135, 204, 235, 73, 117, 22, 56, 66,
			146, 155, 16, 43, 11, 214, 63, 204, 194, 50, 127, 30, 68, 172,
			79, 135, 34, 240, 96, 55, 162, 93, 22, 5, 110, 8, 213, 216,
			38, 56, 237, 5, 94, 15, 232, 75, 65, 81, 33, 102, 204, 241,
			160, 196, 112, 2, 199, 174, 247, 237, 169, 27, 249, 146, 22, 158,
			81, 55, 2, 22, 190, 165, 210, 229, 124, 52, 64, 173, 110, 191,
			15, 131, 32, 28, 9, 42, 107, 34, 124, 178, 76, 82, 151, 250,
			44, 236, 150, 32, 176, 168, 5, 125, 234, 14, 199, 174, 70, 20,
			242, 124, 64, 221, 136, 250, 121, 224, 76, 149, 218, 144, 77, 142,
			34, 32, 220, 99, 197, 78, 67, 74, 81, 101, 71, 114, 76, 65,
			163, 33, 86, 81, 89, 32, 160, 41, 233, 71, 192, 227, 100, 189,
			188, 188, 188, 178, 36, 255, 107, 47, 47, 111, 200, 255, 158, 163,
			23, 235, 235, 235, 235, 75, 43, 171, 75, 107, 43, 237, 213, 181,
			141, 39, 235, 27, 79, 214, 173, 245, 228, 239, 185, 69, 96, 235,
			12, 1, 23, 81, 224, 9, 9, 101, 108, 82, 132, 226, 75, 112,
			74, 129, 134, 124, 20, 197, 100, 252, 148, 74, 46, 238, 177, 240,
			132, 70, 2, 4, 35, 241, 170, 178, 1, 64, 115, 167, 10, 107,
			107, 107, 235, 72, 146, 40, 160, 200, 176, 203, 45, 2, 45, 74,
			225, 171, 132, 237, 156, 158, 158, 90, 1, 21, 29, 139, 69, 221,
			114, 212, 241, 240, 31, 78, 178, 196, 75, 241, 117, 225, 119, 25,
			85, 196, 2, 243, 17, 216, 138, 195, 115, 66, 146, 71, 88, 217,
			128, 42, 27, 12, 71, 130, 78, 132, 180, 180, 237, 160, 209, 114,
			126, 14, 47, 48, 130, 10, 69, 228, 192, 178, 188, 142, 7, 165,
			4, 50, 166, 217, 99, 234, 203, 169, 248, 38, 94, 188, 130, 156,
			94, 63, 172, 213, 138, 197, 75, 199, 201, 24, 46, 44, 23, 55,
			39, 108, 90, 125, 151, 77, 93, 42, 80, 10, 235, 248, 238, 217,
			132, 109, 92, 68, 35, 79, 72, 5, 39, 110, 31, 196, 73, 172,
			241, 220, 240, 7, 226, 164, 4, 210, 160, 205, 223, 215, 165, 19,
			75, 156, 96, 235, 251, 60, 82, 131, 70, 156, 122, 240, 8, 86,
			150, 151, 207, 123, 184, 118, 165, 135, 71, 65, 184, 182, 10, 47,
			118, 169, 104, 157, 113, 65, 7, 248, 186, 194, 119, 130, 62, 109,
			159, 95, 136, 29, 167, 102, 183, 157, 125, 27, 58, 34, 54, 227,
			170, 57, 15, 58, 34, 177, 244, 208, 169, 183, 63, 121, 12, 34,
			240, 190, 229, 240, 51, 40, 20, 10, 170, 167, 216, 17, 150, 127,
			186, 23, 116, 123, 219, 174, 144, 179, 138, 240, 249, 231, 176, 182,
			90, 132, 63, 2, 249, 174, 198, 78, 147, 87, 9, 110, 229, 50,
			84, 208, 94, 159, 157, 114, 41, 18, 191, 172, 149, 229, 229, 137,
			188, 196, 173, 116, 0, 149, 249, 104, 229, 147, 183, 63, 185, 84,
			26, 78, 95, 249, 228, 241, 227, 199, 159, 174, 125, 178, 188, 156,
			126, 255, 199, 180, 195, 34, 10, 135, 97, 240, 50, 145, 178, 254,
			233, 242, 69, 41, 214, 239, 183, 152, 5, 229, 63, 20, 10, 10,
			148, 178, 92, 44, 252, 43, 194, 210, 164, 57, 239, 136, 96, 148,
			131, 112, 37, 114, 22, 39, 228, 200, 0, 40, 158, 11, 128, 199,
			87, 6, 192, 151, 238, 137, 11, 47, 212, 66, 90, 222, 40, 138,
			104, 40, 112, 200, 126, 208, 239, 7, 124, 34, 0, 48, 93, 194,
			64, 246, 194, 207, 224, 234, 9, 223, 19, 230, 240, 179, 113, 175,
			21, 210, 211, 173, 81, 208, 247, 105, 84, 40, 162, 99, 173, 24,
			161, 88, 133, 2, 166, 152, 236, 204, 1, 112, 76, 93, 249, 30,
			132, 2, 61, 143, 71, 42, 215, 99, 183, 37, 2, 69, 235, 24,
			37, 75, 91, 198, 24, 60, 185, 18, 131, 216, 139, 164, 136, 194,
			193, 153, 232, 41, 146, 124, 14, 254, 73, 243, 11, 197, 139, 107,
			179, 75, 69, 117, 140, 70, 161, 40, 51, 160, 220, 218, 239, 187,
			195, 97, 16, 118, 9, 1, 39, 84, 61, 106, 71, 90, 146, 69,
			110, 2, 167, 179, 33, 61, 95, 197, 192, 141, 115, 116, 188, 113,
			33, 240, 85, 146, 193, 127, 199, 68, 28, 171, 178, 160, 141, 181,
			33, 224, 37, 37, 70, 245, 162, 178, 252, 107, 44, 162, 111, 150,
			94, 15, 88, 40, 122, 111, 150, 94, 251, 238, 217, 155, 246, 235,
			30, 27, 69, 111, 54, 94, 15, 130, 240, 205, 198, 107, 78, 189,
			55, 95, 89, 175, 145, 24, 96, 32, 191, 249, 250, 121, 158, 192,
			105, 143, 70, 20, 212, 108, 20, 228, 246, 79, 221, 51, 14, 244,
			37, 18, 11, 158, 214, 253, 14, 27, 69, 224, 7, 221, 64, 112,
			172, 240, 125, 10, 177, 166, 18, 72, 85, 37, 2, 74, 89, 9,
			164, 182, 146, 172, 86, 82, 165, 172, 196, 175, 104, 196, 150, 134,
			174, 239, 171, 173, 145, 56, 101, 137, 52, 234, 122, 61, 75, 158,
			180, 36, 140, 197, 237, 167, 213, 189, 20, 211, 9, 44, 133, 93,
			6, 163, 161, 44, 180, 201, 212, 130, 172, 250, 170, 115, 229, 114,
			94, 83, 44, 17, 169, 159, 13, 149, 100, 165, 41, 255, 60, 15,
			124, 212, 233, 4, 47, 145, 108, 201, 35, 45, 69, 85, 48, 14,
			144, 102, 65, 33, 127, 216, 174, 230, 139, 155, 231, 122, 9, 2,
			20, 209, 95, 142, 130, 136, 250, 22, 84, 64, 29, 233, 168, 96,
			224, 114, 191, 25, 188, 162, 17, 240, 30, 27, 245, 253, 4, 202,
			17, 167, 146, 90, 21, 92, 158, 106, 243, 225, 248, 140, 160, 25,
			69, 92, 128, 16, 119, 120, 161, 136, 249, 213, 197, 80, 66, 32,
			221, 115, 170, 134, 110, 196, 199, 106, 142, 41, 1, 201, 98, 4,
			3, 215, 243, 232, 80, 192, 49, 19, 61, 169, 19, 231, 170, 13,
			113, 226, 3, 127, 203, 14, 112, 67, 96, 157, 14, 167, 170, 222,
			239, 176, 40, 57, 181, 43, 65, 126, 117, 121, 229, 83, 204, 153,
			43, 79, 218, 203, 43, 27, 107, 203, 27, 43, 79, 172, 229, 149,
			231, 249, 56, 186, 57, 200, 118, 154, 116, 135, 46, 23, 4, 228,
			72, 169, 159, 133, 240, 165, 27, 142, 220, 232, 12, 86, 158, 148,
			0, 165, 89, 241, 7, 228, 158, 184, 45, 47, 10, 134, 162, 132,
			212, 239, 28, 217, 113, 1, 139, 70, 114, 150, 38, 121, 18, 178,
			47, 21, 236, 19, 60, 148, 11, 23, 217, 164, 15, 95, 9, 230,
			180, 26, 45, 249, 141, 21, 138, 227, 111, 42, 61, 240, 177, 6,
			236, 85, 208, 239, 187, 242, 227, 162, 225, 210, 97, 171, 236, 51,
			143, 151, 143, 232, 113, 121, 108, 73, 185, 73, 59, 52, 162, 161,
			71, 203, 187, 125, 118, 236, 246, 191, 105, 72, 19, 120, 25, 237,
			41, 79, 40, 249, 154, 164, 167, 146, 78, 146, 104, 74, 242, 51,
			143, 45, 122, 129, 204, 76, 210, 232, 228, 225, 69, 226, 15, 122,
			122, 76, 19, 103, 41, 146, 208, 203, 60, 252, 234, 5, 23, 81,
			71, 206, 156, 112, 136, 121, 220, 26, 170, 188, 134, 174, 172, 150,
			251, 193, 113, 228, 70, 103, 242, 96, 206, 234, 137, 65, 255, 35,
			249, 148, 204, 45, 146, 244, 220, 67, 229, 197, 88, 7, 31, 82,
			15, 30, 46, 62, 91, 90, 28, 44, 45, 250, 237, 197, 189, 141,
			197, 253, 141, 197, 150, 181, 216, 121, 254, 208, 130, 90, 240, 45,
			61, 13, 56, 45, 97, 194, 66, 124, 228, 26, 17, 105, 186, 60,
			26, 238, 81, 248, 146, 249, 174, 12, 213, 135, 28, 190, 122, 225,
			180, 26, 73, 161, 223, 81, 169, 202, 143, 155, 133, 226, 139, 175,
			11, 234, 12, 46, 206, 114, 191, 96, 190, 90, 8, 124, 88, 66,
			171, 202, 238, 48, 144, 235, 145, 244, 74, 119, 202, 202, 214, 242,
			219, 178, 165, 159, 137, 130, 165, 37, 2, 69, 196, 144, 29, 203,
			115, 47, 55, 246, 81, 80, 220, 41, 13, 229, 167, 193, 58, 234,
			224, 219, 85, 31, 89, 242, 129, 113, 149, 144, 83, 232, 229, 145,
			109, 114, 104, 251, 106, 246, 22, 249, 27, 141, 100, 179, 25, 61,
			99, 26, 223, 233, 243, 185, 95, 107, 208, 28, 111, 219, 146, 152,
			103, 29, 25, 234, 18, 93, 30, 132, 222, 36, 231, 32, 151, 147,
			14, 216, 31, 113, 129, 65, 32, 235, 214, 21, 27, 10, 114, 217,
			142, 226, 57, 4, 161, 215, 31, 241, 224, 132, 90, 132, 188, 79,
			166, 208, 186, 172, 153, 253, 78, 127, 117, 155, 92, 83, 205, 41,
			180, 118, 38, 105, 105, 166, 241, 221, 236, 141, 164, 101, 152, 198,
			119, 230, 109, 242, 111, 202, 47, 205, 204, 254, 74, 211, 205, 220,
			191, 106, 80, 103, 225, 82, 72, 187, 174, 8, 78, 232, 249, 189,
			163, 27, 123, 10, 184, 125, 186, 44, 199, 90, 80, 143, 39, 38,
			121, 27, 78, 220, 254, 136, 114, 21, 122, 99, 97, 242, 96, 144,
			139, 160, 223, 135, 158, 123, 66, 33, 156, 212, 41, 69, 199, 19,
			137, 218, 3, 121, 108, 20, 10, 92, 26, 220, 41, 38, 219, 227,
			139, 224, 197, 91, 175, 82, 252, 143, 156, 3, 232, 186, 244, 90,
			203, 154, 83, 191, 210, 244, 239, 230, 99, 192, 180, 41, 233, 247,
			76, 210, 148, 48, 204, 190, 159, 52, 13, 108, 222, 188, 149, 158,
			216, 255, 197, 135, 228, 211, 46, 179, 188, 94, 196, 6, 193, 104,
			32, 67, 183, 63, 242, 130, 114, 215, 147, 145, 91, 78, 46, 37,
			120, 249, 100, 165, 28, 223, 72, 196, 39, 249, 115, 233, 187, 220,
			247, 253, 96, 34, 247, 174, 43, 129, 252, 95, 235, 100, 214, 137,
			101, 153, 215, 137, 30, 248, 242, 240, 127, 174, 169, 7, 190, 153,
			35, 179, 61, 198, 69, 232, 14, 168, 60, 250, 159, 107, 166, 109,
			243, 49, 153, 241, 34, 138, 185, 126, 193, 0, 173, 240, 222, 106,
			238, 226, 77, 128, 149, 22, 156, 102, 50, 20, 37, 246, 131, 142,
			76, 92, 11, 89, 121, 201, 144, 182, 241, 29, 63, 117, 35, 220,
			153, 47, 76, 41, 109, 73, 219, 252, 140, 204, 121, 44, 12, 169,
			135, 250, 166, 223, 169, 111, 60, 216, 92, 32, 51, 114, 187, 53,
			18, 11, 51, 234, 86, 35, 110, 226, 27, 63, 114, 131, 144, 250,
			11, 179, 160, 21, 102, 155, 73, 51, 255, 33, 121, 127, 155, 246,
			169, 160, 77, 117, 149, 116, 17, 152, 252, 103, 132, 236, 82, 113,
			197, 219, 239, 131, 45, 191, 77, 222, 171, 5, 60, 157, 122, 151,
			76, 15, 35, 218, 9, 94, 198, 211, 227, 150, 249, 99, 66, 134,
			110, 151, 126, 35, 216, 183, 52, 140, 133, 204, 97, 79, 27, 59,
			242, 127, 170, 145, 107, 74, 140, 186, 218, 186, 82, 206, 10, 25,
			71, 202, 130, 14, 70, 225, 189, 213, 219, 86, 218, 99, 37, 43,
			223, 28, 143, 50, 31, 144, 27, 33, 125, 41, 190, 153, 208, 111,
			72, 153, 239, 99, 247, 65, 98, 195, 234, 63, 104, 100, 206, 73,
			103, 109, 144, 105, 5, 153, 185, 48, 33, 255, 28, 138, 185, 187,
			111, 173, 152, 188, 235, 50, 87, 136, 177, 75, 133, 121, 103, 98,
			226, 24, 221, 220, 101, 246, 154, 159, 146, 44, 250, 111, 222, 157,
			120, 57, 129, 107, 238, 131, 183, 250, 21, 80, 95, 254, 229, 2,
			153, 54, 179, 217, 204, 145, 70, 126, 171, 201, 59, 163, 108, 198,
			92, 253, 141, 118, 238, 250, 103, 101, 93, 210, 185, 218, 97, 213,
			129, 202, 72, 244, 88, 196, 173, 43, 238, 128, 14, 185, 172, 2,
			241, 73, 251, 248, 198, 36, 224, 208, 101, 39, 52, 10, 145, 234,
			134, 126, 124, 1, 80, 25, 186, 30, 10, 14, 60, 26, 98, 25,
			124, 74, 35, 30, 176, 16, 86, 173, 229, 36, 55, 169, 42, 222,
			97, 163, 208, 79, 238, 35, 106, 78, 213, 174, 183, 108, 232, 4,
			125, 154, 30, 78, 78, 207, 222, 34, 115, 68, 55, 50, 166, 49,
			59, 243, 177, 124, 212, 76, 99, 110, 166, 64, 86, 228, 241, 96,
			246, 90, 230, 161, 150, 91, 132, 10, 236, 86, 237, 244, 126, 19,
			233, 65, 39, 232, 142, 162, 248, 42, 225, 101, 192, 197, 196, 133,
			226, 181, 217, 155, 228, 139, 164, 52, 93, 215, 111, 230, 86, 37,
			16, 206, 182, 242, 145, 194, 211, 125, 229, 106, 42, 239, 212, 229,
			16, 127, 220, 50, 113, 158, 171, 30, 215, 245, 107, 230, 68, 245,
			184, 174, 207, 78, 84, 143, 235, 115, 239, 77, 84, 143, 235, 215,
			111, 144, 165, 184, 120, 24, 183, 244, 187, 57, 144, 154, 147, 239,
			38, 197, 56, 81, 156, 234, 209, 178, 102, 246, 150, 126, 253, 102,
			44, 75, 155, 194, 233, 137, 30, 13, 133, 205, 221, 74, 90, 134,
			105, 220, 154, 191, 67, 62, 151, 122, 116, 211, 152, 215, 11, 185,
			178, 212, 51, 222, 197, 73, 222, 124, 165, 143, 169, 90, 61, 107,
			102, 231, 245, 91, 119, 99, 209, 250, 52, 74, 187, 151, 180, 52,
			211, 152, 255, 209, 71, 73, 203, 48, 141, 249, 7, 15, 137, 39,
			213, 26, 166, 177, 160, 223, 201, 61, 149, 106, 147, 220, 119, 137,
			123, 80, 81, 23, 120, 20, 169, 182, 188, 21, 74, 71, 171, 157,
			97, 106, 94, 192, 193, 151, 223, 217, 216, 58, 35, 107, 102, 23,
			244, 249, 66, 108, 129, 49, 133, 74, 147, 210, 141, 113, 178, 48,
			155, 0, 102, 160, 65, 183, 231, 73, 75, 90, 151, 53, 141, 123,
			250, 221, 220, 206, 37, 224, 83, 104, 197, 217, 88, 222, 142, 203,
			136, 158, 68, 41, 222, 54, 196, 153, 23, 4, 75, 173, 201, 102,
			205, 236, 61, 125, 225, 78, 172, 49, 59, 133, 74, 146, 37, 202,
			106, 166, 113, 47, 93, 162, 172, 97, 26, 247, 230, 239, 144, 29,
			105, 205, 148, 105, 220, 215, 31, 229, 214, 223, 189, 68, 105, 190,
			199, 168, 78, 236, 76, 13, 152, 202, 154, 217, 251, 250, 189, 100,
			177, 166, 166, 81, 110, 178, 88, 83, 154, 105, 220, 255, 209, 98,
			210, 50, 76, 227, 126, 161, 72, 78, 165, 1, 211, 166, 145, 215,
			231, 115, 191, 72, 13, 96, 35, 113, 217, 90, 57, 29, 8, 199,
			122, 113, 159, 4, 61, 12, 154, 212, 170, 227, 179, 148, 42, 179,
			145, 120, 247, 2, 78, 103, 205, 108, 94, 191, 255, 40, 182, 106,
			122, 10, 13, 73, 22, 112, 90, 51, 141, 124, 202, 189, 166, 13,
			211, 200, 155, 183, 201, 43, 105, 241, 140, 105, 60, 208, 111, 231,
			6, 112, 212, 163, 241, 189, 227, 36, 78, 168, 74, 21, 57, 96,
			17, 132, 76, 96, 160, 165, 93, 227, 136, 71, 46, 117, 76, 129,
			187, 29, 218, 63, 75, 172, 131, 136, 118, 221, 200, 239, 83, 46,
			9, 92, 18, 142, 169, 209, 51, 89, 51, 251, 64, 207, 207, 199,
			134, 205, 76, 161, 45, 211, 73, 75, 51, 141, 7, 51, 215, 147,
			150, 97, 26, 15, 110, 153, 196, 34, 24, 0, 217, 71, 153, 37,
			45, 151, 135, 74, 242, 219, 13, 92, 69, 165, 20, 119, 145, 19,
			223, 60, 38, 42, 252, 168, 31, 205, 222, 33, 203, 36, 155, 213,
			48, 81, 149, 244, 155, 185, 143, 46, 36, 170, 212, 149, 84, 82,
			108, 166, 38, 51, 83, 73, 127, 244, 129, 52, 69, 147, 153, 169,
			20, 135, 163, 38, 51, 83, 41, 206, 76, 154, 204, 76, 165, 235,
			55, 200, 99, 162, 103, 117, 51, 187, 156, 249, 68, 203, 21, 206,
			155, 217, 165, 2, 109, 148, 153, 84, 29, 197, 156, 51, 22, 83,
			193, 242, 172, 41, 115, 155, 142, 198, 174, 234, 55, 227, 220, 118,
			185, 177, 93, 42, 98, 75, 117, 105, 233, 170, 190, 172, 0, 213,
			165, 165, 171, 177, 165, 186, 180, 116, 53, 182, 84, 151, 150, 174,
			94, 191, 65, 214, 164, 30, 205, 52, 158, 232, 119, 115, 15, 46,
			253, 140, 175, 210, 134, 153, 244, 137, 190, 122, 51, 150, 136, 153,
			244, 73, 170, 13, 65, 127, 18, 127, 166, 186, 204, 164, 79, 230,
			239, 144, 53, 130, 95, 108, 118, 61, 243, 255, 181, 220, 195, 243,
			184, 200, 203, 247, 183, 80, 225, 49, 44, 152, 131, 214, 103, 111,
			203, 53, 52, 16, 150, 77, 125, 62, 94, 67, 197, 92, 82, 17,
			233, 76, 36, 236, 177, 173, 134, 68, 102, 83, 95, 87, 41, 197,
			144, 200, 108, 198, 182, 26, 18, 153, 205, 185, 27, 73, 203, 48,
			141, 77, 243, 54, 249, 43, 77, 234, 210, 76, 227, 11, 125, 33,
			247, 103, 154, 84, 38, 247, 5, 8, 204, 5, 202, 3, 17, 245,
			168, 188, 73, 151, 59, 191, 73, 10, 97, 97, 205, 247, 147, 165,
			151, 7, 6, 56, 25, 112, 178, 204, 212, 169, 171, 152, 22, 36,
			25, 47, 225, 80, 117, 18, 212, 9, 34, 174, 198, 166, 190, 32,
			238, 95, 232, 155, 243, 177, 189, 136, 251, 23, 169, 47, 136, 251,
			23, 115, 183, 147, 150, 97, 26, 95, 220, 253, 128, 124, 70, 100,
			82, 173, 102, 246, 181, 92, 73, 226, 158, 252, 196, 137, 129, 251,
			214, 42, 92, 4, 31, 83, 110, 117, 118, 94, 130, 159, 69, 240,
			237, 183, 192, 159, 8, 19, 117, 229, 53, 6, 63, 43, 193, 183,
			245, 170, 74, 167, 89, 9, 190, 29, 27, 156, 149, 224, 219, 49,
			248, 89, 9, 190, 109, 222, 38, 43, 82, 149, 102, 26, 123, 122,
			62, 247, 177, 84, 53, 150, 63, 112, 133, 215, 75, 78, 169, 148,
			5, 22, 137, 5, 104, 89, 156, 147, 182, 166, 77, 99, 239, 189,
			219, 73, 11, 229, 205, 255, 56, 105, 25, 166, 177, 7, 63, 33,
			191, 214, 164, 46, 221, 52, 106, 250, 143, 115, 127, 50, 185, 206,
			241, 79, 197, 226, 223, 137, 77, 172, 246, 196, 34, 159, 251, 176,
			229, 15, 51, 38, 22, 151, 92, 190, 186, 66, 30, 118, 34, 76,
			33, 131, 1, 139, 232, 57, 200, 21, 104, 72, 24, 106, 250, 94,
			62, 182, 22, 107, 91, 45, 5, 13, 145, 169, 205, 45, 36, 45,
			195, 52, 106, 247, 126, 68, 62, 37, 250, 116, 198, 204, 30, 100,
			142, 180, 220, 79, 161, 146, 254, 40, 173, 195, 34, 24, 184, 97,
			48, 28, 245, 213, 193, 195, 36, 195, 139, 23, 121, 26, 215, 225,
			96, 246, 22, 89, 39, 217, 105, 73, 231, 154, 250, 231, 185, 18,
			40, 70, 30, 167, 68, 62, 153, 93, 193, 229, 103, 33, 110, 70,
			67, 54, 226, 253, 51, 181, 2, 211, 138, 171, 53, 167, 175, 37,
			45, 221, 52, 154, 239, 231, 146, 150, 97, 26, 205, 197, 13, 140,
			164, 105, 201, 220, 218, 122, 49, 247, 17, 236, 82, 145, 254, 26,
			238, 138, 220, 168, 230, 227, 10, 182, 167, 231, 146, 150, 110, 26,
			109, 50, 159, 180, 12, 211, 104, 127, 248, 144, 148, 165, 108, 221,
			52, 158, 234, 229, 92, 94, 174, 82, 42, 252, 210, 252, 162, 166,
			163, 53, 79, 167, 211, 22, 206, 127, 239, 110, 210, 50, 76, 227,
			233, 79, 150, 146, 157, 248, 127, 7, 0, 0, 255, 255, 142, 238,
			226, 252, 119, 48, 0, 0},
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
	ret, err := discovery.GetDescriptorSet("instances.Instances")
	if err != nil {
		panic(err)
	}
	return ret
}
