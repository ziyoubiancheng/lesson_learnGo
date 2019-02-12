package routers

import (
	"github.com/ziyoubiancheng/lesson_learnGo/lesson_3_beego_web/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
