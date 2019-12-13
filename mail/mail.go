package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Luxurioust/excelize"
	"gopkg.in/gomail.v2"
)

//go get gopkg.in/gomail.v2
//go get github.com/Luxurioust/excelize

type mailMessage struct {
	name    string
	mailTo  string
	subject string
	body    string
}

//SendMail 定义邮箱服务器连接信息
func SendMail(mailTo []string, subject string, body string) error {
	//定义邮箱服务器连接信息，如果是网易邮箱 pass填密码，qq邮箱填授权码

	// mailConn := map[string]string{
	//  "user": "reminance@163.com",
	//  "pass": "hycfhggjhf",
	//  "host": "smtp.163.com",
	//  "port": "465",
	// }

	mailConn := map[string]string{
		"user": "365204662@qq.com",
		"pass": "sddsjcnjqqnlbgcb",
		// "host": "smtp.exmail.qq.com",   //企业邮箱用这个
		"host": "smtp.qq.com",
		"port": "465",
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()

	// m.SetHeader("From", m.FormatAddress(mailConn["user"], "XX官方")) //这种方式可以添加别名，即“XX官方”
	// m.SetHeader("From", "FB Sample"+"<"+mailConn["user"]+">")
	//说明：如果是用网易邮箱账号发送，以下方法别名可以是中文，如果是qq企业邮箱，以下方法用中文别名，会报错，需要用上面此方法转码
	//m.SetHeader("From", "FB Sample"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“FB Sample”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("From", mailConn["user"])
	m.SetHeader("To", mailTo...)    //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err

}

func sendMailFromStruct(messageArr []mailMessage) {
	var size = len(messageArr)
	for i, message := range messageArr {
		fmt.Printf("当前进度[%d/%d], 开始发送给'%s(%s)'\n", i+1, size, message.name, message.mailTo)
		//定义收件人
		mailTo := []string{
			message.mailTo,
		}
		//邮件主题
		subject := message.subject
		//邮件正文
		body := message.body

		err := SendMail(mailTo, subject, body)
		if err != nil {
			log.Println(err)
			fmt.Println("send fail")
			return
		}
		fmt.Printf("发送给'%s(%s)'成功\n", message.name, message.mailTo)
	}
}

func readExcel(file string) (messageArr []mailMessage) {
	mailMessageArr := []mailMessage{}
	xlsx, err := excelize.OpenFile(file)
	if err != nil {
		fmt.Printf("没有找到'%s'excel文件", file)
		os.Exit(1)
	}
	// // Get value from cell by given sheet index and axis.
	// cell, err := xlsx.GetCellValue("Sheet1", "B2")
	// fmt.Println(cell)
	// Get sheet index.
	// index := xlsx.GetSheetIndex("Sheet1")
	// // Get all the rows in a sheet.
	// rows, err := xlsx.GetRows("Sheet" + strconv.Itoa(index))
	rows, err := xlsx.GetRows("Sheet1")
	for i, row := range rows {
		if i == 0 || row[0] == "" {
			continue
		}
		mailMessageArr = append(mailMessageArr, mailMessage{row[0], row[1], row[2], row[3]})
	}
	return mailMessageArr
}

func main() {
	var mailType string
	var user string
	var pass string
	var host string
	var port string
	var filePath string
	flag.StringVar(&mailType, "qq", "qq", "邮箱类型[QQ邮箱填:qq   163邮箱填:netease]\n")
	flag.StringVar(&user, "user", "365204662@qq.com", "邮箱账号[xxx@qq.com或xxx@163.com]\n")
	flag.StringVar(&pass, "pass", "sddsjcnjqqnlbgcb", "邮箱授权码[邮箱设置中开通的授权码]\n")
	flag.StringVar(&host, "host", "smtp.qq.com", "邮箱服务域名[不填的话采用自动适配 只支持QQ邮箱/网易邮箱]\n")
	flag.StringVar(&port, "port", "465", "邮箱服务端口[不填的话采用自动适配]\n")
	flag.StringVar(&filePath, "filePath", "./Workbook.xlsx", "要导入的excel文件路径[可以填全路径或相对路径 \n 如:'D:/file/Workbook.xlsx'  或:'./Workbook.xlsx']\n")

	flag.Parse()

	fmt.Println("邮箱类型:", mailType)
	fmt.Println("邮箱账号:", user)
	fmt.Println("邮箱授权码:", pass)
	fmt.Println("邮箱服务域名:", host)
	fmt.Println("邮箱服务端口:", port)
	fmt.Println("excel文件路径:", filePath)
	fmt.Println("开始发送邮件!")
	var messageArr = readExcel(filePath)
	sendMailFromStruct(messageArr)
	// mailType := flag.String("port", ":8080", "http listen port")
}
