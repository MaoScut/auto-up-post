package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
)

// curl 'https://www.douban.com/group/topic/253898211/add_comment' \
//   -H 'Connection: keep-alive' \
//   -H 'Cache-Control: max-age=0' \
//   -H 'sec-ch-ua: " Not A;Brand";v="99", "Chromium";v="96", "Google Chrome";v="96"' \
//   -H 'sec-ch-ua-mobile: ?0' \
//   -H 'sec-ch-ua-platform: "Windows"' \
//   -H 'Upgrade-Insecure-Requests: 1' \
//   -H 'Origin: https://www.douban.com' \
//   -H 'Content-Type: multipart/form-data; boundary=----WebKitFormBoundary09GVOMf7GoZ56Aky' \
//   -H 'User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36' \
//   -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9' \
//   -H 'Sec-Fetch-Site: same-origin' \
//   -H 'Sec-Fetch-Mode: navigate' \
//   -H 'Sec-Fetch-User: ?1' \
//   -H 'Sec-Fetch-Dest: document' \
//   -H 'Referer: https://www.douban.com/group/topic/253898211/' \
//   -H 'Accept-Language: zh-CN,zh;q=0.9,en;q=0.8,ja;q=0.7' \
//   -H 'Cookie: bid=-9nFIuDkLdw; gr_user_id=bbab060b-95b1-4d5b-9541-3f43a2430c9b; __gads=ID=514249fd5502e601-22e323137cce003b:T=1635774218:RT=1635774218:S=ALNI_MYT6Ke9yyxMZbIgySFjjnS3CO6jVA; ll="118282"; push_noty_num=0; push_doumail_num=0; __utmv=30149280.12528; __yadk_uid=y5S891JHbC8NTVahKu21GEsCzVDlLnpD; douban-fav-remind=1; __utmz=30149280.1637764675.3.3.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; viewed="35231266_3354490_23008813"; _pk_ref.100001.8cb4=%5B%22%22%2C%22%22%2C1637933690%2C%22https%3A%2F%2Fwww.baidu.com%2Flink%3Furl%3DzkxWhBgR1ugcpihpOueUJo6rxzLN3k228DE_yg3fPA_tZkhObM9KT9G5cGy2HCb6%26wd%3D%26eqid%3Db60c6034001c509c00000003619e4e3f%22%5D; _pk_ses.100001.8cb4=*; ap_v=0,6.0; dbcl2="125282692:/+YtWkwJt8o"; ck=CG5e; __utma=30149280.1617365984.1635774221.1637764675.1637933814.4; __utmc=30149280; __utmt=1; _pk_id.100001.8cb4=bc0df6a6c34fff30.1637681490.3.1637933827.1637764682.; __utmb=30149280.17.5.1637933830725' \
//   --data-raw $'------WebKitFormBoundary09GVOMf7GoZ56Aky\r\nContent-Disposition: form-data; name="ck"\r\n\r\nCG5e\r\n------WebKitFormBoundary09GVOMf7GoZ56Aky\r\nContent-Disposition: form-data; name="rv_comment"\r\n\r\n顶\r\n------WebKitFormBoundary09GVOMf7GoZ56Aky\r\nContent-Disposition: form-data; name="img"; filename=""\r\nContent-Type: application/octet-stream\r\n\r\n\r\n------WebKitFormBoundary09GVOMf7GoZ56Aky\r\nContent-Disposition: form-data; name="start"\r\n\r\n0\r\n------WebKitFormBoundary09GVOMf7GoZ56Aky\r\nContent-Disposition: form-data; name="submit_btn"\r\n\r\n发送\r\n------WebKitFormBoundary09GVOMf7GoZ56Aky--\r\n' \
//   --compressed
func main() {

	// 	ck: CG5e
	// rv_comment: 顶
	// img: (binary)
	// start: 0
	// submit_btn: 发送
	body := bytes.Buffer{}
	writer := multipart.NewWriter(&body)
	formData := map[string]string{
		"ck":         "CG5e",
		"rv_comment": "顶",
		"start":      "0",
		"submit_btn": "发送",
	}
	for k, v := range formData {
		fw, err := writer.CreateFormField(k)
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(fw, strings.NewReader(v))
		if err != nil {
			log.Fatal(err)
		}
	}
	writer.Close()
	req, err := http.NewRequest("POST", "https://www.douban.com/group/topic/253898211/add_comment", bytes.NewReader(body.Bytes()))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Add("Origin", "https://www.douban.com")
	req.Header.Add("Referer", "https://www.douban.com/group/topic/253898211/")
	req.Header.Add("Cookie", `bid=-9nFIuDkLdw; gr_user_id=bbab060b-95b1-4d5b-9541-3f43a2430c9b; __gads=ID=514249fd5502e601-22e323137cce003b:T=1635774218:RT=1635774218:S=ALNI_MYT6Ke9yyxMZbIgySFjjnS3CO6jVA; ll="118282"; push_noty_num=0; push_doumail_num=0; __utmv=30149280.12528; __yadk_uid=y5S891JHbC8NTVahKu21GEsCzVDlLnpD; douban-fav-remind=1; __utmz=30149280.1637764675.3.3.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; viewed="35231266_3354490_23008813"; _pk_ref.100001.8cb4=%5B%22%22%2C%22%22%2C1637933690%2C%22https%3A%2F%2Fwww.baidu.com%2Flink%3Furl%3DzkxWhBgR1ugcpihpOueUJo6rxzLN3k228DE_yg3fPA_tZkhObM9KT9G5cGy2HCb6%26wd%3D%26eqid%3Db60c6034001c509c00000003619e4e3f%22%5D; _pk_ses.100001.8cb4=*; ap_v=0,6.0; dbcl2="125282692:/+YtWkwJt8o"; ck=CG5e; __utma=30149280.1617365984.1635774221.1637764675.1637933814.4; __utmc=30149280; __utmt=1; _pk_id.100001.8cb4=bc0df6a6c34fff30.1637681490.3.1637933827.1637764682.; __utmb=30149280.17.5.1637933830725`)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("client.Do failed", err)
	}
	if resp.StatusCode >= 400 {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("read all error", err)
		}
		log.Fatal("status code: ", resp.StatusCode, " body: ", b)
	}
	log.Println("auto up success", " status code: ", resp.StatusCode)
}
