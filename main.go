package main

import "hpv-vaccine-hunter/api"

func main() {
	tk := "wxapptoken:10:5085e1c5427d63045683a42b2e9a948e_a8a0bd6ce6326f71ce782575557e2f71"
	cookie := "_xxhm_=%7B%22id%22%3A23105906%2C%22mobile%22%3A%2217602875725%22%2C%22nickName%22%3A%22%E6%BC%82%E6%B5%AE%E7%BE%A4%E5%B2%9B%F0%9F%8E%88%22%2C%22headerImg%22%3A%22https%3A%2F%2Fthirdwx.qlogo.cn%2Fmmopen%2Fvi_32%2FczD1u6Naufp60MWql9FicPTYQf3xdZqpaLA9EHMSNatANXXCcubvg0OrGFGaFtAmsqFKlCibialiauYBvhNG8xIZPQ%2F132%22%2C%22regionCode%22%3A%22510109%22%2C%22name%22%3A%22%E8%B5%B5*%E6%A2%85%22%2C%22uFrom%22%3A%22depa_vacc_detail%22%2C%22wxSubscribed%22%3A1%2C%22birthday%22%3A%221997-04-29+02%3A00%3A00%22%2C%22sex%22%3A2%2C%22hasPassword%22%3Atrue%2C%22birthdayStr%22%3A%221997-04-29%22%7D; _xzkj_=wxapptoken:10:5085e1c5427d63045683a42b2e9a948e_a8a0bd6ce6326f71ce782575557e2f71; 5377=39accd9e9ed63bc6f3; 5425=42eda48916f889154d; 8681=0a0dff2cb76fdb04f3; tgw_l7_route=6e0a47ce8062c68ea282b9bbb140678e"
	client := api.NewClient(tk, cookie)

	client.GetSecondKillList()
}
