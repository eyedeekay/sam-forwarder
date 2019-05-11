package samtunnelhandler

func DefaultCSS() string {
	return `.server {
    width: 63%;
    max-width: 63%;
    background-color: #9DABD5;
}
.http {
    width: 63%;
    max-width: 63%;
    background-color: #00ffff;
}
.client {
    width: 63%;
    max-width: 63%;
    background-color: #2D4470;
}
.udpserver {
    width: 63%;
    max-width: 63%;
    background-color: #265ea7;
}
.udpclient {
    width: 63%;
    max-width: 63%;
    background-color: #222187;
}
.TunName {
    font-weight: bold;
}
.panel {
    width: 33%;
    max-width: 33%;
}
.prop {

}
body {
    background-color: #9e9e9e;
    color: #070425;
    font-family: "monospace";
}
a {
    color: #080808;
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
}
toggle_visibility_class("prop")
`
}
