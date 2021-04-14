'''
Author: Firefly
Date: 2021-04-10 19:04:05
Descripttion:
LastEditTime: 2021-04-12 00:38:11
'''

import requests
import json
import ocr
import time
import random
url = "https://fangkong.hnu.edu.cn/api/v1/account/login"


headers = {
    "Content-Type": "application/json; charset=UTF-8",
    "Referer": "https://fangkong.hnu.edu.cn/app/",
    "User-Agent": "Mozilla/5.0 (Linux; Android 9; MI 6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.117 Mobile Safari/537.36",
    "Host": "fangkong.hnu.edu.cn",
    "Origin": "https://fangkong.hnu.edu.cn",
}


def get_imgcode():
    getimgvcode_url = "https://fangkong.hnu.edu.cn/api/v1/account/getimgvcode"
    token_res = requests.get(getimgvcode_url, headers=headers).text
    print(token_res)
    data = json.loads(token_res)
    print("get img token: ", data["data"]["Token"])
    return data["data"]["Token"]


def get_img():
    token = get_imgcode()
    IMAGE_URL = "https://fangkong.hnu.edu.cn/imagevcode?token=" + token
    r = requests.get(IMAGE_URL)
    with open('./token.png', 'wb') as f:
        f.write(r.content)
        f.flush()
        f.close()
    return token


info = [["201708010417", "43252219Nft"], [
    "201708010420", "Wgy19981031"], ["201708010419", "531109985@QQ.com"]]


def login(name, pwd):
    print(name,pwd)

    cnt = 6
    while cnt:
        token = get_img()
        a = ocr.img_to_str("./token.png")
        print(a)
        body = {"Code": name, "Password": pwd, "WechatUserinfoCode": "",
                "VerCode": a, "Token": token}
        data = json.dumps(body)
        response = requests.post(url, data=data, headers=headers)
        print(response)

        cookies = response.cookies

        response = response.text
        response = json.loads(response)

        response1 = json.dumps(response, sort_keys=True,
                              indent=4, separators=(',', ':'), ensure_ascii=False)
        print(response1)

        if response["code"] == 404 and response["msg"] == "账号或密码错误":
            break
        if response["code"] == 0:
            print("success")
            break
        print("验证码错误， 尝试 6 次")
        time.sleep(1)
        cnt = cnt - 1

    return getCookies(cookies)


def getCookies(cookie_jar):
    cookie_dict = cookie_jar.get_dict()
    found = ['%s=%s' % (name, value) for (name, value) in cookie_dict.items()]
    return ';'.join(found)


def get_info(cookies):
    url = "https://fangkong.hnu.edu.cn/api/v1/clockinlog/isclockinloginfo"
    header = headers
    header["Cookie"] = cookies
    res = requests.get(url, headers=header)
    res = json.loads(res.text)
    res = json.dumps(res, sort_keys=True,
                     indent=4, separators=(',', ':'), ensure_ascii=False)
    # print(res)

def get_temp():
    return "36." + str(random.randint(4,8))

