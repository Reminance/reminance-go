/**
 * @File : 10_httpDemo
 * @Author: Octal_H
 * @Date : 2019/11/12
 * @Desc :
 */

package main

import "net/http"

func main() {
	/*
		HTTP 文件服务器是常见的 Web 服务之一。开发阶段为了测试，需要自行安装 Apache 或 Nginx 服务器，
		下载安装配置需要大量的时间。使用Go语言实现一个简单的 HTTP 服务器只需要几行代码，如下所示。
	*/
	//使用 http.FileServer 文件服务器将当前目录作为根目录（/目录）的处理器，访问根目录，就会进入当前目录。
	http.Handle("/", http.FileServer(http.Dir(".")))
	//默认的 HTTP 服务侦听在本机 8080 端口
	http.ListenAndServe(":8080", nil)
}
