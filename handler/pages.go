package samtunnelhandler

func DefaultCSS() string {
	return `.server {
    width: 63%;
    min-height: 15%;
    background-color: #9DABD5;
    float: left;
    overflow-wrap: break-word;
}
.client {
    width: 63%;
    min-height: 15%;
    background-color: #2D4470;
    float: left;
    overflow-wrap: break-word;
}
.tcpclient {
    width: 63%;
    min-height: 15%;
    background-color: #2D4470;
    float: left;
    overflow-wrap: break-word;
}
.http {
    width: 63%;
    min-height: 15%;
    background-color: #00ffff;
    float: left;
    overflow-wrap: break-word;
}
.httpclient {
    width: 63%;
    min-height: 15%;
    background-color: #709fa6;
    float: left;
    overflow-wrap: break-word;
}
.udpserver {
    width: 63%;
    min-height: 15%;
    background-color: #265ea7;
    float: left;
    overflow-wrap: break-word;
}
.udpclient {
    width: 63%;
    min-height: 15%;
    background-color: #222187;
    float: left;
    overflow-wrap: break-word;
}
.vpnserver {
    width: 63%;
    min-height: 15%;
    background-color: #265ea7;
    float: left;
    overflow-wrap: break-word;
}
.vpnclient {
    width: 63%;
    min-height: 15%;
    background-color: #222187;
    float: left;
    overflow-wrap: break-word;
}
.TunName {
    font-weight: bold;
}
.panel {
    width: 33%;
    float: right;
}
.prop {

}
.prop:hover {
    box-shadow: inset 0 0 100px 100px rgba(255, 255, 255, 0.1);
}
.global {
    background-color: #00ffff;
}
body {
    background-color: #9e9e9e;
    color: #070425;
    font-family: "monospace";
    font-size: 1rem;
}
a {
    color: #080808;
}
h1 {
    background-color: #9e9e9e;
}
span {
    float: left;
    display: inline;
}
textarea {
    display: inline-block;
    width: 100%;
    resize: none;
    height: 1rem;
    float: right;
    display: inline;
}
.linkstyle {
    align-items: normal;
    background-color: rgba(0,0,0,0);
    border-color: rgb(0, 0, 238);
    border-style: none;
    box-sizing: content-box;
    color: rgb(0, 0, 238);
    cursor: pointer;
    font: inherit;
    height: auto;
    display: inline;
    padding: 0;
    perspective-origin: 0 0;
    text-align: start;
    text-decoration: underline;
    transform-origin: 0 0;
    width: auto;
    -moz-appearance: none;
    -webkit-logical-height: 1em; /* Chrome ignores auto, so we have to use this hack to set the correct height  */
    -webkit-logical-width: auto; /* Chrome ignores auto, but here for completeness */
}

@supports (-moz-appearance:none) { /* Mozilla-only */
    .linkstyle::-moz-focus-inner { /* reset any predefined properties */
        border: none;
        padding: 0;
    }
    .linkstyle:focus { /* add outline to focus pseudo-class */
        outline-style: dotted;
        outline-width: 1px;
    }
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
