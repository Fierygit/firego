'''
Author: Firefly
Date: 2021-04-11 12:42:57
Descripttion:
LastEditTime: 2021-04-11 15:52:07
'''

from PIL import Image
import deal
import data
import numpy as np
import pytesseract


def get_vercode():

    print(Image.open("vcode/" + str(10) + ".png").size)
    img = Image.open("../token.png")
    print(img.size)
    cropped = img.crop((0, 0, 65, 25))  # (left, upper, right, lower)
    cropped.save("./img/cropped.png")
    img = Image.open('./img/cropped.png')

    Img = cropped.convert('L')
    Img.save("./img/grey.png")
    # 自定义灰度界限，大于这个值为黑色，小于这个值为白色
    threshold = 180

    table = []
    for i in range(256):
        if i < threshold:
            table.append(0)
        else:
            table.append(1)

    # 图片二值化
    photo = Img.point(table, '1')
    photo.save("./img/black.png")

    x_len = photo.size[0]
    y_len = photo.size[1]
    print("fdf", x_len)
    # (left, upper, right, lower)
    cropped1 = photo.crop((0, 0, x_len/4, y_len))
    cropped2 = photo.crop((x_len/4, 0, x_len/2, y_len))
    cropped3 = photo.crop((x_len/2, 0, x_len / 4 + x_len / 2, y_len))
    cropped4 = photo.crop((x_len / 4 + x_len / 2, 0, x_len, y_len))

    cropped1.save("./img/cropped1.png")
    cropped2.save("./img/cropped2.png")
    cropped3.save("./img/cropped3.png")
    cropped4.save("./img/cropped4.png")
    text = pytesseract.image_to_string(cropped4)
    print(text)

    image_arr1 = np.array(cropped1, dtype='int').tolist()  # 转化成numpy数组
    image_arr2 = np.array(cropped2, dtype='int').tolist()  # 转化成numpy数组
    image_arr3 = np.array(cropped3, dtype='int').tolist()  # 转化成numpy数组
    image_arr4 = np.array(cropped4, dtype='int').tolist()  # 转化成numpy数组

    image_arr = [image_arr1, image_arr2, image_arr3, image_arr4]
    ans = []
    for i in image_arr:
        ans.append(find_mid(i))

    print("vcode: ", ans)

# for i in image_arr:
#     for j in i:
#         print(j)
#     print()

# # resize image with high-quality
# out = photo.resize((75, 28), Image.ANTIALIAS)
# out.save("./img/black1.png")


def find_mid(arr):
    max_len = 10000000
    max_len_x = 0
    max_len_y = 0
    for i, line in enumerate(arr):
        for j, val in enumerate(line):
            cur = 0
            for ii, line1 in enumerate(arr):
                for jj, val1 in enumerate(line1):
                    if val1 == 0:
                        cur = cur + get_len(i, j, ii, jj)
            if cur <= max_len:
                max_len_x = i
                max_len_y = j
                max_len = cur
    # arr[max_len_x][max_len_y] = 2
    # for i, line in enumerate(arr):
    #     for j, val in enumerate(line):
    #         print(val, end=" ")
    #     print()

    start_x = max_len_x - 7
    start_y = max_len_y - 5
    end_x = max_len_x + 7
    end_y = max_len_y + 4
    # print(max_len_x, max_len_y)
    # print(start_x, start_y, end_x, end_y)
    deald_arr = []
    for i, line in enumerate(arr):
        line1 = []
        for j, val in enumerate(line):
            if i >= start_x and i <= end_x and j >= start_y and j <= end_y:
                line1.append(val)
        if len(line1) != 0:
            deald_arr.append(line1)
            print(line1)
    print()
    x = np.array(deald_arr).flatten()
    res = [0] * 10
    for j in range(10):
        # 一次取字库中标准数据
        y = data.array_map[j]
        # 通过异或计算不同元素的数量
        res[j] = np.sum(x ^ y)
        # 取差异最小的下标

    print(res)
    return str(np.argmin(res))


def get_len(x, y, xx, yy):
    return abs(x - xx) + abs(y - yy)


if __name__ == "__main__":
    get_vercode()
