package main

import (
	"fmt"
	"regexp"
	"testing"
	"time"
)

func TestWriter(t *testing.T) {
	c := make(chan string)
	go writer("./test.txt", c)
	c <- "我是\n"
	defer close(c)
	time.Sleep(time.Second)
}

func TestClubNameRe(t *testing.T) {
	clubNameRe = regexp.MustCompile(`<div class="tit"[\s\S]*[>]+([\s\S]+?)</div>[\s]*<ul>`)
	content := `<div class="basic_info">
	<div class="tit">
			<img src="http://static.shougolf.com/upload/logo/20131007/1548160440.jpg" />天津滨海湖高尔夫球会&nbsp;<label class="btag">2014年中国十佳高尔夫球场</label>
	</div>
	<ul>
			<li style="height:33px;"><span>洞    数</span><label class="tag">18H</label></li>
			<li style="height:33px;"><span>球场风格</span><label class="tag">原生态</label></li>
			<li style="height:33px;"><span>球道草种</span><label class="tag">本特草</label></li>
			<li style="height:33px;"><span>果岭草种</span><label class="tag">本特草</label></li>
			<li style="width: 850px;"><span>配套设施</span><label class="tag">更衣室</label><label class="tag">专卖店</label><label class="tag">中餐厅</label></li>
			<li><span>联系电话</span>022-25271199-8016 -1088</li>
			<li><span>球会网址</span><a title="天津滨海湖高尔夫球会" href="http://www.benhaihulake.com" target="_blank">www.benhaihulake.com</a></li>
			<li style="width: 850px;"><span>联系地址</span>天津天津塘沽区塘黄13829号开发西区滨海湖</li>
	</ul>
	<ul id="infom" style="display:none;">
			<li><span>标准杆数</span>72杆</li>
			<li><span>球场面积</span>6000亩</li>
			<li><span>球道长度</span>7666码</li>
			<li><span>开业时间</span>2008-8-1</li>
			<li><span>设 计 师</span>彼得.戴耶</li>
			<li><span>经营主体</span>鸿铭控股集团</li>
			<li><span>传真号码</span>022-25271099</li>
			<li><span>电子邮箱</span>hanyuting1992@yeah.net</li>
	</ul>
	<div id="infomore" class="para_more" style="background-color: #fff;">
			<a href="javascript:;" onclick="zhankaiinfo(this);">展开更多资料</a></div>
			<script type="text/javascript">
					var d = 1;
					var zhankaiinfo = function (obj) {
							if (d == 1) {
									d = 2;
									$('#infom').show();
									$(obj).html('隐藏更多资料');
							} else {
									d = 1;
									$('#infom').hide();
									$(obj).html('展开更多资料');
							}
					}
			</script>
</div>`
	clubName := clubNameRe.FindAllStringSubmatch(content, -1)
	fmt.Println(clubName[0][1])
}
