'''
Author: Firefly
Date: 2021-04-12 00:28:09
Descripttion: 
LastEditTime: 2021-04-12 00:29:29
'''
# 新建AipOCR
from aip import AipOcr

config = {
    'appId': '23971628',
    'apiKey': 'qTHrcql5df0YrpI3Qa8fSFLK',
    'secretKey': 'P0P6PhGzLiRA5nvi0CgURgjx9GmAiQ7c'
}

client = AipOcr(**config)


def get_file_content(file):
    with open(file, 'rb') as fp:
        return fp.read()


def img_to_str(image_path):
    image = get_file_content(image_path)
    result = client.basicGeneral(image)
    if 'words_result' in result:
        return '\n'.join([w['words'] for w in result['words_result']])


if __name__ == '__main__':
    imagepath = './token.png'
    print(img_to_str(imagepath))
