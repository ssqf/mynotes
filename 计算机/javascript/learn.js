//alert("我的第一个 JavaScript");

var e = document.getElementsByTagName("div")[0]
    //alert(e)
document.write("<p> hello wrold </p>" + e.textContent)
    //e.innerHTML = "新修改的内容"
var d = document.getElementById("d")
    //alert(d)

d.textContent = "属性中获取数据：" + d.getAttribute("dat")
d.style.background = "red"
d.onclick = function() { alert("我被点击了") }