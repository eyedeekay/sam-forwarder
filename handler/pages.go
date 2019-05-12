package samtunnelhandler

func DefaultCSS() string {
	return `.server {
    width: 63%;
    max-width: 63%;
    min-height: 15%;
    background-color: #9DABD5;
    float:left
}
.http {
    width: 63%;
    max-width: 63%;
    min-height: 15%;
    background-color: #00ffff;
    float:left
}
.client {
    width: 63%;
    max-width: 63%;
    min-height: 15%;
    background-color: #2D4470;
    float:left
}
.udpserver {
    width: 63%;
    max-width: 63%;
    min-height: 15%;
    background-color: #265ea7;
    float:left
}
.udpclient {
    width: 63%;
    max-width: 63%;
    min-height: 15%;
    background-color: #222187;
    float:left
}
.TunName {
    font-weight: bold;
}
.panel {
    width: 33%;
    max-width: 33%;
    float: right;

}
.prop {

}
.global {
    background-color: #00ffff;
}
body {
    background-color: #9e9e9e;
    color: #070425;
    font-family: "monospace";
}
a {
    color: #080808;
}
h1 {
    background-color: #709fa6;
}
`
}

func DefaultJS() string {
	return `function toggle_visibility_id(id) {
   var e = document.getElementById(id);
   if(e.style.display == 'block')
      e.style.display = 'none';
   else
      e.style.display = 'block';
}
function toggle_visibility_class(id) {
   var elist = document.getElementsByClassName(id)
   for (let e of elist) {
       if(e.style.display == 'block')
          e.style.display = 'none';
       else
          e.style.display = 'block';
   }
   var tlist = document.getElementsByClassName("TunName")
   for (let t of tlist) {
       t.style.display = 'block';
   }
   var clist = document.getElementsByClassName("control")
   for (let c of clist) {
       c.style.display = 'block';
   }
   var slist = document.getElementsByClassName("status")
   for (let s of slist) {
       s.style.display = 'block';
   }
}
toggle_visibility_class("prop")
`
}
