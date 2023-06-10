package main

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "helloworld/api/helloworld/v1"
	"image"
	_ "image/jpeg"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Person struct {
	Name    string
	Age     string
	Address string
}

func main() {

	// 执行命令获取视频时长
	cmd := exec.Command("ffprobe", "-i", "video.mp4", "-show_entries", "format=duration", "-v", "quiet", "-of", "csv=p=0")

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	// 解析输出，并将其转换为时长
	floatStr := strings.TrimSpace(string(out))
	duration, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Duration: %.2f seconds\n", duration)
}
func mainxx() {
	// 打开要读取的图像文件
	file, err := os.Open("D:\\phpstudy_pro\\WWW\\study\\example-kratos\\client\\cmd\\222.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 获取图像分辨率
	config, _, err := image.DecodeConfig(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(config.Width, config.Height)
}

func mainx() {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("127.0.0.1:9000"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}

	// defer conn.Close()
	cc := v1.NewGreeterClient(conn)
	r, err := cc.SayNihHao(context.Background(), &v1.NiHaoRequest{Name: "test"})
	u := v1.NewUserClient(conn)
	fmt.Println(r)
	s, err := u.GetUser(context.Background(), &v1.GetUserRequest{Id: 1})
	fmt.Println(s, err)
}
