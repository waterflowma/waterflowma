package main

import "fmt"

func main() {
    var siteMap map[string]string /*创建集合 */
    siteMap = make(map[string]string)

    /* map 插入 key - value 对,各个国家对应的首都 */
    siteMap [ "Google" ] = "谷歌"
    siteMap [ "Runoob" ] = "菜鸟教程"
    siteMap [ "Baidu" ] = "百度"
    siteMap [ "Wiki" ] = "维基百科"

    /*使用键输出地图值 */ 
    for site := range siteMap {
        fmt.Println(site, "首都是", siteMap [site])
    }

 // test
    if (ok) {
        fmt.Println("Facebook 的 站点是", name)
    } else {
        fmt.Println("Facebook 站点不存在")
    }
}