def add(cookies):
    body = {"Longitude": "", "Latitude": "",
            "RealProvince": "湖南省",
            "RealCity": "长沙市",
            "RealCounty": "岳麓区",
            "RealAddress": "天马",
            "BackState": 1, "MorningTemp": get_temp(),
            "NightTemp": get_temp(), "tripinfolist": []}
    data = json.dumps(body)

    url = "https://fangkong.hnu.edu.cn/api/v1/clockinlog/add"

    cookies1 = "pgv_pvi=9454653440; CCKF_visitor_id_144692=1306555197; UM_distinctid=17867137c0f7f-0a005df8c224c8-5771031-1fa400-17867137c10135; Hm_lvt_d7e34467518a35dd690511f2596a570e=1618053935; TOKEN=fbfa169c2e42411ba1eca6c22a8c9950; .ASPXAUTH=BECA8770683629B1D8439BDA63F6D7056458F5E1C5C9718917D51F2F2C6380F891DB2F4C98D462D89E39D759BF9F617471811037411CB010CD386238967AD9081CCB999FA99151B6E77DF7266ED828E8114B8A09CDF966C9911994E2522EEDD27AEA183B09AE870AE0A387952D3C73D28CD81FFBC4CF21F7512A0777FEC928050BCCDB6A7C3B4867701A1CDF72019B76FBE6667A8081FE81400E96ACFE36A8624BE064C973E5DCD66C6173F69D92BD491E20EE5993E2871B0FE36E8AD728E0CA2734DE355FBEABDF103D8C4C8EA11DAE50EA28B134CB8B2854C66B88FC966E39DA59F8E4CEB1B83A3C07DEAD65E02753DBF9C85DADE5752C217D77B1F110C8ACD682477A4A5FDD6F7B8B86E142B966732CA0C68658A6A2C596E55FBE801CD66CE3E2C418FF75CDE26F3097FFEC663C637A84D3EEFA1D153735C8D6FD3D1254AED3B6D1F76E9CD68DD2993485160595B7307E263627FFBAAFB218CEC4982CC7F7301B434398C6A3372C66CF794DE01C4365C9AAEE4340C6353FEA78ACB6BB69336307E1008E816350F3C56D9ECF3015F348F1A808DEA44479A06087533FBBD3BD5CA5E8EC1FE37A275042818023B4F3CE37B16B9651DC548FA8F65DB6C8795652F16B75A4909EC68379A63AF37A2F15A77FB02808EC20D84E68F844F799A1F39D6710E73B7E565F767AD8E18E41580904D739A35BE0D92F62094BF262AC69C9A222C760FBC8D4A826739C198B645FE54B9021B136D2016B239E1B2E9732896C60CB47AC0D948A1A509076592644DBA7CEA4E91A08386256035CE02944F0FA44925A142A404C3F92A158ACDFC98C38BE22A78D308CA0229CE7713BDE949DD86848221B6BF072D90F7D61193BA66BAFDFA5C140F089383BCE2A775E5DEF1EB8F70877D4AFDDAD79341BD092D70F895508E1018CC3C615BF6A39B43CDB4CDAB2F1739DCA47CEC749E63B9B851614DAC6895D79402C1219161F67A24F4DB6C0964024058A4270B9242335AB85A1816DC49C95B7A76D6618E138BB545563A0CCB0B41FC91AEACE3FC98BE9335946BF04A3751DDCEDDE77AAEA237F3DEE99464F020B1A81C00C9211FE5E0F4F827080781F8C951343A25CB3B879C3B59ECEEE46C023A156C57767407502C8966856472CFE02EC15F54C6ADBA6EEA17829B795FA3A9AD943759271182C186F878A68D8D138144B9CF9FCF5C02B0A8D36A9CCB0B31335C36618DE8C7628C29BDAEF0EED81F26D19CB7E1DCFC81B84C4429975C9A51BC20F0129D30A7A367003D1F585865EE467E14C3743C034E9CE2F7E20A5606F8C1F16AC29FE79A31410B9564DF6BB0620CC6E67202D61395D3C5BEF550D1E61E5238A4C8F311F3242BEA6BEEDB41AA70DBD1C9438355A499DD5728869016F3B16DD7E95CA2E6CDB1B21C1DF05EA22A9F3B66C113BF401F25E22B504EB8E73B711BA5C27F85D0540925A3121D72A280C6D9FAA7C5ECCE3A197F4CC46A24CD1EF3CA2FEFB3A2DA3015A9B2944ED9F099A5F7208531234A81FE7706A19F36E6AD8B3F978A76CD12D0E7FB2A727459970583DB3C0A955870985B7A734D9857E0DD616852405CE15C475D2F2F6672DFA25E1A731F922AEF54D200A5D610D42074EC308C91E4EFC9F150BE8BCD4123A4F198A6D00B5E4F93C78BCCA49F832A7BE9C985535D1162799EDEC48D7F657F73BB60E8E6CC03D849022B2D58F6D9EAC6E8F55A35F90A7A78692533B6546DBC490CB8578714B3BE4AB8158FEB81C63834D6ACDDE641BBB8EDF3820A4BFB2CB0EF12BC34F3197393CD73F9E76978ABFC8DBA6CEA4B60730B39ED16118D4EDABE1F8324E5C09E45771A1A5750B99BD5C722E85A9A1CB65ECFE9BC79401A477307D708BC68F53BAFBCF318BAAFA135A47B273410BCE9B66; Hm_lpvt_d7e34467518a35dd690511f2596a570e=1618070578"
    header = headers
    header["Cookie"] = cookies
    header["Accept-Language"] = "zh-CN,zh;q=0.9"
    header["Accept-Encoding"] = "gzip, deflate, br"

    res = requests.post(url, data=data,  headers=header)
    print(res)
    print("--------------")
    print(res.text)
    res = json.loads(res.text)
    res = json.dumps(res, sort_keys=True,
                     indent=4, separators=(',', ':'), ensure_ascii=False)

    print(res)


def main():

    print("start deal")

    for i in info:
        cookies = login(i[0], i[1])
        get_info(cookies)
        add(cookies)


if __name__ == "__main__":
    main()
